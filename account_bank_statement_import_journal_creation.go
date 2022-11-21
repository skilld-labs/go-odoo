package odoo

import (
	"fmt"
)

// AccountBankStatementImportJournalCreation represents account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreation struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountControlIds        *Relation  `xmlrpc:"account_control_ids,omptempty"`
	AccountSetupBankDataDone *Bool      `xmlrpc:"account_setup_bank_data_done,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	AtLeastOneInbound        *Bool      `xmlrpc:"at_least_one_inbound,omptempty"`
	AtLeastOneOutbound       *Bool      `xmlrpc:"at_least_one_outbound,omptempty"`
	BankAccNumber            *String    `xmlrpc:"bank_acc_number,omptempty"`
	BankAccountId            *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	BankId                   *Many2One  `xmlrpc:"bank_id,omptempty"`
	BankStatementsSource     *Selection `xmlrpc:"bank_statements_source,omptempty"`
	BelongsToCompany         *Bool      `xmlrpc:"belongs_to_company,omptempty"`
	Code                     *String    `xmlrpc:"code,omptempty"`
	Color                    *Int       `xmlrpc:"color,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCreditAccountId   *Many2One  `xmlrpc:"default_credit_account_id,omptempty"`
	DefaultDebitAccountId    *Many2One  `xmlrpc:"default_debit_account_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	GroupInvoiceLines        *Bool      `xmlrpc:"group_invoice_lines,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	InboundPaymentMethodIds  *Relation  `xmlrpc:"inbound_payment_method_ids,omptempty"`
	JournalId                *Many2One  `xmlrpc:"journal_id,omptempty"`
	KanbanDashboard          *String    `xmlrpc:"kanban_dashboard,omptempty"`
	KanbanDashboardGraph     *String    `xmlrpc:"kanban_dashboard_graph,omptempty"`
	LossAccountId            *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OutboundPaymentMethodIds *Relation  `xmlrpc:"outbound_payment_method_ids,omptempty"`
	ProfitAccountId          *Many2One  `xmlrpc:"profit_account_id,omptempty"`
	RefundSequence           *Bool      `xmlrpc:"refund_sequence,omptempty"`
	RefundSequenceId         *Many2One  `xmlrpc:"refund_sequence_id,omptempty"`
	RefundSequenceNumberNext *Int       `xmlrpc:"refund_sequence_number_next,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	SequenceId               *Many2One  `xmlrpc:"sequence_id,omptempty"`
	SequenceNumberNext       *Int       `xmlrpc:"sequence_number_next,omptempty"`
	ShowOnDashboard          *Bool      `xmlrpc:"show_on_dashboard,omptempty"`
	Type                     *Selection `xmlrpc:"type,omptempty"`
	TypeControlIds           *Relation  `xmlrpc:"type_control_ids,omptempty"`
	UpdatePosted             *Bool      `xmlrpc:"update_posted,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementImportJournalCreations represents array of account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreations []AccountBankStatementImportJournalCreation

// AccountBankStatementImportJournalCreationModel is the odoo model name.
const AccountBankStatementImportJournalCreationModel = "account.bank.statement.import.journal.creation"

// Many2One convert AccountBankStatementImportJournalCreation to *Many2One.
func (absijc *AccountBankStatementImportJournalCreation) Many2One() *Many2One {
	return NewMany2One(absijc.Id.Get(), "")
}

// CreateAccountBankStatementImportJournalCreation creates a new account.bank.statement.import.journal.creation model and returns its id.
func (c *Client) CreateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) (int64, error) {
	return c.Create(AccountBankStatementImportJournalCreationModel, absijc)
}

// UpdateAccountBankStatementImportJournalCreation updates an existing account.bank.statement.import.journal.creation record.
func (c *Client) UpdateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) error {
	return c.UpdateAccountBankStatementImportJournalCreations([]int64{absijc.Id.Get()}, absijc)
}

// UpdateAccountBankStatementImportJournalCreations updates existing account.bank.statement.import.journal.creation records.
// All records (represented by ids) will be updated by absijc values.
func (c *Client) UpdateAccountBankStatementImportJournalCreations(ids []int64, absijc *AccountBankStatementImportJournalCreation) error {
	return c.Update(AccountBankStatementImportJournalCreationModel, ids, absijc)
}

// DeleteAccountBankStatementImportJournalCreation deletes an existing account.bank.statement.import.journal.creation record.
func (c *Client) DeleteAccountBankStatementImportJournalCreation(id int64) error {
	return c.DeleteAccountBankStatementImportJournalCreations([]int64{id})
}

// DeleteAccountBankStatementImportJournalCreations deletes existing account.bank.statement.import.journal.creation records.
func (c *Client) DeleteAccountBankStatementImportJournalCreations(ids []int64) error {
	return c.Delete(AccountBankStatementImportJournalCreationModel, ids)
}

// GetAccountBankStatementImportJournalCreation gets account.bank.statement.import.journal.creation existing record.
func (c *Client) GetAccountBankStatementImportJournalCreation(id int64) (*AccountBankStatementImportJournalCreation, error) {
	absijcs, err := c.GetAccountBankStatementImportJournalCreations([]int64{id})
	if err != nil {
		return nil, err
	}
	if absijcs != nil && len(*absijcs) > 0 {
		return &((*absijcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.import.journal.creation not found", id)
}

// GetAccountBankStatementImportJournalCreations gets account.bank.statement.import.journal.creation existing records.
func (c *Client) GetAccountBankStatementImportJournalCreations(ids []int64) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.Read(AccountBankStatementImportJournalCreationModel, ids, nil, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreation finds account.bank.statement.import.journal.creation record by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreation(criteria *Criteria) (*AccountBankStatementImportJournalCreation, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, NewOptions().Limit(1), absijcs); err != nil {
		return nil, err
	}
	if absijcs != nil && len(*absijcs) > 0 {
		return &((*absijcs)[0]), nil
	}
	return nil, fmt.Errorf("no account.bank.statement.import.journal.creation was found with criteria %v", criteria)
}

// FindAccountBankStatementImportJournalCreations finds account.bank.statement.import.journal.creation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreations(criteria *Criteria, options *Options) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, options, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementImportJournalCreationId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.bank.statement.import.journal.creation was found with criteria %v and options %v", criteria, options)
}
