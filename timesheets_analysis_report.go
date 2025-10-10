package odoo

// TimesheetsAnalysisReport represents timesheets.analysis.report model.
type TimesheetsAnalysisReport struct {
	Amount                     *Float     `xmlrpc:"amount,omitempty"`
	BillableTime               *Float     `xmlrpc:"billable_time,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                       *Time      `xmlrpc:"date,omitempty"`
	DepartmentId               *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                 *Many2One  `xmlrpc:"employee_id,omitempty"`
	HasDepartmentManagerAccess *Bool      `xmlrpc:"has_department_manager_access,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	ManagerId                  *Many2One  `xmlrpc:"manager_id,omitempty"`
	Margin                     *Float     `xmlrpc:"margin,omitempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneId                *Many2One  `xmlrpc:"milestone_id,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NonBillableTime            *Float     `xmlrpc:"non_billable_time,omitempty"`
	OrderId                    *Many2One  `xmlrpc:"order_id,omitempty"`
	ParentTaskId               *Many2One  `xmlrpc:"parent_task_id,omitempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omitempty"`
	ProjectId                  *Many2One  `xmlrpc:"project_id,omitempty"`
	SoLine                     *Many2One  `xmlrpc:"so_line,omitempty"`
	TaskId                     *Many2One  `xmlrpc:"task_id,omitempty"`
	TimesheetInvoiceId         *Many2One  `xmlrpc:"timesheet_invoice_id,omitempty"`
	TimesheetInvoiceType       *Selection `xmlrpc:"timesheet_invoice_type,omitempty"`
	TimesheetRevenues          *Float     `xmlrpc:"timesheet_revenues,omitempty"`
	UnitAmount                 *Float     `xmlrpc:"unit_amount,omitempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omitempty"`
}

// TimesheetsAnalysisReports represents array of timesheets.analysis.report model.
type TimesheetsAnalysisReports []TimesheetsAnalysisReport

// TimesheetsAnalysisReportModel is the odoo model name.
const TimesheetsAnalysisReportModel = "timesheets.analysis.report"

// Many2One convert TimesheetsAnalysisReport to *Many2One.
func (tar *TimesheetsAnalysisReport) Many2One() *Many2One {
	return NewMany2One(tar.Id.Get(), "")
}

// CreateTimesheetsAnalysisReport creates a new timesheets.analysis.report model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReport(tar *TimesheetsAnalysisReport) (int64, error) {
	ids, err := c.CreateTimesheetsAnalysisReports([]*TimesheetsAnalysisReport{tar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimesheetsAnalysisReport creates a new timesheets.analysis.report model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReports(tars []*TimesheetsAnalysisReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range tars {
		vv = append(vv, v)
	}
	return c.Create(TimesheetsAnalysisReportModel, vv, nil)
}

// UpdateTimesheetsAnalysisReport updates an existing timesheets.analysis.report record.
func (c *Client) UpdateTimesheetsAnalysisReport(tar *TimesheetsAnalysisReport) error {
	return c.UpdateTimesheetsAnalysisReports([]int64{tar.Id.Get()}, tar)
}

// UpdateTimesheetsAnalysisReports updates existing timesheets.analysis.report records.
// All records (represented by ids) will be updated by tar values.
func (c *Client) UpdateTimesheetsAnalysisReports(ids []int64, tar *TimesheetsAnalysisReport) error {
	return c.Update(TimesheetsAnalysisReportModel, ids, tar, nil)
}

// DeleteTimesheetsAnalysisReport deletes an existing timesheets.analysis.report record.
func (c *Client) DeleteTimesheetsAnalysisReport(id int64) error {
	return c.DeleteTimesheetsAnalysisReports([]int64{id})
}

// DeleteTimesheetsAnalysisReports deletes existing timesheets.analysis.report records.
func (c *Client) DeleteTimesheetsAnalysisReports(ids []int64) error {
	return c.Delete(TimesheetsAnalysisReportModel, ids)
}

// GetTimesheetsAnalysisReport gets timesheets.analysis.report existing record.
func (c *Client) GetTimesheetsAnalysisReport(id int64) (*TimesheetsAnalysisReport, error) {
	tars, err := c.GetTimesheetsAnalysisReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tars)[0]), nil
}

// GetTimesheetsAnalysisReports gets timesheets.analysis.report existing records.
func (c *Client) GetTimesheetsAnalysisReports(ids []int64) (*TimesheetsAnalysisReports, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.Read(TimesheetsAnalysisReportModel, ids, nil, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReport finds timesheets.analysis.report record by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReport(criteria *Criteria) (*TimesheetsAnalysisReport, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.SearchRead(TimesheetsAnalysisReportModel, criteria, NewOptions().Limit(1), tars); err != nil {
		return nil, err
	}
	return &((*tars)[0]), nil
}

// FindTimesheetsAnalysisReports finds timesheets.analysis.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReports(criteria *Criteria, options *Options) (*TimesheetsAnalysisReports, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.SearchRead(TimesheetsAnalysisReportModel, criteria, options, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TimesheetsAnalysisReportModel, criteria, options)
}

// FindTimesheetsAnalysisReportId finds record id by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimesheetsAnalysisReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
