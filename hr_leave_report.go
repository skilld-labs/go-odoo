package odoo

// HrLeaveReport represents hr.leave.report model.
type HrLeaveReport struct {
	AllocationId               *Many2One  `xmlrpc:"allocation_id,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	DateFrom                   *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                     *Time      `xmlrpc:"date_to,omitempty"`
	DepartmentId               *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                 *Many2One  `xmlrpc:"employee_id,omitempty"`
	HasDepartmentManagerAccess *Bool      `xmlrpc:"has_department_manager_access,omitempty"`
	HolidayStatusId            *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	LeaveId                    *Many2One  `xmlrpc:"leave_id,omitempty"`
	LeaveType                  *Selection `xmlrpc:"leave_type,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NumberOfDays               *Float     `xmlrpc:"number_of_days,omitempty"`
	NumberOfHours              *Float     `xmlrpc:"number_of_hours,omitempty"`
	State                      *Selection `xmlrpc:"state,omitempty"`
}

// HrLeaveReports represents array of hr.leave.report model.
type HrLeaveReports []HrLeaveReport

// HrLeaveReportModel is the odoo model name.
const HrLeaveReportModel = "hr.leave.report"

// Many2One convert HrLeaveReport to *Many2One.
func (hlr *HrLeaveReport) Many2One() *Many2One {
	return NewMany2One(hlr.Id.Get(), "")
}

// CreateHrLeaveReport creates a new hr.leave.report model and returns its id.
func (c *Client) CreateHrLeaveReport(hlr *HrLeaveReport) (int64, error) {
	ids, err := c.CreateHrLeaveReports([]*HrLeaveReport{hlr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveReport creates a new hr.leave.report model and returns its id.
func (c *Client) CreateHrLeaveReports(hlrs []*HrLeaveReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlrs {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveReportModel, vv, nil)
}

// UpdateHrLeaveReport updates an existing hr.leave.report record.
func (c *Client) UpdateHrLeaveReport(hlr *HrLeaveReport) error {
	return c.UpdateHrLeaveReports([]int64{hlr.Id.Get()}, hlr)
}

// UpdateHrLeaveReports updates existing hr.leave.report records.
// All records (represented by ids) will be updated by hlr values.
func (c *Client) UpdateHrLeaveReports(ids []int64, hlr *HrLeaveReport) error {
	return c.Update(HrLeaveReportModel, ids, hlr, nil)
}

// DeleteHrLeaveReport deletes an existing hr.leave.report record.
func (c *Client) DeleteHrLeaveReport(id int64) error {
	return c.DeleteHrLeaveReports([]int64{id})
}

// DeleteHrLeaveReports deletes existing hr.leave.report records.
func (c *Client) DeleteHrLeaveReports(ids []int64) error {
	return c.Delete(HrLeaveReportModel, ids)
}

// GetHrLeaveReport gets hr.leave.report existing record.
func (c *Client) GetHrLeaveReport(id int64) (*HrLeaveReport, error) {
	hlrs, err := c.GetHrLeaveReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlrs)[0]), nil
}

// GetHrLeaveReports gets hr.leave.report existing records.
func (c *Client) GetHrLeaveReports(ids []int64) (*HrLeaveReports, error) {
	hlrs := &HrLeaveReports{}
	if err := c.Read(HrLeaveReportModel, ids, nil, hlrs); err != nil {
		return nil, err
	}
	return hlrs, nil
}

// FindHrLeaveReport finds hr.leave.report record by querying it with criteria.
func (c *Client) FindHrLeaveReport(criteria *Criteria) (*HrLeaveReport, error) {
	hlrs := &HrLeaveReports{}
	if err := c.SearchRead(HrLeaveReportModel, criteria, NewOptions().Limit(1), hlrs); err != nil {
		return nil, err
	}
	return &((*hlrs)[0]), nil
}

// FindHrLeaveReports finds hr.leave.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReports(criteria *Criteria, options *Options) (*HrLeaveReports, error) {
	hlrs := &HrLeaveReports{}
	if err := c.SearchRead(HrLeaveReportModel, criteria, options, hlrs); err != nil {
		return nil, err
	}
	return hlrs, nil
}

// FindHrLeaveReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveReportModel, criteria, options)
}

// FindHrLeaveReportId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
