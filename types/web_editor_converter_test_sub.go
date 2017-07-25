package types

import (
	"time"
)

type WebEditorConverterTestSub struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type WebEditorConverterTestSubNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var WebEditorConverterTestSubModel string = "web_editor.converter.test.sub"

type WebEditorConverterTestSubs []WebEditorConverterTestSub

type WebEditorConverterTestSubsNil []WebEditorConverterTestSubNil

func (s *WebEditorConverterTestSub) NilableType_() interface{} {
	return &WebEditorConverterTestSubNil{}
}

func (ns *WebEditorConverterTestSubNil) Type_() interface{} {
	s := &WebEditorConverterTestSub{}
	return load(ns, s)
}

func (s *WebEditorConverterTestSubs) NilableType_() interface{} {
	return &WebEditorConverterTestSubsNil{}
}

func (ns *WebEditorConverterTestSubsNil) Type_() interface{} {
	s := &WebEditorConverterTestSubs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WebEditorConverterTestSub))
	}
	return s
}
