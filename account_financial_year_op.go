package odoo

import (
	"fmt"
)

// AccountFinancialYearOp represents account.financial.year.op model.
type AccountFinancialYearOp struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	AccountSetupFyDataDone *Bool      `xmlrpc:"account_setup_fy_data_done,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	FiscalyearLastDay      *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth    *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	OpeningDate            *Time      `xmlrpc:"opening_date,omitempty"`
	OpeningMovePosted      *Bool      `xmlrpc:"opening_move_posted,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFinancialYearOps represents array of account.financial.year.op model.
type AccountFinancialYearOps []AccountFinancialYearOp

// AccountFinancialYearOpModel is the odoo model name.
const AccountFinancialYearOpModel = "account.financial.year.op"

// Many2One convert AccountFinancialYearOp to *Many2One.
func (afyo *AccountFinancialYearOp) Many2One() *Many2One {
	return NewMany2One(afyo.Id.Get(), "")
}

// CreateAccountFinancialYearOp creates a new account.financial.year.op model and returns its id.
func (c *Client) CreateAccountFinancialYearOp(afyo *AccountFinancialYearOp) (int64, error) {
	return c.Create(AccountFinancialYearOpModel, afyo)
}

// UpdateAccountFinancialYearOp updates an existing account.financial.year.op record.
func (c *Client) UpdateAccountFinancialYearOp(afyo *AccountFinancialYearOp) error {
	return c.UpdateAccountFinancialYearOps([]int64{afyo.Id.Get()}, afyo)
}

// UpdateAccountFinancialYearOps updates existing account.financial.year.op records.
// All records (represented by ids) will be updated by afyo values.
func (c *Client) UpdateAccountFinancialYearOps(ids []int64, afyo *AccountFinancialYearOp) error {
	return c.Update(AccountFinancialYearOpModel, ids, afyo)
}

// DeleteAccountFinancialYearOp deletes an existing account.financial.year.op record.
func (c *Client) DeleteAccountFinancialYearOp(id int64) error {
	return c.DeleteAccountFinancialYearOps([]int64{id})
}

// DeleteAccountFinancialYearOps deletes existing account.financial.year.op records.
func (c *Client) DeleteAccountFinancialYearOps(ids []int64) error {
	return c.Delete(AccountFinancialYearOpModel, ids)
}

// GetAccountFinancialYearOp gets account.financial.year.op existing record.
func (c *Client) GetAccountFinancialYearOp(id int64) (*AccountFinancialYearOp, error) {
	afyos, err := c.GetAccountFinancialYearOps([]int64{id})
	if err != nil {
		return nil, err
	}
	if afyos != nil && len(*afyos) > 0 {
		return &((*afyos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.financial.year.op not found", id)
}

// GetAccountFinancialYearOps gets account.financial.year.op existing records.
func (c *Client) GetAccountFinancialYearOps(ids []int64) (*AccountFinancialYearOps, error) {
	afyos := &AccountFinancialYearOps{}
	if err := c.Read(AccountFinancialYearOpModel, ids, nil, afyos); err != nil {
		return nil, err
	}
	return afyos, nil
}

// FindAccountFinancialYearOp finds account.financial.year.op record by querying it with criteria.
func (c *Client) FindAccountFinancialYearOp(criteria *Criteria) (*AccountFinancialYearOp, error) {
	afyos := &AccountFinancialYearOps{}
	if err := c.SearchRead(AccountFinancialYearOpModel, criteria, NewOptions().Limit(1), afyos); err != nil {
		return nil, err
	}
	if afyos != nil && len(*afyos) > 0 {
		return &((*afyos)[0]), nil
	}
	return nil, fmt.Errorf("no account.financial.year.op was found with criteria %v", criteria)
}

// FindAccountFinancialYearOps finds account.financial.year.op records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialYearOps(criteria *Criteria, options *Options) (*AccountFinancialYearOps, error) {
	afyos := &AccountFinancialYearOps{}
	if err := c.SearchRead(AccountFinancialYearOpModel, criteria, options, afyos); err != nil {
		return nil, err
	}
	return afyos, nil
}

// FindAccountFinancialYearOpIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialYearOpIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFinancialYearOpModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFinancialYearOpId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialYearOpId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialYearOpModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.financial.year.op was found with criteria %v and options %v", criteria, options)
}
