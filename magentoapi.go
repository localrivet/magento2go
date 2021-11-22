package magento2go

import (
	"magento2go/api/product"
	"magento2go/client"
	"magento2go/models"
	"math/rand"
	"time"
)

func NewMagentoApi(client *client.MagentoCommunity, timeout int64) *magentoApi {
	to := time.Duration(rand.Int63n(timeout)) * time.Second
	return &magentoApi{
		client,
		product.NewProductApi(client, to),
	}
}

type magentoApi struct {
	client  *client.MagentoCommunity
	Product *product.ProductApi
}

func (a *magentoApi) GetAllProducts(currentPage, pageSize int64) ([]*models.CatalogDataProductInterface, error) {
	return a.Product.GetAllProducts(currentPage, pageSize)
}

func (a *magentoApi) GetProductBySku(sku string) (*models.CatalogDataProductInterface, error) {
	return a.Product.GetProductBySku(sku)
}
