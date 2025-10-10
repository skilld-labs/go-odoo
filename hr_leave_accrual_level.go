package odoo

// HrLeaveAccrualLevel represents hr.leave.accrual.level model.
type HrLeaveAccrualLevel struct {
	AccrualPlanId            *Many2One  `xmlrpc:"accrual_plan_id,omitempty"`
	AccrualValidity          *Bool      `xmlrpc:"accrual_validity,omitempty"`
	AccrualValidityCount     *Int       `xmlrpc:"accrual_validity_count,omitempty"`
	AccrualValidityType      *Selection `xmlrpc:"accrual_validity_type,omitempty"`
	AccruedGainTime          *Selection `xmlrpc:"accrued_gain_time,omitempty"`
	ActionWithUnusedAccruals *Selection `xmlrpc:"action_with_unused_accruals,omitempty"`
	AddedValue               *Float     `xmlrpc:"added_value,omitempty"`
	AddedValueType           *Selection `xmlrpc:"added_value_type,omitempty"`
	CanModifyValueType       *Bool      `xmlrpc:"can_modify_value_type,omitempty"`
	CapAccruedTime           *Bool      `xmlrpc:"cap_accrued_time,omitempty"`
	CapAccruedTimeYearly     *Bool      `xmlrpc:"cap_accrued_time_yearly,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	FirstDay                 *Int       `xmlrpc:"first_day,omitempty"`
	FirstDayDisplay          *Selection `xmlrpc:"first_day_display,omitempty"`
	FirstMonth               *Selection `xmlrpc:"first_month,omitempty"`
	FirstMonthDay            *Int       `xmlrpc:"first_month_day,omitempty"`
	FirstMonthDayDisplay     *Selection `xmlrpc:"first_month_day_display,omitempty"`
	Frequency                *Selection `xmlrpc:"frequency,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MaximumLeave             *Float     `xmlrpc:"maximum_leave,omitempty"`
	MaximumLeaveYearly       *Float     `xmlrpc:"maximum_leave_yearly,omitempty"`
	PostponeMaxDays          *Int       `xmlrpc:"postpone_max_days,omitempty"`
	SecondDay                *Int       `xmlrpc:"second_day,omitempty"`
	SecondDayDisplay         *Selection `xmlrpc:"second_day_display,omitempty"`
	SecondMonth              *Selection `xmlrpc:"second_month,omitempty"`
	SecondMonthDay           *Int       `xmlrpc:"second_month_day,omitempty"`
	SecondMonthDayDisplay    *Selection `xmlrpc:"second_month_day_display,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	StartCount               *Int       `xmlrpc:"start_count,omitempty"`
	StartType                *Selection `xmlrpc:"start_type,omitempty"`
	WeekDay                  *Selection `xmlrpc:"week_day,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
	YearlyDay                *Int       `xmlrpc:"yearly_day,omitempty"`
	YearlyDayDisplay         *Selection `xmlrpc:"yearly_day_display,omitempty"`
	YearlyMonth              *Selection `xmlrpc:"yearly_month,omitempty"`
}

// HrLeaveAccrualLevels represents array of hr.leave.accrual.level model.
type HrLeaveAccrualLevels []HrLeaveAccrualLevel

// HrLeaveAccrualLevelModel is the odoo model name.
const HrLeaveAccrualLevelModel = "hr.leave.accrual.level"

// Many2One convert HrLeaveAccrualLevel to *Many2One.
func (hlal *HrLeaveAccrualLevel) Many2One() *Many2One {
	return NewMany2One(hlal.Id.Get(), "")
}

// CreateHrLeaveAccrualLevel creates a new hr.leave.accrual.level model and returns its id.
func (c *Client) CreateHrLeaveAccrualLevel(hlal *HrLeaveAccrualLevel) (int64, error) {
	ids, err := c.CreateHrLeaveAccrualLevels([]*HrLeaveAccrualLevel{hlal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAccrualLevel creates a new hr.leave.accrual.level model and returns its id.
func (c *Client) CreateHrLeaveAccrualLevels(hlals []*HrLeaveAccrualLevel) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlals {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAccrualLevelModel, vv, nil)
}

// UpdateHrLeaveAccrualLevel updates an existing hr.leave.accrual.level record.
func (c *Client) UpdateHrLeaveAccrualLevel(hlal *HrLeaveAccrualLevel) error {
	return c.UpdateHrLeaveAccrualLevels([]int64{hlal.Id.Get()}, hlal)
}

// UpdateHrLeaveAccrualLevels updates existing hr.leave.accrual.level records.
// All records (represented by ids) will be updated by hlal values.
func (c *Client) UpdateHrLeaveAccrualLevels(ids []int64, hlal *HrLeaveAccrualLevel) error {
	return c.Update(HrLeaveAccrualLevelModel, ids, hlal, nil)
}

// DeleteHrLeaveAccrualLevel deletes an existing hr.leave.accrual.level record.
func (c *Client) DeleteHrLeaveAccrualLevel(id int64) error {
	return c.DeleteHrLeaveAccrualLevels([]int64{id})
}

// DeleteHrLeaveAccrualLevels deletes existing hr.leave.accrual.level records.
func (c *Client) DeleteHrLeaveAccrualLevels(ids []int64) error {
	return c.Delete(HrLeaveAccrualLevelModel, ids)
}

// GetHrLeaveAccrualLevel gets hr.leave.accrual.level existing record.
func (c *Client) GetHrLeaveAccrualLevel(id int64) (*HrLeaveAccrualLevel, error) {
	hlals, err := c.GetHrLeaveAccrualLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlals)[0]), nil
}

// GetHrLeaveAccrualLevels gets hr.leave.accrual.level existing records.
func (c *Client) GetHrLeaveAccrualLevels(ids []int64) (*HrLeaveAccrualLevels, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.Read(HrLeaveAccrualLevelModel, ids, nil, hlals); err != nil {
		return nil, err
	}
	return hlals, nil
}

// FindHrLeaveAccrualLevel finds hr.leave.accrual.level record by querying it with criteria.
func (c *Client) FindHrLeaveAccrualLevel(criteria *Criteria) (*HrLeaveAccrualLevel, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.SearchRead(HrLeaveAccrualLevelModel, criteria, NewOptions().Limit(1), hlals); err != nil {
		return nil, err
	}
	return &((*hlals)[0]), nil
}

// FindHrLeaveAccrualLevels finds hr.leave.accrual.level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualLevels(criteria *Criteria, options *Options) (*HrLeaveAccrualLevels, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.SearchRead(HrLeaveAccrualLevelModel, criteria, options, hlals); err != nil {
		return nil, err
	}
	return hlals, nil
}

// FindHrLeaveAccrualLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveAccrualLevelModel, criteria, options)
}

// FindHrLeaveAccrualLevelId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAccrualLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAccrualLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
