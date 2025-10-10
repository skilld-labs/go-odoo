package odoo

// HrSkillLevel represents hr.skill.level model.
type HrSkillLevel struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultLevel  *Bool     `xmlrpc:"default_level,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrSkillLevels represents array of hr.skill.level model.
type HrSkillLevels []HrSkillLevel

// HrSkillLevelModel is the odoo model name.
const HrSkillLevelModel = "hr.skill.level"

// Many2One convert HrSkillLevel to *Many2One.
func (hsl *HrSkillLevel) Many2One() *Many2One {
	return NewMany2One(hsl.Id.Get(), "")
}

// CreateHrSkillLevel creates a new hr.skill.level model and returns its id.
func (c *Client) CreateHrSkillLevel(hsl *HrSkillLevel) (int64, error) {
	ids, err := c.CreateHrSkillLevels([]*HrSkillLevel{hsl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSkillLevel creates a new hr.skill.level model and returns its id.
func (c *Client) CreateHrSkillLevels(hsls []*HrSkillLevel) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsls {
		vv = append(vv, v)
	}
	return c.Create(HrSkillLevelModel, vv, nil)
}

// UpdateHrSkillLevel updates an existing hr.skill.level record.
func (c *Client) UpdateHrSkillLevel(hsl *HrSkillLevel) error {
	return c.UpdateHrSkillLevels([]int64{hsl.Id.Get()}, hsl)
}

// UpdateHrSkillLevels updates existing hr.skill.level records.
// All records (represented by ids) will be updated by hsl values.
func (c *Client) UpdateHrSkillLevels(ids []int64, hsl *HrSkillLevel) error {
	return c.Update(HrSkillLevelModel, ids, hsl, nil)
}

// DeleteHrSkillLevel deletes an existing hr.skill.level record.
func (c *Client) DeleteHrSkillLevel(id int64) error {
	return c.DeleteHrSkillLevels([]int64{id})
}

// DeleteHrSkillLevels deletes existing hr.skill.level records.
func (c *Client) DeleteHrSkillLevels(ids []int64) error {
	return c.Delete(HrSkillLevelModel, ids)
}

// GetHrSkillLevel gets hr.skill.level existing record.
func (c *Client) GetHrSkillLevel(id int64) (*HrSkillLevel, error) {
	hsls, err := c.GetHrSkillLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hsls)[0]), nil
}

// GetHrSkillLevels gets hr.skill.level existing records.
func (c *Client) GetHrSkillLevels(ids []int64) (*HrSkillLevels, error) {
	hsls := &HrSkillLevels{}
	if err := c.Read(HrSkillLevelModel, ids, nil, hsls); err != nil {
		return nil, err
	}
	return hsls, nil
}

// FindHrSkillLevel finds hr.skill.level record by querying it with criteria.
func (c *Client) FindHrSkillLevel(criteria *Criteria) (*HrSkillLevel, error) {
	hsls := &HrSkillLevels{}
	if err := c.SearchRead(HrSkillLevelModel, criteria, NewOptions().Limit(1), hsls); err != nil {
		return nil, err
	}
	return &((*hsls)[0]), nil
}

// FindHrSkillLevels finds hr.skill.level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillLevels(criteria *Criteria, options *Options) (*HrSkillLevels, error) {
	hsls := &HrSkillLevels{}
	if err := c.SearchRead(HrSkillLevelModel, criteria, options, hsls); err != nil {
		return nil, err
	}
	return hsls, nil
}

// FindHrSkillLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrSkillLevelModel, criteria, options)
}

// FindHrSkillLevelId finds record id by querying it with criteria.
func (c *Client) FindHrSkillLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSkillLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
