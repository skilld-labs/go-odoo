package types

import (
	"time"
)

type Report struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ReportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ReportModel string = "report"

type Reports []Report

type ReportsNil []ReportNil

func (s *Report) NilableType_() interface{} {
	return &ReportNil{}
}

func (ns *ReportNil) Type_() interface{} {
	s := &Report{}
	return load(ns, s)
}

func (s *Reports) NilableType_() interface{} {
	return &ReportsNil{}
}

func (ns *ReportsNil) Type_() interface{} {
	s := &Reports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*Report))
	}
	return s
}
