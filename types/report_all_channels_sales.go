package types

import (
	"time"
)

type ReportAllChannelsSales struct {
	AnalyticAccountId Many2One  `xmlrpc:"analytic_account_id"`
	CategId           Many2One  `xmlrpc:"categ_id"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CountryId         Many2One  `xmlrpc:"country_id"`
	DateOrder         time.Time `xmlrpc:"date_order"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Name              string    `xmlrpc:"name"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	PricelistId       Many2One  `xmlrpc:"pricelist_id"`
	PriceSubtotal     float64   `xmlrpc:"price_subtotal"`
	PriceTotal        float64   `xmlrpc:"price_total"`
	ProductId         Many2One  `xmlrpc:"product_id"`
	ProductQty        float64   `xmlrpc:"product_qty"`
	ProductTmplId     Many2One  `xmlrpc:"product_tmpl_id"`
	TeamId            Many2One  `xmlrpc:"team_id"`
	UserId            Many2One  `xmlrpc:"user_id"`
}

type ReportAllChannelsSalesNil struct {
	AnalyticAccountId interface{} `xmlrpc:"analytic_account_id"`
	CategId           interface{} `xmlrpc:"categ_id"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CountryId         interface{} `xmlrpc:"country_id"`
	DateOrder         interface{} `xmlrpc:"date_order"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Name              interface{} `xmlrpc:"name"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	PricelistId       interface{} `xmlrpc:"pricelist_id"`
	PriceSubtotal     interface{} `xmlrpc:"price_subtotal"`
	PriceTotal        interface{} `xmlrpc:"price_total"`
	ProductId         interface{} `xmlrpc:"product_id"`
	ProductQty        interface{} `xmlrpc:"product_qty"`
	ProductTmplId     interface{} `xmlrpc:"product_tmpl_id"`
	TeamId            interface{} `xmlrpc:"team_id"`
	UserId            interface{} `xmlrpc:"user_id"`
}

var ReportAllChannelsSalesModel string = "report.all.channels.sales"

type ReportAllChannelsSaless []ReportAllChannelsSales

type ReportAllChannelsSalessNil []ReportAllChannelsSalesNil

func (s *ReportAllChannelsSales) NilableType_() interface{} {
	return &ReportAllChannelsSalesNil{}
}

func (ns *ReportAllChannelsSalesNil) Type_() interface{} {
	s := &ReportAllChannelsSales{}
	return load(ns, s)
}

func (s *ReportAllChannelsSaless) NilableType_() interface{} {
	return &ReportAllChannelsSalessNil{}
}

func (ns *ReportAllChannelsSalessNil) Type_() interface{} {
	s := &ReportAllChannelsSaless{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAllChannelsSales))
	}
	return s
}
