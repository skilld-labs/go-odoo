package odoo

import (
	"fmt"
)

// IrServerObjectLines represents ir.server.object.lines model.
type IrServerObjectLines struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	Col1        *Many2One  `xmlrpc:"col1,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	ServerId    *Many2One  `xmlrpc:"server_id,omitempty"`
	Type        *Selection `xmlrpc:"type,omitempty"`
	Value       *String    `xmlrpc:"value,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrServerObjectLiness represents array of ir.server.object.lines model.
type IrServerObjectLiness []IrServerObjectLines

// IrServerObjectLinesModel is the odoo model name.
const IrServerObjectLinesModel = "ir.server.object.lines"

// Many2One convert IrServerObjectLines to *Many2One.
func (isol *IrServerObjectLines) Many2One() *Many2One {
	return NewMany2One(isol.Id.Get(), "")
}

// CreateIrServerObjectLines creates a new ir.server.object.lines model and returns its id.
func (c *Client) CreateIrServerObjectLines(isol *IrServerObjectLines) (int64, error) {
	return c.Create(IrServerObjectLinesModel, isol)
}

// UpdateIrServerObjectLines updates an existing ir.server.object.lines record.
func (c *Client) UpdateIrServerObjectLines(isol *IrServerObjectLines) error {
	return c.UpdateIrServerObjectLiness([]int64{isol.Id.Get()}, isol)
}

// UpdateIrServerObjectLiness updates existing ir.server.object.lines records.
// All records (represented by ids) will be updated by isol values.
func (c *Client) UpdateIrServerObjectLiness(ids []int64, isol *IrServerObjectLines) error {
	return c.Update(IrServerObjectLinesModel, ids, isol)
}

// DeleteIrServerObjectLines deletes an existing ir.server.object.lines record.
func (c *Client) DeleteIrServerObjectLines(id int64) error {
	return c.DeleteIrServerObjectLiness([]int64{id})
}

// DeleteIrServerObjectLiness deletes existing ir.server.object.lines records.
func (c *Client) DeleteIrServerObjectLiness(ids []int64) error {
	return c.Delete(IrServerObjectLinesModel, ids)
}

// GetIrServerObjectLines gets ir.server.object.lines existing record.
func (c *Client) GetIrServerObjectLines(id int64) (*IrServerObjectLines, error) {
	isols, err := c.GetIrServerObjectLiness([]int64{id})
	if err != nil {
		return nil, err
	}
	if isols != nil && len(*isols) > 0 {
		return &((*isols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.server.object.lines not found", id)
}

// GetIrServerObjectLiness gets ir.server.object.lines existing records.
func (c *Client) GetIrServerObjectLiness(ids []int64) (*IrServerObjectLiness, error) {
	isols := &IrServerObjectLiness{}
	if err := c.Read(IrServerObjectLinesModel, ids, nil, isols); err != nil {
		return nil, err
	}
	return isols, nil
}

// FindIrServerObjectLines finds ir.server.object.lines record by querying it with criteria.
func (c *Client) FindIrServerObjectLines(criteria *Criteria) (*IrServerObjectLines, error) {
	isols := &IrServerObjectLiness{}
	if err := c.SearchRead(IrServerObjectLinesModel, criteria, NewOptions().Limit(1), isols); err != nil {
		return nil, err
	}
	if isols != nil && len(*isols) > 0 {
		return &((*isols)[0]), nil
	}
	return nil, fmt.Errorf("no ir.server.object.lines was found with criteria %v", criteria)
}

// FindIrServerObjectLiness finds ir.server.object.lines records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrServerObjectLiness(criteria *Criteria, options *Options) (*IrServerObjectLiness, error) {
	isols := &IrServerObjectLiness{}
	if err := c.SearchRead(IrServerObjectLinesModel, criteria, options, isols); err != nil {
		return nil, err
	}
	return isols, nil
}

// FindIrServerObjectLinesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrServerObjectLinesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrServerObjectLinesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrServerObjectLinesId finds record id by querying it with criteria.
func (c *Client) FindIrServerObjectLinesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrServerObjectLinesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.server.object.lines was found with criteria %v and options %v", criteria, options)
}
