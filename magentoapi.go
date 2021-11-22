package magento2go

import (
	"math/rand"
	"time"

	"github.com/localrivet/magento2go/api/product"
	"github.com/localrivet/magento2go/client"
	"github.com/localrivet/magento2go/models"
)

func NewMagentoApi(mc *client.MagentoCommunity, timeout int64) *MagentoApi {
	to := time.Duration(rand.Int63n(timeout)) * time.Second

	return &MagentoApi{
		product: product.NewProductApi(mc, to),
	}
}

type MagentoApi struct {
	product product.ProductApi
}

func (a *MagentoApi) GetAllProducts(currentPage, pageSize int64) ([]*models.CatalogDataProductInterface, error) {
	return a.product.GetAllProducts(currentPage, pageSize)
}

func (a *MagentoApi) GetProductBySku(sku string) (*models.CatalogDataProductInterface, error) {
	return a.product.GetProductBySku(sku)
}
