package fronted

import (
	"context"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/fronted/web/controllers"
	"github.com/scnlk12/high-concurrency-flash-sale-system/repositories"
	"github.com/scnlk12/high-concurrency-flash-sale-system/services"
)

func main() {

	// 1. 创建iris实例
	app := iris.New()

	// 2. 设置错误模式 在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 3. 注册模板
	template := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	// 4. 设置模板目标
	app.StaticWeb("/assets", "./fronted/web/public/assets")
	// 5. 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {

	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建session
	sess := sessions.New(sessions.Config{
		Cookie: "helloworld",
		Expires: 60 * time.Minute,
	})

	// 注册控制器
	user := repositories.NewUserRepository("user", db)
	userSerevice := services.NewUserService(user)
	userPro := mvc.New(app.Party("/user"))
	userPro.Register(userSerevice, ctx, sess.Start)
	userPro.Handle(new(controllers.UserController))

	app.Run(
		iris.Addr("0.0.0.0:8082"),
		// iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
