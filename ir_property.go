package odoo

import (
	"fmt"
)

// IrProperty represents ir.property model.
type IrProperty struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	FieldsId       *Many2One  `xmlrpc:"fields_id,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	ResId          *String    `xmlrpc:"res_id,omitempty"`
	Type           *Selection `xmlrpc:"type,omitempty"`
	ValueBinary    *String    `xmlrpc:"value_binary,omitempty"`
	ValueDatetime  *Time      `xmlrpc:"value_datetime,omitempty"`
	ValueFloat     *Float     `xmlrpc:"value_float,omitempty"`
	ValueInteger   *Int       `xmlrpc:"value_integer,omitempty"`
	ValueReference *String    `xmlrpc:"value_reference,omitempty"`
	ValueText      *String    `xmlrpc:"value_text,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrPropertys represents array of ir.property model.
type IrPropertys []IrProperty

// IrPropertyModel is the odoo model name.
const IrPropertyModel = "ir.property"

// Many2One convert IrProperty to *Many2One.
func (ip *IrProperty) Many2One() *Many2One {
	return NewMany2One(ip.Id.Get(), "")
}

// CreateIrProperty creates a new ir.property model and returns its id.
func (c *Client) CreateIrProperty(ip *IrProperty) (int64, error) {
	return c.Create(IrPropertyModel, ip)
}

// UpdateIrProperty updates an existing ir.property record.
func (c *Client) UpdateIrProperty(ip *IrProperty) error {
	return c.UpdateIrPropertys([]int64{ip.Id.Get()}, ip)
}

// UpdateIrPropertys updates existing ir.property records.
// All records (represented by ids) will be updated by ip values.
func (c *Client) UpdateIrPropertys(ids []int64, ip *IrProperty) error {
	return c.Update(IrPropertyModel, ids, ip)
}

// DeleteIrProperty deletes an existing ir.property record.
func (c *Client) DeleteIrProperty(id int64) error {
	return c.DeleteIrPropertys([]int64{id})
}

// DeleteIrPropertys deletes existing ir.property records.
func (c *Client) DeleteIrPropertys(ids []int64) error {
	return c.Delete(IrPropertyModel, ids)
}

// GetIrProperty gets ir.property existing record.
func (c *Client) GetIrProperty(id int64) (*IrProperty, error) {
	ips, err := c.GetIrPropertys([]int64{id})
	if err != nil {
		return nil, err
	}
	if ips != nil && len(*ips) > 0 {
		return &((*ips)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.property not found", id)
}

// GetIrPropertys gets ir.property existing records.
func (c *Client) GetIrPropertys(ids []int64) (*IrPropertys, error) {
	ips := &IrPropertys{}
	if err := c.Read(IrPropertyModel, ids, nil, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrProperty finds ir.property record by querying it with criteria.
func (c *Client) FindIrProperty(criteria *Criteria) (*IrProperty, error) {
	ips := &IrPropertys{}
	if err := c.SearchRead(IrPropertyModel, criteria, NewOptions().Limit(1), ips); err != nil {
		return nil, err
	}
	if ips != nil && len(*ips) > 0 {
		return &((*ips)[0]), nil
	}
	return nil, fmt.Errorf("ir.property was not found")
}

// FindIrPropertys finds ir.property records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrPropertys(criteria *Criteria, options *Options) (*IrPropertys, error) {
	ips := &IrPropertys{}
	if err := c.SearchRead(IrPropertyModel, criteria, options, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrPropertyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrPropertyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrPropertyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrPropertyId finds record id by querying it with criteria.
func (c *Client) FindIrPropertyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrPropertyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.property was not found")
}
