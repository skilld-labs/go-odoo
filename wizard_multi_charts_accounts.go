package odoo

import (
	"fmt"
)

// WizardMultiChartsAccounts represents wizard.multi.charts.accounts model.
type WizardMultiChartsAccounts struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	BankAccountCodePrefix *String   `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankAccountIds        *Relation `xmlrpc:"bank_account_ids,omptempty"`
	CashAccountCodePrefix *String   `xmlrpc:"cash_account_code_prefix,omptempty"`
	ChartTemplateId       *Many2One `xmlrpc:"chart_template_id,omptempty"`
	CodeDigits            *Int      `xmlrpc:"code_digits,omptempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omptempty"`
	CompleteTaxSet        *Bool     `xmlrpc:"complete_tax_set,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId            *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	OnlyOneChartTemplate  *Bool     `xmlrpc:"only_one_chart_template,omptempty"`
	PurchaseTaxId         *Many2One `xmlrpc:"purchase_tax_id,omptempty"`
	PurchaseTaxRate       *Float    `xmlrpc:"purchase_tax_rate,omptempty"`
	SaleTaxId             *Many2One `xmlrpc:"sale_tax_id,omptempty"`
	SaleTaxRate           *Float    `xmlrpc:"sale_tax_rate,omptempty"`
	TransferAccountId     *Many2One `xmlrpc:"transfer_account_id,omptempty"`
	UseAngloSaxon         *Bool     `xmlrpc:"use_anglo_saxon,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(WizardMultiChartsAccountsModel, []interface{}{wmca})
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
	return c.Create(WizardMultiChartsAccountsModel, vv)
}

// UpdateWizardMultiChartsAccounts updates an existing wizard.multi.charts.accounts record.
func (c *Client) UpdateWizardMultiChartsAccounts(wmca *WizardMultiChartsAccounts) error {
	return c.UpdateWizardMultiChartsAccountss([]int64{wmca.Id.Get()}, wmca)
}

// UpdateWizardMultiChartsAccountss updates existing wizard.multi.charts.accounts records.
// All records (represented by ids) will be updated by wmca values.
func (c *Client) UpdateWizardMultiChartsAccountss(ids []int64, wmca *WizardMultiChartsAccounts) error {
	return c.Update(WizardMultiChartsAccountsModel, ids, wmca)
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
	if wmcas != nil && len(*wmcas) > 0 {
		return &((*wmcas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of wizard.multi.charts.accounts not found", id)
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
	if wmcas != nil && len(*wmcas) > 0 {
		return &((*wmcas)[0]), nil
	}
	return nil, fmt.Errorf("wizard.multi.charts.accounts was not found with criteria %v", criteria)
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
	ids, err := c.Search(WizardMultiChartsAccountsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWizardMultiChartsAccountsId finds record id by querying it with criteria.
func (c *Client) FindWizardMultiChartsAccountsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WizardMultiChartsAccountsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("wizard.multi.charts.accounts was not found with criteria %v and options %v", criteria, options)
}
