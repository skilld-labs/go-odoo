package odoo

import (
	"fmt"
)

// IrSequence represents ir.sequence model.
type IrSequence struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	Active           *Bool      `xmlrpc:"active,omptempty"`
	Code             *String    `xmlrpc:"code,omptempty"`
	CompanyId        *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateRangeIds     *Relation  `xmlrpc:"date_range_ids,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Implementation   *Selection `xmlrpc:"implementation,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	NumberIncrement  *Int       `xmlrpc:"number_increment,omptempty"`
	NumberNext       *Int       `xmlrpc:"number_next,omptempty"`
	NumberNextActual *Int       `xmlrpc:"number_next_actual,omptempty"`
	Padding          *Int       `xmlrpc:"padding,omptempty"`
	Prefix           *String    `xmlrpc:"prefix,omptempty"`
	Suffix           *String    `xmlrpc:"suffix,omptempty"`
	UseDateRange     *Bool      `xmlrpc:"use_date_range,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrSequences represents array of ir.sequence model.
type IrSequences []IrSequence

// IrSequenceModel is the odoo model name.
const IrSequenceModel = "ir.sequence"

// Many2One convert IrSequence to *Many2One.
func (is *IrSequence) Many2One() *Many2One {
	return NewMany2One(is.Id.Get(), "")
}

// CreateIrSequence creates a new ir.sequence model and returns its id.
func (c *Client) CreateIrSequence(is *IrSequence) (int64, error) {
	return c.Create(IrSequenceModel, is)
}

// UpdateIrSequence updates an existing ir.sequence record.
func (c *Client) UpdateIrSequence(is *IrSequence) error {
	return c.UpdateIrSequences([]int64{is.Id.Get()}, is)
}

// UpdateIrSequences updates existing ir.sequence records.
// All records (represented by ids) will be updated by is values.
func (c *Client) UpdateIrSequences(ids []int64, is *IrSequence) error {
	return c.Update(IrSequenceModel, ids, is)
}

// DeleteIrSequence deletes an existing ir.sequence record.
func (c *Client) DeleteIrSequence(id int64) error {
	return c.DeleteIrSequences([]int64{id})
}

// DeleteIrSequences deletes existing ir.sequence records.
func (c *Client) DeleteIrSequences(ids []int64) error {
	return c.Delete(IrSequenceModel, ids)
}

// GetIrSequence gets ir.sequence existing record.
func (c *Client) GetIrSequence(id int64) (*IrSequence, error) {
	iss, err := c.GetIrSequences([]int64{id})
	if err != nil {
		return nil, err
	}
	if iss != nil && len(*iss) > 0 {
		return &((*iss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.sequence not found", id)
}

// GetIrSequences gets ir.sequence existing records.
func (c *Client) GetIrSequences(ids []int64) (*IrSequences, error) {
	iss := &IrSequences{}
	if err := c.Read(IrSequenceModel, ids, nil, iss); err != nil {
		return nil, err
	}
	return iss, nil
}

// FindIrSequence finds ir.sequence record by querying it with criteria.
func (c *Client) FindIrSequence(criteria *Criteria) (*IrSequence, error) {
	iss := &IrSequences{}
	if err := c.SearchRead(IrSequenceModel, criteria, NewOptions().Limit(1), iss); err != nil {
		return nil, err
	}
	if iss != nil && len(*iss) > 0 {
		return &((*iss)[0]), nil
	}
	return nil, fmt.Errorf("ir.sequence was not found with criteria %v", criteria)
}

// FindIrSequences finds ir.sequence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrSequences(criteria *Criteria, options *Options) (*IrSequences, error) {
	iss := &IrSequences{}
	if err := c.SearchRead(IrSequenceModel, criteria, options, iss); err != nil {
		return nil, err
	}
	return iss, nil
}

// FindIrSequenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrSequenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrSequenceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrSequenceId finds record id by querying it with criteria.
func (c *Client) FindIrSequenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrSequenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.sequence was not found with criteria %v and options %v", criteria, options)
}
