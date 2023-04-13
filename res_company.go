package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountNo                         *String    `xmlrpc:"account_no,omptempty"`
	AccountOpeningDate                *Time      `xmlrpc:"account_opening_date,omptempty"`
	AccountOpeningJournalId           *Many2One  `xmlrpc:"account_opening_journal_id,omptempty"`
	AccountOpeningMoveId              *Many2One  `xmlrpc:"account_opening_move_id,omptempty"`
	AccountSetupBankDataDone          *Bool      `xmlrpc:"account_setup_bank_data_done,omptempty"`
	AccountSetupBarClosed             *Bool      `xmlrpc:"account_setup_bar_closed,omptempty"`
	AccountSetupCoaDone               *Bool      `xmlrpc:"account_setup_coa_done,omptempty"`
	AccountSetupCompanyDataDone       *Bool      `xmlrpc:"account_setup_company_data_done,omptempty"`
	AccountSetupFyDataDone            *Bool      `xmlrpc:"account_setup_fy_data_done,omptempty"`
	AccountsCodeDigits                *Int       `xmlrpc:"accounts_code_digits,omptempty"`
	AngloSaxonAccounting              *Bool      `xmlrpc:"anglo_saxon_accounting,omptempty"`
	Ape                               *String    `xmlrpc:"ape,omptempty"`
	BankAccountCodePrefix             *String    `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankIds                           *Relation  `xmlrpc:"bank_ids,omptempty"`
	BankJournalIds                    *Relation  `xmlrpc:"bank_journal_ids,omptempty"`
	CashAccountCodePrefix             *String    `xmlrpc:"cash_account_code_prefix,omptempty"`
	ChartTemplateId                   *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	ChildIds                          *Relation  `xmlrpc:"child_ids,omptempty"`
	City                              *String    `xmlrpc:"city,omptempty"`
	CompanyRegistry                   *String    `xmlrpc:"company_registry,omptempty"`
	CountryId                         *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId         *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omptempty"`
	Email                             *String    `xmlrpc:"email,omptempty"`
	ExpectsChartOfAccounts            *Bool      `xmlrpc:"expects_chart_of_accounts,omptempty"`
	ExpenseCurrencyExchangeAccountId  *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExternalReportLayout              *Selection `xmlrpc:"external_report_layout,omptempty"`
	FiscalyearLastDay                 *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth               *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Id                                *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId   *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	InternalTransitLocationId         *Many2One  `xmlrpc:"internal_transit_location_id,omptempty"`
	LeaveTimesheetProjectId           *Many2One  `xmlrpc:"leave_timesheet_project_id,omptempty"`
	LeaveTimesheetTaskId              *Many2One  `xmlrpc:"leave_timesheet_task_id,omptempty"`
	Logo                              *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                           *String    `xmlrpc:"logo_web,omptempty"`
	Name                              *String    `xmlrpc:"name,omptempty"`
	OverdueMsg                        *String    `xmlrpc:"overdue_msg,omptempty"`
	PaperformatId                     *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                          *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerId                         *Many2One  `xmlrpc:"partner_id,omptempty"`
	PeriodLockDate                    *Time      `xmlrpc:"period_lock_date,omptempty"`
	Phone                             *String    `xmlrpc:"phone,omptempty"`
	PoDoubleValidation                *Selection `xmlrpc:"po_double_validation,omptempty"`
	PoDoubleValidationAmount          *Float     `xmlrpc:"po_double_validation_amount,omptempty"`
	PoLead                            *Float     `xmlrpc:"po_lead,omptempty"`
	PoLock                            *Selection `xmlrpc:"po_lock,omptempty"`
	ProjectTimeModeId                 *Many2One  `xmlrpc:"project_time_mode_id,omptempty"`
	PropagationMinimumDelta           *Int       `xmlrpc:"propagation_minimum_delta,omptempty"`
	PropertyStockAccountInputCategId  *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId   *Many2One  `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	ReportFooter                      *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader                      *String    `xmlrpc:"report_header,omptempty"`
	ResourceCalendarId                *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceCalendarIds               *Relation  `xmlrpc:"resource_calendar_ids,omptempty"`
	SaleNote                          *String    `xmlrpc:"sale_note,omptempty"`
	SecurityLead                      *Float     `xmlrpc:"security_lead,omptempty"`
	Sequence                          *Int       `xmlrpc:"sequence,omptempty"`
	Siret                             *String    `xmlrpc:"siret,omptempty"`
	SocialFacebook                    *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                      *String    `xmlrpc:"social_github,omptempty"`
	SocialGoogleplus                  *String    `xmlrpc:"social_googleplus,omptempty"`
	SocialLinkedin                    *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                     *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                     *String    `xmlrpc:"social_youtube,omptempty"`
	StateId                           *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                            *String    `xmlrpc:"street,omptempty"`
	Street2                           *String    `xmlrpc:"street2,omptempty"`
	TaxCalculationRoundingMethod      *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId             *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                    *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TransferAccountId                 *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	UserIds                           *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                               *String    `xmlrpc:"vat,omptempty"`
	VatCheckVies                      *Bool      `xmlrpc:"vat_check_vies,omptempty"`
	Website                           *String    `xmlrpc:"website,omptempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                               *String    `xmlrpc:"zip,omptempty"`
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
	ids, err := c.Create(ResCompanyModel, []interface{}{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv)
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
	return nil, fmt.Errorf("res.company was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("res.company was not found with criteria %v and options %v", criteria, options)
}
