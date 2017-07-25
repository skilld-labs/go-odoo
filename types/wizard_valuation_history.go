package types

import (
	"time"
)

type WizardValuationHistory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ChooseDate  bool      `xmlrpc:"choose_date"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Date        time.Time `xmlrpc:"date"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type WizardValuationHistoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ChooseDate  bool        `xmlrpc:"choose_date"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Date        interface{} `xmlrpc:"date"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var WizardValuationHistoryModel string = "wizard.valuation.history"

type WizardValuationHistorys []WizardValuationHistory

type WizardValuationHistorysNil []WizardValuationHistoryNil

func (s *WizardValuationHistory) NilableType_() interface{} {
	return &WizardValuationHistoryNil{}
}

func (ns *WizardValuationHistoryNil) Type_() interface{} {
	s := &WizardValuationHistory{}
	return load(ns, s)
}

func (s *WizardValuationHistorys) NilableType_() interface{} {
	return &WizardValuationHistorysNil{}
}

func (ns *WizardValuationHistorysNil) Type_() interface{} {
	s := &WizardValuationHistorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WizardValuationHistory))
	}
	return s
}
