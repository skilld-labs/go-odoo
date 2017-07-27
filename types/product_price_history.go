package types

import (
	"time"
)

type ProductPriceHistory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	Cost        float64   `xmlrpc:"cost"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Datetime    time.Time `xmlrpc:"datetime"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	ProductId   Many2One  `xmlrpc:"product_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductPriceHistoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	Cost        interface{} `xmlrpc:"cost"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Datetime    interface{} `xmlrpc:"datetime"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	ProductId   interface{} `xmlrpc:"product_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductPriceHistoryModel string = "product.price.history"

type ProductPriceHistorys []ProductPriceHistory

type ProductPriceHistorysNil []ProductPriceHistoryNil

func (s *ProductPriceHistory) NilableType_() interface{} {
	return &ProductPriceHistoryNil{}
}

func (ns *ProductPriceHistoryNil) Type_() interface{} {
	s := &ProductPriceHistory{}
	return load(ns, s)
}

func (s *ProductPriceHistorys) NilableType_() interface{} {
	return &ProductPriceHistorysNil{}
}

func (ns *ProductPriceHistorysNil) Type_() interface{} {
	s := &ProductPriceHistorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPriceHistory))
	}
	return s
}
