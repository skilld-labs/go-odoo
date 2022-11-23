package odoo

import (
	"fmt"
)

// IrModel represents ir.model model.
type IrModel struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	AccessIds         *Relation  `xmlrpc:"access_ids,omitempty"`
	Count             *Int       `xmlrpc:"count,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	FieldId           *Relation  `xmlrpc:"field_id,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Info              *String    `xmlrpc:"info,omitempty"`
	InheritedModelIds *Relation  `xmlrpc:"inherited_model_ids,omitempty"`
	IsMailActivity    *Bool      `xmlrpc:"is_mail_activity,omitempty"`
	IsMailBlacklist   *Bool      `xmlrpc:"is_mail_blacklist,omitempty"`
	IsMailThread      *Bool      `xmlrpc:"is_mail_thread,omitempty"`
	IsMailThreadSms   *Bool      `xmlrpc:"is_mail_thread_sms,omitempty"`
	Model             *String    `xmlrpc:"model,omitempty"`
	Modules           *String    `xmlrpc:"modules,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	Order             *String    `xmlrpc:"order,omitempty"`
	RuleIds           *Relation  `xmlrpc:"rule_ids,omitempty"`
	State             *Selection `xmlrpc:"state,omitempty"`
	Transient         *Bool      `xmlrpc:"transient,omitempty"`
	ViewIds           *Relation  `xmlrpc:"view_ids,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrModels represents array of ir.model model.
type IrModels []IrModel

// IrModelModel is the odoo model name.
const IrModelModel = "ir.model"

// Many2One convert IrModel to *Many2One.
func (im *IrModel) Many2One() *Many2One {
	return NewMany2One(im.Id.Get(), "")
}

// CreateIrModel creates a new ir.model model and returns its id.
func (c *Client) CreateIrModel(im *IrModel) (int64, error) {
	return c.Create(IrModelModel, im)
}

// UpdateIrModel updates an existing ir.model record.
func (c *Client) UpdateIrModel(im *IrModel) error {
	return c.UpdateIrModels([]int64{im.Id.Get()}, im)
}

// UpdateIrModels updates existing ir.model records.
// All records (represented by IDs) will be updated by im values.
func (c *Client) UpdateIrModels(ids []int64, im *IrModel) error {
	return c.Update(IrModelModel, ids, im)
}

// DeleteIrModel deletes an existing ir.model record.
func (c *Client) DeleteIrModel(id int64) error {
	return c.DeleteIrModels([]int64{id})
}

// DeleteIrModels deletes existing ir.model records.
func (c *Client) DeleteIrModels(ids []int64) error {
	return c.Delete(IrModelModel, ids)
}

// GetIrModel gets ir.model existing record.
func (c *Client) GetIrModel(id int64) (*IrModel, error) {
	ims, err := c.GetIrModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.model not found", id)
}

// GetIrModels gets ir.model existing records.
func (c *Client) GetIrModels(ids []int64) (*IrModels, error) {
	ims := &IrModels{}
	if err := c.Read(IrModelModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrModel finds ir.model record by querying it with criteria.
func (c *Client) FindIrModel(criteria *Criteria) (*IrModel, error) {
	ims := &IrModels{}
	if err := c.SearchRead(IrModelModel, criteria, NewOptions().Limit(1), ims); err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("ir.model was not found")
}

// FindIrModels finds ir.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModels(criteria *Criteria, options *Options) (*IrModels, error) {
	ims := &IrModels{}
	if err := c.SearchRead(IrModelModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrModelIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModelId finds record id by querying it with criteria.
func (c *Client) FindIrModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.model was not found")
}
