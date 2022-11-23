package odoo

import (
	"fmt"
)

// IrQweb represents ir.qweb model.
type IrQweb struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebs represents array of ir.qweb model.
type IrQwebs []IrQweb

// IrQwebModel is the odoo model name.
const IrQwebModel = "ir.qweb"

// Many2One convert IrQweb to *Many2One.
func (iq *IrQweb) Many2One() *Many2One {
	return NewMany2One(iq.Id.Get(), "")
}

// CreateIrQweb creates a new ir.qweb model and returns its id.
func (c *Client) CreateIrQweb(iq *IrQweb) (int64, error) {
	return c.Create(IrQwebModel, iq)
}

// UpdateIrQweb updates an existing ir.qweb record.
func (c *Client) UpdateIrQweb(iq *IrQweb) error {
	return c.UpdateIrQwebs([]int64{iq.Id.Get()}, iq)
}

// UpdateIrQwebs updates existing ir.qweb records.
// All records (represented by IDs) will be updated by iq values.
func (c *Client) UpdateIrQwebs(ids []int64, iq *IrQweb) error {
	return c.Update(IrQwebModel, ids, iq)
}

// DeleteIrQweb deletes an existing ir.qweb record.
func (c *Client) DeleteIrQweb(id int64) error {
	return c.DeleteIrQwebs([]int64{id})
}

// DeleteIrQwebs deletes existing ir.qweb records.
func (c *Client) DeleteIrQwebs(ids []int64) error {
	return c.Delete(IrQwebModel, ids)
}

// GetIrQweb gets ir.qweb existing record.
func (c *Client) GetIrQweb(id int64) (*IrQweb, error) {
	iqs, err := c.GetIrQwebs([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqs != nil && len(*iqs) > 0 {
		return &((*iqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb not found", id)
}

// GetIrQwebs gets ir.qweb existing records.
func (c *Client) GetIrQwebs(ids []int64) (*IrQwebs, error) {
	iqs := &IrQwebs{}
	if err := c.Read(IrQwebModel, ids, nil, iqs); err != nil {
		return nil, err
	}
	return iqs, nil
}

// FindIrQweb finds ir.qweb record by querying it with criteria.
func (c *Client) FindIrQweb(criteria *Criteria) (*IrQweb, error) {
	iqs := &IrQwebs{}
	if err := c.SearchRead(IrQwebModel, criteria, NewOptions().Limit(1), iqs); err != nil {
		return nil, err
	}
	if iqs != nil && len(*iqs) > 0 {
		return &((*iqs)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb was not found")
}

// FindIrQwebs finds ir.qweb records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebs(criteria *Criteria, options *Options) (*IrQwebs, error) {
	iqs := &IrQwebs{}
	if err := c.SearchRead(IrQwebModel, criteria, options, iqs); err != nil {
		return nil, err
	}
	return iqs, nil
}

// FindIrQwebIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebId finds record id by querying it with criteria.
func (c *Client) FindIrQwebId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb was not found")
}
