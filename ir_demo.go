package odoo

import (
	"fmt"
)

// IrDemo represents ir.demo model.
type IrDemo struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrDemos represents array of ir.demo model.
type IrDemos []IrDemo

// IrDemoModel is the odoo model name.
const IrDemoModel = "ir.demo"

// Many2One convert IrDemo to *Many2One.
func (ID *IrDemo) Many2One() *Many2One {
	return NewMany2One(ID.Id.Get(), "")
}

// CreateIrDemo creates a new ir.demo model and returns its id.
func (c *Client) CreateIrDemo(ID *IrDemo) (int64, error) {
	return c.Create(IrDemoModel, ID)
}

// UpdateIrDemo updates an existing ir.demo record.
func (c *Client) UpdateIrDemo(ID *IrDemo) error {
	return c.UpdateIrDemos([]int64{ID.Id.Get()}, ID)
}

// UpdateIrDemos updates existing ir.demo records.
// All records (represented by ids) will be updated by ID values.
func (c *Client) UpdateIrDemos(ids []int64, ID *IrDemo) error {
	return c.Update(IrDemoModel, ids, ID)
}

// DeleteIrDemo deletes an existing ir.demo record.
func (c *Client) DeleteIrDemo(id int64) error {
	return c.DeleteIrDemos([]int64{id})
}

// DeleteIrDemos deletes existing ir.demo records.
func (c *Client) DeleteIrDemos(ids []int64) error {
	return c.Delete(IrDemoModel, ids)
}

// GetIrDemo gets ir.demo existing record.
func (c *Client) GetIrDemo(id int64) (*IrDemo, error) {
	IDs, err := c.GetIrDemos([]int64{id})
	if err != nil {
		return nil, err
	}
	if IDs != nil && len(*IDs) > 0 {
		return &((*IDs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.demo not found", id)
}

// GetIrDemos gets ir.demo existing records.
func (c *Client) GetIrDemos(ids []int64) (*IrDemos, error) {
	IDs := &IrDemos{}
	if err := c.Read(IrDemoModel, ids, nil, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDemo finds ir.demo record by querying it with criteria.
func (c *Client) FindIrDemo(criteria *Criteria) (*IrDemo, error) {
	IDs := &IrDemos{}
	if err := c.SearchRead(IrDemoModel, criteria, NewOptions().Limit(1), IDs); err != nil {
		return nil, err
	}
	if IDs != nil && len(*IDs) > 0 {
		return &((*IDs)[0]), nil
	}
	return nil, fmt.Errorf("ir.demo was not found")
}

// FindIrDemos finds ir.demo records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemos(criteria *Criteria, options *Options) (*IrDemos, error) {
	IDs := &IrDemos{}
	if err := c.SearchRead(IrDemoModel, criteria, options, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDemoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemoIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrDemoModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrDemoId finds record id by querying it with criteria.
func (c *Client) FindIrDemoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrDemoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.demo was not found")
}
