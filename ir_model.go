package odoo

import (
	"errors"
	"fmt"
)

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
	IsMailThread      *Bool      `xmlrpc:"is_mail_thread,omptempty"`
	Model             *String    `xmlrpc:"model,omptempty"`
	Modules           *String    `xmlrpc:"modules,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	Transient         *Bool      `xmlrpc:"transient,omptempty"`
	ViewIds           *Relation  `xmlrpc:"view_ids,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

type IrModels []IrModel

const IrModelModel = "ir.model"

func (c *Client) CreateIrModel(im *IrModel) (int64, error) {
	return c.Create(IrModelModel, im)
}

func (c *Client) UpdateIrModel(im *IrModel) error {
	return c.UpdateIrModels([]int64{im.Id.Get()}, im)
}

func (c *Client) UpdateIrModels(ids []int64, im *IrModel) error {
	return c.Update(IrModelModel, ids, im)
}

func (c *Client) DeleteIrModel(id int64) error {
	return c.DeleteIrModels([]int64{id})
}

func (c *Client) DeleteIrModels(ids []int64) error {
	return c.Delete(IrModelModel, ids)
}

func (c *Client) GetIrModel(id int64) (*IrModel, error) {
	ims, err := c.GetIrModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, errors.New(fmt.Sprintf("id %v of %s not found", id, IrModelModel))
}

func (c *Client) GetIrModels(ids []int64) (*IrModels, error) {
	ims := &IrModels{}
	if err := c.Read(IrModelModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

func (c *Client) FindIrModels(criteria *Criteria, options *Options) (*IrModels, error) {
	ims := &IrModels{}
	if err := c.SearchRead(IrModelModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}
