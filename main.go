package main

import "github.com/kataras/iris"

func main() {
	// 创建实例对象
	app := iris.New()

	// 错误等级
	app.Logger().SetLevel("debug")

	// 注册模板目录
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// 注册控制器
	

	app.Run(
		iris.Addr("localhost:8080"),
	)
}