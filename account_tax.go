package odoo

import (
	"fmt"
)

// AccountTax represents account.tax model.
type AccountTax struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId          *Many2One  `xmlrpc:"account_id,omptempty"`
	Active             *Bool      `xmlrpc:"active,omptempty"`
	Amount             *Float     `xmlrpc:"amount,omptempty"`
	AmountType         *Selection `xmlrpc:"amount_type,omptempty"`
	Analytic           *Bool      `xmlrpc:"analytic,omptempty"`
	CashBasisAccount   *Many2One  `xmlrpc:"cash_basis_account,omptempty"`
	ChildrenTaxIds     *Relation  `xmlrpc:"children_tax_ids,omptempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description        *String    `xmlrpc:"description,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	HideTaxExigibility *Bool      `xmlrpc:"hide_tax_exigibility,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IncludeBaseAmount  *Bool      `xmlrpc:"include_base_amount,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	PriceInclude       *Bool      `xmlrpc:"price_include,omptempty"`
	RefundAccountId    *Many2One  `xmlrpc:"refund_account_id,omptempty"`
	Sequence           *Int       `xmlrpc:"sequence,omptempty"`
	TagIds             *Relation  `xmlrpc:"tag_ids,omptempty"`
	TaxAdjustment      *Bool      `xmlrpc:"tax_adjustment,omptempty"`
	TaxExigibility     *Selection `xmlrpc:"tax_exigibility,omptempty"`
	TaxGroupId         *Many2One  `xmlrpc:"tax_group_id,omptempty"`
	TypeTaxUse         *Selection `xmlrpc:"type_tax_use,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountTaxs represents array of account.tax model.
type AccountTaxs []AccountTax

// AccountTaxModel is the odoo model name.
const AccountTaxModel = "account.tax"

// Many2One convert AccountTax to *Many2One.
func (at *AccountTax) Many2One() *Many2One {
	return NewMany2One(at.Id.Get(), "")
}

// CreateAccountTax creates a new account.tax model and returns its id.
func (c *Client) CreateAccountTax(at *AccountTax) (int64, error) {
	return c.Create(AccountTaxModel, at)
}

// UpdateAccountTax updates an existing account.tax record.
func (c *Client) UpdateAccountTax(at *AccountTax) error {
	return c.UpdateAccountTaxs([]int64{at.Id.Get()}, at)
}

// UpdateAccountTaxs updates existing account.tax records.
// All records (represented by ids) will be updated by at values.
func (c *Client) UpdateAccountTaxs(ids []int64, at *AccountTax) error {
	return c.Update(AccountTaxModel, ids, at)
}

// DeleteAccountTax deletes an existing account.tax record.
func (c *Client) DeleteAccountTax(id int64) error {
	return c.DeleteAccountTaxs([]int64{id})
}

// DeleteAccountTaxs deletes existing account.tax records.
func (c *Client) DeleteAccountTaxs(ids []int64) error {
	return c.Delete(AccountTaxModel, ids)
}

// GetAccountTax gets account.tax existing record.
func (c *Client) GetAccountTax(id int64) (*AccountTax, error) {
	ats, err := c.GetAccountTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if ats != nil && len(*ats) > 0 {
		return &((*ats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax not found", id)
}

// GetAccountTaxs gets account.tax existing records.
func (c *Client) GetAccountTaxs(ids []int64) (*AccountTaxs, error) {
	ats := &AccountTaxs{}
	if err := c.Read(AccountTaxModel, ids, nil, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAccountTax finds account.tax record by querying it with criteria.
func (c *Client) FindAccountTax(criteria *Criteria) (*AccountTax, error) {
	ats := &AccountTaxs{}
	if err := c.SearchRead(AccountTaxModel, criteria, NewOptions().Limit(1), ats); err != nil {
		return nil, err
	}
	if ats != nil && len(*ats) > 0 {
		return &((*ats)[0]), nil
	}
	return nil, fmt.Errorf("no account.tax was found with criteria %v", criteria)
}

// FindAccountTaxs finds account.tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxs(criteria *Criteria, options *Options) (*AccountTaxs, error) {
	ats := &AccountTaxs{}
	if err := c.SearchRead(AccountTaxModel, criteria, options, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAccountTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.tax was found with criteria %v and options %v", criteria, options)
}
