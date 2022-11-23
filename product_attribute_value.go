package odoo

import (
	"fmt"
)

// ProductAttributeValue represents product.attribute.value model.
type ProductAttributeValue struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AttributeId         *Many2One  `xmlrpc:"attribute_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	DisplayType         *Selection `xmlrpc:"display_type,omitempty"`
	HtmlColor           *String    `xmlrpc:"html_color,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IsCustom            *Bool      `xmlrpc:"is_custom,omitempty"`
	IsUsedOnProducts    *Bool      `xmlrpc:"is_used_on_products,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	PavAttributeLineIds *Relation  `xmlrpc:"pav_attribute_line_ids,omitempty"`
	Sequence            *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductAttributeValues represents array of product.attribute.value model.
type ProductAttributeValues []ProductAttributeValue

// ProductAttributeValueModel is the odoo model name.
const ProductAttributeValueModel = "product.attribute.value"

// Many2One convert ProductAttributeValue to *Many2One.
func (pav *ProductAttributeValue) Many2One() *Many2One {
	return NewMany2One(pav.Id.Get(), "")
}

// CreateProductAttributeValue creates a new product.attribute.value model and returns its id.
func (c *Client) CreateProductAttributeValue(pav *ProductAttributeValue) (int64, error) {
	return c.Create(ProductAttributeValueModel, pav)
}

// UpdateProductAttributeValue updates an existing product.attribute.value record.
func (c *Client) UpdateProductAttributeValue(pav *ProductAttributeValue) error {
	return c.UpdateProductAttributeValues([]int64{pav.Id.Get()}, pav)
}

// UpdateProductAttributeValues updates existing product.attribute.value records.
// All records (represented by IDs) will be updated by pav values.
func (c *Client) UpdateProductAttributeValues(ids []int64, pav *ProductAttributeValue) error {
	return c.Update(ProductAttributeValueModel, ids, pav)
}

// DeleteProductAttributeValue deletes an existing product.attribute.value record.
func (c *Client) DeleteProductAttributeValue(id int64) error {
	return c.DeleteProductAttributeValues([]int64{id})
}

// DeleteProductAttributeValues deletes existing product.attribute.value records.
func (c *Client) DeleteProductAttributeValues(ids []int64) error {
	return c.Delete(ProductAttributeValueModel, ids)
}

// GetProductAttributeValue gets product.attribute.value existing record.
func (c *Client) GetProductAttributeValue(id int64) (*ProductAttributeValue, error) {
	pavs, err := c.GetProductAttributeValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if pavs != nil && len(*pavs) > 0 {
		return &((*pavs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.attribute.value not found", id)
}

// GetProductAttributeValues gets product.attribute.value existing records.
func (c *Client) GetProductAttributeValues(ids []int64) (*ProductAttributeValues, error) {
	pavs := &ProductAttributeValues{}
	if err := c.Read(ProductAttributeValueModel, ids, nil, pavs); err != nil {
		return nil, err
	}
	return pavs, nil
}

// FindProductAttributeValue finds product.attribute.value record by querying it with criteria.
func (c *Client) FindProductAttributeValue(criteria *Criteria) (*ProductAttributeValue, error) {
	pavs := &ProductAttributeValues{}
	if err := c.SearchRead(ProductAttributeValueModel, criteria, NewOptions().Limit(1), pavs); err != nil {
		return nil, err
	}
	if pavs != nil && len(*pavs) > 0 {
		return &((*pavs)[0]), nil
	}
	return nil, fmt.Errorf("product.attribute.value was not found")
}

// FindProductAttributeValues finds product.attribute.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeValues(criteria *Criteria, options *Options) (*ProductAttributeValues, error) {
	pavs := &ProductAttributeValues{}
	if err := c.SearchRead(ProductAttributeValueModel, criteria, options, pavs); err != nil {
		return nil, err
	}
	return pavs, nil
}

// FindProductAttributeValueIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductAttributeValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductAttributeValueId finds record id by querying it with criteria.
func (c *Client) FindProductAttributeValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributeValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.attribute.value was not found")
}
