package odoo

import (
	"fmt"
)

// ProductAttributeLine represents product.attribute.line model.
type ProductAttributeLine struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	AttributeId   *Many2One `xmlrpc:"attribute_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	ProductTmplId *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ValueIds      *Relation `xmlrpc:"value_ids,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductAttributeLines represents array of product.attribute.line model.
type ProductAttributeLines []ProductAttributeLine

// ProductAttributeLineModel is the odoo model name.
const ProductAttributeLineModel = "product.attribute.line"

// Many2One convert ProductAttributeLine to *Many2One.
func (pal *ProductAttributeLine) Many2One() *Many2One {
	return NewMany2One(pal.Id.Get(), "")
}

// CreateProductAttributeLine creates a new product.attribute.line model and returns its id.
func (c *Client) CreateProductAttributeLine(pal *ProductAttributeLine) (int64, error) {
	return c.Create(ProductAttributeLineModel, pal)
}

// UpdateProductAttributeLine updates an existing product.attribute.line record.
func (c *Client) UpdateProductAttributeLine(pal *ProductAttributeLine) error {
	return c.UpdateProductAttributeLines([]int64{pal.Id.Get()}, pal)
}

// UpdateProductAttributeLines updates existing product.attribute.line records.
// All records (represented by ids) will be updated by pal values.
func (c *Client) UpdateProductAttributeLines(ids []int64, pal *ProductAttributeLine) error {
	return c.Update(ProductAttributeLineModel, ids, pal)
}

// DeleteProductAttributeLine deletes an existing product.attribute.line record.
func (c *Client) DeleteProductAttributeLine(id int64) error {
	return c.DeleteProductAttributeLines([]int64{id})
}

// DeleteProductAttributeLines deletes existing product.attribute.line records.
func (c *Client) DeleteProductAttributeLines(ids []int64) error {
	return c.Delete(ProductAttributeLineModel, ids)
}

// GetProductAttributeLine gets product.attribute.line existing record.
func (c *Client) GetProductAttributeLine(id int64) (*ProductAttributeLine, error) {
	pals, err := c.GetProductAttributeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if pals != nil && len(*pals) > 0 {
		return &((*pals)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.attribute.line not found", id)
}

// GetProductAttributeLines gets product.attribute.line existing records.
func (c *Client) GetProductAttributeLines(ids []int64) (*ProductAttributeLines, error) {
	pals := &ProductAttributeLines{}
	if err := c.Read(ProductAttributeLineModel, ids, nil, pals); err != nil {
		return nil, err
	}
	return pals, nil
}

// FindProductAttributeLine finds product.attribute.line record by querying it with criteria.
func (c *Client) FindProductAttributeLine(criteria *Criteria) (*ProductAttributeLine, error) {
	pals := &ProductAttributeLines{}
	if err := c.SearchRead(ProductAttributeLineModel, criteria, NewOptions().Limit(1), pals); err != nil {
		return nil, err
	}
	if pals != nil && len(*pals) > 0 {
		return &((*pals)[0]), nil
	}
	return nil, fmt.Errorf("no product.attribute.line was found with criteria %v", criteria)
}

// FindProductAttributeLines finds product.attribute.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeLines(criteria *Criteria, options *Options) (*ProductAttributeLines, error) {
	pals := &ProductAttributeLines{}
	if err := c.SearchRead(ProductAttributeLineModel, criteria, options, pals); err != nil {
		return nil, err
	}
	return pals, nil
}

// FindProductAttributeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductAttributeLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductAttributeLineId finds record id by querying it with criteria.
func (c *Client) FindProductAttributeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no product.attribute.line was found with criteria %v and options %v", criteria, options)
}
