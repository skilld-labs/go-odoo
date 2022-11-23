package odoo

import (
	"fmt"
)

// ProductAttribute represents product.attribute model.
type ProductAttribute struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	AttributeLineIds *Relation `xmlrpc:"attribute_line_ids,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	CreateVariant    *Bool     `xmlrpc:"create_variant,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	ValueIds         *Relation `xmlrpc:"value_ids,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductAttributes represents array of product.attribute model.
type ProductAttributes []ProductAttribute

// ProductAttributeModel is the odoo model name.
const ProductAttributeModel = "product.attribute"

// Many2One convert ProductAttribute to *Many2One.
func (pa *ProductAttribute) Many2One() *Many2One {
	return NewMany2One(pa.Id.Get(), "")
}

// CreateProductAttribute creates a new product.attribute model and returns its id.
func (c *Client) CreateProductAttribute(pa *ProductAttribute) (int64, error) {
	return c.Create(ProductAttributeModel, pa)
}

// UpdateProductAttribute updates an existing product.attribute record.
func (c *Client) UpdateProductAttribute(pa *ProductAttribute) error {
	return c.UpdateProductAttributes([]int64{pa.Id.Get()}, pa)
}

// UpdateProductAttributes updates existing product.attribute records.
// All records (represented by ids) will be updated by pa values.
func (c *Client) UpdateProductAttributes(ids []int64, pa *ProductAttribute) error {
	return c.Update(ProductAttributeModel, ids, pa)
}

// DeleteProductAttribute deletes an existing product.attribute record.
func (c *Client) DeleteProductAttribute(id int64) error {
	return c.DeleteProductAttributes([]int64{id})
}

// DeleteProductAttributes deletes existing product.attribute records.
func (c *Client) DeleteProductAttributes(ids []int64) error {
	return c.Delete(ProductAttributeModel, ids)
}

// GetProductAttribute gets product.attribute existing record.
func (c *Client) GetProductAttribute(id int64) (*ProductAttribute, error) {
	pas, err := c.GetProductAttributes([]int64{id})
	if err != nil {
		return nil, err
	}
	if pas != nil && len(*pas) > 0 {
		return &((*pas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.attribute not found", id)
}

// GetProductAttributes gets product.attribute existing records.
func (c *Client) GetProductAttributes(ids []int64) (*ProductAttributes, error) {
	pas := &ProductAttributes{}
	if err := c.Read(ProductAttributeModel, ids, nil, pas); err != nil {
		return nil, err
	}
	return pas, nil
}

// FindProductAttribute finds product.attribute record by querying it with criteria.
func (c *Client) FindProductAttribute(criteria *Criteria) (*ProductAttribute, error) {
	pas := &ProductAttributes{}
	if err := c.SearchRead(ProductAttributeModel, criteria, NewOptions().Limit(1), pas); err != nil {
		return nil, err
	}
	if pas != nil && len(*pas) > 0 {
		return &((*pas)[0]), nil
	}
	return nil, fmt.Errorf("product.attribute was not found with criteria %v", criteria)
}

// FindProductAttributes finds product.attribute records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributes(criteria *Criteria, options *Options) (*ProductAttributes, error) {
	pas := &ProductAttributes{}
	if err := c.SearchRead(ProductAttributeModel, criteria, options, pas); err != nil {
		return nil, err
	}
	return pas, nil
}

// FindProductAttributeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductAttributeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductAttributeId finds record id by querying it with criteria.
func (c *Client) FindProductAttributeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.attribute was not found with criteria %v and options %v", criteria, options)
}
