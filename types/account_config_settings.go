package types

import (
	"time"
)

type AccountConfigSettings struct {
	LastUpdate                          time.Time `xmlrpc:"__last_update"`
	BankAccountCodePrefix               string    `xmlrpc:"bank_account_code_prefix"`
	CashAccountCodePrefix               string    `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId                     Many2One  `xmlrpc:"chart_template_id"`
	CodeDigits                          int64     `xmlrpc:"code_digits"`
	CompanyFooter                       string    `xmlrpc:"company_footer"`
	CompanyId                           Many2One  `xmlrpc:"company_id"`
	CompleteTaxSet                      bool      `xmlrpc:"complete_tax_set"`
	CreateDate                          time.Time `xmlrpc:"create_date"`
	CreateUid                           Many2One  `xmlrpc:"create_uid"`
	CurrencyExchangeJournalId           Many2One  `xmlrpc:"currency_exchange_journal_id"`
	CurrencyId                          Many2One  `xmlrpc:"currency_id"`
	DefaultPurchaseTaxId                Many2One  `xmlrpc:"default_purchase_tax_id"`
	DefaultSaleTaxId                    Many2One  `xmlrpc:"default_sale_tax_id"`
	DisplayName                         string    `xmlrpc:"display_name"`
	ExpectsChartOfAccounts              bool      `xmlrpc:"expects_chart_of_accounts"`
	FiscalyearLastDay                   int64     `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth                 string    `xmlrpc:"fiscalyear_last_month"`
	FiscalyearLockDate                  time.Time `xmlrpc:"fiscalyear_lock_date"`
	GroupAnalyticAccountForPurchases    bool      `xmlrpc:"group_analytic_account_for_purchases"`
	GroupAnalyticAccountForSales        bool      `xmlrpc:"group_analytic_account_for_sales"`
	GroupAnalyticAccounting             bool      `xmlrpc:"group_analytic_accounting"`
	GroupMultiCurrency                  bool      `xmlrpc:"group_multi_currency"`
	GroupProformaInvoices               bool      `xmlrpc:"group_proforma_invoices"`
	GroupWarningAccount                 string    `xmlrpc:"group_warning_account"`
	HasChartOfAccounts                  bool      `xmlrpc:"has_chart_of_accounts"`
	HasDefaultCompany                   bool      `xmlrpc:"has_default_company"`
	Id                                  int64     `xmlrpc:"id"`
	ModuleAccountAccountant             bool      `xmlrpc:"module_account_accountant"`
	ModuleAccountAsset                  bool      `xmlrpc:"module_account_asset"`
	ModuleAccountBankStatementImportCsv bool      `xmlrpc:"module_account_bank_statement_import_csv"`
	ModuleAccountBankStatementImportOfx bool      `xmlrpc:"module_account_bank_statement_import_ofx"`
	ModuleAccountBankStatementImportQif bool      `xmlrpc:"module_account_bank_statement_import_qif"`
	ModuleAccountBatchDeposit           bool      `xmlrpc:"module_account_batch_deposit"`
	ModuleAccountBudget                 bool      `xmlrpc:"module_account_budget"`
	ModuleAccountDeferredRevenue        bool      `xmlrpc:"module_account_deferred_revenue"`
	ModuleAccountPlaid                  bool      `xmlrpc:"module_account_plaid"`
	ModuleAccountReports                bool      `xmlrpc:"module_account_reports"`
	ModuleAccountReportsFollowup        bool      `xmlrpc:"module_account_reports_followup"`
	ModuleAccountSepa                   bool      `xmlrpc:"module_account_sepa"`
	ModuleAccountTaxCashBasis           bool      `xmlrpc:"module_account_tax_cash_basis"`
	ModuleAccountYodlee                 bool      `xmlrpc:"module_account_yodlee"`
	ModuleL10nUsCheckPrinting           bool      `xmlrpc:"module_l10n_us_check_printing"`
	OverdueMsg                          string    `xmlrpc:"overdue_msg"`
	PeriodLockDate                      time.Time `xmlrpc:"period_lock_date"`
	PurchaseTaxId                       Many2One  `xmlrpc:"purchase_tax_id"`
	PurchaseTaxRate                     float64   `xmlrpc:"purchase_tax_rate"`
	SaleTaxId                           Many2One  `xmlrpc:"sale_tax_id"`
	SaleTaxRate                         float64   `xmlrpc:"sale_tax_rate"`
	TaxCalculationRoundingMethod        string    `xmlrpc:"tax_calculation_rounding_method"`
	TemplateTransferAccountId           Many2One  `xmlrpc:"template_transfer_account_id"`
	TransferAccountId                   Many2One  `xmlrpc:"transfer_account_id"`
	UseAngloSaxon                       bool      `xmlrpc:"use_anglo_saxon"`
	WriteDate                           time.Time `xmlrpc:"write_date"`
	WriteUid                            Many2One  `xmlrpc:"write_uid"`
}

