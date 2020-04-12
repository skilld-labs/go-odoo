package odoo

import (
	"fmt"
)

// IrModel represents ir.model model
type IrModel struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AccessIds         *Relation  `xmlrpc:"access_ids,omptempty"`
	Count             *Int       `xmlrpc:"count,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	FieldId           *Relation  `xmlrpc:"field_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	Info              *String    `xmlrpc:"info,omptempty"`
	InheritedModelIds *Relation  `xmlrpc:"inherited_model_ids,omptempty"`
	Model             *String    `xmlrpc:"model,omptempty"`
	Modules           *String    `xmlrpc:"modules,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	RuleIds           *Relation  `xmlrpc:"rule_ids,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	Transient         *Bool      `xmlrpc:"transient,omptempty"`
	ViewIds           *Relation  `xmlrpc:"view_ids,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrModels represents array of ir.model model
type IrModels []IrModel

// IrModelModel is the odoo model name
const IrModelModel = "ir.model"

// CreateIrModel creates a new ir.model model and returns its id.
func (c *Client) CreateIrModel(im *IrModel) (int64, error) {
	return c.Create(IrModelModel, im)
}

// UpdateIrModel pdates an existing ir.model record.
func (c *Client) UpdateIrModel(im *IrModel) error {
	return c.UpdateIrModels([]int64{im.Id.Get()}, im)
}

// UpdateIrModels updates existing ir.model records.
// All records (represented by ids) will be updated by im values.
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
	return nil, fmt.Errorf("id %v of %s not found", id, IrModelModel)
}

// GetIrModels gets ir.model existing records.
func (c *Client) GetIrModels(ids []int64) (*IrModels, error) {
	ims := &IrModels{}
	if err := c.Read(IrModelModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
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
