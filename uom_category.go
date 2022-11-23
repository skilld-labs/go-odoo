package odoo

import (
	"fmt"
)

// UomCategory represents uom.category model.
type UomCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// UomCategorys represents array of uom.category model.
type UomCategorys []UomCategory

// UomCategoryModel is the odoo model name.
const UomCategoryModel = "uom.category"

// Many2One convert UomCategory to *Many2One.
func (uc *UomCategory) Many2One() *Many2One {
	return NewMany2One(uc.Id.Get(), "")
}

// CreateUomCategory creates a new uom.category model and returns its id.
func (c *Client) CreateUomCategory(uc *UomCategory) (int64, error) {
	return c.Create(UomCategoryModel, uc)
}

// UpdateUomCategory updates an existing uom.category record.
func (c *Client) UpdateUomCategory(uc *UomCategory) error {
	return c.UpdateUomCategorys([]int64{uc.Id.Get()}, uc)
}

// UpdateUomCategorys updates existing uom.category records.
// All records (represented by IDs) will be updated by uc values.
func (c *Client) UpdateUomCategorys(ids []int64, uc *UomCategory) error {
	return c.Update(UomCategoryModel, ids, uc)
}

// DeleteUomCategory deletes an existing uom.category record.
func (c *Client) DeleteUomCategory(id int64) error {
	return c.DeleteUomCategorys([]int64{id})
}

// DeleteUomCategorys deletes existing uom.category records.
func (c *Client) DeleteUomCategorys(ids []int64) error {
	return c.Delete(UomCategoryModel, ids)
}

// GetUomCategory gets uom.category existing record.
func (c *Client) GetUomCategory(id int64) (*UomCategory, error) {
	ucs, err := c.GetUomCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if ucs != nil && len(*ucs) > 0 {
		return &((*ucs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of uom.category not found", id)
}

// GetUomCategorys gets uom.category existing records.
func (c *Client) GetUomCategorys(ids []int64) (*UomCategorys, error) {
	ucs := &UomCategorys{}
	if err := c.Read(UomCategoryModel, ids, nil, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUomCategory finds uom.category record by querying it with criteria.
func (c *Client) FindUomCategory(criteria *Criteria) (*UomCategory, error) {
	ucs := &UomCategorys{}
	if err := c.SearchRead(UomCategoryModel, criteria, NewOptions().Limit(1), ucs); err != nil {
		return nil, err
	}
	if ucs != nil && len(*ucs) > 0 {
		return &((*ucs)[0]), nil
	}
	return nil, fmt.Errorf("uom.category was not found")
}

// FindUomCategorys finds uom.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomCategorys(criteria *Criteria, options *Options) (*UomCategorys, error) {
	ucs := &UomCategorys{}
	if err := c.SearchRead(UomCategoryModel, criteria, options, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUomCategoryIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UomCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUomCategoryId finds record id by querying it with criteria.
func (c *Client) FindUomCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UomCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("uom.category was not found")
}
