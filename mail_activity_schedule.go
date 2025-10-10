package odoo

// MailActivitySchedule represents mail.activity.schedule model.
type MailActivitySchedule struct {
	ActivityCategory         *Selection `xmlrpc:"activity_category,omitempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	ChainingType             *Selection `xmlrpc:"chaining_type,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateDeadline             *Time      `xmlrpc:"date_deadline,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Error                    *String    `xmlrpc:"error,omitempty"`
	HasError                 *Bool      `xmlrpc:"has_error,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsBatchMode              *Bool      `xmlrpc:"is_batch_mode,omitempty"`
	Note                     *String    `xmlrpc:"note,omitempty"`
	PlanAvailableIds         *Relation  `xmlrpc:"plan_available_ids,omitempty"`
	PlanDate                 *Time      `xmlrpc:"plan_date,omitempty"`
	PlanDepartmentFilterable *Bool      `xmlrpc:"plan_department_filterable,omitempty"`
	PlanHasUserOnDemand      *Bool      `xmlrpc:"plan_has_user_on_demand,omitempty"`
	PlanId                   *Many2One  `xmlrpc:"plan_id,omitempty"`
	PlanOnDemandUserId       *Many2One  `xmlrpc:"plan_on_demand_user_id,omitempty"`
	PlanSummary              *String    `xmlrpc:"plan_summary,omitempty"`
	ResIds                   *String    `xmlrpc:"res_ids,omitempty"`
	ResModel                 *String    `xmlrpc:"res_model,omitempty"`
	ResModelId               *Many2One  `xmlrpc:"res_model_id,omitempty"`
	Summary                  *String    `xmlrpc:"summary,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailActivitySchedules represents array of mail.activity.schedule model.
type MailActivitySchedules []MailActivitySchedule

// MailActivityScheduleModel is the odoo model name.
const MailActivityScheduleModel = "mail.activity.schedule"

// Many2One convert MailActivitySchedule to *Many2One.
func (mas *MailActivitySchedule) Many2One() *Many2One {
	return NewMany2One(mas.Id.Get(), "")
}

// CreateMailActivitySchedule creates a new mail.activity.schedule model and returns its id.
func (c *Client) CreateMailActivitySchedule(mas *MailActivitySchedule) (int64, error) {
	ids, err := c.CreateMailActivitySchedules([]*MailActivitySchedule{mas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivitySchedule creates a new mail.activity.schedule model and returns its id.
func (c *Client) CreateMailActivitySchedules(mass []*MailActivitySchedule) ([]int64, error) {
	var vv []interface{}
	for _, v := range mass {
		vv = append(vv, v)
	}
	return c.Create(MailActivityScheduleModel, vv, nil)
}

// UpdateMailActivitySchedule updates an existing mail.activity.schedule record.
func (c *Client) UpdateMailActivitySchedule(mas *MailActivitySchedule) error {
	return c.UpdateMailActivitySchedules([]int64{mas.Id.Get()}, mas)
}

// UpdateMailActivitySchedules updates existing mail.activity.schedule records.
// All records (represented by ids) will be updated by mas values.
func (c *Client) UpdateMailActivitySchedules(ids []int64, mas *MailActivitySchedule) error {
	return c.Update(MailActivityScheduleModel, ids, mas, nil)
}

// DeleteMailActivitySchedule deletes an existing mail.activity.schedule record.
func (c *Client) DeleteMailActivitySchedule(id int64) error {
	return c.DeleteMailActivitySchedules([]int64{id})
}

// DeleteMailActivitySchedules deletes existing mail.activity.schedule records.
func (c *Client) DeleteMailActivitySchedules(ids []int64) error {
	return c.Delete(MailActivityScheduleModel, ids)
}

// GetMailActivitySchedule gets mail.activity.schedule existing record.
func (c *Client) GetMailActivitySchedule(id int64) (*MailActivitySchedule, error) {
	mass, err := c.GetMailActivitySchedules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mass)[0]), nil
}

// GetMailActivitySchedules gets mail.activity.schedule existing records.
func (c *Client) GetMailActivitySchedules(ids []int64) (*MailActivitySchedules, error) {
	mass := &MailActivitySchedules{}
	if err := c.Read(MailActivityScheduleModel, ids, nil, mass); err != nil {
		return nil, err
	}
	return mass, nil
}

// FindMailActivitySchedule finds mail.activity.schedule record by querying it with criteria.
func (c *Client) FindMailActivitySchedule(criteria *Criteria) (*MailActivitySchedule, error) {
	mass := &MailActivitySchedules{}
	if err := c.SearchRead(MailActivityScheduleModel, criteria, NewOptions().Limit(1), mass); err != nil {
		return nil, err
	}
	return &((*mass)[0]), nil
}

// FindMailActivitySchedules finds mail.activity.schedule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivitySchedules(criteria *Criteria, options *Options) (*MailActivitySchedules, error) {
	mass := &MailActivitySchedules{}
	if err := c.SearchRead(MailActivityScheduleModel, criteria, options, mass); err != nil {
		return nil, err
	}
	return mass, nil
}

// FindMailActivityScheduleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityScheduleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityScheduleModel, criteria, options)
}

// FindMailActivityScheduleId finds record id by querying it with criteria.
func (c *Client) FindMailActivityScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
