package odoo

import (
	"fmt"
)

// IrModelFieldsSelection represents ir.model.fields.selection model.
type IrModelFieldsSelection struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FieldId     *Many2One `xmlrpc:"field_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrModelFieldsSelections represents array of ir.model.fields.selection model.
type IrModelFieldsSelections []IrModelFieldsSelection

// IrModelFieldsSelectionModel is the odoo model name.
const IrModelFieldsSelectionModel = "ir.model.fields.selection"

// Many2One convert IrModelFieldsSelection to *Many2One.
func (imfs *IrModelFieldsSelection) Many2One() *Many2One {
	return NewMany2One(imfs.Id.Get(), "")
}

// CreateIrModelFieldsSelection creates a new ir.model.fields.selection model and returns its id.
func (c *Client) CreateIrModelFieldsSelection(imfs *IrModelFieldsSelection) (int64, error) {
	return c.Create(IrModelFieldsSelectionModel, imfs)
}

// UpdateIrModelFieldsSelection updates an existing ir.model.fields.selection record.
func (c *Client) UpdateIrModelFieldsSelection(imfs *IrModelFieldsSelection) error {
	return c.UpdateIrModelFieldsSelections([]int64{imfs.Id.Get()}, imfs)
}

// UpdateIrModelFieldsSelections updates existing ir.model.fields.selection records.
// All records (represented by IDs) will be updated by imfs values.
func (c *Client) UpdateIrModelFieldsSelections(ids []int64, imfs *IrModelFieldsSelection) error {
	return c.Update(IrModelFieldsSelectionModel, ids, imfs)
}

// DeleteIrModelFieldsSelection deletes an existing ir.model.fields.selection record.
func (c *Client) DeleteIrModelFieldsSelection(id int64) error {
	return c.DeleteIrModelFieldsSelections([]int64{id})
}

// DeleteIrModelFieldsSelections deletes existing ir.model.fields.selection records.
func (c *Client) DeleteIrModelFieldsSelections(ids []int64) error {
	return c.Delete(IrModelFieldsSelectionModel, ids)
}

// GetIrModelFieldsSelection gets ir.model.fields.selection existing record.
func (c *Client) GetIrModelFieldsSelection(id int64) (*IrModelFieldsSelection, error) {
	imfss, err := c.GetIrModelFieldsSelections([]int64{id})
	if err != nil {
		return nil, err
	}
	if imfss != nil && len(*imfss) > 0 {
		return &((*imfss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.model.fields.selection not found", id)
}

// GetIrModelFieldsSelections gets ir.model.fields.selection existing records.
func (c *Client) GetIrModelFieldsSelections(ids []int64) (*IrModelFieldsSelections, error) {
	imfss := &IrModelFieldsSelections{}
	if err := c.Read(IrModelFieldsSelectionModel, ids, nil, imfss); err != nil {
		return nil, err
	}
	return imfss, nil
}

// FindIrModelFieldsSelection finds ir.model.fields.selection record by querying it with criteria.
func (c *Client) FindIrModelFieldsSelection(criteria *Criteria) (*IrModelFieldsSelection, error) {
	imfss := &IrModelFieldsSelections{}
	if err := c.SearchRead(IrModelFieldsSelectionModel, criteria, NewOptions().Limit(1), imfss); err != nil {
		return nil, err
	}
	if imfss != nil && len(*imfss) > 0 {
		return &((*imfss)[0]), nil
	}
	return nil, fmt.Errorf("ir.model.fields.selection was not found")
}

// FindIrModelFieldsSelections finds ir.model.fields.selection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldsSelections(criteria *Criteria, options *Options) (*IrModelFieldsSelections, error) {
	imfss := &IrModelFieldsSelections{}
	if err := c.SearchRead(IrModelFieldsSelectionModel, criteria, options, imfss); err != nil {
		return nil, err
	}
	return imfss, nil
}

// FindIrModelFieldsSelectionIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldsSelectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModelFieldsSelectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModelFieldsSelectionId finds record id by querying it with criteria.
func (c *Client) FindIrModelFieldsSelectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelFieldsSelectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.model.fields.selection was not found")
}
