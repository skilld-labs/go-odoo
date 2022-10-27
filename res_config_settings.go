package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccountCashBasisBaseAccountId         *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountCheckPrintingDateLabel         *Bool      `xmlrpc:"account_check_printing_date_label,omitempty"`
	AccountCheckPrintingLayout            *Selection `xmlrpc:"account_check_printing_layout,omitempty"`
	AccountCheckPrintingMarginLeft        *Float     `xmlrpc:"account_check_printing_margin_left,omitempty"`
	AccountCheckPrintingMarginRight       *Float     `xmlrpc:"account_check_printing_margin_right,omitempty"`
	AccountCheckPrintingMarginTop         *Float     `xmlrpc:"account_check_printing_margin_top,omitempty"`
	AccountCheckPrintingMultiStub         *Bool      `xmlrpc:"account_check_printing_multi_stub,omitempty"`
	ActiveUserCount                       *Int       `xmlrpc:"active_user_count,omitempty"`
	AliasDomain                           *String    `xmlrpc:"alias_domain,omitempty"`
	AuthSignupResetPassword               *Bool      `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId              *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                   *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                      *Bool      `xmlrpc:"automatic_invoice,omitempty"`
	ChartTemplateId                       *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CompanyCount                          *Int       `xmlrpc:"company_count,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyInformations                   *String    `xmlrpc:"company_informations,omitempty"`
	CompanyName                           *String    `xmlrpc:"company_name,omitempty"`
	CompanySoTemplateId                   *Many2One  `xmlrpc:"company_so_template_id,omitempty"`
	ConfirmationTemplateId                *Many2One  `xmlrpc:"confirmation_template_id,omitempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId             *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultInvoicePolicy                  *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DepositDefaultProductId               *Many2One  `xmlrpc:"deposit_default_product_id,omitempty"`
	DigestEmails                          *Bool      `xmlrpc:"digest_emails,omitempty"`
	DigestId                              *Many2One  `xmlrpc:"digest_id,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	ExpenseCurrencyExchangeAccountId      *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalEmailServerDefault            *Bool      `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	FailCounter                           *Int       `xmlrpc:"fail_counter,omitempty"`
	GroupAnalyticAccounting               *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAnalyticTags                     *Bool      `xmlrpc:"group_analytic_tags,omitempty"`
	GroupAutoDoneSetting                  *Bool      `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupCashRounding                     *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupMultiCurrency                    *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupProductPricelist                 *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                   *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                    *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupSaleDeliveryAddress              *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                *Bool      `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSalePricelist                    *Bool      `xmlrpc:"group_sale_pricelist,omitempty"`
	GroupShowLineSubtotalsTaxExcluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omitempty"`
	GroupShowLineSubtotalsTaxIncluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omitempty"`
	GroupShowPurchaseReceipts             *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                 *Bool      `xmlrpc:"group_show_sale_receipts,omitempty"`
	GroupStockPackaging                   *Bool      `xmlrpc:"group_stock_packaging,omitempty"`
	GroupUom                              *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupWarningAccount                   *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningSale                      *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	HasAccountingEntries                  *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                    *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId       *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                            *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InvoiceIsEmail                        *Bool      `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsPrint                        *Bool      `xmlrpc:"invoice_is_print,omitempty"`
	InvoiceIsSnailmail                    *Bool      `xmlrpc:"invoice_is_snailmail,omitempty"`
	InvoiceTerms                          *String    `xmlrpc:"invoice_terms,omitempty"`
	LanguageCount                         *Int       `xmlrpc:"language_count,omitempty"`
	ModuleAccountAccountant               *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementImportCamt  *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv   *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx   *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif   *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment             *Bool      `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                   *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting            *Bool      `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountInterCompanyRules        *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                *Bool      `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract           *Bool      `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountPayment                  *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPlaid                    *Bool      `xmlrpc:"module_account_plaid,omitempty"`
	ModuleAccountReports                  *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepa                     *Bool      `xmlrpc:"module_account_sepa,omitempty"`
	ModuleAccountSepaDirectDebit          *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAccountTaxcloud                 *Bool      `xmlrpc:"module_account_taxcloud,omitempty"`
	ModuleAccountYodlee                   *Bool      `xmlrpc:"module_account_yodlee,omitempty"`
	ModuleAuthLdap                        *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                       *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGengo                       *Bool      `xmlrpc:"module_base_gengo,omitempty"`
	ModuleBaseGeolocalize                 *Bool      `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                      *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCurrencyRateLive                *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                        *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                   *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                     *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                *Bool      `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                   *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliveryUps                     *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                    *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                  *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleDrive                     *Bool      `xmlrpc:"module_google_drive,omitempty"`
	ModuleGoogleRecaptcha                 *Bool      `xmlrpc:"module_google_recaptcha,omitempty"`
	ModuleGoogleSpreadsheet               *Bool      `xmlrpc:"module_google_spreadsheet,omitempty"`
	ModuleL10NEuService                   *Bool      `xmlrpc:"module_l10n_eu_service,omitempty"`
	ModuleMicrosoftCalendar               *Bool      `xmlrpc:"module_microsoft_calendar,omitempty"`
	ModulePad                             *Bool      `xmlrpc:"module_pad,omitempty"`
	ModulePartnerAutocomplete             *Bool      `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModuleProductEmailTemplate            *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductMargin                   *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleSaleAmazon                      *Bool      `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleCoupon                      *Bool      `xmlrpc:"module_sale_coupon,omitempty"`
	ModuleSaleMargin                      *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSaleProductConfigurator         *Bool      `xmlrpc:"module_sale_product_configurator,omitempty"`
	ModuleSaleProductMatrix               *Bool      `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSaleQuotationBuilder            *Bool      `xmlrpc:"module_sale_quotation_builder,omitempty"`
	ModuleSnailmailAccount                *Bool      `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleVoip                            *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                     *Bool      `xmlrpc:"module_web_unsplash,omitempty"`
	PaperformatId                         *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	PartnerAutocompleteInsufficientCredit *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PortalConfirmationPay                 *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	ProductPricelistSetting               *Selection `xmlrpc:"product_pricelist_setting,omitempty"`
	ProductVolumeVolumeInCubicFeet        *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                    *Selection `xmlrpc:"product_weight_in_lbs,omitempty"`
	PurchaseTaxId                         *Many2One  `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                *Bool      `xmlrpc:"qr_code,omitempty"`
	QuotationValidityDays                 *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	ReportFooter                          *String    `xmlrpc:"report_footer,omitempty"`
	SaleTaxId                             *Many2One  `xmlrpc:"sale_tax_id,omitempty"`
	ShowEffect                            *Bool      `xmlrpc:"show_effect,omitempty"`
	ShowLineSubtotalsTaxSelection         *Selection `xmlrpc:"show_line_subtotals_tax_selection,omitempty"`
	SnailmailColor                        *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                        *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                       *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	TaxCalculationRoundingMethod          *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                 *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                        *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TemplateId                            *Many2One  `xmlrpc:"template_id,omitempty"`
	UnsplashAccessKey                     *String    `xmlrpc:"unsplash_access_key,omitempty"`
	UseInvoiceTerms                       *Bool      `xmlrpc:"use_invoice_terms,omitempty"`
	UseQuotationValidityDays              *Bool      `xmlrpc:"use_quotation_validity_days,omitempty"`
	UserDefaultRights                     *Bool      `xmlrpc:"user_default_rights,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	return c.Create(ResConfigSettingsModel, rcs)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config.settings not found", id)
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.config.settings was not found")
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config.settings was not found")
}
