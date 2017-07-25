package types

import (
	"time"
)

type IrMailServer struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Active         bool      `xmlrpc:"active"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	Name           string    `xmlrpc:"name"`
	Sequence       int64     `xmlrpc:"sequence"`
	SmtpDebug      bool      `xmlrpc:"smtp_debug"`
	SmtpEncryption string    `xmlrpc:"smtp_encryption"`
	SmtpHost       string    `xmlrpc:"smtp_host"`
	SmtpPass       string    `xmlrpc:"smtp_pass"`
	SmtpPort       int64     `xmlrpc:"smtp_port"`
	SmtpUser       string    `xmlrpc:"smtp_user"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type IrMailServerNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Active         bool        `xmlrpc:"active"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	Name           interface{} `xmlrpc:"name"`
	Sequence       interface{} `xmlrpc:"sequence"`
	SmtpDebug      bool        `xmlrpc:"smtp_debug"`
	SmtpEncryption interface{} `xmlrpc:"smtp_encryption"`
	SmtpHost       interface{} `xmlrpc:"smtp_host"`
	SmtpPass       interface{} `xmlrpc:"smtp_pass"`
	SmtpPort       interface{} `xmlrpc:"smtp_port"`
	SmtpUser       interface{} `xmlrpc:"smtp_user"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var IrMailServerModel string = "ir.mail_server"

type IrMailServers []IrMailServer

type IrMailServersNil []IrMailServerNil

func (s *IrMailServer) NilableType_() interface{} {
	return &IrMailServerNil{}
}

func (ns *IrMailServerNil) Type_() interface{} {
	s := &IrMailServer{}
	return load(ns, s)
}

func (s *IrMailServers) NilableType_() interface{} {
	return &IrMailServersNil{}
}

func (ns *IrMailServersNil) Type_() interface{} {
	s := &IrMailServers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrMailServer))
	}
	return s
}
