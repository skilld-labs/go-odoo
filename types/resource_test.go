package types

import (
	"time"
)

type ResourceTest struct {
	CompanyId          Many2One  `xmlrpc:"company_id"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	Name               string    `xmlrpc:"name"`
	ResourceCalendarId Many2One  `xmlrpc:"resource_calendar_id"`
	ResourceId         Many2One  `xmlrpc:"resource_id"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type ResourceTestNil struct {
	CompanyId          interface{} `xmlrpc:"company_id"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	Name               interface{} `xmlrpc:"name"`
	ResourceCalendarId interface{} `xmlrpc:"resource_calendar_id"`
	ResourceId         interface{} `xmlrpc:"resource_id"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var ResourceTestModel string = "resource.test"

type ResourceTests []ResourceTest

type ResourceTestsNil []ResourceTestNil

func (s *ResourceTest) NilableType_() interface{} {
	return &ResourceTestNil{}
}

func (ns *ResourceTestNil) Type_() interface{} {
	s := &ResourceTest{}
	return load(ns, s)
}

func (s *ResourceTests) NilableType_() interface{} {
	return &ResourceTestsNil{}
}

func (ns *ResourceTestsNil) Type_() interface{} {
	s := &ResourceTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceTest))
	}
	return s
}
