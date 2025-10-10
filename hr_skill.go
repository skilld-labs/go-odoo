package odoo

// HrSkill represents hr.skill model.
type HrSkill struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	SkillTypeId *Many2One `xmlrpc:"skill_type_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrSkills represents array of hr.skill model.
type HrSkills []HrSkill

// HrSkillModel is the odoo model name.
const HrSkillModel = "hr.skill"

// Many2One convert HrSkill to *Many2One.
func (hs *HrSkill) Many2One() *Many2One {
	return NewMany2One(hs.Id.Get(), "")
}

// CreateHrSkill creates a new hr.skill model and returns its id.
func (c *Client) CreateHrSkill(hs *HrSkill) (int64, error) {
	ids, err := c.CreateHrSkills([]*HrSkill{hs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSkill creates a new hr.skill model and returns its id.
func (c *Client) CreateHrSkills(hss []*HrSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hss {
		vv = append(vv, v)
	}
	return c.Create(HrSkillModel, vv, nil)
}

// UpdateHrSkill updates an existing hr.skill record.
func (c *Client) UpdateHrSkill(hs *HrSkill) error {
	return c.UpdateHrSkills([]int64{hs.Id.Get()}, hs)
}

// UpdateHrSkills updates existing hr.skill records.
// All records (represented by ids) will be updated by hs values.
func (c *Client) UpdateHrSkills(ids []int64, hs *HrSkill) error {
	return c.Update(HrSkillModel, ids, hs, nil)
}

// DeleteHrSkill deletes an existing hr.skill record.
func (c *Client) DeleteHrSkill(id int64) error {
	return c.DeleteHrSkills([]int64{id})
}

// DeleteHrSkills deletes existing hr.skill records.
func (c *Client) DeleteHrSkills(ids []int64) error {
	return c.Delete(HrSkillModel, ids)
}

// GetHrSkill gets hr.skill existing record.
func (c *Client) GetHrSkill(id int64) (*HrSkill, error) {
	hss, err := c.GetHrSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hss)[0]), nil
}

// GetHrSkills gets hr.skill existing records.
func (c *Client) GetHrSkills(ids []int64) (*HrSkills, error) {
	hss := &HrSkills{}
	if err := c.Read(HrSkillModel, ids, nil, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHrSkill finds hr.skill record by querying it with criteria.
func (c *Client) FindHrSkill(criteria *Criteria) (*HrSkill, error) {
	hss := &HrSkills{}
	if err := c.SearchRead(HrSkillModel, criteria, NewOptions().Limit(1), hss); err != nil {
		return nil, err
	}
	return &((*hss)[0]), nil
}

// FindHrSkills finds hr.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkills(criteria *Criteria, options *Options) (*HrSkills, error) {
	hss := &HrSkills{}
	if err := c.SearchRead(HrSkillModel, criteria, options, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHrSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrSkillModel, criteria, options)
}

// FindHrSkillId finds record id by querying it with criteria.
func (c *Client) FindHrSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
