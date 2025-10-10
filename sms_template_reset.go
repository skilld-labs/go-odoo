package odoo

// SmsTemplateReset represents sms.template.reset model.
type SmsTemplateReset struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	TemplateIds *Relation `xmlrpc:"template_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsTemplateResets represents array of sms.template.reset model.
type SmsTemplateResets []SmsTemplateReset

// SmsTemplateResetModel is the odoo model name.
const SmsTemplateResetModel = "sms.template.reset"

// Many2One convert SmsTemplateReset to *Many2One.
func (str *SmsTemplateReset) Many2One() *Many2One {
	return NewMany2One(str.Id.Get(), "")
}

// CreateSmsTemplateReset creates a new sms.template.reset model and returns its id.
func (c *Client) CreateSmsTemplateReset(str *SmsTemplateReset) (int64, error) {
	ids, err := c.CreateSmsTemplateResets([]*SmsTemplateReset{str})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsTemplateReset creates a new sms.template.reset model and returns its id.
func (c *Client) CreateSmsTemplateResets(strs []*SmsTemplateReset) ([]int64, error) {
	var vv []interface{}
	for _, v := range strs {
		vv = append(vv, v)
	}
	return c.Create(SmsTemplateResetModel, vv, nil)
}

// UpdateSmsTemplateReset updates an existing sms.template.reset record.
func (c *Client) UpdateSmsTemplateReset(str *SmsTemplateReset) error {
	return c.UpdateSmsTemplateResets([]int64{str.Id.Get()}, str)
}

// UpdateSmsTemplateResets updates existing sms.template.reset records.
// All records (represented by ids) will be updated by str values.
func (c *Client) UpdateSmsTemplateResets(ids []int64, str *SmsTemplateReset) error {
	return c.Update(SmsTemplateResetModel, ids, str, nil)
}

// DeleteSmsTemplateReset deletes an existing sms.template.reset record.
func (c *Client) DeleteSmsTemplateReset(id int64) error {
	return c.DeleteSmsTemplateResets([]int64{id})
}

// DeleteSmsTemplateResets deletes existing sms.template.reset records.
func (c *Client) DeleteSmsTemplateResets(ids []int64) error {
	return c.Delete(SmsTemplateResetModel, ids)
}

// GetSmsTemplateReset gets sms.template.reset existing record.
func (c *Client) GetSmsTemplateReset(id int64) (*SmsTemplateReset, error) {
	strs, err := c.GetSmsTemplateResets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*strs)[0]), nil
}

// GetSmsTemplateResets gets sms.template.reset existing records.
func (c *Client) GetSmsTemplateResets(ids []int64) (*SmsTemplateResets, error) {
	strs := &SmsTemplateResets{}
	if err := c.Read(SmsTemplateResetModel, ids, nil, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindSmsTemplateReset finds sms.template.reset record by querying it with criteria.
func (c *Client) FindSmsTemplateReset(criteria *Criteria) (*SmsTemplateReset, error) {
	strs := &SmsTemplateResets{}
	if err := c.SearchRead(SmsTemplateResetModel, criteria, NewOptions().Limit(1), strs); err != nil {
		return nil, err
	}
	return &((*strs)[0]), nil
}

// FindSmsTemplateResets finds sms.template.reset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplateResets(criteria *Criteria, options *Options) (*SmsTemplateResets, error) {
	strs := &SmsTemplateResets{}
	if err := c.SearchRead(SmsTemplateResetModel, criteria, options, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindSmsTemplateResetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplateResetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsTemplateResetModel, criteria, options)
}

// FindSmsTemplateResetId finds record id by querying it with criteria.
func (c *Client) FindSmsTemplateResetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTemplateResetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
