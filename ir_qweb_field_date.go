package odoo

import (
	"fmt"
)

// IrQwebFieldDate represents ir.qweb.field.date model.
type IrQwebFieldDate struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrQwebFieldDates represents array of ir.qweb.field.date model.
type IrQwebFieldDates []IrQwebFieldDate

// IrQwebFieldDateModel is the odoo model name.
const IrQwebFieldDateModel = "ir.qweb.field.date"

// Many2One convert IrQwebFieldDate to *Many2One.
func (iqfd *IrQwebFieldDate) Many2One() *Many2One {
	return NewMany2One(iqfd.Id.Get(), "")
}

// CreateIrQwebFieldDate creates a new ir.qweb.field.date model and returns its id.
func (c *Client) CreateIrQwebFieldDate(iqfd *IrQwebFieldDate) (int64, error) {
	return c.Create(IrQwebFieldDateModel, iqfd)
}

// UpdateIrQwebFieldDate updates an existing ir.qweb.field.date record.
func (c *Client) UpdateIrQwebFieldDate(iqfd *IrQwebFieldDate) error {
	return c.UpdateIrQwebFieldDates([]int64{iqfd.Id.Get()}, iqfd)
}

// UpdateIrQwebFieldDates updates existing ir.qweb.field.date records.
// All records (represented by ids) will be updated by iqfd values.
func (c *Client) UpdateIrQwebFieldDates(ids []int64, iqfd *IrQwebFieldDate) error {
	return c.Update(IrQwebFieldDateModel, ids, iqfd)
}

// DeleteIrQwebFieldDate deletes an existing ir.qweb.field.date record.
func (c *Client) DeleteIrQwebFieldDate(id int64) error {
	return c.DeleteIrQwebFieldDates([]int64{id})
}

// DeleteIrQwebFieldDates deletes existing ir.qweb.field.date records.
func (c *Client) DeleteIrQwebFieldDates(ids []int64) error {
	return c.Delete(IrQwebFieldDateModel, ids)
}

// GetIrQwebFieldDate gets ir.qweb.field.date existing record.
func (c *Client) GetIrQwebFieldDate(id int64) (*IrQwebFieldDate, error) {
	iqfds, err := c.GetIrQwebFieldDates([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfds != nil && len(*iqfds) > 0 {
		return &((*iqfds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.date not found", id)
}

// GetIrQwebFieldDates gets ir.qweb.field.date existing records.
func (c *Client) GetIrQwebFieldDates(ids []int64) (*IrQwebFieldDates, error) {
	iqfds := &IrQwebFieldDates{}
	if err := c.Read(IrQwebFieldDateModel, ids, nil, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDate finds ir.qweb.field.date record by querying it with criteria.
func (c *Client) FindIrQwebFieldDate(criteria *Criteria) (*IrQwebFieldDate, error) {
	iqfds := &IrQwebFieldDates{}
	if err := c.SearchRead(IrQwebFieldDateModel, criteria, NewOptions().Limit(1), iqfds); err != nil {
		return nil, err
	}
	if iqfds != nil && len(*iqfds) > 0 {
		return &((*iqfds)[0]), nil
	}
	return nil, fmt.Errorf("no ir.qweb.field.date was found with criteria %v", criteria)
}

// FindIrQwebFieldDates finds ir.qweb.field.date records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDates(criteria *Criteria, options *Options) (*IrQwebFieldDates, error) {
	iqfds := &IrQwebFieldDates{}
	if err := c.SearchRead(IrQwebFieldDateModel, criteria, options, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldDateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldDateId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldDateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldDateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.qweb.field.date was found with criteria %v and options %v", criteria, options)
}
