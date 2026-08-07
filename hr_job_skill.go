package odoo

// HrJobSkill represents hr.job.skill model.
type HrJobSkill struct {
	CertificationSkillTypeCount *Int      `xmlrpc:"certification_skill_type_count,omitempty"`
	Color                       *Int      `xmlrpc:"color,omitempty"`
	CreateDate                  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String   `xmlrpc:"display_name,omitempty"`
	DisplayWarningMessage       *Bool     `xmlrpc:"display_warning_message,omitempty"`
	Id                          *Int      `xmlrpc:"id,omitempty"`
	IsCertification             *Bool     `xmlrpc:"is_certification,omitempty"`
	JobId                       *Many2One `xmlrpc:"job_id,omitempty"`
	LevelProgress               *Int      `xmlrpc:"level_progress,omitempty"`
	LevelsCount                 *Int      `xmlrpc:"levels_count,omitempty"`
	SkillId                     *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillLevelId                *Many2One `xmlrpc:"skill_level_id,omitempty"`
	SkillTypeId                 *Many2One `xmlrpc:"skill_type_id,omitempty"`
	ValidFrom                   *Time     `xmlrpc:"valid_from,omitempty"`
	ValidTo                     *Time     `xmlrpc:"valid_to,omitempty"`
	WriteDate                   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrJobSkills represents array of hr.job.skill model.
type HrJobSkills []HrJobSkill

// HrJobSkillModel is the odoo model name.
const HrJobSkillModel = "hr.job.skill"

// Many2One convert HrJobSkill to *Many2One.
func (hjs *HrJobSkill) Many2One() *Many2One {
	return NewMany2One(hjs.Id.Get(), "")
}

// CreateHrJobSkill creates a new hr.job.skill model and returns its id.
func (c *Client) CreateHrJobSkill(hjs *HrJobSkill) (int64, error) {
	ids, err := c.CreateHrJobSkills([]*HrJobSkill{hjs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrJobSkill creates a new hr.job.skill model and returns its id.
func (c *Client) CreateHrJobSkills(hjss []*HrJobSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hjss {
		vv = append(vv, v)
	}
	return c.Create(HrJobSkillModel, vv, nil)
}

// UpdateHrJobSkill updates an existing hr.job.skill record.
func (c *Client) UpdateHrJobSkill(hjs *HrJobSkill) error {
	return c.UpdateHrJobSkills([]int64{hjs.Id.Get()}, hjs)
}

// UpdateHrJobSkills updates existing hr.job.skill records.
// All records (represented by ids) will be updated by hjs values.
func (c *Client) UpdateHrJobSkills(ids []int64, hjs *HrJobSkill) error {
	return c.Update(HrJobSkillModel, ids, hjs, nil)
}

// DeleteHrJobSkill deletes an existing hr.job.skill record.
func (c *Client) DeleteHrJobSkill(id int64) error {
	return c.DeleteHrJobSkills([]int64{id})
}

// DeleteHrJobSkills deletes existing hr.job.skill records.
func (c *Client) DeleteHrJobSkills(ids []int64) error {
	return c.Delete(HrJobSkillModel, ids)
}

// GetHrJobSkill gets hr.job.skill existing record.
func (c *Client) GetHrJobSkill(id int64) (*HrJobSkill, error) {
	hjss, err := c.GetHrJobSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hjss)[0]), nil
}

// GetHrJobSkills gets hr.job.skill existing records.
func (c *Client) GetHrJobSkills(ids []int64) (*HrJobSkills, error) {
	hjss := &HrJobSkills{}
	if err := c.Read(HrJobSkillModel, ids, nil, hjss); err != nil {
		return nil, err
	}
	return hjss, nil
}

// FindHrJobSkill finds hr.job.skill record by querying it with criteria.
func (c *Client) FindHrJobSkill(criteria *Criteria) (*HrJobSkill, error) {
	hjss := &HrJobSkills{}
	if err := c.SearchRead(HrJobSkillModel, criteria, NewOptions().Limit(1), hjss); err != nil {
		return nil, err
	}
	return &((*hjss)[0]), nil
}

// FindHrJobSkills finds hr.job.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobSkills(criteria *Criteria, options *Options) (*HrJobSkills, error) {
	hjss := &HrJobSkills{}
	if err := c.SearchRead(HrJobSkillModel, criteria, options, hjss); err != nil {
		return nil, err
	}
	return hjss, nil
}

// FindHrJobSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrJobSkillModel, criteria, options)
}

// FindHrJobSkillId finds record id by querying it with criteria.
func (c *Client) FindHrJobSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
