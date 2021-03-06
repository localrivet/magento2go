// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogInventoryDataStockItemInterface Interface StockItem
//
// swagger:model catalog-inventory-data-stock-item-interface
type CatalogInventoryDataStockItemInterface struct {

	// Backorders status
	// Required: true
	Backorders *int64 `json:"backorders"`

	// Whether Quantity Increments is enabled
	// Required: true
	EnableQtyIncrements *bool `json:"enable_qty_increments"`

	// extension attributes
	ExtensionAttributes CatalogInventoryDataStockItemExtensionInterface `json:"extension_attributes,omitempty"`

	// is decimal divided
	// Required: true
	IsDecimalDivided *bool `json:"is_decimal_divided"`

	// Stock Availability
	// Required: true
	IsInStock *bool `json:"is_in_stock"`

	// is qty decimal
	// Required: true
	IsQtyDecimal *bool `json:"is_qty_decimal"`

	// item id
	ItemID int64 `json:"item_id,omitempty"`

	// low stock date
	// Required: true
	LowStockDate *string `json:"low_stock_date"`

	// Can Manage Stock
	// Required: true
	ManageStock *bool `json:"manage_stock"`

	// Maximum Qty Allowed in Shopping Cart data wrapper
	// Required: true
	MaxSaleQty *float64 `json:"max_sale_qty"`

	// Minimal quantity available for item status in stock
	// Required: true
	MinQty *float64 `json:"min_qty"`

	// Minimum Qty Allowed in Shopping Cart or NULL when there is no limitation
	// Required: true
	MinSaleQty *float64 `json:"min_sale_qty"`

	// Notify for Quantity Below data wrapper
	// Required: true
	NotifyStockQty *float64 `json:"notify_stock_qty"`

	// product id
	ProductID int64 `json:"product_id,omitempty"`

	// qty
	// Required: true
	Qty *float64 `json:"qty"`

	// Quantity Increments data wrapper
	// Required: true
	QtyIncrements *float64 `json:"qty_increments"`

	// show default notification message
	// Required: true
	ShowDefaultNotificationMessage *bool `json:"show_default_notification_message"`

	// Stock identifier
	StockID int64 `json:"stock_id,omitempty"`

	// stock status changed auto
	// Required: true
	StockStatusChangedAuto *int64 `json:"stock_status_changed_auto"`

	// use config backorders
	// Required: true
	UseConfigBackorders *bool `json:"use_config_backorders"`

	// use config enable qty inc
	// Required: true
	UseConfigEnableQtyInc *bool `json:"use_config_enable_qty_inc"`

	// use config manage stock
	// Required: true
	UseConfigManageStock *bool `json:"use_config_manage_stock"`

	// use config max sale qty
	// Required: true
	UseConfigMaxSaleQty *bool `json:"use_config_max_sale_qty"`

	// use config min qty
	// Required: true
	UseConfigMinQty *bool `json:"use_config_min_qty"`

	// use config min sale qty
	// Required: true
	UseConfigMinSaleQty *int64 `json:"use_config_min_sale_qty"`

	// use config notify stock qty
	// Required: true
	UseConfigNotifyStockQty *bool `json:"use_config_notify_stock_qty"`

	// use config qty increments
	// Required: true
	UseConfigQtyIncrements *bool `json:"use_config_qty_increments"`
}

// Validate validates this catalog inventory data stock item interface
func (m *CatalogInventoryDataStockItemInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackorders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableQtyIncrements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDecimalDivided(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsInStock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsQtyDecimal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLowStockDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManageStock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSaleQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinSaleQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyStockQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQtyIncrements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowDefaultNotificationMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStockStatusChangedAuto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigBackorders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigEnableQtyInc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigManageStock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigMaxSaleQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigMinQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigMinSaleQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigNotifyStockQty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseConfigQtyIncrements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateBackorders(formats strfmt.Registry) error {

	if err := validate.Required("backorders", "body", m.Backorders); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateEnableQtyIncrements(formats strfmt.Registry) error {

	if err := validate.Required("enable_qty_increments", "body", m.EnableQtyIncrements); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateIsDecimalDivided(formats strfmt.Registry) error {

	if err := validate.Required("is_decimal_divided", "body", m.IsDecimalDivided); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateIsInStock(formats strfmt.Registry) error {

	if err := validate.Required("is_in_stock", "body", m.IsInStock); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateIsQtyDecimal(formats strfmt.Registry) error {

	if err := validate.Required("is_qty_decimal", "body", m.IsQtyDecimal); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateLowStockDate(formats strfmt.Registry) error {

	if err := validate.Required("low_stock_date", "body", m.LowStockDate); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateManageStock(formats strfmt.Registry) error {

	if err := validate.Required("manage_stock", "body", m.ManageStock); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateMaxSaleQty(formats strfmt.Registry) error {

	if err := validate.Required("max_sale_qty", "body", m.MaxSaleQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateMinQty(formats strfmt.Registry) error {

	if err := validate.Required("min_qty", "body", m.MinQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateMinSaleQty(formats strfmt.Registry) error {

	if err := validate.Required("min_sale_qty", "body", m.MinSaleQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateNotifyStockQty(formats strfmt.Registry) error {

	if err := validate.Required("notify_stock_qty", "body", m.NotifyStockQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateQty(formats strfmt.Registry) error {

	if err := validate.Required("qty", "body", m.Qty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateQtyIncrements(formats strfmt.Registry) error {

	if err := validate.Required("qty_increments", "body", m.QtyIncrements); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateShowDefaultNotificationMessage(formats strfmt.Registry) error {

	if err := validate.Required("show_default_notification_message", "body", m.ShowDefaultNotificationMessage); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateStockStatusChangedAuto(formats strfmt.Registry) error {

	if err := validate.Required("stock_status_changed_auto", "body", m.StockStatusChangedAuto); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigBackorders(formats strfmt.Registry) error {

	if err := validate.Required("use_config_backorders", "body", m.UseConfigBackorders); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigEnableQtyInc(formats strfmt.Registry) error {

	if err := validate.Required("use_config_enable_qty_inc", "body", m.UseConfigEnableQtyInc); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigManageStock(formats strfmt.Registry) error {

	if err := validate.Required("use_config_manage_stock", "body", m.UseConfigManageStock); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigMaxSaleQty(formats strfmt.Registry) error {

	if err := validate.Required("use_config_max_sale_qty", "body", m.UseConfigMaxSaleQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigMinQty(formats strfmt.Registry) error {

	if err := validate.Required("use_config_min_qty", "body", m.UseConfigMinQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigMinSaleQty(formats strfmt.Registry) error {

	if err := validate.Required("use_config_min_sale_qty", "body", m.UseConfigMinSaleQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigNotifyStockQty(formats strfmt.Registry) error {

	if err := validate.Required("use_config_notify_stock_qty", "body", m.UseConfigNotifyStockQty); err != nil {
		return err
	}

	return nil
}

func (m *CatalogInventoryDataStockItemInterface) validateUseConfigQtyIncrements(formats strfmt.Registry) error {

	if err := validate.Required("use_config_qty_increments", "body", m.UseConfigQtyIncrements); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this catalog inventory data stock item interface based on context it is used
func (m *CatalogInventoryDataStockItemInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogInventoryDataStockItemInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogInventoryDataStockItemInterface) UnmarshalBinary(b []byte) error {
	var res CatalogInventoryDataStockItemInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
