package types

import (
	"time"
)

type BarcodeRule struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	Alias                 string    `xmlrpc:"alias"`
	BarcodeNomenclatureId Many2One  `xmlrpc:"barcode_nomenclature_id"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Encoding              string    `xmlrpc:"encoding"`
	Id                    int64     `xmlrpc:"id"`
	Name                  string    `xmlrpc:"name"`
	Pattern               string    `xmlrpc:"pattern"`
	Sequence              int64     `xmlrpc:"sequence"`
	Type                  string    `xmlrpc:"type"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type BarcodeRuleNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	Alias                 interface{} `xmlrpc:"alias"`
	BarcodeNomenclatureId interface{} `xmlrpc:"barcode_nomenclature_id"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Encoding              interface{} `xmlrpc:"encoding"`
	Id                    interface{} `xmlrpc:"id"`
	Name                  interface{} `xmlrpc:"name"`
	Pattern               interface{} `xmlrpc:"pattern"`
	Sequence              interface{} `xmlrpc:"sequence"`
	Type                  interface{} `xmlrpc:"type"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var BarcodeRuleModel string = "barcode.rule"

type BarcodeRules []BarcodeRule

type BarcodeRulesNil []BarcodeRuleNil

func (s *BarcodeRule) NilableType_() interface{} {
	return &BarcodeRuleNil{}
}

func (ns *BarcodeRuleNil) Type_() interface{} {
	s := &BarcodeRule{}
	return load(ns, s)
}

func (s *BarcodeRules) NilableType_() interface{} {
	return &BarcodeRulesNil{}
}

func (ns *BarcodeRulesNil) Type_() interface{} {
	s := &BarcodeRules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BarcodeRule))
	}
	return s
}
