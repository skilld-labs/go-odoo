package types

import (
	"time"
)

type ResCompany struct {
	LastUpdate                        time.Time `xmlrpc:"__last_update"`
	AccountNo                         string    `xmlrpc:"account_no"`
	AccountsCodeDigits                int64     `xmlrpc:"accounts_code_digits"`
	AngloSaxonAccounting              bool      `xmlrpc:"anglo_saxon_accounting"`
	Ape                               string    `xmlrpc:"ape"`
	BankAccountCodePrefix             string    `xmlrpc:"bank_account_code_prefix"`
	BankIds                           []int64   `xmlrpc:"bank_ids"`
	BankJournalIds                    []int64   `xmlrpc:"bank_journal_ids"`
	CashAccountCodePrefix             string    `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId                   Many2One  `xmlrpc:"chart_template_id"`
	ChildIds                          []int64   `xmlrpc:"child_ids"`
	City                              string    `xmlrpc:"city"`
	CompanyRegistry                   string    `xmlrpc:"company_registry"`
	CountryId                         Many2One  `xmlrpc:"country_id"`
	CreateDate                        time.Time `xmlrpc:"create_date"`
	CreateUid                         Many2One  `xmlrpc:"create_uid"`
	CurrencyExchangeJournalId         Many2One  `xmlrpc:"currency_exchange_journal_id"`
	CurrencyId                        Many2One  `xmlrpc:"currency_id"`
	CustomFooter                      bool      `xmlrpc:"custom_footer"`
	DisplayName                       string    `xmlrpc:"display_name"`
	Email                             string    `xmlrpc:"email"`
	ExpectsChartOfAccounts            bool      `xmlrpc:"expects_chart_of_accounts"`
	ExpenseCurrencyExchangeAccountId  Many2One  `xmlrpc:"expense_currency_exchange_account_id"`
	Fax                               string    `xmlrpc:"fax"`
	FiscalyearLastDay                 int64     `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth               string    `xmlrpc:"fiscalyear_last_month"`
	FiscalyearLockDate                time.Time `xmlrpc:"fiscalyear_lock_date"`
	Font                              Many2One  `xmlrpc:"font"`
	Id                                int64     `xmlrpc:"id"`
	IncomeCurrencyExchangeAccountId   Many2One  `xmlrpc:"income_currency_exchange_account_id"`
	InternalTransitLocationId         Many2One  `xmlrpc:"internal_transit_location_id"`
	Ldaps                             []int64   `xmlrpc:"ldaps"`
	Logo                              string    `xmlrpc:"logo"`
	LogoWeb                           string    `xmlrpc:"logo_web"`
	Name                              string    `xmlrpc:"name"`
	OverdueMsg                        string    `xmlrpc:"overdue_msg"`
	PaperformatId                     Many2One  `xmlrpc:"paperformat_id"`
	ParentId                          Many2One  `xmlrpc:"parent_id"`
	PartnerId                         Many2One  `xmlrpc:"partner_id"`
	PeriodLockDate                    time.Time `xmlrpc:"period_lock_date"`
	Phone                             string    `xmlrpc:"phone"`
	PoDoubleValidation                string    `xmlrpc:"po_double_validation"`
	PoDoubleValidationAmount          float64   `xmlrpc:"po_double_validation_amount"`
	PoLead                            float64   `xmlrpc:"po_lead"`
	PoLock                            string    `xmlrpc:"po_lock"`
	ProjectTimeModeId                 Many2One  `xmlrpc:"project_time_mode_id"`
	PropagationMinimumDelta           int64     `xmlrpc:"propagation_minimum_delta"`
	PropertyStockAccountInputCategId  Many2One  `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId Many2One  `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockValuationAccountId   Many2One  `xmlrpc:"property_stock_valuation_account_id"`
	RmlFooter                         string    `xmlrpc:"rml_footer"`
	RmlFooterReadonly                 string    `xmlrpc:"rml_footer_readonly"`
	RmlHeader                         string    `xmlrpc:"rml_header"`
	RmlHeader1                        string    `xmlrpc:"rml_header1"`
	RmlHeader2                        string    `xmlrpc:"rml_header2"`
	RmlHeader3                        string    `xmlrpc:"rml_header3"`
	RmlPaperFormat                    string    `xmlrpc:"rml_paper_format"`
	SaleNote                          string    `xmlrpc:"sale_note"`
	SecurityLead                      float64   `xmlrpc:"security_lead"`
	Sequence                          int64     `xmlrpc:"sequence"`
	Siret                             string    `xmlrpc:"siret"`
	StateId                           Many2One  `xmlrpc:"state_id"`
	Street                            string    `xmlrpc:"street"`
	Street2                           string    `xmlrpc:"street2"`
	TaxCalculationRoundingMethod      string    `xmlrpc:"tax_calculation_rounding_method"`
	TransferAccountId                 Many2One  `xmlrpc:"transfer_account_id"`
	UserIds                           []int64   `xmlrpc:"user_ids"`
	Vat                               string    `xmlrpc:"vat"`
	VatCheckVies                      bool      `xmlrpc:"vat_check_vies"`
	Website                           string    `xmlrpc:"website"`
	WriteDate                         time.Time `xmlrpc:"write_date"`
	WriteUid                          Many2One  `xmlrpc:"write_uid"`
	Zip                               string    `xmlrpc:"zip"`
}

