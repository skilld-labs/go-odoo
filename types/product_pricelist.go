package types

import (
	"time"
)

type ProductPricelist struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Active          bool      `xmlrpc:"active"`
	CompanyId       Many2One  `xmlrpc:"company_id"`
	CountryGroupIds []int64   `xmlrpc:"country_group_ids"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	CurrencyId      Many2One  `xmlrpc:"currency_id"`
	DiscountPolicy  string    `xmlrpc:"discount_policy"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	ItemIds         []int64   `xmlrpc:"item_ids"`
	Name            string    `xmlrpc:"name"`
	Sequence        int64     `xmlrpc:"sequence"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type ProductPricelistNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Active          bool        `xmlrpc:"active"`
	CompanyId       interface{} `xmlrpc:"company_id"`
	CountryGroupIds interface{} `xmlrpc:"country_group_ids"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	CurrencyId      interface{} `xmlrpc:"currency_id"`
	DiscountPolicy  interface{} `xmlrpc:"discount_policy"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	ItemIds         interface{} `xmlrpc:"item_ids"`
	Name            interface{} `xmlrpc:"name"`
	Sequence        interface{} `xmlrpc:"sequence"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var ProductPricelistModel string = "product.pricelist"

type ProductPricelists []ProductPricelist

type ProductPricelistsNil []ProductPricelistNil

func (s *ProductPricelist) NilableType_() interface{} {
	return &ProductPricelistNil{}
}

func (ns *ProductPricelistNil) Type_() interface{} {
	s := &ProductPricelist{}
	return load(ns, s)
}

func (s *ProductPricelists) NilableType_() interface{} {
	return &ProductPricelistsNil{}
}

func (ns *ProductPricelistsNil) Type_() interface{} {
	s := &ProductPricelists{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPricelist))
	}
	return s
}
