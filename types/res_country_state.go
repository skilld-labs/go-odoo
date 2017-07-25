package types

import (
	"time"
)

type ResCountryState struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Code        string    `xmlrpc:"code"`
	CountryId   Many2One  `xmlrpc:"country_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResCountryStateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Code        interface{} `xmlrpc:"code"`
	CountryId   interface{} `xmlrpc:"country_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResCountryStateModel string = "res.country.state"

type ResCountryStates []ResCountryState

type ResCountryStatesNil []ResCountryStateNil

func (s *ResCountryState) NilableType_() interface{} {
	return &ResCountryStateNil{}
}

func (ns *ResCountryStateNil) Type_() interface{} {
	s := &ResCountryState{}
	return load(ns, s)
}

func (s *ResCountryStates) NilableType_() interface{} {
	return &ResCountryStatesNil{}
}

func (ns *ResCountryStatesNil) Type_() interface{} {
	s := &ResCountryStates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCountryState))
	}
	return s
}
