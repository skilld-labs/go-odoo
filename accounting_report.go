package odoo

// AccountingReport represents accounting.report model.
type AccountingReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	AccountReportId *Many2One  `xmlrpc:"account_report_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omitempty"`
	DateFromCmp     *Time      `xmlrpc:"date_from_cmp,omitempty"`
	DateTo          *Time      `xmlrpc:"date_to,omitempty"`
	DateToCmp       *Time      `xmlrpc:"date_to_cmp,omitempty"`
	DebitCredit     *Bool      `xmlrpc:"debit_credit,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	EnableFilter    *Bool      `xmlrpc:"enable_filter,omitempty"`
	FilterCmp       *Selection `xmlrpc:"filter_cmp,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	JournalIds      *Relation  `xmlrpc:"journal_ids,omitempty"`
	LabelFilter     *String    `xmlrpc:"label_filter,omitempty"`
	TargetMove      *Selection `xmlrpc:"target_move,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountingReports represents array of accounting.report model.
type AccountingReports []AccountingReport

// AccountingReportModel is the odoo model name.
const AccountingReportModel = "accounting.report"

// Many2One convert AccountingReport to *Many2One.
func (ar *AccountingReport) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAccountingReport creates a new accounting.report model and returns its id.
func (c *Client) CreateAccountingReport(ar *AccountingReport) (int64, error) {
	ids, err := c.CreateAccountingReports([]*AccountingReport{ar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountingReports creates a new accounting.report model and returns its id.
func (c *Client) CreateAccountingReports(ars []*AccountingReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ars {
		vv = append(vv, v)
	}
	return c.Create(AccountingReportModel, vv, nil)
}

// UpdateAccountingReport updates an existing accounting.report record.
func (c *Client) UpdateAccountingReport(ar *AccountingReport) error {
	return c.UpdateAccountingReports([]int64{ar.Id.Get()}, ar)
}

// UpdateAccountingReports updates existing accounting.report records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAccountingReports(ids []int64, ar *AccountingReport) error {
	return c.Update(AccountingReportModel, ids, ar, nil)
}

// DeleteAccountingReport deletes an existing accounting.report record.
func (c *Client) DeleteAccountingReport(id int64) error {
	return c.DeleteAccountingReports([]int64{id})
}

// DeleteAccountingReports deletes existing accounting.report records.
func (c *Client) DeleteAccountingReports(ids []int64) error {
	return c.Delete(AccountingReportModel, ids)
}

// GetAccountingReport gets accounting.report existing record.
func (c *Client) GetAccountingReport(id int64) (*AccountingReport, error) {
	ars, err := c.GetAccountingReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// GetAccountingReports gets accounting.report existing records.
func (c *Client) GetAccountingReports(ids []int64) (*AccountingReports, error) {
	ars := &AccountingReports{}
	if err := c.Read(AccountingReportModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountingReport finds accounting.report record by querying it with criteria.
func (c *Client) FindAccountingReport(criteria *Criteria) (*AccountingReport, error) {
	ars := &AccountingReports{}
	if err := c.SearchRead(AccountingReportModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// FindAccountingReports finds accounting.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountingReports(criteria *Criteria, options *Options) (*AccountingReports, error) {
	ars := &AccountingReports{}
	if err := c.SearchRead(AccountingReportModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountingReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountingReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountingReportModel, criteria, options)
}

// FindAccountingReportId finds record id by querying it with criteria.
func (c *Client) FindAccountingReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountingReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
