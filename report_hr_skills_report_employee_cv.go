package odoo

// ReportHrSkillsReportEmployeeCv represents report.hr_skills.report_employee_cv model.
type ReportHrSkillsReportEmployeeCv struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportHrSkillsReportEmployeeCvs represents array of report.hr_skills.report_employee_cv model.
type ReportHrSkillsReportEmployeeCvs []ReportHrSkillsReportEmployeeCv

// ReportHrSkillsReportEmployeeCvModel is the odoo model name.
const ReportHrSkillsReportEmployeeCvModel = "report.hr_skills.report_employee_cv"

// Many2One convert ReportHrSkillsReportEmployeeCv to *Many2One.
func (rhr *ReportHrSkillsReportEmployeeCv) Many2One() *Many2One {
	return NewMany2One(rhr.Id.Get(), "")
}

// CreateReportHrSkillsReportEmployeeCv creates a new report.hr_skills.report_employee_cv model and returns its id.
func (c *Client) CreateReportHrSkillsReportEmployeeCv(rhr *ReportHrSkillsReportEmployeeCv) (int64, error) {
	ids, err := c.CreateReportHrSkillsReportEmployeeCvs([]*ReportHrSkillsReportEmployeeCv{rhr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportHrSkillsReportEmployeeCv creates a new report.hr_skills.report_employee_cv model and returns its id.
func (c *Client) CreateReportHrSkillsReportEmployeeCvs(rhrs []*ReportHrSkillsReportEmployeeCv) ([]int64, error) {
	var vv []interface{}
	for _, v := range rhrs {
		vv = append(vv, v)
	}
	return c.Create(ReportHrSkillsReportEmployeeCvModel, vv, nil)
}

// UpdateReportHrSkillsReportEmployeeCv updates an existing report.hr_skills.report_employee_cv record.
func (c *Client) UpdateReportHrSkillsReportEmployeeCv(rhr *ReportHrSkillsReportEmployeeCv) error {
	return c.UpdateReportHrSkillsReportEmployeeCvs([]int64{rhr.Id.Get()}, rhr)
}

// UpdateReportHrSkillsReportEmployeeCvs updates existing report.hr_skills.report_employee_cv records.
// All records (represented by ids) will be updated by rhr values.
func (c *Client) UpdateReportHrSkillsReportEmployeeCvs(ids []int64, rhr *ReportHrSkillsReportEmployeeCv) error {
	return c.Update(ReportHrSkillsReportEmployeeCvModel, ids, rhr, nil)
}

// DeleteReportHrSkillsReportEmployeeCv deletes an existing report.hr_skills.report_employee_cv record.
func (c *Client) DeleteReportHrSkillsReportEmployeeCv(id int64) error {
	return c.DeleteReportHrSkillsReportEmployeeCvs([]int64{id})
}

// DeleteReportHrSkillsReportEmployeeCvs deletes existing report.hr_skills.report_employee_cv records.
func (c *Client) DeleteReportHrSkillsReportEmployeeCvs(ids []int64) error {
	return c.Delete(ReportHrSkillsReportEmployeeCvModel, ids)
}

// GetReportHrSkillsReportEmployeeCv gets report.hr_skills.report_employee_cv existing record.
func (c *Client) GetReportHrSkillsReportEmployeeCv(id int64) (*ReportHrSkillsReportEmployeeCv, error) {
	rhrs, err := c.GetReportHrSkillsReportEmployeeCvs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rhrs)[0]), nil
}

// GetReportHrSkillsReportEmployeeCvs gets report.hr_skills.report_employee_cv existing records.
func (c *Client) GetReportHrSkillsReportEmployeeCvs(ids []int64) (*ReportHrSkillsReportEmployeeCvs, error) {
	rhrs := &ReportHrSkillsReportEmployeeCvs{}
	if err := c.Read(ReportHrSkillsReportEmployeeCvModel, ids, nil, rhrs); err != nil {
		return nil, err
	}
	return rhrs, nil
}

// FindReportHrSkillsReportEmployeeCv finds report.hr_skills.report_employee_cv record by querying it with criteria.
func (c *Client) FindReportHrSkillsReportEmployeeCv(criteria *Criteria) (*ReportHrSkillsReportEmployeeCv, error) {
	rhrs := &ReportHrSkillsReportEmployeeCvs{}
	if err := c.SearchRead(ReportHrSkillsReportEmployeeCvModel, criteria, NewOptions().Limit(1), rhrs); err != nil {
		return nil, err
	}
	return &((*rhrs)[0]), nil
}

// FindReportHrSkillsReportEmployeeCvs finds report.hr_skills.report_employee_cv records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportHrSkillsReportEmployeeCvs(criteria *Criteria, options *Options) (*ReportHrSkillsReportEmployeeCvs, error) {
	rhrs := &ReportHrSkillsReportEmployeeCvs{}
	if err := c.SearchRead(ReportHrSkillsReportEmployeeCvModel, criteria, options, rhrs); err != nil {
		return nil, err
	}
	return rhrs, nil
}

// FindReportHrSkillsReportEmployeeCvIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportHrSkillsReportEmployeeCvIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportHrSkillsReportEmployeeCvModel, criteria, options)
}

// FindReportHrSkillsReportEmployeeCvId finds record id by querying it with criteria.
func (c *Client) FindReportHrSkillsReportEmployeeCvId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportHrSkillsReportEmployeeCvModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
