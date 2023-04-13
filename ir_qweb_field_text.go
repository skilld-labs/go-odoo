package odoo

import (
	"fmt"
)

// IrQwebFieldText represents ir.qweb.field.text model.
type IrQwebFieldText struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrQwebFieldTexts represents array of ir.qweb.field.text model.
type IrQwebFieldTexts []IrQwebFieldText

// IrQwebFieldTextModel is the odoo model name.
const IrQwebFieldTextModel = "ir.qweb.field.text"

// Many2One convert IrQwebFieldText to *Many2One.
func (iqft *IrQwebFieldText) Many2One() *Many2One {
	return NewMany2One(iqft.Id.Get(), "")
}

// CreateIrQwebFieldText creates a new ir.qweb.field.text model and returns its id.
func (c *Client) CreateIrQwebFieldText(iqft *IrQwebFieldText) (int64, error) {
	ids, err := c.CreateIrQwebFieldTexts([]*IrQwebFieldText{iqft})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldText creates a new ir.qweb.field.text model and returns its id.
func (c *Client) CreateIrQwebFieldTexts(iqfts []*IrQwebFieldText) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfts {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldTextModel, vv)
}

// UpdateIrQwebFieldText updates an existing ir.qweb.field.text record.
func (c *Client) UpdateIrQwebFieldText(iqft *IrQwebFieldText) error {
	return c.UpdateIrQwebFieldTexts([]int64{iqft.Id.Get()}, iqft)
}

// UpdateIrQwebFieldTexts updates existing ir.qweb.field.text records.
// All records (represented by ids) will be updated by iqft values.
func (c *Client) UpdateIrQwebFieldTexts(ids []int64, iqft *IrQwebFieldText) error {
	return c.Update(IrQwebFieldTextModel, ids, iqft)
}

// DeleteIrQwebFieldText deletes an existing ir.qweb.field.text record.
func (c *Client) DeleteIrQwebFieldText(id int64) error {
	return c.DeleteIrQwebFieldTexts([]int64{id})
}

// DeleteIrQwebFieldTexts deletes existing ir.qweb.field.text records.
func (c *Client) DeleteIrQwebFieldTexts(ids []int64) error {
	return c.Delete(IrQwebFieldTextModel, ids)
}

// GetIrQwebFieldText gets ir.qweb.field.text existing record.
func (c *Client) GetIrQwebFieldText(id int64) (*IrQwebFieldText, error) {
	iqfts, err := c.GetIrQwebFieldTexts([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfts != nil && len(*iqfts) > 0 {
		return &((*iqfts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.text not found", id)
}

// GetIrQwebFieldTexts gets ir.qweb.field.text existing records.
func (c *Client) GetIrQwebFieldTexts(ids []int64) (*IrQwebFieldTexts, error) {
	iqfts := &IrQwebFieldTexts{}
	if err := c.Read(IrQwebFieldTextModel, ids, nil, iqfts); err != nil {
		return nil, err
	}
	return iqfts, nil
}

// FindIrQwebFieldText finds ir.qweb.field.text record by querying it with criteria.
func (c *Client) FindIrQwebFieldText(criteria *Criteria) (*IrQwebFieldText, error) {
	iqfts := &IrQwebFieldTexts{}
	if err := c.SearchRead(IrQwebFieldTextModel, criteria, NewOptions().Limit(1), iqfts); err != nil {
		return nil, err
	}
	if iqfts != nil && len(*iqfts) > 0 {
		return &((*iqfts)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.text was not found with criteria %v", criteria)
}

// FindIrQwebFieldTexts finds ir.qweb.field.text records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldTexts(criteria *Criteria, options *Options) (*IrQwebFieldTexts, error) {
	iqfts := &IrQwebFieldTexts{}
	if err := c.SearchRead(IrQwebFieldTextModel, criteria, options, iqfts); err != nil {
		return nil, err
	}
	return iqfts, nil
}

// FindIrQwebFieldTextIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldTextIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldTextModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldTextId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldTextId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldTextModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.text was not found with criteria %v and options %v", criteria, options)
}
