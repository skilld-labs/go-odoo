package types

import (
	"time"
)

type MailAliasMixin struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	AliasContact        string    `xmlrpc:"alias_contact"`
	AliasDefaults       string    `xmlrpc:"alias_defaults"`
	AliasDomain         string    `xmlrpc:"alias_domain"`
	AliasForceThreadId  int64     `xmlrpc:"alias_force_thread_id"`
	AliasId             Many2One  `xmlrpc:"alias_id"`
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

type MailAliasMixinNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	AliasContact        interface{} `xmlrpc:"alias_contact"`
	AliasDefaults       interface{} `xmlrpc:"alias_defaults"`
	AliasDomain         interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId  interface{} `xmlrpc:"alias_force_thread_id"`
	AliasId             interface{} `xmlrpc:"alias_id"`
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

var MailAliasMixinModel string = "mail.alias.mixin"

type MailAliasMixins []MailAliasMixin

type MailAliasMixinsNil []MailAliasMixinNil

func (s *MailAliasMixin) NilableType_() interface{} {
	return &MailAliasMixinNil{}
}

func (ns *MailAliasMixinNil) Type_() interface{} {
	s := &MailAliasMixin{}
	return load(ns, s)
}

func (s *MailAliasMixins) NilableType_() interface{} {
	return &MailAliasMixinsNil{}
}

func (ns *MailAliasMixinsNil) Type_() interface{} {
	s := &MailAliasMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailAliasMixin))
	}
	return s
}
