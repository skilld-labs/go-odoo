package types

import (
	"time"
)

type PrintPrenumberedChecks struct {
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	NextCheckNumber int64     `xmlrpc:"next_check_number"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type PrintPrenumberedChecksNil struct {
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	NextCheckNumber interface{} `xmlrpc:"next_check_number"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var PrintPrenumberedChecksModel string = "print.prenumbered.checks"

type PrintPrenumberedCheckss []PrintPrenumberedChecks

type PrintPrenumberedCheckssNil []PrintPrenumberedChecksNil

func (s *PrintPrenumberedChecks) NilableType_() interface{} {
	return &PrintPrenumberedChecksNil{}
}

func (ns *PrintPrenumberedChecksNil) Type_() interface{} {
	s := &PrintPrenumberedChecks{}
	return load(ns, s)
}

func (s *PrintPrenumberedCheckss) NilableType_() interface{} {
	return &PrintPrenumberedCheckssNil{}
}

func (ns *PrintPrenumberedCheckssNil) Type_() interface{} {
	s := &PrintPrenumberedCheckss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PrintPrenumberedChecks))
	}
	return s
}
