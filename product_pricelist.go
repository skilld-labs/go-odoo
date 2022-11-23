package odoo

import (
	"fmt"
)

// ProductPricelist represents product.pricelist model.
type ProductPricelist struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Active          *Bool      `xmlrpc:"active,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryGroupIds *Relation  `xmlrpc:"country_group_ids,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omptempty"`
	DiscountPolicy  *Selection `xmlrpc:"discount_policy,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	ItemIds         *Relation  `xmlrpc:"item_ids,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	Sequence        *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductPricelists represents array of product.pricelist model.
type ProductPricelists []ProductPricelist

// ProductPricelistModel is the odoo model name.
const ProductPricelistModel = "product.pricelist"

// Many2One convert ProductPricelist to *Many2One.
func (pp *ProductPricelist) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPricelist creates a new product.pricelist model and returns its id.
func (c *Client) CreateProductPricelist(pp *ProductPricelist) (int64, error) {
	return c.Create(ProductPricelistModel, pp)
}

// UpdateProductPricelist updates an existing product.pricelist record.
func (c *Client) UpdateProductPricelist(pp *ProductPricelist) error {
	return c.UpdateProductPricelists([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPricelists updates existing product.pricelist records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPricelists(ids []int64, pp *ProductPricelist) error {
	return c.Update(ProductPricelistModel, ids, pp)
}

// DeleteProductPricelist deletes an existing product.pricelist record.
func (c *Client) DeleteProductPricelist(id int64) error {
	return c.DeleteProductPricelists([]int64{id})
}

// DeleteProductPricelists deletes existing product.pricelist records.
func (c *Client) DeleteProductPricelists(ids []int64) error {
	return c.Delete(ProductPricelistModel, ids)
}

// GetProductPricelist gets product.pricelist existing record.
func (c *Client) GetProductPricelist(id int64) (*ProductPricelist, error) {
	pps, err := c.GetProductPricelists([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.pricelist not found", id)
}

// GetProductPricelists gets product.pricelist existing records.
func (c *Client) GetProductPricelists(ids []int64) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.Read(ProductPricelistModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelist finds product.pricelist record by querying it with criteria.
func (c *Client) FindProductPricelist(criteria *Criteria) (*ProductPricelist, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.pricelist was not found with criteria %v", criteria)
}

// FindProductPricelists finds product.pricelist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelists(criteria *Criteria, options *Options) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPricelistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPricelistId finds record id by querying it with criteria.
func (c *Client) FindProductPricelistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricelistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.pricelist was not found with criteria %v and options %v", criteria, options)
}
