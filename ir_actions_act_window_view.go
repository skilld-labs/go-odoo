package odoo

import (
	"fmt"
)

// IrActionsActWindowView represents ir.actions.act_window.view model.
type IrActionsActWindowView struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	ActWindowId *Many2One  `xmlrpc:"act_window_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Multi       *Bool      `xmlrpc:"multi,omitempty"`
	Sequence    *Int       `xmlrpc:"sequence,omitempty"`
	ViewId      *Many2One  `xmlrpc:"view_id,omitempty"`
	ViewMode    *Selection `xmlrpc:"view_mode,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrActionsActWindowViews represents array of ir.actions.act_window.view model.
type IrActionsActWindowViews []IrActionsActWindowView

// IrActionsActWindowViewModel is the odoo model name.
const IrActionsActWindowViewModel = "ir.actions.act_window.view"

// Many2One convert IrActionsActWindowView to *Many2One.
func (iaav *IrActionsActWindowView) Many2One() *Many2One {
	return NewMany2One(iaav.Id.Get(), "")
}

// CreateIrActionsActWindowView creates a new ir.actions.act_window.view model and returns its id.
func (c *Client) CreateIrActionsActWindowView(iaav *IrActionsActWindowView) (int64, error) {
	return c.Create(IrActionsActWindowViewModel, iaav)
}

// UpdateIrActionsActWindowView updates an existing ir.actions.act_window.view record.
func (c *Client) UpdateIrActionsActWindowView(iaav *IrActionsActWindowView) error {
	return c.UpdateIrActionsActWindowViews([]int64{iaav.Id.Get()}, iaav)
}

// UpdateIrActionsActWindowViews updates existing ir.actions.act_window.view records.
// All records (represented by ids) will be updated by iaav values.
func (c *Client) UpdateIrActionsActWindowViews(ids []int64, iaav *IrActionsActWindowView) error {
	return c.Update(IrActionsActWindowViewModel, ids, iaav)
}

// DeleteIrActionsActWindowView deletes an existing ir.actions.act_window.view record.
func (c *Client) DeleteIrActionsActWindowView(id int64) error {
	return c.DeleteIrActionsActWindowViews([]int64{id})
}

// DeleteIrActionsActWindowViews deletes existing ir.actions.act_window.view records.
func (c *Client) DeleteIrActionsActWindowViews(ids []int64) error {
	return c.Delete(IrActionsActWindowViewModel, ids)
}

// GetIrActionsActWindowView gets ir.actions.act_window.view existing record.
func (c *Client) GetIrActionsActWindowView(id int64) (*IrActionsActWindowView, error) {
	iaavs, err := c.GetIrActionsActWindowViews([]int64{id})
	if err != nil {
		return nil, err
	}
	if iaavs != nil && len(*iaavs) > 0 {
		return &((*iaavs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.act_window.view not found", id)
}

// GetIrActionsActWindowViews gets ir.actions.act_window.view existing records.
func (c *Client) GetIrActionsActWindowViews(ids []int64) (*IrActionsActWindowViews, error) {
	iaavs := &IrActionsActWindowViews{}
	if err := c.Read(IrActionsActWindowViewModel, ids, nil, iaavs); err != nil {
		return nil, err
	}
	return iaavs, nil
}

// FindIrActionsActWindowView finds ir.actions.act_window.view record by querying it with criteria.
func (c *Client) FindIrActionsActWindowView(criteria *Criteria) (*IrActionsActWindowView, error) {
	iaavs := &IrActionsActWindowViews{}
	if err := c.SearchRead(IrActionsActWindowViewModel, criteria, NewOptions().Limit(1), iaavs); err != nil {
		return nil, err
	}
	if iaavs != nil && len(*iaavs) > 0 {
		return &((*iaavs)[0]), nil
	}
	return nil, fmt.Errorf("ir.actions.act_window.view was not found")
}

// FindIrActionsActWindowViews finds ir.actions.act_window.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowViews(criteria *Criteria, options *Options) (*IrActionsActWindowViews, error) {
	iaavs := &IrActionsActWindowViews{}
	if err := c.SearchRead(IrActionsActWindowViewModel, criteria, options, iaavs); err != nil {
		return nil, err
	}
	return iaavs, nil
}

// FindIrActionsActWindowViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsActWindowViewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsActWindowViewId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActWindowViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActWindowViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.actions.act_window.view was not found")
}
