package repositories

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
)

type IUserRepository interface {
	Conn() error
	Select(userName string) (*datamodels.User, error)
	Insert(user *datamodels.User) (int64, error)
}

type UserManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

func NewUserRepository(table string, conn *sql.DB) IUserRepository {
	return &UserManagerRepository{
		table: table,
		mysqlConn: conn,
	}	
}

// 实现接口
func (u *UserManagerRepository) Conn() error {
	if u.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}

	if u.table == "" {
		u.table = "user"
	}

	return nil
}

// 查询
func (u *UserManagerRepository) Select(userName string) (*datamodels.User, error) {
	if userName == "" {
		return &datamodels.User{}, errors.New("查询条件不能为空!")
	}
	
	// 判断mysql连接
	if err := u.Conn(); err != nil {
		return nil, err
	}

	// sql
	sql := "select * from " + u.table + " where userName =?"
	rows, err := u.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在!")
	}

	user := &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return user, nil
}

// 插入
func (u *UserManagerRepository) Insert(user *datamodels.User) (int64, error) {
	// 判断mysql连接
	if err := u.Conn(); err != nil {
		return 0, err
	}
	// sql
	sql := "insert " + u.table + " set nickName=?, userName=?,password=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 根据userId查询用户信息
func (u *UserManagerRepository) SelectById(userId int64) (*datamodels.User, error) {
	sql := "select * from " + u.table + " where userId=" + strconv.FormatInt(userId, 10)
	row, err := u.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	
	res := common.GetResultRow(row)
	if len(res) == 0 {
		return nil, errors.New("用户不存在!")
	}

	user := &datamodels.User{}
	common.DataToStructByTagSql(res, user)
	return user, err
}