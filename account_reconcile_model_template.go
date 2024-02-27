package odoo

// AccountReconcileModelTemplate represents account.reconcile.model.template model.
type AccountReconcileModelTemplate struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId        *Many2One  `xmlrpc:"account_id,omitempty"`
	Amount           *Float     `xmlrpc:"amount,omitempty"`
	AmountType       *Selection `xmlrpc:"amount_type,omitempty"`
	ChartTemplateId  *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	HasSecondLine    *Bool      `xmlrpc:"has_second_line,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Label            *String    `xmlrpc:"label,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	SecondAccountId  *Many2One  `xmlrpc:"second_account_id,omitempty"`
	SecondAmount     *Float     `xmlrpc:"second_amount,omitempty"`
	SecondAmountType *Selection `xmlrpc:"second_amount_type,omitempty"`
	SecondLabel      *String    `xmlrpc:"second_label,omitempty"`
	SecondTaxId      *Many2One  `xmlrpc:"second_tax_id,omitempty"`
	Sequence         *Int       `xmlrpc:"sequence,omitempty"`
	TaxId            *Many2One  `xmlrpc:"tax_id,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModelTemplates represents array of account.reconcile.model.template model.
type AccountReconcileModelTemplates []AccountReconcileModelTemplate

// AccountReconcileModelTemplateModel is the odoo model name.
const AccountReconcileModelTemplateModel = "account.reconcile.model.template"

// Many2One convert AccountReconcileModelTemplate to *Many2One.
func (armt *AccountReconcileModelTemplate) Many2One() *Many2One {
	return NewMany2One(armt.Id.Get(), "")
}

// CreateAccountReconcileModelTemplate creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) (int64, error) {
	ids, err := c.CreateAccountReconcileModelTemplates([]*AccountReconcileModelTemplate{armt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModelTemplates creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplates(armts []*AccountReconcileModelTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range armts {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelTemplateModel, vv, nil)
}

// UpdateAccountReconcileModelTemplate updates an existing account.reconcile.model.template record.
func (c *Client) UpdateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) error {
	return c.UpdateAccountReconcileModelTemplates([]int64{armt.Id.Get()}, armt)
}

// UpdateAccountReconcileModelTemplates updates existing account.reconcile.model.template records.
// All records (represented by ids) will be updated by armt values.
func (c *Client) UpdateAccountReconcileModelTemplates(ids []int64, armt *AccountReconcileModelTemplate) error {
	return c.Update(AccountReconcileModelTemplateModel, ids, armt, nil)
}

// DeleteAccountReconcileModelTemplate deletes an existing account.reconcile.model.template record.
func (c *Client) DeleteAccountReconcileModelTemplate(id int64) error {
	return c.DeleteAccountReconcileModelTemplates([]int64{id})
}

// DeleteAccountReconcileModelTemplates deletes existing account.reconcile.model.template records.
func (c *Client) DeleteAccountReconcileModelTemplates(ids []int64) error {
	return c.Delete(AccountReconcileModelTemplateModel, ids)
}

// GetAccountReconcileModelTemplate gets account.reconcile.model.template existing record.
func (c *Client) GetAccountReconcileModelTemplate(id int64) (*AccountReconcileModelTemplate, error) {
	armts, err := c.GetAccountReconcileModelTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*armts)[0]), nil
}

// GetAccountReconcileModelTemplates gets account.reconcile.model.template existing records.
func (c *Client) GetAccountReconcileModelTemplates(ids []int64) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.Read(AccountReconcileModelTemplateModel, ids, nil, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplate finds account.reconcile.model.template record by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplate(criteria *Criteria) (*AccountReconcileModelTemplate, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, NewOptions().Limit(1), armts); err != nil {
		return nil, err
	}
	return &((*armts)[0]), nil
}

// FindAccountReconcileModelTemplates finds account.reconcile.model.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplates(criteria *Criteria, options *Options) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, options, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelTemplateModel, criteria, options)
}

// FindAccountReconcileModelTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
