package types

import (
	"time"
)

type SmsApi struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type SmsApiNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var SmsApiModel string = "sms.api"

type SmsApis []SmsApi

type SmsApisNil []SmsApiNil

func (s *SmsApi) NilableType_() interface{} {
	return &SmsApiNil{}
}

func (ns *SmsApiNil) Type_() interface{} {
	s := &SmsApi{}
	return load(ns, s)
}

func (s *SmsApis) NilableType_() interface{} {
	return &SmsApisNil{}
}

func (ns *SmsApisNil) Type_() interface{} {
	s := &SmsApis{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SmsApi))
	}
	return s
}
