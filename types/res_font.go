package types

import (
	"time"
)

type ResFont struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Family      string    `xmlrpc:"family"`
	Id          int64     `xmlrpc:"id"`
	Mode        string    `xmlrpc:"mode"`
	Name        string    `xmlrpc:"name"`
	Path        string    `xmlrpc:"path"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResFontNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Family      interface{} `xmlrpc:"family"`
	Id          interface{} `xmlrpc:"id"`
	Mode        interface{} `xmlrpc:"mode"`
	Name        interface{} `xmlrpc:"name"`
	Path        interface{} `xmlrpc:"path"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResFontModel string = "res.font"

type ResFonts []ResFont

type ResFontsNil []ResFontNil

func (s *ResFont) NilableType_() interface{} {
	return &ResFontNil{}
}

func (ns *ResFontNil) Type_() interface{} {
	s := &ResFont{}
	return load(ns, s)
}

func (s *ResFonts) NilableType_() interface{} {
	return &ResFontsNil{}
}

func (ns *ResFontsNil) Type_() interface{} {
	s := &ResFonts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResFont))
	}
	return s
}
