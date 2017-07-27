package types

import (
	"time"
)

type WebEditorConverterTest struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	Binary       string    `xmlrpc:"binary"`
	Char         string    `xmlrpc:"char"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	Date         time.Time `xmlrpc:"date"`
	Datetime     time.Time `xmlrpc:"datetime"`
	DisplayName  string    `xmlrpc:"display_name"`
	Float        float64   `xmlrpc:"float"`
	Html         string    `xmlrpc:"html"`
	Id           int64     `xmlrpc:"id"`
	Integer      int64     `xmlrpc:"integer"`
	Many2one     Many2One  `xmlrpc:"many2one"`
	Numeric      float64   `xmlrpc:"numeric"`
	Selection    string    `xmlrpc:"selection"`
	SelectionStr string    `xmlrpc:"selection_str"`
	Text         string    `xmlrpc:"text"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type WebEditorConverterTestNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	Binary       interface{} `xmlrpc:"binary"`
	Char         interface{} `xmlrpc:"char"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	Date         interface{} `xmlrpc:"date"`
	Datetime     interface{} `xmlrpc:"datetime"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Float        interface{} `xmlrpc:"float"`
	Html         interface{} `xmlrpc:"html"`
	Id           interface{} `xmlrpc:"id"`
	Integer      interface{} `xmlrpc:"integer"`
	Many2one     interface{} `xmlrpc:"many2one"`
	Numeric      interface{} `xmlrpc:"numeric"`
	Selection    interface{} `xmlrpc:"selection"`
	SelectionStr interface{} `xmlrpc:"selection_str"`
	Text         interface{} `xmlrpc:"text"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var WebEditorConverterTestModel string = "web_editor.converter.test"

type WebEditorConverterTests []WebEditorConverterTest

type WebEditorConverterTestsNil []WebEditorConverterTestNil

func (s *WebEditorConverterTest) NilableType_() interface{} {
	return &WebEditorConverterTestNil{}
}

func (ns *WebEditorConverterTestNil) Type_() interface{} {
	s := &WebEditorConverterTest{}
	return load(ns, s)
}

func (s *WebEditorConverterTests) NilableType_() interface{} {
	return &WebEditorConverterTestsNil{}
}

func (ns *WebEditorConverterTestsNil) Type_() interface{} {
	s := &WebEditorConverterTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WebEditorConverterTest))
	}
	return s
}
