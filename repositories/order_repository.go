package repositories

import (
	"database/sql"
	"strconv"

	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
)

type IOrderRepository interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

// 创建类
func NewOrderManagerRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table: table,
		mysqlConn: db,
	}
}

// 实现接口中定义的方法
func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

// 插入
func (o *OrderManagerRepository) Insert(order *datamodels.Order) (int64, error) {
	// 1. 判断连接是否有问题
	if err := o.Conn(); err != nil {
		return 0, err
	}

	sql := "insert into " + o.table + "set userId=?, productId=?, orderStatus=?"

	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// 删除
func (o *OrderManagerRepository) Delete(orderId int64) bool {
	// 判断mysql连接
	if err := o.Conn(); err != nil {
		return false
	}
	// sql
	sql := "delete from " + o.table + " where orderId=?"
	stmt, err := o.mysqlConn.Prepare(sql)

	if err != nil {
		return false
	}

	_, err = stmt.Exec(orderId)
	if err != nil {
		return false
	}
	return true
}

// 更新
func (o *OrderManagerRepository) Update(order *datamodels.Order) error {
	// 判断mysql连接
	if err := o.Conn(); err != nil {
		return err
	}

	// sql
	sql := "update " + o.table + " set userId=?, productId=?, orderStatus=? where orderId=" + strconv.FormatInt(order.Id, 10)

	stmt, err := o.mysqlConn.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return err
	}
	return nil
}

// 查询
func (o *OrderManagerRepository) SelectByKey(orderId int64) (*datamodels.Order, error) {
	// mysql连接
	if err := o.Conn(); err != nil {
		return &datamodels.Order{}, err
	}
	// sql
	sql := "select * from " + o.table + " where orderId=" + strconv.FormatInt(orderId, 10)

	row, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}

	res := common.GetResultRow(row)
	if len(res) == 0 {
		return nil, nil
	}
	order := &datamodels.Order{}
	common.DataToStructByTagSql(res, order)
	return order, nil
}

// 查询所有
func (o *OrderManagerRepository) SelectAll() ([]*datamodels.Order, error) {
	// mysql连接
	if err := o.Conn(); err != nil {
		return nil, err
	}

	// sql
	sql := "select * from " + o.table

	// Query
	row, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	res := common.GetResultRows(row)
	if len(res) == 0 {
		return nil, nil
	}
	var orderArr []*datamodels.Order
	for _, v := range res {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArr = append(orderArr, order)
	}

	return orderArr, nil
}

// SelectAllWithInfo
func (o *OrderManagerRepository) SelectAllWithInfo() (map[int]map[string]string, error) {
	// mysql
	if err := o.Conn(); err != nil {
		return nil, err
	}

	// sql
	sql := "select o.orderId, p.productName, o.orderStatus from order as o left join product as p on o.productId = p.productId"

	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	res := common.GetResultRows(rows)
	return res, nil
}