package odoo

// ReportAccountReportFinancial represents report.account.report_financial model.
type ReportAccountReportFinancial struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportFinancials represents array of report.account.report_financial model.
type ReportAccountReportFinancials []ReportAccountReportFinancial

// ReportAccountReportFinancialModel is the odoo model name.
const ReportAccountReportFinancialModel = "report.account.report_financial"

// Many2One convert ReportAccountReportFinancial to *Many2One.
func (rar *ReportAccountReportFinancial) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportFinancial creates a new report.account.report_financial model and returns its id.
func (c *Client) CreateReportAccountReportFinancial(rar *ReportAccountReportFinancial) (int64, error) {
	ids, err := c.CreateReportAccountReportFinancials([]*ReportAccountReportFinancial{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportFinancial creates a new report.account.report_financial model and returns its id.
func (c *Client) CreateReportAccountReportFinancials(rars []*ReportAccountReportFinancial) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportFinancialModel, vv, nil)
}

// UpdateReportAccountReportFinancial updates an existing report.account.report_financial record.
func (c *Client) UpdateReportAccountReportFinancial(rar *ReportAccountReportFinancial) error {
	return c.UpdateReportAccountReportFinancials([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportFinancials updates existing report.account.report_financial records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportFinancials(ids []int64, rar *ReportAccountReportFinancial) error {
	return c.Update(ReportAccountReportFinancialModel, ids, rar, nil)
}

// DeleteReportAccountReportFinancial deletes an existing report.account.report_financial record.
func (c *Client) DeleteReportAccountReportFinancial(id int64) error {
	return c.DeleteReportAccountReportFinancials([]int64{id})
}

// DeleteReportAccountReportFinancials deletes existing report.account.report_financial records.
func (c *Client) DeleteReportAccountReportFinancials(ids []int64) error {
	return c.Delete(ReportAccountReportFinancialModel, ids)
}

// GetReportAccountReportFinancial gets report.account.report_financial existing record.
func (c *Client) GetReportAccountReportFinancial(id int64) (*ReportAccountReportFinancial, error) {
	rars, err := c.GetReportAccountReportFinancials([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// GetReportAccountReportFinancials gets report.account.report_financial existing records.
func (c *Client) GetReportAccountReportFinancials(ids []int64) (*ReportAccountReportFinancials, error) {
	rars := &ReportAccountReportFinancials{}
	if err := c.Read(ReportAccountReportFinancialModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportFinancial finds report.account.report_financial record by querying it with criteria.
func (c *Client) FindReportAccountReportFinancial(criteria *Criteria) (*ReportAccountReportFinancial, error) {
	rars := &ReportAccountReportFinancials{}
	if err := c.SearchRead(ReportAccountReportFinancialModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// FindReportAccountReportFinancials finds report.account.report_financial records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportFinancials(criteria *Criteria, options *Options) (*ReportAccountReportFinancials, error) {
	rars := &ReportAccountReportFinancials{}
	if err := c.SearchRead(ReportAccountReportFinancialModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportFinancialIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportFinancialIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAccountReportFinancialModel, criteria, options)
}

// FindReportAccountReportFinancialId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportFinancialId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportFinancialModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
