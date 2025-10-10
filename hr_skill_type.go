package odoo

// HrSkillType represents hr.skill.type model.
type HrSkillType struct {
	Active        *Bool     `xmlrpc:"active,omitempty"`
	Color         *Int      `xmlrpc:"color,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	SkillIds      *Relation `xmlrpc:"skill_ids,omitempty"`
	SkillLevelIds *Relation `xmlrpc:"skill_level_ids,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrSkillTypes represents array of hr.skill.type model.
type HrSkillTypes []HrSkillType

// HrSkillTypeModel is the odoo model name.
const HrSkillTypeModel = "hr.skill.type"

// Many2One convert HrSkillType to *Many2One.
func (hst *HrSkillType) Many2One() *Many2One {
	return NewMany2One(hst.Id.Get(), "")
}

// CreateHrSkillType creates a new hr.skill.type model and returns its id.
func (c *Client) CreateHrSkillType(hst *HrSkillType) (int64, error) {
	ids, err := c.CreateHrSkillTypes([]*HrSkillType{hst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSkillType creates a new hr.skill.type model and returns its id.
func (c *Client) CreateHrSkillTypes(hsts []*HrSkillType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsts {
		vv = append(vv, v)
	}
	return c.Create(HrSkillTypeModel, vv, nil)
}

// UpdateHrSkillType updates an existing hr.skill.type record.
func (c *Client) UpdateHrSkillType(hst *HrSkillType) error {
	return c.UpdateHrSkillTypes([]int64{hst.Id.Get()}, hst)
}

// UpdateHrSkillTypes updates existing hr.skill.type records.
// All records (represented by ids) will be updated by hst values.
func (c *Client) UpdateHrSkillTypes(ids []int64, hst *HrSkillType) error {
	return c.Update(HrSkillTypeModel, ids, hst, nil)
}

// DeleteHrSkillType deletes an existing hr.skill.type record.
func (c *Client) DeleteHrSkillType(id int64) error {
	return c.DeleteHrSkillTypes([]int64{id})
}

// DeleteHrSkillTypes deletes existing hr.skill.type records.
func (c *Client) DeleteHrSkillTypes(ids []int64) error {
	return c.Delete(HrSkillTypeModel, ids)
}

// GetHrSkillType gets hr.skill.type existing record.
func (c *Client) GetHrSkillType(id int64) (*HrSkillType, error) {
	hsts, err := c.GetHrSkillTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hsts)[0]), nil
}

// GetHrSkillTypes gets hr.skill.type existing records.
func (c *Client) GetHrSkillTypes(ids []int64) (*HrSkillTypes, error) {
	hsts := &HrSkillTypes{}
	if err := c.Read(HrSkillTypeModel, ids, nil, hsts); err != nil {
		return nil, err
	}
	return hsts, nil
}

// FindHrSkillType finds hr.skill.type record by querying it with criteria.
func (c *Client) FindHrSkillType(criteria *Criteria) (*HrSkillType, error) {
	hsts := &HrSkillTypes{}
	if err := c.SearchRead(HrSkillTypeModel, criteria, NewOptions().Limit(1), hsts); err != nil {
		return nil, err
	}
	return &((*hsts)[0]), nil
}

// FindHrSkillTypes finds hr.skill.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillTypes(criteria *Criteria, options *Options) (*HrSkillTypes, error) {
	hsts := &HrSkillTypes{}
	if err := c.SearchRead(HrSkillTypeModel, criteria, options, hsts); err != nil {
		return nil, err
	}
	return hsts, nil
}

// FindHrSkillTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrSkillTypeModel, criteria, options)
}

// FindHrSkillTypeId finds record id by querying it with criteria.
func (c *Client) FindHrSkillTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSkillTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
