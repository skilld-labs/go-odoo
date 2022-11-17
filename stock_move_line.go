package odoo

import (
	"fmt"
)

// StockMoveLine represents stock.move.line model.
type StockMoveLine struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omitempty"`
	ConsumeLineIds          *Relation  `xmlrpc:"consume_line_ids,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                    *Time      `xmlrpc:"date,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty"`
	FromLoc                 *String    `xmlrpc:"from_loc,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty"`
	InEntirePackage         *Bool      `xmlrpc:"in_entire_package,omitempty"`
	IsInitialDemandEditable *Bool      `xmlrpc:"is_initial_demand_editable,omitempty"`
	IsLocked                *Bool      `xmlrpc:"is_locked,omitempty"`
	LocationDestId          *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId              *Many2One  `xmlrpc:"location_id,omitempty"`
	LotId                   *Many2One  `xmlrpc:"lot_id,omitempty"`
	LotName                 *String    `xmlrpc:"lot_name,omitempty"`
	LotsVisible             *Bool      `xmlrpc:"lots_visible,omitempty"`
	MoveId                  *Many2One  `xmlrpc:"move_id,omitempty"`
	OrderedQty              *Float     `xmlrpc:"ordered_qty,omitempty"`
	OwnerId                 *Many2One  `xmlrpc:"owner_id,omitempty"`
	PackageId               *Many2One  `xmlrpc:"package_id,omitempty"`
	PickingId               *Many2One  `xmlrpc:"picking_id,omitempty"`
	ProduceLineIds          *Relation  `xmlrpc:"produce_line_ids,omitempty"`
	ProductId               *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty              *Float     `xmlrpc:"product_qty,omitempty"`
	ProductUomId            *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductUomQty           *Float     `xmlrpc:"product_uom_qty,omitempty"`
	QtyDone                 *Float     `xmlrpc:"qty_done,omitempty"`
	Reference               *String    `xmlrpc:"reference,omitempty"`
	ResultPackageId         *Many2One  `xmlrpc:"result_package_id,omitempty"`
	State                   *Selection `xmlrpc:"state,omitempty"`
	ToLoc                   *String    `xmlrpc:"to_loc,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockMoveLines represents array of stock.move.line model.
type StockMoveLines []StockMoveLine

// StockMoveLineModel is the odoo model name.
const StockMoveLineModel = "stock.move.line"

// Many2One convert StockMoveLine to *Many2One.
func (sml *StockMoveLine) Many2One() *Many2One {
	return NewMany2One(sml.Id.Get(), "")
}

// CreateStockMoveLine creates a new stock.move.line model and returns its id.
func (c *Client) CreateStockMoveLine(sml *StockMoveLine) (int64, error) {
	return c.Create(StockMoveLineModel, sml)
}

// UpdateStockMoveLine updates an existing stock.move.line record.
func (c *Client) UpdateStockMoveLine(sml *StockMoveLine) error {
	return c.UpdateStockMoveLines([]int64{sml.Id.Get()}, sml)
}

// UpdateStockMoveLines updates existing stock.move.line records.
// All records (represented by ids) will be updated by sml values.
func (c *Client) UpdateStockMoveLines(ids []int64, sml *StockMoveLine) error {
	return c.Update(StockMoveLineModel, ids, sml)
}

// DeleteStockMoveLine deletes an existing stock.move.line record.
func (c *Client) DeleteStockMoveLine(id int64) error {
	return c.DeleteStockMoveLines([]int64{id})
}

// DeleteStockMoveLines deletes existing stock.move.line records.
func (c *Client) DeleteStockMoveLines(ids []int64) error {
	return c.Delete(StockMoveLineModel, ids)
}

// GetStockMoveLine gets stock.move.line existing record.
func (c *Client) GetStockMoveLine(id int64) (*StockMoveLine, error) {
	smls, err := c.GetStockMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if smls != nil && len(*smls) > 0 {
		return &((*smls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.move.line not found", id)
}

// GetStockMoveLines gets stock.move.line existing records.
func (c *Client) GetStockMoveLines(ids []int64) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.Read(StockMoveLineModel, ids, nil, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLine finds stock.move.line record by querying it with criteria.
func (c *Client) FindStockMoveLine(criteria *Criteria) (*StockMoveLine, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, NewOptions().Limit(1), smls); err != nil {
		return nil, err
	}
	if smls != nil && len(*smls) > 0 {
		return &((*smls)[0]), nil
	}
	return nil, fmt.Errorf("no stock.move.line was found with criteria %v", criteria)
}

// FindStockMoveLines finds stock.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLines(criteria *Criteria, options *Options) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, options, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockMoveLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockMoveLineId finds record id by querying it with criteria.
func (c *Client) FindStockMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.move.line was found with criteria %v and options %v", criteria, options)
}
