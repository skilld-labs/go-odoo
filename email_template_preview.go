package odoo

import (
	"fmt"
)

// EmailTemplatePreview represents email_template.preview model.
type EmailTemplatePreview struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AttachmentIds       *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AutoDelete          *Bool      `xmlrpc:"auto_delete,omptempty"`
	BodyHtml            *String    `xmlrpc:"body_html,omptempty"`
	Copyvalue           *String    `xmlrpc:"copyvalue,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	EmailCc             *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom           *String    `xmlrpc:"email_from,omptempty"`
	EmailTo             *String    `xmlrpc:"email_to,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	Lang                *String    `xmlrpc:"lang,omptempty"`
	MailServerId        *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	Model               *String    `xmlrpc:"model,omptempty"`
	ModelId             *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelObjectField    *Many2One  `xmlrpc:"model_object_field,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	NullValue           *String    `xmlrpc:"null_value,omptempty"`
	PartnerIds          *Relation  `xmlrpc:"partner_ids,omptempty"`
	PartnerTo           *String    `xmlrpc:"partner_to,omptempty"`
	RefIrActWindow      *Many2One  `xmlrpc:"ref_ir_act_window,omptempty"`
	ReplyTo             *String    `xmlrpc:"reply_to,omptempty"`
	ReportName          *String    `xmlrpc:"report_name,omptempty"`
	ReportTemplate      *Many2One  `xmlrpc:"report_template,omptempty"`
	ResId               *Selection `xmlrpc:"res_id,omptempty"`
	ScheduledDate       *String    `xmlrpc:"scheduled_date,omptempty"`
	SubModelObjectField *Many2One  `xmlrpc:"sub_model_object_field,omptempty"`
	SubObject           *Many2One  `xmlrpc:"sub_object,omptempty"`
	Subject             *String    `xmlrpc:"subject,omptempty"`
	UseDefaultTo        *Bool      `xmlrpc:"use_default_to,omptempty"`
	UserSignature       *Bool      `xmlrpc:"user_signature,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EmailTemplatePreviews represents array of email_template.preview model.
type EmailTemplatePreviews []EmailTemplatePreview

// EmailTemplatePreviewModel is the odoo model name.
const EmailTemplatePreviewModel = "email_template.preview"

// Many2One convert EmailTemplatePreview to *Many2One.
func (ep *EmailTemplatePreview) Many2One() *Many2One {
	return NewMany2One(ep.Id.Get(), "")
}

// CreateEmailTemplatePreview creates a new email_template.preview model and returns its id.
func (c *Client) CreateEmailTemplatePreview(ep *EmailTemplatePreview) (int64, error) {
	return c.Create(EmailTemplatePreviewModel, ep)
}

// UpdateEmailTemplatePreview updates an existing email_template.preview record.
func (c *Client) UpdateEmailTemplatePreview(ep *EmailTemplatePreview) error {
	return c.UpdateEmailTemplatePreviews([]int64{ep.Id.Get()}, ep)
}

// UpdateEmailTemplatePreviews updates existing email_template.preview records.
// All records (represented by ids) will be updated by ep values.
func (c *Client) UpdateEmailTemplatePreviews(ids []int64, ep *EmailTemplatePreview) error {
	return c.Update(EmailTemplatePreviewModel, ids, ep)
}

// DeleteEmailTemplatePreview deletes an existing email_template.preview record.
func (c *Client) DeleteEmailTemplatePreview(id int64) error {
	return c.DeleteEmailTemplatePreviews([]int64{id})
}

// DeleteEmailTemplatePreviews deletes existing email_template.preview records.
func (c *Client) DeleteEmailTemplatePreviews(ids []int64) error {
	return c.Delete(EmailTemplatePreviewModel, ids)
}

// GetEmailTemplatePreview gets email_template.preview existing record.
func (c *Client) GetEmailTemplatePreview(id int64) (*EmailTemplatePreview, error) {
	eps, err := c.GetEmailTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if eps != nil && len(*eps) > 0 {
		return &((*eps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of email_template.preview not found", id)
}

// GetEmailTemplatePreviews gets email_template.preview existing records.
func (c *Client) GetEmailTemplatePreviews(ids []int64) (*EmailTemplatePreviews, error) {
	eps := &EmailTemplatePreviews{}
	if err := c.Read(EmailTemplatePreviewModel, ids, nil, eps); err != nil {
		return nil, err
	}
	return eps, nil
}

// FindEmailTemplatePreview finds email_template.preview record by querying it with criteria.
func (c *Client) FindEmailTemplatePreview(criteria *Criteria) (*EmailTemplatePreview, error) {
	eps := &EmailTemplatePreviews{}
	if err := c.SearchRead(EmailTemplatePreviewModel, criteria, NewOptions().Limit(1), eps); err != nil {
		return nil, err
	}
	if eps != nil && len(*eps) > 0 {
		return &((*eps)[0]), nil
	}
	return nil, fmt.Errorf("no email_template.preview was found with criteria %v", criteria)
}

// FindEmailTemplatePreviews finds email_template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEmailTemplatePreviews(criteria *Criteria, options *Options) (*EmailTemplatePreviews, error) {
	eps := &EmailTemplatePreviews{}
	if err := c.SearchRead(EmailTemplatePreviewModel, criteria, options, eps); err != nil {
		return nil, err
	}
	return eps, nil
}

// FindEmailTemplatePreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEmailTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EmailTemplatePreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEmailTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindEmailTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EmailTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no email_template.preview was found with criteria %v and options %v", criteria, options)
}
