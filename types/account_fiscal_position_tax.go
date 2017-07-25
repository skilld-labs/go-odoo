package types

import (
	"time"
)

type AccountFiscalPositionTax struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PositionId  Many2One  `xmlrpc:"position_id"`
	TaxDestId   Many2One  `xmlrpc:"tax_dest_id"`
	TaxSrcId    Many2One  `xmlrpc:"tax_src_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountFiscalPositionTaxNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PositionId  interface{} `xmlrpc:"position_id"`
	TaxDestId   interface{} `xmlrpc:"tax_dest_id"`
	TaxSrcId    interface{} `xmlrpc:"tax_src_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountFiscalPositionTaxModel string = "account.fiscal.position.tax"

type AccountFiscalPositionTaxs []AccountFiscalPositionTax

type AccountFiscalPositionTaxsNil []AccountFiscalPositionTaxNil

func (s *AccountFiscalPositionTax) NilableType_() interface{} {
	return &AccountFiscalPositionTaxNil{}
}

func (ns *AccountFiscalPositionTaxNil) Type_() interface{} {
	s := &AccountFiscalPositionTax{}
	return load(ns, s)
}

func (s *AccountFiscalPositionTaxs) NilableType_() interface{} {
	return &AccountFiscalPositionTaxsNil{}
}

func (ns *AccountFiscalPositionTaxsNil) Type_() interface{} {
	s := &AccountFiscalPositionTaxs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPositionTax))
	}
	return s
}
