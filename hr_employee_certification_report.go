package odoo

// HrEmployeeCertificationReport represents hr.employee.certification.report model.
type HrEmployeeCertificationReport struct {
	Active                     *Bool     `xmlrpc:"active,omitempty"`
	CompanyId                  *Many2One `xmlrpc:"company_id,omitempty"`
	DepartmentId               *Many2One `xmlrpc:"department_id,omitempty"`
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId                 *Many2One `xmlrpc:"employee_id,omitempty"`
	HasDepartmentManagerAccess *Bool     `xmlrpc:"has_department_manager_access,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
	LevelProgress              *Float    `xmlrpc:"level_progress,omitempty"`
	SkillId                    *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillLevel                 *String   `xmlrpc:"skill_level,omitempty"`
	SkillTypeId                *Many2One `xmlrpc:"skill_type_id,omitempty"`
}

// HrEmployeeCertificationReports represents array of hr.employee.certification.report model.
type HrEmployeeCertificationReports []HrEmployeeCertificationReport

// HrEmployeeCertificationReportModel is the odoo model name.
const HrEmployeeCertificationReportModel = "hr.employee.certification.report"

// Many2One convert HrEmployeeCertificationReport to *Many2One.
func (hecr *HrEmployeeCertificationReport) Many2One() *Many2One {
	return NewMany2One(hecr.Id.Get(), "")
}

// CreateHrEmployeeCertificationReport creates a new hr.employee.certification.report model and returns its id.
func (c *Client) CreateHrEmployeeCertificationReport(hecr *HrEmployeeCertificationReport) (int64, error) {
	ids, err := c.CreateHrEmployeeCertificationReports([]*HrEmployeeCertificationReport{hecr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeCertificationReport creates a new hr.employee.certification.report model and returns its id.
func (c *Client) CreateHrEmployeeCertificationReports(hecrs []*HrEmployeeCertificationReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hecrs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeCertificationReportModel, vv, nil)
}

// UpdateHrEmployeeCertificationReport updates an existing hr.employee.certification.report record.
func (c *Client) UpdateHrEmployeeCertificationReport(hecr *HrEmployeeCertificationReport) error {
	return c.UpdateHrEmployeeCertificationReports([]int64{hecr.Id.Get()}, hecr)
}

// UpdateHrEmployeeCertificationReports updates existing hr.employee.certification.report records.
// All records (represented by ids) will be updated by hecr values.
func (c *Client) UpdateHrEmployeeCertificationReports(ids []int64, hecr *HrEmployeeCertificationReport) error {
	return c.Update(HrEmployeeCertificationReportModel, ids, hecr, nil)
}

// DeleteHrEmployeeCertificationReport deletes an existing hr.employee.certification.report record.
func (c *Client) DeleteHrEmployeeCertificationReport(id int64) error {
	return c.DeleteHrEmployeeCertificationReports([]int64{id})
}

// DeleteHrEmployeeCertificationReports deletes existing hr.employee.certification.report records.
func (c *Client) DeleteHrEmployeeCertificationReports(ids []int64) error {
	return c.Delete(HrEmployeeCertificationReportModel, ids)
}

// GetHrEmployeeCertificationReport gets hr.employee.certification.report existing record.
func (c *Client) GetHrEmployeeCertificationReport(id int64) (*HrEmployeeCertificationReport, error) {
	hecrs, err := c.GetHrEmployeeCertificationReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hecrs)[0]), nil
}

// GetHrEmployeeCertificationReports gets hr.employee.certification.report existing records.
func (c *Client) GetHrEmployeeCertificationReports(ids []int64) (*HrEmployeeCertificationReports, error) {
	hecrs := &HrEmployeeCertificationReports{}
	if err := c.Read(HrEmployeeCertificationReportModel, ids, nil, hecrs); err != nil {
		return nil, err
	}
	return hecrs, nil
}

// FindHrEmployeeCertificationReport finds hr.employee.certification.report record by querying it with criteria.
func (c *Client) FindHrEmployeeCertificationReport(criteria *Criteria) (*HrEmployeeCertificationReport, error) {
	hecrs := &HrEmployeeCertificationReports{}
	if err := c.SearchRead(HrEmployeeCertificationReportModel, criteria, NewOptions().Limit(1), hecrs); err != nil {
		return nil, err
	}
	return &((*hecrs)[0]), nil
}

// FindHrEmployeeCertificationReports finds hr.employee.certification.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCertificationReports(criteria *Criteria, options *Options) (*HrEmployeeCertificationReports, error) {
	hecrs := &HrEmployeeCertificationReports{}
	if err := c.SearchRead(HrEmployeeCertificationReportModel, criteria, options, hecrs); err != nil {
		return nil, err
	}
	return hecrs, nil
}

// FindHrEmployeeCertificationReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCertificationReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeCertificationReportModel, criteria, options)
}

// FindHrEmployeeCertificationReportId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeCertificationReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeCertificationReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
