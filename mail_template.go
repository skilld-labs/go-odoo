package odoo

// MailTemplate represents mail.template model.
type MailTemplate struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	AttachmentIds       *Relation `xmlrpc:"attachment_ids,omitempty"`
	AutoDelete          *Bool     `xmlrpc:"auto_delete,omitempty"`
	BodyHtml            *String   `xmlrpc:"body_html,omitempty"`
	Copyvalue           *String   `xmlrpc:"copyvalue,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	EmailCc             *String   `xmlrpc:"email_cc,omitempty"`
	EmailFrom           *String   `xmlrpc:"email_from,omitempty"`
	EmailTo             *String   `xmlrpc:"email_to,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	Lang                *String   `xmlrpc:"lang,omitempty"`
	MailServerId        *Many2One `xmlrpc:"mail_server_id,omitempty"`
	Model               *String   `xmlrpc:"model,omitempty"`
	ModelId             *Many2One `xmlrpc:"model_id,omitempty"`
	ModelObjectField    *Many2One `xmlrpc:"model_object_field,omitempty"`
	Name                *String   `xmlrpc:"name,omitempty"`
	NullValue           *String   `xmlrpc:"null_value,omitempty"`
	PartnerTo           *String   `xmlrpc:"partner_to,omitempty"`
	RefIrActWindow      *Many2One `xmlrpc:"ref_ir_act_window,omitempty"`
	ReplyTo             *String   `xmlrpc:"reply_to,omitempty"`
	ReportName          *String   `xmlrpc:"report_name,omitempty"`
	ReportTemplate      *Many2One `xmlrpc:"report_template,omitempty"`
	ScheduledDate       *String   `xmlrpc:"scheduled_date,omitempty"`
	SubModelObjectField *Many2One `xmlrpc:"sub_model_object_field,omitempty"`
	SubObject           *Many2One `xmlrpc:"sub_object,omitempty"`
	Subject             *String   `xmlrpc:"subject,omitempty"`
	UseDefaultTo        *Bool     `xmlrpc:"use_default_to,omitempty"`
	UserSignature       *Bool     `xmlrpc:"user_signature,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateMailTemplates([]*MailTemplate{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTemplate creates a new mail.template model and returns its id.
func (c *Client) CreateMailTemplates(mts []*MailTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MailTemplateModel, vv, nil)
}

// UpdateMailTemplate updates an existing mail.template record.
func (c *Client) UpdateMailTemplate(mt *MailTemplate) error {
	return c.UpdateMailTemplates([]int64{mt.Id.Get()}, mt)
}

// UpdateMailTemplates updates existing mail.template records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailTemplates(ids []int64, mt *MailTemplate) error {
	return c.Update(MailTemplateModel, ids, mt, nil)
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
	return &((*mts)[0]), nil
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
	return &((*mts)[0]), nil
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
	return c.Search(MailTemplateModel, criteria, options)
}

// FindMailTemplateId finds record id by querying it with criteria.
func (c *Client) FindMailTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
