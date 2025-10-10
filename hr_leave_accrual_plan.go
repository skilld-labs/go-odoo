package odoo

// HrLeaveAccrualPlan represents hr.leave.accrual.plan model.
type HrLeaveAccrualPlan struct {
	AccruedGainTime     *Selection `xmlrpc:"accrued_gain_time,omitempty"`
	Active              *Bool      `xmlrpc:"active,omitempty"`
	AddedValueType      *Selection `xmlrpc:"added_value_type,omitempty"`
	AllocationIds       *Relation  `xmlrpc:"allocation_ids,omitempty"`
	CarryoverDate       *Selection `xmlrpc:"carryover_date,omitempty"`
	CarryoverDay        *Int       `xmlrpc:"carryover_day,omitempty"`
	CarryoverDayDisplay *Selection `xmlrpc:"carryover_day_display,omitempty"`
	CarryoverMonth      *Selection `xmlrpc:"carryover_month,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EmployeesCount      *Int       `xmlrpc:"employees_count,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IsBasedOnWorkedTime *Bool      `xmlrpc:"is_based_on_worked_time,omitempty"`
	LevelCount          *Int       `xmlrpc:"level_count,omitempty"`
	LevelIds            *Relation  `xmlrpc:"level_ids,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	ShowTransitionMode  *Bool      `xmlrpc:"show_transition_mode,omitempty"`
	TimeOffTypeId       *Many2One  `xmlrpc:"time_off_type_id,omitempty"`
	TransitionMode      *Selection `xmlrpc:"transition_mode,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveAccrualPlans represents array of hr.leave.accrual.plan model.
type HrLeaveAccrualPlans []HrLeaveAccrualPlan

// HrLeaveAccrualPlanModel is the odoo model name.
const HrLeaveAccrualPlanModel = "hr.leave.accrual.plan"

// Many2One convert HrLeaveAccrualPlan to *Many2One.
func (hlap *HrLeaveAccrualPlan) Many2One() *Many2One {
	return NewMany2One(hlap.Id.Get(), "")
}

// CreateHrLeaveAccrualPlan creates a new hr.leave.accrual.plan model and returns its id.
func (c *Client) CreateHrLeaveAccrualPlan(hlap *HrLeaveAccrualPlan) (int64, error) {
	ids, err := c.CreateHrLeaveAccrualPlans([]*HrLeaveAccrualPlan{hlap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAccrualPlan creates a new hr.leave.accrual.plan model and returns its id.
func (c *Client) CreateHrLeaveAccrualPlans(hlaps []*HrLeaveAccrualPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlaps {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAccrualPlanModel, vv, nil)
}

// UpdateHrLeaveAccrualPlan updates an existing hr.leave.accrual.plan record.
func (c *Client) UpdateHrLeaveAccrualPlan(hlap *HrLeaveAccrualPlan) error {
	return c.UpdateHrLeaveAccrualPlans([]int64{hlap.Id.Get()}, hlap)
}

// UpdateHrLeaveAccrualPlans updates existing hr.leave.accrual.plan records.
// All records (represented by ids) will be updated by hlap values.
func (c *Client) UpdateHrLeaveAccrualPlans(ids []int64, hlap *HrLeaveAccrualPlan) error {
	return c.Update(HrLeaveAccrualPlanModel, ids, hlap, nil)
}

// DeleteHrLeaveAccrualPlan deletes an existing hr.leave.accrual.plan record.
func (c *Client) DeleteHrLeaveAccrualPlan(id int64) error {
	return c.DeleteHrLeaveAccrualPlans([]int64{id})
}

// DeleteHrLeaveAccrualPlans deletes existing hr.leave.accrual.plan records.
func (c *Client) DeleteHrLeaveAccrualPlans(ids []int64) error {
	return c.Delete(HrLeaveAccrualPlanModel, ids)
}

// GetHrLeaveAccrualPlan gets hr.leave.accrual.plan existing record.
func (c *Client) GetHrLeaveAccrualPlan(id int64) (*HrLeaveAccrualPlan, error) {
	hlaps, err := c.GetHrLeaveAccrualPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlaps)[0]), nil
}

// GetHrLeaveAccrualPlans gets hr.leave.accrual.plan existing records.
func (c *Client) GetHrLeaveAccrualPlans(ids []int64) (*HrLeaveAccrualPlans, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.Read(HrLeaveAccrualPlanModel, ids, nil, hlaps); err != nil {
		return nil, err
	}
	return hlaps, nil
}

// FindHrLeaveAccrualPlan finds hr.leave.accrual.plan record by querying it with criteria.
func (c *Client) FindHrLeaveAccrualPlan(criteria *Criteria) (*HrLeaveAccrualPlan, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.SearchRead(HrLeaveAccrualPlanModel, criteria, NewOptions().Limit(1), hlaps); err != nil {
		return nil, err
	}
	return &((*hlaps)[0]), nil
}

// FindHrLeaveAccrualPlans finds hr.leave.accrual.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualPlans(criteria *Criteria, options *Options) (*HrLeaveAccrualPlans, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.SearchRead(HrLeaveAccrualPlanModel, criteria, options, hlaps); err != nil {
		return nil, err
	}
	return hlaps, nil
}

// FindHrLeaveAccrualPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveAccrualPlanModel, criteria, options)
}

// FindHrLeaveAccrualPlanId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAccrualPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAccrualPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
