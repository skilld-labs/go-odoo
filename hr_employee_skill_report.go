package odoo

// HrEmployeeSkillReport represents hr.employee.skill.report model.
type HrEmployeeSkillReport struct {
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

// HrEmployeeSkillReports represents array of hr.employee.skill.report model.
type HrEmployeeSkillReports []HrEmployeeSkillReport

// HrEmployeeSkillReportModel is the odoo model name.
const HrEmployeeSkillReportModel = "hr.employee.skill.report"

// Many2One convert HrEmployeeSkillReport to *Many2One.
func (hesr *HrEmployeeSkillReport) Many2One() *Many2One {
	return NewMany2One(hesr.Id.Get(), "")
}

// CreateHrEmployeeSkillReport creates a new hr.employee.skill.report model and returns its id.
func (c *Client) CreateHrEmployeeSkillReport(hesr *HrEmployeeSkillReport) (int64, error) {
	ids, err := c.CreateHrEmployeeSkillReports([]*HrEmployeeSkillReport{hesr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeSkillReport creates a new hr.employee.skill.report model and returns its id.
func (c *Client) CreateHrEmployeeSkillReports(hesrs []*HrEmployeeSkillReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hesrs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeSkillReportModel, vv, nil)
}

// UpdateHrEmployeeSkillReport updates an existing hr.employee.skill.report record.
func (c *Client) UpdateHrEmployeeSkillReport(hesr *HrEmployeeSkillReport) error {
	return c.UpdateHrEmployeeSkillReports([]int64{hesr.Id.Get()}, hesr)
}

// UpdateHrEmployeeSkillReports updates existing hr.employee.skill.report records.
// All records (represented by ids) will be updated by hesr values.
func (c *Client) UpdateHrEmployeeSkillReports(ids []int64, hesr *HrEmployeeSkillReport) error {
	return c.Update(HrEmployeeSkillReportModel, ids, hesr, nil)
}

// DeleteHrEmployeeSkillReport deletes an existing hr.employee.skill.report record.
func (c *Client) DeleteHrEmployeeSkillReport(id int64) error {
	return c.DeleteHrEmployeeSkillReports([]int64{id})
}

// DeleteHrEmployeeSkillReports deletes existing hr.employee.skill.report records.
func (c *Client) DeleteHrEmployeeSkillReports(ids []int64) error {
	return c.Delete(HrEmployeeSkillReportModel, ids)
}

// GetHrEmployeeSkillReport gets hr.employee.skill.report existing record.
func (c *Client) GetHrEmployeeSkillReport(id int64) (*HrEmployeeSkillReport, error) {
	hesrs, err := c.GetHrEmployeeSkillReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hesrs)[0]), nil
}

// GetHrEmployeeSkillReports gets hr.employee.skill.report existing records.
func (c *Client) GetHrEmployeeSkillReports(ids []int64) (*HrEmployeeSkillReports, error) {
	hesrs := &HrEmployeeSkillReports{}
	if err := c.Read(HrEmployeeSkillReportModel, ids, nil, hesrs); err != nil {
		return nil, err
	}
	return hesrs, nil
}

// FindHrEmployeeSkillReport finds hr.employee.skill.report record by querying it with criteria.
func (c *Client) FindHrEmployeeSkillReport(criteria *Criteria) (*HrEmployeeSkillReport, error) {
	hesrs := &HrEmployeeSkillReports{}
	if err := c.SearchRead(HrEmployeeSkillReportModel, criteria, NewOptions().Limit(1), hesrs); err != nil {
		return nil, err
	}
	return &((*hesrs)[0]), nil
}

// FindHrEmployeeSkillReports finds hr.employee.skill.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillReports(criteria *Criteria, options *Options) (*HrEmployeeSkillReports, error) {
	hesrs := &HrEmployeeSkillReports{}
	if err := c.SearchRead(HrEmployeeSkillReportModel, criteria, options, hesrs); err != nil {
		return nil, err
	}
	return hesrs, nil
}

// FindHrEmployeeSkillReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeSkillReportModel, criteria, options)
}

// FindHrEmployeeSkillReportId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeSkillReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeSkillReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
