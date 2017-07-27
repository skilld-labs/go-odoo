package types

import (
	"time"
)

type IrActionsActUrl struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Help        string    `xmlrpc:"help"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Target      string    `xmlrpc:"target"`
	Type        string    `xmlrpc:"type"`
	Url         string    `xmlrpc:"url"`
	Usage       string    `xmlrpc:"usage"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	XmlId       string    `xmlrpc:"xml_id"`
}

type IrActionsActUrlNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Help        interface{} `xmlrpc:"help"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Target      interface{} `xmlrpc:"target"`
	Type        interface{} `xmlrpc:"type"`
	Url         interface{} `xmlrpc:"url"`
	Usage       interface{} `xmlrpc:"usage"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	XmlId       interface{} `xmlrpc:"xml_id"`
}

var IrActionsActUrlModel string = "ir.actions.act_url"

type IrActionsActUrls []IrActionsActUrl

type IrActionsActUrlsNil []IrActionsActUrlNil

func (s *IrActionsActUrl) NilableType_() interface{} {
	return &IrActionsActUrlNil{}
}

func (ns *IrActionsActUrlNil) Type_() interface{} {
	s := &IrActionsActUrl{}
	return load(ns, s)
}

func (s *IrActionsActUrls) NilableType_() interface{} {
	return &IrActionsActUrlsNil{}
}

func (ns *IrActionsActUrlsNil) Type_() interface{} {
	s := &IrActionsActUrls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActUrl))
	}
	return s
}
