package odoo

// AccountChartTemplate represents account.chart.template model.
type AccountChartTemplate struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountChartTemplates represents array of account.chart.template model.
type AccountChartTemplates []AccountChartTemplate

// AccountChartTemplateModel is the odoo model name.
const AccountChartTemplateModel = "account.chart.template"

// Many2One convert AccountChartTemplate to *Many2One.
func (act *AccountChartTemplate) Many2One() *Many2One {
	return NewMany2One(act.Id.Get(), "")
}

// CreateAccountChartTemplate creates a new account.chart.template model and returns its id.
func (c *Client) CreateAccountChartTemplate(act *AccountChartTemplate) (int64, error) {
	ids, err := c.CreateAccountChartTemplates([]*AccountChartTemplate{act})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountChartTemplate creates a new account.chart.template model and returns its id.
func (c *Client) CreateAccountChartTemplates(acts []*AccountChartTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range acts {
		vv = append(vv, v)
	}
	return c.Create(AccountChartTemplateModel, vv, nil)
}

// UpdateAccountChartTemplate updates an existing account.chart.template record.
func (c *Client) UpdateAccountChartTemplate(act *AccountChartTemplate) error {
	return c.UpdateAccountChartTemplates([]int64{act.Id.Get()}, act)
}

// UpdateAccountChartTemplates updates existing account.chart.template records.
// All records (represented by ids) will be updated by act values.
func (c *Client) UpdateAccountChartTemplates(ids []int64, act *AccountChartTemplate) error {
	return c.Update(AccountChartTemplateModel, ids, act, nil)
}

// DeleteAccountChartTemplate deletes an existing account.chart.template record.
func (c *Client) DeleteAccountChartTemplate(id int64) error {
	return c.DeleteAccountChartTemplates([]int64{id})
}

// DeleteAccountChartTemplates deletes existing account.chart.template records.
func (c *Client) DeleteAccountChartTemplates(ids []int64) error {
	return c.Delete(AccountChartTemplateModel, ids)
}

// GetAccountChartTemplate gets account.chart.template existing record.
func (c *Client) GetAccountChartTemplate(id int64) (*AccountChartTemplate, error) {
	acts, err := c.GetAccountChartTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acts)[0]), nil
}

// GetAccountChartTemplates gets account.chart.template existing records.
func (c *Client) GetAccountChartTemplates(ids []int64) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.Read(AccountChartTemplateModel, ids, nil, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplate finds account.chart.template record by querying it with criteria.
func (c *Client) FindAccountChartTemplate(criteria *Criteria) (*AccountChartTemplate, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, NewOptions().Limit(1), acts); err != nil {
		return nil, err
	}
	return &((*acts)[0]), nil
}

// FindAccountChartTemplates finds account.chart.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplates(criteria *Criteria, options *Options) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, options, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountChartTemplateModel, criteria, options)
}

// FindAccountChartTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountChartTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountChartTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
