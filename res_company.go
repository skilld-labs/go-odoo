package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omitempty"`
	AccountNo                         *String    `xmlrpc:"account_no,omitempty"`
	AccountOpeningDate                *Time      `xmlrpc:"account_opening_date,omitempty"`
	AccountOpeningJournalId           *Many2One  `xmlrpc:"account_opening_journal_id,omitempty"`
	AccountOpeningMoveId              *Many2One  `xmlrpc:"account_opening_move_id,omitempty"`
	AccountSetupBankDataDone          *Bool      `xmlrpc:"account_setup_bank_data_done,omitempty"`
	AccountSetupBarClosed             *Bool      `xmlrpc:"account_setup_bar_closed,omitempty"`
	AccountSetupCoaDone               *Bool      `xmlrpc:"account_setup_coa_done,omitempty"`
	AccountSetupCompanyDataDone       *Bool      `xmlrpc:"account_setup_company_data_done,omitempty"`
	AccountSetupFyDataDone            *Bool      `xmlrpc:"account_setup_fy_data_done,omitempty"`
	AccountsCodeDigits                *Int       `xmlrpc:"accounts_code_digits,omitempty"`
	AngloSaxonAccounting              *Bool      `xmlrpc:"anglo_saxon_accounting,omitempty"`
	Ape                               *String    `xmlrpc:"ape,omitempty"`
	BankAccountCodePrefix             *String    `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankIds                           *Relation  `xmlrpc:"bank_ids,omitempty"`
	BankJournalIds                    *Relation  `xmlrpc:"bank_journal_ids,omitempty"`
	CashAccountCodePrefix             *String    `xmlrpc:"cash_account_code_prefix,omitempty"`
	ChartTemplateId                   *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	ChildIds                          *Relation  `xmlrpc:"child_ids,omitempty"`
	City                              *String    `xmlrpc:"city,omitempty"`
	CompanyRegistry                   *String    `xmlrpc:"company_registry,omitempty"`
	CountryId                         *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId         *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                        *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omitempty"`
	Email                             *String    `xmlrpc:"email,omitempty"`
	ExpectsChartOfAccounts            *Bool      `xmlrpc:"expects_chart_of_accounts,omitempty"`
	ExpenseCurrencyExchangeAccountId  *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalReportLayout              *Selection `xmlrpc:"external_report_layout,omitempty"`
	FiscalyearLastDay                 *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth               *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                *Time      `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Id                                *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId   *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	InternalTransitLocationId         *Many2One  `xmlrpc:"internal_transit_location_id,omitempty"`
	LeaveTimesheetProjectId           *Many2One  `xmlrpc:"leave_timesheet_project_id,omitempty"`
	LeaveTimesheetTaskId              *Many2One  `xmlrpc:"leave_timesheet_task_id,omitempty"`
	Logo                              *String    `xmlrpc:"logo,omitempty"`
	LogoWeb                           *String    `xmlrpc:"logo_web,omitempty"`
	Name                              *String    `xmlrpc:"name,omitempty"`
	OverdueMsg                        *String    `xmlrpc:"overdue_msg,omitempty"`
	PaperformatId                     *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	ParentId                          *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerId                         *Many2One  `xmlrpc:"partner_id,omitempty"`
	PeriodLockDate                    *Time      `xmlrpc:"period_lock_date,omitempty"`
	Phone                             *String    `xmlrpc:"phone,omitempty"`
	PoDoubleValidation                *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount          *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                            *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                            *Selection `xmlrpc:"po_lock,omitempty"`
	ProjectTimeModeId                 *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PropagationMinimumDelta           *Int       `xmlrpc:"propagation_minimum_delta,omitempty"`
	PropertyStockAccountInputCategId  *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockValuationAccountId   *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	ReportFooter                      *String    `xmlrpc:"report_footer,omitempty"`
	ReportHeader                      *String    `xmlrpc:"report_header,omitempty"`
	ResourceCalendarId                *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceCalendarIds               *Relation  `xmlrpc:"resource_calendar_ids,omitempty"`
	SaleNote                          *String    `xmlrpc:"sale_note,omitempty"`
	SecurityLead                      *Float     `xmlrpc:"security_lead,omitempty"`
	Sequence                          *Int       `xmlrpc:"sequence,omitempty"`
	Siret                             *String    `xmlrpc:"siret,omitempty"`
	SocialFacebook                    *String    `xmlrpc:"social_facebook,omitempty"`
	SocialGithub                      *String    `xmlrpc:"social_github,omitempty"`
	SocialGoogleplus                  *String    `xmlrpc:"social_googleplus,omitempty"`
	SocialLinkedin                    *String    `xmlrpc:"social_linkedin,omitempty"`
	SocialTwitter                     *String    `xmlrpc:"social_twitter,omitempty"`
	SocialYoutube                     *String    `xmlrpc:"social_youtube,omitempty"`
	StateId                           *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                            *String    `xmlrpc:"street,omitempty"`
	Street2                           *String    `xmlrpc:"street2,omitempty"`
	TaxCalculationRoundingMethod      *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId             *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                    *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TransferAccountId                 *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	UserIds                           *Relation  `xmlrpc:"user_ids,omitempty"`
	Vat                               *String    `xmlrpc:"vat,omitempty"`
	VatCheckVies                      *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	Website                           *String    `xmlrpc:"website,omitempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omitempty"`
	Zip                               *String    `xmlrpc:"zip,omitempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	return c.Create(ResCompanyModel, rc)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("no res.company was found with criteria %v", criteria)
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no res.company was found with criteria %v and options %v", criteria, options)
}
