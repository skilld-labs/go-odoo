package odoo

import (
	"fmt"
)

// MailMail represents mail.mail model.
type MailMail struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	AddSign            *Bool      `xmlrpc:"add_sign,omitempty"`
	AttachmentIds      *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorAvatar       *String    `xmlrpc:"author_avatar,omitempty"`
	AuthorId           *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete         *Bool      `xmlrpc:"auto_delete,omitempty"`
	Body               *String    `xmlrpc:"body,omitempty"`
	BodyHtml           *String    `xmlrpc:"body_html,omitempty"`
	CannedResponseIds  *Relation  `xmlrpc:"canned_response_ids,omitempty"`
	ChannelIds         *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds           *Relation  `xmlrpc:"child_ids,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date               *Time      `xmlrpc:"date,omitempty"`
	Description        *String    `xmlrpc:"description,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	EmailCc            *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom          *String    `xmlrpc:"email_from,omitempty"`
	EmailLayoutXmlid   *String    `xmlrpc:"email_layout_xmlid,omitempty"`
	EmailTo            *String    `xmlrpc:"email_to,omitempty"`
	FailureReason      *String    `xmlrpc:"failure_reason,omitempty"`
	FetchmailServerId  *Many2One  `xmlrpc:"fetchmail_server_id,omitempty"`
	HasError           *Bool      `xmlrpc:"has_error,omitempty"`
	HasSmsError        *Bool      `xmlrpc:"has_sms_error,omitempty"`
	Headers            *String    `xmlrpc:"headers,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	IsInternal         *Bool      `xmlrpc:"is_internal,omitempty"`
	LetterIds          *Relation  `xmlrpc:"letter_ids,omitempty"`
	MailActivityTypeId *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailIds            *Relation  `xmlrpc:"mail_ids,omitempty"`
	MailMessageId      *Many2One  `xmlrpc:"mail_message_id,omitempty"`
	MailServerId       *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MessageId          *String    `xmlrpc:"message_id,omitempty"`
	MessageType        *Selection `xmlrpc:"message_type,omitempty"`
	Model              *String    `xmlrpc:"model,omitempty"`
	ModerationStatus   *Selection `xmlrpc:"moderation_status,omitempty"`
	ModeratorId        *Many2One  `xmlrpc:"moderator_id,omitempty"`
	NeedModeration     *Bool      `xmlrpc:"need_moderation,omitempty"`
	Needaction         *Bool      `xmlrpc:"needaction,omitempty"`
	NoAutoThread       *Bool      `xmlrpc:"no_auto_thread,omitempty"`
	Notification       *Bool      `xmlrpc:"notification,omitempty"`
	NotificationIds    *Relation  `xmlrpc:"notification_ids,omitempty"`
	NotifiedPartnerIds *Relation  `xmlrpc:"notified_partner_ids,omitempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omitempty"`
	RecipientIds       *Relation  `xmlrpc:"recipient_ids,omitempty"`
	RecordName         *String    `xmlrpc:"record_name,omitempty"`
	References         *String    `xmlrpc:"references,omitempty"`
	ReplyTo            *String    `xmlrpc:"reply_to,omitempty"`
	ResId              *Many2One  `xmlrpc:"res_id,omitempty"`
	ScheduledDate      *String    `xmlrpc:"scheduled_date,omitempty"`
	SnailmailError     *Bool      `xmlrpc:"snailmail_error,omitempty"`
	Starred            *Bool      `xmlrpc:"starred,omitempty"`
	StarredPartnerIds  *Relation  `xmlrpc:"starred_partner_ids,omitempty"`
	State              *Selection `xmlrpc:"state,omitempty"`
	Subject            *String    `xmlrpc:"subject,omitempty"`
	SubtypeId          *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TrackingValueIds   *Relation  `xmlrpc:"tracking_value_ids,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailMails represents array of mail.mail model.
type MailMails []MailMail

// MailMailModel is the odoo model name.
const MailMailModel = "mail.mail"

// Many2One convert MailMail to *Many2One.
func (mm *MailMail) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMail creates a new mail.mail model and returns its id.
func (c *Client) CreateMailMail(mm *MailMail) (int64, error) {
	return c.Create(MailMailModel, mm)
}

// UpdateMailMail updates an existing mail.mail record.
func (c *Client) UpdateMailMail(mm *MailMail) error {
	return c.UpdateMailMails([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMails updates existing mail.mail records.
// All records (represented by IDs) will be updated by mm values.
func (c *Client) UpdateMailMails(ids []int64, mm *MailMail) error {
	return c.Update(MailMailModel, ids, mm)
}

// DeleteMailMail deletes an existing mail.mail record.
func (c *Client) DeleteMailMail(id int64) error {
	return c.DeleteMailMails([]int64{id})
}

// DeleteMailMails deletes existing mail.mail records.
func (c *Client) DeleteMailMails(ids []int64) error {
	return c.Delete(MailMailModel, ids)
}

// GetMailMail gets mail.mail existing record.
func (c *Client) GetMailMail(id int64) (*MailMail, error) {
	mms, err := c.GetMailMails([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mail not found", id)
}

// GetMailMails gets mail.mail existing records.
func (c *Client) GetMailMails(ids []int64) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.Read(MailMailModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMail finds mail.mail record by querying it with criteria.
func (c *Client) FindMailMail(criteria *Criteria) (*MailMail, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mail.mail was not found")
}

// FindMailMails finds mail.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMails(criteria *Criteria, options *Options) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMailIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMailId finds record id by querying it with criteria.
func (c *Client) FindMailMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.mail was not found")
}
