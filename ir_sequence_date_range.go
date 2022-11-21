package odoo

import (
	"fmt"
)

// IrSequenceDateRange represents ir.sequence.date_range model.
type IrSequenceDateRange struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom         *Time     `xmlrpc:"date_from,omptempty"`
	DateTo           *Time     `xmlrpc:"date_to,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	NumberNext       *Int      `xmlrpc:"number_next,omptempty"`
	NumberNextActual *Int      `xmlrpc:"number_next_actual,omptempty"`
	SequenceId       *Many2One `xmlrpc:"sequence_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrSequenceDateRanges represents array of ir.sequence.date_range model.
type IrSequenceDateRanges []IrSequenceDateRange

// IrSequenceDateRangeModel is the odoo model name.
const IrSequenceDateRangeModel = "ir.sequence.date_range"

// Many2One convert IrSequenceDateRange to *Many2One.
func (isd *IrSequenceDateRange) Many2One() *Many2One {
	return NewMany2One(isd.Id.Get(), "")
}

// CreateIrSequenceDateRange creates a new ir.sequence.date_range model and returns its id.
func (c *Client) CreateIrSequenceDateRange(isd *IrSequenceDateRange) (int64, error) {
	return c.Create(IrSequenceDateRangeModel, isd)
}

// UpdateIrSequenceDateRange updates an existing ir.sequence.date_range record.
func (c *Client) UpdateIrSequenceDateRange(isd *IrSequenceDateRange) error {
	return c.UpdateIrSequenceDateRanges([]int64{isd.Id.Get()}, isd)
}

// UpdateIrSequenceDateRanges updates existing ir.sequence.date_range records.
// All records (represented by ids) will be updated by isd values.
func (c *Client) UpdateIrSequenceDateRanges(ids []int64, isd *IrSequenceDateRange) error {
	return c.Update(IrSequenceDateRangeModel, ids, isd)
}

// DeleteIrSequenceDateRange deletes an existing ir.sequence.date_range record.
func (c *Client) DeleteIrSequenceDateRange(id int64) error {
	return c.DeleteIrSequenceDateRanges([]int64{id})
}

// DeleteIrSequenceDateRanges deletes existing ir.sequence.date_range records.
func (c *Client) DeleteIrSequenceDateRanges(ids []int64) error {
	return c.Delete(IrSequenceDateRangeModel, ids)
}

// GetIrSequenceDateRange gets ir.sequence.date_range existing record.
func (c *Client) GetIrSequenceDateRange(id int64) (*IrSequenceDateRange, error) {
	isds, err := c.GetIrSequenceDateRanges([]int64{id})
	if err != nil {
		return nil, err
	}
	if isds != nil && len(*isds) > 0 {
		return &((*isds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.sequence.date_range not found", id)
}

// GetIrSequenceDateRanges gets ir.sequence.date_range existing records.
func (c *Client) GetIrSequenceDateRanges(ids []int64) (*IrSequenceDateRanges, error) {
	isds := &IrSequenceDateRanges{}
	if err := c.Read(IrSequenceDateRangeModel, ids, nil, isds); err != nil {
		return nil, err
	}
	return isds, nil
}

// FindIrSequenceDateRange finds ir.sequence.date_range record by querying it with criteria.
func (c *Client) FindIrSequenceDateRange(criteria *Criteria) (*IrSequenceDateRange, error) {
	isds := &IrSequenceDateRanges{}
	if err := c.SearchRead(IrSequenceDateRangeModel, criteria, NewOptions().Limit(1), isds); err != nil {
		return nil, err
	}
	if isds != nil && len(*isds) > 0 {
		return &((*isds)[0]), nil
	}
	return nil, fmt.Errorf("no ir.sequence.date_range was found with criteria %v", criteria)
}

// FindIrSequenceDateRanges finds ir.sequence.date_range records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrSequenceDateRanges(criteria *Criteria, options *Options) (*IrSequenceDateRanges, error) {
	isds := &IrSequenceDateRanges{}
	if err := c.SearchRead(IrSequenceDateRangeModel, criteria, options, isds); err != nil {
		return nil, err
	}
	return isds, nil
}

// FindIrSequenceDateRangeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrSequenceDateRangeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrSequenceDateRangeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrSequenceDateRangeId finds record id by querying it with criteria.
func (c *Client) FindIrSequenceDateRangeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrSequenceDateRangeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.sequence.date_range was found with criteria %v and options %v", criteria, options)
}
