package odoo

import (
	"fmt"
)

// MailModeration represents mail.moderation model.
type MailModeration struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	ChannelId   *Many2One  `xmlrpc:"channel_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Email       *String    `xmlrpc:"email,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Status      *Selection `xmlrpc:"status,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailModerations represents array of mail.moderation model.
type MailModerations []MailModeration

// MailModerationModel is the odoo model name.
const MailModerationModel = "mail.moderation"

// Many2One convert MailModeration to *Many2One.
func (mm *MailModeration) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailModeration creates a new mail.moderation model and returns its id.
func (c *Client) CreateMailModeration(mm *MailModeration) (int64, error) {
	return c.Create(MailModerationModel, mm)
}

// UpdateMailModeration updates an existing mail.moderation record.
func (c *Client) UpdateMailModeration(mm *MailModeration) error {
	return c.UpdateMailModerations([]int64{mm.Id.Get()}, mm)
}

// UpdateMailModerations updates existing mail.moderation records.
// All records (represented by IDs) will be updated by mm values.
func (c *Client) UpdateMailModerations(ids []int64, mm *MailModeration) error {
	return c.Update(MailModerationModel, ids, mm)
}

// DeleteMailModeration deletes an existing mail.moderation record.
func (c *Client) DeleteMailModeration(id int64) error {
	return c.DeleteMailModerations([]int64{id})
}

// DeleteMailModerations deletes existing mail.moderation records.
func (c *Client) DeleteMailModerations(ids []int64) error {
	return c.Delete(MailModerationModel, ids)
}

// GetMailModeration gets mail.moderation existing record.
func (c *Client) GetMailModeration(id int64) (*MailModeration, error) {
	mms, err := c.GetMailModerations([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.moderation not found", id)
}

// GetMailModerations gets mail.moderation existing records.
func (c *Client) GetMailModerations(ids []int64) (*MailModerations, error) {
	mms := &MailModerations{}
	if err := c.Read(MailModerationModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailModeration finds mail.moderation record by querying it with criteria.
func (c *Client) FindMailModeration(criteria *Criteria) (*MailModeration, error) {
	mms := &MailModerations{}
	if err := c.SearchRead(MailModerationModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mail.moderation was not found")
}

// FindMailModerations finds mail.moderation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailModerations(criteria *Criteria, options *Options) (*MailModerations, error) {
	mms := &MailModerations{}
	if err := c.SearchRead(MailModerationModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailModerationIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailModerationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailModerationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailModerationId finds record id by querying it with criteria.
func (c *Client) FindMailModerationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailModerationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.moderation was not found")
}
