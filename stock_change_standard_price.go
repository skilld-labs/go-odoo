package odoo

import (
	"fmt"
)

// StockChangeStandardPrice represents stock.change.standard.price model.
type StockChangeStandardPrice struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omitempty"`
	CounterpartAccountId         *Many2One `xmlrpc:"counterpart_account_id,omitempty"`
	CounterpartAccountIdRequired *Bool     `xmlrpc:"counterpart_account_id_required,omitempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omitempty"`
	Id                           *Int      `xmlrpc:"id,omitempty"`
	NewPrice                     *Float    `xmlrpc:"new_price,omitempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockChangeStandardPrices represents array of stock.change.standard.price model.
type StockChangeStandardPrices []StockChangeStandardPrice

// StockChangeStandardPriceModel is the odoo model name.
const StockChangeStandardPriceModel = "stock.change.standard.price"

// Many2One convert StockChangeStandardPrice to *Many2One.
func (scsp *StockChangeStandardPrice) Many2One() *Many2One {
	return NewMany2One(scsp.Id.Get(), "")
}

// CreateStockChangeStandardPrice creates a new stock.change.standard.price model and returns its id.
func (c *Client) CreateStockChangeStandardPrice(scsp *StockChangeStandardPrice) (int64, error) {
	return c.Create(StockChangeStandardPriceModel, scsp)
}

// UpdateStockChangeStandardPrice updates an existing stock.change.standard.price record.
func (c *Client) UpdateStockChangeStandardPrice(scsp *StockChangeStandardPrice) error {
	return c.UpdateStockChangeStandardPrices([]int64{scsp.Id.Get()}, scsp)
}

// UpdateStockChangeStandardPrices updates existing stock.change.standard.price records.
// All records (represented by ids) will be updated by scsp values.
func (c *Client) UpdateStockChangeStandardPrices(ids []int64, scsp *StockChangeStandardPrice) error {
	return c.Update(StockChangeStandardPriceModel, ids, scsp)
}

// DeleteStockChangeStandardPrice deletes an existing stock.change.standard.price record.
func (c *Client) DeleteStockChangeStandardPrice(id int64) error {
	return c.DeleteStockChangeStandardPrices([]int64{id})
}

// DeleteStockChangeStandardPrices deletes existing stock.change.standard.price records.
func (c *Client) DeleteStockChangeStandardPrices(ids []int64) error {
	return c.Delete(StockChangeStandardPriceModel, ids)
}

// GetStockChangeStandardPrice gets stock.change.standard.price existing record.
func (c *Client) GetStockChangeStandardPrice(id int64) (*StockChangeStandardPrice, error) {
	scsps, err := c.GetStockChangeStandardPrices([]int64{id})
	if err != nil {
		return nil, err
	}
	if scsps != nil && len(*scsps) > 0 {
		return &((*scsps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.change.standard.price not found", id)
}

// GetStockChangeStandardPrices gets stock.change.standard.price existing records.
func (c *Client) GetStockChangeStandardPrices(ids []int64) (*StockChangeStandardPrices, error) {
	scsps := &StockChangeStandardPrices{}
	if err := c.Read(StockChangeStandardPriceModel, ids, nil, scsps); err != nil {
		return nil, err
	}
	return scsps, nil
}

// FindStockChangeStandardPrice finds stock.change.standard.price record by querying it with criteria.
func (c *Client) FindStockChangeStandardPrice(criteria *Criteria) (*StockChangeStandardPrice, error) {
	scsps := &StockChangeStandardPrices{}
	if err := c.SearchRead(StockChangeStandardPriceModel, criteria, NewOptions().Limit(1), scsps); err != nil {
		return nil, err
	}
	if scsps != nil && len(*scsps) > 0 {
		return &((*scsps)[0]), nil
	}
	return nil, fmt.Errorf("no stock.change.standard.price was found with criteria %v", criteria)
}

// FindStockChangeStandardPrices finds stock.change.standard.price records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeStandardPrices(criteria *Criteria, options *Options) (*StockChangeStandardPrices, error) {
	scsps := &StockChangeStandardPrices{}
	if err := c.SearchRead(StockChangeStandardPriceModel, criteria, options, scsps); err != nil {
		return nil, err
	}
	return scsps, nil
}

// FindStockChangeStandardPriceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeStandardPriceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockChangeStandardPriceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockChangeStandardPriceId finds record id by querying it with criteria.
func (c *Client) FindStockChangeStandardPriceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockChangeStandardPriceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.change.standard.price was found with criteria %v and options %v", criteria, options)
}
