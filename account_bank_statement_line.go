package odoo

import (
	"fmt"
)

// AccountBankStatementLine represents account.bank.statement.line model.
type AccountBankStatementLine struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId         *Many2One  `xmlrpc:"account_id,omptempty"`
	Amount            *Float     `xmlrpc:"amount,omptempty"`
	AmountCurrency    *Float     `xmlrpc:"amount_currency,omptempty"`
	BankAccountId     *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date              *Time      `xmlrpc:"date,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	JournalCurrencyId *Many2One  `xmlrpc:"journal_currency_id,omptempty"`
	JournalEntryIds   *Relation  `xmlrpc:"journal_entry_ids,omptempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omptempty"`
	MoveName          *String    `xmlrpc:"move_name,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Note              *String    `xmlrpc:"note,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerName       *String    `xmlrpc:"partner_name,omptempty"`
	Ref               *String    `xmlrpc:"ref,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	StatementId       *Many2One  `xmlrpc:"statement_id,omptempty"`
	UniqueImportId    *String    `xmlrpc:"unique_import_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementLines represents array of account.bank.statement.line model.
type AccountBankStatementLines []AccountBankStatementLine

// AccountBankStatementLineModel is the odoo model name.
const AccountBankStatementLineModel = "account.bank.statement.line"

// Many2One convert AccountBankStatementLine to *Many2One.
func (absl *AccountBankStatementLine) Many2One() *Many2One {
	return NewMany2One(absl.Id.Get(), "")
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLine(absl *AccountBankStatementLine) (int64, error) {
	ids, err := c.CreateAccountBankStatementLines([]*AccountBankStatementLine{absl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLines(absls []*AccountBankStatementLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range absls {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementLineModel, vv)
}

// UpdateAccountBankStatementLine updates an existing account.bank.statement.line record.
func (c *Client) UpdateAccountBankStatementLine(absl *AccountBankStatementLine) error {
	return c.UpdateAccountBankStatementLines([]int64{absl.Id.Get()}, absl)
}

// UpdateAccountBankStatementLines updates existing account.bank.statement.line records.
// All records (represented by ids) will be updated by absl values.
func (c *Client) UpdateAccountBankStatementLines(ids []int64, absl *AccountBankStatementLine) error {
	return c.Update(AccountBankStatementLineModel, ids, absl)
}

// DeleteAccountBankStatementLine deletes an existing account.bank.statement.line record.
func (c *Client) DeleteAccountBankStatementLine(id int64) error {
	return c.DeleteAccountBankStatementLines([]int64{id})
}

// DeleteAccountBankStatementLines deletes existing account.bank.statement.line records.
func (c *Client) DeleteAccountBankStatementLines(ids []int64) error {
	return c.Delete(AccountBankStatementLineModel, ids)
}

// GetAccountBankStatementLine gets account.bank.statement.line existing record.
func (c *Client) GetAccountBankStatementLine(id int64) (*AccountBankStatementLine, error) {
	absls, err := c.GetAccountBankStatementLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if absls != nil && len(*absls) > 0 {
		return &((*absls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.line not found", id)
}

// GetAccountBankStatementLines gets account.bank.statement.line existing records.
func (c *Client) GetAccountBankStatementLines(ids []int64) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.Read(AccountBankStatementLineModel, ids, nil, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLine finds account.bank.statement.line record by querying it with criteria.
func (c *Client) FindAccountBankStatementLine(criteria *Criteria) (*AccountBankStatementLine, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, NewOptions().Limit(1), absls); err != nil {
		return nil, err
	}
	if absls != nil && len(*absls) > 0 {
		return &((*absls)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement.line was not found with criteria %v", criteria)
}

// FindAccountBankStatementLines finds account.bank.statement.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLines(criteria *Criteria, options *Options) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, options, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement.line was not found with criteria %v and options %v", criteria, options)
}