type ResCompanyNil struct {
	LastUpdate                        interface{} `xmlrpc:"__last_update"`
	AccountNo                         interface{} `xmlrpc:"account_no"`
	AccountsCodeDigits                interface{} `xmlrpc:"accounts_code_digits"`
	AngloSaxonAccounting              bool        `xmlrpc:"anglo_saxon_accounting"`
	Ape                               interface{} `xmlrpc:"ape"`
	BankAccountCodePrefix             interface{} `xmlrpc:"bank_account_code_prefix"`
	BankIds                           interface{} `xmlrpc:"bank_ids"`
	BankJournalIds                    interface{} `xmlrpc:"bank_journal_ids"`
	CashAccountCodePrefix             interface{} `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId                   interface{} `xmlrpc:"chart_template_id"`
	ChildIds                          interface{} `xmlrpc:"child_ids"`
	City                              interface{} `xmlrpc:"city"`
	CompanyRegistry                   interface{} `xmlrpc:"company_registry"`
	CountryId                         interface{} `xmlrpc:"country_id"`
	CreateDate                        interface{} `xmlrpc:"create_date"`
	CreateUid                         interface{} `xmlrpc:"create_uid"`
	CurrencyExchangeJournalId         interface{} `xmlrpc:"currency_exchange_journal_id"`
	CurrencyId                        interface{} `xmlrpc:"currency_id"`
	CustomFooter                      bool        `xmlrpc:"custom_footer"`
	DisplayName                       interface{} `xmlrpc:"display_name"`
	Email                             interface{} `xmlrpc:"email"`
	ExpectsChartOfAccounts            bool        `xmlrpc:"expects_chart_of_accounts"`
	ExpenseCurrencyExchangeAccountId  interface{} `xmlrpc:"expense_currency_exchange_account_id"`
	Fax                               interface{} `xmlrpc:"fax"`
	FiscalyearLastDay                 interface{} `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth               interface{} `xmlrpc:"fiscalyear_last_month"`
	FiscalyearLockDate                interface{} `xmlrpc:"fiscalyear_lock_date"`
	Font                              interface{} `xmlrpc:"font"`
	Id                                interface{} `xmlrpc:"id"`
	IncomeCurrencyExchangeAccountId   interface{} `xmlrpc:"income_currency_exchange_account_id"`
	InternalTransitLocationId         interface{} `xmlrpc:"internal_transit_location_id"`
	Ldaps                             interface{} `xmlrpc:"ldaps"`
	Logo                              interface{} `xmlrpc:"logo"`
	LogoWeb                           interface{} `xmlrpc:"logo_web"`
	Name                              interface{} `xmlrpc:"name"`
	OverdueMsg                        interface{} `xmlrpc:"overdue_msg"`
	PaperformatId                     interface{} `xmlrpc:"paperformat_id"`
	ParentId                          interface{} `xmlrpc:"parent_id"`
	PartnerId                         interface{} `xmlrpc:"partner_id"`
	PeriodLockDate                    interface{} `xmlrpc:"period_lock_date"`
	Phone                             interface{} `xmlrpc:"phone"`
	PoDoubleValidation                interface{} `xmlrpc:"po_double_validation"`
	PoDoubleValidationAmount          interface{} `xmlrpc:"po_double_validation_amount"`
	PoLead                            interface{} `xmlrpc:"po_lead"`
	PoLock                            interface{} `xmlrpc:"po_lock"`
	ProjectTimeModeId                 interface{} `xmlrpc:"project_time_mode_id"`
	PropagationMinimumDelta           interface{} `xmlrpc:"propagation_minimum_delta"`
	PropertyStockAccountInputCategId  interface{} `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId interface{} `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockValuationAccountId   interface{} `xmlrpc:"property_stock_valuation_account_id"`
	RmlFooter                         interface{} `xmlrpc:"rml_footer"`
	RmlFooterReadonly                 interface{} `xmlrpc:"rml_footer_readonly"`
	RmlHeader                         interface{} `xmlrpc:"rml_header"`
	RmlHeader1                        interface{} `xmlrpc:"rml_header1"`
	RmlHeader2                        interface{} `xmlrpc:"rml_header2"`
	RmlHeader3                        interface{} `xmlrpc:"rml_header3"`
	RmlPaperFormat                    interface{} `xmlrpc:"rml_paper_format"`
	SaleNote                          interface{} `xmlrpc:"sale_note"`
	SecurityLead                      interface{} `xmlrpc:"security_lead"`
	Sequence                          interface{} `xmlrpc:"sequence"`
	Siret                             interface{} `xmlrpc:"siret"`
	StateId                           interface{} `xmlrpc:"state_id"`
	Street                            interface{} `xmlrpc:"street"`
	Street2                           interface{} `xmlrpc:"street2"`
	TaxCalculationRoundingMethod      interface{} `xmlrpc:"tax_calculation_rounding_method"`
	TransferAccountId                 interface{} `xmlrpc:"transfer_account_id"`
	UserIds                           interface{} `xmlrpc:"user_ids"`
	Vat                               interface{} `xmlrpc:"vat"`
	VatCheckVies                      bool        `xmlrpc:"vat_check_vies"`
	Website                           interface{} `xmlrpc:"website"`
	WriteDate                         interface{} `xmlrpc:"write_date"`
	WriteUid                          interface{} `xmlrpc:"write_uid"`
	Zip                               interface{} `xmlrpc:"zip"`
}

var ResCompanyModel string = "res.company"

type ResCompanys []ResCompany

type ResCompanysNil []ResCompanyNil

func (s *ResCompany) NilableType_() interface{} {
	return &ResCompanyNil{}
}

func (ns *ResCompanyNil) Type_() interface{} {
	s := &ResCompany{}
	return load(ns, s)
}

func (s *ResCompanys) NilableType_() interface{} {
	return &ResCompanysNil{}
}

func (ns *ResCompanysNil) Type_() interface{} {
	s := &ResCompanys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCompany))
	}
	return s
}
