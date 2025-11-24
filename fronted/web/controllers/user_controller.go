package controllers

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/services"
)

// 登录 注册
type UserController struct {
	Ctx iris.Context
	Service services.IUserService
	Session *sessions.Session
}

// 展示注册页面
func (c *UserController) GetRegister() mvc.View {
	return mvc.View {
		Name: "user/register.html",
	}
}

// 处理注册操作
func (c *UserController) PostRegister() {
	// 获取少量字段时可用
	// 字段较多时 使用结构体标签进行映射
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	// ozzo-validation 表单校验

	user := &datamodels.User{
		UserName: userName,
		NickName: nickName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)

	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	
	c.Ctx.Redirect("/user/login")
	return 
}

// 用户登录

// 请求登录页面
func (u *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

// 登录操作
func (u *UserController) PostLogin() mvc.Response {
	var (
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
	)

	user, isOk := u.Service.IsPwdSuccess(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}
	common.GlobalCookie(u.Ctx, "uid", strconv.FormatInt(user.Id, 10))
	u.Session.Set("userId", strconv.FormatInt(user.Id, 10))
	return mvc.Response {
		Path: "/product/",
	}
}
