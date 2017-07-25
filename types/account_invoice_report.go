package types

import (
	"time"
)

type AccountInvoiceReport struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AccountId                Many2One  `xmlrpc:"account_id"`
	AccountLineId            Many2One  `xmlrpc:"account_line_id"`
	CategId                  Many2One  `xmlrpc:"categ_id"`
	CommercialPartnerId      Many2One  `xmlrpc:"commercial_partner_id"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CountryId                Many2One  `xmlrpc:"country_id"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	CurrencyRate             float64   `xmlrpc:"currency_rate"`
	Date                     time.Time `xmlrpc:"date"`
	DateDue                  time.Time `xmlrpc:"date_due"`
	DisplayName              string    `xmlrpc:"display_name"`
	FiscalPositionId         Many2One  `xmlrpc:"fiscal_position_id"`
	Id                       int64     `xmlrpc:"id"`
	JournalId                Many2One  `xmlrpc:"journal_id"`
	Nbr                      int64     `xmlrpc:"nbr"`
	PartnerBankId            Many2One  `xmlrpc:"partner_bank_id"`
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PaymentTermId            Many2One  `xmlrpc:"payment_term_id"`
	PriceAverage             float64   `xmlrpc:"price_average"`
	PriceTotal               float64   `xmlrpc:"price_total"`
	ProductId                Many2One  `xmlrpc:"product_id"`
	ProductQty               float64   `xmlrpc:"product_qty"`
	Residual                 float64   `xmlrpc:"residual"`
	State                    string    `xmlrpc:"state"`
	TeamId                   Many2One  `xmlrpc:"team_id"`
	Type                     string    `xmlrpc:"type"`
	UomName                  string    `xmlrpc:"uom_name"`
	UserCurrencyPriceAverage float64   `xmlrpc:"user_currency_price_average"`
	UserCurrencyPriceTotal   float64   `xmlrpc:"user_currency_price_total"`
	UserCurrencyResidual     float64   `xmlrpc:"user_currency_residual"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	Volume                   float64   `xmlrpc:"volume"`
	Weight                   float64   `xmlrpc:"weight"`
}

type AccountInvoiceReportNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AccountId                interface{} `xmlrpc:"account_id"`
	AccountLineId            interface{} `xmlrpc:"account_line_id"`
	CategId                  interface{} `xmlrpc:"categ_id"`
	CommercialPartnerId      interface{} `xmlrpc:"commercial_partner_id"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CountryId                interface{} `xmlrpc:"country_id"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	CurrencyRate             interface{} `xmlrpc:"currency_rate"`
	Date                     interface{} `xmlrpc:"date"`
	DateDue                  interface{} `xmlrpc:"date_due"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	FiscalPositionId         interface{} `xmlrpc:"fiscal_position_id"`
	Id                       interface{} `xmlrpc:"id"`
	JournalId                interface{} `xmlrpc:"journal_id"`
	Nbr                      interface{} `xmlrpc:"nbr"`
	PartnerBankId            interface{} `xmlrpc:"partner_bank_id"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PaymentTermId            interface{} `xmlrpc:"payment_term_id"`
	PriceAverage             interface{} `xmlrpc:"price_average"`
	PriceTotal               interface{} `xmlrpc:"price_total"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProductQty               interface{} `xmlrpc:"product_qty"`
	Residual                 interface{} `xmlrpc:"residual"`
	State                    interface{} `xmlrpc:"state"`
	TeamId                   interface{} `xmlrpc:"team_id"`
	Type                     interface{} `xmlrpc:"type"`
	UomName                  interface{} `xmlrpc:"uom_name"`
	UserCurrencyPriceAverage interface{} `xmlrpc:"user_currency_price_average"`
	UserCurrencyPriceTotal   interface{} `xmlrpc:"user_currency_price_total"`
	UserCurrencyResidual     interface{} `xmlrpc:"user_currency_residual"`
	UserId                   interface{} `xmlrpc:"user_id"`
	Volume                   interface{} `xmlrpc:"volume"`
	Weight                   interface{} `xmlrpc:"weight"`
}

var AccountInvoiceReportModel string = "account.invoice.report"

type AccountInvoiceReports []AccountInvoiceReport

type AccountInvoiceReportsNil []AccountInvoiceReportNil

func (s *AccountInvoiceReport) NilableType_() interface{} {
	return &AccountInvoiceReportNil{}
}

func (ns *AccountInvoiceReportNil) Type_() interface{} {
	s := &AccountInvoiceReport{}
	return load(ns, s)
}

func (s *AccountInvoiceReports) NilableType_() interface{} {
	return &AccountInvoiceReportsNil{}
}

func (ns *AccountInvoiceReportsNil) Type_() interface{} {
	s := &AccountInvoiceReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceReport))
	}
	return s
}
