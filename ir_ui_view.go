package odoo

import (
	"fmt"
)

// IrUiView represents ir.ui.view model.
type IrUiView struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	Active             *Bool      `xmlrpc:"active,omptempty"`
	Arch               *String    `xmlrpc:"arch,omptempty"`
	ArchBase           *String    `xmlrpc:"arch_base,omptempty"`
	ArchDb             *String    `xmlrpc:"arch_db,omptempty"`
	ArchFs             *String    `xmlrpc:"arch_fs,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	FieldParent        *String    `xmlrpc:"field_parent,omptempty"`
	GroupsId           *Relation  `xmlrpc:"groups_id,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	InheritChildrenIds *Relation  `xmlrpc:"inherit_children_ids,omptempty"`
	InheritId          *Many2One  `xmlrpc:"inherit_id,omptempty"`
	Key                *String    `xmlrpc:"key,omptempty"`
	Mode               *Selection `xmlrpc:"mode,omptempty"`
	Model              *String    `xmlrpc:"model,omptempty"`
	ModelDataId        *Many2One  `xmlrpc:"model_data_id,omptempty"`
	ModelIds           *Relation  `xmlrpc:"model_ids,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	Priority           *Int       `xmlrpc:"priority,omptempty"`
	Type               *Selection `xmlrpc:"type,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId              *String    `xmlrpc:"xml_id,omptempty"`
}

// IrUiViews represents array of ir.ui.view model.
type IrUiViews []IrUiView

// IrUiViewModel is the odoo model name.
const IrUiViewModel = "ir.ui.view"

// Many2One convert IrUiView to *Many2One.
func (iuv *IrUiView) Many2One() *Many2One {
	return NewMany2One(iuv.Id.Get(), "")
}

// CreateIrUiView creates a new ir.ui.view model and returns its id.
func (c *Client) CreateIrUiView(iuv *IrUiView) (int64, error) {
	return c.Create(IrUiViewModel, iuv)
}

// UpdateIrUiView updates an existing ir.ui.view record.
func (c *Client) UpdateIrUiView(iuv *IrUiView) error {
	return c.UpdateIrUiViews([]int64{iuv.Id.Get()}, iuv)
}

// UpdateIrUiViews updates existing ir.ui.view records.
// All records (represented by ids) will be updated by iuv values.
func (c *Client) UpdateIrUiViews(ids []int64, iuv *IrUiView) error {
	return c.Update(IrUiViewModel, ids, iuv)
}

// DeleteIrUiView deletes an existing ir.ui.view record.
func (c *Client) DeleteIrUiView(id int64) error {
	return c.DeleteIrUiViews([]int64{id})
}

// DeleteIrUiViews deletes existing ir.ui.view records.
func (c *Client) DeleteIrUiViews(ids []int64) error {
	return c.Delete(IrUiViewModel, ids)
}

// GetIrUiView gets ir.ui.view existing record.
func (c *Client) GetIrUiView(id int64) (*IrUiView, error) {
	iuvs, err := c.GetIrUiViews([]int64{id})
	if err != nil {
		return nil, err
	}
	if iuvs != nil && len(*iuvs) > 0 {
		return &((*iuvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.ui.view not found", id)
}

// GetIrUiViews gets ir.ui.view existing records.
func (c *Client) GetIrUiViews(ids []int64) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.Read(IrUiViewModel, ids, nil, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiView finds ir.ui.view record by querying it with criteria.
func (c *Client) FindIrUiView(criteria *Criteria) (*IrUiView, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, NewOptions().Limit(1), iuvs); err != nil {
		return nil, err
	}
	if iuvs != nil && len(*iuvs) > 0 {
		return &((*iuvs)[0]), nil
	}
	return nil, fmt.Errorf("no ir.ui.view was found with criteria %v", criteria)
}

// FindIrUiViews finds ir.ui.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViews(criteria *Criteria, options *Options) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, options, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrUiViewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrUiViewId finds record id by querying it with criteria.
func (c *Client) FindIrUiViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrUiViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.ui.view was found with criteria %v and options %v", criteria, options)
}
