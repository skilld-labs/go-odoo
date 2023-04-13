package odoo

import (
	"fmt"
)

// StockQuantPackage represents stock.quant.package model.
type StockQuantPackage struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId                    *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrentDestinationLocationId *Many2One `xmlrpc:"current_destination_location_id,omptempty"`
	CurrentPickingId             *Bool     `xmlrpc:"current_picking_id,omptempty"`
	CurrentPickingMoveLineIds    *Relation `xmlrpc:"current_picking_move_line_ids,omptempty"`
	CurrentSourceLocationId      *Many2One `xmlrpc:"current_source_location_id,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	IsProcessed                  *Bool     `xmlrpc:"is_processed,omptempty"`
	LocationId                   *Many2One `xmlrpc:"location_id,omptempty"`
	MoveLineIds                  *Relation `xmlrpc:"move_line_ids,omptempty"`
	Name                         *String   `xmlrpc:"name,omptempty"`
	OwnerId                      *Many2One `xmlrpc:"owner_id,omptempty"`
	PackagingId                  *Many2One `xmlrpc:"packaging_id,omptempty"`
	QuantIds                     *Relation `xmlrpc:"quant_ids,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockQuantPackages represents array of stock.quant.package model.
type StockQuantPackages []StockQuantPackage

// StockQuantPackageModel is the odoo model name.
const StockQuantPackageModel = "stock.quant.package"

// Many2One convert StockQuantPackage to *Many2One.
func (sqp *StockQuantPackage) Many2One() *Many2One {
	return NewMany2One(sqp.Id.Get(), "")
}

// CreateStockQuantPackage creates a new stock.quant.package model and returns its id.
func (c *Client) CreateStockQuantPackage(sqp *StockQuantPackage) (int64, error) {
	ids, err := c.CreateStockQuantPackages([]*StockQuantPackage{sqp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockQuantPackage creates a new stock.quant.package model and returns its id.
func (c *Client) CreateStockQuantPackages(sqps []*StockQuantPackage) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqps {
		vv = append(vv, v)
	}
	return c.Create(StockQuantPackageModel, vv)
}

// UpdateStockQuantPackage updates an existing stock.quant.package record.
func (c *Client) UpdateStockQuantPackage(sqp *StockQuantPackage) error {
	return c.UpdateStockQuantPackages([]int64{sqp.Id.Get()}, sqp)
}

// UpdateStockQuantPackages updates existing stock.quant.package records.
// All records (represented by ids) will be updated by sqp values.
func (c *Client) UpdateStockQuantPackages(ids []int64, sqp *StockQuantPackage) error {
	return c.Update(StockQuantPackageModel, ids, sqp)
}

// DeleteStockQuantPackage deletes an existing stock.quant.package record.
func (c *Client) DeleteStockQuantPackage(id int64) error {
	return c.DeleteStockQuantPackages([]int64{id})
}

// DeleteStockQuantPackages deletes existing stock.quant.package records.
func (c *Client) DeleteStockQuantPackages(ids []int64) error {
	return c.Delete(StockQuantPackageModel, ids)
}

// GetStockQuantPackage gets stock.quant.package existing record.
func (c *Client) GetStockQuantPackage(id int64) (*StockQuantPackage, error) {
	sqps, err := c.GetStockQuantPackages([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqps != nil && len(*sqps) > 0 {
		return &((*sqps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.quant.package not found", id)
}

// GetStockQuantPackages gets stock.quant.package existing records.
func (c *Client) GetStockQuantPackages(ids []int64) (*StockQuantPackages, error) {
	sqps := &StockQuantPackages{}
	if err := c.Read(StockQuantPackageModel, ids, nil, sqps); err != nil {
		return nil, err
	}
	return sqps, nil
}

// FindStockQuantPackage finds stock.quant.package record by querying it with criteria.
func (c *Client) FindStockQuantPackage(criteria *Criteria) (*StockQuantPackage, error) {
	sqps := &StockQuantPackages{}
	if err := c.SearchRead(StockQuantPackageModel, criteria, NewOptions().Limit(1), sqps); err != nil {
		return nil, err
	}
	if sqps != nil && len(*sqps) > 0 {
		return &((*sqps)[0]), nil
	}
	return nil, fmt.Errorf("stock.quant.package was not found with criteria %v", criteria)
}

// FindStockQuantPackages finds stock.quant.package records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantPackages(criteria *Criteria, options *Options) (*StockQuantPackages, error) {
	sqps := &StockQuantPackages{}
	if err := c.SearchRead(StockQuantPackageModel, criteria, options, sqps); err != nil {
		return nil, err
	}
	return sqps, nil
}

// FindStockQuantPackageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantPackageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockQuantPackageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockQuantPackageId finds record id by querying it with criteria.
func (c *Client) FindStockQuantPackageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockQuantPackageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.quant.package was not found with criteria %v and options %v", criteria, options)
}
