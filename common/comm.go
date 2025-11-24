package common

import (
	"reflect"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

// 根据结构体中的sql标签映射数据到结构体中并转换类型
func DataToStructByTagSql(data map[string]string, obj interface{}) {
	objVal := reflect.ValueOf(obj).Elem()
	for i := 0; i < objVal.NumField(); i++ {
		// 获取sql对应的值
		value := data[objVal.Type().Field(i).Tag.Get("sql")]
		// 获取对应字段的名称
		name := objVal.Type().Field(i).Name
		// 获取对应字段类型
		structFieldType := objVal.Field(i).Type()
		// 获取变量类型 也可以直接写 string类型
		val := reflect.ValueOf(value)
		var err error
		if structFieldType != val.Type() {
			// 类型转换
			val, err = TypeConversion(value, structFieldType.Name())
			if err != nil {
				
			}
		}
		// 设置类型值
		objVal.FieldByName(name).Set(val)
	}

}

// 类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error)  {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		
	}
}


func GlobalCookie(ctx iris.Context, key string, value string) {

}