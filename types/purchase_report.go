package types

import (
	"time"
)

type PurchaseReport struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	AccountAnalyticId   Many2One  `xmlrpc:"account_analytic_id"`
	CategoryId          Many2One  `xmlrpc:"category_id"`
	CommercialPartnerId Many2One  `xmlrpc:"commercial_partner_id"`
	CompanyId           Many2One  `xmlrpc:"company_id"`
	CountryId           Many2One  `xmlrpc:"country_id"`
	CurrencyId          Many2One  `xmlrpc:"currency_id"`
	DateApprove         time.Time `xmlrpc:"date_approve"`
	DateOrder           time.Time `xmlrpc:"date_order"`
	Delay               float64   `xmlrpc:"delay"`
	DelayPass           float64   `xmlrpc:"delay_pass"`
	DisplayName         string    `xmlrpc:"display_name"`
	FiscalPositionId    Many2One  `xmlrpc:"fiscal_position_id"`
	Id                  int64     `xmlrpc:"id"`
	NbrLines            int64     `xmlrpc:"nbr_lines"`
	Negociation         float64   `xmlrpc:"negociation"`
	PartnerId           Many2One  `xmlrpc:"partner_id"`
	PickingTypeId       Many2One  `xmlrpc:"picking_type_id"`
	PriceAverage        float64   `xmlrpc:"price_average"`
	PriceStandard       float64   `xmlrpc:"price_standard"`
	PriceTotal          float64   `xmlrpc:"price_total"`
	ProductId           Many2One  `xmlrpc:"product_id"`
	ProductTmplId       Many2One  `xmlrpc:"product_tmpl_id"`
	ProductUom          Many2One  `xmlrpc:"product_uom"`
	State               string    `xmlrpc:"state"`
	UnitQuantity        float64   `xmlrpc:"unit_quantity"`
	UserId              Many2One  `xmlrpc:"user_id"`
	Volume              float64   `xmlrpc:"volume"`
	Weight              float64   `xmlrpc:"weight"`
}

type PurchaseReportNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	AccountAnalyticId   interface{} `xmlrpc:"account_analytic_id"`
	CategoryId          interface{} `xmlrpc:"category_id"`
	CommercialPartnerId interface{} `xmlrpc:"commercial_partner_id"`
	CompanyId           interface{} `xmlrpc:"company_id"`
	CountryId           interface{} `xmlrpc:"country_id"`
	CurrencyId          interface{} `xmlrpc:"currency_id"`
	DateApprove         interface{} `xmlrpc:"date_approve"`
	DateOrder           interface{} `xmlrpc:"date_order"`
	Delay               interface{} `xmlrpc:"delay"`
	DelayPass           interface{} `xmlrpc:"delay_pass"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	FiscalPositionId    interface{} `xmlrpc:"fiscal_position_id"`
	Id                  interface{} `xmlrpc:"id"`
	NbrLines            interface{} `xmlrpc:"nbr_lines"`
	Negociation         interface{} `xmlrpc:"negociation"`
	PartnerId           interface{} `xmlrpc:"partner_id"`
	PickingTypeId       interface{} `xmlrpc:"picking_type_id"`
	PriceAverage        interface{} `xmlrpc:"price_average"`
	PriceStandard       interface{} `xmlrpc:"price_standard"`
	PriceTotal          interface{} `xmlrpc:"price_total"`
	ProductId           interface{} `xmlrpc:"product_id"`
	ProductTmplId       interface{} `xmlrpc:"product_tmpl_id"`
	ProductUom          interface{} `xmlrpc:"product_uom"`
	State               interface{} `xmlrpc:"state"`
	UnitQuantity        interface{} `xmlrpc:"unit_quantity"`
	UserId              interface{} `xmlrpc:"user_id"`
	Volume              interface{} `xmlrpc:"volume"`
	Weight              interface{} `xmlrpc:"weight"`
}

var PurchaseReportModel string = "purchase.report"

type PurchaseReports []PurchaseReport

type PurchaseReportsNil []PurchaseReportNil

func (s *PurchaseReport) NilableType_() interface{} {
	return &PurchaseReportNil{}
}

func (ns *PurchaseReportNil) Type_() interface{} {
	s := &PurchaseReport{}
	return load(ns, s)
}

func (s *PurchaseReports) NilableType_() interface{} {
	return &PurchaseReportsNil{}
}

func (ns *PurchaseReportsNil) Type_() interface{} {
	s := &PurchaseReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PurchaseReport))
	}
	return s
}