type AccountConfigSettingsNil struct {
	LastUpdate                          interface{} `xmlrpc:"__last_update"`
	BankAccountCodePrefix               interface{} `xmlrpc:"bank_account_code_prefix"`
	CashAccountCodePrefix               interface{} `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId                     interface{} `xmlrpc:"chart_template_id"`
	CodeDigits                          interface{} `xmlrpc:"code_digits"`
	CompanyFooter                       interface{} `xmlrpc:"company_footer"`
	CompanyId                           interface{} `xmlrpc:"company_id"`
	CompleteTaxSet                      bool        `xmlrpc:"complete_tax_set"`
	CreateDate                          interface{} `xmlrpc:"create_date"`
	CreateUid                           interface{} `xmlrpc:"create_uid"`
	CurrencyExchangeJournalId           interface{} `xmlrpc:"currency_exchange_journal_id"`
	CurrencyId                          interface{} `xmlrpc:"currency_id"`
	DefaultPurchaseTaxId                interface{} `xmlrpc:"default_purchase_tax_id"`
	DefaultSaleTaxId                    interface{} `xmlrpc:"default_sale_tax_id"`
	DisplayName                         interface{} `xmlrpc:"display_name"`
	ExpectsChartOfAccounts              bool        `xmlrpc:"expects_chart_of_accounts"`
	FiscalyearLastDay                   interface{} `xmlrpc:"fiscalyear_last_day"`
	FiscalyearLastMonth                 interface{} `xmlrpc:"fiscalyear_last_month"`
	FiscalyearLockDate                  interface{} `xmlrpc:"fiscalyear_lock_date"`
	GroupAnalyticAccountForPurchases    bool        `xmlrpc:"group_analytic_account_for_purchases"`
	GroupAnalyticAccountForSales        bool        `xmlrpc:"group_analytic_account_for_sales"`
	GroupAnalyticAccounting             bool        `xmlrpc:"group_analytic_accounting"`
	GroupMultiCurrency                  bool        `xmlrpc:"group_multi_currency"`
	GroupProformaInvoices               bool        `xmlrpc:"group_proforma_invoices"`
	GroupWarningAccount                 interface{} `xmlrpc:"group_warning_account"`
	HasChartOfAccounts                  bool        `xmlrpc:"has_chart_of_accounts"`
	HasDefaultCompany                   bool        `xmlrpc:"has_default_company"`
	Id                                  interface{} `xmlrpc:"id"`
	ModuleAccountAccountant             bool        `xmlrpc:"module_account_accountant"`
	ModuleAccountAsset                  bool        `xmlrpc:"module_account_asset"`
	ModuleAccountBankStatementImportCsv bool        `xmlrpc:"module_account_bank_statement_import_csv"`
	ModuleAccountBankStatementImportOfx bool        `xmlrpc:"module_account_bank_statement_import_ofx"`
	ModuleAccountBankStatementImportQif bool        `xmlrpc:"module_account_bank_statement_import_qif"`
	ModuleAccountBatchDeposit           bool        `xmlrpc:"module_account_batch_deposit"`
	ModuleAccountBudget                 bool        `xmlrpc:"module_account_budget"`
	ModuleAccountDeferredRevenue        bool        `xmlrpc:"module_account_deferred_revenue"`
	ModuleAccountPlaid                  bool        `xmlrpc:"module_account_plaid"`
	ModuleAccountReports                bool        `xmlrpc:"module_account_reports"`
	ModuleAccountReportsFollowup        bool        `xmlrpc:"module_account_reports_followup"`
	ModuleAccountSepa                   bool        `xmlrpc:"module_account_sepa"`
	ModuleAccountTaxCashBasis           bool        `xmlrpc:"module_account_tax_cash_basis"`
	ModuleAccountYodlee                 bool        `xmlrpc:"module_account_yodlee"`
	ModuleL10nUsCheckPrinting           bool        `xmlrpc:"module_l10n_us_check_printing"`
	OverdueMsg                          interface{} `xmlrpc:"overdue_msg"`
	PeriodLockDate                      interface{} `xmlrpc:"period_lock_date"`
	PurchaseTaxId                       interface{} `xmlrpc:"purchase_tax_id"`
	PurchaseTaxRate                     interface{} `xmlrpc:"purchase_tax_rate"`
	SaleTaxId                           interface{} `xmlrpc:"sale_tax_id"`
	SaleTaxRate                         interface{} `xmlrpc:"sale_tax_rate"`
	TaxCalculationRoundingMethod        interface{} `xmlrpc:"tax_calculation_rounding_method"`
	TemplateTransferAccountId           interface{} `xmlrpc:"template_transfer_account_id"`
	TransferAccountId                   interface{} `xmlrpc:"transfer_account_id"`
	UseAngloSaxon                       bool        `xmlrpc:"use_anglo_saxon"`
	WriteDate                           interface{} `xmlrpc:"write_date"`
	WriteUid                            interface{} `xmlrpc:"write_uid"`
}

var AccountConfigSettingsModel string = "account.config.settings"

type AccountConfigSettingss []AccountConfigSettings

type AccountConfigSettingssNil []AccountConfigSettingsNil

func (s *AccountConfigSettings) NilableType_() interface{} {
	return &AccountConfigSettingsNil{}
}

func (ns *AccountConfigSettingsNil) Type_() interface{} {
	s := &AccountConfigSettings{}
	return load(ns, s)
}

func (s *AccountConfigSettingss) NilableType_() interface{} {
	return &AccountConfigSettingssNil{}
}

func (ns *AccountConfigSettingssNil) Type_() interface{} {
	s := &AccountConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountConfigSettings))
	}
	return s
}
