package datamodels

type Order struct {
	Id int64 `sql:"orderId"`
	UserId int64 `sql:"userId"`
	// TODO 商品id为int64 订单与商品为1：1关系 待扩展
	ProductId int64 `sql:"productId"`
	OrderStatus int64 `sql: orderStatus`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)