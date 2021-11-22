package product

import (
	"time"

	"github.com/localrivet/magento2go/api/errs"
	"github.com/localrivet/magento2go/client"
	"github.com/localrivet/magento2go/client/catalog_product_repository_v1"
	"github.com/localrivet/magento2go/models"
)

// NewProductApi
func NewProductApi(mc *client.MagentoCommunity, timeout time.Duration) ProductApi {
	return &productApi{
		client:  mc,
		timeout: timeout,
	}
}

type ProductApi interface {
	GetProductBySku(sku string) (*models.CatalogDataProductInterface, error)
	GetAllProducts(currentPage, pageSize int64) ([]*models.CatalogDataProductInterface, error)
}

type productApi struct {
	client  *client.MagentoCommunity
	timeout time.Duration
}

// GetProductBySku
func (a *productApi) GetProductBySku(sku string) (*models.CatalogDataProductInterface, error) {
	params := catalog_product_repository_v1.NewCatalogProductRepositoryV1GetGetParams()
	params.Sku = sku
	params.SetTimeout(time.Second * 300)
	result, err := a.client.CatalogProductRepositoryV1.CatalogProductRepositoryV1GetGet(params)
	if err != nil {
		return nil, errs.InterceptError(err)
	}

	return result.Payload, err
}

// GetAllProducts
func (a *productApi) GetAllProducts(currentPage, pageSize int64) ([]*models.CatalogDataProductInterface, error) {
	params := catalog_product_repository_v1.NewCatalogProductRepositoryV1GetListGetParams()
	params.SearchCriteriaPageSize = &pageSize
	params.SearchCriteriaCurrentPage = &currentPage
	params.SetTimeout(time.Second * 300)
	result, err := a.client.CatalogProductRepositoryV1.CatalogProductRepositoryV1GetListGet(params)
	if err != nil {
		return nil, errs.InterceptError(err)
	}
	return result.Payload.Items, nil
}
