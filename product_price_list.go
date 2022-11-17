package odoo

import (
	"fmt"
)

// ProductPriceList represents product.price_list model.
type ProductPriceList struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PriceList   *Many2One `xmlrpc:"price_list,omitempty"`
	Qty1        *Int      `xmlrpc:"qty1,omitempty"`
	Qty2        *Int      `xmlrpc:"qty2,omitempty"`
	Qty3        *Int      `xmlrpc:"qty3,omitempty"`
	Qty4        *Int      `xmlrpc:"qty4,omitempty"`
	Qty5        *Int      `xmlrpc:"qty5,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductPriceLists represents array of product.price_list model.
type ProductPriceLists []ProductPriceList

// ProductPriceListModel is the odoo model name.
const ProductPriceListModel = "product.price_list"

// Many2One convert ProductPriceList to *Many2One.
func (pp *ProductPriceList) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPriceList creates a new product.price_list model and returns its id.
func (c *Client) CreateProductPriceList(pp *ProductPriceList) (int64, error) {
	return c.Create(ProductPriceListModel, pp)
}

// UpdateProductPriceList updates an existing product.price_list record.
func (c *Client) UpdateProductPriceList(pp *ProductPriceList) error {
	return c.UpdateProductPriceLists([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPriceLists updates existing product.price_list records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPriceLists(ids []int64, pp *ProductPriceList) error {
	return c.Update(ProductPriceListModel, ids, pp)
}

// DeleteProductPriceList deletes an existing product.price_list record.
func (c *Client) DeleteProductPriceList(id int64) error {
	return c.DeleteProductPriceLists([]int64{id})
}

// DeleteProductPriceLists deletes existing product.price_list records.
func (c *Client) DeleteProductPriceLists(ids []int64) error {
	return c.Delete(ProductPriceListModel, ids)
}

// GetProductPriceList gets product.price_list existing record.
func (c *Client) GetProductPriceList(id int64) (*ProductPriceList, error) {
	pps, err := c.GetProductPriceLists([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.price_list not found", id)
}

// GetProductPriceLists gets product.price_list existing records.
func (c *Client) GetProductPriceLists(ids []int64) (*ProductPriceLists, error) {
	pps := &ProductPriceLists{}
	if err := c.Read(ProductPriceListModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPriceList finds product.price_list record by querying it with criteria.
func (c *Client) FindProductPriceList(criteria *Criteria) (*ProductPriceList, error) {
	pps := &ProductPriceLists{}
	if err := c.SearchRead(ProductPriceListModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("no product.price_list was found with criteria %v", criteria)
}

// FindProductPriceLists finds product.price_list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPriceLists(criteria *Criteria, options *Options) (*ProductPriceLists, error) {
	pps := &ProductPriceLists{}
	if err := c.SearchRead(ProductPriceListModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPriceListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPriceListIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPriceListModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPriceListId finds record id by querying it with criteria.
func (c *Client) FindProductPriceListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPriceListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no product.price_list was found with criteria %v and options %v", criteria, options)
}
