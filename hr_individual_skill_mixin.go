package odoo

// HrIndividualSkillMixin represents hr.individual.skill.mixin model.
type HrIndividualSkillMixin struct {
	CertificationSkillTypeCount *Int      `xmlrpc:"certification_skill_type_count,omitempty"`
	Color                       *Int      `xmlrpc:"color,omitempty"`
	DisplayName                 *String   `xmlrpc:"display_name,omitempty"`
	DisplayWarningMessage       *Bool     `xmlrpc:"display_warning_message,omitempty"`
	Id                          *Int      `xmlrpc:"id,omitempty"`
	IsCertification             *Bool     `xmlrpc:"is_certification,omitempty"`
	LevelProgress               *Int      `xmlrpc:"level_progress,omitempty"`
	LevelsCount                 *Int      `xmlrpc:"levels_count,omitempty"`
	SkillId                     *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillLevelId                *Many2One `xmlrpc:"skill_level_id,omitempty"`
	SkillTypeId                 *Many2One `xmlrpc:"skill_type_id,omitempty"`
	ValidFrom                   *Time     `xmlrpc:"valid_from,omitempty"`
	ValidTo                     *Time     `xmlrpc:"valid_to,omitempty"`
}

// HrIndividualSkillMixins represents array of hr.individual.skill.mixin model.
type HrIndividualSkillMixins []HrIndividualSkillMixin

// HrIndividualSkillMixinModel is the odoo model name.
const HrIndividualSkillMixinModel = "hr.individual.skill.mixin"

// Many2One convert HrIndividualSkillMixin to *Many2One.
func (hism *HrIndividualSkillMixin) Many2One() *Many2One {
	return NewMany2One(hism.Id.Get(), "")
}

// CreateHrIndividualSkillMixin creates a new hr.individual.skill.mixin model and returns its id.
func (c *Client) CreateHrIndividualSkillMixin(hism *HrIndividualSkillMixin) (int64, error) {
	ids, err := c.CreateHrIndividualSkillMixins([]*HrIndividualSkillMixin{hism})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrIndividualSkillMixin creates a new hr.individual.skill.mixin model and returns its id.
func (c *Client) CreateHrIndividualSkillMixins(hisms []*HrIndividualSkillMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range hisms {
		vv = append(vv, v)
	}
	return c.Create(HrIndividualSkillMixinModel, vv, nil)
}

// UpdateHrIndividualSkillMixin updates an existing hr.individual.skill.mixin record.
func (c *Client) UpdateHrIndividualSkillMixin(hism *HrIndividualSkillMixin) error {
	return c.UpdateHrIndividualSkillMixins([]int64{hism.Id.Get()}, hism)
}

// UpdateHrIndividualSkillMixins updates existing hr.individual.skill.mixin records.
// All records (represented by ids) will be updated by hism values.
func (c *Client) UpdateHrIndividualSkillMixins(ids []int64, hism *HrIndividualSkillMixin) error {
	return c.Update(HrIndividualSkillMixinModel, ids, hism, nil)
}

// DeleteHrIndividualSkillMixin deletes an existing hr.individual.skill.mixin record.
func (c *Client) DeleteHrIndividualSkillMixin(id int64) error {
	return c.DeleteHrIndividualSkillMixins([]int64{id})
}

// DeleteHrIndividualSkillMixins deletes existing hr.individual.skill.mixin records.
func (c *Client) DeleteHrIndividualSkillMixins(ids []int64) error {
	return c.Delete(HrIndividualSkillMixinModel, ids)
}

// GetHrIndividualSkillMixin gets hr.individual.skill.mixin existing record.
func (c *Client) GetHrIndividualSkillMixin(id int64) (*HrIndividualSkillMixin, error) {
	hisms, err := c.GetHrIndividualSkillMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hisms)[0]), nil
}

// GetHrIndividualSkillMixins gets hr.individual.skill.mixin existing records.
func (c *Client) GetHrIndividualSkillMixins(ids []int64) (*HrIndividualSkillMixins, error) {
	hisms := &HrIndividualSkillMixins{}
	if err := c.Read(HrIndividualSkillMixinModel, ids, nil, hisms); err != nil {
		return nil, err
	}
	return hisms, nil
}

// FindHrIndividualSkillMixin finds hr.individual.skill.mixin record by querying it with criteria.
func (c *Client) FindHrIndividualSkillMixin(criteria *Criteria) (*HrIndividualSkillMixin, error) {
	hisms := &HrIndividualSkillMixins{}
	if err := c.SearchRead(HrIndividualSkillMixinModel, criteria, NewOptions().Limit(1), hisms); err != nil {
		return nil, err
	}
	return &((*hisms)[0]), nil
}

// FindHrIndividualSkillMixins finds hr.individual.skill.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrIndividualSkillMixins(criteria *Criteria, options *Options) (*HrIndividualSkillMixins, error) {
	hisms := &HrIndividualSkillMixins{}
	if err := c.SearchRead(HrIndividualSkillMixinModel, criteria, options, hisms); err != nil {
		return nil, err
	}
	return hisms, nil
}

// FindHrIndividualSkillMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrIndividualSkillMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrIndividualSkillMixinModel, criteria, options)
}

// FindHrIndividualSkillMixinId finds record id by querying it with criteria.
func (c *Client) FindHrIndividualSkillMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrIndividualSkillMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
