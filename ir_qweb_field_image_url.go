package odoo

import (
	"fmt"
)

// IrQwebFieldImageUrl represents ir.qweb.field.image_url model.
type IrQwebFieldImageUrl struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldImageUrls represents array of ir.qweb.field.image_url model.
type IrQwebFieldImageUrls []IrQwebFieldImageUrl

// IrQwebFieldImageUrlModel is the odoo model name.
const IrQwebFieldImageUrlModel = "ir.qweb.field.image_url"

// Many2One convert IrQwebFieldImageUrl to *Many2One.
func (iqfi *IrQwebFieldImageUrl) Many2One() *Many2One {
	return NewMany2One(iqfi.Id.Get(), "")
}

// CreateIrQwebFieldImageUrl creates a new ir.qweb.field.image_url model and returns its id.
func (c *Client) CreateIrQwebFieldImageUrl(iqfi *IrQwebFieldImageUrl) (int64, error) {
	return c.Create(IrQwebFieldImageUrlModel, iqfi)
}

// UpdateIrQwebFieldImageUrl updates an existing ir.qweb.field.image_url record.
func (c *Client) UpdateIrQwebFieldImageUrl(iqfi *IrQwebFieldImageUrl) error {
	return c.UpdateIrQwebFieldImageUrls([]int64{iqfi.Id.Get()}, iqfi)
}

// UpdateIrQwebFieldImageUrls updates existing ir.qweb.field.image_url records.
// All records (represented by ids) will be updated by iqfi values.
func (c *Client) UpdateIrQwebFieldImageUrls(ids []int64, iqfi *IrQwebFieldImageUrl) error {
	return c.Update(IrQwebFieldImageUrlModel, ids, iqfi)
}

// DeleteIrQwebFieldImageUrl deletes an existing ir.qweb.field.image_url record.
func (c *Client) DeleteIrQwebFieldImageUrl(id int64) error {
	return c.DeleteIrQwebFieldImageUrls([]int64{id})
}

// DeleteIrQwebFieldImageUrls deletes existing ir.qweb.field.image_url records.
func (c *Client) DeleteIrQwebFieldImageUrls(ids []int64) error {
	return c.Delete(IrQwebFieldImageUrlModel, ids)
}

// GetIrQwebFieldImageUrl gets ir.qweb.field.image_url existing record.
func (c *Client) GetIrQwebFieldImageUrl(id int64) (*IrQwebFieldImageUrl, error) {
	iqfis, err := c.GetIrQwebFieldImageUrls([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.image_url not found", id)
}

// GetIrQwebFieldImageUrls gets ir.qweb.field.image_url existing records.
func (c *Client) GetIrQwebFieldImageUrls(ids []int64) (*IrQwebFieldImageUrls, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.Read(IrQwebFieldImageUrlModel, ids, nil, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImageUrl finds ir.qweb.field.image_url record by querying it with criteria.
func (c *Client) FindIrQwebFieldImageUrl(criteria *Criteria) (*IrQwebFieldImageUrl, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.SearchRead(IrQwebFieldImageUrlModel, criteria, NewOptions().Limit(1), iqfis); err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.image_url was not found")
}

// FindIrQwebFieldImageUrls finds ir.qweb.field.image_url records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImageUrls(criteria *Criteria, options *Options) (*IrQwebFieldImageUrls, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.SearchRead(IrQwebFieldImageUrlModel, criteria, options, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImageUrlIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImageUrlIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldImageUrlModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldImageUrlId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldImageUrlId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldImageUrlModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.image_url was not found")
}
