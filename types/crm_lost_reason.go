package types

import (
	"time"
)

type CrmLostReason struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CrmLostReasonNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CrmLostReasonModel string = "crm.lost.reason"

type CrmLostReasons []CrmLostReason

type CrmLostReasonsNil []CrmLostReasonNil

func (s *CrmLostReason) NilableType_() interface{} {
	return &CrmLostReasonNil{}
}

func (ns *CrmLostReasonNil) Type_() interface{} {
	s := &CrmLostReason{}
	return load(ns, s)
}

func (s *CrmLostReasons) NilableType_() interface{} {
	return &CrmLostReasonsNil{}
}

func (ns *CrmLostReasonsNil) Type_() interface{} {
	s := &CrmLostReasons{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLostReason))
	}
	return s
}
