package odoo

import (
	"fmt"
)

// ProductAttributeCustomValue represents product.attribute.custom.value model.
type ProductAttributeCustomValue struct {
	LastUpdate                            *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate                            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One `xmlrpc:"create_uid,omitempty"`
	CustomProductTemplateAttributeValueId *Many2One `xmlrpc:"custom_product_template_attribute_value_id,omitempty"`
	CustomValue                           *String   `xmlrpc:"custom_value,omitempty"`
	DisplayName                           *String   `xmlrpc:"display_name,omitempty"`
	Id                                    *Int      `xmlrpc:"id,omitempty"`
	Name                                  *String   `xmlrpc:"name,omitempty"`
	SaleOrderLineId                       *Many2One `xmlrpc:"sale_order_line_id,omitempty"`
	WriteDate                             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductAttributeCustomValues represents array of product.attribute.custom.value model.
type ProductAttributeCustomValues []ProductAttributeCustomValue

// ProductAttributeCustomValueModel is the odoo model name.
const ProductAttributeCustomValueModel = "product.attribute.custom.value"

// Many2One convert ProductAttributeCustomValue to *Many2One.
func (pacv *ProductAttributeCustomValue) Many2One() *Many2One {
	return NewMany2One(pacv.Id.Get(), "")
}

// CreateProductAttributeCustomValue creates a new product.attribute.custom.value model and returns its id.
func (c *Client) CreateProductAttributeCustomValue(pacv *ProductAttributeCustomValue) (int64, error) {
	return c.Create(ProductAttributeCustomValueModel, pacv)
}

// UpdateProductAttributeCustomValue updates an existing product.attribute.custom.value record.
func (c *Client) UpdateProductAttributeCustomValue(pacv *ProductAttributeCustomValue) error {
	return c.UpdateProductAttributeCustomValues([]int64{pacv.Id.Get()}, pacv)
}

// UpdateProductAttributeCustomValues updates existing product.attribute.custom.value records.
// All records (represented by IDs) will be updated by pacv values.
func (c *Client) UpdateProductAttributeCustomValues(ids []int64, pacv *ProductAttributeCustomValue) error {
	return c.Update(ProductAttributeCustomValueModel, ids, pacv)
}

// DeleteProductAttributeCustomValue deletes an existing product.attribute.custom.value record.
func (c *Client) DeleteProductAttributeCustomValue(id int64) error {
	return c.DeleteProductAttributeCustomValues([]int64{id})
}

// DeleteProductAttributeCustomValues deletes existing product.attribute.custom.value records.
func (c *Client) DeleteProductAttributeCustomValues(ids []int64) error {
	return c.Delete(ProductAttributeCustomValueModel, ids)
}

// GetProductAttributeCustomValue gets product.attribute.custom.value existing record.
func (c *Client) GetProductAttributeCustomValue(id int64) (*ProductAttributeCustomValue, error) {
	pacvs, err := c.GetProductAttributeCustomValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if pacvs != nil && len(*pacvs) > 0 {
		return &((*pacvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.attribute.custom.value not found", id)
}

// GetProductAttributeCustomValues gets product.attribute.custom.value existing records.
func (c *Client) GetProductAttributeCustomValues(ids []int64) (*ProductAttributeCustomValues, error) {
	pacvs := &ProductAttributeCustomValues{}
	if err := c.Read(ProductAttributeCustomValueModel, ids, nil, pacvs); err != nil {
		return nil, err
	}
	return pacvs, nil
}

// FindProductAttributeCustomValue finds product.attribute.custom.value record by querying it with criteria.
func (c *Client) FindProductAttributeCustomValue(criteria *Criteria) (*ProductAttributeCustomValue, error) {
	pacvs := &ProductAttributeCustomValues{}
	if err := c.SearchRead(ProductAttributeCustomValueModel, criteria, NewOptions().Limit(1), pacvs); err != nil {
		return nil, err
	}
	if pacvs != nil && len(*pacvs) > 0 {
		return &((*pacvs)[0]), nil
	}
	return nil, fmt.Errorf("product.attribute.custom.value was not found")
}

// FindProductAttributeCustomValues finds product.attribute.custom.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeCustomValues(criteria *Criteria, options *Options) (*ProductAttributeCustomValues, error) {
	pacvs := &ProductAttributeCustomValues{}
	if err := c.SearchRead(ProductAttributeCustomValueModel, criteria, options, pacvs); err != nil {
		return nil, err
	}
	return pacvs, nil
}

// FindProductAttributeCustomValueIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeCustomValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductAttributeCustomValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductAttributeCustomValueId finds record id by querying it with criteria.
func (c *Client) FindProductAttributeCustomValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributeCustomValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.attribute.custom.value was not found")
}
