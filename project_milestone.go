package odoo

// ProjectMilestone represents project.milestone model.
type ProjectMilestone struct {
	AllowBillable            *Bool     `xmlrpc:"allow_billable,omitempty"`
	CanBeMarkedAsDone        *Bool     `xmlrpc:"can_be_marked_as_done,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	Deadline                 *Time     `xmlrpc:"deadline,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	DoneTaskCount            *Int      `xmlrpc:"done_task_count,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsDeadlineExceeded       *Bool     `xmlrpc:"is_deadline_exceeded,omitempty"`
	IsDeadlineFuture         *Bool     `xmlrpc:"is_deadline_future,omitempty"`
	IsReached                *Bool     `xmlrpc:"is_reached,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	ProductUom               *Many2One `xmlrpc:"product_uom,omitempty"`
	ProductUomQty            *Float    `xmlrpc:"product_uom_qty,omitempty"`
	ProjectId                *Many2One `xmlrpc:"project_id,omitempty"`
	ProjectPartnerId         *Many2One `xmlrpc:"project_partner_id,omitempty"`
	QuantityPercentage       *Float    `xmlrpc:"quantity_percentage,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	ReachedDate              *Time     `xmlrpc:"reached_date,omitempty"`
	SaleLineDisplayName      *String   `xmlrpc:"sale_line_display_name,omitempty"`
	SaleLineId               *Many2One `xmlrpc:"sale_line_id,omitempty"`
	TaskCount                *Int      `xmlrpc:"task_count,omitempty"`
	TaskIds                  *Relation `xmlrpc:"task_ids,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectMilestones represents array of project.milestone model.
type ProjectMilestones []ProjectMilestone

// ProjectMilestoneModel is the odoo model name.
const ProjectMilestoneModel = "project.milestone"

// Many2One convert ProjectMilestone to *Many2One.
func (pm *ProjectMilestone) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreateProjectMilestone creates a new project.milestone model and returns its id.
func (c *Client) CreateProjectMilestone(pm *ProjectMilestone) (int64, error) {
	ids, err := c.CreateProjectMilestones([]*ProjectMilestone{pm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectMilestone creates a new project.milestone model and returns its id.
func (c *Client) CreateProjectMilestones(pms []*ProjectMilestone) ([]int64, error) {
	var vv []interface{}
	for _, v := range pms {
		vv = append(vv, v)
	}
	return c.Create(ProjectMilestoneModel, vv, nil)
}

// UpdateProjectMilestone updates an existing project.milestone record.
func (c *Client) UpdateProjectMilestone(pm *ProjectMilestone) error {
	return c.UpdateProjectMilestones([]int64{pm.Id.Get()}, pm)
}

// UpdateProjectMilestones updates existing project.milestone records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdateProjectMilestones(ids []int64, pm *ProjectMilestone) error {
	return c.Update(ProjectMilestoneModel, ids, pm, nil)
}

// DeleteProjectMilestone deletes an existing project.milestone record.
func (c *Client) DeleteProjectMilestone(id int64) error {
	return c.DeleteProjectMilestones([]int64{id})
}

// DeleteProjectMilestones deletes existing project.milestone records.
func (c *Client) DeleteProjectMilestones(ids []int64) error {
	return c.Delete(ProjectMilestoneModel, ids)
}

// GetProjectMilestone gets project.milestone existing record.
func (c *Client) GetProjectMilestone(id int64) (*ProjectMilestone, error) {
	pms, err := c.GetProjectMilestones([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// GetProjectMilestones gets project.milestone existing records.
func (c *Client) GetProjectMilestones(ids []int64) (*ProjectMilestones, error) {
	pms := &ProjectMilestones{}
	if err := c.Read(ProjectMilestoneModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProjectMilestone finds project.milestone record by querying it with criteria.
func (c *Client) FindProjectMilestone(criteria *Criteria) (*ProjectMilestone, error) {
	pms := &ProjectMilestones{}
	if err := c.SearchRead(ProjectMilestoneModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// FindProjectMilestones finds project.milestone records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectMilestones(criteria *Criteria, options *Options) (*ProjectMilestones, error) {
	pms := &ProjectMilestones{}
	if err := c.SearchRead(ProjectMilestoneModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProjectMilestoneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectMilestoneIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectMilestoneModel, criteria, options)
}

// FindProjectMilestoneId finds record id by querying it with criteria.
func (c *Client) FindProjectMilestoneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectMilestoneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
