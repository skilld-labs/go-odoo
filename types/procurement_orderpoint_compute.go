package types

import (
	"time"
)

type ProcurementOrderpointCompute struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProcurementOrderpointComputeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProcurementOrderpointComputeModel string = "procurement.orderpoint.compute"

type ProcurementOrderpointComputes []ProcurementOrderpointCompute

type ProcurementOrderpointComputesNil []ProcurementOrderpointComputeNil

func (s *ProcurementOrderpointCompute) NilableType_() interface{} {
	return &ProcurementOrderpointComputeNil{}
}

func (ns *ProcurementOrderpointComputeNil) Type_() interface{} {
	s := &ProcurementOrderpointCompute{}
	return load(ns, s)
}

func (s *ProcurementOrderpointComputes) NilableType_() interface{} {
	return &ProcurementOrderpointComputesNil{}
}

func (ns *ProcurementOrderpointComputesNil) Type_() interface{} {
	s := &ProcurementOrderpointComputes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProcurementOrderpointCompute))
	}
	return s
}
