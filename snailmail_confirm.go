package odoo

import (
	"fmt"
)

// SnailmailConfirm represents snailmail.confirm model.
type SnailmailConfirm struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	ModelName   *String `xmlrpc:"model_name,omitempty"`
}

// SnailmailConfirms represents array of snailmail.confirm model.
type SnailmailConfirms []SnailmailConfirm

// SnailmailConfirmModel is the odoo model name.
const SnailmailConfirmModel = "snailmail.confirm"

// Many2One convert SnailmailConfirm to *Many2One.
func (sc *SnailmailConfirm) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSnailmailConfirm creates a new snailmail.confirm model and returns its id.
func (c *Client) CreateSnailmailConfirm(sc *SnailmailConfirm) (int64, error) {
	return c.Create(SnailmailConfirmModel, sc)
}

// UpdateSnailmailConfirm updates an existing snailmail.confirm record.
func (c *Client) UpdateSnailmailConfirm(sc *SnailmailConfirm) error {
	return c.UpdateSnailmailConfirms([]int64{sc.Id.Get()}, sc)
}

// UpdateSnailmailConfirms updates existing snailmail.confirm records.
// All records (represented by IDs) will be updated by sc values.
func (c *Client) UpdateSnailmailConfirms(ids []int64, sc *SnailmailConfirm) error {
	return c.Update(SnailmailConfirmModel, ids, sc)
}

// DeleteSnailmailConfirm deletes an existing snailmail.confirm record.
func (c *Client) DeleteSnailmailConfirm(id int64) error {
	return c.DeleteSnailmailConfirms([]int64{id})
}

// DeleteSnailmailConfirms deletes existing snailmail.confirm records.
func (c *Client) DeleteSnailmailConfirms(ids []int64) error {
	return c.Delete(SnailmailConfirmModel, ids)
}

// GetSnailmailConfirm gets snailmail.confirm existing record.
func (c *Client) GetSnailmailConfirm(id int64) (*SnailmailConfirm, error) {
	scs, err := c.GetSnailmailConfirms([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of snailmail.confirm not found", id)
}

// GetSnailmailConfirms gets snailmail.confirm existing records.
func (c *Client) GetSnailmailConfirms(ids []int64) (*SnailmailConfirms, error) {
	scs := &SnailmailConfirms{}
	if err := c.Read(SnailmailConfirmModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSnailmailConfirm finds snailmail.confirm record by querying it with criteria.
func (c *Client) FindSnailmailConfirm(criteria *Criteria) (*SnailmailConfirm, error) {
	scs := &SnailmailConfirms{}
	if err := c.SearchRead(SnailmailConfirmModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("snailmail.confirm was not found")
}

// FindSnailmailConfirms finds snailmail.confirm records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailConfirms(criteria *Criteria, options *Options) (*SnailmailConfirms, error) {
	scs := &SnailmailConfirms{}
	if err := c.SearchRead(SnailmailConfirmModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSnailmailConfirmIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailConfirmIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SnailmailConfirmModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSnailmailConfirmId finds record id by querying it with criteria.
func (c *Client) FindSnailmailConfirmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailConfirmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("snailmail.confirm was not found")
}
