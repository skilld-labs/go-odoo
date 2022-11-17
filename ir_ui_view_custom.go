package odoo

import (
	"fmt"
)

// IrUiViewCustom represents ir.ui.view.custom model.
type IrUiViewCustom struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Arch        *String   `xmlrpc:"arch,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	RefId       *Many2One `xmlrpc:"ref_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrUiViewCustoms represents array of ir.ui.view.custom model.
type IrUiViewCustoms []IrUiViewCustom

// IrUiViewCustomModel is the odoo model name.
const IrUiViewCustomModel = "ir.ui.view.custom"

// Many2One convert IrUiViewCustom to *Many2One.
func (iuvc *IrUiViewCustom) Many2One() *Many2One {
	return NewMany2One(iuvc.Id.Get(), "")
}

// CreateIrUiViewCustom creates a new ir.ui.view.custom model and returns its id.
func (c *Client) CreateIrUiViewCustom(iuvc *IrUiViewCustom) (int64, error) {
	return c.Create(IrUiViewCustomModel, iuvc)
}

// UpdateIrUiViewCustom updates an existing ir.ui.view.custom record.
func (c *Client) UpdateIrUiViewCustom(iuvc *IrUiViewCustom) error {
	return c.UpdateIrUiViewCustoms([]int64{iuvc.Id.Get()}, iuvc)
}

// UpdateIrUiViewCustoms updates existing ir.ui.view.custom records.
// All records (represented by ids) will be updated by iuvc values.
func (c *Client) UpdateIrUiViewCustoms(ids []int64, iuvc *IrUiViewCustom) error {
	return c.Update(IrUiViewCustomModel, ids, iuvc)
}

// DeleteIrUiViewCustom deletes an existing ir.ui.view.custom record.
func (c *Client) DeleteIrUiViewCustom(id int64) error {
	return c.DeleteIrUiViewCustoms([]int64{id})
}

// DeleteIrUiViewCustoms deletes existing ir.ui.view.custom records.
func (c *Client) DeleteIrUiViewCustoms(ids []int64) error {
	return c.Delete(IrUiViewCustomModel, ids)
}

// GetIrUiViewCustom gets ir.ui.view.custom existing record.
func (c *Client) GetIrUiViewCustom(id int64) (*IrUiViewCustom, error) {
	iuvcs, err := c.GetIrUiViewCustoms([]int64{id})
	if err != nil {
		return nil, err
	}
	if iuvcs != nil && len(*iuvcs) > 0 {
		return &((*iuvcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.ui.view.custom not found", id)
}

// GetIrUiViewCustoms gets ir.ui.view.custom existing records.
func (c *Client) GetIrUiViewCustoms(ids []int64) (*IrUiViewCustoms, error) {
	iuvcs := &IrUiViewCustoms{}
	if err := c.Read(IrUiViewCustomModel, ids, nil, iuvcs); err != nil {
		return nil, err
	}
	return iuvcs, nil
}

// FindIrUiViewCustom finds ir.ui.view.custom record by querying it with criteria.
func (c *Client) FindIrUiViewCustom(criteria *Criteria) (*IrUiViewCustom, error) {
	iuvcs := &IrUiViewCustoms{}
	if err := c.SearchRead(IrUiViewCustomModel, criteria, NewOptions().Limit(1), iuvcs); err != nil {
		return nil, err
	}
	if iuvcs != nil && len(*iuvcs) > 0 {
		return &((*iuvcs)[0]), nil
	}
	return nil, fmt.Errorf("no ir.ui.view.custom was found with criteria %v", criteria)
}

// FindIrUiViewCustoms finds ir.ui.view.custom records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViewCustoms(criteria *Criteria, options *Options) (*IrUiViewCustoms, error) {
	iuvcs := &IrUiViewCustoms{}
	if err := c.SearchRead(IrUiViewCustomModel, criteria, options, iuvcs); err != nil {
		return nil, err
	}
	return iuvcs, nil
}

// FindIrUiViewCustomIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViewCustomIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrUiViewCustomModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrUiViewCustomId finds record id by querying it with criteria.
func (c *Client) FindIrUiViewCustomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrUiViewCustomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.ui.view.custom was found with criteria %v and options %v", criteria, options)
}
