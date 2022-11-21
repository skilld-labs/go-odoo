package odoo

import (
	"fmt"
)

// IrActionsActWindow represents ir.actions.act_window model.
type IrActionsActWindow struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AutoSearch     *Bool      `xmlrpc:"auto_search,omptempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omptempty"`
	Context        *String    `xmlrpc:"context,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Domain         *String    `xmlrpc:"domain,omptempty"`
	Filter         *Bool      `xmlrpc:"filter,omptempty"`
	GroupsId       *Relation  `xmlrpc:"groups_id,omptempty"`
	Help           *String    `xmlrpc:"help,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Limit          *Int       `xmlrpc:"limit,omptempty"`
	Multi          *Bool      `xmlrpc:"multi,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	ResId          *Int       `xmlrpc:"res_id,omptempty"`
	ResModel       *String    `xmlrpc:"res_model,omptempty"`
	SearchView     *String    `xmlrpc:"search_view,omptempty"`
	SearchViewId   *Many2One  `xmlrpc:"search_view_id,omptempty"`
	SrcModel       *String    `xmlrpc:"src_model,omptempty"`
	Target         *Selection `xmlrpc:"target,omptempty"`
	Type           *String    `xmlrpc:"type,omptempty"`
	Usage          *String    `xmlrpc:"usage,omptempty"`
	ViewId         *Many2One  `xmlrpc:"view_id,omptempty"`
	ViewIds        *Relation  `xmlrpc:"view_ids,omptempty"`
	ViewMode       *String    `xmlrpc:"view_mode,omptempty"`
	ViewType       *Selection `xmlrpc:"view_type,omptempty"`
	Views          *String    `xmlrpc:"views,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId          *String    `xmlrpc:"xml_id,omptempty"`
}

// IrActionsActWindows represents array of ir.actions.act_window model.
type IrActionsActWindows []IrActionsActWindow

// IrActionsActWindowModel is the odoo model name.
const IrActionsActWindowModel = "ir.actions.act_window"

// Many2One convert IrActionsActWindow to *Many2One.
func (iaa *IrActionsActWindow) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActWindow creates a new ir.actions.act_window model and returns its id.
func (c *Client) CreateIrActionsActWindow(iaa *IrActionsActWindow) (int64, error) {
	return c.Create(IrActionsActWindowModel, iaa)
}

// UpdateIrActionsActWindow updates an existing ir.actions.act_window record.
func (c *Client) UpdateIrActionsActWindow(iaa *IrActionsActWindow) error {
	return c.UpdateIrActionsActWindows([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActWindows updates existing ir.actions.act_window records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActWindows(ids []int64, iaa *IrActionsActWindow) error {
	return c.Update(IrActionsActWindowModel, ids, iaa)
}

// DeleteIrActionsActWindow deletes an existing ir.actions.act_window record.
func (c *Client) DeleteIrActionsActWindow(id int64) error {
	return c.DeleteIrActionsActWindows([]int64{id})
}

// DeleteIrActionsActWindows deletes existing ir.actions.act_window records.
func (c *Client) DeleteIrActionsActWindows(ids []int64) error {
	return c.Delete(IrActionsActWindowModel, ids)
}

// GetIrActionsActWindow gets ir.actions.act_window existing record.
func (c *Client) GetIrActionsActWindow(id int64) (*IrActionsActWindow, error) {
	iaas, err := c.GetIrActionsActWindows([]int64{id})
	if err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.act_window not found", id)
}

// GetIrActionsActWindows gets ir.actions.act_window existing records.
func (c *Client) GetIrActionsActWindows(ids []int64) (*IrActionsActWindows, error) {
	iaas := &IrActionsActWindows{}
	if err := c.Read(IrActionsActWindowModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindow finds ir.actions.act_window record by querying it with criteria.
func (c *Client) FindIrActionsActWindow(criteria *Criteria) (*IrActionsActWindow, error) {
	iaas := &IrActionsActWindows{}
	if err := c.SearchRead(IrActionsActWindowModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("no ir.actions.act_window was found with criteria %v", criteria)
}

// FindIrActionsActWindows finds ir.actions.act_window records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindows(criteria *Criteria, options *Options) (*IrActionsActWindows, error) {
	iaas := &IrActionsActWindows{}
	if err := c.SearchRead(IrActionsActWindowModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindowIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsActWindowModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsActWindowId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActWindowId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActWindowModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.actions.act_window was found with criteria %v and options %v", criteria, options)
}
