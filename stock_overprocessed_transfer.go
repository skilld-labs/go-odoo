package odoo

import (
	"fmt"
)

// StockOverprocessedTransfer represents stock.overprocessed.transfer model.
type StockOverprocessedTransfer struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	OverprocessedProductName *String   `xmlrpc:"overprocessed_product_name,omptempty"`
	PickingId                *Many2One `xmlrpc:"picking_id,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockOverprocessedTransfers represents array of stock.overprocessed.transfer model.
type StockOverprocessedTransfers []StockOverprocessedTransfer

// StockOverprocessedTransferModel is the odoo model name.
const StockOverprocessedTransferModel = "stock.overprocessed.transfer"

// Many2One convert StockOverprocessedTransfer to *Many2One.
func (sot *StockOverprocessedTransfer) Many2One() *Many2One {
	return NewMany2One(sot.Id.Get(), "")
}

// CreateStockOverprocessedTransfer creates a new stock.overprocessed.transfer model and returns its id.
func (c *Client) CreateStockOverprocessedTransfer(sot *StockOverprocessedTransfer) (int64, error) {
	return c.Create(StockOverprocessedTransferModel, sot)
}

// UpdateStockOverprocessedTransfer updates an existing stock.overprocessed.transfer record.
func (c *Client) UpdateStockOverprocessedTransfer(sot *StockOverprocessedTransfer) error {
	return c.UpdateStockOverprocessedTransfers([]int64{sot.Id.Get()}, sot)
}

// UpdateStockOverprocessedTransfers updates existing stock.overprocessed.transfer records.
// All records (represented by ids) will be updated by sot values.
func (c *Client) UpdateStockOverprocessedTransfers(ids []int64, sot *StockOverprocessedTransfer) error {
	return c.Update(StockOverprocessedTransferModel, ids, sot)
}

// DeleteStockOverprocessedTransfer deletes an existing stock.overprocessed.transfer record.
func (c *Client) DeleteStockOverprocessedTransfer(id int64) error {
	return c.DeleteStockOverprocessedTransfers([]int64{id})
}

// DeleteStockOverprocessedTransfers deletes existing stock.overprocessed.transfer records.
func (c *Client) DeleteStockOverprocessedTransfers(ids []int64) error {
	return c.Delete(StockOverprocessedTransferModel, ids)
}

// GetStockOverprocessedTransfer gets stock.overprocessed.transfer existing record.
func (c *Client) GetStockOverprocessedTransfer(id int64) (*StockOverprocessedTransfer, error) {
	sots, err := c.GetStockOverprocessedTransfers([]int64{id})
	if err != nil {
		return nil, err
	}
	if sots != nil && len(*sots) > 0 {
		return &((*sots)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.overprocessed.transfer not found", id)
}

// GetStockOverprocessedTransfers gets stock.overprocessed.transfer existing records.
func (c *Client) GetStockOverprocessedTransfers(ids []int64) (*StockOverprocessedTransfers, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.Read(StockOverprocessedTransferModel, ids, nil, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindStockOverprocessedTransfer finds stock.overprocessed.transfer record by querying it with criteria.
func (c *Client) FindStockOverprocessedTransfer(criteria *Criteria) (*StockOverprocessedTransfer, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.SearchRead(StockOverprocessedTransferModel, criteria, NewOptions().Limit(1), sots); err != nil {
		return nil, err
	}
	if sots != nil && len(*sots) > 0 {
		return &((*sots)[0]), nil
	}
	return nil, fmt.Errorf("no stock.overprocessed.transfer was found with criteria %v", criteria)
}

// FindStockOverprocessedTransfers finds stock.overprocessed.transfer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOverprocessedTransfers(criteria *Criteria, options *Options) (*StockOverprocessedTransfers, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.SearchRead(StockOverprocessedTransferModel, criteria, options, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindStockOverprocessedTransferIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOverprocessedTransferIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockOverprocessedTransferModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockOverprocessedTransferId finds record id by querying it with criteria.
func (c *Client) FindStockOverprocessedTransferId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockOverprocessedTransferModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.overprocessed.transfer was found with criteria %v and options %v", criteria, options)
}
