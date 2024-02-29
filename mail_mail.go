package odoo

// MailMail represents mail.mail model.
type MailMail struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	AttachmentIds        *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorAvatar         *String    `xmlrpc:"author_avatar,omitempty"`
	AuthorId             *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete           *Bool      `xmlrpc:"auto_delete,omitempty"`
	Body                 *String    `xmlrpc:"body,omitempty"`
	BodyHtml             *String    `xmlrpc:"body_html,omitempty"`
	ChannelIds           *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds             *Relation  `xmlrpc:"child_ids,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                 *Time      `xmlrpc:"date,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	EmailCc              *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom            *String    `xmlrpc:"email_from,omitempty"`
	EmailTo              *String    `xmlrpc:"email_to,omitempty"`
	FailureReason        *String    `xmlrpc:"failure_reason,omitempty"`
	FetchmailServerId    *Many2One  `xmlrpc:"fetchmail_server_id,omitempty"`
	Headers              *String    `xmlrpc:"headers,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	MailActivityTypeId   *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailMessageId        *Many2One  `xmlrpc:"mail_message_id,omitempty"`
	MailServerId         *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MailingId            *Many2One  `xmlrpc:"mailing_id,omitempty"`
	MessageId            *String    `xmlrpc:"message_id,omitempty"`
	MessageType          *Selection `xmlrpc:"message_type,omitempty"`
	Model                *String    `xmlrpc:"model,omitempty"`
	Needaction           *Bool      `xmlrpc:"needaction,omitempty"`
	NeedactionPartnerIds *Relation  `xmlrpc:"needaction_partner_ids,omitempty"`
	NoAutoThread         *Bool      `xmlrpc:"no_auto_thread,omitempty"`
	Notification         *Bool      `xmlrpc:"notification,omitempty"`
	NotificationIds      *Relation  `xmlrpc:"notification_ids,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds           *Relation  `xmlrpc:"partner_ids,omitempty"`
	RatingIds            *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingValue          *Float     `xmlrpc:"rating_value,omitempty"`
	RecipientIds         *Relation  `xmlrpc:"recipient_ids,omitempty"`
	RecordName           *String    `xmlrpc:"record_name,omitempty"`
	References           *String    `xmlrpc:"references,omitempty"`
	ReplyTo              *String    `xmlrpc:"reply_to,omitempty"`
	ResId                *Int       `xmlrpc:"res_id,omitempty"`
	ScheduledDate        *String    `xmlrpc:"scheduled_date,omitempty"`
	Starred              *Bool      `xmlrpc:"starred,omitempty"`
	StarredPartnerIds    *Relation  `xmlrpc:"starred_partner_ids,omitempty"`
	State                *Selection `xmlrpc:"state,omitempty"`
	StatisticsIds        *Relation  `xmlrpc:"statistics_ids,omitempty"`
	Subject              *String    `xmlrpc:"subject,omitempty"`
	SubtypeId            *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TrackingValueIds     *Relation  `xmlrpc:"tracking_value_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateMailMails([]*MailMail{mm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMails creates a new mail.mail model and returns its id.
func (c *Client) CreateMailMails(mms []*MailMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range mms {
		vv = append(vv, v)
	}
	return c.Create(MailMailModel, vv, nil)
}

// UpdateMailMail updates an existing mail.mail record.
func (c *Client) UpdateMailMail(mm *MailMail) error {
	return c.UpdateMailMails([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMails updates existing mail.mail records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMails(ids []int64, mm *MailMail) error {
	return c.Update(MailMailModel, ids, mm, nil)
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
	return &((*mms)[0]), nil
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
	return &((*mms)[0]), nil
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

// FindMailMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMailModel, criteria, options)
}

// FindMailMailId finds record id by querying it with criteria.
func (c *Client) FindMailMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
