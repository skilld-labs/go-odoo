package odoo

// MailScheduledMessage represents mail.scheduled.message model.
type MailScheduledMessage struct {
	AttachmentIds          *Relation `xmlrpc:"attachment_ids,omitempty"`
	AuthorId               *Many2One `xmlrpc:"author_id,omitempty"`
	Body                   *String   `xmlrpc:"body,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	IsNote                 *Bool     `xmlrpc:"is_note,omitempty"`
	Model                  *String   `xmlrpc:"model,omitempty"`
	NotificationParameters *String   `xmlrpc:"notification_parameters,omitempty"`
	PartnerIds             *Relation `xmlrpc:"partner_ids,omitempty"`
	ResId                  *Many2One `xmlrpc:"res_id,omitempty"`
	ScheduledDate          *Time     `xmlrpc:"scheduled_date,omitempty"`
	Subject                *String   `xmlrpc:"subject,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailScheduledMessages represents array of mail.scheduled.message model.
type MailScheduledMessages []MailScheduledMessage

// MailScheduledMessageModel is the odoo model name.
const MailScheduledMessageModel = "mail.scheduled.message"

// Many2One convert MailScheduledMessage to *Many2One.
func (msm *MailScheduledMessage) Many2One() *Many2One {
	return NewMany2One(msm.Id.Get(), "")
}

// CreateMailScheduledMessage creates a new mail.scheduled.message model and returns its id.
func (c *Client) CreateMailScheduledMessage(msm *MailScheduledMessage) (int64, error) {
	ids, err := c.CreateMailScheduledMessages([]*MailScheduledMessage{msm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailScheduledMessage creates a new mail.scheduled.message model and returns its id.
func (c *Client) CreateMailScheduledMessages(msms []*MailScheduledMessage) ([]int64, error) {
	var vv []interface{}
	for _, v := range msms {
		vv = append(vv, v)
	}
	return c.Create(MailScheduledMessageModel, vv, nil)
}

// UpdateMailScheduledMessage updates an existing mail.scheduled.message record.
func (c *Client) UpdateMailScheduledMessage(msm *MailScheduledMessage) error {
	return c.UpdateMailScheduledMessages([]int64{msm.Id.Get()}, msm)
}

// UpdateMailScheduledMessages updates existing mail.scheduled.message records.
// All records (represented by ids) will be updated by msm values.
func (c *Client) UpdateMailScheduledMessages(ids []int64, msm *MailScheduledMessage) error {
	return c.Update(MailScheduledMessageModel, ids, msm, nil)
}

// DeleteMailScheduledMessage deletes an existing mail.scheduled.message record.
func (c *Client) DeleteMailScheduledMessage(id int64) error {
	return c.DeleteMailScheduledMessages([]int64{id})
}

// DeleteMailScheduledMessages deletes existing mail.scheduled.message records.
func (c *Client) DeleteMailScheduledMessages(ids []int64) error {
	return c.Delete(MailScheduledMessageModel, ids)
}

// GetMailScheduledMessage gets mail.scheduled.message existing record.
func (c *Client) GetMailScheduledMessage(id int64) (*MailScheduledMessage, error) {
	msms, err := c.GetMailScheduledMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*msms)[0]), nil
}

// GetMailScheduledMessages gets mail.scheduled.message existing records.
func (c *Client) GetMailScheduledMessages(ids []int64) (*MailScheduledMessages, error) {
	msms := &MailScheduledMessages{}
	if err := c.Read(MailScheduledMessageModel, ids, nil, msms); err != nil {
		return nil, err
	}
	return msms, nil
}

// FindMailScheduledMessage finds mail.scheduled.message record by querying it with criteria.
func (c *Client) FindMailScheduledMessage(criteria *Criteria) (*MailScheduledMessage, error) {
	msms := &MailScheduledMessages{}
	if err := c.SearchRead(MailScheduledMessageModel, criteria, NewOptions().Limit(1), msms); err != nil {
		return nil, err
	}
	return &((*msms)[0]), nil
}

// FindMailScheduledMessages finds mail.scheduled.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailScheduledMessages(criteria *Criteria, options *Options) (*MailScheduledMessages, error) {
	msms := &MailScheduledMessages{}
	if err := c.SearchRead(MailScheduledMessageModel, criteria, options, msms); err != nil {
		return nil, err
	}
	return msms, nil
}

// FindMailScheduledMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailScheduledMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailScheduledMessageModel, criteria, options)
}

// FindMailScheduledMessageId finds record id by querying it with criteria.
func (c *Client) FindMailScheduledMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailScheduledMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
