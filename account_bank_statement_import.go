package odoo

import (
	"fmt"
)

// AccountBankStatementImport represents account.bank.statement.import model.
type AccountBankStatementImport struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DataFile    *String   `xmlrpc:"data_file,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Filename    *String   `xmlrpc:"filename,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementImports represents array of account.bank.statement.import model.
type AccountBankStatementImports []AccountBankStatementImport

// AccountBankStatementImportModel is the odoo model name.
const AccountBankStatementImportModel = "account.bank.statement.import"

// Many2One convert AccountBankStatementImport to *Many2One.
func (absi *AccountBankStatementImport) Many2One() *Many2One {
	return NewMany2One(absi.Id.Get(), "")
}

// CreateAccountBankStatementImport creates a new account.bank.statement.import model and returns its id.
func (c *Client) CreateAccountBankStatementImport(absi *AccountBankStatementImport) (int64, error) {
	return c.Create(AccountBankStatementImportModel, absi)
}

// UpdateAccountBankStatementImport updates an existing account.bank.statement.import record.
func (c *Client) UpdateAccountBankStatementImport(absi *AccountBankStatementImport) error {
	return c.UpdateAccountBankStatementImports([]int64{absi.Id.Get()}, absi)
}

// UpdateAccountBankStatementImports updates existing account.bank.statement.import records.
// All records (represented by ids) will be updated by absi values.
func (c *Client) UpdateAccountBankStatementImports(ids []int64, absi *AccountBankStatementImport) error {
	return c.Update(AccountBankStatementImportModel, ids, absi)
}

// DeleteAccountBankStatementImport deletes an existing account.bank.statement.import record.
func (c *Client) DeleteAccountBankStatementImport(id int64) error {
	return c.DeleteAccountBankStatementImports([]int64{id})
}

// DeleteAccountBankStatementImports deletes existing account.bank.statement.import records.
func (c *Client) DeleteAccountBankStatementImports(ids []int64) error {
	return c.Delete(AccountBankStatementImportModel, ids)
}

// GetAccountBankStatementImport gets account.bank.statement.import existing record.
func (c *Client) GetAccountBankStatementImport(id int64) (*AccountBankStatementImport, error) {
	absis, err := c.GetAccountBankStatementImports([]int64{id})
	if err != nil {
		return nil, err
	}
	if absis != nil && len(*absis) > 0 {
		return &((*absis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.import not found", id)
}

// GetAccountBankStatementImports gets account.bank.statement.import existing records.
func (c *Client) GetAccountBankStatementImports(ids []int64) (*AccountBankStatementImports, error) {
	absis := &AccountBankStatementImports{}
	if err := c.Read(AccountBankStatementImportModel, ids, nil, absis); err != nil {
		return nil, err
	}
	return absis, nil
}

// FindAccountBankStatementImport finds account.bank.statement.import record by querying it with criteria.
func (c *Client) FindAccountBankStatementImport(criteria *Criteria) (*AccountBankStatementImport, error) {
	absis := &AccountBankStatementImports{}
	if err := c.SearchRead(AccountBankStatementImportModel, criteria, NewOptions().Limit(1), absis); err != nil {
		return nil, err
	}
	if absis != nil && len(*absis) > 0 {
		return &((*absis)[0]), nil
	}
	return nil, fmt.Errorf("no account.bank.statement.import was found with criteria %v", criteria)
}

// FindAccountBankStatementImports finds account.bank.statement.import records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImports(criteria *Criteria, options *Options) (*AccountBankStatementImports, error) {
	absis := &AccountBankStatementImports{}
	if err := c.SearchRead(AccountBankStatementImportModel, criteria, options, absis); err != nil {
		return nil, err
	}
	return absis, nil
}

// FindAccountBankStatementImportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementImportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementImportId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementImportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementImportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.bank.statement.import was found with criteria %v and options %v", criteria, options)
}
