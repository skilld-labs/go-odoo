package types

import (
	"time"
)

type SaleReport struct {
	AmtInvoiced         float64   `xmlrpc:"amt_invoiced"`
	AmtToInvoice        float64   `xmlrpc:"amt_to_invoice"`
	AnalyticAccountId   Many2One  `xmlrpc:"analytic_account_id"`
	CategId             Many2One  `xmlrpc:"categ_id"`
	CommercialPartnerId Many2One  `xmlrpc:"commercial_partner_id"`
	CompanyId           Many2One  `xmlrpc:"company_id"`
	ConfirmationDate    time.Time `xmlrpc:"confirmation_date"`
	CountryId           Many2One  `xmlrpc:"country_id"`
	Date                time.Time `xmlrpc:"date"`
	DisplayName         string    `xmlrpc:"display_name"`
	Id                  int64     `xmlrpc:"id"`
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	Name                string    `xmlrpc:"name"`
	Nbr                 int64     `xmlrpc:"nbr"`
	PartnerId           Many2One  `xmlrpc:"partner_id"`
	PricelistId         Many2One  `xmlrpc:"pricelist_id"`
	PriceSubtotal       float64   `xmlrpc:"price_subtotal"`
	PriceTotal          float64   `xmlrpc:"price_total"`
	ProductId           Many2One  `xmlrpc:"product_id"`
	ProductTmplId       Many2One  `xmlrpc:"product_tmpl_id"`
	ProductUom          Many2One  `xmlrpc:"product_uom"`
	ProductUomQty       float64   `xmlrpc:"product_uom_qty"`
	QtyDelivered        float64   `xmlrpc:"qty_delivered"`
	QtyInvoiced         float64   `xmlrpc:"qty_invoiced"`
	QtyToInvoice        float64   `xmlrpc:"qty_to_invoice"`
	State               string    `xmlrpc:"state"`
	TeamId              Many2One  `xmlrpc:"team_id"`
	UserId              Many2One  `xmlrpc:"user_id"`
	Volume              float64   `xmlrpc:"volume"`
	WarehouseId         Many2One  `xmlrpc:"warehouse_id"`
	Weight              float64   `xmlrpc:"weight"`
}

type SaleReportNil struct {
	AmtInvoiced         interface{} `xmlrpc:"amt_invoiced"`
	AmtToInvoice        interface{} `xmlrpc:"amt_to_invoice"`
	AnalyticAccountId   interface{} `xmlrpc:"analytic_account_id"`
	CategId             interface{} `xmlrpc:"categ_id"`
	CommercialPartnerId interface{} `xmlrpc:"commercial_partner_id"`
	CompanyId           interface{} `xmlrpc:"company_id"`
	ConfirmationDate    interface{} `xmlrpc:"confirmation_date"`
	CountryId           interface{} `xmlrpc:"country_id"`
	Date                interface{} `xmlrpc:"date"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	Id                  interface{} `xmlrpc:"id"`
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	Name                interface{} `xmlrpc:"name"`
	Nbr                 interface{} `xmlrpc:"nbr"`
	PartnerId           interface{} `xmlrpc:"partner_id"`
	PricelistId         interface{} `xmlrpc:"pricelist_id"`
	PriceSubtotal       interface{} `xmlrpc:"price_subtotal"`
	PriceTotal          interface{} `xmlrpc:"price_total"`
	ProductId           interface{} `xmlrpc:"product_id"`
	ProductTmplId       interface{} `xmlrpc:"product_tmpl_id"`
	ProductUom          interface{} `xmlrpc:"product_uom"`
	ProductUomQty       interface{} `xmlrpc:"product_uom_qty"`
	QtyDelivered        interface{} `xmlrpc:"qty_delivered"`
	QtyInvoiced         interface{} `xmlrpc:"qty_invoiced"`
	QtyToInvoice        interface{} `xmlrpc:"qty_to_invoice"`
	State               interface{} `xmlrpc:"state"`
	TeamId              interface{} `xmlrpc:"team_id"`
	UserId              interface{} `xmlrpc:"user_id"`
	Volume              interface{} `xmlrpc:"volume"`
	WarehouseId         interface{} `xmlrpc:"warehouse_id"`
	Weight              interface{} `xmlrpc:"weight"`
}

var SaleReportModel string = "sale.report"

type SaleReports []SaleReport

type SaleReportsNil []SaleReportNil

func (s *SaleReport) NilableType_() interface{} {
	return &SaleReportNil{}
}

func (ns *SaleReportNil) Type_() interface{} {
	s := &SaleReport{}
	return load(ns, s)
}

func (s *SaleReports) NilableType_() interface{} {
	return &SaleReportsNil{}
}

func (ns *SaleReportsNil) Type_() interface{} {
	s := &SaleReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleReport))
	}
	return s
}
