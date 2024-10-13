package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	order_service "github.com/AmitSuresh/orderapi/src/infra/proto/order_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateOrder_Success(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Use the mock generated in the 'controller' package
	mockClient := NewMockOrderServiceClient(ctrl)

	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the CreateOrder call
	mockClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(&order_service.OrderResp{Id: "12345"}, nil)

	// Register the route
	CreateOrder(router, mockClient)

	// Prepare a valid request payload
	body := `{"CustomerName":"John Doe", "Quantity":2, "Sku":"SKU123"}`

	// Perform the request
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Contains(t, w.Body.String(), "Customer of id 12345 is created")
}

func TestCreateOrder_InvalidRequest(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockOrderServiceClient(ctrl)
	mockClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.InvalidArgument, "Internal server error"))
	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the route
	CreateOrder(router, mockClient)

	// Perform the request with an invalid payload
	body := `{"CustomerName": "asdf", "Quantity": 1, "Sku": "afs-ass-sda"}` // Invalid input
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error_message"`)
}

func TestCreateOrder_AlreadyExists(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockOrderServiceClient(ctrl)

	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the CreateOrder call to return AlreadyExists error
	mockClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.AlreadyExists, "Order already exists"))

	// Register the route
	CreateOrder(router, mockClient)

	// Prepare a valid request payload
	body := `{"CustomerName":"John Doe", "Quantity":2, "Sku":"SKU123"}`
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Order already exists")
}

func TestCreateOrder_InvalidStruct(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockOrderServiceClient(ctrl)

	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the route
	CreateOrder(router, mockClient)

	// Perform the request with an invalid payload
	body := `{"CustomerName": 11, "Quantity": "sdf", "Sku": "afs-ass-sda"}` // Invalid input
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// assert.Contains(t, w.Body.String(), `"error_message":"Invalid input"`)
	// assert.Contains(t, w.Body.String(), `cannot unmarshal`)
}

func TestCreateOrder_Unknown(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockOrderServiceClient(ctrl)
	mockClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.Unknown, "Something unknown"))
	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the route
	CreateOrder(router, mockClient)

	// Perform the request with an invalid payload
	body := `{"CustomerName": "11", "Quantity": 1, "Sku": "afs-ass-sda"}` // Invalid input
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// assert.Contains(t, w.Body.String(), `"error_message":"Invalid input"`)
	// assert.Contains(t, w.Body.String(), `cannot unmarshal`)
}

func TestCreateOrder_OtherError(t *testing.T) {
	// Set up the mock controller and client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockOrderServiceClient(ctrl)
	mockClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("some internal error"))
	// Create the gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the route
	CreateOrder(router, mockClient)

	// Perform the request with an invalid payload
	body := `{"CustomerName": "11", "Quantity": 1, "Sku": "afs-ass-sda"}` // Invalid input
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// assert.Contains(t, w.Body.String(), `"error_message":"Invalid input"`)
	// assert.Contains(t, w.Body.String(), `cannot unmarshal`)
}
