package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                           *Time      `xmlrpc:"__last_update,omptempty"`
	AccountHideSetupBar                  *Bool      `xmlrpc:"account_hide_setup_bar,omptempty"`
	AliasDomain                          *String    `xmlrpc:"alias_domain,omptempty"`
	AuthSignupResetPassword              *Bool      `xmlrpc:"auth_signup_reset_password,omptempty"`
	AuthSignupTemplateUserId             *Many2One  `xmlrpc:"auth_signup_template_user_id,omptempty"`
	AuthSignupUninvited                  *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	AutoDoneSetting                      *Bool      `xmlrpc:"auto_done_setting,omptempty"`
	ChartTemplateId                      *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CodeDigits                           *Int       `xmlrpc:"code_digits,omptempty"`
	CompanyCurrencyId                    *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                            *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanySharePartner                  *Bool      `xmlrpc:"company_share_partner,omptempty"`
	CompanyShareProduct                  *Bool      `xmlrpc:"company_share_product,omptempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmAliasPrefix                       *String    `xmlrpc:"crm_alias_prefix,omptempty"`
	CurrencyExchangeJournalId            *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                           *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCustomReportFooter            *Bool      `xmlrpc:"default_custom_report_footer,omptempty"`
	DefaultDepositProductId              *Many2One  `xmlrpc:"default_deposit_product_id,omptempty"`
	DefaultExternalEmailServer           *Bool      `xmlrpc:"default_external_email_server,omptempty"`
	DefaultInvoicePolicy                 *Selection `xmlrpc:"default_invoice_policy,omptempty"`
	DefaultPickingPolicy                 *Selection `xmlrpc:"default_picking_policy,omptempty"`
	DefaultPurchaseMethod                *Selection `xmlrpc:"default_purchase_method,omptempty"`
	DefaultPurchaseTaxId                 *Many2One  `xmlrpc:"default_purchase_tax_id,omptempty"`
	DefaultSaleTaxId                     *Many2One  `xmlrpc:"default_sale_tax_id,omptempty"`
	DefaultUserRights                    *Bool      `xmlrpc:"default_user_rights,omptempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omptempty"`
	ExternalReportLayout                 *Selection `xmlrpc:"external_report_layout,omptempty"`
	FailCounter                          *Int       `xmlrpc:"fail_counter,omptempty"`
	GenerateLeadFromAlias                *Bool      `xmlrpc:"generate_lead_from_alias,omptempty"`
	GroupAnalyticAccountForPurchases     *Bool      `xmlrpc:"group_analytic_account_for_purchases,omptempty"`
	GroupAnalyticAccounting              *Bool      `xmlrpc:"group_analytic_accounting,omptempty"`
	GroupCashRounding                    *Bool      `xmlrpc:"group_cash_rounding,omptempty"`
	GroupDiscountPerSoLine               *Bool      `xmlrpc:"group_discount_per_so_line,omptempty"`
	GroupDisplayIncoterm                 *Bool      `xmlrpc:"group_display_incoterm,omptempty"`
	GroupManageVendorPrice               *Bool      `xmlrpc:"group_manage_vendor_price,omptempty"`
	GroupMassMailingCampaign             *Bool      `xmlrpc:"group_mass_mailing_campaign,omptempty"`
	GroupMultiCompany                    *Bool      `xmlrpc:"group_multi_company,omptempty"`
	GroupMultiCurrency                   *Bool      `xmlrpc:"group_multi_currency,omptempty"`
	GroupPricelistItem                   *Bool      `xmlrpc:"group_pricelist_item,omptempty"`
	GroupProductPricelist                *Bool      `xmlrpc:"group_product_pricelist,omptempty"`
	GroupProductVariant                  *Bool      `xmlrpc:"group_product_variant,omptempty"`
	GroupProformaSales                   *Bool      `xmlrpc:"group_proforma_sales,omptempty"`
	GroupRouteSoLines                    *Bool      `xmlrpc:"group_route_so_lines,omptempty"`
	GroupSaleDeliveryAddress             *Bool      `xmlrpc:"group_sale_delivery_address,omptempty"`
	GroupSaleLayout                      *Bool      `xmlrpc:"group_sale_layout,omptempty"`
	GroupSalePricelist                   *Bool      `xmlrpc:"group_sale_pricelist,omptempty"`
	GroupShowPriceSubtotal               *Bool      `xmlrpc:"group_show_price_subtotal,omptempty"`
	GroupShowPriceTotal                  *Bool      `xmlrpc:"group_show_price_total,omptempty"`
	GroupStockAdvLocation                *Bool      `xmlrpc:"group_stock_adv_location,omptempty"`
	GroupStockMultiLocations             *Bool      `xmlrpc:"group_stock_multi_locations,omptempty"`
	GroupStockMultiWarehouses            *Bool      `xmlrpc:"group_stock_multi_warehouses,omptempty"`
	GroupStockPackaging                  *Bool      `xmlrpc:"group_stock_packaging,omptempty"`
	GroupStockProductionLot              *Bool      `xmlrpc:"group_stock_production_lot,omptempty"`
	GroupStockTrackingLot                *Bool      `xmlrpc:"group_stock_tracking_lot,omptempty"`
	GroupStockTrackingOwner              *Bool      `xmlrpc:"group_stock_tracking_owner,omptempty"`
	GroupSubtaskProject                  *Bool      `xmlrpc:"group_subtask_project,omptempty"`
	GroupUom                             *Bool      `xmlrpc:"group_uom,omptempty"`
	GroupUseLead                         *Bool      `xmlrpc:"group_use_lead,omptempty"`
	GroupWarningAccount                  *Bool      `xmlrpc:"group_warning_account,omptempty"`
	GroupWarningPurchase                 *Bool      `xmlrpc:"group_warning_purchase,omptempty"`
	GroupWarningSale                     *Bool      `xmlrpc:"group_warning_sale,omptempty"`
	GroupWarningStock                    *Bool      `xmlrpc:"group_warning_stock,omptempty"`
	HasAccountingEntries                 *Bool      `xmlrpc:"has_accounting_entries,omptempty"`
	HasChartOfAccounts                   *Bool      `xmlrpc:"has_chart_of_accounts,omptempty"`
	Id                                   *Int       `xmlrpc:"id,omptempty"`
	IsInstalledSale                      *Bool      `xmlrpc:"is_installed_sale,omptempty"`
	LeaveTimesheetProjectId              *Many2One  `xmlrpc:"leave_timesheet_project_id,omptempty"`
	LeaveTimesheetTaskId                 *Many2One  `xmlrpc:"leave_timesheet_task_id,omptempty"`
	LockConfirmedPo                      *Bool      `xmlrpc:"lock_confirmed_po,omptempty"`
	ModuleAccount3wayMatch               *Bool      `xmlrpc:"module_account_3way_match,omptempty"`
	ModuleAccountAccountant              *Bool      `xmlrpc:"module_account_accountant,omptempty"`
	ModuleAccountAsset                   *Bool      `xmlrpc:"module_account_asset,omptempty"`
	ModuleAccountBankStatementImportCamt *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omptempty"`
	ModuleAccountBankStatementImportCsv  *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omptempty"`
	ModuleAccountBankStatementImportOfx  *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omptempty"`
	ModuleAccountBankStatementImportQif  *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omptempty"`
	ModuleAccountBatchDeposit            *Bool      `xmlrpc:"module_account_batch_deposit,omptempty"`
	ModuleAccountBudget                  *Bool      `xmlrpc:"module_account_budget,omptempty"`
	ModuleAccountDeferredRevenue         *Bool      `xmlrpc:"module_account_deferred_revenue,omptempty"`
	ModuleAccountPayment                 *Bool      `xmlrpc:"module_account_payment,omptempty"`
	ModuleAccountPlaid                   *Bool      `xmlrpc:"module_account_plaid,omptempty"`
	ModuleAccountReports                 *Bool      `xmlrpc:"module_account_reports,omptempty"`
	ModuleAccountReportsFollowup         *Bool      `xmlrpc:"module_account_reports_followup,omptempty"`
	ModuleAccountSepa                    *Bool      `xmlrpc:"module_account_sepa,omptempty"`
	ModuleAccountSepaDirectDebit         *Bool      `xmlrpc:"module_account_sepa_direct_debit,omptempty"`
	ModuleAccountTaxcloud                *Bool      `xmlrpc:"module_account_taxcloud,omptempty"`
	ModuleAccountYodlee                  *Bool      `xmlrpc:"module_account_yodlee,omptempty"`
	ModuleAuthLdap                       *Bool      `xmlrpc:"module_auth_ldap,omptempty"`
	ModuleAuthOauth                      *Bool      `xmlrpc:"module_auth_oauth,omptempty"`
	ModuleBaseGengo                      *Bool      `xmlrpc:"module_base_gengo,omptempty"`
	ModuleBaseImport                     *Bool      `xmlrpc:"module_base_import,omptempty"`
	ModuleCrmPhoneValidation             *Bool      `xmlrpc:"module_crm_phone_validation,omptempty"`
	ModuleCurrencyRateLive               *Bool      `xmlrpc:"module_currency_rate_live,omptempty"`
	ModuleDelivery                       *Bool      `xmlrpc:"module_delivery,omptempty"`
	ModuleDeliveryBpost                  *Bool      `xmlrpc:"module_delivery_bpost,omptempty"`
	ModuleDeliveryDhl                    *Bool      `xmlrpc:"module_delivery_dhl,omptempty"`
	ModuleDeliveryFedex                  *Bool      `xmlrpc:"module_delivery_fedex,omptempty"`
	ModuleDeliveryUps                    *Bool      `xmlrpc:"module_delivery_ups,omptempty"`
	ModuleDeliveryUsps                   *Bool      `xmlrpc:"module_delivery_usps,omptempty"`
	ModuleGoogleCalendar                 *Bool      `xmlrpc:"module_google_calendar,omptempty"`
	ModuleGoogleDrive                    *Bool      `xmlrpc:"module_google_drive,omptempty"`
	ModuleGoogleSpreadsheet              *Bool      `xmlrpc:"module_google_spreadsheet,omptempty"`
	ModuleHrOrgChart                     *Bool      `xmlrpc:"module_hr_org_chart,omptempty"`
	ModuleHrTimesheet                    *Bool      `xmlrpc:"module_hr_timesheet,omptempty"`
	ModuleInterCompanyRules              *Bool      `xmlrpc:"module_inter_company_rules,omptempty"`
	ModuleL10NEuService                  *Bool      `xmlrpc:"module_l10n_eu_service,omptempty"`
	ModuleL10NUsCheckPrinting            *Bool      `xmlrpc:"module_l10n_us_check_printing,omptempty"`
	ModulePad                            *Bool      `xmlrpc:"module_pad,omptempty"`
	ModulePrintDocsaway                  *Bool      `xmlrpc:"module_print_docsaway,omptempty"`
	ModuleProcurementJit                 *Selection `xmlrpc:"module_procurement_jit,omptempty"`
	ModuleProductEmailTemplate           *Bool      `xmlrpc:"module_product_email_template,omptempty"`
	ModuleProductExpiry                  *Bool      `xmlrpc:"module_product_expiry,omptempty"`
	ModuleProductMargin                  *Bool      `xmlrpc:"module_product_margin,omptempty"`
	ModuleProjectForecast                *Bool      `xmlrpc:"module_project_forecast,omptempty"`
	ModuleProjectTimesheetHolidays       *Bool      `xmlrpc:"module_project_timesheet_holidays,omptempty"`
	ModuleProjectTimesheetSynchro        *Bool      `xmlrpc:"module_project_timesheet_synchro,omptempty"`
	ModulePurchaseRequisition            *Bool      `xmlrpc:"module_purchase_requisition,omptempty"`
	ModuleRatingProject                  *Bool      `xmlrpc:"module_rating_project,omptempty"`
	ModuleSaleCoupon                     *Bool      `xmlrpc:"module_sale_coupon,omptempty"`
	ModuleSaleMargin                     *Bool      `xmlrpc:"module_sale_margin,omptempty"`
	ModuleSaleOrderDates                 *Bool      `xmlrpc:"module_sale_order_dates,omptempty"`
	ModuleSalePayment                    *Bool      `xmlrpc:"module_sale_payment,omptempty"`
	ModuleSaleTimesheet                  *Bool      `xmlrpc:"module_sale_timesheet,omptempty"`
	ModuleStockBarcode                   *Bool      `xmlrpc:"module_stock_barcode,omptempty"`
	ModuleStockDropshipping              *Bool      `xmlrpc:"module_stock_dropshipping,omptempty"`
	ModuleStockLandedCosts               *Bool      `xmlrpc:"module_stock_landed_costs,omptempty"`
	ModuleStockPickingBatch              *Bool      `xmlrpc:"module_stock_picking_batch,omptempty"`
	ModuleVoip                           *Bool      `xmlrpc:"module_voip,omptempty"`
	ModuleWebClearbit                    *Bool      `xmlrpc:"module_web_clearbit,omptempty"`
	ModuleWebsiteQuote                   *Bool      `xmlrpc:"module_website_quote,omptempty"`
	ModuleWebsiteSaleDigital             *Bool      `xmlrpc:"module_website_sale_digital,omptempty"`
	MultiSalesPrice                      *Bool      `xmlrpc:"multi_sales_price,omptempty"`
	MultiSalesPriceMethod                *Selection `xmlrpc:"multi_sales_price_method,omptempty"`
	PaperformatId                        *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	PoDoubleValidation                   *Selection `xmlrpc:"po_double_validation,omptempty"`
	PoDoubleValidationAmount             *Float     `xmlrpc:"po_double_validation_amount,omptempty"`
	PoLead                               *Float     `xmlrpc:"po_lead,omptempty"`
	PoLock                               *Selection `xmlrpc:"po_lock,omptempty"`
	PoOrderApproval                      *Bool      `xmlrpc:"po_order_approval,omptempty"`
	PortalConfirmation                   *Bool      `xmlrpc:"portal_confirmation,omptempty"`
	PortalConfirmationOptions            *Selection `xmlrpc:"portal_confirmation_options,omptempty"`
	ProjectTimeModeId                    *Many2One  `xmlrpc:"project_time_mode_id,omptempty"`
	PropagationMinimumDelta              *Int       `xmlrpc:"propagation_minimum_delta,omptempty"`
	ReportFooter                         *String    `xmlrpc:"report_footer,omptempty"`
	ResourceCalendarId                   *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	SaleNote                             *String    `xmlrpc:"sale_note,omptempty"`
	SalePricelistSetting                 *Selection `xmlrpc:"sale_pricelist_setting,omptempty"`
	SaleShowTax                          *Selection `xmlrpc:"sale_show_tax,omptempty"`
	SecurityLead                         *Float     `xmlrpc:"security_lead,omptempty"`
	TaxCalculationRoundingMethod         *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                       *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	UsePoLead                            *Bool      `xmlrpc:"use_po_lead,omptempty"`
	UsePropagationMinimumDelta           *Bool      `xmlrpc:"use_propagation_minimum_delta,omptempty"`
	UseSaleNote                          *Bool      `xmlrpc:"use_sale_note,omptempty"`
	UseSecurityLead                      *Bool      `xmlrpc:"use_security_lead,omptempty"`
	VatCheckVies                         *Bool      `xmlrpc:"vat_check_vies,omptempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return nil, fmt.Errorf("no res.config.settings was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no res.config.settings was found with criteria %v and options %v", criteria, options)
}
