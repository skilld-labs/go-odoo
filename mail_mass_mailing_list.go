package odoo

import (
	"fmt"
)

// MailMassMailingList represents mail.mass_mailing.list model.
type MailMassMailingList struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	ContactNbr  *Int      `xmlrpc:"contact_nbr,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMassMailingLists represents array of mail.mass_mailing.list model.
type MailMassMailingLists []MailMassMailingList

// MailMassMailingListModel is the odoo model name.
const MailMassMailingListModel = "mail.mass_mailing.list"

// Many2One convert MailMassMailingList to *Many2One.
func (mml *MailMassMailingList) Many2One() *Many2One {
	return NewMany2One(mml.Id.Get(), "")
}

// CreateMailMassMailingList creates a new mail.mass_mailing.list model and returns its id.
func (c *Client) CreateMailMassMailingList(mml *MailMassMailingList) (int64, error) {
	return c.Create(MailMassMailingListModel, mml)
}

// UpdateMailMassMailingList updates an existing mail.mass_mailing.list record.
func (c *Client) UpdateMailMassMailingList(mml *MailMassMailingList) error {
	return c.UpdateMailMassMailingLists([]int64{mml.Id.Get()}, mml)
}

// UpdateMailMassMailingLists updates existing mail.mass_mailing.list records.
// All records (represented by ids) will be updated by mml values.
func (c *Client) UpdateMailMassMailingLists(ids []int64, mml *MailMassMailingList) error {
	return c.Update(MailMassMailingListModel, ids, mml)
}

// DeleteMailMassMailingList deletes an existing mail.mass_mailing.list record.
func (c *Client) DeleteMailMassMailingList(id int64) error {
	return c.DeleteMailMassMailingLists([]int64{id})
}

// DeleteMailMassMailingLists deletes existing mail.mass_mailing.list records.
func (c *Client) DeleteMailMassMailingLists(ids []int64) error {
	return c.Delete(MailMassMailingListModel, ids)
}

// GetMailMassMailingList gets mail.mass_mailing.list existing record.
func (c *Client) GetMailMassMailingList(id int64) (*MailMassMailingList, error) {
	mmls, err := c.GetMailMassMailingLists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmls != nil && len(*mmls) > 0 {
		return &((*mmls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mass_mailing.list not found", id)
}

// GetMailMassMailingLists gets mail.mass_mailing.list existing records.
func (c *Client) GetMailMassMailingLists(ids []int64) (*MailMassMailingLists, error) {
	mmls := &MailMassMailingLists{}
	if err := c.Read(MailMassMailingListModel, ids, nil, mmls); err != nil {
		return nil, err
	}
	return mmls, nil
}

// FindMailMassMailingList finds mail.mass_mailing.list record by querying it with criteria.
func (c *Client) FindMailMassMailingList(criteria *Criteria) (*MailMassMailingList, error) {
	mmls := &MailMassMailingLists{}
	if err := c.SearchRead(MailMassMailingListModel, criteria, NewOptions().Limit(1), mmls); err != nil {
		return nil, err
	}
	if mmls != nil && len(*mmls) > 0 {
		return &((*mmls)[0]), nil
	}
	return nil, fmt.Errorf("no mail.mass_mailing.list was found with criteria %v", criteria)
}

// FindMailMassMailingLists finds mail.mass_mailing.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingLists(criteria *Criteria, options *Options) (*MailMassMailingLists, error) {
	mmls := &MailMassMailingLists{}
	if err := c.SearchRead(MailMassMailingListModel, criteria, options, mmls); err != nil {
		return nil, err
	}
	return mmls, nil
}

// FindMailMassMailingListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingListIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMassMailingListModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMassMailingListId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no mail.mass_mailing.list was found with criteria %v and options %v", criteria, options)
}
