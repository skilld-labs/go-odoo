package odoo

// HrManagerDepartmentReport represents hr.manager.department.report model.
type HrManagerDepartmentReport struct {
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId                 *Many2One `xmlrpc:"employee_id,omitempty"`
	HasDepartmentManagerAccess *Bool     `xmlrpc:"has_department_manager_access,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
}

// HrManagerDepartmentReports represents array of hr.manager.department.report model.
type HrManagerDepartmentReports []HrManagerDepartmentReport

// HrManagerDepartmentReportModel is the odoo model name.
const HrManagerDepartmentReportModel = "hr.manager.department.report"

// Many2One convert HrManagerDepartmentReport to *Many2One.
func (hmdr *HrManagerDepartmentReport) Many2One() *Many2One {
	return NewMany2One(hmdr.Id.Get(), "")
}

// CreateHrManagerDepartmentReport creates a new hr.manager.department.report model and returns its id.
func (c *Client) CreateHrManagerDepartmentReport(hmdr *HrManagerDepartmentReport) (int64, error) {
	ids, err := c.CreateHrManagerDepartmentReports([]*HrManagerDepartmentReport{hmdr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrManagerDepartmentReport creates a new hr.manager.department.report model and returns its id.
func (c *Client) CreateHrManagerDepartmentReports(hmdrs []*HrManagerDepartmentReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hmdrs {
		vv = append(vv, v)
	}
	return c.Create(HrManagerDepartmentReportModel, vv, nil)
}

// UpdateHrManagerDepartmentReport updates an existing hr.manager.department.report record.
func (c *Client) UpdateHrManagerDepartmentReport(hmdr *HrManagerDepartmentReport) error {
	return c.UpdateHrManagerDepartmentReports([]int64{hmdr.Id.Get()}, hmdr)
}

// UpdateHrManagerDepartmentReports updates existing hr.manager.department.report records.
// All records (represented by ids) will be updated by hmdr values.
func (c *Client) UpdateHrManagerDepartmentReports(ids []int64, hmdr *HrManagerDepartmentReport) error {
	return c.Update(HrManagerDepartmentReportModel, ids, hmdr, nil)
}

// DeleteHrManagerDepartmentReport deletes an existing hr.manager.department.report record.
func (c *Client) DeleteHrManagerDepartmentReport(id int64) error {
	return c.DeleteHrManagerDepartmentReports([]int64{id})
}

// DeleteHrManagerDepartmentReports deletes existing hr.manager.department.report records.
func (c *Client) DeleteHrManagerDepartmentReports(ids []int64) error {
	return c.Delete(HrManagerDepartmentReportModel, ids)
}

// GetHrManagerDepartmentReport gets hr.manager.department.report existing record.
func (c *Client) GetHrManagerDepartmentReport(id int64) (*HrManagerDepartmentReport, error) {
	hmdrs, err := c.GetHrManagerDepartmentReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hmdrs)[0]), nil
}

// GetHrManagerDepartmentReports gets hr.manager.department.report existing records.
func (c *Client) GetHrManagerDepartmentReports(ids []int64) (*HrManagerDepartmentReports, error) {
	hmdrs := &HrManagerDepartmentReports{}
	if err := c.Read(HrManagerDepartmentReportModel, ids, nil, hmdrs); err != nil {
		return nil, err
	}
	return hmdrs, nil
}

// FindHrManagerDepartmentReport finds hr.manager.department.report record by querying it with criteria.
func (c *Client) FindHrManagerDepartmentReport(criteria *Criteria) (*HrManagerDepartmentReport, error) {
	hmdrs := &HrManagerDepartmentReports{}
	if err := c.SearchRead(HrManagerDepartmentReportModel, criteria, NewOptions().Limit(1), hmdrs); err != nil {
		return nil, err
	}
	return &((*hmdrs)[0]), nil
}

// FindHrManagerDepartmentReports finds hr.manager.department.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrManagerDepartmentReports(criteria *Criteria, options *Options) (*HrManagerDepartmentReports, error) {
	hmdrs := &HrManagerDepartmentReports{}
	if err := c.SearchRead(HrManagerDepartmentReportModel, criteria, options, hmdrs); err != nil {
		return nil, err
	}
	return hmdrs, nil
}

// FindHrManagerDepartmentReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrManagerDepartmentReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrManagerDepartmentReportModel, criteria, options)
}

// FindHrManagerDepartmentReportId finds record id by querying it with criteria.
func (c *Client) FindHrManagerDepartmentReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrManagerDepartmentReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
