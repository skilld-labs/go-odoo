package odoo

// AccountFinancialReport represents account.financial.report model.
type AccountFinancialReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	AccountIds      *Relation  `xmlrpc:"account_ids,omitempty"`
	AccountReportId *Many2One  `xmlrpc:"account_report_id,omitempty"`
	AccountTypeIds  *Relation  `xmlrpc:"account_type_ids,omitempty"`
	ChildrenIds     *Relation  `xmlrpc:"children_ids,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayDetail   *Selection `xmlrpc:"display_detail,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Level           *Int       `xmlrpc:"level,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	ParentId        *Many2One  `xmlrpc:"parent_id,omitempty"`
	Sequence        *Int       `xmlrpc:"sequence,omitempty"`
	Sign            *Selection `xmlrpc:"sign,omitempty"`
	StyleOverwrite  *Selection `xmlrpc:"style_overwrite,omitempty"`
	Type            *Selection `xmlrpc:"type,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFinancialReports represents array of account.financial.report model.
type AccountFinancialReports []AccountFinancialReport

// AccountFinancialReportModel is the odoo model name.
const AccountFinancialReportModel = "account.financial.report"

// Many2One convert AccountFinancialReport to *Many2One.
func (afr *AccountFinancialReport) Many2One() *Many2One {
	return NewMany2One(afr.Id.Get(), "")
}

// CreateAccountFinancialReport creates a new account.financial.report model and returns its id.
func (c *Client) CreateAccountFinancialReport(afr *AccountFinancialReport) (int64, error) {
	ids, err := c.CreateAccountFinancialReports([]*AccountFinancialReport{afr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFinancialReports creates a new account.financial.report model and returns its id.
func (c *Client) CreateAccountFinancialReports(afrs []*AccountFinancialReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range afrs {
		vv = append(vv, v)
	}
	return c.Create(AccountFinancialReportModel, vv, nil)
}

// UpdateAccountFinancialReport updates an existing account.financial.report record.
func (c *Client) UpdateAccountFinancialReport(afr *AccountFinancialReport) error {
	return c.UpdateAccountFinancialReports([]int64{afr.Id.Get()}, afr)
}

// UpdateAccountFinancialReports updates existing account.financial.report records.
// All records (represented by ids) will be updated by afr values.
func (c *Client) UpdateAccountFinancialReports(ids []int64, afr *AccountFinancialReport) error {
	return c.Update(AccountFinancialReportModel, ids, afr, nil)
}

// DeleteAccountFinancialReport deletes an existing account.financial.report record.
func (c *Client) DeleteAccountFinancialReport(id int64) error {
	return c.DeleteAccountFinancialReports([]int64{id})
}

// DeleteAccountFinancialReports deletes existing account.financial.report records.
func (c *Client) DeleteAccountFinancialReports(ids []int64) error {
	return c.Delete(AccountFinancialReportModel, ids)
}

// GetAccountFinancialReport gets account.financial.report existing record.
func (c *Client) GetAccountFinancialReport(id int64) (*AccountFinancialReport, error) {
	afrs, err := c.GetAccountFinancialReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afrs)[0]), nil
}

// GetAccountFinancialReports gets account.financial.report existing records.
func (c *Client) GetAccountFinancialReports(ids []int64) (*AccountFinancialReports, error) {
	afrs := &AccountFinancialReports{}
	if err := c.Read(AccountFinancialReportModel, ids, nil, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFinancialReport finds account.financial.report record by querying it with criteria.
func (c *Client) FindAccountFinancialReport(criteria *Criteria) (*AccountFinancialReport, error) {
	afrs := &AccountFinancialReports{}
	if err := c.SearchRead(AccountFinancialReportModel, criteria, NewOptions().Limit(1), afrs); err != nil {
		return nil, err
	}
	return &((*afrs)[0]), nil
}

// FindAccountFinancialReports finds account.financial.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialReports(criteria *Criteria, options *Options) (*AccountFinancialReports, error) {
	afrs := &AccountFinancialReports{}
	if err := c.SearchRead(AccountFinancialReportModel, criteria, options, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFinancialReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFinancialReportModel, criteria, options)
}

// FindAccountFinancialReportId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
