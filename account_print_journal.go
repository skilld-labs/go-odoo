package odoo

import (
	"fmt"
)

// AccountPrintJournal represents account.print.journal model.
type AccountPrintJournal struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AmountCurrency *Bool      `xmlrpc:"amount_currency,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omptempty"`
	DateTo         *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	JournalIds     *Relation  `xmlrpc:"journal_ids,omptempty"`
	SortSelection  *Selection `xmlrpc:"sort_selection,omptempty"`
	TargetMove     *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountPrintJournals represents array of account.print.journal model.
type AccountPrintJournals []AccountPrintJournal

// AccountPrintJournalModel is the odoo model name.
const AccountPrintJournalModel = "account.print.journal"

// Many2One convert AccountPrintJournal to *Many2One.
func (apj *AccountPrintJournal) Many2One() *Many2One {
	return NewMany2One(apj.Id.Get(), "")
}

// CreateAccountPrintJournal creates a new account.print.journal model and returns its id.
func (c *Client) CreateAccountPrintJournal(apj *AccountPrintJournal) (int64, error) {
	return c.Create(AccountPrintJournalModel, apj)
}

// UpdateAccountPrintJournal updates an existing account.print.journal record.
func (c *Client) UpdateAccountPrintJournal(apj *AccountPrintJournal) error {
	return c.UpdateAccountPrintJournals([]int64{apj.Id.Get()}, apj)
}

// UpdateAccountPrintJournals updates existing account.print.journal records.
// All records (represented by ids) will be updated by apj values.
func (c *Client) UpdateAccountPrintJournals(ids []int64, apj *AccountPrintJournal) error {
	return c.Update(AccountPrintJournalModel, ids, apj)
}

// DeleteAccountPrintJournal deletes an existing account.print.journal record.
func (c *Client) DeleteAccountPrintJournal(id int64) error {
	return c.DeleteAccountPrintJournals([]int64{id})
}

// DeleteAccountPrintJournals deletes existing account.print.journal records.
func (c *Client) DeleteAccountPrintJournals(ids []int64) error {
	return c.Delete(AccountPrintJournalModel, ids)
}

// GetAccountPrintJournal gets account.print.journal existing record.
func (c *Client) GetAccountPrintJournal(id int64) (*AccountPrintJournal, error) {
	apjs, err := c.GetAccountPrintJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if apjs != nil && len(*apjs) > 0 {
		return &((*apjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.print.journal not found", id)
}

// GetAccountPrintJournals gets account.print.journal existing records.
func (c *Client) GetAccountPrintJournals(ids []int64) (*AccountPrintJournals, error) {
	apjs := &AccountPrintJournals{}
	if err := c.Read(AccountPrintJournalModel, ids, nil, apjs); err != nil {
		return nil, err
	}
	return apjs, nil
}

// FindAccountPrintJournal finds account.print.journal record by querying it with criteria.
func (c *Client) FindAccountPrintJournal(criteria *Criteria) (*AccountPrintJournal, error) {
	apjs := &AccountPrintJournals{}
	if err := c.SearchRead(AccountPrintJournalModel, criteria, NewOptions().Limit(1), apjs); err != nil {
		return nil, err
	}
	if apjs != nil && len(*apjs) > 0 {
		return &((*apjs)[0]), nil
	}
	return nil, fmt.Errorf("account.print.journal was not found with criteria %v", criteria)
}

// FindAccountPrintJournals finds account.print.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPrintJournals(criteria *Criteria, options *Options) (*AccountPrintJournals, error) {
	apjs := &AccountPrintJournals{}
	if err := c.SearchRead(AccountPrintJournalModel, criteria, options, apjs); err != nil {
		return nil, err
	}
	return apjs, nil
}

// FindAccountPrintJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPrintJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPrintJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPrintJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountPrintJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPrintJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.print.journal was not found with criteria %v and options %v", criteria, options)
}
