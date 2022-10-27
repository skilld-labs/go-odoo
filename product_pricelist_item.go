package odoo

import (
	"fmt"
)

// ProductPricelistItem represents product.pricelist.item model.
type ProductPricelistItem struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	Active          *Bool      `xmlrpc:"active,omitempty"`
	AppliedOn       *Selection `xmlrpc:"applied_on,omitempty"`
	Base            *Selection `xmlrpc:"base,omitempty"`
	BasePricelistId *Many2One  `xmlrpc:"base_pricelist_id,omitempty"`
	CategId         *Many2One  `xmlrpc:"categ_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	ComputePrice    *Selection `xmlrpc:"compute_price,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateEnd         *Time      `xmlrpc:"date_end,omitempty"`
	DateStart       *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	FixedPrice      *Float     `xmlrpc:"fixed_price,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	MinQuantity     *Float     `xmlrpc:"min_quantity,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	PercentPrice    *Float     `xmlrpc:"percent_price,omitempty"`
	Price           *String    `xmlrpc:"price,omitempty"`
	PriceDiscount   *Float     `xmlrpc:"price_discount,omitempty"`
	PriceMaxMargin  *Float     `xmlrpc:"price_max_margin,omitempty"`
	PriceMinMargin  *Float     `xmlrpc:"price_min_margin,omitempty"`
	PriceRound      *Float     `xmlrpc:"price_round,omitempty"`
	PriceSurcharge  *Float     `xmlrpc:"price_surcharge,omitempty"`
	PricelistId     *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTmplId   *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductPricelistItems represents array of product.pricelist.item model.
type ProductPricelistItems []ProductPricelistItem

// ProductPricelistItemModel is the odoo model name.
const ProductPricelistItemModel = "product.pricelist.item"

// Many2One convert ProductPricelistItem to *Many2One.
func (ppi *ProductPricelistItem) Many2One() *Many2One {
	return NewMany2One(ppi.Id.Get(), "")
}

// CreateProductPricelistItem creates a new product.pricelist.item model and returns its id.
func (c *Client) CreateProductPricelistItem(ppi *ProductPricelistItem) (int64, error) {
	return c.Create(ProductPricelistItemModel, ppi)
}

// UpdateProductPricelistItem updates an existing product.pricelist.item record.
func (c *Client) UpdateProductPricelistItem(ppi *ProductPricelistItem) error {
	return c.UpdateProductPricelistItems([]int64{ppi.Id.Get()}, ppi)
}

// UpdateProductPricelistItems updates existing product.pricelist.item records.
// All records (represented by ids) will be updated by ppi values.
func (c *Client) UpdateProductPricelistItems(ids []int64, ppi *ProductPricelistItem) error {
	return c.Update(ProductPricelistItemModel, ids, ppi)
}

// DeleteProductPricelistItem deletes an existing product.pricelist.item record.
func (c *Client) DeleteProductPricelistItem(id int64) error {
	return c.DeleteProductPricelistItems([]int64{id})
}

// DeleteProductPricelistItems deletes existing product.pricelist.item records.
func (c *Client) DeleteProductPricelistItems(ids []int64) error {
	return c.Delete(ProductPricelistItemModel, ids)
}

// GetProductPricelistItem gets product.pricelist.item existing record.
func (c *Client) GetProductPricelistItem(id int64) (*ProductPricelistItem, error) {
	ppis, err := c.GetProductPricelistItems([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppis != nil && len(*ppis) > 0 {
		return &((*ppis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.pricelist.item not found", id)
}

// GetProductPricelistItems gets product.pricelist.item existing records.
func (c *Client) GetProductPricelistItems(ids []int64) (*ProductPricelistItems, error) {
	ppis := &ProductPricelistItems{}
	if err := c.Read(ProductPricelistItemModel, ids, nil, ppis); err != nil {
		return nil, err
	}
	return ppis, nil
}

// FindProductPricelistItem finds product.pricelist.item record by querying it with criteria.
func (c *Client) FindProductPricelistItem(criteria *Criteria) (*ProductPricelistItem, error) {
	ppis := &ProductPricelistItems{}
	if err := c.SearchRead(ProductPricelistItemModel, criteria, NewOptions().Limit(1), ppis); err != nil {
		return nil, err
	}
	if ppis != nil && len(*ppis) > 0 {
		return &((*ppis)[0]), nil
	}
	return nil, fmt.Errorf("product.pricelist.item was not found")
}

// FindProductPricelistItems finds product.pricelist.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistItems(criteria *Criteria, options *Options) (*ProductPricelistItems, error) {
	ppis := &ProductPricelistItems{}
	if err := c.SearchRead(ProductPricelistItemModel, criteria, options, ppis); err != nil {
		return nil, err
	}
	return ppis, nil
}

// FindProductPricelistItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPricelistItemModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPricelistItemId finds record id by querying it with criteria.
func (c *Client) FindProductPricelistItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricelistItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.pricelist.item was not found")
}
