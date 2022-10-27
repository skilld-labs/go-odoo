package odoo

import (
	"fmt"
)

// IrQwebFieldDuration represents ir.qweb.field.duration model.
type IrQwebFieldDuration struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldDurations represents array of ir.qweb.field.duration model.
type IrQwebFieldDurations []IrQwebFieldDuration

// IrQwebFieldDurationModel is the odoo model name.
const IrQwebFieldDurationModel = "ir.qweb.field.duration"

// Many2One convert IrQwebFieldDuration to *Many2One.
func (iqfd *IrQwebFieldDuration) Many2One() *Many2One {
	return NewMany2One(iqfd.Id.Get(), "")
}

// CreateIrQwebFieldDuration creates a new ir.qweb.field.duration model and returns its id.
func (c *Client) CreateIrQwebFieldDuration(iqfd *IrQwebFieldDuration) (int64, error) {
	return c.Create(IrQwebFieldDurationModel, iqfd)
}

// UpdateIrQwebFieldDuration updates an existing ir.qweb.field.duration record.
func (c *Client) UpdateIrQwebFieldDuration(iqfd *IrQwebFieldDuration) error {
	return c.UpdateIrQwebFieldDurations([]int64{iqfd.Id.Get()}, iqfd)
}

// UpdateIrQwebFieldDurations updates existing ir.qweb.field.duration records.
// All records (represented by ids) will be updated by iqfd values.
func (c *Client) UpdateIrQwebFieldDurations(ids []int64, iqfd *IrQwebFieldDuration) error {
	return c.Update(IrQwebFieldDurationModel, ids, iqfd)
}

// DeleteIrQwebFieldDuration deletes an existing ir.qweb.field.duration record.
func (c *Client) DeleteIrQwebFieldDuration(id int64) error {
	return c.DeleteIrQwebFieldDurations([]int64{id})
}

// DeleteIrQwebFieldDurations deletes existing ir.qweb.field.duration records.
func (c *Client) DeleteIrQwebFieldDurations(ids []int64) error {
	return c.Delete(IrQwebFieldDurationModel, ids)
}

// GetIrQwebFieldDuration gets ir.qweb.field.duration existing record.
func (c *Client) GetIrQwebFieldDuration(id int64) (*IrQwebFieldDuration, error) {
	iqfds, err := c.GetIrQwebFieldDurations([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfds != nil && len(*iqfds) > 0 {
		return &((*iqfds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.duration not found", id)
}

// GetIrQwebFieldDurations gets ir.qweb.field.duration existing records.
func (c *Client) GetIrQwebFieldDurations(ids []int64) (*IrQwebFieldDurations, error) {
	iqfds := &IrQwebFieldDurations{}
	if err := c.Read(IrQwebFieldDurationModel, ids, nil, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDuration finds ir.qweb.field.duration record by querying it with criteria.
func (c *Client) FindIrQwebFieldDuration(criteria *Criteria) (*IrQwebFieldDuration, error) {
	iqfds := &IrQwebFieldDurations{}
	if err := c.SearchRead(IrQwebFieldDurationModel, criteria, NewOptions().Limit(1), iqfds); err != nil {
		return nil, err
	}
	if iqfds != nil && len(*iqfds) > 0 {
		return &((*iqfds)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.duration was not found")
}

// FindIrQwebFieldDurations finds ir.qweb.field.duration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDurations(criteria *Criteria, options *Options) (*IrQwebFieldDurations, error) {
	iqfds := &IrQwebFieldDurations{}
	if err := c.SearchRead(IrQwebFieldDurationModel, criteria, options, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDurationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDurationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldDurationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldDurationId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldDurationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldDurationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.duration was not found")
}
