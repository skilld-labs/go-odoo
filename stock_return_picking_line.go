package odoo

import (
	"fmt"
)

// StockReturnPickingLine represents stock.return.picking.line model.
type StockReturnPickingLine struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MoveId      *Many2One `xmlrpc:"move_id,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	Quantity    *Float    `xmlrpc:"quantity,omptempty"`
	ToRefund    *Bool     `xmlrpc:"to_refund,omptempty"`
	UomId       *Many2One `xmlrpc:"uom_id,omptempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockReturnPickingLines represents array of stock.return.picking.line model.
type StockReturnPickingLines []StockReturnPickingLine

// StockReturnPickingLineModel is the odoo model name.
const StockReturnPickingLineModel = "stock.return.picking.line"

// Many2One convert StockReturnPickingLine to *Many2One.
func (srpl *StockReturnPickingLine) Many2One() *Many2One {
	return NewMany2One(srpl.Id.Get(), "")
}

// CreateStockReturnPickingLine creates a new stock.return.picking.line model and returns its id.
func (c *Client) CreateStockReturnPickingLine(srpl *StockReturnPickingLine) (int64, error) {
	return c.Create(StockReturnPickingLineModel, srpl)
}

// UpdateStockReturnPickingLine updates an existing stock.return.picking.line record.
func (c *Client) UpdateStockReturnPickingLine(srpl *StockReturnPickingLine) error {
	return c.UpdateStockReturnPickingLines([]int64{srpl.Id.Get()}, srpl)
}

// UpdateStockReturnPickingLines updates existing stock.return.picking.line records.
// All records (represented by ids) will be updated by srpl values.
func (c *Client) UpdateStockReturnPickingLines(ids []int64, srpl *StockReturnPickingLine) error {
	return c.Update(StockReturnPickingLineModel, ids, srpl)
}

// DeleteStockReturnPickingLine deletes an existing stock.return.picking.line record.
func (c *Client) DeleteStockReturnPickingLine(id int64) error {
	return c.DeleteStockReturnPickingLines([]int64{id})
}

// DeleteStockReturnPickingLines deletes existing stock.return.picking.line records.
func (c *Client) DeleteStockReturnPickingLines(ids []int64) error {
	return c.Delete(StockReturnPickingLineModel, ids)
}

// GetStockReturnPickingLine gets stock.return.picking.line existing record.
func (c *Client) GetStockReturnPickingLine(id int64) (*StockReturnPickingLine, error) {
	srpls, err := c.GetStockReturnPickingLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if srpls != nil && len(*srpls) > 0 {
		return &((*srpls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.return.picking.line not found", id)
}

// GetStockReturnPickingLines gets stock.return.picking.line existing records.
func (c *Client) GetStockReturnPickingLines(ids []int64) (*StockReturnPickingLines, error) {
	srpls := &StockReturnPickingLines{}
	if err := c.Read(StockReturnPickingLineModel, ids, nil, srpls); err != nil {
		return nil, err
	}
	return srpls, nil
}

// FindStockReturnPickingLine finds stock.return.picking.line record by querying it with criteria.
func (c *Client) FindStockReturnPickingLine(criteria *Criteria) (*StockReturnPickingLine, error) {
	srpls := &StockReturnPickingLines{}
	if err := c.SearchRead(StockReturnPickingLineModel, criteria, NewOptions().Limit(1), srpls); err != nil {
		return nil, err
	}
	if srpls != nil && len(*srpls) > 0 {
		return &((*srpls)[0]), nil
	}
	return nil, fmt.Errorf("no stock.return.picking.line was found with criteria %v", criteria)
}

// FindStockReturnPickingLines finds stock.return.picking.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReturnPickingLines(criteria *Criteria, options *Options) (*StockReturnPickingLines, error) {
	srpls := &StockReturnPickingLines{}
	if err := c.SearchRead(StockReturnPickingLineModel, criteria, options, srpls); err != nil {
		return nil, err
	}
	return srpls, nil
}

// FindStockReturnPickingLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReturnPickingLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockReturnPickingLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockReturnPickingLineId finds record id by querying it with criteria.
func (c *Client) FindStockReturnPickingLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReturnPickingLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.return.picking.line was found with criteria %v and options %v", criteria, options)
}
