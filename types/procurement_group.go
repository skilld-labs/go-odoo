package types

import (
	"time"
)

type ProcurementGroup struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	MoveType       string    `xmlrpc:"move_type"`
	Name           string    `xmlrpc:"name"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	ProcurementIds []int64   `xmlrpc:"procurement_ids"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ProcurementGroupNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	MoveType       interface{} `xmlrpc:"move_type"`
	Name           interface{} `xmlrpc:"name"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	ProcurementIds interface{} `xmlrpc:"procurement_ids"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
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
