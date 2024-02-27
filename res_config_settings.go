package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                           *Time      `xmlrpc:"__last_update,omitempty"`
	AccountHideSetupBar                  *Bool      `xmlrpc:"account_hide_setup_bar,omitempty"`
	AliasDomain                          *String    `xmlrpc:"alias_domain,omitempty"`
	AuthSignupUninvited                  *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutoDoneSetting                      *Bool      `xmlrpc:"auto_done_setting,omitempty"`
	ChartTemplateId                      *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CodeDigits                           *Int       `xmlrpc:"code_digits,omitempty"`
	CompanyCurrencyId                    *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                            *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanySharePartner                  *Bool      `xmlrpc:"company_share_partner,omitempty"`
	CompanyShareProduct                  *Bool      `xmlrpc:"company_share_product,omitempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmAliasPrefix                       *String    `xmlrpc:"crm_alias_prefix,omitempty"`
	CurrencyExchangeJournalId            *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                           *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCustomReportFooter            *Bool      `xmlrpc:"default_custom_report_footer,omitempty"`
	DefaultDepositProductId              *Many2One  `xmlrpc:"default_deposit_product_id,omitempty"`
	DefaultExternalEmailServer           *Bool      `xmlrpc:"default_external_email_server,omitempty"`
	DefaultInvoicePolicy                 *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                 *Selection `xmlrpc:"default_picking_policy,omitempty"`
	DefaultPurchaseMethod                *Selection `xmlrpc:"default_purchase_method,omitempty"`
	DefaultPurchaseTaxId                 *Many2One  `xmlrpc:"default_purchase_tax_id,omitempty"`
	DefaultSaleTaxId                     *Many2One  `xmlrpc:"default_sale_tax_id,omitempty"`
	DefaultUserRights                    *Bool      `xmlrpc:"default_user_rights,omitempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omitempty"`
	ExternalReportLayout                 *Selection `xmlrpc:"external_report_layout,omitempty"`
	FailCounter                          *Int       `xmlrpc:"fail_counter,omitempty"`
	GenerateLeadFromAlias                *Bool      `xmlrpc:"generate_lead_from_alias,omitempty"`
	GroupAnalyticAccountForPurchases     *Bool      `xmlrpc:"group_analytic_account_for_purchases,omitempty"`
	GroupAnalyticAccounting              *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupCashRounding                    *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine               *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupDisplayIncoterm                 *Bool      `xmlrpc:"group_display_incoterm,omitempty"`
	GroupManageVendorPrice               *Bool      `xmlrpc:"group_manage_vendor_price,omitempty"`
	GroupMassMailingCampaign             *Bool      `xmlrpc:"group_mass_mailing_campaign,omitempty"`
	GroupMultiCompany                    *Bool      `xmlrpc:"group_multi_company,omitempty"`
	GroupMultiCurrency                   *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupPricelistItem                   *Bool      `xmlrpc:"group_pricelist_item,omitempty"`
	GroupProductPricelist                *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                  *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                   *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupRouteSoLines                    *Bool      `xmlrpc:"group_route_so_lines,omitempty"`
	GroupSaleDeliveryAddress             *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleLayout                      *Bool      `xmlrpc:"group_sale_layout,omitempty"`
	GroupSalePricelist                   *Bool      `xmlrpc:"group_sale_pricelist,omitempty"`
	GroupShowPriceSubtotal               *Bool      `xmlrpc:"group_show_price_subtotal,omitempty"`
	GroupShowPriceTotal                  *Bool      `xmlrpc:"group_show_price_total,omitempty"`
	GroupStockAdvLocation                *Bool      `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockMultiLocations             *Bool      `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockMultiWarehouses            *Bool      `xmlrpc:"group_stock_multi_warehouses,omitempty"`
	GroupStockPackaging                  *Bool      `xmlrpc:"group_stock_packaging,omitempty"`
	GroupStockProductionLot              *Bool      `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockTrackingLot                *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner              *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupSubtaskProject                  *Bool      `xmlrpc:"group_subtask_project,omitempty"`
	GroupUom                             *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupUseLead                         *Bool      `xmlrpc:"group_use_lead,omitempty"`
	GroupWarningAccount                  *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningPurchase                 *Bool      `xmlrpc:"group_warning_purchase,omitempty"`
	GroupWarningSale                     *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                    *Bool      `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                 *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                   *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	Id                                   *Int       `xmlrpc:"id,omitempty"`
	IsInstalledSale                      *Bool      `xmlrpc:"is_installed_sale,omitempty"`
	Ldaps                                *Relation  `xmlrpc:"ldaps,omitempty"`
	LeaveTimesheetProjectId              *Many2One  `xmlrpc:"leave_timesheet_project_id,omitempty"`
	LeaveTimesheetTaskId                 *Many2One  `xmlrpc:"leave_timesheet_task_id,omitempty"`
	LockConfirmedPo                      *Bool      `xmlrpc:"lock_confirmed_po,omitempty"`
	ModuleAccount3WayMatch               *Bool      `xmlrpc:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant              *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountAsset                   *Bool      `xmlrpc:"module_account_asset,omitempty"`
	ModuleAccountBankStatementImportCamt *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv  *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx  *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif  *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchDeposit            *Bool      `xmlrpc:"module_account_batch_deposit,omitempty"`
	ModuleAccountBudget                  *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountDeferredRevenue         *Bool      `xmlrpc:"module_account_deferred_revenue,omitempty"`
	ModuleAccountPayment                 *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPlaid                   *Bool      `xmlrpc:"module_account_plaid,omitempty"`
	ModuleAccountReports                 *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountReportsFollowup         *Bool      `xmlrpc:"module_account_reports_followup,omitempty"`
	ModuleAccountSepa                    *Bool      `xmlrpc:"module_account_sepa,omitempty"`
	ModuleAccountSepaDirectDebit         *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAccountTaxcloud                *Bool      `xmlrpc:"module_account_taxcloud,omitempty"`
	ModuleAccountYodlee                  *Bool      `xmlrpc:"module_account_yodlee,omitempty"`
	ModuleAuthLdap                       *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                      *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGengo                      *Bool      `xmlrpc:"module_base_gengo,omitempty"`
	ModuleBaseImport                     *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCrmPhoneValidation             *Bool      `xmlrpc:"module_crm_phone_validation,omitempty"`
	ModuleCurrencyRateLive               *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                       *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                  *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                    *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryFedex                  *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliveryUps                    *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                   *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                 *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleDrive                    *Bool      `xmlrpc:"module_google_drive,omitempty"`
	ModuleGoogleSpreadsheet              *Bool      `xmlrpc:"module_google_spreadsheet,omitempty"`
	ModuleHrOrgChart                     *Bool      `xmlrpc:"module_hr_org_chart,omitempty"`
	ModuleHrTimesheet                    *Bool      `xmlrpc:"module_hr_timesheet,omitempty"`
	ModuleInterCompanyRules              *Bool      `xmlrpc:"module_inter_company_rules,omitempty"`
	ModuleL10NEuService                  *Bool      `xmlrpc:"module_l10n_eu_service,omitempty"`
	ModuleL10NUsCheckPrinting            *Bool      `xmlrpc:"module_l10n_us_check_printing,omitempty"`
	ModulePad                            *Bool      `xmlrpc:"module_pad,omitempty"`
	ModulePrintDocsaway                  *Bool      `xmlrpc:"module_print_docsaway,omitempty"`
	ModuleProcurementJit                 *Selection `xmlrpc:"module_procurement_jit,omitempty"`
	ModuleProductEmailTemplate           *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                  *Bool      `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductMargin                  *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleProjectForecast                *Bool      `xmlrpc:"module_project_forecast,omitempty"`
	ModuleProjectTimesheetHolidays       *Bool      `xmlrpc:"module_project_timesheet_holidays,omitempty"`
	ModuleProjectTimesheetSynchro        *Bool      `xmlrpc:"module_project_timesheet_synchro,omitempty"`
	ModulePurchaseRequisition            *Bool      `xmlrpc:"module_purchase_requisition,omitempty"`
	ModuleRatingProject                  *Bool      `xmlrpc:"module_rating_project,omitempty"`
	ModuleSaleCoupon                     *Bool      `xmlrpc:"module_sale_coupon,omitempty"`
	ModuleSaleMargin                     *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSaleOrderDates                 *Bool      `xmlrpc:"module_sale_order_dates,omitempty"`
	ModuleSalePayment                    *Bool      `xmlrpc:"module_sale_payment,omitempty"`
	ModuleSaleTimesheet                  *Bool      `xmlrpc:"module_sale_timesheet,omitempty"`
	ModuleStockBarcode                   *Bool      `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockDropshipping              *Bool      `xmlrpc:"module_stock_dropshipping,omitempty"`
	ModuleStockLandedCosts               *Bool      `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch              *Bool      `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleVoip                           *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebClearbit                    *Bool      `xmlrpc:"module_web_clearbit,omitempty"`
	ModuleWebsiteQuote                   *Bool      `xmlrpc:"module_website_quote,omitempty"`
	ModuleWebsiteSaleDigital             *Bool      `xmlrpc:"module_website_sale_digital,omitempty"`
	MultiSalesPrice                      *Bool      `xmlrpc:"multi_sales_price,omitempty"`
	MultiSalesPriceMethod                *Selection `xmlrpc:"multi_sales_price_method,omitempty"`
	PaperformatId                        *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	PoDoubleValidation                   *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount             *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                               *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                               *Selection `xmlrpc:"po_lock,omitempty"`
	PoOrderApproval                      *Bool      `xmlrpc:"po_order_approval,omitempty"`
	PortalConfirmation                   *Bool      `xmlrpc:"portal_confirmation,omitempty"`
	PortalConfirmationOptions            *Selection `xmlrpc:"portal_confirmation_options,omitempty"`
	ProjectTimeModeId                    *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PropagationMinimumDelta              *Int       `xmlrpc:"propagation_minimum_delta,omitempty"`
	ReportFooter                         *String    `xmlrpc:"report_footer,omitempty"`
	ResourceCalendarId                   *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SaleNote                             *String    `xmlrpc:"sale_note,omitempty"`
	SalePricelistSetting                 *Selection `xmlrpc:"sale_pricelist_setting,omitempty"`
	SaleShowTax                          *Selection `xmlrpc:"sale_show_tax,omitempty"`
	SecurityLead                         *Float     `xmlrpc:"security_lead,omitempty"`
	TaxCalculationRoundingMethod         *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                       *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	UsePoLead                            *Bool      `xmlrpc:"use_po_lead,omitempty"`
	UsePropagationMinimumDelta           *Bool      `xmlrpc:"use_propagation_minimum_delta,omitempty"`
	UseSaleNote                          *Bool      `xmlrpc:"use_sale_note,omitempty"`
	UseSecurityLead                      *Bool      `xmlrpc:"use_security_lead,omitempty"`
	VatCheckVies                         *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
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
	return &((*rcss)[0]), nil
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
	return &((*rcss)[0]), nil
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
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
