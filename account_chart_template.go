package odoo

// AccountChartTemplate represents account.chart.template model.
type AccountChartTemplate struct {
	LastUpdate                        *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds                        *Relation `xmlrpc:"account_ids,omptempty"`
	BankAccountCodePrefix             *String   `xmlrpc:"bank_account_code_prefix,omptempty"`
	CashAccountCodePrefix             *String   `xmlrpc:"cash_account_code_prefix,omptempty"`
	CodeDigits                        *Int      `xmlrpc:"code_digits,omptempty"`
	CompanyId                         *Many2One `xmlrpc:"company_id,omptempty"`
	CompleteTaxSet                    *Bool     `xmlrpc:"complete_tax_set,omptempty"`
	CreateDate                        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId                        *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName                       *String   `xmlrpc:"display_name,omptempty"`
	ExpenseCurrencyExchangeAccountId  *Many2One `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	Id                                *Int      `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId   *Many2One `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	Name                              *String   `xmlrpc:"name,omptempty"`
	ParentId                          *Many2One `xmlrpc:"parent_id,omptempty"`
	PropertyAccountExpenseCategId     *Many2One `xmlrpc:"property_account_expense_categ_id,omptempty"`
	PropertyAccountExpenseId          *Many2One `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeCategId      *Many2One `xmlrpc:"property_account_income_categ_id,omptempty"`
	PropertyAccountIncomeId           *Many2One `xmlrpc:"property_account_income_id,omptempty"`
	PropertyAccountPayableId          *Many2One `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountReceivableId       *Many2One `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyStockAccountInputCategId  *Many2One `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId *Many2One `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId   *Many2One `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	TaxTemplateIds                    *Relation `xmlrpc:"tax_template_ids,omptempty"`
	TransferAccountId                 *Many2One `xmlrpc:"transfer_account_id,omptempty"`
	UseAngloSaxon                     *Bool     `xmlrpc:"use_anglo_saxon,omptempty"`
	Visible                           *Bool     `xmlrpc:"visible,omptempty"`
	WriteDate                         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One `xmlrpc:"write_uid,omptempty"`
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
