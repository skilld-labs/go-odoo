package odoo

import (
	"fmt"
)

// IrQwebFieldFloat represents ir.qweb.field.float model.
type IrQwebFieldFloat struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldFloats represents array of ir.qweb.field.float model.
type IrQwebFieldFloats []IrQwebFieldFloat

// IrQwebFieldFloatModel is the odoo model name.
const IrQwebFieldFloatModel = "ir.qweb.field.float"

// Many2One convert IrQwebFieldFloat to *Many2One.
func (iqff *IrQwebFieldFloat) Many2One() *Many2One {
	return NewMany2One(iqff.Id.Get(), "")
}

// CreateIrQwebFieldFloat creates a new ir.qweb.field.float model and returns its id.
func (c *Client) CreateIrQwebFieldFloat(iqff *IrQwebFieldFloat) (int64, error) {
	return c.Create(IrQwebFieldFloatModel, iqff)
}

// UpdateIrQwebFieldFloat updates an existing ir.qweb.field.float record.
func (c *Client) UpdateIrQwebFieldFloat(iqff *IrQwebFieldFloat) error {
	return c.UpdateIrQwebFieldFloats([]int64{iqff.Id.Get()}, iqff)
}

// UpdateIrQwebFieldFloats updates existing ir.qweb.field.float records.
// All records (represented by ids) will be updated by iqff values.
func (c *Client) UpdateIrQwebFieldFloats(ids []int64, iqff *IrQwebFieldFloat) error {
	return c.Update(IrQwebFieldFloatModel, ids, iqff)
}

// DeleteIrQwebFieldFloat deletes an existing ir.qweb.field.float record.
func (c *Client) DeleteIrQwebFieldFloat(id int64) error {
	return c.DeleteIrQwebFieldFloats([]int64{id})
}

// DeleteIrQwebFieldFloats deletes existing ir.qweb.field.float records.
func (c *Client) DeleteIrQwebFieldFloats(ids []int64) error {
	return c.Delete(IrQwebFieldFloatModel, ids)
}

// GetIrQwebFieldFloat gets ir.qweb.field.float existing record.
func (c *Client) GetIrQwebFieldFloat(id int64) (*IrQwebFieldFloat, error) {
	iqffs, err := c.GetIrQwebFieldFloats([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqffs != nil && len(*iqffs) > 0 {
		return &((*iqffs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.float not found", id)
}

// GetIrQwebFieldFloats gets ir.qweb.field.float existing records.
func (c *Client) GetIrQwebFieldFloats(ids []int64) (*IrQwebFieldFloats, error) {
	iqffs := &IrQwebFieldFloats{}
	if err := c.Read(IrQwebFieldFloatModel, ids, nil, iqffs); err != nil {
		return nil, err
	}
	return iqffs, nil
}

// FindIrQwebFieldFloat finds ir.qweb.field.float record by querying it with criteria.
func (c *Client) FindIrQwebFieldFloat(criteria *Criteria) (*IrQwebFieldFloat, error) {
	iqffs := &IrQwebFieldFloats{}
	if err := c.SearchRead(IrQwebFieldFloatModel, criteria, NewOptions().Limit(1), iqffs); err != nil {
		return nil, err
	}
	if iqffs != nil && len(*iqffs) > 0 {
		return &((*iqffs)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.float was not found")
}

// FindIrQwebFieldFloats finds ir.qweb.field.float records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldFloats(criteria *Criteria, options *Options) (*IrQwebFieldFloats, error) {
	iqffs := &IrQwebFieldFloats{}
	if err := c.SearchRead(IrQwebFieldFloatModel, criteria, options, iqffs); err != nil {
		return nil, err
	}
	return iqffs, nil
}

// FindIrQwebFieldFloatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldFloatIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldFloatModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldFloatId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldFloatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldFloatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.float was not found")
}
