package errs

import (
	"errors"
	"magento2go/client/catalog_product_repository_v1"
)

func InterceptError(err error) error {
	toInterface := func(t interface{}) interface{} {
		return t
	}
	response := toInterface(err)

	switch toInterface(response).(type) {
	case *catalog_product_repository_v1.CatalogProductRepositoryV1GetGetDefault:
		if i, ok := response.(*catalog_product_repository_v1.CatalogProductRepositoryV1GetGetDefault); ok {
			err = errors.New(*i.Payload.Message)
		}
	case *catalog_product_repository_v1.CatalogProductRepositoryV1GetListGetDefault:
		if i, ok := response.(*catalog_product_repository_v1.CatalogProductRepositoryV1GetListGetDefault); ok {
			err = errors.New(*i.Payload.Message)
		}
	}

	return err
}
