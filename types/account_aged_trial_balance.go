package types

import (
	"time"
)

type AccountAgedTrialBalance struct {
	CompanyId       Many2One  `xmlrpc:"company_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DateFrom        time.Time `xmlrpc:"date_from"`
	DateTo          time.Time `xmlrpc:"date_to"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	JournalIds      []int64   `xmlrpc:"journal_ids"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	PeriodLength    int64     `xmlrpc:"period_length"`
	ResultSelection string    `xmlrpc:"result_selection"`
	TargetMove      string    `xmlrpc:"target_move"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountAgedTrialBalanceNil struct {
	CompanyId       interface{} `xmlrpc:"company_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DateFrom        interface{} `xmlrpc:"date_from"`
	DateTo          interface{} `xmlrpc:"date_to"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	JournalIds      interface{} `xmlrpc:"journal_ids"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	PeriodLength    interface{} `xmlrpc:"period_length"`
	ResultSelection interface{} `xmlrpc:"result_selection"`
	TargetMove      interface{} `xmlrpc:"target_move"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountAgedTrialBalanceModel string = "account.aged.trial.balance"

type AccountAgedTrialBalances []AccountAgedTrialBalance

type AccountAgedTrialBalancesNil []AccountAgedTrialBalanceNil

func (s *AccountAgedTrialBalance) NilableType_() interface{} {
	return &AccountAgedTrialBalanceNil{}
}

func (ns *AccountAgedTrialBalanceNil) Type_() interface{} {
	s := &AccountAgedTrialBalance{}
	return load(ns, s)
}

func (s *AccountAgedTrialBalances) NilableType_() interface{} {
	return &AccountAgedTrialBalancesNil{}
}

func (ns *AccountAgedTrialBalancesNil) Type_() interface{} {
	s := &AccountAgedTrialBalances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAgedTrialBalance))
	}
	return s
}
