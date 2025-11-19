package services

import (
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/repositories"
)

type IOrderService interface {
	GetOrderById(int64) (*datamodels.Order, error)
	DeleteOrderById(int64) bool
	UpdateOrder(*datamodels.Order) error
	InsertOrder(*datamodels.Order) (int64, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}

// 实现接口
type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func NewOrderService (repository repositories.IOrderRepository) IOrderService {
	return &OrderService{
		OrderRepository: repository,
	}
}

func (o *OrderService) GetOrderById(orderId int64) (*datamodels.Order, error) {
	return o.OrderRepository.SelectByKey(orderId)
}

func (o *OrderService) DeleteOrderById(orderId int64) bool {
	return o.OrderRepository.Delete(orderId)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo() (map[int]map[string]string, error) {
	return o.OrderRepository.SelectAllWithInfo()
}