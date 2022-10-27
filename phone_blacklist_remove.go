package odoo

import (
	"fmt"
)

// PhoneBlacklistRemove represents phone.blacklist.remove model.
type PhoneBlacklistRemove struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Phone       *String   `xmlrpc:"phone,omitempty"`
	Reason      *String   `xmlrpc:"reason,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PhoneBlacklistRemoves represents array of phone.blacklist.remove model.
type PhoneBlacklistRemoves []PhoneBlacklistRemove

// PhoneBlacklistRemoveModel is the odoo model name.
const PhoneBlacklistRemoveModel = "phone.blacklist.remove"

// Many2One convert PhoneBlacklistRemove to *Many2One.
func (pbr *PhoneBlacklistRemove) Many2One() *Many2One {
	return NewMany2One(pbr.Id.Get(), "")
}

// CreatePhoneBlacklistRemove creates a new phone.blacklist.remove model and returns its id.
func (c *Client) CreatePhoneBlacklistRemove(pbr *PhoneBlacklistRemove) (int64, error) {
	return c.Create(PhoneBlacklistRemoveModel, pbr)
}

// UpdatePhoneBlacklistRemove updates an existing phone.blacklist.remove record.
func (c *Client) UpdatePhoneBlacklistRemove(pbr *PhoneBlacklistRemove) error {
	return c.UpdatePhoneBlacklistRemoves([]int64{pbr.Id.Get()}, pbr)
}

// UpdatePhoneBlacklistRemoves updates existing phone.blacklist.remove records.
// All records (represented by ids) will be updated by pbr values.
func (c *Client) UpdatePhoneBlacklistRemoves(ids []int64, pbr *PhoneBlacklistRemove) error {
	return c.Update(PhoneBlacklistRemoveModel, ids, pbr)
}

// DeletePhoneBlacklistRemove deletes an existing phone.blacklist.remove record.
func (c *Client) DeletePhoneBlacklistRemove(id int64) error {
	return c.DeletePhoneBlacklistRemoves([]int64{id})
}

// DeletePhoneBlacklistRemoves deletes existing phone.blacklist.remove records.
func (c *Client) DeletePhoneBlacklistRemoves(ids []int64) error {
	return c.Delete(PhoneBlacklistRemoveModel, ids)
}

// GetPhoneBlacklistRemove gets phone.blacklist.remove existing record.
func (c *Client) GetPhoneBlacklistRemove(id int64) (*PhoneBlacklistRemove, error) {
	pbrs, err := c.GetPhoneBlacklistRemoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if pbrs != nil && len(*pbrs) > 0 {
		return &((*pbrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of phone.blacklist.remove not found", id)
}

// GetPhoneBlacklistRemoves gets phone.blacklist.remove existing records.
func (c *Client) GetPhoneBlacklistRemoves(ids []int64) (*PhoneBlacklistRemoves, error) {
	pbrs := &PhoneBlacklistRemoves{}
	if err := c.Read(PhoneBlacklistRemoveModel, ids, nil, pbrs); err != nil {
		return nil, err
	}
	return pbrs, nil
}

// FindPhoneBlacklistRemove finds phone.blacklist.remove record by querying it with criteria.
func (c *Client) FindPhoneBlacklistRemove(criteria *Criteria) (*PhoneBlacklistRemove, error) {
	pbrs := &PhoneBlacklistRemoves{}
	if err := c.SearchRead(PhoneBlacklistRemoveModel, criteria, NewOptions().Limit(1), pbrs); err != nil {
		return nil, err
	}
	if pbrs != nil && len(*pbrs) > 0 {
		return &((*pbrs)[0]), nil
	}
	return nil, fmt.Errorf("phone.blacklist.remove was not found")
}

// FindPhoneBlacklistRemoves finds phone.blacklist.remove records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneBlacklistRemoves(criteria *Criteria, options *Options) (*PhoneBlacklistRemoves, error) {
	pbrs := &PhoneBlacklistRemoves{}
	if err := c.SearchRead(PhoneBlacklistRemoveModel, criteria, options, pbrs); err != nil {
		return nil, err
	}
	return pbrs, nil
}

// FindPhoneBlacklistRemoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneBlacklistRemoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PhoneBlacklistRemoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPhoneBlacklistRemoveId finds record id by querying it with criteria.
func (c *Client) FindPhoneBlacklistRemoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PhoneBlacklistRemoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("phone.blacklist.remove was not found")
}
