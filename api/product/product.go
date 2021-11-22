package product

import (
	"magento2go/api/errs"
	"magento2go/client"
	"magento2go/client/catalog_product_repository_v1"
	"magento2go/models"
	"time"
)

type ProductApi struct {
	client  *client.MagentoCommunity
	timeout time.Duration
}

// NewProductApi
func NewProductApi(client *client.MagentoCommunity, timeout time.Duration) *ProductApi {
	return &ProductApi{
		client,
		timeout,
	}
}

// GetProductBySku
func (a *ProductApi) GetProductBySku(sku string) (*models.CatalogDataProductInterface, error) {
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
func (a *ProductApi) GetAllProducts(currentPage, pageSize int64) ([]*models.CatalogDataProductInterface, error) {
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
