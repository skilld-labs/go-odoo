package types

import (
	"time"
)

type MailMessageSubtype struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	Default       bool      `xmlrpc:"default"`
	Description   string    `xmlrpc:"description"`
	DisplayName   string    `xmlrpc:"display_name"`
	Hidden        bool      `xmlrpc:"hidden"`
	Id            int64     `xmlrpc:"id"`
	Internal      bool      `xmlrpc:"internal"`
	Name          string    `xmlrpc:"name"`
	ParentId      Many2One  `xmlrpc:"parent_id"`
	RelationField string    `xmlrpc:"relation_field"`
	ResModel      string    `xmlrpc:"res_model"`
	Sequence      int64     `xmlrpc:"sequence"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type MailMessageSubtypeNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	Default       bool        `xmlrpc:"default"`
	Description   interface{} `xmlrpc:"description"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Hidden        bool        `xmlrpc:"hidden"`
	Id            interface{} `xmlrpc:"id"`
	Internal      bool        `xmlrpc:"internal"`
	Name          interface{} `xmlrpc:"name"`
	ParentId      interface{} `xmlrpc:"parent_id"`
	RelationField interface{} `xmlrpc:"relation_field"`
	ResModel      interface{} `xmlrpc:"res_model"`
	Sequence      interface{} `xmlrpc:"sequence"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var MailMessageSubtypeModel string = "mail.message.subtype"

type MailMessageSubtypes []MailMessageSubtype

type MailMessageSubtypesNil []MailMessageSubtypeNil

func (s *MailMessageSubtype) NilableType_() interface{} {
	return &MailMessageSubtypeNil{}
}

func (ns *MailMessageSubtypeNil) Type_() interface{} {
	s := &MailMessageSubtype{}
	return load(ns, s)
}

func (s *MailMessageSubtypes) NilableType_() interface{} {
	return &MailMessageSubtypesNil{}
}

func (ns *MailMessageSubtypesNil) Type_() interface{} {
	s := &MailMessageSubtypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMessageSubtype))
	}
	return s
}
