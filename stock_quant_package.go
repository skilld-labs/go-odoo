package odoo

import (
	"fmt"
)

// StockQuantPackage represents stock.quant.package model.
type StockQuantPackage struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId                    *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentDestinationLocationId *Many2One `xmlrpc:"current_destination_location_id,omitempty"`
	CurrentPickingId             *Bool     `xmlrpc:"current_picking_id,omitempty"`
	CurrentPickingMoveLineIds    *Relation `xmlrpc:"current_picking_move_line_ids,omitempty"`
	CurrentSourceLocationId      *Many2One `xmlrpc:"current_source_location_id,omitempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omitempty"`
	Id                           *Int      `xmlrpc:"id,omitempty"`
	IsProcessed                  *Bool     `xmlrpc:"is_processed,omitempty"`
	LocationId                   *Many2One `xmlrpc:"location_id,omitempty"`
	MoveLineIds                  *Relation `xmlrpc:"move_line_ids,omitempty"`
	Name                         *String   `xmlrpc:"name,omitempty"`
	OwnerId                      *Many2One `xmlrpc:"owner_id,omitempty"`
	PackagingId                  *Many2One `xmlrpc:"packaging_id,omitempty"`
	QuantIds                     *Relation `xmlrpc:"quant_ids,omitempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(StockQuantPackageModel, sqp)
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
	return nil, fmt.Errorf("no stock.quant.package was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no stock.quant.package was found with criteria %v and options %v", criteria, options)
}
