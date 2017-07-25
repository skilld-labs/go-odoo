package types

import (
	"time"
)

type LinkTrackerCode struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Code        string    `xmlrpc:"code"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LinkId      Many2One  `xmlrpc:"link_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type LinkTrackerCodeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Code        interface{} `xmlrpc:"code"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LinkId      interface{} `xmlrpc:"link_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var LinkTrackerCodeModel string = "link.tracker.code"

type LinkTrackerCodes []LinkTrackerCode

type LinkTrackerCodesNil []LinkTrackerCodeNil

func (s *LinkTrackerCode) NilableType_() interface{} {
	return &LinkTrackerCodeNil{}
}

func (ns *LinkTrackerCodeNil) Type_() interface{} {
	s := &LinkTrackerCode{}
	return load(ns, s)
}

func (s *LinkTrackerCodes) NilableType_() interface{} {
	return &LinkTrackerCodesNil{}
}

func (ns *LinkTrackerCodesNil) Type_() interface{} {
	s := &LinkTrackerCodes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*LinkTrackerCode))
	}
	return s
}
