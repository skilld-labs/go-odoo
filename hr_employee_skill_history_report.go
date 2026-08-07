package odoo

// HrEmployeeSkillHistoryReport represents hr.employee.skill.history.report model.
type HrEmployeeSkillHistoryReport struct {
	Date          *Time     `xmlrpc:"date,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LevelProgress *Float    `xmlrpc:"level_progress,omitempty"`
	SkillId       *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omitempty"`
}

// HrEmployeeSkillHistoryReports represents array of hr.employee.skill.history.report model.
type HrEmployeeSkillHistoryReports []HrEmployeeSkillHistoryReport

// HrEmployeeSkillHistoryReportModel is the odoo model name.
const HrEmployeeSkillHistoryReportModel = "hr.employee.skill.history.report"

// Many2One convert HrEmployeeSkillHistoryReport to *Many2One.
func (heshr *HrEmployeeSkillHistoryReport) Many2One() *Many2One {
	return NewMany2One(heshr.Id.Get(), "")
}

// CreateHrEmployeeSkillHistoryReport creates a new hr.employee.skill.history.report model and returns its id.
func (c *Client) CreateHrEmployeeSkillHistoryReport(heshr *HrEmployeeSkillHistoryReport) (int64, error) {
	ids, err := c.CreateHrEmployeeSkillHistoryReports([]*HrEmployeeSkillHistoryReport{heshr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeSkillHistoryReport creates a new hr.employee.skill.history.report model and returns its id.
func (c *Client) CreateHrEmployeeSkillHistoryReports(heshrs []*HrEmployeeSkillHistoryReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range heshrs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeSkillHistoryReportModel, vv, nil)
}

// UpdateHrEmployeeSkillHistoryReport updates an existing hr.employee.skill.history.report record.
func (c *Client) UpdateHrEmployeeSkillHistoryReport(heshr *HrEmployeeSkillHistoryReport) error {
	return c.UpdateHrEmployeeSkillHistoryReports([]int64{heshr.Id.Get()}, heshr)
}

// UpdateHrEmployeeSkillHistoryReports updates existing hr.employee.skill.history.report records.
// All records (represented by ids) will be updated by heshr values.
func (c *Client) UpdateHrEmployeeSkillHistoryReports(ids []int64, heshr *HrEmployeeSkillHistoryReport) error {
	return c.Update(HrEmployeeSkillHistoryReportModel, ids, heshr, nil)
}

// DeleteHrEmployeeSkillHistoryReport deletes an existing hr.employee.skill.history.report record.
func (c *Client) DeleteHrEmployeeSkillHistoryReport(id int64) error {
	return c.DeleteHrEmployeeSkillHistoryReports([]int64{id})
}

// DeleteHrEmployeeSkillHistoryReports deletes existing hr.employee.skill.history.report records.
func (c *Client) DeleteHrEmployeeSkillHistoryReports(ids []int64) error {
	return c.Delete(HrEmployeeSkillHistoryReportModel, ids)
}

// GetHrEmployeeSkillHistoryReport gets hr.employee.skill.history.report existing record.
func (c *Client) GetHrEmployeeSkillHistoryReport(id int64) (*HrEmployeeSkillHistoryReport, error) {
	heshrs, err := c.GetHrEmployeeSkillHistoryReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*heshrs)[0]), nil
}

// GetHrEmployeeSkillHistoryReports gets hr.employee.skill.history.report existing records.
func (c *Client) GetHrEmployeeSkillHistoryReports(ids []int64) (*HrEmployeeSkillHistoryReports, error) {
	heshrs := &HrEmployeeSkillHistoryReports{}
	if err := c.Read(HrEmployeeSkillHistoryReportModel, ids, nil, heshrs); err != nil {
		return nil, err
	}
	return heshrs, nil
}

// FindHrEmployeeSkillHistoryReport finds hr.employee.skill.history.report record by querying it with criteria.
func (c *Client) FindHrEmployeeSkillHistoryReport(criteria *Criteria) (*HrEmployeeSkillHistoryReport, error) {
	heshrs := &HrEmployeeSkillHistoryReports{}
	if err := c.SearchRead(HrEmployeeSkillHistoryReportModel, criteria, NewOptions().Limit(1), heshrs); err != nil {
		return nil, err
	}
	return &((*heshrs)[0]), nil
}

// FindHrEmployeeSkillHistoryReports finds hr.employee.skill.history.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillHistoryReports(criteria *Criteria, options *Options) (*HrEmployeeSkillHistoryReports, error) {
	heshrs := &HrEmployeeSkillHistoryReports{}
	if err := c.SearchRead(HrEmployeeSkillHistoryReportModel, criteria, options, heshrs); err != nil {
		return nil, err
	}
	return heshrs, nil
}

// FindHrEmployeeSkillHistoryReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillHistoryReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeSkillHistoryReportModel, criteria, options)
}

// FindHrEmployeeSkillHistoryReportId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeSkillHistoryReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeSkillHistoryReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
