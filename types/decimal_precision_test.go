package types

import (
	"time"
)

type DecimalPrecisionTest struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Float       float64   `xmlrpc:"float"`
	Float2      float64   `xmlrpc:"float_2"`
	Float4      float64   `xmlrpc:"float_4"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type DecimalPrecisionTestNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Float       interface{} `xmlrpc:"float"`
	Float2      interface{} `xmlrpc:"float_2"`
	Float4      interface{} `xmlrpc:"float_4"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var DecimalPrecisionTestModel string = "decimal.precision.test"

type DecimalPrecisionTests []DecimalPrecisionTest

type DecimalPrecisionTestsNil []DecimalPrecisionTestNil

func (s *DecimalPrecisionTest) NilableType_() interface{} {
	return &DecimalPrecisionTestNil{}
}

func (ns *DecimalPrecisionTestNil) Type_() interface{} {
	s := &DecimalPrecisionTest{}
	return load(ns, s)
}

func (s *DecimalPrecisionTests) NilableType_() interface{} {
	return &DecimalPrecisionTestsNil{}
}

func (ns *DecimalPrecisionTestsNil) Type_() interface{} {
	s := &DecimalPrecisionTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*DecimalPrecisionTest))
	}
	return s
}
