package odoo

import (
	"fmt"
)

// AccountJournal represents account.journal model.
type AccountJournal struct {
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

// AccountJournals represents array of account.journal model.
type AccountJournals []AccountJournal

// AccountJournalModel is the odoo model name.
const AccountJournalModel = "account.journal"

// Many2One convert AccountJournal to *Many2One.
func (aj *AccountJournal) Many2One() *Many2One {
	return NewMany2One(aj.Id.Get(), "")
}

// CreateAccountJournal creates a new account.journal model and returns its id.
func (c *Client) CreateAccountJournal(aj *AccountJournal) (int64, error) {
	return c.Create(AccountJournalModel, aj)
}

// UpdateAccountJournal updates an existing account.journal record.
func (c *Client) UpdateAccountJournal(aj *AccountJournal) error {
	return c.UpdateAccountJournals([]int64{aj.Id.Get()}, aj)
}

// UpdateAccountJournals updates existing account.journal records.
// All records (represented by ids) will be updated by aj values.
func (c *Client) UpdateAccountJournals(ids []int64, aj *AccountJournal) error {
	return c.Update(AccountJournalModel, ids, aj)
}

// DeleteAccountJournal deletes an existing account.journal record.
func (c *Client) DeleteAccountJournal(id int64) error {
	return c.DeleteAccountJournals([]int64{id})
}

// DeleteAccountJournals deletes existing account.journal records.
func (c *Client) DeleteAccountJournals(ids []int64) error {
	return c.Delete(AccountJournalModel, ids)
}

// GetAccountJournal gets account.journal existing record.
func (c *Client) GetAccountJournal(id int64) (*AccountJournal, error) {
	ajs, err := c.GetAccountJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if ajs != nil && len(*ajs) > 0 {
		return &((*ajs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.journal was not found", id)
}

// GetAccountJournals gets account.journal existing records.
func (c *Client) GetAccountJournals(ids []int64) (*AccountJournals, error) {
	ajs := &AccountJournals{}
	if err := c.Read(AccountJournalModel, ids, nil, ajs); err != nil {
		return nil, err
	}
	return ajs, nil
}

// FindAccountJournal finds account.journal record by querying it with criteria.
func (c *Client) FindAccountJournal(criteria *Criteria) (*AccountJournal, error) {
	ajs := &AccountJournals{}
	if err := c.SearchRead(AccountJournalModel, criteria, NewOptions().Limit(1), ajs); err != nil {
		return nil, err
	}
	if ajs != nil && len(*ajs) > 0 {
		return &((*ajs)[0]), nil
	}
	return nil, fmt.Errorf("account.journal was not found")
}

// FindAccountJournals finds account.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournals(criteria *Criteria, options *Options) (*AccountJournals, error) {
	ajs := &AccountJournals{}
	if err := c.SearchRead(AccountJournalModel, criteria, options, ajs); err != nil {
		return nil, err
	}
	return ajs, nil
}

// FindAccountJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalId(criteria *Criteria) (int64, error) {
	ids, err := c.Search(AccountJournalModel, criteria, NewOptions().Limit(1))
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.journal was not found")
}
