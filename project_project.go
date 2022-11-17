package odoo

import (
	"fmt"
)

// ProjectProject represents project.project model.
type ProjectProject struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	AllowTimesheets          *Bool      `xmlrpc:"allow_timesheets,omitempty"`
	AnalyticAccountId        *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	Balance                  *Float     `xmlrpc:"balance,omitempty"`
	Code                     *String    `xmlrpc:"code,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyUomId             *Many2One  `xmlrpc:"company_uom_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Credit                   *Float     `xmlrpc:"credit,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty"`
	DateStart                *Time      `xmlrpc:"date_start,omitempty"`
	Debit                    *Float     `xmlrpc:"debit,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	DocCount                 *Int       `xmlrpc:"doc_count,omitempty"`
	FavoriteUserIds          *Relation  `xmlrpc:"favorite_user_ids,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsFavorite               *Bool      `xmlrpc:"is_favorite,omitempty"`
	LabelTasks               *String    `xmlrpc:"label_tasks,omitempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omitempty"`
	MachineProjectName       *String    `xmlrpc:"machine_project_name,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PortalUrl                *String    `xmlrpc:"portal_url,omitempty"`
	PrivacyVisibility        *Selection `xmlrpc:"privacy_visibility,omitempty"`
	ProjectCount             *Int       `xmlrpc:"project_count,omitempty"`
	ProjectCreated           *Bool      `xmlrpc:"project_created,omitempty"`
	ProjectIds               *Relation  `xmlrpc:"project_ids,omitempty"`
	ResourceCalendarId       *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SaleLineId               *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	SubtaskProjectId         *Many2One  `xmlrpc:"subtask_project_id,omitempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omitempty"`
	TaskCount                *Int       `xmlrpc:"task_count,omitempty"`
	TaskIds                  *Relation  `xmlrpc:"task_ids,omitempty"`
	TaskNeedactionCount      *Int       `xmlrpc:"task_needaction_count,omitempty"`
	Tasks                    *Relation  `xmlrpc:"tasks,omitempty"`
	TypeIds                  *Relation  `xmlrpc:"type_ids,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProjectProjects represents array of project.project model.
type ProjectProjects []ProjectProject

// ProjectProjectModel is the odoo model name.
const ProjectProjectModel = "project.project"

// Many2One convert ProjectProject to *Many2One.
func (pp *ProjectProject) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProject(pp *ProjectProject) (int64, error) {
	return c.Create(ProjectProjectModel, pp)
}

// UpdateProjectProject updates an existing project.project record.
func (c *Client) UpdateProjectProject(pp *ProjectProject) error {
	return c.UpdateProjectProjects([]int64{pp.Id.Get()}, pp)
}

// UpdateProjectProjects updates existing project.project records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProjectProjects(ids []int64, pp *ProjectProject) error {
	return c.Update(ProjectProjectModel, ids, pp)
}

// DeleteProjectProject deletes an existing project.project record.
func (c *Client) DeleteProjectProject(id int64) error {
	return c.DeleteProjectProjects([]int64{id})
}

// DeleteProjectProjects deletes existing project.project records.
func (c *Client) DeleteProjectProjects(ids []int64) error {
	return c.Delete(ProjectProjectModel, ids)
}

// GetProjectProject gets project.project existing record.
func (c *Client) GetProjectProject(id int64) (*ProjectProject, error) {
	pps, err := c.GetProjectProjects([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.project not found", id)
}

// GetProjectProjects gets project.project existing records.
func (c *Client) GetProjectProjects(ids []int64) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.Read(ProjectProjectModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProject finds project.project record by querying it with criteria.
func (c *Client) FindProjectProject(criteria *Criteria) (*ProjectProject, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("no project.project was found with criteria %v", criteria)
}

// FindProjectProjects finds project.project records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjects(criteria *Criteria, options *Options) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProjectIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectProjectId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no project.project was found with criteria %v and options %v", criteria, options)
}
