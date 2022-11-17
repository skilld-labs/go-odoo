package odoo

import (
	"fmt"
)

// ProductUomCateg represents product.uom.categ model.
type ProductUomCateg struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductUomCategs represents array of product.uom.categ model.
type ProductUomCategs []ProductUomCateg

// ProductUomCategModel is the odoo model name.
const ProductUomCategModel = "product.uom.categ"

// Many2One convert ProductUomCateg to *Many2One.
func (puc *ProductUomCateg) Many2One() *Many2One {
	return NewMany2One(puc.Id.Get(), "")
}

// CreateProductUomCateg creates a new product.uom.categ model and returns its id.
func (c *Client) CreateProductUomCateg(puc *ProductUomCateg) (int64, error) {
	return c.Create(ProductUomCategModel, puc)
}

// UpdateProductUomCateg updates an existing product.uom.categ record.
func (c *Client) UpdateProductUomCateg(puc *ProductUomCateg) error {
	return c.UpdateProductUomCategs([]int64{puc.Id.Get()}, puc)
}

// UpdateProductUomCategs updates existing product.uom.categ records.
// All records (represented by ids) will be updated by puc values.
func (c *Client) UpdateProductUomCategs(ids []int64, puc *ProductUomCateg) error {
	return c.Update(ProductUomCategModel, ids, puc)
}

// DeleteProductUomCateg deletes an existing product.uom.categ record.
func (c *Client) DeleteProductUomCateg(id int64) error {
	return c.DeleteProductUomCategs([]int64{id})
}

// DeleteProductUomCategs deletes existing product.uom.categ records.
func (c *Client) DeleteProductUomCategs(ids []int64) error {
	return c.Delete(ProductUomCategModel, ids)
}

// GetProductUomCateg gets product.uom.categ existing record.
func (c *Client) GetProductUomCateg(id int64) (*ProductUomCateg, error) {
	pucs, err := c.GetProductUomCategs([]int64{id})
	if err != nil {
		return nil, err
	}
	if pucs != nil && len(*pucs) > 0 {
		return &((*pucs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.uom.categ not found", id)
}

// GetProductUomCategs gets product.uom.categ existing records.
func (c *Client) GetProductUomCategs(ids []int64) (*ProductUomCategs, error) {
	pucs := &ProductUomCategs{}
	if err := c.Read(ProductUomCategModel, ids, nil, pucs); err != nil {
		return nil, err
	}
	return pucs, nil
}

// FindProductUomCateg finds product.uom.categ record by querying it with criteria.
func (c *Client) FindProductUomCateg(criteria *Criteria) (*ProductUomCateg, error) {
	pucs := &ProductUomCategs{}
	if err := c.SearchRead(ProductUomCategModel, criteria, NewOptions().Limit(1), pucs); err != nil {
		return nil, err
	}
	if pucs != nil && len(*pucs) > 0 {
		return &((*pucs)[0]), nil
	}
	return nil, fmt.Errorf("no product.uom.categ was found with criteria %v", criteria)
}

// FindProductUomCategs finds product.uom.categ records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUomCategs(criteria *Criteria, options *Options) (*ProductUomCategs, error) {
	pucs := &ProductUomCategs{}
	if err := c.SearchRead(ProductUomCategModel, criteria, options, pucs); err != nil {
		return nil, err
	}
	return pucs, nil
}

// FindProductUomCategIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUomCategIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductUomCategModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductUomCategId finds record id by querying it with criteria.
func (c *Client) FindProductUomCategId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductUomCategModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no product.uom.categ was found with criteria %v and options %v", criteria, options)
}
