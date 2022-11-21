package odoo

import (
	"fmt"
)

// UtmMedium represents utm.medium model.
type UtmMedium struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UtmMediums represents array of utm.medium model.
type UtmMediums []UtmMedium

// UtmMediumModel is the odoo model name.
const UtmMediumModel = "utm.medium"

// Many2One convert UtmMedium to *Many2One.
func (um *UtmMedium) Many2One() *Many2One {
	return NewMany2One(um.Id.Get(), "")
}

// CreateUtmMedium creates a new utm.medium model and returns its id.
func (c *Client) CreateUtmMedium(um *UtmMedium) (int64, error) {
	return c.Create(UtmMediumModel, um)
}

// UpdateUtmMedium updates an existing utm.medium record.
func (c *Client) UpdateUtmMedium(um *UtmMedium) error {
	return c.UpdateUtmMediums([]int64{um.Id.Get()}, um)
}

// UpdateUtmMediums updates existing utm.medium records.
// All records (represented by ids) will be updated by um values.
func (c *Client) UpdateUtmMediums(ids []int64, um *UtmMedium) error {
	return c.Update(UtmMediumModel, ids, um)
}

// DeleteUtmMedium deletes an existing utm.medium record.
func (c *Client) DeleteUtmMedium(id int64) error {
	return c.DeleteUtmMediums([]int64{id})
}

// DeleteUtmMediums deletes existing utm.medium records.
func (c *Client) DeleteUtmMediums(ids []int64) error {
	return c.Delete(UtmMediumModel, ids)
}

// GetUtmMedium gets utm.medium existing record.
func (c *Client) GetUtmMedium(id int64) (*UtmMedium, error) {
	ums, err := c.GetUtmMediums([]int64{id})
	if err != nil {
		return nil, err
	}
	if ums != nil && len(*ums) > 0 {
		return &((*ums)[0]), nil
	}
	return nil, fmt.Errorf("id %v of utm.medium not found", id)
}

// GetUtmMediums gets utm.medium existing records.
func (c *Client) GetUtmMediums(ids []int64) (*UtmMediums, error) {
	ums := &UtmMediums{}
	if err := c.Read(UtmMediumModel, ids, nil, ums); err != nil {
		return nil, err
	}
	return ums, nil
}

// FindUtmMedium finds utm.medium record by querying it with criteria.
func (c *Client) FindUtmMedium(criteria *Criteria) (*UtmMedium, error) {
	ums := &UtmMediums{}
	if err := c.SearchRead(UtmMediumModel, criteria, NewOptions().Limit(1), ums); err != nil {
		return nil, err
	}
	if ums != nil && len(*ums) > 0 {
		return &((*ums)[0]), nil
	}
	return nil, fmt.Errorf("no utm.medium was found with criteria %v", criteria)
}

// FindUtmMediums finds utm.medium records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmMediums(criteria *Criteria, options *Options) (*UtmMediums, error) {
	ums := &UtmMediums{}
	if err := c.SearchRead(UtmMediumModel, criteria, options, ums); err != nil {
		return nil, err
	}
	return ums, nil
}

// FindUtmMediumIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmMediumIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UtmMediumModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUtmMediumId finds record id by querying it with criteria.
func (c *Client) FindUtmMediumId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmMediumModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no utm.medium was found with criteria %v and options %v", criteria, options)
}
