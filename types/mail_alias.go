package types

import (
	"time"
)

type MailAlias struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	AliasContact        string    `xmlrpc:"alias_contact"`
	AliasDefaults       string    `xmlrpc:"alias_defaults"`
	AliasDomain         string    `xmlrpc:"alias_domain"`
	AliasForceThreadId  int64     `xmlrpc:"alias_force_thread_id"`
	AliasModelId        Many2One  `xmlrpc:"alias_model_id"`
	AliasName           string    `xmlrpc:"alias_name"`
	AliasParentModelId  Many2One  `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId int64     `xmlrpc:"alias_parent_thread_id"`
	AliasUserId         Many2One  `xmlrpc:"alias_user_id"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	DisplayName         string    `xmlrpc:"display_name"`
	Id                  int64     `xmlrpc:"id"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
}

type MailAliasNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	AliasContact        interface{} `xmlrpc:"alias_contact"`
	AliasDefaults       interface{} `xmlrpc:"alias_defaults"`
	AliasDomain         interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId  interface{} `xmlrpc:"alias_force_thread_id"`
	AliasModelId        interface{} `xmlrpc:"alias_model_id"`
	AliasName           interface{} `xmlrpc:"alias_name"`
	AliasParentModelId  interface{} `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId interface{} `xmlrpc:"alias_parent_thread_id"`
	AliasUserId         interface{} `xmlrpc:"alias_user_id"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	Id                  interface{} `xmlrpc:"id"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
}

var MailAliasModel string = "mail.alias"

type MailAliass []MailAlias

type MailAliassNil []MailAliasNil

func (s *MailAlias) NilableType_() interface{} {
	return &MailAliasNil{}
}

func (ns *MailAliasNil) Type_() interface{} {
	s := &MailAlias{}
	return load(ns, s)
}

func (s *MailAliass) NilableType_() interface{} {
	return &MailAliassNil{}
}

func (ns *MailAliassNil) Type_() interface{} {
	s := &MailAliass{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailAlias))
	}
	return s
}
