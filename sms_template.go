package odoo

// SmsTemplate represents sms.template model.
type SmsTemplate struct {
	Body            *String   `xmlrpc:"body,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Lang            *String   `xmlrpc:"lang,omitempty"`
	Model           *String   `xmlrpc:"model,omitempty"`
	ModelId         *Many2One `xmlrpc:"model_id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	RenderModel     *String   `xmlrpc:"render_model,omitempty"`
	SidebarActionId *Many2One `xmlrpc:"sidebar_action_id,omitempty"`
	TemplateFs      *String   `xmlrpc:"template_fs,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsTemplates represents array of sms.template model.
type SmsTemplates []SmsTemplate

// SmsTemplateModel is the odoo model name.
const SmsTemplateModel = "sms.template"

// Many2One convert SmsTemplate to *Many2One.
func (st *SmsTemplate) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSmsTemplate creates a new sms.template model and returns its id.
func (c *Client) CreateSmsTemplate(st *SmsTemplate) (int64, error) {
	ids, err := c.CreateSmsTemplates([]*SmsTemplate{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsTemplate creates a new sms.template model and returns its id.
func (c *Client) CreateSmsTemplates(sts []*SmsTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SmsTemplateModel, vv, nil)
}

// UpdateSmsTemplate updates an existing sms.template record.
func (c *Client) UpdateSmsTemplate(st *SmsTemplate) error {
	return c.UpdateSmsTemplates([]int64{st.Id.Get()}, st)
}

// UpdateSmsTemplates updates existing sms.template records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSmsTemplates(ids []int64, st *SmsTemplate) error {
	return c.Update(SmsTemplateModel, ids, st, nil)
}

// DeleteSmsTemplate deletes an existing sms.template record.
func (c *Client) DeleteSmsTemplate(id int64) error {
	return c.DeleteSmsTemplates([]int64{id})
}

// DeleteSmsTemplates deletes existing sms.template records.
func (c *Client) DeleteSmsTemplates(ids []int64) error {
	return c.Delete(SmsTemplateModel, ids)
}

// GetSmsTemplate gets sms.template existing record.
func (c *Client) GetSmsTemplate(id int64) (*SmsTemplate, error) {
	sts, err := c.GetSmsTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// GetSmsTemplates gets sms.template existing records.
func (c *Client) GetSmsTemplates(ids []int64) (*SmsTemplates, error) {
	sts := &SmsTemplates{}
	if err := c.Read(SmsTemplateModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTemplate finds sms.template record by querying it with criteria.
func (c *Client) FindSmsTemplate(criteria *Criteria) (*SmsTemplate, error) {
	sts := &SmsTemplates{}
	if err := c.SearchRead(SmsTemplateModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// FindSmsTemplates finds sms.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplates(criteria *Criteria, options *Options) (*SmsTemplates, error) {
	sts := &SmsTemplates{}
	if err := c.SearchRead(SmsTemplateModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsTemplateModel, criteria, options)
}

// FindSmsTemplateId finds record id by querying it with criteria.
func (c *Client) FindSmsTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
