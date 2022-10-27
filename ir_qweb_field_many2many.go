package odoo

import (
	"fmt"
)

// IrQwebFieldMany2Many represents ir.qweb.field.many2many model.
type IrQwebFieldMany2Many struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldMany2Manys represents array of ir.qweb.field.many2many model.
type IrQwebFieldMany2Manys []IrQwebFieldMany2Many

// IrQwebFieldMany2ManyModel is the odoo model name.
const IrQwebFieldMany2ManyModel = "ir.qweb.field.many2many"

// Many2One convert IrQwebFieldMany2Many to *Many2One.
func (iqfm *IrQwebFieldMany2Many) Many2One() *Many2One {
	return NewMany2One(iqfm.Id.Get(), "")
}

// CreateIrQwebFieldMany2Many creates a new ir.qweb.field.many2many model and returns its id.
func (c *Client) CreateIrQwebFieldMany2Many(iqfm *IrQwebFieldMany2Many) (int64, error) {
	return c.Create(IrQwebFieldMany2ManyModel, iqfm)
}

// UpdateIrQwebFieldMany2Many updates an existing ir.qweb.field.many2many record.
func (c *Client) UpdateIrQwebFieldMany2Many(iqfm *IrQwebFieldMany2Many) error {
	return c.UpdateIrQwebFieldMany2Manys([]int64{iqfm.Id.Get()}, iqfm)
}

// UpdateIrQwebFieldMany2Manys updates existing ir.qweb.field.many2many records.
// All records (represented by ids) will be updated by iqfm values.
func (c *Client) UpdateIrQwebFieldMany2Manys(ids []int64, iqfm *IrQwebFieldMany2Many) error {
	return c.Update(IrQwebFieldMany2ManyModel, ids, iqfm)
}

// DeleteIrQwebFieldMany2Many deletes an existing ir.qweb.field.many2many record.
func (c *Client) DeleteIrQwebFieldMany2Many(id int64) error {
	return c.DeleteIrQwebFieldMany2Manys([]int64{id})
}

// DeleteIrQwebFieldMany2Manys deletes existing ir.qweb.field.many2many records.
func (c *Client) DeleteIrQwebFieldMany2Manys(ids []int64) error {
	return c.Delete(IrQwebFieldMany2ManyModel, ids)
}

// GetIrQwebFieldMany2Many gets ir.qweb.field.many2many existing record.
func (c *Client) GetIrQwebFieldMany2Many(id int64) (*IrQwebFieldMany2Many, error) {
	iqfms, err := c.GetIrQwebFieldMany2Manys([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.many2many not found", id)
}

// GetIrQwebFieldMany2Manys gets ir.qweb.field.many2many existing records.
func (c *Client) GetIrQwebFieldMany2Manys(ids []int64) (*IrQwebFieldMany2Manys, error) {
	iqfms := &IrQwebFieldMany2Manys{}
	if err := c.Read(IrQwebFieldMany2ManyModel, ids, nil, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMany2Many finds ir.qweb.field.many2many record by querying it with criteria.
func (c *Client) FindIrQwebFieldMany2Many(criteria *Criteria) (*IrQwebFieldMany2Many, error) {
	iqfms := &IrQwebFieldMany2Manys{}
	if err := c.SearchRead(IrQwebFieldMany2ManyModel, criteria, NewOptions().Limit(1), iqfms); err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.many2many was not found")
}

// FindIrQwebFieldMany2Manys finds ir.qweb.field.many2many records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMany2Manys(criteria *Criteria, options *Options) (*IrQwebFieldMany2Manys, error) {
	iqfms := &IrQwebFieldMany2Manys{}
	if err := c.SearchRead(IrQwebFieldMany2ManyModel, criteria, options, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMany2ManyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMany2ManyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldMany2ManyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldMany2ManyId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldMany2ManyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldMany2ManyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.many2many was not found")
}
