package types

import (
	"time"
)

type SmsSendSms struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Message     string    `xmlrpc:"message"`
	Recipients  string    `xmlrpc:"recipients"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type SmsSendSmsNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Message     interface{} `xmlrpc:"message"`
	Recipients  interface{} `xmlrpc:"recipients"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var SmsSendSmsModel string = "sms.send_sms"

type SmsSendSmss []SmsSendSms

type SmsSendSmssNil []SmsSendSmsNil

func (s *SmsSendSms) NilableType_() interface{} {
	return &SmsSendSmsNil{}
}

func (ns *SmsSendSmsNil) Type_() interface{} {
	s := &SmsSendSms{}
	return load(ns, s)
}

func (s *SmsSendSmss) NilableType_() interface{} {
	return &SmsSendSmssNil{}
}

func (ns *SmsSendSmssNil) Type_() interface{} {
	s := &SmsSendSmss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SmsSendSms))
	}
	return s
}
