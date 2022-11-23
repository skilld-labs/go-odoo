package odoo

import (
	"fmt"
)

// MailBlacklist represents mail.blacklist model.
type MailBlacklist struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Email                    *String   `xmlrpc:"email,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailBlacklists represents array of mail.blacklist model.
type MailBlacklists []MailBlacklist

// MailBlacklistModel is the odoo model name.
const MailBlacklistModel = "mail.blacklist"

// Many2One convert MailBlacklist to *Many2One.
func (mb *MailBlacklist) Many2One() *Many2One {
	return NewMany2One(mb.Id.Get(), "")
}

// CreateMailBlacklist creates a new mail.blacklist model and returns its id.
func (c *Client) CreateMailBlacklist(mb *MailBlacklist) (int64, error) {
	return c.Create(MailBlacklistModel, mb)
}

// UpdateMailBlacklist updates an existing mail.blacklist record.
func (c *Client) UpdateMailBlacklist(mb *MailBlacklist) error {
	return c.UpdateMailBlacklists([]int64{mb.Id.Get()}, mb)
}

// UpdateMailBlacklists updates existing mail.blacklist records.
// All records (represented by IDs) will be updated by mb values.
func (c *Client) UpdateMailBlacklists(ids []int64, mb *MailBlacklist) error {
	return c.Update(MailBlacklistModel, ids, mb)
}

// DeleteMailBlacklist deletes an existing mail.blacklist record.
func (c *Client) DeleteMailBlacklist(id int64) error {
	return c.DeleteMailBlacklists([]int64{id})
}

// DeleteMailBlacklists deletes existing mail.blacklist records.
func (c *Client) DeleteMailBlacklists(ids []int64) error {
	return c.Delete(MailBlacklistModel, ids)
}

// GetMailBlacklist gets mail.blacklist existing record.
func (c *Client) GetMailBlacklist(id int64) (*MailBlacklist, error) {
	mbs, err := c.GetMailBlacklists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.blacklist not found", id)
}

// GetMailBlacklists gets mail.blacklist existing records.
func (c *Client) GetMailBlacklists(ids []int64) (*MailBlacklists, error) {
	mbs := &MailBlacklists{}
	if err := c.Read(MailBlacklistModel, ids, nil, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBlacklist finds mail.blacklist record by querying it with criteria.
func (c *Client) FindMailBlacklist(criteria *Criteria) (*MailBlacklist, error) {
	mbs := &MailBlacklists{}
	if err := c.SearchRead(MailBlacklistModel, criteria, NewOptions().Limit(1), mbs); err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("mail.blacklist was not found")
}

// FindMailBlacklists finds mail.blacklist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklists(criteria *Criteria, options *Options) (*MailBlacklists, error) {
	mbs := &MailBlacklists{}
	if err := c.SearchRead(MailBlacklistModel, criteria, options, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBlacklistIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailBlacklistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailBlacklistId finds record id by querying it with criteria.
func (c *Client) FindMailBlacklistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailBlacklistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.blacklist was not found")
}
