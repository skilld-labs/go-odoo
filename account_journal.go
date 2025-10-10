package odoo

// AccountJournal represents account.journal model.
type AccountJournal struct {
	AccessToken                  *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omitempty"`
	AccountControlIds            *Relation  `xmlrpc:"account_control_ids,omitempty"`
	AccountingDate               *Time      `xmlrpc:"accounting_date,omitempty"`
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId      *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AliasDefaults                *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId                *Many2One  `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail                   *String    `xmlrpc:"alias_email,omitempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omitempty"`
	AutocheckOnPost              *Bool      `xmlrpc:"autocheck_on_post,omitempty"`
	AvailablePaymentMethodIds    *Relation  `xmlrpc:"available_payment_method_ids,omitempty"`
	BankAccNumber                *String    `xmlrpc:"bank_acc_number,omitempty"`
	BankAccountId                *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	BankId                       *Many2One  `xmlrpc:"bank_id,omitempty"`
	BankStatementsSource         *Selection `xmlrpc:"bank_statements_source,omitempty"`
	Code                         *String    `xmlrpc:"code,omitempty"`
	Color                        *Int       `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyPartnerId             *Many2One  `xmlrpc:"company_partner_id,omitempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrentStatementBalance      *Float     `xmlrpc:"current_statement_balance,omitempty"`
	DefaultAccountId             *Many2One  `xmlrpc:"default_account_id,omitempty"`
	DefaultAccountType           *String    `xmlrpc:"default_account_type,omitempty"`
	DisplayAliasFields           *Bool      `xmlrpc:"display_alias_fields,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	EntriesCount                 *Int       `xmlrpc:"entries_count,omitempty"`
	HasEntries                   *Bool      `xmlrpc:"has_entries,omitempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omitempty"`
	HasPostedEntries             *Bool      `xmlrpc:"has_posted_entries,omitempty"`
	HasSequenceHoles             *Bool      `xmlrpc:"has_sequence_holes,omitempty"`
	HasStatementLines            *Bool      `xmlrpc:"has_statement_lines,omitempty"`
	HasUnhashedEntries           *Bool      `xmlrpc:"has_unhashed_entries,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	InboundPaymentMethodLineIds  *Relation  `xmlrpc:"inbound_payment_method_line_ids,omitempty"`
	InvoiceReferenceModel        *Selection `xmlrpc:"invoice_reference_model,omitempty"`
	InvoiceReferenceType         *Selection `xmlrpc:"invoice_reference_type,omitempty"`
	JournalGroupIds              *Relation  `xmlrpc:"journal_group_ids,omitempty"`
	JsonActivityData             *String    `xmlrpc:"json_activity_data,omitempty"`
	KanbanDashboard              *String    `xmlrpc:"kanban_dashboard,omitempty"`
	KanbanDashboardGraph         *String    `xmlrpc:"kanban_dashboard_graph,omitempty"`
	LastStatementId              *Many2One  `xmlrpc:"last_statement_id,omitempty"`
	LossAccountId                *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	OutboundPaymentMethodLineIds *Relation  `xmlrpc:"outbound_payment_method_line_ids,omitempty"`
	PaymentSequence              *Bool      `xmlrpc:"payment_sequence,omitempty"`
	ProfitAccountId              *Many2One  `xmlrpc:"profit_account_id,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	RefundSequence               *Bool      `xmlrpc:"refund_sequence,omitempty"`
	RestrictModeHashTable        *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	SelectedPaymentMethodCodes   *String    `xmlrpc:"selected_payment_method_codes,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	SequenceOverrideRegex        *String    `xmlrpc:"sequence_override_regex,omitempty"`
	ShowOnDashboard              *Bool      `xmlrpc:"show_on_dashboard,omitempty"`
	SuspenseAccountId            *Many2One  `xmlrpc:"suspense_account_id,omitempty"`
	Type                         *Selection `xmlrpc:"type,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
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

// CreateAccountJournal creates a new account.journal model and returns its id.
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
