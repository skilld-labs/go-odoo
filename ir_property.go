package odoo

import (
	"fmt"
)

// IrProperty represents ir.property model.
type IrProperty struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	FieldsId       *Many2One  `xmlrpc:"fields_id,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	ResId          *String    `xmlrpc:"res_id,omptempty"`
	Type           *Selection `xmlrpc:"type,omptempty"`
	ValueBinary    *String    `xmlrpc:"value_binary,omptempty"`
	ValueDatetime  *Time      `xmlrpc:"value_datetime,omptempty"`
	ValueFloat     *Float     `xmlrpc:"value_float,omptempty"`
	ValueInteger   *Int       `xmlrpc:"value_integer,omptempty"`
	ValueReference *String    `xmlrpc:"value_reference,omptempty"`
	ValueText      *String    `xmlrpc:"value_text,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return nil, fmt.Errorf("no ir.property was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no ir.property was found with criteria %v and options %v", criteria, options)
}
