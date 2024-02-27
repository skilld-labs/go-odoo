package odoo

// WizardMultiChartsAccounts represents wizard.multi.charts.accounts model.
type WizardMultiChartsAccounts struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omitempty"`
	BankAccountCodePrefix *String   `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankAccountIds        *Relation `xmlrpc:"bank_account_ids,omitempty"`
	CashAccountCodePrefix *String   `xmlrpc:"cash_account_code_prefix,omitempty"`
	ChartTemplateId       *Many2One `xmlrpc:"chart_template_id,omitempty"`
	CodeDigits            *Int      `xmlrpc:"code_digits,omitempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omitempty"`
	CompleteTaxSet        *Bool     `xmlrpc:"complete_tax_set,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId            *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	OnlyOneChartTemplate  *Bool     `xmlrpc:"only_one_chart_template,omitempty"`
	PurchaseTaxId         *Many2One `xmlrpc:"purchase_tax_id,omitempty"`
	PurchaseTaxRate       *Float    `xmlrpc:"purchase_tax_rate,omitempty"`
	SaleTaxId             *Many2One `xmlrpc:"sale_tax_id,omitempty"`
	SaleTaxRate           *Float    `xmlrpc:"sale_tax_rate,omitempty"`
	TransferAccountId     *Many2One `xmlrpc:"transfer_account_id,omitempty"`
	UseAngloSaxon         *Bool     `xmlrpc:"use_anglo_saxon,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WizardMultiChartsAccountss represents array of wizard.multi.charts.accounts model.
type WizardMultiChartsAccountss []WizardMultiChartsAccounts

// WizardMultiChartsAccountsModel is the odoo model name.
const WizardMultiChartsAccountsModel = "wizard.multi.charts.accounts"

// Many2One convert WizardMultiChartsAccounts to *Many2One.
func (wmca *WizardMultiChartsAccounts) Many2One() *Many2One {
	return NewMany2One(wmca.Id.Get(), "")
}

// CreateWizardMultiChartsAccounts creates a new wizard.multi.charts.accounts model and returns its id.
func (c *Client) CreateWizardMultiChartsAccounts(wmca *WizardMultiChartsAccounts) (int64, error) {
	ids, err := c.CreateWizardMultiChartsAccountss([]*WizardMultiChartsAccounts{wmca})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWizardMultiChartsAccounts creates a new wizard.multi.charts.accounts model and returns its id.
func (c *Client) CreateWizardMultiChartsAccountss(wmcas []*WizardMultiChartsAccounts) ([]int64, error) {
	var vv []interface{}
	for _, v := range wmcas {
		vv = append(vv, v)
	}
	return c.Create(WizardMultiChartsAccountsModel, vv, nil)
}

// UpdateWizardMultiChartsAccounts updates an existing wizard.multi.charts.accounts record.
func (c *Client) UpdateWizardMultiChartsAccounts(wmca *WizardMultiChartsAccounts) error {
	return c.UpdateWizardMultiChartsAccountss([]int64{wmca.Id.Get()}, wmca)
}

// UpdateWizardMultiChartsAccountss updates existing wizard.multi.charts.accounts records.
// All records (represented by ids) will be updated by wmca values.
func (c *Client) UpdateWizardMultiChartsAccountss(ids []int64, wmca *WizardMultiChartsAccounts) error {
	return c.Update(WizardMultiChartsAccountsModel, ids, wmca, nil)
}

// DeleteWizardMultiChartsAccounts deletes an existing wizard.multi.charts.accounts record.
func (c *Client) DeleteWizardMultiChartsAccounts(id int64) error {
	return c.DeleteWizardMultiChartsAccountss([]int64{id})
}

// DeleteWizardMultiChartsAccountss deletes existing wizard.multi.charts.accounts records.
func (c *Client) DeleteWizardMultiChartsAccountss(ids []int64) error {
	return c.Delete(WizardMultiChartsAccountsModel, ids)
}

// GetWizardMultiChartsAccounts gets wizard.multi.charts.accounts existing record.
func (c *Client) GetWizardMultiChartsAccounts(id int64) (*WizardMultiChartsAccounts, error) {
	wmcas, err := c.GetWizardMultiChartsAccountss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wmcas)[0]), nil
}

// GetWizardMultiChartsAccountss gets wizard.multi.charts.accounts existing records.
func (c *Client) GetWizardMultiChartsAccountss(ids []int64) (*WizardMultiChartsAccountss, error) {
	wmcas := &WizardMultiChartsAccountss{}
	if err := c.Read(WizardMultiChartsAccountsModel, ids, nil, wmcas); err != nil {
		return nil, err
	}
	return wmcas, nil
}

// FindWizardMultiChartsAccounts finds wizard.multi.charts.accounts record by querying it with criteria.
func (c *Client) FindWizardMultiChartsAccounts(criteria *Criteria) (*WizardMultiChartsAccounts, error) {
	wmcas := &WizardMultiChartsAccountss{}
	if err := c.SearchRead(WizardMultiChartsAccountsModel, criteria, NewOptions().Limit(1), wmcas); err != nil {
		return nil, err
	}
	return &((*wmcas)[0]), nil
}

// FindWizardMultiChartsAccountss finds wizard.multi.charts.accounts records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardMultiChartsAccountss(criteria *Criteria, options *Options) (*WizardMultiChartsAccountss, error) {
	wmcas := &WizardMultiChartsAccountss{}
	if err := c.SearchRead(WizardMultiChartsAccountsModel, criteria, options, wmcas); err != nil {
		return nil, err
	}
	return wmcas, nil
}

// FindWizardMultiChartsAccountsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardMultiChartsAccountsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WizardMultiChartsAccountsModel, criteria, options)
}

// FindWizardMultiChartsAccountsId finds record id by querying it with criteria.
func (c *Client) FindWizardMultiChartsAccountsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WizardMultiChartsAccountsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
