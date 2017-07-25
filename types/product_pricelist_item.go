package types

import (
	"time"
)

type ProductPricelistItem struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	AppliedOn       string    `xmlrpc:"applied_on"`
	Base            string    `xmlrpc:"base"`
	BasePricelistId Many2One  `xmlrpc:"base_pricelist_id"`
	CategId         Many2One  `xmlrpc:"categ_id"`
	CompanyId       Many2One  `xmlrpc:"company_id"`
	ComputePrice    string    `xmlrpc:"compute_price"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	CurrencyId      Many2One  `xmlrpc:"currency_id"`
	DateEnd         time.Time `xmlrpc:"date_end"`
	DateStart       time.Time `xmlrpc:"date_start"`
	DisplayName     string    `xmlrpc:"display_name"`
	FixedPrice      float64   `xmlrpc:"fixed_price"`
	Id              int64     `xmlrpc:"id"`
	MinQuantity     int64     `xmlrpc:"min_quantity"`
	Name            string    `xmlrpc:"name"`
	PercentPrice    float64   `xmlrpc:"percent_price"`
	Price           string    `xmlrpc:"price"`
	PriceDiscount   float64   `xmlrpc:"price_discount"`
	PriceMaxMargin  float64   `xmlrpc:"price_max_margin"`
	PriceMinMargin  float64   `xmlrpc:"price_min_margin"`
	PriceRound      float64   `xmlrpc:"price_round"`
	PriceSurcharge  float64   `xmlrpc:"price_surcharge"`
	PricelistId     Many2One  `xmlrpc:"pricelist_id"`
	ProductId       Many2One  `xmlrpc:"product_id"`
	ProductTmplId   Many2One  `xmlrpc:"product_tmpl_id"`
	Sequence        int64     `xmlrpc:"sequence"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type ProductPricelistItemNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	AppliedOn       interface{} `xmlrpc:"applied_on"`
	Base            interface{} `xmlrpc:"base"`
	BasePricelistId interface{} `xmlrpc:"base_pricelist_id"`
	CategId         interface{} `xmlrpc:"categ_id"`
	CompanyId       interface{} `xmlrpc:"company_id"`
	ComputePrice    interface{} `xmlrpc:"compute_price"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	CurrencyId      interface{} `xmlrpc:"currency_id"`
	DateEnd         interface{} `xmlrpc:"date_end"`
	DateStart       interface{} `xmlrpc:"date_start"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	FixedPrice      interface{} `xmlrpc:"fixed_price"`
	Id              interface{} `xmlrpc:"id"`
	MinQuantity     interface{} `xmlrpc:"min_quantity"`
	Name            interface{} `xmlrpc:"name"`
	PercentPrice    interface{} `xmlrpc:"percent_price"`
	Price           interface{} `xmlrpc:"price"`
	PriceDiscount   interface{} `xmlrpc:"price_discount"`
	PriceMaxMargin  interface{} `xmlrpc:"price_max_margin"`
	PriceMinMargin  interface{} `xmlrpc:"price_min_margin"`
	PriceRound      interface{} `xmlrpc:"price_round"`
	PriceSurcharge  interface{} `xmlrpc:"price_surcharge"`
	PricelistId     interface{} `xmlrpc:"pricelist_id"`
	ProductId       interface{} `xmlrpc:"product_id"`
	ProductTmplId   interface{} `xmlrpc:"product_tmpl_id"`
	Sequence        interface{} `xmlrpc:"sequence"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var ProductPricelistItemModel string = "product.pricelist.item"

type ProductPricelistItems []ProductPricelistItem

type ProductPricelistItemsNil []ProductPricelistItemNil

func (s *ProductPricelistItem) NilableType_() interface{} {
	return &ProductPricelistItemNil{}
}

func (ns *ProductPricelistItemNil) Type_() interface{} {
	s := &ProductPricelistItem{}
	return load(ns, s)
}

func (s *ProductPricelistItems) NilableType_() interface{} {
	return &ProductPricelistItemsNil{}
}

func (ns *ProductPricelistItemsNil) Type_() interface{} {
	s := &ProductPricelistItems{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPricelistItem))
	}
	return s
}
