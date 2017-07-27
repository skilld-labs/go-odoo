package types

import (
	"time"
)

type AccountBankStatement struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AllLinesReconciled       bool      `xmlrpc:"all_lines_reconciled"`
	BalanceEnd               float64   `xmlrpc:"balance_end"`
	BalanceEndReal           float64   `xmlrpc:"balance_end_real"`
	BalanceStart             float64   `xmlrpc:"balance_start"`
	CashboxEndId             Many2One  `xmlrpc:"cashbox_end_id"`
	CashboxStartId           Many2One  `xmlrpc:"cashbox_start_id"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	Date                     time.Time `xmlrpc:"date"`
	DateDone                 time.Time `xmlrpc:"date_done"`
	Difference               float64   `xmlrpc:"difference"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
	IsDifferenceZero         bool      `xmlrpc:"is_difference_zero"`
	JournalId                Many2One  `xmlrpc:"journal_id"`
	JournalType              string    `xmlrpc:"journal_type"`
	LineIds                  []int64   `xmlrpc:"line_ids"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	MoveLineIds              []int64   `xmlrpc:"move_line_ids"`
	Name                     string    `xmlrpc:"name"`
	Reference                string    `xmlrpc:"reference"`
	State                    string    `xmlrpc:"state"`
	TotalEntryEncoding       float64   `xmlrpc:"total_entry_encoding"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type AccountBankStatementNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AllLinesReconciled       bool        `xmlrpc:"all_lines_reconciled"`
	BalanceEnd               interface{} `xmlrpc:"balance_end"`
	BalanceEndReal           interface{} `xmlrpc:"balance_end_real"`
	BalanceStart             interface{} `xmlrpc:"balance_start"`
	CashboxEndId             interface{} `xmlrpc:"cashbox_end_id"`
	CashboxStartId           interface{} `xmlrpc:"cashbox_start_id"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	Date                     interface{} `xmlrpc:"date"`
	DateDone                 interface{} `xmlrpc:"date_done"`
	Difference               interface{} `xmlrpc:"difference"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
	IsDifferenceZero         bool        `xmlrpc:"is_difference_zero"`
	JournalId                interface{} `xmlrpc:"journal_id"`
	JournalType              interface{} `xmlrpc:"journal_type"`
	LineIds                  interface{} `xmlrpc:"line_ids"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	MoveLineIds              interface{} `xmlrpc:"move_line_ids"`
	Name                     interface{} `xmlrpc:"name"`
	Reference                interface{} `xmlrpc:"reference"`
	State                    interface{} `xmlrpc:"state"`
	TotalEntryEncoding       interface{} `xmlrpc:"total_entry_encoding"`
	UserId                   interface{} `xmlrpc:"user_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var AccountBankStatementModel string = "account.bank.statement"

type AccountBankStatements []AccountBankStatement

type AccountBankStatementsNil []AccountBankStatementNil

func (s *AccountBankStatement) NilableType_() interface{} {
	return &AccountBankStatementNil{}
}

func (ns *AccountBankStatementNil) Type_() interface{} {
	s := &AccountBankStatement{}
	return load(ns, s)
}

func (s *AccountBankStatements) NilableType_() interface{} {
	return &AccountBankStatementsNil{}
}

func (ns *AccountBankStatementsNil) Type_() interface{} {
	s := &AccountBankStatements{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankStatement))
	}
	return s
}
