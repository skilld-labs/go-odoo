package odoo

// AccountReport represents account.report model.
type AccountReport struct {
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	AvailabilityCondition    *Selection `xmlrpc:"availability_condition,omitempty"`
	ChartTemplate            *Selection `xmlrpc:"chart_template,omitempty"`
	ColumnIds                *Relation  `xmlrpc:"column_ids,omitempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyTranslation      *Selection `xmlrpc:"currency_translation,omitempty"`
	DefaultOpeningDateFilter *Selection `xmlrpc:"default_opening_date_filter,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	FilterAccountType        *Selection `xmlrpc:"filter_account_type,omitempty"`
	FilterAmlIrFilters       *Bool      `xmlrpc:"filter_aml_ir_filters,omitempty"`
	FilterAnalytic           *Bool      `xmlrpc:"filter_analytic,omitempty"`
	FilterBudgets            *Bool      `xmlrpc:"filter_budgets,omitempty"`
	FilterDateRange          *Bool      `xmlrpc:"filter_date_range,omitempty"`
	FilterFiscalPosition     *Bool      `xmlrpc:"filter_fiscal_position,omitempty"`
	FilterGrowthComparison   *Bool      `xmlrpc:"filter_growth_comparison,omitempty"`
	FilterHide0Lines         *Selection `xmlrpc:"filter_hide_0_lines,omitempty"`
	FilterHierarchy          *Selection `xmlrpc:"filter_hierarchy,omitempty"`
	FilterJournals           *Bool      `xmlrpc:"filter_journals,omitempty"`
	FilterMultiCompany       *Selection `xmlrpc:"filter_multi_company,omitempty"`
	FilterPartner            *Bool      `xmlrpc:"filter_partner,omitempty"`
	FilterPeriodComparison   *Bool      `xmlrpc:"filter_period_comparison,omitempty"`
	FilterShowDraft          *Bool      `xmlrpc:"filter_show_draft,omitempty"`
	FilterUnfoldAll          *Bool      `xmlrpc:"filter_unfold_all,omitempty"`
	FilterUnreconciled       *Bool      `xmlrpc:"filter_unreconciled,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IntegerRounding          *Selection `xmlrpc:"integer_rounding,omitempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omitempty"`
	LoadMoreLimit            *Int       `xmlrpc:"load_more_limit,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	OnlyTaxExigible          *Bool      `xmlrpc:"only_tax_exigible,omitempty"`
	PrefixGroupsThreshold    *Int       `xmlrpc:"prefix_groups_threshold,omitempty"`
	RootReportId             *Many2One  `xmlrpc:"root_report_id,omitempty"`
	SearchBar                *Bool      `xmlrpc:"search_bar,omitempty"`
	SectionMainReportIds     *Relation  `xmlrpc:"section_main_report_ids,omitempty"`
	SectionReportIds         *Relation  `xmlrpc:"section_report_ids,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	UseSections              *Bool      `xmlrpc:"use_sections,omitempty"`
	VariantReportIds         *Relation  `xmlrpc:"variant_report_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReports represents array of account.report model.
type AccountReports []AccountReport

// AccountReportModel is the odoo model name.
const AccountReportModel = "account.report"

// Many2One convert AccountReport to *Many2One.
func (ar *AccountReport) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAccountReport creates a new account.report model and returns its id.
func (c *Client) CreateAccountReport(ar *AccountReport) (int64, error) {
	ids, err := c.CreateAccountReports([]*AccountReport{ar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReport creates a new account.report model and returns its id.
func (c *Client) CreateAccountReports(ars []*AccountReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ars {
		vv = append(vv, v)
	}
	return c.Create(AccountReportModel, vv, nil)
}

// UpdateAccountReport updates an existing account.report record.
func (c *Client) UpdateAccountReport(ar *AccountReport) error {
	return c.UpdateAccountReports([]int64{ar.Id.Get()}, ar)
}

// UpdateAccountReports updates existing account.report records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAccountReports(ids []int64, ar *AccountReport) error {
	return c.Update(AccountReportModel, ids, ar, nil)
}

// DeleteAccountReport deletes an existing account.report record.
func (c *Client) DeleteAccountReport(id int64) error {
	return c.DeleteAccountReports([]int64{id})
}

// DeleteAccountReports deletes existing account.report records.
func (c *Client) DeleteAccountReports(ids []int64) error {
	return c.Delete(AccountReportModel, ids)
}

// GetAccountReport gets account.report existing record.
func (c *Client) GetAccountReport(id int64) (*AccountReport, error) {
	ars, err := c.GetAccountReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// GetAccountReports gets account.report existing records.
func (c *Client) GetAccountReports(ids []int64) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.Read(AccountReportModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReport finds account.report record by querying it with criteria.
func (c *Client) FindAccountReport(criteria *Criteria) (*AccountReport, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// FindAccountReports finds account.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReports(criteria *Criteria, options *Options) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportModel, criteria, options)
}

// FindAccountReportId finds record id by querying it with criteria.
func (c *Client) FindAccountReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
