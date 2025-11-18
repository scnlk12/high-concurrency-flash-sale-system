package datamodels

// 创建类 商品模型
type Product struct {
	// 结构体标签
	ProductId   int64    `json:"ProductId" sql:"productId" gf_test:"ProductId"`
	ProductName string `json:"ProductName" sql:"productName" gf_test:"ProductName"`
	ProductNum  int    `json:"ProductNum" sql:"productNum" gf_test:"ProductNum"`
	ProductImg  string `json:"ProductImg" sql:"productImg" gf_test:"ProductImg"`
	ProductUrl  string `json:"ProductUrl" sql:"productUrl" gf_test:"ProductUrl"`
}
