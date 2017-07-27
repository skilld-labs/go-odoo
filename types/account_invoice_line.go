package types

import (
	"time"
)

type AccountInvoiceLine struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	AccountAnalyticId      Many2One  `xmlrpc:"account_analytic_id"`
	AccountId              Many2One  `xmlrpc:"account_id"`
	AnalyticTagIds         []int64   `xmlrpc:"analytic_tag_ids"`
	CompanyCurrencyId      Many2One  `xmlrpc:"company_currency_id"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	CurrencyId             Many2One  `xmlrpc:"currency_id"`
	Discount               float64   `xmlrpc:"discount"`
	DisplayName            string    `xmlrpc:"display_name"`
	Id                     int64     `xmlrpc:"id"`
	InvoiceId              Many2One  `xmlrpc:"invoice_id"`
	InvoiceLineTaxIds      []int64   `xmlrpc:"invoice_line_tax_ids"`
	LayoutCategoryId       Many2One  `xmlrpc:"layout_category_id"`
	LayoutCategorySequence int64     `xmlrpc:"layout_category_sequence"`
	Name                   string    `xmlrpc:"name"`
	Origin                 string    `xmlrpc:"origin"`
	PartnerId              Many2One  `xmlrpc:"partner_id"`
	PriceSubtotal          float64   `xmlrpc:"price_subtotal"`
	PriceSubtotalSigned    float64   `xmlrpc:"price_subtotal_signed"`
	PriceUnit              float64   `xmlrpc:"price_unit"`
	ProductId              Many2One  `xmlrpc:"product_id"`
	PurchaseId             Many2One  `xmlrpc:"purchase_id"`
	PurchaseLineId         Many2One  `xmlrpc:"purchase_line_id"`
	Quantity               float64   `xmlrpc:"quantity"`
	SaleLineIds            []int64   `xmlrpc:"sale_line_ids"`
	Sequence               int64     `xmlrpc:"sequence"`
	UomId                  Many2One  `xmlrpc:"uom_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type AccountInvoiceLineNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	AccountAnalyticId      interface{} `xmlrpc:"account_analytic_id"`
	AccountId              interface{} `xmlrpc:"account_id"`
	AnalyticTagIds         interface{} `xmlrpc:"analytic_tag_ids"`
	CompanyCurrencyId      interface{} `xmlrpc:"company_currency_id"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	Discount               interface{} `xmlrpc:"discount"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Id                     interface{} `xmlrpc:"id"`
	InvoiceId              interface{} `xmlrpc:"invoice_id"`
	InvoiceLineTaxIds      interface{} `xmlrpc:"invoice_line_tax_ids"`
	LayoutCategoryId       interface{} `xmlrpc:"layout_category_id"`
	LayoutCategorySequence interface{} `xmlrpc:"layout_category_sequence"`
	Name                   interface{} `xmlrpc:"name"`
	Origin                 interface{} `xmlrpc:"origin"`
	PartnerId              interface{} `xmlrpc:"partner_id"`
	PriceSubtotal          interface{} `xmlrpc:"price_subtotal"`
	PriceSubtotalSigned    interface{} `xmlrpc:"price_subtotal_signed"`
	PriceUnit              interface{} `xmlrpc:"price_unit"`
	ProductId              interface{} `xmlrpc:"product_id"`
	PurchaseId             interface{} `xmlrpc:"purchase_id"`
	PurchaseLineId         interface{} `xmlrpc:"purchase_line_id"`
	Quantity               interface{} `xmlrpc:"quantity"`
	SaleLineIds            interface{} `xmlrpc:"sale_line_ids"`
	Sequence               interface{} `xmlrpc:"sequence"`
	UomId                  interface{} `xmlrpc:"uom_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceLineModel string = "account.invoice.line"

type AccountInvoiceLines []AccountInvoiceLine

type AccountInvoiceLinesNil []AccountInvoiceLineNil

func (s *AccountInvoiceLine) NilableType_() interface{} {
	return &AccountInvoiceLineNil{}
}

func (ns *AccountInvoiceLineNil) Type_() interface{} {
	s := &AccountInvoiceLine{}
	return load(ns, s)
}

func (s *AccountInvoiceLines) NilableType_() interface{} {
	return &AccountInvoiceLinesNil{}
}

func (ns *AccountInvoiceLinesNil) Type_() interface{} {
	s := &AccountInvoiceLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceLine))
	}
	return s
}
