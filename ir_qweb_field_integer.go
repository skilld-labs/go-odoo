package odoo

import (
	"fmt"
)

// IrQwebFieldInteger represents ir.qweb.field.integer model.
type IrQwebFieldInteger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldIntegers represents array of ir.qweb.field.integer model.
type IrQwebFieldIntegers []IrQwebFieldInteger

// IrQwebFieldIntegerModel is the odoo model name.
const IrQwebFieldIntegerModel = "ir.qweb.field.integer"

// Many2One convert IrQwebFieldInteger to *Many2One.
func (iqfi *IrQwebFieldInteger) Many2One() *Many2One {
	return NewMany2One(iqfi.Id.Get(), "")
}

// CreateIrQwebFieldInteger creates a new ir.qweb.field.integer model and returns its id.
func (c *Client) CreateIrQwebFieldInteger(iqfi *IrQwebFieldInteger) (int64, error) {
	return c.Create(IrQwebFieldIntegerModel, iqfi)
}

// UpdateIrQwebFieldInteger updates an existing ir.qweb.field.integer record.
func (c *Client) UpdateIrQwebFieldInteger(iqfi *IrQwebFieldInteger) error {
	return c.UpdateIrQwebFieldIntegers([]int64{iqfi.Id.Get()}, iqfi)
}

// UpdateIrQwebFieldIntegers updates existing ir.qweb.field.integer records.
// All records (represented by ids) will be updated by iqfi values.
func (c *Client) UpdateIrQwebFieldIntegers(ids []int64, iqfi *IrQwebFieldInteger) error {
	return c.Update(IrQwebFieldIntegerModel, ids, iqfi)
}

// DeleteIrQwebFieldInteger deletes an existing ir.qweb.field.integer record.
func (c *Client) DeleteIrQwebFieldInteger(id int64) error {
	return c.DeleteIrQwebFieldIntegers([]int64{id})
}

// DeleteIrQwebFieldIntegers deletes existing ir.qweb.field.integer records.
func (c *Client) DeleteIrQwebFieldIntegers(ids []int64) error {
	return c.Delete(IrQwebFieldIntegerModel, ids)
}

// GetIrQwebFieldInteger gets ir.qweb.field.integer existing record.
func (c *Client) GetIrQwebFieldInteger(id int64) (*IrQwebFieldInteger, error) {
	iqfis, err := c.GetIrQwebFieldIntegers([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.integer not found", id)
}

// GetIrQwebFieldIntegers gets ir.qweb.field.integer existing records.
func (c *Client) GetIrQwebFieldIntegers(ids []int64) (*IrQwebFieldIntegers, error) {
	iqfis := &IrQwebFieldIntegers{}
	if err := c.Read(IrQwebFieldIntegerModel, ids, nil, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldInteger finds ir.qweb.field.integer record by querying it with criteria.
func (c *Client) FindIrQwebFieldInteger(criteria *Criteria) (*IrQwebFieldInteger, error) {
	iqfis := &IrQwebFieldIntegers{}
	if err := c.SearchRead(IrQwebFieldIntegerModel, criteria, NewOptions().Limit(1), iqfis); err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.integer was not found")
}

// FindIrQwebFieldIntegers finds ir.qweb.field.integer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldIntegers(criteria *Criteria, options *Options) (*IrQwebFieldIntegers, error) {
	iqfis := &IrQwebFieldIntegers{}
	if err := c.SearchRead(IrQwebFieldIntegerModel, criteria, options, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldIntegerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldIntegerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldIntegerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldIntegerId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldIntegerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldIntegerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.integer was not found")
}
