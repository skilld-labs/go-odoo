package odoo

import (
	"fmt"
)

// AccountReconcileModel represents account.reconcile.model model.
type AccountReconcileModel struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId               *Many2One  `xmlrpc:"account_id,omptempty"`
	Amount                  *Float     `xmlrpc:"amount,omptempty"`
	AmountType              *Selection `xmlrpc:"amount_type,omptempty"`
	AnalyticAccountId       *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CompanyId               *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	HasSecondLine           *Bool      `xmlrpc:"has_second_line,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	JournalId               *Many2One  `xmlrpc:"journal_id,omptempty"`
	Label                   *String    `xmlrpc:"label,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	SecondAccountId         *Many2One  `xmlrpc:"second_account_id,omptempty"`
	SecondAmount            *Float     `xmlrpc:"second_amount,omptempty"`
	SecondAmountType        *Selection `xmlrpc:"second_amount_type,omptempty"`
	SecondAnalyticAccountId *Many2One  `xmlrpc:"second_analytic_account_id,omptempty"`
	SecondJournalId         *Many2One  `xmlrpc:"second_journal_id,omptempty"`
	SecondLabel             *String    `xmlrpc:"second_label,omptempty"`
	SecondTaxId             *Many2One  `xmlrpc:"second_tax_id,omptempty"`
	Sequence                *Int       `xmlrpc:"sequence,omptempty"`
	TaxId                   *Many2One  `xmlrpc:"tax_id,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReconcileModels represents array of account.reconcile.model model.
type AccountReconcileModels []AccountReconcileModel

// AccountReconcileModelModel is the odoo model name.
const AccountReconcileModelModel = "account.reconcile.model"

// Many2One convert AccountReconcileModel to *Many2One.
func (arm *AccountReconcileModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModel(arm *AccountReconcileModel) (int64, error) {
	return c.Create(AccountReconcileModelModel, arm)
}

// UpdateAccountReconcileModel updates an existing account.reconcile.model record.
func (c *Client) UpdateAccountReconcileModel(arm *AccountReconcileModel) error {
	return c.UpdateAccountReconcileModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReconcileModels updates existing account.reconcile.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReconcileModels(ids []int64, arm *AccountReconcileModel) error {
	return c.Update(AccountReconcileModelModel, ids, arm)
}

// DeleteAccountReconcileModel deletes an existing account.reconcile.model record.
func (c *Client) DeleteAccountReconcileModel(id int64) error {
	return c.DeleteAccountReconcileModels([]int64{id})
}

// DeleteAccountReconcileModels deletes existing account.reconcile.model records.
func (c *Client) DeleteAccountReconcileModels(ids []int64) error {
	return c.Delete(AccountReconcileModelModel, ids)
}

// GetAccountReconcileModel gets account.reconcile.model existing record.
func (c *Client) GetAccountReconcileModel(id int64) (*AccountReconcileModel, error) {
	arms, err := c.GetAccountReconcileModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model not found", id)
}

// GetAccountReconcileModels gets account.reconcile.model existing records.
func (c *Client) GetAccountReconcileModels(ids []int64) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.Read(AccountReconcileModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModel finds account.reconcile.model record by querying it with criteria.
func (c *Client) FindAccountReconcileModel(criteria *Criteria) (*AccountReconcileModel, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model was not found with criteria %v", criteria)
}

// FindAccountReconcileModels finds account.reconcile.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModels(criteria *Criteria, options *Options) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model was not found with criteria %v and options %v", criteria, options)
}
