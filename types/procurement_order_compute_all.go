package types

import (
	"time"
)

type ProcurementOrderComputeAll struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProcurementOrderComputeAllNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProcurementOrderComputeAllModel string = "procurement.order.compute.all"

type ProcurementOrderComputeAlls []ProcurementOrderComputeAll

type ProcurementOrderComputeAllsNil []ProcurementOrderComputeAllNil

func (s *ProcurementOrderComputeAll) NilableType_() interface{} {
	return &ProcurementOrderComputeAllNil{}
}

func (ns *ProcurementOrderComputeAllNil) Type_() interface{} {
	s := &ProcurementOrderComputeAll{}
	return load(ns, s)
}

func (s *ProcurementOrderComputeAlls) NilableType_() interface{} {
	return &ProcurementOrderComputeAllsNil{}
}

func (ns *ProcurementOrderComputeAllsNil) Type_() interface{} {
	s := &ProcurementOrderComputeAlls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProcurementOrderComputeAll))
	}
	return s
}
