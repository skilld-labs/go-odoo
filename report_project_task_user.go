package odoo

// ReportProjectTaskUser represents report.project.task.user model.
type ReportProjectTaskUser struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	DateDeadline        *Time      `xmlrpc:"date_deadline,omitempty"`
	DateEnd             *Time      `xmlrpc:"date_end,omitempty"`
	DateLastStageUpdate *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	DateStart           *Time      `xmlrpc:"date_start,omitempty"`
	DelayEndingsDays    *Float     `xmlrpc:"delay_endings_days,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	HoursDelay          *Float     `xmlrpc:"hours_delay,omitempty"`
	HoursEffective      *Float     `xmlrpc:"hours_effective,omitempty"`
	HoursPlanned        *Float     `xmlrpc:"hours_planned,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	Nbr                 *Int       `xmlrpc:"nbr,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	Priority            *Selection `xmlrpc:"priority,omitempty"`
	Progress            *Float     `xmlrpc:"progress,omitempty"`
	ProjectId           *Many2One  `xmlrpc:"project_id,omitempty"`
	RemainingHours      *Float     `xmlrpc:"remaining_hours,omitempty"`
	StageId             *Many2One  `xmlrpc:"stage_id,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	TotalHours          *Float     `xmlrpc:"total_hours,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
	WorkingDaysClose    *Float     `xmlrpc:"working_days_close,omitempty"`
	WorkingDaysOpen     *Float     `xmlrpc:"working_days_open,omitempty"`
}

// ReportProjectTaskUsers represents array of report.project.task.user model.
type ReportProjectTaskUsers []ReportProjectTaskUser

// ReportProjectTaskUserModel is the odoo model name.
const ReportProjectTaskUserModel = "report.project.task.user"

// Many2One convert ReportProjectTaskUser to *Many2One.
func (rptu *ReportProjectTaskUser) Many2One() *Many2One {
	return NewMany2One(rptu.Id.Get(), "")
}

// CreateReportProjectTaskUser creates a new report.project.task.user model and returns its id.
func (c *Client) CreateReportProjectTaskUser(rptu *ReportProjectTaskUser) (int64, error) {
	ids, err := c.CreateReportProjectTaskUsers([]*ReportProjectTaskUser{rptu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProjectTaskUser creates a new report.project.task.user model and returns its id.
func (c *Client) CreateReportProjectTaskUsers(rptus []*ReportProjectTaskUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range rptus {
		vv = append(vv, v)
	}
	return c.Create(ReportProjectTaskUserModel, vv, nil)
}

// UpdateReportProjectTaskUser updates an existing report.project.task.user record.
func (c *Client) UpdateReportProjectTaskUser(rptu *ReportProjectTaskUser) error {
	return c.UpdateReportProjectTaskUsers([]int64{rptu.Id.Get()}, rptu)
}

// UpdateReportProjectTaskUsers updates existing report.project.task.user records.
// All records (represented by ids) will be updated by rptu values.
func (c *Client) UpdateReportProjectTaskUsers(ids []int64, rptu *ReportProjectTaskUser) error {
	return c.Update(ReportProjectTaskUserModel, ids, rptu, nil)
}

// DeleteReportProjectTaskUser deletes an existing report.project.task.user record.
func (c *Client) DeleteReportProjectTaskUser(id int64) error {
	return c.DeleteReportProjectTaskUsers([]int64{id})
}

// DeleteReportProjectTaskUsers deletes existing report.project.task.user records.
func (c *Client) DeleteReportProjectTaskUsers(ids []int64) error {
	return c.Delete(ReportProjectTaskUserModel, ids)
}

// GetReportProjectTaskUser gets report.project.task.user existing record.
func (c *Client) GetReportProjectTaskUser(id int64) (*ReportProjectTaskUser, error) {
	rptus, err := c.GetReportProjectTaskUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rptus)[0]), nil
}

// GetReportProjectTaskUsers gets report.project.task.user existing records.
func (c *Client) GetReportProjectTaskUsers(ids []int64) (*ReportProjectTaskUsers, error) {
	rptus := &ReportProjectTaskUsers{}
	if err := c.Read(ReportProjectTaskUserModel, ids, nil, rptus); err != nil {
		return nil, err
	}
	return rptus, nil
}

// FindReportProjectTaskUser finds report.project.task.user record by querying it with criteria.
func (c *Client) FindReportProjectTaskUser(criteria *Criteria) (*ReportProjectTaskUser, error) {
	rptus := &ReportProjectTaskUsers{}
	if err := c.SearchRead(ReportProjectTaskUserModel, criteria, NewOptions().Limit(1), rptus); err != nil {
		return nil, err
	}
	return &((*rptus)[0]), nil
}

// FindReportProjectTaskUsers finds report.project.task.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProjectTaskUsers(criteria *Criteria, options *Options) (*ReportProjectTaskUsers, error) {
	rptus := &ReportProjectTaskUsers{}
	if err := c.SearchRead(ReportProjectTaskUserModel, criteria, options, rptus); err != nil {
		return nil, err
	}
	return rptus, nil
}

// FindReportProjectTaskUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProjectTaskUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProjectTaskUserModel, criteria, options)
}

// FindReportProjectTaskUserId finds record id by querying it with criteria.
func (c *Client) FindReportProjectTaskUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProjectTaskUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
