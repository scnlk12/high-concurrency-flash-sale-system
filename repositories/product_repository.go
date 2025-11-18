package repositories

import (
	"database/sql"

	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
)

// 开发接口
type IProduct interface {
	// 连接数据库
	Conn() error
	// 增删改查
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

// 实现接口
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

// 数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// 插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "INSERT product SET productName=?, productNum=?, productImg=?, productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if err != nil {
		return 0, err
	}

	// 获取productId
	return res.LastInsertId()
}
