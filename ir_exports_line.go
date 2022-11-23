package odoo

import (
	"fmt"
)

// IrExportsLine represents ir.exports.line model.
type IrExportsLine struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	ExportId    *Many2One `xmlrpc:"export_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrExportsLines represents array of ir.exports.line model.
type IrExportsLines []IrExportsLine

// IrExportsLineModel is the odoo model name.
const IrExportsLineModel = "ir.exports.line"

// Many2One convert IrExportsLine to *Many2One.
func (iel *IrExportsLine) Many2One() *Many2One {
	return NewMany2One(iel.Id.Get(), "")
}

// CreateIrExportsLine creates a new ir.exports.line model and returns its id.
func (c *Client) CreateIrExportsLine(iel *IrExportsLine) (int64, error) {
	return c.Create(IrExportsLineModel, iel)
}

// UpdateIrExportsLine updates an existing ir.exports.line record.
func (c *Client) UpdateIrExportsLine(iel *IrExportsLine) error {
	return c.UpdateIrExportsLines([]int64{iel.Id.Get()}, iel)
}

// UpdateIrExportsLines updates existing ir.exports.line records.
// All records (represented by IDs) will be updated by iel values.
func (c *Client) UpdateIrExportsLines(ids []int64, iel *IrExportsLine) error {
	return c.Update(IrExportsLineModel, ids, iel)
}

// DeleteIrExportsLine deletes an existing ir.exports.line record.
func (c *Client) DeleteIrExportsLine(id int64) error {
	return c.DeleteIrExportsLines([]int64{id})
}

// DeleteIrExportsLines deletes existing ir.exports.line records.
func (c *Client) DeleteIrExportsLines(ids []int64) error {
	return c.Delete(IrExportsLineModel, ids)
}

// GetIrExportsLine gets ir.exports.line existing record.
func (c *Client) GetIrExportsLine(id int64) (*IrExportsLine, error) {
	iels, err := c.GetIrExportsLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if iels != nil && len(*iels) > 0 {
		return &((*iels)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.exports.line not found", id)
}

// GetIrExportsLines gets ir.exports.line existing records.
func (c *Client) GetIrExportsLines(ids []int64) (*IrExportsLines, error) {
	iels := &IrExportsLines{}
	if err := c.Read(IrExportsLineModel, ids, nil, iels); err != nil {
		return nil, err
	}
	return iels, nil
}

// FindIrExportsLine finds ir.exports.line record by querying it with criteria.
func (c *Client) FindIrExportsLine(criteria *Criteria) (*IrExportsLine, error) {
	iels := &IrExportsLines{}
	if err := c.SearchRead(IrExportsLineModel, criteria, NewOptions().Limit(1), iels); err != nil {
		return nil, err
	}
	if iels != nil && len(*iels) > 0 {
		return &((*iels)[0]), nil
	}
	return nil, fmt.Errorf("ir.exports.line was not found")
}

// FindIrExportsLines finds ir.exports.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrExportsLines(criteria *Criteria, options *Options) (*IrExportsLines, error) {
	iels := &IrExportsLines{}
	if err := c.SearchRead(IrExportsLineModel, criteria, options, iels); err != nil {
		return nil, err
	}
	return iels, nil
}

// FindIrExportsLineIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrExportsLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrExportsLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrExportsLineId finds record id by querying it with criteria.
func (c *Client) FindIrExportsLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrExportsLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.exports.line was not found")
}
