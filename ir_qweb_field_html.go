package odoo

import (
	"fmt"
)

// IrQwebFieldHtml represents ir.qweb.field.html model.
type IrQwebFieldHtml struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrQwebFieldHtmls represents array of ir.qweb.field.html model.
type IrQwebFieldHtmls []IrQwebFieldHtml

// IrQwebFieldHtmlModel is the odoo model name.
const IrQwebFieldHtmlModel = "ir.qweb.field.html"

// Many2One convert IrQwebFieldHtml to *Many2One.
func (iqfh *IrQwebFieldHtml) Many2One() *Many2One {
	return NewMany2One(iqfh.Id.Get(), "")
}

// CreateIrQwebFieldHtml creates a new ir.qweb.field.html model and returns its id.
func (c *Client) CreateIrQwebFieldHtml(iqfh *IrQwebFieldHtml) (int64, error) {
	return c.Create(IrQwebFieldHtmlModel, iqfh)
}

// UpdateIrQwebFieldHtml updates an existing ir.qweb.field.html record.
func (c *Client) UpdateIrQwebFieldHtml(iqfh *IrQwebFieldHtml) error {
	return c.UpdateIrQwebFieldHtmls([]int64{iqfh.Id.Get()}, iqfh)
}

// UpdateIrQwebFieldHtmls updates existing ir.qweb.field.html records.
// All records (represented by ids) will be updated by iqfh values.
func (c *Client) UpdateIrQwebFieldHtmls(ids []int64, iqfh *IrQwebFieldHtml) error {
	return c.Update(IrQwebFieldHtmlModel, ids, iqfh)
}

// DeleteIrQwebFieldHtml deletes an existing ir.qweb.field.html record.
func (c *Client) DeleteIrQwebFieldHtml(id int64) error {
	return c.DeleteIrQwebFieldHtmls([]int64{id})
}

// DeleteIrQwebFieldHtmls deletes existing ir.qweb.field.html records.
func (c *Client) DeleteIrQwebFieldHtmls(ids []int64) error {
	return c.Delete(IrQwebFieldHtmlModel, ids)
}

// GetIrQwebFieldHtml gets ir.qweb.field.html existing record.
func (c *Client) GetIrQwebFieldHtml(id int64) (*IrQwebFieldHtml, error) {
	iqfhs, err := c.GetIrQwebFieldHtmls([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfhs != nil && len(*iqfhs) > 0 {
		return &((*iqfhs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.html not found", id)
}

// GetIrQwebFieldHtmls gets ir.qweb.field.html existing records.
func (c *Client) GetIrQwebFieldHtmls(ids []int64) (*IrQwebFieldHtmls, error) {
	iqfhs := &IrQwebFieldHtmls{}
	if err := c.Read(IrQwebFieldHtmlModel, ids, nil, iqfhs); err != nil {
		return nil, err
	}
	return iqfhs, nil
}

// FindIrQwebFieldHtml finds ir.qweb.field.html record by querying it with criteria.
func (c *Client) FindIrQwebFieldHtml(criteria *Criteria) (*IrQwebFieldHtml, error) {
	iqfhs := &IrQwebFieldHtmls{}
	if err := c.SearchRead(IrQwebFieldHtmlModel, criteria, NewOptions().Limit(1), iqfhs); err != nil {
		return nil, err
	}
	if iqfhs != nil && len(*iqfhs) > 0 {
		return &((*iqfhs)[0]), nil
	}
	return nil, fmt.Errorf("no ir.qweb.field.html was found with criteria %v", criteria)
}

// FindIrQwebFieldHtmls finds ir.qweb.field.html records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldHtmls(criteria *Criteria, options *Options) (*IrQwebFieldHtmls, error) {
	iqfhs := &IrQwebFieldHtmls{}
	if err := c.SearchRead(IrQwebFieldHtmlModel, criteria, options, iqfhs); err != nil {
		return nil, err
	}
	return iqfhs, nil
}

// FindIrQwebFieldHtmlIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldHtmlIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldHtmlModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldHtmlId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldHtmlId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldHtmlModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.qweb.field.html was found with criteria %v and options %v", criteria, options)
}
