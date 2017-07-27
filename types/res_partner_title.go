package types

import (
	"time"
)

type ResPartnerTitle struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Shortcut    string    `xmlrpc:"shortcut"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResPartnerTitleNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Shortcut    interface{} `xmlrpc:"shortcut"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResPartnerTitleModel string = "res.partner.title"

type ResPartnerTitles []ResPartnerTitle

type ResPartnerTitlesNil []ResPartnerTitleNil

func (s *ResPartnerTitle) NilableType_() interface{} {
	return &ResPartnerTitleNil{}
}

func (ns *ResPartnerTitleNil) Type_() interface{} {
	s := &ResPartnerTitle{}
	return load(ns, s)
}

func (s *ResPartnerTitles) NilableType_() interface{} {
	return &ResPartnerTitlesNil{}
}

func (ns *ResPartnerTitlesNil) Type_() interface{} {
	s := &ResPartnerTitles{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResPartnerTitle))
	}
	return s
}
