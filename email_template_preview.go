package odoo

// EmailTemplatePreview represents email_template.preview model.
type EmailTemplatePreview struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AttachmentIds       *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AutoDelete          *Bool      `xmlrpc:"auto_delete,omitempty"`
	BodyHtml            *String    `xmlrpc:"body_html,omitempty"`
	Copyvalue           *String    `xmlrpc:"copyvalue,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EmailCc             *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom           *String    `xmlrpc:"email_from,omitempty"`
	EmailTo             *String    `xmlrpc:"email_to,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Lang                *String    `xmlrpc:"lang,omitempty"`
	MailServerId        *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	Model               *String    `xmlrpc:"model,omitempty"`
	ModelId             *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelObjectField    *Many2One  `xmlrpc:"model_object_field,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	NullValue           *String    `xmlrpc:"null_value,omitempty"`
	PartnerIds          *Relation  `xmlrpc:"partner_ids,omitempty"`
	PartnerTo           *String    `xmlrpc:"partner_to,omitempty"`
	RefIrActWindow      *Many2One  `xmlrpc:"ref_ir_act_window,omitempty"`
	ReplyTo             *String    `xmlrpc:"reply_to,omitempty"`
	ReportName          *String    `xmlrpc:"report_name,omitempty"`
	ReportTemplate      *Many2One  `xmlrpc:"report_template,omitempty"`
	ResId               *Selection `xmlrpc:"res_id,omitempty"`
	ScheduledDate       *String    `xmlrpc:"scheduled_date,omitempty"`
	SubModelObjectField *Many2One  `xmlrpc:"sub_model_object_field,omitempty"`
	SubObject           *Many2One  `xmlrpc:"sub_object,omitempty"`
	Subject             *String    `xmlrpc:"subject,omitempty"`
	UseDefaultTo        *Bool      `xmlrpc:"use_default_to,omitempty"`
	UserSignature       *Bool      `xmlrpc:"user_signature,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateEmailTemplatePreviews([]*EmailTemplatePreview{ep})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEmailTemplatePreview creates a new email_template.preview model and returns its id.
func (c *Client) CreateEmailTemplatePreviews(eps []*EmailTemplatePreview) ([]int64, error) {
	var vv []interface{}
	for _, v := range eps {
		vv = append(vv, v)
	}
	return c.Create(EmailTemplatePreviewModel, vv, nil)
}

// UpdateEmailTemplatePreview updates an existing email_template.preview record.
func (c *Client) UpdateEmailTemplatePreview(ep *EmailTemplatePreview) error {
	return c.UpdateEmailTemplatePreviews([]int64{ep.Id.Get()}, ep)
}

// UpdateEmailTemplatePreviews updates existing email_template.preview records.
// All records (represented by ids) will be updated by ep values.
func (c *Client) UpdateEmailTemplatePreviews(ids []int64, ep *EmailTemplatePreview) error {
	return c.Update(EmailTemplatePreviewModel, ids, ep, nil)
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
	return &((*eps)[0]), nil
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
	return &((*eps)[0]), nil
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
	return c.Search(EmailTemplatePreviewModel, criteria, options)
}

// FindEmailTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindEmailTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EmailTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
