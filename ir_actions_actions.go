package odoo

import (
	"fmt"
)

// IrActionsActions represents ir.actions.actions model.
type IrActionsActions struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	BindingModelId   *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType      *Selection `xmlrpc:"binding_type,omitempty"`
	BindingViewTypes *String    `xmlrpc:"binding_view_types,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Help             *String    `xmlrpc:"help,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	Type             *String    `xmlrpc:"type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId            *String    `xmlrpc:"xml_id,omitempty"`
}

// IrActionsActionss represents array of ir.actions.actions model.
type IrActionsActionss []IrActionsActions

// IrActionsActionsModel is the odoo model name.
const IrActionsActionsModel = "ir.actions.actions"

// Many2One convert IrActionsActions to *Many2One.
func (iaa *IrActionsActions) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActions creates a new ir.actions.actions model and returns its id.
func (c *Client) CreateIrActionsActions(iaa *IrActionsActions) (int64, error) {
	return c.Create(IrActionsActionsModel, iaa)
}

// UpdateIrActionsActions updates an existing ir.actions.actions record.
func (c *Client) UpdateIrActionsActions(iaa *IrActionsActions) error {
	return c.UpdateIrActionsActionss([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActionss updates existing ir.actions.actions records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActionss(ids []int64, iaa *IrActionsActions) error {
	return c.Update(IrActionsActionsModel, ids, iaa)
}

// DeleteIrActionsActions deletes an existing ir.actions.actions record.
func (c *Client) DeleteIrActionsActions(id int64) error {
	return c.DeleteIrActionsActionss([]int64{id})
}

// DeleteIrActionsActionss deletes existing ir.actions.actions records.
func (c *Client) DeleteIrActionsActionss(ids []int64) error {
	return c.Delete(IrActionsActionsModel, ids)
}

// GetIrActionsActions gets ir.actions.actions existing record.
func (c *Client) GetIrActionsActions(id int64) (*IrActionsActions, error) {
	iaas, err := c.GetIrActionsActionss([]int64{id})
	if err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.actions not found", id)
}

// GetIrActionsActionss gets ir.actions.actions existing records.
func (c *Client) GetIrActionsActionss(ids []int64) (*IrActionsActionss, error) {
	iaas := &IrActionsActionss{}
	if err := c.Read(IrActionsActionsModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActions finds ir.actions.actions record by querying it with criteria.
func (c *Client) FindIrActionsActions(criteria *Criteria) (*IrActionsActions, error) {
	iaas := &IrActionsActionss{}
	if err := c.SearchRead(IrActionsActionsModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("ir.actions.actions was not found")
}

// FindIrActionsActionss finds ir.actions.actions records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActionss(criteria *Criteria, options *Options) (*IrActionsActionss, error) {
	iaas := &IrActionsActionss{}
	if err := c.SearchRead(IrActionsActionsModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActionsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActionsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsActionsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsActionsId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActionsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActionsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.actions.actions was not found")
}
