package types

import (
	"time"
)

type AccountChartTemplate struct {
	LastUpdate                        time.Time `xmlrpc:"__last_update"`
	AccountIds                        []int64   `xmlrpc:"account_ids"`
	BankAccountCodePrefix             string    `xmlrpc:"bank_account_code_prefix"`
	CashAccountCodePrefix             string    `xmlrpc:"cash_account_code_prefix"`
	CodeDigits                        int64     `xmlrpc:"code_digits"`
	CompanyId                         Many2One  `xmlrpc:"company_id"`
	CompleteTaxSet                    bool      `xmlrpc:"complete_tax_set"`
	CreateDate                        time.Time `xmlrpc:"create_date"`
	CreateUid                         Many2One  `xmlrpc:"create_uid"`
	CurrencyId                        Many2One  `xmlrpc:"currency_id"`
	DisplayName                       string    `xmlrpc:"display_name"`
	ExpenseCurrencyExchangeAccountId  Many2One  `xmlrpc:"expense_currency_exchange_account_id"`
	Id                                int64     `xmlrpc:"id"`
	IncomeCurrencyExchangeAccountId   Many2One  `xmlrpc:"income_currency_exchange_account_id"`
	Name                              string    `xmlrpc:"name"`
	ParentId                          Many2One  `xmlrpc:"parent_id"`
	PropertyAccountExpenseCategId     Many2One  `xmlrpc:"property_account_expense_categ_id"`
	PropertyAccountExpenseId          Many2One  `xmlrpc:"property_account_expense_id"`
	PropertyAccountIncomeCategId      Many2One  `xmlrpc:"property_account_income_categ_id"`
	PropertyAccountIncomeId           Many2One  `xmlrpc:"property_account_income_id"`
	PropertyAccountPayableId          Many2One  `xmlrpc:"property_account_payable_id"`
	PropertyAccountReceivableId       Many2One  `xmlrpc:"property_account_receivable_id"`
	PropertyStockAccountInputCategId  Many2One  `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId Many2One  `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockValuationAccountId   Many2One  `xmlrpc:"property_stock_valuation_account_id"`
	TaxTemplateIds                    []int64   `xmlrpc:"tax_template_ids"`
	TransferAccountId                 Many2One  `xmlrpc:"transfer_account_id"`
	UseAngloSaxon                     bool      `xmlrpc:"use_anglo_saxon"`
	Visible                           bool      `xmlrpc:"visible"`
	WriteDate                         time.Time `xmlrpc:"write_date"`
	WriteUid                          Many2One  `xmlrpc:"write_uid"`
}

type AccountChartTemplateNil struct {
	LastUpdate                        interface{} `xmlrpc:"__last_update"`
	AccountIds                        interface{} `xmlrpc:"account_ids"`
	BankAccountCodePrefix             interface{} `xmlrpc:"bank_account_code_prefix"`
	CashAccountCodePrefix             interface{} `xmlrpc:"cash_account_code_prefix"`
	CodeDigits                        interface{} `xmlrpc:"code_digits"`
	CompanyId                         interface{} `xmlrpc:"company_id"`
	CompleteTaxSet                    bool        `xmlrpc:"complete_tax_set"`
	CreateDate                        interface{} `xmlrpc:"create_date"`
	CreateUid                         interface{} `xmlrpc:"create_uid"`
	CurrencyId                        interface{} `xmlrpc:"currency_id"`
	DisplayName                       interface{} `xmlrpc:"display_name"`
	ExpenseCurrencyExchangeAccountId  interface{} `xmlrpc:"expense_currency_exchange_account_id"`
	Id                                interface{} `xmlrpc:"id"`
	IncomeCurrencyExchangeAccountId   interface{} `xmlrpc:"income_currency_exchange_account_id"`
	Name                              interface{} `xmlrpc:"name"`
	ParentId                          interface{} `xmlrpc:"parent_id"`
	PropertyAccountExpenseCategId     interface{} `xmlrpc:"property_account_expense_categ_id"`
	PropertyAccountExpenseId          interface{} `xmlrpc:"property_account_expense_id"`
	PropertyAccountIncomeCategId      interface{} `xmlrpc:"property_account_income_categ_id"`
	PropertyAccountIncomeId           interface{} `xmlrpc:"property_account_income_id"`
	PropertyAccountPayableId          interface{} `xmlrpc:"property_account_payable_id"`
	PropertyAccountReceivableId       interface{} `xmlrpc:"property_account_receivable_id"`
	PropertyStockAccountInputCategId  interface{} `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId interface{} `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockValuationAccountId   interface{} `xmlrpc:"property_stock_valuation_account_id"`
	TaxTemplateIds                    interface{} `xmlrpc:"tax_template_ids"`
	TransferAccountId                 interface{} `xmlrpc:"transfer_account_id"`
	UseAngloSaxon                     bool        `xmlrpc:"use_anglo_saxon"`
	Visible                           bool        `xmlrpc:"visible"`
	WriteDate                         interface{} `xmlrpc:"write_date"`
	WriteUid                          interface{} `xmlrpc:"write_uid"`
}

var AccountChartTemplateModel string = "account.chart.template"

type AccountChartTemplates []AccountChartTemplate

type AccountChartTemplatesNil []AccountChartTemplateNil

func (s *AccountChartTemplate) NilableType_() interface{} {
	return &AccountChartTemplateNil{}
}

func (ns *AccountChartTemplateNil) Type_() interface{} {
	s := &AccountChartTemplate{}
	return load(ns, s)
}

func (s *AccountChartTemplates) NilableType_() interface{} {
	return &AccountChartTemplatesNil{}
}

func (ns *AccountChartTemplatesNil) Type_() interface{} {
	s := &AccountChartTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountChartTemplate))
	}
	return s
}
