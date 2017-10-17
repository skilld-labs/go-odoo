package types

import (
	"time"
)

type AutosalesConfigLine struct {
	AutosalesConfigId  Many2One  `xmlrpc:"autosales_config_id"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	PercentProductBase float64   `xmlrpc:"percent_product_base"`
	ProductAuto        Many2One  `xmlrpc:"product_auto"`
	ProductBase        Many2One  `xmlrpc:"product_base"`
	Sequence           int64     `xmlrpc:"sequence"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type AutosalesConfigLineNil struct {
	AutosalesConfigId  interface{} `xmlrpc:"autosales_config_id"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	PercentProductBase interface{} `xmlrpc:"percent_product_base"`
	ProductAuto        interface{} `xmlrpc:"product_auto"`
	ProductBase        interface{} `xmlrpc:"product_base"`
	Sequence           interface{} `xmlrpc:"sequence"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var AutosalesConfigLineModel string = "autosales.config.line"

type AutosalesConfigLines []AutosalesConfigLine

type AutosalesConfigLinesNil []AutosalesConfigLineNil

func (s *AutosalesConfigLine) NilableType_() interface{} {
	return &AutosalesConfigLineNil{}
}

func (ns *AutosalesConfigLineNil) Type_() interface{} {
	s := &AutosalesConfigLine{}
	return load(ns, s)
}

func (s *AutosalesConfigLines) NilableType_() interface{} {
	return &AutosalesConfigLinesNil{}
}

func (ns *AutosalesConfigLinesNil) Type_() interface{} {
	s := &AutosalesConfigLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AutosalesConfigLine))
	}
	return s
}
