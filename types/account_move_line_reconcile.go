package types

import (
	"time"
)

type AccountMoveLineReconcile struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Credit      float64   `xmlrpc:"credit"`
	Debit       float64   `xmlrpc:"debit"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	TransNbr    int64     `xmlrpc:"trans_nbr"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	Writeoff    float64   `xmlrpc:"writeoff"`
}

type AccountMoveLineReconcileNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Credit      interface{} `xmlrpc:"credit"`
	Debit       interface{} `xmlrpc:"debit"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	TransNbr    interface{} `xmlrpc:"trans_nbr"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	Writeoff    interface{} `xmlrpc:"writeoff"`
}

var AccountMoveLineReconcileModel string = "account.move.line.reconcile"

type AccountMoveLineReconciles []AccountMoveLineReconcile

type AccountMoveLineReconcilesNil []AccountMoveLineReconcileNil

func (s *AccountMoveLineReconcile) NilableType_() interface{} {
	return &AccountMoveLineReconcileNil{}
}

func (ns *AccountMoveLineReconcileNil) Type_() interface{} {
	s := &AccountMoveLineReconcile{}
	return load(ns, s)
}

func (s *AccountMoveLineReconciles) NilableType_() interface{} {
	return &AccountMoveLineReconcilesNil{}
}

func (ns *AccountMoveLineReconcilesNil) Type_() interface{} {
	s := &AccountMoveLineReconciles{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountMoveLineReconcile))
	}
	return s
}
