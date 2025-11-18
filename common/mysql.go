package common

import (
	"database/sql"
	"log"
)

// 创建mysql连接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/product?charset=utf8")
	return
}

// 获取返回值
func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]string)
	for rows.Next() {
		// 将行数据保存到record字典
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				record[columns[i]] = string(v.([]byte))
			} else {
				record[columns[i]] = ""
			}
		}
	}
	return record
}

// 获取所有返回值
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	// 返回所有列
	columns, _ := rows.Columns()
	// 这里表示一行所有列的值 用[]byte表示
	vals := make([][]byte, len(columns))
	// 一行填充数据
	scans := make([]interface{}, len(columns))
	// scans引用vals 把数据填充到[]bytes中
	for k := range vals {
		scans[k] = &vals[k]
	}

	result := make(map[int]map[string]string)
	i := 0

	for rows.Next() {
		// 扫描一行数据
		err := rows.Scan(scans...)
		if err != nil {
			log.Printf("扫描行数据失败: %v\n", err)
			continue
		}
		row := make(map[string]string)
		for j, val := range vals {
			key := columns[j]
			if val != nil {
				row[key] = string(val)
			} else {
				row[key] = ""
			}
		}

		result[i] = row
		i++
	}

	return result
}