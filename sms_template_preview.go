package odoo

import (
	"fmt"
)

// SmsTemplatePreview represents sms.template.preview model.
type SmsTemplatePreview struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	Body          *String    `xmlrpc:"body,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	Lang          *Selection `xmlrpc:"lang,omitempty"`
	ModelId       *Many2One  `xmlrpc:"model_id,omitempty"`
	NoRecord      *Bool      `xmlrpc:"no_record,omitempty"`
	ResourceRef   *String    `xmlrpc:"resource_ref,omitempty"`
	SmsTemplateId *Many2One  `xmlrpc:"sms_template_id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SmsTemplatePreviews represents array of sms.template.preview model.
type SmsTemplatePreviews []SmsTemplatePreview

// SmsTemplatePreviewModel is the odoo model name.
const SmsTemplatePreviewModel = "sms.template.preview"

// Many2One convert SmsTemplatePreview to *Many2One.
func (stp *SmsTemplatePreview) Many2One() *Many2One {
	return NewMany2One(stp.Id.Get(), "")
}

// CreateSmsTemplatePreview creates a new sms.template.preview model and returns its id.
func (c *Client) CreateSmsTemplatePreview(stp *SmsTemplatePreview) (int64, error) {
	return c.Create(SmsTemplatePreviewModel, stp)
}

// UpdateSmsTemplatePreview updates an existing sms.template.preview record.
func (c *Client) UpdateSmsTemplatePreview(stp *SmsTemplatePreview) error {
	return c.UpdateSmsTemplatePreviews([]int64{stp.Id.Get()}, stp)
}

// UpdateSmsTemplatePreviews updates existing sms.template.preview records.
// All records (represented by IDs) will be updated by stp values.
func (c *Client) UpdateSmsTemplatePreviews(ids []int64, stp *SmsTemplatePreview) error {
	return c.Update(SmsTemplatePreviewModel, ids, stp)
}

// DeleteSmsTemplatePreview deletes an existing sms.template.preview record.
func (c *Client) DeleteSmsTemplatePreview(id int64) error {
	return c.DeleteSmsTemplatePreviews([]int64{id})
}

// DeleteSmsTemplatePreviews deletes existing sms.template.preview records.
func (c *Client) DeleteSmsTemplatePreviews(ids []int64) error {
	return c.Delete(SmsTemplatePreviewModel, ids)
}

// GetSmsTemplatePreview gets sms.template.preview existing record.
func (c *Client) GetSmsTemplatePreview(id int64) (*SmsTemplatePreview, error) {
	stps, err := c.GetSmsTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if stps != nil && len(*stps) > 0 {
		return &((*stps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.template.preview not found", id)
}

// GetSmsTemplatePreviews gets sms.template.preview existing records.
func (c *Client) GetSmsTemplatePreviews(ids []int64) (*SmsTemplatePreviews, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.Read(SmsTemplatePreviewModel, ids, nil, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSmsTemplatePreview finds sms.template.preview record by querying it with criteria.
func (c *Client) FindSmsTemplatePreview(criteria *Criteria) (*SmsTemplatePreview, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.SearchRead(SmsTemplatePreviewModel, criteria, NewOptions().Limit(1), stps); err != nil {
		return nil, err
	}
	if stps != nil && len(*stps) > 0 {
		return &((*stps)[0]), nil
	}
	return nil, fmt.Errorf("sms.template.preview was not found")
}

// FindSmsTemplatePreviews finds sms.template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplatePreviews(criteria *Criteria, options *Options) (*SmsTemplatePreviews, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.SearchRead(SmsTemplatePreviewModel, criteria, options, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSmsTemplatePreviewIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsTemplatePreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindSmsTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.template.preview was not found")
}
