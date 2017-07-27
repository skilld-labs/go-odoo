package types

import (
	"time"
)

type AccountFiscalPositionAccountTemplate struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	AccountDestId Many2One  `xmlrpc:"account_dest_id"`
	AccountSrcId  Many2One  `xmlrpc:"account_src_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	PositionId    Many2One  `xmlrpc:"position_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type AccountFiscalPositionAccountTemplateNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	AccountDestId interface{} `xmlrpc:"account_dest_id"`
	AccountSrcId  interface{} `xmlrpc:"account_src_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	PositionId    interface{} `xmlrpc:"position_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var AccountFiscalPositionAccountTemplateModel string = "account.fiscal.position.account.template"

type AccountFiscalPositionAccountTemplates []AccountFiscalPositionAccountTemplate

type AccountFiscalPositionAccountTemplatesNil []AccountFiscalPositionAccountTemplateNil

func (s *AccountFiscalPositionAccountTemplate) NilableType_() interface{} {
	return &AccountFiscalPositionAccountTemplateNil{}
}

func (ns *AccountFiscalPositionAccountTemplateNil) Type_() interface{} {
	s := &AccountFiscalPositionAccountTemplate{}
	return load(ns, s)
}

func (s *AccountFiscalPositionAccountTemplates) NilableType_() interface{} {
	return &AccountFiscalPositionAccountTemplatesNil{}
}

func (ns *AccountFiscalPositionAccountTemplatesNil) Type_() interface{} {
	s := &AccountFiscalPositionAccountTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPositionAccountTemplate))
	}
	return s
}
