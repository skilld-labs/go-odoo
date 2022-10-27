package odoo

import (
	"fmt"
)

// MailThreadCc represents mail.thread.cc model.
type MailThreadCc struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	EmailCc                  *String   `xmlrpc:"email_cc,omitempty"`
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
}

// MailThreadCcs represents array of mail.thread.cc model.
type MailThreadCcs []MailThreadCc

// MailThreadCcModel is the odoo model name.
const MailThreadCcModel = "mail.thread.cc"

// Many2One convert MailThreadCc to *Many2One.
func (mtc *MailThreadCc) Many2One() *Many2One {
	return NewMany2One(mtc.Id.Get(), "")
}

// CreateMailThreadCc creates a new mail.thread.cc model and returns its id.
func (c *Client) CreateMailThreadCc(mtc *MailThreadCc) (int64, error) {
	return c.Create(MailThreadCcModel, mtc)
}

// UpdateMailThreadCc updates an existing mail.thread.cc record.
func (c *Client) UpdateMailThreadCc(mtc *MailThreadCc) error {
	return c.UpdateMailThreadCcs([]int64{mtc.Id.Get()}, mtc)
}

// UpdateMailThreadCcs updates existing mail.thread.cc records.
// All records (represented by ids) will be updated by mtc values.
func (c *Client) UpdateMailThreadCcs(ids []int64, mtc *MailThreadCc) error {
	return c.Update(MailThreadCcModel, ids, mtc)
}

// DeleteMailThreadCc deletes an existing mail.thread.cc record.
func (c *Client) DeleteMailThreadCc(id int64) error {
	return c.DeleteMailThreadCcs([]int64{id})
}

// DeleteMailThreadCcs deletes existing mail.thread.cc records.
func (c *Client) DeleteMailThreadCcs(ids []int64) error {
	return c.Delete(MailThreadCcModel, ids)
}

// GetMailThreadCc gets mail.thread.cc existing record.
func (c *Client) GetMailThreadCc(id int64) (*MailThreadCc, error) {
	mtcs, err := c.GetMailThreadCcs([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtcs != nil && len(*mtcs) > 0 {
		return &((*mtcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.thread.cc not found", id)
}

// GetMailThreadCcs gets mail.thread.cc existing records.
func (c *Client) GetMailThreadCcs(ids []int64) (*MailThreadCcs, error) {
	mtcs := &MailThreadCcs{}
	if err := c.Read(MailThreadCcModel, ids, nil, mtcs); err != nil {
		return nil, err
	}
	return mtcs, nil
}

// FindMailThreadCc finds mail.thread.cc record by querying it with criteria.
func (c *Client) FindMailThreadCc(criteria *Criteria) (*MailThreadCc, error) {
	mtcs := &MailThreadCcs{}
	if err := c.SearchRead(MailThreadCcModel, criteria, NewOptions().Limit(1), mtcs); err != nil {
		return nil, err
	}
	if mtcs != nil && len(*mtcs) > 0 {
		return &((*mtcs)[0]), nil
	}
	return nil, fmt.Errorf("mail.thread.cc was not found")
}

// FindMailThreadCcs finds mail.thread.cc records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadCcs(criteria *Criteria, options *Options) (*MailThreadCcs, error) {
	mtcs := &MailThreadCcs{}
	if err := c.SearchRead(MailThreadCcModel, criteria, options, mtcs); err != nil {
		return nil, err
	}
	return mtcs, nil
}

// FindMailThreadCcIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadCcIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailThreadCcModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailThreadCcId finds record id by querying it with criteria.
func (c *Client) FindMailThreadCcId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadCcModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.thread.cc was not found")
}
