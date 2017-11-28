package types

import (
	"time"
)

type AccountFinancialYearOp struct {
	AccountSetupFyDataDone bool        `xmlrpc:"account_setup_fy_data_done"`
	CompanyId              Many2One    `xmlrpc:"company_id"`
	CreateDate             time.Time   `xmlrpc:"create_date"`
	CreateUid              Many2One    `xmlrpc:"create_uid"`
	DisplayName            string      `xmlrpc:"display_name"`
	FiscalyearLastDay      int64       `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth    interface{} `xmlrpc:"fiscalyear_last_month"`
	Id                     int64       `xmlrpc:"id"`
	LastUpdate             time.Time   `xmlrpc:"__last_update"`
	OpeningDate            time.Time   `xmlrpc:"opening_date"`
	OpeningMovePosted      bool        `xmlrpc:"opening_move_posted"`
	WriteDate              time.Time   `xmlrpc:"write_date"`
	WriteUid               Many2One    `xmlrpc:"write_uid"`
}

type AccountFinancialYearOpNil struct {
	AccountSetupFyDataDone bool        `xmlrpc:"account_setup_fy_data_done"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	FiscalyearLastDay      interface{} `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth    interface{} `xmlrpc:"fiscalyear_last_month"`
	Id                     interface{} `xmlrpc:"id"`
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	OpeningDate            interface{} `xmlrpc:"opening_date"`
	OpeningMovePosted      bool        `xmlrpc:"opening_move_posted"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var AccountFinancialYearOpModel string = "account.financial.year.op"

type AccountFinancialYearOps []AccountFinancialYearOp

type AccountFinancialYearOpsNil []AccountFinancialYearOpNil

func (s *AccountFinancialYearOp) NilableType_() interface{} {
	return &AccountFinancialYearOpNil{}
}

func (ns *AccountFinancialYearOpNil) Type_() interface{} {
	s := &AccountFinancialYearOp{}
	return load(ns, s)
}

func (s *AccountFinancialYearOps) NilableType_() interface{} {
	return &AccountFinancialYearOpsNil{}
}

func (ns *AccountFinancialYearOpsNil) Type_() interface{} {
	s := &AccountFinancialYearOps{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFinancialYearOp))
	}
	return s
}
