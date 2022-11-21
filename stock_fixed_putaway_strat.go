package odoo

import (
	"fmt"
)

// StockFixedPutawayStrat represents stock.fixed.putaway.strat model.
type StockFixedPutawayStrat struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	FixedLocationId *Many2One `xmlrpc:"fixed_location_id,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	PutawayId       *Many2One `xmlrpc:"putaway_id,omptempty"`
	Sequence        *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockFixedPutawayStrats represents array of stock.fixed.putaway.strat model.
type StockFixedPutawayStrats []StockFixedPutawayStrat

// StockFixedPutawayStratModel is the odoo model name.
const StockFixedPutawayStratModel = "stock.fixed.putaway.strat"

// Many2One convert StockFixedPutawayStrat to *Many2One.
func (sfps *StockFixedPutawayStrat) Many2One() *Many2One {
	return NewMany2One(sfps.Id.Get(), "")
}

// CreateStockFixedPutawayStrat creates a new stock.fixed.putaway.strat model and returns its id.
func (c *Client) CreateStockFixedPutawayStrat(sfps *StockFixedPutawayStrat) (int64, error) {
	return c.Create(StockFixedPutawayStratModel, sfps)
}

// UpdateStockFixedPutawayStrat updates an existing stock.fixed.putaway.strat record.
func (c *Client) UpdateStockFixedPutawayStrat(sfps *StockFixedPutawayStrat) error {
	return c.UpdateStockFixedPutawayStrats([]int64{sfps.Id.Get()}, sfps)
}

// UpdateStockFixedPutawayStrats updates existing stock.fixed.putaway.strat records.
// All records (represented by ids) will be updated by sfps values.
func (c *Client) UpdateStockFixedPutawayStrats(ids []int64, sfps *StockFixedPutawayStrat) error {
	return c.Update(StockFixedPutawayStratModel, ids, sfps)
}

// DeleteStockFixedPutawayStrat deletes an existing stock.fixed.putaway.strat record.
func (c *Client) DeleteStockFixedPutawayStrat(id int64) error {
	return c.DeleteStockFixedPutawayStrats([]int64{id})
}

// DeleteStockFixedPutawayStrats deletes existing stock.fixed.putaway.strat records.
func (c *Client) DeleteStockFixedPutawayStrats(ids []int64) error {
	return c.Delete(StockFixedPutawayStratModel, ids)
}

// GetStockFixedPutawayStrat gets stock.fixed.putaway.strat existing record.
func (c *Client) GetStockFixedPutawayStrat(id int64) (*StockFixedPutawayStrat, error) {
	sfpss, err := c.GetStockFixedPutawayStrats([]int64{id})
	if err != nil {
		return nil, err
	}
	if sfpss != nil && len(*sfpss) > 0 {
		return &((*sfpss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.fixed.putaway.strat not found", id)
}

// GetStockFixedPutawayStrats gets stock.fixed.putaway.strat existing records.
func (c *Client) GetStockFixedPutawayStrats(ids []int64) (*StockFixedPutawayStrats, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.Read(StockFixedPutawayStratModel, ids, nil, sfpss); err != nil {
		return nil, err
	}
	return sfpss, nil
}

// FindStockFixedPutawayStrat finds stock.fixed.putaway.strat record by querying it with criteria.
func (c *Client) FindStockFixedPutawayStrat(criteria *Criteria) (*StockFixedPutawayStrat, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.SearchRead(StockFixedPutawayStratModel, criteria, NewOptions().Limit(1), sfpss); err != nil {
		return nil, err
	}
	if sfpss != nil && len(*sfpss) > 0 {
		return &((*sfpss)[0]), nil
	}
	return nil, fmt.Errorf("no stock.fixed.putaway.strat was found with criteria %v", criteria)
}

// FindStockFixedPutawayStrats finds stock.fixed.putaway.strat records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockFixedPutawayStrats(criteria *Criteria, options *Options) (*StockFixedPutawayStrats, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.SearchRead(StockFixedPutawayStratModel, criteria, options, sfpss); err != nil {
		return nil, err
	}
	return sfpss, nil
}

// FindStockFixedPutawayStratIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockFixedPutawayStratIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockFixedPutawayStratModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockFixedPutawayStratId finds record id by querying it with criteria.
func (c *Client) FindStockFixedPutawayStratId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockFixedPutawayStratModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.fixed.putaway.strat was found with criteria %v and options %v", criteria, options)
}
