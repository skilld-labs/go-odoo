package odoo

import (
	"fmt"
)

// AccountInvoiceTax represents account.invoice.tax model.
type AccountInvoiceTax struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId *Many2One `xmlrpc:"account_analytic_id,omptempty"`
	AccountId         *Many2One `xmlrpc:"account_id,omptempty"`
	Amount            *Float    `xmlrpc:"amount,omptempty"`
	AmountRounding    *Float    `xmlrpc:"amount_rounding,omptempty"`
	AmountTotal       *Float    `xmlrpc:"amount_total,omptempty"`
	Base              *Float    `xmlrpc:"base,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	InvoiceId         *Many2One `xmlrpc:"invoice_id,omptempty"`
	Manual            *Bool     `xmlrpc:"manual,omptempty"`
	Name              *String   `xmlrpc:"name,omptempty"`
	Sequence          *Int      `xmlrpc:"sequence,omptempty"`
	TaxId             *Many2One `xmlrpc:"tax_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountInvoiceTaxs represents array of account.invoice.tax model.
type AccountInvoiceTaxs []AccountInvoiceTax

// AccountInvoiceTaxModel is the odoo model name.
const AccountInvoiceTaxModel = "account.invoice.tax"

// Many2One convert AccountInvoiceTax to *Many2One.
func (ait *AccountInvoiceTax) Many2One() *Many2One {
	return NewMany2One(ait.Id.Get(), "")
}

// CreateAccountInvoiceTax creates a new account.invoice.tax model and returns its id.
func (c *Client) CreateAccountInvoiceTax(ait *AccountInvoiceTax) (int64, error) {
	return c.Create(AccountInvoiceTaxModel, ait)
}

// UpdateAccountInvoiceTax updates an existing account.invoice.tax record.
func (c *Client) UpdateAccountInvoiceTax(ait *AccountInvoiceTax) error {
	return c.UpdateAccountInvoiceTaxs([]int64{ait.Id.Get()}, ait)
}

// UpdateAccountInvoiceTaxs updates existing account.invoice.tax records.
// All records (represented by ids) will be updated by ait values.
func (c *Client) UpdateAccountInvoiceTaxs(ids []int64, ait *AccountInvoiceTax) error {
	return c.Update(AccountInvoiceTaxModel, ids, ait)
}

// DeleteAccountInvoiceTax deletes an existing account.invoice.tax record.
func (c *Client) DeleteAccountInvoiceTax(id int64) error {
	return c.DeleteAccountInvoiceTaxs([]int64{id})
}

// DeleteAccountInvoiceTaxs deletes existing account.invoice.tax records.
func (c *Client) DeleteAccountInvoiceTaxs(ids []int64) error {
	return c.Delete(AccountInvoiceTaxModel, ids)
}

// GetAccountInvoiceTax gets account.invoice.tax existing record.
func (c *Client) GetAccountInvoiceTax(id int64) (*AccountInvoiceTax, error) {
	aits, err := c.GetAccountInvoiceTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if aits != nil && len(*aits) > 0 {
		return &((*aits)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice.tax not found", id)
}

// GetAccountInvoiceTaxs gets account.invoice.tax existing records.
func (c *Client) GetAccountInvoiceTaxs(ids []int64) (*AccountInvoiceTaxs, error) {
	aits := &AccountInvoiceTaxs{}
	if err := c.Read(AccountInvoiceTaxModel, ids, nil, aits); err != nil {
		return nil, err
	}
	return aits, nil
}

// FindAccountInvoiceTax finds account.invoice.tax record by querying it with criteria.
func (c *Client) FindAccountInvoiceTax(criteria *Criteria) (*AccountInvoiceTax, error) {
	aits := &AccountInvoiceTaxs{}
	if err := c.SearchRead(AccountInvoiceTaxModel, criteria, NewOptions().Limit(1), aits); err != nil {
		return nil, err
	}
	if aits != nil && len(*aits) > 0 {
		return &((*aits)[0]), nil
	}
	return nil, fmt.Errorf("account.invoice.tax was not found with criteria %v", criteria)
}

// FindAccountInvoiceTaxs finds account.invoice.tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceTaxs(criteria *Criteria, options *Options) (*AccountInvoiceTaxs, error) {
	aits := &AccountInvoiceTaxs{}
	if err := c.SearchRead(AccountInvoiceTaxModel, criteria, options, aits); err != nil {
		return nil, err
	}
	return aits, nil
}

// FindAccountInvoiceTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceTaxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceTaxId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.invoice.tax was not found with criteria %v and options %v", criteria, options)
}
