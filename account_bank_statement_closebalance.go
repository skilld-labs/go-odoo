package odoo

import (
	"fmt"
)

// AccountBankStatementClosebalance represents account.bank.statement.closebalance model.
type AccountBankStatementClosebalance struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementClosebalances represents array of account.bank.statement.closebalance model.
type AccountBankStatementClosebalances []AccountBankStatementClosebalance

// AccountBankStatementClosebalanceModel is the odoo model name.
const AccountBankStatementClosebalanceModel = "account.bank.statement.closebalance"

// Many2One convert AccountBankStatementClosebalance to *Many2One.
func (absc *AccountBankStatementClosebalance) Many2One() *Many2One {
	return NewMany2One(absc.Id.Get(), "")
}

// CreateAccountBankStatementClosebalance creates a new account.bank.statement.closebalance model and returns its id.
func (c *Client) CreateAccountBankStatementClosebalance(absc *AccountBankStatementClosebalance) (int64, error) {
	ids, err := c.Create(AccountBankStatementClosebalanceModel, []interface{}{absc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementClosebalance creates a new account.bank.statement.closebalance model and returns its id.
func (c *Client) CreateAccountBankStatementClosebalances(abscs []*AccountBankStatementClosebalance) ([]int64, error) {
	var vv []interface{}
	for _, v := range abscs {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementClosebalanceModel, vv)
}

// UpdateAccountBankStatementClosebalance updates an existing account.bank.statement.closebalance record.
func (c *Client) UpdateAccountBankStatementClosebalance(absc *AccountBankStatementClosebalance) error {
	return c.UpdateAccountBankStatementClosebalances([]int64{absc.Id.Get()}, absc)
}

// UpdateAccountBankStatementClosebalances updates existing account.bank.statement.closebalance records.
// All records (represented by ids) will be updated by absc values.
func (c *Client) UpdateAccountBankStatementClosebalances(ids []int64, absc *AccountBankStatementClosebalance) error {
	return c.Update(AccountBankStatementClosebalanceModel, ids, absc)
}

// DeleteAccountBankStatementClosebalance deletes an existing account.bank.statement.closebalance record.
func (c *Client) DeleteAccountBankStatementClosebalance(id int64) error {
	return c.DeleteAccountBankStatementClosebalances([]int64{id})
}

// DeleteAccountBankStatementClosebalances deletes existing account.bank.statement.closebalance records.
func (c *Client) DeleteAccountBankStatementClosebalances(ids []int64) error {
	return c.Delete(AccountBankStatementClosebalanceModel, ids)
}

// GetAccountBankStatementClosebalance gets account.bank.statement.closebalance existing record.
func (c *Client) GetAccountBankStatementClosebalance(id int64) (*AccountBankStatementClosebalance, error) {
	abscs, err := c.GetAccountBankStatementClosebalances([]int64{id})
	if err != nil {
		return nil, err
	}
	if abscs != nil && len(*abscs) > 0 {
		return &((*abscs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.closebalance not found", id)
}

// GetAccountBankStatementClosebalances gets account.bank.statement.closebalance existing records.
func (c *Client) GetAccountBankStatementClosebalances(ids []int64) (*AccountBankStatementClosebalances, error) {
	abscs := &AccountBankStatementClosebalances{}
	if err := c.Read(AccountBankStatementClosebalanceModel, ids, nil, abscs); err != nil {
		return nil, err
	}
	return abscs, nil
}

// FindAccountBankStatementClosebalance finds account.bank.statement.closebalance record by querying it with criteria.
func (c *Client) FindAccountBankStatementClosebalance(criteria *Criteria) (*AccountBankStatementClosebalance, error) {
	abscs := &AccountBankStatementClosebalances{}
	if err := c.SearchRead(AccountBankStatementClosebalanceModel, criteria, NewOptions().Limit(1), abscs); err != nil {
		return nil, err
	}
	if abscs != nil && len(*abscs) > 0 {
		return &((*abscs)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement.closebalance was not found with criteria %v", criteria)
}

// FindAccountBankStatementClosebalances finds account.bank.statement.closebalance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementClosebalances(criteria *Criteria, options *Options) (*AccountBankStatementClosebalances, error) {
	abscs := &AccountBankStatementClosebalances{}
	if err := c.SearchRead(AccountBankStatementClosebalanceModel, criteria, options, abscs); err != nil {
		return nil, err
	}
	return abscs, nil
}

// FindAccountBankStatementClosebalanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementClosebalanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementClosebalanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementClosebalanceId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementClosebalanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementClosebalanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement.closebalance was not found with criteria %v and options %v", criteria, options)
}
