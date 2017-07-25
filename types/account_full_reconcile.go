package types

import (
	"time"
)

type AccountFullReconcile struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	DisplayName          string    `xmlrpc:"display_name"`
	ExchangeMoveId       Many2One  `xmlrpc:"exchange_move_id"`
	ExchangePartialRecId Many2One  `xmlrpc:"exchange_partial_rec_id"`
	Id                   int64     `xmlrpc:"id"`
	Name                 string    `xmlrpc:"name"`
	PartialReconcileIds  []int64   `xmlrpc:"partial_reconcile_ids"`
	ReconciledLineIds    []int64   `xmlrpc:"reconciled_line_ids"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type AccountFullReconcileNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	ExchangeMoveId       interface{} `xmlrpc:"exchange_move_id"`
	ExchangePartialRecId interface{} `xmlrpc:"exchange_partial_rec_id"`
	Id                   interface{} `xmlrpc:"id"`
	Name                 interface{} `xmlrpc:"name"`
	PartialReconcileIds  interface{} `xmlrpc:"partial_reconcile_ids"`
	ReconciledLineIds    interface{} `xmlrpc:"reconciled_line_ids"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var AccountFullReconcileModel string = "account.full.reconcile"

type AccountFullReconciles []AccountFullReconcile

type AccountFullReconcilesNil []AccountFullReconcileNil

func (s *AccountFullReconcile) NilableType_() interface{} {
	return &AccountFullReconcileNil{}
}

func (ns *AccountFullReconcileNil) Type_() interface{} {
	s := &AccountFullReconcile{}
	return load(ns, s)
}

func (s *AccountFullReconciles) NilableType_() interface{} {
	return &AccountFullReconcilesNil{}
}

func (ns *AccountFullReconcilesNil) Type_() interface{} {
	s := &AccountFullReconciles{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFullReconcile))
	}
	return s
}
