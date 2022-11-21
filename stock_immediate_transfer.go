package odoo

import (
	"fmt"
)

// StockImmediateTransfer represents stock.immediate.transfer model.
type StockImmediateTransfer struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	PickIds     *Relation `xmlrpc:"pick_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockImmediateTransfers represents array of stock.immediate.transfer model.
type StockImmediateTransfers []StockImmediateTransfer

// StockImmediateTransferModel is the odoo model name.
const StockImmediateTransferModel = "stock.immediate.transfer"

// Many2One convert StockImmediateTransfer to *Many2One.
func (sit *StockImmediateTransfer) Many2One() *Many2One {
	return NewMany2One(sit.Id.Get(), "")
}

// CreateStockImmediateTransfer creates a new stock.immediate.transfer model and returns its id.
func (c *Client) CreateStockImmediateTransfer(sit *StockImmediateTransfer) (int64, error) {
	return c.Create(StockImmediateTransferModel, sit)
}

// UpdateStockImmediateTransfer updates an existing stock.immediate.transfer record.
func (c *Client) UpdateStockImmediateTransfer(sit *StockImmediateTransfer) error {
	return c.UpdateStockImmediateTransfers([]int64{sit.Id.Get()}, sit)
}

// UpdateStockImmediateTransfers updates existing stock.immediate.transfer records.
// All records (represented by ids) will be updated by sit values.
func (c *Client) UpdateStockImmediateTransfers(ids []int64, sit *StockImmediateTransfer) error {
	return c.Update(StockImmediateTransferModel, ids, sit)
}

// DeleteStockImmediateTransfer deletes an existing stock.immediate.transfer record.
func (c *Client) DeleteStockImmediateTransfer(id int64) error {
	return c.DeleteStockImmediateTransfers([]int64{id})
}

// DeleteStockImmediateTransfers deletes existing stock.immediate.transfer records.
func (c *Client) DeleteStockImmediateTransfers(ids []int64) error {
	return c.Delete(StockImmediateTransferModel, ids)
}

// GetStockImmediateTransfer gets stock.immediate.transfer existing record.
func (c *Client) GetStockImmediateTransfer(id int64) (*StockImmediateTransfer, error) {
	sits, err := c.GetStockImmediateTransfers([]int64{id})
	if err != nil {
		return nil, err
	}
	if sits != nil && len(*sits) > 0 {
		return &((*sits)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.immediate.transfer not found", id)
}

// GetStockImmediateTransfers gets stock.immediate.transfer existing records.
func (c *Client) GetStockImmediateTransfers(ids []int64) (*StockImmediateTransfers, error) {
	sits := &StockImmediateTransfers{}
	if err := c.Read(StockImmediateTransferModel, ids, nil, sits); err != nil {
		return nil, err
	}
	return sits, nil
}

// FindStockImmediateTransfer finds stock.immediate.transfer record by querying it with criteria.
func (c *Client) FindStockImmediateTransfer(criteria *Criteria) (*StockImmediateTransfer, error) {
	sits := &StockImmediateTransfers{}
	if err := c.SearchRead(StockImmediateTransferModel, criteria, NewOptions().Limit(1), sits); err != nil {
		return nil, err
	}
	if sits != nil && len(*sits) > 0 {
		return &((*sits)[0]), nil
	}
	return nil, fmt.Errorf("no stock.immediate.transfer was found with criteria %v", criteria)
}

// FindStockImmediateTransfers finds stock.immediate.transfer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockImmediateTransfers(criteria *Criteria, options *Options) (*StockImmediateTransfers, error) {
	sits := &StockImmediateTransfers{}
	if err := c.SearchRead(StockImmediateTransferModel, criteria, options, sits); err != nil {
		return nil, err
	}
	return sits, nil
}

// FindStockImmediateTransferIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockImmediateTransferIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockImmediateTransferModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockImmediateTransferId finds record id by querying it with criteria.
func (c *Client) FindStockImmediateTransferId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockImmediateTransferModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.immediate.transfer was found with criteria %v and options %v", criteria, options)
}
