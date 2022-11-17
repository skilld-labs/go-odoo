package odoo

import (
	"fmt"
)

// AccountBankStatementCashbox represents account.bank.statement.cashbox model.
type AccountBankStatementCashbox struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CashboxLinesIds *Relation `xmlrpc:"cashbox_lines_ids,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatementCashboxs represents array of account.bank.statement.cashbox model.
type AccountBankStatementCashboxs []AccountBankStatementCashbox

// AccountBankStatementCashboxModel is the odoo model name.
const AccountBankStatementCashboxModel = "account.bank.statement.cashbox"

// Many2One convert AccountBankStatementCashbox to *Many2One.
func (absc *AccountBankStatementCashbox) Many2One() *Many2One {
	return NewMany2One(absc.Id.Get(), "")
}

// CreateAccountBankStatementCashbox creates a new account.bank.statement.cashbox model and returns its id.
func (c *Client) CreateAccountBankStatementCashbox(absc *AccountBankStatementCashbox) (int64, error) {
	return c.Create(AccountBankStatementCashboxModel, absc)
}

// UpdateAccountBankStatementCashbox updates an existing account.bank.statement.cashbox record.
func (c *Client) UpdateAccountBankStatementCashbox(absc *AccountBankStatementCashbox) error {
	return c.UpdateAccountBankStatementCashboxs([]int64{absc.Id.Get()}, absc)
}

// UpdateAccountBankStatementCashboxs updates existing account.bank.statement.cashbox records.
// All records (represented by ids) will be updated by absc values.
func (c *Client) UpdateAccountBankStatementCashboxs(ids []int64, absc *AccountBankStatementCashbox) error {
	return c.Update(AccountBankStatementCashboxModel, ids, absc)
}

// DeleteAccountBankStatementCashbox deletes an existing account.bank.statement.cashbox record.
func (c *Client) DeleteAccountBankStatementCashbox(id int64) error {
	return c.DeleteAccountBankStatementCashboxs([]int64{id})
}

// DeleteAccountBankStatementCashboxs deletes existing account.bank.statement.cashbox records.
func (c *Client) DeleteAccountBankStatementCashboxs(ids []int64) error {
	return c.Delete(AccountBankStatementCashboxModel, ids)
}

// GetAccountBankStatementCashbox gets account.bank.statement.cashbox existing record.
func (c *Client) GetAccountBankStatementCashbox(id int64) (*AccountBankStatementCashbox, error) {
	abscs, err := c.GetAccountBankStatementCashboxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if abscs != nil && len(*abscs) > 0 {
		return &((*abscs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.cashbox not found", id)
}

// GetAccountBankStatementCashboxs gets account.bank.statement.cashbox existing records.
func (c *Client) GetAccountBankStatementCashboxs(ids []int64) (*AccountBankStatementCashboxs, error) {
	abscs := &AccountBankStatementCashboxs{}
	if err := c.Read(AccountBankStatementCashboxModel, ids, nil, abscs); err != nil {
		return nil, err
	}
	return abscs, nil
}

// FindAccountBankStatementCashbox finds account.bank.statement.cashbox record by querying it with criteria.
func (c *Client) FindAccountBankStatementCashbox(criteria *Criteria) (*AccountBankStatementCashbox, error) {
	abscs := &AccountBankStatementCashboxs{}
	if err := c.SearchRead(AccountBankStatementCashboxModel, criteria, NewOptions().Limit(1), abscs); err != nil {
		return nil, err
	}
	if abscs != nil && len(*abscs) > 0 {
		return &((*abscs)[0]), nil
	}
	return nil, fmt.Errorf("no account.bank.statement.cashbox was found with criteria %v", criteria)
}

// FindAccountBankStatementCashboxs finds account.bank.statement.cashbox records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementCashboxs(criteria *Criteria, options *Options) (*AccountBankStatementCashboxs, error) {
	abscs := &AccountBankStatementCashboxs{}
	if err := c.SearchRead(AccountBankStatementCashboxModel, criteria, options, abscs); err != nil {
		return nil, err
	}
	return abscs, nil
}

// FindAccountBankStatementCashboxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementCashboxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementCashboxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementCashboxId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementCashboxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementCashboxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.bank.statement.cashbox was found with criteria %v and options %v", criteria, options)
}
