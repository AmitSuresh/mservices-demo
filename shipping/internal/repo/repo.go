package repo

import (
	"sync"
	"time"

	"github.com/AmitSuresh/shipping/pkg/config"
	"github.com/AmitSuresh/shipping/pkg/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type repoHandler struct {
	db *gorm.DB
	l  *zap.Logger
}

type RepoHandler interface {
	QueryOrders(n_limit int) ([]*OrderShipping, error)
	QueryAllOrders(n_limit int) ([]*Order, error)
}

func NewRepo(l *zap.Logger, cfg *config.Config) RepoHandler {
	db, err := repo.NewDB(cfg)
	if err != nil {
		l.Fatal(err.Error())
		return nil
	}
	db.AutoMigrate(&OrderShipping{}, &Order{})
	return &repoHandler{
		db: db,
		l:  l,
	}
}

func (repo *repoHandler) MigrateAll(db *gorm.DB, i ...interface{}) {
	db.AutoMigrate(i...)
}

func (repo *repoHandler) QueryOrders(n_limit int) ([]*OrderShipping, error) {
	tx := repo.db.Begin()

	if tx.Error != nil {
		repo.l.Error("tx.Error is", zap.Error(tx.Error))
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			repo.l.Error("panicked during tx", zap.Any("r", r))
			tx.Rollback()
		}
	}()

	currTime := time.Now()
	eta := currTime.AddDate(0, 0, 30).Add(6 * time.Hour).Format(time.RFC3339)
	//repo.l.Info("in", zap.Any("eta", eta))
	errChan := make(chan error, 1)
	doneChan := make(chan struct{})

	ordersToShip := &OrderShippings{
		mu:            &sync.Mutex{},
		shippedOrders: []*OrderShipping{},
	}

	go func(eta *string, err chan error, sss *OrderShippings) {
		defer close(errChan)

		orders := make([]Order, 0, n_limit)
		errFromTx := tx.Model(&Order{}).
			Where("Shipped = ?", "false").
			Limit(n_limit).
			FindInBatches(&orders, 100, func(tx *gorm.DB, batch int) error {

				for _, order := range orders {
					shippedOrder := &OrderShipping{
						OrderId: order.Id,
						Order:   order,
						Shipped: true,
						ETA:     *eta,
					}
					sss.mu.Lock()
					sss.shippedOrders = append(sss.shippedOrders, shippedOrder)
					repo.l.Info("in", zap.Any("shippedOrder", shippedOrder))
					sss.mu.Unlock()
				}

				return nil
			}).Error

		if errFromTx != nil {
			repo.l.Error("error while d.Error", zap.Error(errFromTx))
			errChan <- errFromTx
			return
		}
		close(doneChan)
	}(&eta, errChan, ordersToShip)

	select {
	case err := <-errChan:
		if err != nil {
			repo.l.Error("error while select 1", zap.Error(err))
			tx.Rollback()
			return nil, err
		}
	case <-doneChan:
		if len(ordersToShip.shippedOrders) > 0 {
			repo.l.Info("")
			if err := tx.CreateInBatches(ordersToShip.shippedOrders, 200).Error; err != nil {
				repo.l.Error("error while CreateInBatches", zap.Error(err))
				return nil, err
			}
		}
	}
	repo.l.Info("in", zap.Any("ordersToShip.shippedOrders", ordersToShip.shippedOrders))

	if err := tx.Commit().Error; err != nil {
		repo.l.Error("error while committing", zap.Error(err))
		return nil, err
	}

	return ordersToShip.shippedOrders, nil
}

func (repo *repoHandler) QueryAllOrders(n_limit int) ([]*Order, error) {
	tx := repo.db.Begin()

	if tx.Error != nil {
		repo.l.Error("tx.Error is", zap.Error(tx.Error))
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			repo.l.Error("panicked during tx", zap.Any("r", r))
			tx.Rollback()
		}
	}()

	var orders []*Order
	err := tx.Limit(n_limit).Find(&orders).Error
	if err != nil {
		tx.Rollback()
		repo.l.Error("Failed to query orders", zap.Error(err))
		return nil, err
	}

	if err := tx.Rollback().Error; err != nil {
		repo.l.Error("Failed to rollback transaction", zap.Error(err))
		return nil, err
	}

	return orders, nil
}
