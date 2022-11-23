package odoo

import (
	"fmt"
)

// MailTemplate represents mail.template model.
type MailTemplate struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	AttachmentIds       *Relation `xmlrpc:"attachment_ids,omptempty"`
	AutoDelete          *Bool     `xmlrpc:"auto_delete,omptempty"`
	BodyHtml            *String   `xmlrpc:"body_html,omptempty"`
	Copyvalue           *String   `xmlrpc:"copyvalue,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	EmailCc             *String   `xmlrpc:"email_cc,omptempty"`
	EmailFrom           *String   `xmlrpc:"email_from,omptempty"`
	EmailTo             *String   `xmlrpc:"email_to,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	Lang                *String   `xmlrpc:"lang,omptempty"`
	MailServerId        *Many2One `xmlrpc:"mail_server_id,omptempty"`
	Model               *String   `xmlrpc:"model,omptempty"`
	ModelId             *Many2One `xmlrpc:"model_id,omptempty"`
	ModelObjectField    *Many2One `xmlrpc:"model_object_field,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	NullValue           *String   `xmlrpc:"null_value,omptempty"`
	PartnerTo           *String   `xmlrpc:"partner_to,omptempty"`
	RefIrActWindow      *Many2One `xmlrpc:"ref_ir_act_window,omptempty"`
	ReplyTo             *String   `xmlrpc:"reply_to,omptempty"`
	ReportName          *String   `xmlrpc:"report_name,omptempty"`
	ReportTemplate      *Many2One `xmlrpc:"report_template,omptempty"`
	ScheduledDate       *String   `xmlrpc:"scheduled_date,omptempty"`
	SubModelObjectField *Many2One `xmlrpc:"sub_model_object_field,omptempty"`
	SubObject           *Many2One `xmlrpc:"sub_object,omptempty"`
	Subject             *String   `xmlrpc:"subject,omptempty"`
	UseDefaultTo        *Bool     `xmlrpc:"use_default_to,omptempty"`
	UserSignature       *Bool     `xmlrpc:"user_signature,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailTemplates represents array of mail.template model.
type MailTemplates []MailTemplate

// MailTemplateModel is the odoo model name.
const MailTemplateModel = "mail.template"

// Many2One convert MailTemplate to *Many2One.
func (mt *MailTemplate) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMailTemplate creates a new mail.template model and returns its id.
func (c *Client) CreateMailTemplate(mt *MailTemplate) (int64, error) {
	return c.Create(MailTemplateModel, mt)
}

// UpdateMailTemplate updates an existing mail.template record.
func (c *Client) UpdateMailTemplate(mt *MailTemplate) error {
	return c.UpdateMailTemplates([]int64{mt.Id.Get()}, mt)
}

// UpdateMailTemplates updates existing mail.template records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailTemplates(ids []int64, mt *MailTemplate) error {
	return c.Update(MailTemplateModel, ids, mt)
}

// DeleteMailTemplate deletes an existing mail.template record.
func (c *Client) DeleteMailTemplate(id int64) error {
	return c.DeleteMailTemplates([]int64{id})
}

// DeleteMailTemplates deletes existing mail.template records.
func (c *Client) DeleteMailTemplates(ids []int64) error {
	return c.Delete(MailTemplateModel, ids)
}

// GetMailTemplate gets mail.template existing record.
func (c *Client) GetMailTemplate(id int64) (*MailTemplate, error) {
	mts, err := c.GetMailTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.template not found", id)
}

// GetMailTemplates gets mail.template existing records.
func (c *Client) GetMailTemplates(ids []int64) (*MailTemplates, error) {
	mts := &MailTemplates{}
	if err := c.Read(MailTemplateModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailTemplate finds mail.template record by querying it with criteria.
func (c *Client) FindMailTemplate(criteria *Criteria) (*MailTemplate, error) {
	mts := &MailTemplates{}
	if err := c.SearchRead(MailTemplateModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("mail.template was not found with criteria %v", criteria)
}

// FindMailTemplates finds mail.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplates(criteria *Criteria, options *Options) (*MailTemplates, error) {
	mts := &MailTemplates{}
	if err := c.SearchRead(MailTemplateModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTemplateId finds record id by querying it with criteria.
func (c *Client) FindMailTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.template was not found with criteria %v and options %v", criteria, options)
}
