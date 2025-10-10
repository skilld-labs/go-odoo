package odoo

// ProjectUpdate represents project.update model.
type ProjectUpdate struct {
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllocatedTime               *Int       `xmlrpc:"allocated_time,omitempty"`
	ClosedTaskCount             *Int       `xmlrpc:"closed_task_count,omitempty"`
	ClosedTaskPercentage        *Int       `xmlrpc:"closed_task_percentage,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	DisplayTimesheetStats       *Bool      `xmlrpc:"display_timesheet_stats,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NameCropped                 *String    `xmlrpc:"name_cropped,omitempty"`
	Progress                    *Int       `xmlrpc:"progress,omitempty"`
	ProgressPercentage          *Float     `xmlrpc:"progress_percentage,omitempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	Status                      *Selection `xmlrpc:"status,omitempty"`
	TaskCount                   *Int       `xmlrpc:"task_count,omitempty"`
	TimesheetPercentage         *Int       `xmlrpc:"timesheet_percentage,omitempty"`
	TimesheetTime               *Int       `xmlrpc:"timesheet_time,omitempty"`
	UomId                       *Many2One  `xmlrpc:"uom_id,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProjectUpdates represents array of project.update model.
type ProjectUpdates []ProjectUpdate

// ProjectUpdateModel is the odoo model name.
const ProjectUpdateModel = "project.update"

// Many2One convert ProjectUpdate to *Many2One.
func (pu *ProjectUpdate) Many2One() *Many2One {
	return NewMany2One(pu.Id.Get(), "")
}

// CreateProjectUpdate creates a new project.update model and returns its id.
func (c *Client) CreateProjectUpdate(pu *ProjectUpdate) (int64, error) {
	ids, err := c.CreateProjectUpdates([]*ProjectUpdate{pu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectUpdate creates a new project.update model and returns its id.
func (c *Client) CreateProjectUpdates(pus []*ProjectUpdate) ([]int64, error) {
	var vv []interface{}
	for _, v := range pus {
		vv = append(vv, v)
	}
	return c.Create(ProjectUpdateModel, vv, nil)
}

// UpdateProjectUpdate updates an existing project.update record.
func (c *Client) UpdateProjectUpdate(pu *ProjectUpdate) error {
	return c.UpdateProjectUpdates([]int64{pu.Id.Get()}, pu)
}

// UpdateProjectUpdates updates existing project.update records.
// All records (represented by ids) will be updated by pu values.
func (c *Client) UpdateProjectUpdates(ids []int64, pu *ProjectUpdate) error {
	return c.Update(ProjectUpdateModel, ids, pu, nil)
}

// DeleteProjectUpdate deletes an existing project.update record.
func (c *Client) DeleteProjectUpdate(id int64) error {
	return c.DeleteProjectUpdates([]int64{id})
}

// DeleteProjectUpdates deletes existing project.update records.
func (c *Client) DeleteProjectUpdates(ids []int64) error {
	return c.Delete(ProjectUpdateModel, ids)
}

// GetProjectUpdate gets project.update existing record.
func (c *Client) GetProjectUpdate(id int64) (*ProjectUpdate, error) {
	pus, err := c.GetProjectUpdates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pus)[0]), nil
}

// GetProjectUpdates gets project.update existing records.
func (c *Client) GetProjectUpdates(ids []int64) (*ProjectUpdates, error) {
	pus := &ProjectUpdates{}
	if err := c.Read(ProjectUpdateModel, ids, nil, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProjectUpdate finds project.update record by querying it with criteria.
func (c *Client) FindProjectUpdate(criteria *Criteria) (*ProjectUpdate, error) {
	pus := &ProjectUpdates{}
	if err := c.SearchRead(ProjectUpdateModel, criteria, NewOptions().Limit(1), pus); err != nil {
		return nil, err
	}
	return &((*pus)[0]), nil
}

// FindProjectUpdates finds project.update records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectUpdates(criteria *Criteria, options *Options) (*ProjectUpdates, error) {
	pus := &ProjectUpdates{}
	if err := c.SearchRead(ProjectUpdateModel, criteria, options, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProjectUpdateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectUpdateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectUpdateModel, criteria, options)
}

// FindProjectUpdateId finds record id by querying it with criteria.
func (c *Client) FindProjectUpdateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectUpdateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
