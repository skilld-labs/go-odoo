package types

import (
	"time"
)

type ProcurementGroup struct {
	CreateDate  time.Time   `xmlrpc:"create_date"`
	CreateUid   Many2One    `xmlrpc:"create_uid"`
	DisplayName string      `xmlrpc:"display_name"`
	Id          int64       `xmlrpc:"id"`
	LastUpdate  time.Time   `xmlrpc:"__last_update"`
	MoveType    interface{} `xmlrpc:"move_type"`
	Name        string      `xmlrpc:"name"`
	PartnerId   Many2One    `xmlrpc:"partner_id"`
	SaleId      Many2One    `xmlrpc:"sale_id"`
	WriteDate   time.Time   `xmlrpc:"write_date"`
	WriteUid    Many2One    `xmlrpc:"write_uid"`
}

type ProcurementGroupNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	MoveType    interface{} `xmlrpc:"move_type"`
	Name        interface{} `xmlrpc:"name"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	SaleId      interface{} `xmlrpc:"sale_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProcurementGroupModel string = "procurement.group"

type ProcurementGroups []ProcurementGroup

type ProcurementGroupsNil []ProcurementGroupNil

func (s *ProcurementGroup) NilableType_() interface{} {
	return &ProcurementGroupNil{}
}

func (ns *ProcurementGroupNil) Type_() interface{} {
	s := &ProcurementGroup{}
	return load(ns, s)
}

func (s *ProcurementGroups) NilableType_() interface{} {
	return &ProcurementGroupsNil{}
}

func (ns *ProcurementGroupsNil) Type_() interface{} {
	s := &ProcurementGroups{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProcurementGroup))
	}
	return s
}
