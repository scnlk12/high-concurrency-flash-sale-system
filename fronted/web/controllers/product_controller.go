package controllers

import (
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/services"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.ProductService
	OrderService   services.OrderService
	Session        *sessions.Session
}

var (
	htmlOutPath = "./fronted/web/htmlProductShow/" // 生成的html保存目录
	templatePath = "./fronted/web/views/template/" // 静态文件模板目录
)

// 生成html静态文件控制器
func (p *ProductController) GetGenerateHtml() {
	// 获取productId
	productString := p.Ctx.URLParam("productId")
	productId, err := strconv.Atoi(productString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	// 1. 获取模板文件地址
	contentTmp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	// 2. 获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	// 3. 获取模板渲染数据
	product, err := p.ProductService.GetProductById(int64(productId))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	// 4. 生成静态文件
	generateStaticHtml(p.Ctx, contentTmp, fileName, product)
}

// 生成html静态文件
func generateStaticHtml(ctx iris.Context, template *template.Template,
fileName string, product *datamodels.Product) {
	// 判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}
	}

	// 生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

// 判断文件是否存在
func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 获取订单详情页面
func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductById(1)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	productString := p.Ctx.URLParam("productId")
	userId := p.Ctx.GetCookie("uid")
	productId, err := strconv.Atoi(productString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductById(int64(productId))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	var orderId int64
	showMessage := "抢购失败!"
	// 判断商品数量是否满足要求
	if product.ProductNum > 0 {
		// 扣除商品数量
		product.ProductNum -= 1
		// 在高并发大流量下会出现超卖情况
		err := p.ProductService.UpdateProduct(product)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		}

		// 创建订单
		userId, err := strconv.Atoi(userId)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		}
		order := &datamodels.Order{
			UserId: int64(userId),
			ProductId: int64(productId),
			OrderStatus: datamodels.OrderSuccess,
		}
		orderId, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		} else {
			showMessage = "抢购成功!"
		}
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderId": orderId,
			"showMessage": showMessage,
		},
	}
}
