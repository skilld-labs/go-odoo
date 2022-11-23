package odoo

import (
	"fmt"
)

// ProductPutaway represents product.putaway model.
type ProductPutaway struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	FixedLocationIds *Relation `xmlrpc:"fixed_location_ids,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductPutaways represents array of product.putaway model.
type ProductPutaways []ProductPutaway

// ProductPutawayModel is the odoo model name.
const ProductPutawayModel = "product.putaway"

// Many2One convert ProductPutaway to *Many2One.
func (pp *ProductPutaway) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPutaway creates a new product.putaway model and returns its id.
func (c *Client) CreateProductPutaway(pp *ProductPutaway) (int64, error) {
	return c.Create(ProductPutawayModel, pp)
}

// UpdateProductPutaway updates an existing product.putaway record.
func (c *Client) UpdateProductPutaway(pp *ProductPutaway) error {
	return c.UpdateProductPutaways([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPutaways updates existing product.putaway records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPutaways(ids []int64, pp *ProductPutaway) error {
	return c.Update(ProductPutawayModel, ids, pp)
}

// DeleteProductPutaway deletes an existing product.putaway record.
func (c *Client) DeleteProductPutaway(id int64) error {
	return c.DeleteProductPutaways([]int64{id})
}

// DeleteProductPutaways deletes existing product.putaway records.
func (c *Client) DeleteProductPutaways(ids []int64) error {
	return c.Delete(ProductPutawayModel, ids)
}

// GetProductPutaway gets product.putaway existing record.
func (c *Client) GetProductPutaway(id int64) (*ProductPutaway, error) {
	pps, err := c.GetProductPutaways([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.putaway not found", id)
}

// GetProductPutaways gets product.putaway existing records.
func (c *Client) GetProductPutaways(ids []int64) (*ProductPutaways, error) {
	pps := &ProductPutaways{}
	if err := c.Read(ProductPutawayModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPutaway finds product.putaway record by querying it with criteria.
func (c *Client) FindProductPutaway(criteria *Criteria) (*ProductPutaway, error) {
	pps := &ProductPutaways{}
	if err := c.SearchRead(ProductPutawayModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.putaway was not found with criteria %v", criteria)
}

// FindProductPutaways finds product.putaway records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPutaways(criteria *Criteria, options *Options) (*ProductPutaways, error) {
	pps := &ProductPutaways{}
	if err := c.SearchRead(ProductPutawayModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPutawayIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPutawayIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPutawayModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPutawayId finds record id by querying it with criteria.
func (c *Client) FindProductPutawayId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPutawayModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.putaway was not found with criteria %v and options %v", criteria, options)
}
