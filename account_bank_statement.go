package odoo

import (
	"fmt"
)

// AccountBankStatement represents account.bank.statement model.
type AccountBankStatement struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	AllLinesReconciled       *Bool      `xmlrpc:"all_lines_reconciled,omitempty"`
	BalanceEnd               *Float     `xmlrpc:"balance_end,omitempty"`
	BalanceEndReal           *Float     `xmlrpc:"balance_end_real,omitempty"`
	BalanceStart             *Float     `xmlrpc:"balance_start,omitempty"`
	CashboxEndId             *Many2One  `xmlrpc:"cashbox_end_id,omitempty"`
	CashboxStartId           *Many2One  `xmlrpc:"cashbox_start_id,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode              *String    `xmlrpc:"country_code,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty"`
	DateDone                 *Time      `xmlrpc:"date_done,omitempty"`
	Difference               *Float     `xmlrpc:"difference,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsDifferenceZero         *Bool      `xmlrpc:"is_difference_zero,omitempty"`
	IsValidBalanceStart      *Bool      `xmlrpc:"is_valid_balance_start,omitempty"`
	JournalId                *Many2One  `xmlrpc:"journal_id,omitempty"`
	JournalType              *Selection `xmlrpc:"journal_type,omitempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveLineCount            *Int       `xmlrpc:"move_line_count,omitempty"`
	MoveLineIds              *Relation  `xmlrpc:"move_line_ids,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	PreviousStatementId      *Many2One  `xmlrpc:"previous_statement_id,omitempty"`
	Reference                *String    `xmlrpc:"reference,omitempty"`
	SequenceNumber           *Int       `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix           *String    `xmlrpc:"sequence_prefix,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	TotalEntryEncoding       *Float     `xmlrpc:"total_entry_encoding,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatements represents array of account.bank.statement model.
type AccountBankStatements []AccountBankStatement

// AccountBankStatementModel is the odoo model name.
const AccountBankStatementModel = "account.bank.statement"

// Many2One convert AccountBankStatement to *Many2One.
func (abs *AccountBankStatement) Many2One() *Many2One {
	return NewMany2One(abs.Id.Get(), "")
}

// CreateAccountBankStatement creates a new account.bank.statement model and returns its id.
func (c *Client) CreateAccountBankStatement(abs *AccountBankStatement) (int64, error) {
	return c.Create(AccountBankStatementModel, abs)
}

// UpdateAccountBankStatement updates an existing account.bank.statement record.
func (c *Client) UpdateAccountBankStatement(abs *AccountBankStatement) error {
	return c.UpdateAccountBankStatements([]int64{abs.Id.Get()}, abs)
}

// UpdateAccountBankStatements updates existing account.bank.statement records.
// All records (represented by ids) will be updated by abs values.
func (c *Client) UpdateAccountBankStatements(ids []int64, abs *AccountBankStatement) error {
	return c.Update(AccountBankStatementModel, ids, abs)
}

// DeleteAccountBankStatement deletes an existing account.bank.statement record.
func (c *Client) DeleteAccountBankStatement(id int64) error {
	return c.DeleteAccountBankStatements([]int64{id})
}

// DeleteAccountBankStatements deletes existing account.bank.statement records.
func (c *Client) DeleteAccountBankStatements(ids []int64) error {
	return c.Delete(AccountBankStatementModel, ids)
}

// GetAccountBankStatement gets account.bank.statement existing record.
func (c *Client) GetAccountBankStatement(id int64) (*AccountBankStatement, error) {
	abss, err := c.GetAccountBankStatements([]int64{id})
	if err != nil {
		return nil, err
	}
	if abss != nil && len(*abss) > 0 {
		return &((*abss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement not found", id)
}

// GetAccountBankStatements gets account.bank.statement existing records.
func (c *Client) GetAccountBankStatements(ids []int64) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.Read(AccountBankStatementModel, ids, nil, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatement finds account.bank.statement record by querying it with criteria.
func (c *Client) FindAccountBankStatement(criteria *Criteria) (*AccountBankStatement, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, NewOptions().Limit(1), abss); err != nil {
		return nil, err
	}
	if abss != nil && len(*abss) > 0 {
		return &((*abss)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement was not found")
}

// FindAccountBankStatements finds account.bank.statement records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatements(criteria *Criteria, options *Options) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, options, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatementIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement was not found")
}
