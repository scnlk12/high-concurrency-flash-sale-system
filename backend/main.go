package main

import (
	"context"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/scnlk12/high-concurrency-flash-sale-system/backend/web/controllers"
	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/repositories"
	"github.com/scnlk12/high-concurrency-flash-sale-system/services"
)

func main() {
	// 创建实例对象
	app := iris.New()

	// 错误等级 设置错误模式 在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 注册模板目录
	template := iris.HTML("./backend/web/views", ".html").Layout(
		"shared/layout.html").Reload(true)

	app.RegisterView(template)

	// 设置模板目录
	app.StaticWeb("/assets", "./backend/web/assets")

	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))
	
	// 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		// iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}