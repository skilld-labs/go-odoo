package types

import (
	"time"
)

type PurchaseOrderLine struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AccountAnalyticId Many2One  `xmlrpc:"account_analytic_id"`
	AnalyticTagIds    []int64   `xmlrpc:"analytic_tag_ids"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	DateOrder         time.Time `xmlrpc:"date_order"`
	DatePlanned       time.Time `xmlrpc:"date_planned"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	InvoiceLines      []int64   `xmlrpc:"invoice_lines"`
	MoveIds           []int64   `xmlrpc:"move_ids"`
	Name              string    `xmlrpc:"name"`
	OrderId           Many2One  `xmlrpc:"order_id"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	PriceSubtotal     float64   `xmlrpc:"price_subtotal"`
	PriceTax          float64   `xmlrpc:"price_tax"`
	PriceTotal        float64   `xmlrpc:"price_total"`
	PriceUnit         float64   `xmlrpc:"price_unit"`
	ProcurementIds    []int64   `xmlrpc:"procurement_ids"`
	ProductId         Many2One  `xmlrpc:"product_id"`
	ProductQty        float64   `xmlrpc:"product_qty"`
	ProductUom        Many2One  `xmlrpc:"product_uom"`
	QtyInvoiced       float64   `xmlrpc:"qty_invoiced"`
	QtyReceived       float64   `xmlrpc:"qty_received"`
	Sequence          int64     `xmlrpc:"sequence"`
	State             string    `xmlrpc:"state"`
	TaxesId           []int64   `xmlrpc:"taxes_id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type PurchaseOrderLineNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AccountAnalyticId interface{} `xmlrpc:"account_analytic_id"`
	AnalyticTagIds    interface{} `xmlrpc:"analytic_tag_ids"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	DateOrder         interface{} `xmlrpc:"date_order"`
	DatePlanned       interface{} `xmlrpc:"date_planned"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	InvoiceLines      interface{} `xmlrpc:"invoice_lines"`
	MoveIds           interface{} `xmlrpc:"move_ids"`
	Name              interface{} `xmlrpc:"name"`
	OrderId           interface{} `xmlrpc:"order_id"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	PriceSubtotal     interface{} `xmlrpc:"price_subtotal"`
	PriceTax          interface{} `xmlrpc:"price_tax"`
	PriceTotal        interface{} `xmlrpc:"price_total"`
	PriceUnit         interface{} `xmlrpc:"price_unit"`
	ProcurementIds    interface{} `xmlrpc:"procurement_ids"`
	ProductId         interface{} `xmlrpc:"product_id"`
	ProductQty        interface{} `xmlrpc:"product_qty"`
	ProductUom        interface{} `xmlrpc:"product_uom"`
	QtyInvoiced       interface{} `xmlrpc:"qty_invoiced"`
	QtyReceived       interface{} `xmlrpc:"qty_received"`
	Sequence          interface{} `xmlrpc:"sequence"`
	State             interface{} `xmlrpc:"state"`
	TaxesId           interface{} `xmlrpc:"taxes_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var PurchaseOrderLineModel string = "purchase.order.line"

type PurchaseOrderLines []PurchaseOrderLine

type PurchaseOrderLinesNil []PurchaseOrderLineNil

func (s *PurchaseOrderLine) NilableType_() interface{} {
	return &PurchaseOrderLineNil{}
}

func (ns *PurchaseOrderLineNil) Type_() interface{} {
	s := &PurchaseOrderLine{}
	return load(ns, s)
}

func (s *PurchaseOrderLines) NilableType_() interface{} {
	return &PurchaseOrderLinesNil{}
}

func (ns *PurchaseOrderLinesNil) Type_() interface{} {
	s := &PurchaseOrderLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PurchaseOrderLine))
	}
	return s
}
