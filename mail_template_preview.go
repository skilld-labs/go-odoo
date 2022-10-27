package odoo

import (
	"fmt"
)

// MailTemplatePreview represents mail.template.preview model.
type MailTemplatePreview struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	AttachmentIds  *Relation  `xmlrpc:"attachment_ids,omitempty"`
	BodyHtml       *String    `xmlrpc:"body_html,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	EmailCc        *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom      *String    `xmlrpc:"email_from,omitempty"`
	EmailTo        *String    `xmlrpc:"email_to,omitempty"`
	ErrorMsg       *String    `xmlrpc:"error_msg,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Lang           *Selection `xmlrpc:"lang,omitempty"`
	MailTemplateId *Many2One  `xmlrpc:"mail_template_id,omitempty"`
	ModelId        *Many2One  `xmlrpc:"model_id,omitempty"`
	NoRecord       *Bool      `xmlrpc:"no_record,omitempty"`
	PartnerIds     *Relation  `xmlrpc:"partner_ids,omitempty"`
	ReplyTo        *String    `xmlrpc:"reply_to,omitempty"`
	ResourceRef    *String    `xmlrpc:"resource_ref,omitempty"`
	ScheduledDate  *String    `xmlrpc:"scheduled_date,omitempty"`
	Subject        *String    `xmlrpc:"subject,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailTemplatePreviews represents array of mail.template.preview model.
type MailTemplatePreviews []MailTemplatePreview

// MailTemplatePreviewModel is the odoo model name.
const MailTemplatePreviewModel = "mail.template.preview"

// Many2One convert MailTemplatePreview to *Many2One.
func (mtp *MailTemplatePreview) Many2One() *Many2One {
	return NewMany2One(mtp.Id.Get(), "")
}

// CreateMailTemplatePreview creates a new mail.template.preview model and returns its id.
func (c *Client) CreateMailTemplatePreview(mtp *MailTemplatePreview) (int64, error) {
	return c.Create(MailTemplatePreviewModel, mtp)
}

// UpdateMailTemplatePreview updates an existing mail.template.preview record.
func (c *Client) UpdateMailTemplatePreview(mtp *MailTemplatePreview) error {
	return c.UpdateMailTemplatePreviews([]int64{mtp.Id.Get()}, mtp)
}

// UpdateMailTemplatePreviews updates existing mail.template.preview records.
// All records (represented by ids) will be updated by mtp values.
func (c *Client) UpdateMailTemplatePreviews(ids []int64, mtp *MailTemplatePreview) error {
	return c.Update(MailTemplatePreviewModel, ids, mtp)
}

// DeleteMailTemplatePreview deletes an existing mail.template.preview record.
func (c *Client) DeleteMailTemplatePreview(id int64) error {
	return c.DeleteMailTemplatePreviews([]int64{id})
}

// DeleteMailTemplatePreviews deletes existing mail.template.preview records.
func (c *Client) DeleteMailTemplatePreviews(ids []int64) error {
	return c.Delete(MailTemplatePreviewModel, ids)
}

// GetMailTemplatePreview gets mail.template.preview existing record.
func (c *Client) GetMailTemplatePreview(id int64) (*MailTemplatePreview, error) {
	mtps, err := c.GetMailTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.template.preview not found", id)
}

// GetMailTemplatePreviews gets mail.template.preview existing records.
func (c *Client) GetMailTemplatePreviews(ids []int64) (*MailTemplatePreviews, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.Read(MailTemplatePreviewModel, ids, nil, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailTemplatePreview finds mail.template.preview record by querying it with criteria.
func (c *Client) FindMailTemplatePreview(criteria *Criteria) (*MailTemplatePreview, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.SearchRead(MailTemplatePreviewModel, criteria, NewOptions().Limit(1), mtps); err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("mail.template.preview was not found")
}

// FindMailTemplatePreviews finds mail.template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplatePreviews(criteria *Criteria, options *Options) (*MailTemplatePreviews, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.SearchRead(MailTemplatePreviewModel, criteria, options, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailTemplatePreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTemplatePreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindMailTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.template.preview was not found")
}
