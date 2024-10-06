package repo

import (
	"github.com/AmitSuresh/shipping/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	some1 := &gorm.Config{TranslateError: true}
	opts := []gorm.Option{}
	opts = append(opts, some1)

	db, err := gorm.Open(postgres.Open(cfg.DbDSN), opts...)
	return db, err
}
