package odoo

import (
	"fmt"
)

// IrQwebFieldSelection represents ir.qweb.field.selection model.
type IrQwebFieldSelection struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldSelections represents array of ir.qweb.field.selection model.
type IrQwebFieldSelections []IrQwebFieldSelection

// IrQwebFieldSelectionModel is the odoo model name.
const IrQwebFieldSelectionModel = "ir.qweb.field.selection"

// Many2One convert IrQwebFieldSelection to *Many2One.
func (iqfs *IrQwebFieldSelection) Many2One() *Many2One {
	return NewMany2One(iqfs.Id.Get(), "")
}

// CreateIrQwebFieldSelection creates a new ir.qweb.field.selection model and returns its id.
func (c *Client) CreateIrQwebFieldSelection(iqfs *IrQwebFieldSelection) (int64, error) {
	return c.Create(IrQwebFieldSelectionModel, iqfs)
}

// UpdateIrQwebFieldSelection updates an existing ir.qweb.field.selection record.
func (c *Client) UpdateIrQwebFieldSelection(iqfs *IrQwebFieldSelection) error {
	return c.UpdateIrQwebFieldSelections([]int64{iqfs.Id.Get()}, iqfs)
}

// UpdateIrQwebFieldSelections updates existing ir.qweb.field.selection records.
// All records (represented by ids) will be updated by iqfs values.
func (c *Client) UpdateIrQwebFieldSelections(ids []int64, iqfs *IrQwebFieldSelection) error {
	return c.Update(IrQwebFieldSelectionModel, ids, iqfs)
}

// DeleteIrQwebFieldSelection deletes an existing ir.qweb.field.selection record.
func (c *Client) DeleteIrQwebFieldSelection(id int64) error {
	return c.DeleteIrQwebFieldSelections([]int64{id})
}

// DeleteIrQwebFieldSelections deletes existing ir.qweb.field.selection records.
func (c *Client) DeleteIrQwebFieldSelections(ids []int64) error {
	return c.Delete(IrQwebFieldSelectionModel, ids)
}

// GetIrQwebFieldSelection gets ir.qweb.field.selection existing record.
func (c *Client) GetIrQwebFieldSelection(id int64) (*IrQwebFieldSelection, error) {
	iqfss, err := c.GetIrQwebFieldSelections([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfss != nil && len(*iqfss) > 0 {
		return &((*iqfss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.selection not found", id)
}

// GetIrQwebFieldSelections gets ir.qweb.field.selection existing records.
func (c *Client) GetIrQwebFieldSelections(ids []int64) (*IrQwebFieldSelections, error) {
	iqfss := &IrQwebFieldSelections{}
	if err := c.Read(IrQwebFieldSelectionModel, ids, nil, iqfss); err != nil {
		return nil, err
	}
	return iqfss, nil
}

// FindIrQwebFieldSelection finds ir.qweb.field.selection record by querying it with criteria.
func (c *Client) FindIrQwebFieldSelection(criteria *Criteria) (*IrQwebFieldSelection, error) {
	iqfss := &IrQwebFieldSelections{}
	if err := c.SearchRead(IrQwebFieldSelectionModel, criteria, NewOptions().Limit(1), iqfss); err != nil {
		return nil, err
	}
	if iqfss != nil && len(*iqfss) > 0 {
		return &((*iqfss)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.selection was not found")
}

// FindIrQwebFieldSelections finds ir.qweb.field.selection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldSelections(criteria *Criteria, options *Options) (*IrQwebFieldSelections, error) {
	iqfss := &IrQwebFieldSelections{}
	if err := c.SearchRead(IrQwebFieldSelectionModel, criteria, options, iqfss); err != nil {
		return nil, err
	}
	return iqfss, nil
}

// FindIrQwebFieldSelectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldSelectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldSelectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldSelectionId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldSelectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldSelectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.selection was not found")
}
