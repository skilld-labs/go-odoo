package types

import (
	"time"
)

type DecimalPrecision struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Digits      int64     `xmlrpc:"digits"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type DecimalPrecisionNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Digits      interface{} `xmlrpc:"digits"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var DecimalPrecisionModel string = "decimal.precision"

type DecimalPrecisions []DecimalPrecision

type DecimalPrecisionsNil []DecimalPrecisionNil

func (s *DecimalPrecision) NilableType_() interface{} {
	return &DecimalPrecisionNil{}
}

func (ns *DecimalPrecisionNil) Type_() interface{} {
	s := &DecimalPrecision{}
	return load(ns, s)
}

func (s *DecimalPrecisions) NilableType_() interface{} {
	return &DecimalPrecisionsNil{}
}

func (ns *DecimalPrecisionsNil) Type_() interface{} {
	s := &DecimalPrecisions{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*DecimalPrecision))
	}
	return s
}
