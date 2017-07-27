package types

import (
	"time"
)

type ResUsersLog struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResUsersLogNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResUsersLogModel string = "res.users.log"

type ResUsersLogs []ResUsersLog

type ResUsersLogsNil []ResUsersLogNil

func (s *ResUsersLog) NilableType_() interface{} {
	return &ResUsersLogNil{}
}

func (ns *ResUsersLogNil) Type_() interface{} {
	s := &ResUsersLog{}
	return load(ns, s)
}

func (s *ResUsersLogs) NilableType_() interface{} {
	return &ResUsersLogsNil{}
}

func (ns *ResUsersLogsNil) Type_() interface{} {
	s := &ResUsersLogs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResUsersLog))
	}
	return s
}
