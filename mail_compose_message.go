package odoo

import (
	"fmt"
)

// MailComposeMessage represents mail.compose.message model.
type MailComposeMessage struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	ActiveDomain       *String    `xmlrpc:"active_domain,omitempty"`
	AddSign            *Bool      `xmlrpc:"add_sign,omitempty"`
	AttachmentIds      *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorId           *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete         *Bool      `xmlrpc:"auto_delete,omitempty"`
	AutoDeleteMessage  *Bool      `xmlrpc:"auto_delete_message,omitempty"`
	Body               *String    `xmlrpc:"body,omitempty"`
	CompositionMode    *Selection `xmlrpc:"composition_mode,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom          *String    `xmlrpc:"email_from,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	IsLog              *Bool      `xmlrpc:"is_log,omitempty"`
	Layout             *String    `xmlrpc:"layout,omitempty"`
	MailActivityTypeId *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailServerId       *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MessageType        *Selection `xmlrpc:"message_type,omitempty"`
	Model              *String    `xmlrpc:"model,omitempty"`
	NoAutoThread       *Bool      `xmlrpc:"no_auto_thread,omitempty"`
	Notify             *Bool      `xmlrpc:"notify,omitempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omitempty"`
	RecordName         *String    `xmlrpc:"record_name,omitempty"`
	ReplyTo            *String    `xmlrpc:"reply_to,omitempty"`
	ResId              *Int       `xmlrpc:"res_id,omitempty"`
	Subject            *String    `xmlrpc:"subject,omitempty"`
	SubtypeId          *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TemplateId         *Many2One  `xmlrpc:"template_id,omitempty"`
	UseActiveDomain    *Bool      `xmlrpc:"use_active_domain,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailComposeMessages represents array of mail.compose.message model.
type MailComposeMessages []MailComposeMessage

// MailComposeMessageModel is the odoo model name.
const MailComposeMessageModel = "mail.compose.message"

// Many2One convert MailComposeMessage to *Many2One.
func (mcm *MailComposeMessage) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessage(mcm *MailComposeMessage) (int64, error) {
	return c.Create(MailComposeMessageModel, mcm)
}

// UpdateMailComposeMessage updates an existing mail.compose.message record.
func (c *Client) UpdateMailComposeMessage(mcm *MailComposeMessage) error {
	return c.UpdateMailComposeMessages([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposeMessages updates existing mail.compose.message records.
// All records (represented by IDs) will be updated by mcm values.
func (c *Client) UpdateMailComposeMessages(ids []int64, mcm *MailComposeMessage) error {
	return c.Update(MailComposeMessageModel, ids, mcm)
}

// DeleteMailComposeMessage deletes an existing mail.compose.message record.
func (c *Client) DeleteMailComposeMessage(id int64) error {
	return c.DeleteMailComposeMessages([]int64{id})
}

// DeleteMailComposeMessages deletes existing mail.compose.message records.
func (c *Client) DeleteMailComposeMessages(ids []int64) error {
	return c.Delete(MailComposeMessageModel, ids)
}

// GetMailComposeMessage gets mail.compose.message existing record.
func (c *Client) GetMailComposeMessage(id int64) (*MailComposeMessage, error) {
	mcms, err := c.GetMailComposeMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.compose.message not found", id)
}

// GetMailComposeMessages gets mail.compose.message existing records.
func (c *Client) GetMailComposeMessages(ids []int64) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.Read(MailComposeMessageModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessage finds mail.compose.message record by querying it with criteria.
func (c *Client) FindMailComposeMessage(criteria *Criteria) (*MailComposeMessage, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("mail.compose.message was not found")
}

// FindMailComposeMessages finds mail.compose.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessages(criteria *Criteria, options *Options) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessageIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailComposeMessageId finds record id by querying it with criteria.
func (c *Client) FindMailComposeMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.compose.message was not found")
}
