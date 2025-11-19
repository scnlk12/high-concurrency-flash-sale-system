package controllers

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/scnlk12/high-concurrency-flash-sale-system/common"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/services"
)

type OrderController struct {
	Ctx iris.Context
	OrderService services.IOrderService
}

// 查询订单页面
func (o *OrderController) Get() mvc.View {
	orderArr, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败!")
	}

	return mvc.View {
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderArr,
		},
	}
}

// 修改订单
func (o *OrderController) PostUpdate()  {
	order := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	// 表单对应标签为 gf_test 与数据库对应字段做映射
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "gf_test"})
	if err := dec.Decode(o.Ctx.Request().Form, order); err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	err := o.OrderService.UpdateOrder(order)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order/all")
}

// 请求订单新增页面
func (o *OrderController) GetAdd() mvc.View {
	return mvc.View{
		Name: "order/add.html",
	}
}

// 新增订单逻辑
func (o *OrderController) PostAdd()  {
	order := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "gf_test"})
	if err := dec.Decode(o.Ctx.Request().Form, order); err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}

	_, err := o.OrderService.InsertOrder(order)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order/all")
}

// 修改订单页面
func (o *OrderController) GetManager() mvc.View {
	// 从路由中获取当前访问订单的orderId
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	order, err := o.OrderService.GetOrderById(id)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name: "order/manager.html",
		Data: iris.Map{
			"order": order,
		},
	}
}

// 删除订单
func (o *OrderController) GetDelete() {
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	isOk := o.OrderService.DeleteOrderById(id)
	if isOk {
		o.Ctx.Application().Logger().Debug("删除订单成功, id为" + idString)
	} else {
		o.Ctx.Application().Logger().Debug("删除订单失败, id为" + idString)
	}
	o.Ctx.Redirect("/order/all")
}