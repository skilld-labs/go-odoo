package odoo

import (
	"fmt"
)

// IrQwebFieldImage represents ir.qweb.field.image model.
type IrQwebFieldImage struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldImages represents array of ir.qweb.field.image model.
type IrQwebFieldImages []IrQwebFieldImage

// IrQwebFieldImageModel is the odoo model name.
const IrQwebFieldImageModel = "ir.qweb.field.image"

// Many2One convert IrQwebFieldImage to *Many2One.
func (iqfi *IrQwebFieldImage) Many2One() *Many2One {
	return NewMany2One(iqfi.Id.Get(), "")
}

// CreateIrQwebFieldImage creates a new ir.qweb.field.image model and returns its id.
func (c *Client) CreateIrQwebFieldImage(iqfi *IrQwebFieldImage) (int64, error) {
	return c.Create(IrQwebFieldImageModel, iqfi)
}

// UpdateIrQwebFieldImage updates an existing ir.qweb.field.image record.
func (c *Client) UpdateIrQwebFieldImage(iqfi *IrQwebFieldImage) error {
	return c.UpdateIrQwebFieldImages([]int64{iqfi.Id.Get()}, iqfi)
}

// UpdateIrQwebFieldImages updates existing ir.qweb.field.image records.
// All records (represented by IDs) will be updated by iqfi values.
func (c *Client) UpdateIrQwebFieldImages(ids []int64, iqfi *IrQwebFieldImage) error {
	return c.Update(IrQwebFieldImageModel, ids, iqfi)
}

// DeleteIrQwebFieldImage deletes an existing ir.qweb.field.image record.
func (c *Client) DeleteIrQwebFieldImage(id int64) error {
	return c.DeleteIrQwebFieldImages([]int64{id})
}

// DeleteIrQwebFieldImages deletes existing ir.qweb.field.image records.
func (c *Client) DeleteIrQwebFieldImages(ids []int64) error {
	return c.Delete(IrQwebFieldImageModel, ids)
}

// GetIrQwebFieldImage gets ir.qweb.field.image existing record.
func (c *Client) GetIrQwebFieldImage(id int64) (*IrQwebFieldImage, error) {
	iqfis, err := c.GetIrQwebFieldImages([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.image not found", id)
}

// GetIrQwebFieldImages gets ir.qweb.field.image existing records.
func (c *Client) GetIrQwebFieldImages(ids []int64) (*IrQwebFieldImages, error) {
	iqfis := &IrQwebFieldImages{}
	if err := c.Read(IrQwebFieldImageModel, ids, nil, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImage finds ir.qweb.field.image record by querying it with criteria.
func (c *Client) FindIrQwebFieldImage(criteria *Criteria) (*IrQwebFieldImage, error) {
	iqfis := &IrQwebFieldImages{}
	if err := c.SearchRead(IrQwebFieldImageModel, criteria, NewOptions().Limit(1), iqfis); err != nil {
		return nil, err
	}
	if iqfis != nil && len(*iqfis) > 0 {
		return &((*iqfis)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.image was not found")
}

// FindIrQwebFieldImages finds ir.qweb.field.image records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImages(criteria *Criteria, options *Options) (*IrQwebFieldImages, error) {
	iqfis := &IrQwebFieldImages{}
	if err := c.SearchRead(IrQwebFieldImageModel, criteria, options, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImageIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldImageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldImageId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldImageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldImageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.image was not found")
}
