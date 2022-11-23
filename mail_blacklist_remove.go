package odoo

import (
	"fmt"
)

// MailBlacklistRemove represents mail.blacklist.remove model.
type MailBlacklistRemove struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Email       *String   `xmlrpc:"email,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Reason      *String   `xmlrpc:"reason,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailBlacklistRemoves represents array of mail.blacklist.remove model.
type MailBlacklistRemoves []MailBlacklistRemove

// MailBlacklistRemoveModel is the odoo model name.
const MailBlacklistRemoveModel = "mail.blacklist.remove"

// Many2One convert MailBlacklistRemove to *Many2One.
func (mbr *MailBlacklistRemove) Many2One() *Many2One {
	return NewMany2One(mbr.Id.Get(), "")
}

// CreateMailBlacklistRemove creates a new mail.blacklist.remove model and returns its id.
func (c *Client) CreateMailBlacklistRemove(mbr *MailBlacklistRemove) (int64, error) {
	return c.Create(MailBlacklistRemoveModel, mbr)
}

// UpdateMailBlacklistRemove updates an existing mail.blacklist.remove record.
func (c *Client) UpdateMailBlacklistRemove(mbr *MailBlacklistRemove) error {
	return c.UpdateMailBlacklistRemoves([]int64{mbr.Id.Get()}, mbr)
}

// UpdateMailBlacklistRemoves updates existing mail.blacklist.remove records.
// All records (represented by IDs) will be updated by mbr values.
func (c *Client) UpdateMailBlacklistRemoves(ids []int64, mbr *MailBlacklistRemove) error {
	return c.Update(MailBlacklistRemoveModel, ids, mbr)
}

// DeleteMailBlacklistRemove deletes an existing mail.blacklist.remove record.
func (c *Client) DeleteMailBlacklistRemove(id int64) error {
	return c.DeleteMailBlacklistRemoves([]int64{id})
}

// DeleteMailBlacklistRemoves deletes existing mail.blacklist.remove records.
func (c *Client) DeleteMailBlacklistRemoves(ids []int64) error {
	return c.Delete(MailBlacklistRemoveModel, ids)
}

// GetMailBlacklistRemove gets mail.blacklist.remove existing record.
func (c *Client) GetMailBlacklistRemove(id int64) (*MailBlacklistRemove, error) {
	mbrs, err := c.GetMailBlacklistRemoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if mbrs != nil && len(*mbrs) > 0 {
		return &((*mbrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.blacklist.remove not found", id)
}

// GetMailBlacklistRemoves gets mail.blacklist.remove existing records.
func (c *Client) GetMailBlacklistRemoves(ids []int64) (*MailBlacklistRemoves, error) {
	mbrs := &MailBlacklistRemoves{}
	if err := c.Read(MailBlacklistRemoveModel, ids, nil, mbrs); err != nil {
		return nil, err
	}
	return mbrs, nil
}

// FindMailBlacklistRemove finds mail.blacklist.remove record by querying it with criteria.
func (c *Client) FindMailBlacklistRemove(criteria *Criteria) (*MailBlacklistRemove, error) {
	mbrs := &MailBlacklistRemoves{}
	if err := c.SearchRead(MailBlacklistRemoveModel, criteria, NewOptions().Limit(1), mbrs); err != nil {
		return nil, err
	}
	if mbrs != nil && len(*mbrs) > 0 {
		return &((*mbrs)[0]), nil
	}
	return nil, fmt.Errorf("mail.blacklist.remove was not found")
}

// FindMailBlacklistRemoves finds mail.blacklist.remove records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklistRemoves(criteria *Criteria, options *Options) (*MailBlacklistRemoves, error) {
	mbrs := &MailBlacklistRemoves{}
	if err := c.SearchRead(MailBlacklistRemoveModel, criteria, options, mbrs); err != nil {
		return nil, err
	}
	return mbrs, nil
}

// FindMailBlacklistRemoveIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklistRemoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailBlacklistRemoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailBlacklistRemoveId finds record id by querying it with criteria.
func (c *Client) FindMailBlacklistRemoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailBlacklistRemoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.blacklist.remove was not found")
}
