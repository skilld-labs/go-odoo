package odoo

// CrmRecurringPlan represents crm.recurring.plan model.
type CrmRecurringPlan struct {
	Active         *Bool     `xmlrpc:"active,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	NumberOfMonths *Int      `xmlrpc:"number_of_months,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmRecurringPlans represents array of crm.recurring.plan model.
type CrmRecurringPlans []CrmRecurringPlan

// CrmRecurringPlanModel is the odoo model name.
const CrmRecurringPlanModel = "crm.recurring.plan"

// Many2One convert CrmRecurringPlan to *Many2One.
func (crp *CrmRecurringPlan) Many2One() *Many2One {
	return NewMany2One(crp.Id.Get(), "")
}

// CreateCrmRecurringPlan creates a new crm.recurring.plan model and returns its id.
func (c *Client) CreateCrmRecurringPlan(crp *CrmRecurringPlan) (int64, error) {
	ids, err := c.CreateCrmRecurringPlans([]*CrmRecurringPlan{crp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmRecurringPlan creates a new crm.recurring.plan model and returns its id.
func (c *Client) CreateCrmRecurringPlans(crps []*CrmRecurringPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range crps {
		vv = append(vv, v)
	}
	return c.Create(CrmRecurringPlanModel, vv, nil)
}

// UpdateCrmRecurringPlan updates an existing crm.recurring.plan record.
func (c *Client) UpdateCrmRecurringPlan(crp *CrmRecurringPlan) error {
	return c.UpdateCrmRecurringPlans([]int64{crp.Id.Get()}, crp)
}

// UpdateCrmRecurringPlans updates existing crm.recurring.plan records.
// All records (represented by ids) will be updated by crp values.
func (c *Client) UpdateCrmRecurringPlans(ids []int64, crp *CrmRecurringPlan) error {
	return c.Update(CrmRecurringPlanModel, ids, crp, nil)
}

// DeleteCrmRecurringPlan deletes an existing crm.recurring.plan record.
func (c *Client) DeleteCrmRecurringPlan(id int64) error {
	return c.DeleteCrmRecurringPlans([]int64{id})
}

// DeleteCrmRecurringPlans deletes existing crm.recurring.plan records.
func (c *Client) DeleteCrmRecurringPlans(ids []int64) error {
	return c.Delete(CrmRecurringPlanModel, ids)
}

// GetCrmRecurringPlan gets crm.recurring.plan existing record.
func (c *Client) GetCrmRecurringPlan(id int64) (*CrmRecurringPlan, error) {
	crps, err := c.GetCrmRecurringPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*crps)[0]), nil
}

// GetCrmRecurringPlans gets crm.recurring.plan existing records.
func (c *Client) GetCrmRecurringPlans(ids []int64) (*CrmRecurringPlans, error) {
	crps := &CrmRecurringPlans{}
	if err := c.Read(CrmRecurringPlanModel, ids, nil, crps); err != nil {
		return nil, err
	}
	return crps, nil
}

// FindCrmRecurringPlan finds crm.recurring.plan record by querying it with criteria.
func (c *Client) FindCrmRecurringPlan(criteria *Criteria) (*CrmRecurringPlan, error) {
	crps := &CrmRecurringPlans{}
	if err := c.SearchRead(CrmRecurringPlanModel, criteria, NewOptions().Limit(1), crps); err != nil {
		return nil, err
	}
	return &((*crps)[0]), nil
}

// FindCrmRecurringPlans finds crm.recurring.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmRecurringPlans(criteria *Criteria, options *Options) (*CrmRecurringPlans, error) {
	crps := &CrmRecurringPlans{}
	if err := c.SearchRead(CrmRecurringPlanModel, criteria, options, crps); err != nil {
		return nil, err
	}
	return crps, nil
}

// FindCrmRecurringPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmRecurringPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmRecurringPlanModel, criteria, options)
}

// FindCrmRecurringPlanId finds record id by querying it with criteria.
func (c *Client) FindCrmRecurringPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmRecurringPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
