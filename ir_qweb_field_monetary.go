package odoo

import (
	"fmt"
)

// IrQwebFieldMonetary represents ir.qweb.field.monetary model.
type IrQwebFieldMonetary struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldMonetarys represents array of ir.qweb.field.monetary model.
type IrQwebFieldMonetarys []IrQwebFieldMonetary

// IrQwebFieldMonetaryModel is the odoo model name.
const IrQwebFieldMonetaryModel = "ir.qweb.field.monetary"

// Many2One convert IrQwebFieldMonetary to *Many2One.
func (iqfm *IrQwebFieldMonetary) Many2One() *Many2One {
	return NewMany2One(iqfm.Id.Get(), "")
}

// CreateIrQwebFieldMonetary creates a new ir.qweb.field.monetary model and returns its id.
func (c *Client) CreateIrQwebFieldMonetary(iqfm *IrQwebFieldMonetary) (int64, error) {
	return c.Create(IrQwebFieldMonetaryModel, iqfm)
}

// UpdateIrQwebFieldMonetary updates an existing ir.qweb.field.monetary record.
func (c *Client) UpdateIrQwebFieldMonetary(iqfm *IrQwebFieldMonetary) error {
	return c.UpdateIrQwebFieldMonetarys([]int64{iqfm.Id.Get()}, iqfm)
}

// UpdateIrQwebFieldMonetarys updates existing ir.qweb.field.monetary records.
// All records (represented by ids) will be updated by iqfm values.
func (c *Client) UpdateIrQwebFieldMonetarys(ids []int64, iqfm *IrQwebFieldMonetary) error {
	return c.Update(IrQwebFieldMonetaryModel, ids, iqfm)
}

// DeleteIrQwebFieldMonetary deletes an existing ir.qweb.field.monetary record.
func (c *Client) DeleteIrQwebFieldMonetary(id int64) error {
	return c.DeleteIrQwebFieldMonetarys([]int64{id})
}

// DeleteIrQwebFieldMonetarys deletes existing ir.qweb.field.monetary records.
func (c *Client) DeleteIrQwebFieldMonetarys(ids []int64) error {
	return c.Delete(IrQwebFieldMonetaryModel, ids)
}

// GetIrQwebFieldMonetary gets ir.qweb.field.monetary existing record.
func (c *Client) GetIrQwebFieldMonetary(id int64) (*IrQwebFieldMonetary, error) {
	iqfms, err := c.GetIrQwebFieldMonetarys([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.monetary not found", id)
}

// GetIrQwebFieldMonetarys gets ir.qweb.field.monetary existing records.
func (c *Client) GetIrQwebFieldMonetarys(ids []int64) (*IrQwebFieldMonetarys, error) {
	iqfms := &IrQwebFieldMonetarys{}
	if err := c.Read(IrQwebFieldMonetaryModel, ids, nil, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMonetary finds ir.qweb.field.monetary record by querying it with criteria.
func (c *Client) FindIrQwebFieldMonetary(criteria *Criteria) (*IrQwebFieldMonetary, error) {
	iqfms := &IrQwebFieldMonetarys{}
	if err := c.SearchRead(IrQwebFieldMonetaryModel, criteria, NewOptions().Limit(1), iqfms); err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.monetary was not found")
}

// FindIrQwebFieldMonetarys finds ir.qweb.field.monetary records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMonetarys(criteria *Criteria, options *Options) (*IrQwebFieldMonetarys, error) {
	iqfms := &IrQwebFieldMonetarys{}
	if err := c.SearchRead(IrQwebFieldMonetaryModel, criteria, options, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMonetaryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMonetaryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldMonetaryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldMonetaryId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldMonetaryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldMonetaryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.monetary was not found")
}
