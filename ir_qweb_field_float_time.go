package odoo

import (
	"fmt"
)

// IrQwebFieldFloatTime represents ir.qweb.field.float_time model.
type IrQwebFieldFloatTime struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldFloatTimes represents array of ir.qweb.field.float_time model.
type IrQwebFieldFloatTimes []IrQwebFieldFloatTime

// IrQwebFieldFloatTimeModel is the odoo model name.
const IrQwebFieldFloatTimeModel = "ir.qweb.field.float_time"

// Many2One convert IrQwebFieldFloatTime to *Many2One.
func (iqff *IrQwebFieldFloatTime) Many2One() *Many2One {
	return NewMany2One(iqff.Id.Get(), "")
}

// CreateIrQwebFieldFloatTime creates a new ir.qweb.field.float_time model and returns its id.
func (c *Client) CreateIrQwebFieldFloatTime(iqff *IrQwebFieldFloatTime) (int64, error) {
	return c.Create(IrQwebFieldFloatTimeModel, iqff)
}

// UpdateIrQwebFieldFloatTime updates an existing ir.qweb.field.float_time record.
func (c *Client) UpdateIrQwebFieldFloatTime(iqff *IrQwebFieldFloatTime) error {
	return c.UpdateIrQwebFieldFloatTimes([]int64{iqff.Id.Get()}, iqff)
}

// UpdateIrQwebFieldFloatTimes updates existing ir.qweb.field.float_time records.
// All records (represented by ids) will be updated by iqff values.
func (c *Client) UpdateIrQwebFieldFloatTimes(ids []int64, iqff *IrQwebFieldFloatTime) error {
	return c.Update(IrQwebFieldFloatTimeModel, ids, iqff)
}

// DeleteIrQwebFieldFloatTime deletes an existing ir.qweb.field.float_time record.
func (c *Client) DeleteIrQwebFieldFloatTime(id int64) error {
	return c.DeleteIrQwebFieldFloatTimes([]int64{id})
}

// DeleteIrQwebFieldFloatTimes deletes existing ir.qweb.field.float_time records.
func (c *Client) DeleteIrQwebFieldFloatTimes(ids []int64) error {
	return c.Delete(IrQwebFieldFloatTimeModel, ids)
}

// GetIrQwebFieldFloatTime gets ir.qweb.field.float_time existing record.
func (c *Client) GetIrQwebFieldFloatTime(id int64) (*IrQwebFieldFloatTime, error) {
	iqffs, err := c.GetIrQwebFieldFloatTimes([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqffs != nil && len(*iqffs) > 0 {
		return &((*iqffs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.float_time not found", id)
}

// GetIrQwebFieldFloatTimes gets ir.qweb.field.float_time existing records.
func (c *Client) GetIrQwebFieldFloatTimes(ids []int64) (*IrQwebFieldFloatTimes, error) {
	iqffs := &IrQwebFieldFloatTimes{}
	if err := c.Read(IrQwebFieldFloatTimeModel, ids, nil, iqffs); err != nil {
		return nil, err
	}
	return iqffs, nil
}

// FindIrQwebFieldFloatTime finds ir.qweb.field.float_time record by querying it with criteria.
func (c *Client) FindIrQwebFieldFloatTime(criteria *Criteria) (*IrQwebFieldFloatTime, error) {
	iqffs := &IrQwebFieldFloatTimes{}
	if err := c.SearchRead(IrQwebFieldFloatTimeModel, criteria, NewOptions().Limit(1), iqffs); err != nil {
		return nil, err
	}
	if iqffs != nil && len(*iqffs) > 0 {
		return &((*iqffs)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.float_time was not found")
}

// FindIrQwebFieldFloatTimes finds ir.qweb.field.float_time records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldFloatTimes(criteria *Criteria, options *Options) (*IrQwebFieldFloatTimes, error) {
	iqffs := &IrQwebFieldFloatTimes{}
	if err := c.SearchRead(IrQwebFieldFloatTimeModel, criteria, options, iqffs); err != nil {
		return nil, err
	}
	return iqffs, nil
}

// FindIrQwebFieldFloatTimeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldFloatTimeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldFloatTimeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldFloatTimeId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldFloatTimeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldFloatTimeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.float_time was not found")
}
