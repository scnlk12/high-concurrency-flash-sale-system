package services

import (
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/repositories"
)

type IProductService interface {
	GetProductById(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductById(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

// 初始化函数
func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{productRepository: repository}
}

func (ps *ProductService) GetProductById(productId int64) (*datamodels.Product, error) {
	return ps.productRepository.SelectByKey(productId)
}

func (ps *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return ps.productRepository.SelectAll()
}

func (ps *ProductService) DeleteProductById(productId int64) bool {
	return ps.productRepository.Delete(productId)
}

func (ps *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return ps.productRepository.Insert(product)
}

func (ps *ProductService) UpdateProduct(product *datamodels.Product) error {
	return ps.productRepository.Update(product)
}