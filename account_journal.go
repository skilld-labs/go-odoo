package odoo

// AccountJournal represents account.journal model.
type AccountJournal struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	AccountControlIds        *Relation  `xmlrpc:"account_control_ids,omitempty"`
	AccountSetupBankDataDone *Bool      `xmlrpc:"account_setup_bank_data_done,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	AtLeastOneInbound        *Bool      `xmlrpc:"at_least_one_inbound,omitempty"`
	AtLeastOneOutbound       *Bool      `xmlrpc:"at_least_one_outbound,omitempty"`
	BankAccNumber            *String    `xmlrpc:"bank_acc_number,omitempty"`
	BankAccountId            *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	BankId                   *Many2One  `xmlrpc:"bank_id,omitempty"`
	BankStatementsSource     *Selection `xmlrpc:"bank_statements_source,omitempty"`
	BelongsToCompany         *Bool      `xmlrpc:"belongs_to_company,omitempty"`
	Code                     *String    `xmlrpc:"code,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCreditAccountId   *Many2One  `xmlrpc:"default_credit_account_id,omitempty"`
	DefaultDebitAccountId    *Many2One  `xmlrpc:"default_debit_account_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	GroupInvoiceLines        *Bool      `xmlrpc:"group_invoice_lines,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	InboundPaymentMethodIds  *Relation  `xmlrpc:"inbound_payment_method_ids,omitempty"`
	KanbanDashboard          *String    `xmlrpc:"kanban_dashboard,omitempty"`
	KanbanDashboardGraph     *String    `xmlrpc:"kanban_dashboard_graph,omitempty"`
	LossAccountId            *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	OutboundPaymentMethodIds *Relation  `xmlrpc:"outbound_payment_method_ids,omitempty"`
	ProfitAccountId          *Many2One  `xmlrpc:"profit_account_id,omitempty"`
	RefundSequence           *Bool      `xmlrpc:"refund_sequence,omitempty"`
	RefundSequenceId         *Many2One  `xmlrpc:"refund_sequence_id,omitempty"`
	RefundSequenceNumberNext *Int       `xmlrpc:"refund_sequence_number_next,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	SequenceId               *Many2One  `xmlrpc:"sequence_id,omitempty"`
	SequenceNumberNext       *Int       `xmlrpc:"sequence_number_next,omitempty"`
	ShowOnDashboard          *Bool      `xmlrpc:"show_on_dashboard,omitempty"`
	Type                     *Selection `xmlrpc:"type,omitempty"`
	TypeControlIds           *Relation  `xmlrpc:"type_control_ids,omitempty"`
	UpdatePosted             *Bool      `xmlrpc:"update_posted,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateAccountJournals([]*AccountJournal{aj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountJournals creates a new account.journal model and returns its id.
func (c *Client) CreateAccountJournals(ajs []*AccountJournal) ([]int64, error) {
	var vv []interface{}
	for _, v := range ajs {
		vv = append(vv, v)
	}
	return c.Create(AccountJournalModel, vv, nil)
}

// UpdateAccountJournal updates an existing account.journal record.
func (c *Client) UpdateAccountJournal(aj *AccountJournal) error {
	return c.UpdateAccountJournals([]int64{aj.Id.Get()}, aj)
}

// UpdateAccountJournals updates existing account.journal records.
// All records (represented by ids) will be updated by aj values.
func (c *Client) UpdateAccountJournals(ids []int64, aj *AccountJournal) error {
	return c.Update(AccountJournalModel, ids, aj, nil)
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
	return &((*ajs)[0]), nil
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
	return &((*ajs)[0]), nil
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
	return c.Search(AccountJournalModel, criteria, options)
}

// FindAccountJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
