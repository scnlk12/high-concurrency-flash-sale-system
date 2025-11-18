package repositories

import (
	"database/sql"
	"strconv"

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
	SelectByKey(int64) (*datamodels.Product, error)
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

// 删除
func (p *ProductManager) Delete (productId int64) bool {
	// 判断连接是否成功
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "DELETE from product where productId=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productId)

	if(err != nil) {
		return false
	}
	return true
}

// 更新
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	// 判断连接是否成功
	if err = p.Conn(); err != nil {
		return
	}

	sql := "update product set productName=?, productNum=?, productImg=?, productUrl=? where productId=" + strconv.FormatInt(product.ProductId, 10)

	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if err != nil {
		return
	}

	return nil
}

// 根据商品id查询商品
func (p *ProductManager) SelectByKey(productId int64) (product *datamodels.Product, err error) {
	// 判断连接是否存在
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	sql := "select * from" + p.table + "where productId = " + strconv.FormatInt(productId, 10)

	row, err := p.mysqlConn.Query(sql)
	defer row.Close()

	if err != nil {
		return &datamodels.Product{}, err
	}

	// 获取结果
	res := common.GetResultRow(row)

	if len(res) == 0 {
		return &datamodels.Product{}, err
	}

	// map[string]string -> Struct Product
	common.DataToStructByTagSql(res, product)

	return
}

// 获取所有商品
func (p *ProductManager) SelectAll() (productArr []*datamodels.Product, err error)  {
	// 判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from " + p.table

	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArr = append(productArr, product)
	}

	return 
}