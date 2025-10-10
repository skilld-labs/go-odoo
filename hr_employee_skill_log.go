package odoo

// HrEmployeeSkillLog represents hr.employee.skill.log model.
type HrEmployeeSkillLog struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	Date          *Time     `xmlrpc:"date,omitempty"`
	DepartmentId  *Many2One `xmlrpc:"department_id,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omitempty"`
	SkillId       *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillLevelId  *Many2One `xmlrpc:"skill_level_id,omitempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeeSkillLogs represents array of hr.employee.skill.log model.
type HrEmployeeSkillLogs []HrEmployeeSkillLog

// HrEmployeeSkillLogModel is the odoo model name.
const HrEmployeeSkillLogModel = "hr.employee.skill.log"

// Many2One convert HrEmployeeSkillLog to *Many2One.
func (hesl *HrEmployeeSkillLog) Many2One() *Many2One {
	return NewMany2One(hesl.Id.Get(), "")
}

// CreateHrEmployeeSkillLog creates a new hr.employee.skill.log model and returns its id.
func (c *Client) CreateHrEmployeeSkillLog(hesl *HrEmployeeSkillLog) (int64, error) {
	ids, err := c.CreateHrEmployeeSkillLogs([]*HrEmployeeSkillLog{hesl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeSkillLog creates a new hr.employee.skill.log model and returns its id.
func (c *Client) CreateHrEmployeeSkillLogs(hesls []*HrEmployeeSkillLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range hesls {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeSkillLogModel, vv, nil)
}

// UpdateHrEmployeeSkillLog updates an existing hr.employee.skill.log record.
func (c *Client) UpdateHrEmployeeSkillLog(hesl *HrEmployeeSkillLog) error {
	return c.UpdateHrEmployeeSkillLogs([]int64{hesl.Id.Get()}, hesl)
}

// UpdateHrEmployeeSkillLogs updates existing hr.employee.skill.log records.
// All records (represented by ids) will be updated by hesl values.
func (c *Client) UpdateHrEmployeeSkillLogs(ids []int64, hesl *HrEmployeeSkillLog) error {
	return c.Update(HrEmployeeSkillLogModel, ids, hesl, nil)
}

// DeleteHrEmployeeSkillLog deletes an existing hr.employee.skill.log record.
func (c *Client) DeleteHrEmployeeSkillLog(id int64) error {
	return c.DeleteHrEmployeeSkillLogs([]int64{id})
}

// DeleteHrEmployeeSkillLogs deletes existing hr.employee.skill.log records.
func (c *Client) DeleteHrEmployeeSkillLogs(ids []int64) error {
	return c.Delete(HrEmployeeSkillLogModel, ids)
}

// GetHrEmployeeSkillLog gets hr.employee.skill.log existing record.
func (c *Client) GetHrEmployeeSkillLog(id int64) (*HrEmployeeSkillLog, error) {
	hesls, err := c.GetHrEmployeeSkillLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hesls)[0]), nil
}

// GetHrEmployeeSkillLogs gets hr.employee.skill.log existing records.
func (c *Client) GetHrEmployeeSkillLogs(ids []int64) (*HrEmployeeSkillLogs, error) {
	hesls := &HrEmployeeSkillLogs{}
	if err := c.Read(HrEmployeeSkillLogModel, ids, nil, hesls); err != nil {
		return nil, err
	}
	return hesls, nil
}

// FindHrEmployeeSkillLog finds hr.employee.skill.log record by querying it with criteria.
func (c *Client) FindHrEmployeeSkillLog(criteria *Criteria) (*HrEmployeeSkillLog, error) {
	hesls := &HrEmployeeSkillLogs{}
	if err := c.SearchRead(HrEmployeeSkillLogModel, criteria, NewOptions().Limit(1), hesls); err != nil {
		return nil, err
	}
	return &((*hesls)[0]), nil
}

// FindHrEmployeeSkillLogs finds hr.employee.skill.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillLogs(criteria *Criteria, options *Options) (*HrEmployeeSkillLogs, error) {
	hesls := &HrEmployeeSkillLogs{}
	if err := c.SearchRead(HrEmployeeSkillLogModel, criteria, options, hesls); err != nil {
		return nil, err
	}
	return hesls, nil
}

// FindHrEmployeeSkillLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeSkillLogModel, criteria, options)
}

// FindHrEmployeeSkillLogId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeSkillLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeSkillLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
