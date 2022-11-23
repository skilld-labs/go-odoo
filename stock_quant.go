package odoo

import (
	"fmt"
)

// StockQuant represents stock.quant model.
type StockQuant struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId        *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	InDate           *Time     `xmlrpc:"in_date,omptempty"`
	LocationId       *Many2One `xmlrpc:"location_id,omptempty"`
	LotId            *Many2One `xmlrpc:"lot_id,omptempty"`
	OwnerId          *Many2One `xmlrpc:"owner_id,omptempty"`
	PackageId        *Many2One `xmlrpc:"package_id,omptempty"`
	ProductId        *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId    *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUomId     *Many2One `xmlrpc:"product_uom_id,omptempty"`
	Quantity         *Float    `xmlrpc:"quantity,omptempty"`
	ReservedQuantity *Float    `xmlrpc:"reserved_quantity,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockQuants represents array of stock.quant model.
type StockQuants []StockQuant

// StockQuantModel is the odoo model name.
const StockQuantModel = "stock.quant"

// Many2One convert StockQuant to *Many2One.
func (sq *StockQuant) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

// CreateStockQuant creates a new stock.quant model and returns its id.
func (c *Client) CreateStockQuant(sq *StockQuant) (int64, error) {
	return c.Create(StockQuantModel, sq)
}

// UpdateStockQuant updates an existing stock.quant record.
func (c *Client) UpdateStockQuant(sq *StockQuant) error {
	return c.UpdateStockQuants([]int64{sq.Id.Get()}, sq)
}

// UpdateStockQuants updates existing stock.quant records.
// All records (represented by ids) will be updated by sq values.
func (c *Client) UpdateStockQuants(ids []int64, sq *StockQuant) error {
	return c.Update(StockQuantModel, ids, sq)
}

// DeleteStockQuant deletes an existing stock.quant record.
func (c *Client) DeleteStockQuant(id int64) error {
	return c.DeleteStockQuants([]int64{id})
}

// DeleteStockQuants deletes existing stock.quant records.
func (c *Client) DeleteStockQuants(ids []int64) error {
	return c.Delete(StockQuantModel, ids)
}

// GetStockQuant gets stock.quant existing record.
func (c *Client) GetStockQuant(id int64) (*StockQuant, error) {
	sqs, err := c.GetStockQuants([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.quant not found", id)
}

// GetStockQuants gets stock.quant existing records.
func (c *Client) GetStockQuants(ids []int64) (*StockQuants, error) {
	sqs := &StockQuants{}
	if err := c.Read(StockQuantModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindStockQuant finds stock.quant record by querying it with criteria.
func (c *Client) FindStockQuant(criteria *Criteria) (*StockQuant, error) {
	sqs := &StockQuants{}
	if err := c.SearchRead(StockQuantModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("stock.quant was not found with criteria %v", criteria)
}

// FindStockQuants finds stock.quant records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuants(criteria *Criteria, options *Options) (*StockQuants, error) {
	sqs := &StockQuants{}
	if err := c.SearchRead(StockQuantModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindStockQuantIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockQuantModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockQuantId finds record id by querying it with criteria.
func (c *Client) FindStockQuantId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockQuantModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.quant was not found with criteria %v and options %v", criteria, options)
}
