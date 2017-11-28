package types

import (
	"time"
)

type SaleOrderLine struct {
	AmtInvoiced            float64     `xmlrpc:"amt_invoiced"`
	AmtToInvoice           float64     `xmlrpc:"amt_to_invoice"`
	AnalyticTagIds         []int64     `xmlrpc:"analytic_tag_ids"`
	AutosalesBaseOrderLine Many2One    `xmlrpc:"autosales_base_order_line"`
	AutosalesLine          bool        `xmlrpc:"autosales_line"`
	CompanyId              Many2One    `xmlrpc:"company_id"`
	CreateDate             time.Time   `xmlrpc:"create_date"`
	CreateUid              Many2One    `xmlrpc:"create_uid"`
	CurrencyId             Many2One    `xmlrpc:"currency_id"`
	CustomerLead           float64     `xmlrpc:"customer_lead"`
	Discount               float64     `xmlrpc:"discount"`
	DisplayName            string      `xmlrpc:"display_name"`
	Id                     int64       `xmlrpc:"id"`
	InvoiceLines           []int64     `xmlrpc:"invoice_lines"`
	InvoiceStatus          interface{} `xmlrpc:"invoice_status"`
	IsDownpayment          bool        `xmlrpc:"is_downpayment"`
	IsService              bool        `xmlrpc:"is_service"`
	LastUpdate             time.Time   `xmlrpc:"__last_update"`
	LayoutCategoryId       Many2One    `xmlrpc:"layout_category_id"`
	LayoutCategorySequence int64       `xmlrpc:"layout_category_sequence"`
	MoveIds                []int64     `xmlrpc:"move_ids"`
	Name                   string      `xmlrpc:"name"`
	OrderId                Many2One    `xmlrpc:"order_id"`
	OrderPartnerId         Many2One    `xmlrpc:"order_partner_id"`
	PriceReduce            float64     `xmlrpc:"price_reduce"`
	PriceReduceTaxexcl     float64     `xmlrpc:"price_reduce_taxexcl"`
	PriceReduceTaxinc      float64     `xmlrpc:"price_reduce_taxinc"`
	PriceSubtotal          float64     `xmlrpc:"price_subtotal"`
	PriceTax               float64     `xmlrpc:"price_tax"`
	PriceTotal             float64     `xmlrpc:"price_total"`
	PriceUnit              float64     `xmlrpc:"price_unit"`
	ProductId              Many2One    `xmlrpc:"product_id"`
	ProductImage           string      `xmlrpc:"product_image"`
	ProductPackaging       Many2One    `xmlrpc:"product_packaging"`
	ProductUom             Many2One    `xmlrpc:"product_uom"`
	ProductUomQty          float64     `xmlrpc:"product_uom_qty"`
	QtyDelivered           float64     `xmlrpc:"qty_delivered"`
	QtyDeliveredUpdateable bool        `xmlrpc:"qty_delivered_updateable"`
	QtyInvoiced            float64     `xmlrpc:"qty_invoiced"`
	QtyToInvoice           float64     `xmlrpc:"qty_to_invoice"`
	RouteId                Many2One    `xmlrpc:"route_id"`
	SalesmanId             Many2One    `xmlrpc:"salesman_id"`
	Sequence               int64       `xmlrpc:"sequence"`
	State                  interface{} `xmlrpc:"state"`
	TaskId                 Many2One    `xmlrpc:"task_id"`
	TaxId                  []int64     `xmlrpc:"tax_id"`
	WriteDate              time.Time   `xmlrpc:"write_date"`
	WriteUid               Many2One    `xmlrpc:"write_uid"`
}

type SaleOrderLineNil struct {
	AmtInvoiced            interface{} `xmlrpc:"amt_invoiced"`
	AmtToInvoice           interface{} `xmlrpc:"amt_to_invoice"`
	AnalyticTagIds         interface{} `xmlrpc:"analytic_tag_ids"`
	AutosalesBaseOrderLine interface{} `xmlrpc:"autosales_base_order_line"`
	AutosalesLine          bool        `xmlrpc:"autosales_line"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	CustomerLead           interface{} `xmlrpc:"customer_lead"`
	Discount               interface{} `xmlrpc:"discount"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Id                     interface{} `xmlrpc:"id"`
	InvoiceLines           interface{} `xmlrpc:"invoice_lines"`
	InvoiceStatus          interface{} `xmlrpc:"invoice_status"`
	IsDownpayment          bool        `xmlrpc:"is_downpayment"`
	IsService              bool        `xmlrpc:"is_service"`
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	LayoutCategoryId       interface{} `xmlrpc:"layout_category_id"`
	LayoutCategorySequence interface{} `xmlrpc:"layout_category_sequence"`
	MoveIds                interface{} `xmlrpc:"move_ids"`
	Name                   interface{} `xmlrpc:"name"`
	OrderId                interface{} `xmlrpc:"order_id"`
	OrderPartnerId         interface{} `xmlrpc:"order_partner_id"`
	PriceReduce            interface{} `xmlrpc:"price_reduce"`
	PriceReduceTaxexcl     interface{} `xmlrpc:"price_reduce_taxexcl"`
	PriceReduceTaxinc      interface{} `xmlrpc:"price_reduce_taxinc"`
	PriceSubtotal          interface{} `xmlrpc:"price_subtotal"`
	PriceTax               interface{} `xmlrpc:"price_tax"`
	PriceTotal             interface{} `xmlrpc:"price_total"`
	PriceUnit              interface{} `xmlrpc:"price_unit"`
	ProductId              interface{} `xmlrpc:"product_id"`
	ProductImage           interface{} `xmlrpc:"product_image"`
	ProductPackaging       interface{} `xmlrpc:"product_packaging"`
	ProductUom             interface{} `xmlrpc:"product_uom"`
	ProductUomQty          interface{} `xmlrpc:"product_uom_qty"`
	QtyDelivered           interface{} `xmlrpc:"qty_delivered"`
	QtyDeliveredUpdateable bool        `xmlrpc:"qty_delivered_updateable"`
	QtyInvoiced            interface{} `xmlrpc:"qty_invoiced"`
	QtyToInvoice           interface{} `xmlrpc:"qty_to_invoice"`
	RouteId                interface{} `xmlrpc:"route_id"`
	SalesmanId             interface{} `xmlrpc:"salesman_id"`
	Sequence               interface{} `xmlrpc:"sequence"`
	State                  interface{} `xmlrpc:"state"`
	TaskId                 interface{} `xmlrpc:"task_id"`
	TaxId                  interface{} `xmlrpc:"tax_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var SaleOrderLineModel string = "sale.order.line"

type SaleOrderLines []SaleOrderLine

type SaleOrderLinesNil []SaleOrderLineNil

func (s *SaleOrderLine) NilableType_() interface{} {
	return &SaleOrderLineNil{}
}

func (ns *SaleOrderLineNil) Type_() interface{} {
	s := &SaleOrderLine{}
	return load(ns, s)
}

func (s *SaleOrderLines) NilableType_() interface{} {
	return &SaleOrderLinesNil{}
}

func (ns *SaleOrderLinesNil) Type_() interface{} {
	s := &SaleOrderLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleOrderLine))
	}
	return s
}
