package odoo

// MailActivityPlan represents mail.activity.plan model.
type MailActivityPlan struct {
	Active               *Bool      `xmlrpc:"active,omitempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DepartmentAssignable *Bool      `xmlrpc:"department_assignable,omitempty"`
	DepartmentId         *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	HasUserOnDemand      *Bool      `xmlrpc:"has_user_on_demand,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	ResModel             *Selection `xmlrpc:"res_model,omitempty"`
	ResModelId           *Many2One  `xmlrpc:"res_model_id,omitempty"`
	StepsCount           *Int       `xmlrpc:"steps_count,omitempty"`
	TemplateIds          *Relation  `xmlrpc:"template_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailActivityPlans represents array of mail.activity.plan model.
type MailActivityPlans []MailActivityPlan

// MailActivityPlanModel is the odoo model name.
const MailActivityPlanModel = "mail.activity.plan"

// Many2One convert MailActivityPlan to *Many2One.
func (MAP *MailActivityPlan) Many2One() *Many2One {
	return NewMany2One(MAP.Id.Get(), "")
}

// CreateMailActivityPlan creates a new mail.activity.plan model and returns its id.
func (c *Client) CreateMailActivityPlan(MAP *MailActivityPlan) (int64, error) {
	ids, err := c.CreateMailActivityPlans([]*MailActivityPlan{MAP})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityPlan creates a new mail.activity.plan model and returns its id.
func (c *Client) CreateMailActivityPlans(MAPs []*MailActivityPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range MAPs {
		vv = append(vv, v)
	}
	return c.Create(MailActivityPlanModel, vv, nil)
}

// UpdateMailActivityPlan updates an existing mail.activity.plan record.
func (c *Client) UpdateMailActivityPlan(MAP *MailActivityPlan) error {
	return c.UpdateMailActivityPlans([]int64{MAP.Id.Get()}, MAP)
}

// UpdateMailActivityPlans updates existing mail.activity.plan records.
// All records (represented by ids) will be updated by MAP values.
func (c *Client) UpdateMailActivityPlans(ids []int64, MAP *MailActivityPlan) error {
	return c.Update(MailActivityPlanModel, ids, MAP, nil)
}

// DeleteMailActivityPlan deletes an existing mail.activity.plan record.
func (c *Client) DeleteMailActivityPlan(id int64) error {
	return c.DeleteMailActivityPlans([]int64{id})
}

// DeleteMailActivityPlans deletes existing mail.activity.plan records.
func (c *Client) DeleteMailActivityPlans(ids []int64) error {
	return c.Delete(MailActivityPlanModel, ids)
}

// GetMailActivityPlan gets mail.activity.plan existing record.
func (c *Client) GetMailActivityPlan(id int64) (*MailActivityPlan, error) {
	MAPs, err := c.GetMailActivityPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*MAPs)[0]), nil
}

// GetMailActivityPlans gets mail.activity.plan existing records.
func (c *Client) GetMailActivityPlans(ids []int64) (*MailActivityPlans, error) {
	MAPs := &MailActivityPlans{}
	if err := c.Read(MailActivityPlanModel, ids, nil, MAPs); err != nil {
		return nil, err
	}
	return MAPs, nil
}

// FindMailActivityPlan finds mail.activity.plan record by querying it with criteria.
func (c *Client) FindMailActivityPlan(criteria *Criteria) (*MailActivityPlan, error) {
	MAPs := &MailActivityPlans{}
	if err := c.SearchRead(MailActivityPlanModel, criteria, NewOptions().Limit(1), MAPs); err != nil {
		return nil, err
	}
	return &((*MAPs)[0]), nil
}

// FindMailActivityPlans finds mail.activity.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlans(criteria *Criteria, options *Options) (*MailActivityPlans, error) {
	MAPs := &MailActivityPlans{}
	if err := c.SearchRead(MailActivityPlanModel, criteria, options, MAPs); err != nil {
		return nil, err
	}
	return MAPs, nil
}

// FindMailActivityPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityPlanModel, criteria, options)
}

// FindMailActivityPlanId finds record id by querying it with criteria.
func (c *Client) FindMailActivityPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
