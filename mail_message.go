package odoo

import (
	"fmt"
)

// MailMessage represents mail.message model.
type MailMessage struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	AttachmentIds        *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorAvatar         *String    `xmlrpc:"author_avatar,omitempty"`
	AuthorId             *Many2One  `xmlrpc:"author_id,omitempty"`
	Body                 *String    `xmlrpc:"body,omitempty"`
	ChannelIds           *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds             *Relation  `xmlrpc:"child_ids,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                 *Time      `xmlrpc:"date,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom            *String    `xmlrpc:"email_from,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	MailActivityTypeId   *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailServerId         *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MessageId            *String    `xmlrpc:"message_id,omitempty"`
	MessageType          *Selection `xmlrpc:"message_type,omitempty"`
	Model                *String    `xmlrpc:"model,omitempty"`
	Needaction           *Bool      `xmlrpc:"needaction,omitempty"`
	NeedactionPartnerIds *Relation  `xmlrpc:"needaction_partner_ids,omitempty"`
	NoAutoThread         *Bool      `xmlrpc:"no_auto_thread,omitempty"`
	NotificationIds      *Relation  `xmlrpc:"notification_ids,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds           *Relation  `xmlrpc:"partner_ids,omitempty"`
	RatingIds            *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingValue          *Float     `xmlrpc:"rating_value,omitempty"`
	RecordName           *String    `xmlrpc:"record_name,omitempty"`
	ReplyTo              *String    `xmlrpc:"reply_to,omitempty"`
	ResId                *Int       `xmlrpc:"res_id,omitempty"`
	Starred              *Bool      `xmlrpc:"starred,omitempty"`
	StarredPartnerIds    *Relation  `xmlrpc:"starred_partner_ids,omitempty"`
	Subject              *String    `xmlrpc:"subject,omitempty"`
	SubtypeId            *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TrackingValueIds     *Relation  `xmlrpc:"tracking_value_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailMessages represents array of mail.message model.
type MailMessages []MailMessage

// MailMessageModel is the odoo model name.
const MailMessageModel = "mail.message"

// Many2One convert MailMessage to *Many2One.
func (mm *MailMessage) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMessage creates a new mail.message model and returns its id.
func (c *Client) CreateMailMessage(mm *MailMessage) (int64, error) {
	return c.Create(MailMessageModel, mm)
}

// UpdateMailMessage updates an existing mail.message record.
func (c *Client) UpdateMailMessage(mm *MailMessage) error {
	return c.UpdateMailMessages([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMessages updates existing mail.message records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMessages(ids []int64, mm *MailMessage) error {
	return c.Update(MailMessageModel, ids, mm)
}

// DeleteMailMessage deletes an existing mail.message record.
func (c *Client) DeleteMailMessage(id int64) error {
	return c.DeleteMailMessages([]int64{id})
}

// DeleteMailMessages deletes existing mail.message records.
func (c *Client) DeleteMailMessages(ids []int64) error {
	return c.Delete(MailMessageModel, ids)
}

// GetMailMessage gets mail.message existing record.
func (c *Client) GetMailMessage(id int64) (*MailMessage, error) {
	mms, err := c.GetMailMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.message not found", id)
}

// GetMailMessages gets mail.message existing records.
func (c *Client) GetMailMessages(ids []int64) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.Read(MailMessageModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessage finds mail.message record by querying it with criteria.
func (c *Client) FindMailMessage(criteria *Criteria) (*MailMessage, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("no mail.message was found with criteria %v", criteria)
}

// FindMailMessages finds mail.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessages(criteria *Criteria, options *Options) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMessageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMessageId finds record id by querying it with criteria.
func (c *Client) FindMailMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no mail.message was found with criteria %v and options %v", criteria, options)
}
