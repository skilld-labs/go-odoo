package types

import (
	"time"
)

type IrCron struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Active         bool      `xmlrpc:"active"`
	Args           string    `xmlrpc:"args"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Doall          bool      `xmlrpc:"doall"`
	Function       string    `xmlrpc:"function"`
	Id             int64     `xmlrpc:"id"`
	IntervalNumber int64     `xmlrpc:"interval_number"`
	IntervalType   string    `xmlrpc:"interval_type"`
	Model          string    `xmlrpc:"model"`
	Name           string    `xmlrpc:"name"`
	Nextcall       time.Time `xmlrpc:"nextcall"`
	Numbercall     int64     `xmlrpc:"numbercall"`
	Priority       int64     `xmlrpc:"priority"`
	UserId         Many2One  `xmlrpc:"user_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type IrCronNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Active         bool        `xmlrpc:"active"`
	Args           interface{} `xmlrpc:"args"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Doall          bool        `xmlrpc:"doall"`
	Function       interface{} `xmlrpc:"function"`
	Id             interface{} `xmlrpc:"id"`
	IntervalNumber interface{} `xmlrpc:"interval_number"`
	IntervalType   interface{} `xmlrpc:"interval_type"`
	Model          interface{} `xmlrpc:"model"`
	Name           interface{} `xmlrpc:"name"`
	Nextcall       interface{} `xmlrpc:"nextcall"`
	Numbercall     interface{} `xmlrpc:"numbercall"`
	Priority       interface{} `xmlrpc:"priority"`
	UserId         interface{} `xmlrpc:"user_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var IrCronModel string = "ir.cron"

type IrCrons []IrCron

type IrCronsNil []IrCronNil

func (s *IrCron) NilableType_() interface{} {
	return &IrCronNil{}
}

func (ns *IrCronNil) Type_() interface{} {
	s := &IrCron{}
	return load(ns, s)
}

func (s *IrCrons) NilableType_() interface{} {
	return &IrCronsNil{}
}

func (ns *IrCronsNil) Type_() interface{} {
	s := &IrCrons{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrCron))
	}
	return s
}
