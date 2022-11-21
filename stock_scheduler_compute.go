package odoo

import (
	"fmt"
)

// StockSchedulerCompute represents stock.scheduler.compute model.
type StockSchedulerCompute struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockSchedulerComputes represents array of stock.scheduler.compute model.
type StockSchedulerComputes []StockSchedulerCompute

// StockSchedulerComputeModel is the odoo model name.
const StockSchedulerComputeModel = "stock.scheduler.compute"

// Many2One convert StockSchedulerCompute to *Many2One.
func (ssc *StockSchedulerCompute) Many2One() *Many2One {
	return NewMany2One(ssc.Id.Get(), "")
}

// CreateStockSchedulerCompute creates a new stock.scheduler.compute model and returns its id.
func (c *Client) CreateStockSchedulerCompute(ssc *StockSchedulerCompute) (int64, error) {
	return c.Create(StockSchedulerComputeModel, ssc)
}

// UpdateStockSchedulerCompute updates an existing stock.scheduler.compute record.
func (c *Client) UpdateStockSchedulerCompute(ssc *StockSchedulerCompute) error {
	return c.UpdateStockSchedulerComputes([]int64{ssc.Id.Get()}, ssc)
}

// UpdateStockSchedulerComputes updates existing stock.scheduler.compute records.
// All records (represented by ids) will be updated by ssc values.
func (c *Client) UpdateStockSchedulerComputes(ids []int64, ssc *StockSchedulerCompute) error {
	return c.Update(StockSchedulerComputeModel, ids, ssc)
}

// DeleteStockSchedulerCompute deletes an existing stock.scheduler.compute record.
func (c *Client) DeleteStockSchedulerCompute(id int64) error {
	return c.DeleteStockSchedulerComputes([]int64{id})
}

// DeleteStockSchedulerComputes deletes existing stock.scheduler.compute records.
func (c *Client) DeleteStockSchedulerComputes(ids []int64) error {
	return c.Delete(StockSchedulerComputeModel, ids)
}

// GetStockSchedulerCompute gets stock.scheduler.compute existing record.
func (c *Client) GetStockSchedulerCompute(id int64) (*StockSchedulerCompute, error) {
	sscs, err := c.GetStockSchedulerComputes([]int64{id})
	if err != nil {
		return nil, err
	}
	if sscs != nil && len(*sscs) > 0 {
		return &((*sscs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.scheduler.compute not found", id)
}

// GetStockSchedulerComputes gets stock.scheduler.compute existing records.
func (c *Client) GetStockSchedulerComputes(ids []int64) (*StockSchedulerComputes, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.Read(StockSchedulerComputeModel, ids, nil, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockSchedulerCompute finds stock.scheduler.compute record by querying it with criteria.
func (c *Client) FindStockSchedulerCompute(criteria *Criteria) (*StockSchedulerCompute, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.SearchRead(StockSchedulerComputeModel, criteria, NewOptions().Limit(1), sscs); err != nil {
		return nil, err
	}
	if sscs != nil && len(*sscs) > 0 {
		return &((*sscs)[0]), nil
	}
	return nil, fmt.Errorf("no stock.scheduler.compute was found with criteria %v", criteria)
}

// FindStockSchedulerComputes finds stock.scheduler.compute records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockSchedulerComputes(criteria *Criteria, options *Options) (*StockSchedulerComputes, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.SearchRead(StockSchedulerComputeModel, criteria, options, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockSchedulerComputeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockSchedulerComputeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockSchedulerComputeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockSchedulerComputeId finds record id by querying it with criteria.
func (c *Client) FindStockSchedulerComputeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockSchedulerComputeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.scheduler.compute was found with criteria %v and options %v", criteria, options)
}
