package odoo

// ProjectProject represents project.project model.
type ProjectProject struct {
	AccessInstructionMessage     *String     `xmlrpc:"access_instruction_message,omitempty"`
	AccessToken                  *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                    *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                *String     `xmlrpc:"access_warning,omitempty"`
	AccountId                    *Many2One   `xmlrpc:"account_id,omitempty"`
	Active                       *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId      *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline         *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary              *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon             *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AliasBouncedContent          *String     `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact                 *Selection  `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                *String     `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                  *String     `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId                *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail                   *String     `xmlrpc:"alias_email,omitempty"`
	AliasForceThreadId           *Int        `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasFullName                *String     `xmlrpc:"alias_full_name,omitempty"`
	AliasId                      *Many2One   `xmlrpc:"alias_id,omitempty"`
	AliasIncomingLocal           *Bool       `xmlrpc:"alias_incoming_local,omitempty"`
	AliasModelId                 *Many2One   `xmlrpc:"alias_model_id,omitempty"`
	AliasName                    *String     `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId           *Many2One   `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId          *Int        `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasStatus                  *Selection  `xmlrpc:"alias_status,omitempty"`
	AllocatedHours               *Float      `xmlrpc:"allocated_hours,omitempty"`
	AllowBillable                *Bool       `xmlrpc:"allow_billable,omitempty"`
	AllowMilestones              *Bool       `xmlrpc:"allow_milestones,omitempty"`
	AllowTaskDependencies        *Bool       `xmlrpc:"allow_task_dependencies,omitempty"`
	AllowTimesheets              *Bool       `xmlrpc:"allow_timesheets,omitempty"`
	AnalyticAccountActive        *Bool       `xmlrpc:"analytic_account_active,omitempty"`
	AnalyticAccountBalance       *Float      `xmlrpc:"analytic_account_balance,omitempty"`
	AutoAccountId                *Many2One   `xmlrpc:"auto_account_id,omitempty"`
	BillingType                  *Selection  `xmlrpc:"billing_type,omitempty"`
	CanMarkMilestoneAsDone       *Bool       `xmlrpc:"can_mark_milestone_as_done,omitempty"`
	ClosedTaskCount              *Int        `xmlrpc:"closed_task_count,omitempty"`
	CollaboratorCount            *Int        `xmlrpc:"collaborator_count,omitempty"`
	CollaboratorIds              *Relation   `xmlrpc:"collaborator_ids,omitempty"`
	Color                        *Int        `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                         *Time       `xmlrpc:"date,omitempty"`
	DateStart                    *Time       `xmlrpc:"date_start,omitempty"`
	Description                  *String     `xmlrpc:"description,omitempty"`
	DisplayName                  *String     `xmlrpc:"display_name,omitempty"`
	DisplaySalesStatButtons      *Bool       `xmlrpc:"display_sales_stat_buttons,omitempty"`
	DurationTracking             interface{} `xmlrpc:"duration_tracking,omitempty"`
	EffectiveHours               *Float      `xmlrpc:"effective_hours,omitempty"`
	EncodeUomInDays              *Bool       `xmlrpc:"encode_uom_in_days,omitempty"`
	FavoriteUserIds              *Relation   `xmlrpc:"favorite_user_ids,omitempty"`
	HasAnySoToInvoice            *Bool       `xmlrpc:"has_any_so_to_invoice,omitempty"`
	HasAnySoWithNothingToInvoice *Bool       `xmlrpc:"has_any_so_with_nothing_to_invoice,omitempty"`
	HasMessage                   *Bool       `xmlrpc:"has_message,omitempty"`
	Id                           *Int        `xmlrpc:"id,omitempty"`
	InvoiceCount                 *Int        `xmlrpc:"invoice_count,omitempty"`
	IsFavorite                   *Bool       `xmlrpc:"is_favorite,omitempty"`
	IsInternalProject            *Bool       `xmlrpc:"is_internal_project,omitempty"`
	IsMilestoneDeadlineExceeded  *Bool       `xmlrpc:"is_milestone_deadline_exceeded,omitempty"`
	IsMilestoneExceeded          *Bool       `xmlrpc:"is_milestone_exceeded,omitempty"`
	IsProjectOvertime            *Bool       `xmlrpc:"is_project_overtime,omitempty"`
	LabelTasks                   *String     `xmlrpc:"label_tasks,omitempty"`
	LastUpdateColor              *Int        `xmlrpc:"last_update_color,omitempty"`
	LastUpdateId                 *Many2One   `xmlrpc:"last_update_id,omitempty"`
	LastUpdateStatus             *Selection  `xmlrpc:"last_update_status,omitempty"`
	MessageAttachmentCount       *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneCount               *Int        `xmlrpc:"milestone_count,omitempty"`
	MilestoneCountReached        *Int        `xmlrpc:"milestone_count_reached,omitempty"`
	MilestoneIds                 *Relation   `xmlrpc:"milestone_ids,omitempty"`
	MilestoneProgress            *Int        `xmlrpc:"milestone_progress,omitempty"`
	MyActivityDateDeadline       *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                         *String     `xmlrpc:"name,omitempty"`
	NextMilestoneId              *Many2One   `xmlrpc:"next_milestone_id,omitempty"`
	OpenTaskCount                *Int        `xmlrpc:"open_task_count,omitempty"`
	PartnerId                    *Many2One   `xmlrpc:"partner_id,omitempty"`
	PricingType                  *Selection  `xmlrpc:"pricing_type,omitempty"`
	PrivacyVisibility            *Selection  `xmlrpc:"privacy_visibility,omitempty"`
	PrivacyVisibilityWarning     *String     `xmlrpc:"privacy_visibility_warning,omitempty"`
	PurchaseOrdersCount          *Int        `xmlrpc:"purchase_orders_count,omitempty"`
	RatingActive                 *Bool       `xmlrpc:"rating_active,omitempty"`
	RatingAvg                    *Float      `xmlrpc:"rating_avg,omitempty"`
	RatingAvgPercentage          *Float      `xmlrpc:"rating_avg_percentage,omitempty"`
	RatingCount                  *Int        `xmlrpc:"rating_count,omitempty"`
	RatingIds                    *Relation   `xmlrpc:"rating_ids,omitempty"`
	RatingPercentageSatisfaction *Int        `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	RatingRequestDeadline        *Time       `xmlrpc:"rating_request_deadline,omitempty"`
	RatingStatus                 *Selection  `xmlrpc:"rating_status,omitempty"`
	RatingStatusPeriod           *Selection  `xmlrpc:"rating_status_period,omitempty"`
	ReinvoicedSaleOrderId        *Many2One   `xmlrpc:"reinvoiced_sale_order_id,omitempty"`
	RemainingHours               *Float      `xmlrpc:"remaining_hours,omitempty"`
	ResourceCalendarId           *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	SaleLineEmployeeIds          *Relation   `xmlrpc:"sale_line_employee_ids,omitempty"`
	SaleLineId                   *Many2One   `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderCount               *Int        `xmlrpc:"sale_order_count,omitempty"`
	SaleOrderId                  *Many2One   `xmlrpc:"sale_order_id,omitempty"`
	SaleOrderLineCount           *Int        `xmlrpc:"sale_order_line_count,omitempty"`
	SaleOrderState               *Selection  `xmlrpc:"sale_order_state,omitempty"`
	Sequence                     *Int        `xmlrpc:"sequence,omitempty"`
	StageId                      *Many2One   `xmlrpc:"stage_id,omitempty"`
	TagIds                       *Relation   `xmlrpc:"tag_ids,omitempty"`
	TaskCompletionPercentage     *Float      `xmlrpc:"task_completion_percentage,omitempty"`
	TaskCount                    *Int        `xmlrpc:"task_count,omitempty"`
	TaskIds                      *Relation   `xmlrpc:"task_ids,omitempty"`
	TaskPropertiesDefinition     interface{} `xmlrpc:"task_properties_definition,omitempty"`
	Tasks                        *Relation   `xmlrpc:"tasks,omitempty"`
	TimesheetEncodeUomId         *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                 *Relation   `xmlrpc:"timesheet_ids,omitempty"`
	TimesheetProductId           *Many2One   `xmlrpc:"timesheet_product_id,omitempty"`
	TotalTimesheetTime           *Int        `xmlrpc:"total_timesheet_time,omitempty"`
	TypeIds                      *Relation   `xmlrpc:"type_ids,omitempty"`
	UpdateCount                  *Int        `xmlrpc:"update_count,omitempty"`
	UpdateIds                    *Relation   `xmlrpc:"update_ids,omitempty"`
	UserId                       *Many2One   `xmlrpc:"user_id,omitempty"`
	VendorBillCount              *Int        `xmlrpc:"vendor_bill_count,omitempty"`
	WarningEmployeeRate          *Bool       `xmlrpc:"warning_employee_rate,omitempty"`
	WebsiteMessageIds            *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One   `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateProjectProjects([]*ProjectProject{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProjects(pps []*ProjectProject) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectModel, vv, nil)
}

// UpdateProjectProject updates an existing project.project record.
func (c *Client) UpdateProjectProject(pp *ProjectProject) error {
	return c.UpdateProjectProjects([]int64{pp.Id.Get()}, pp)
}

// UpdateProjectProjects updates existing project.project records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProjectProjects(ids []int64, pp *ProjectProject) error {
	return c.Update(ProjectProjectModel, ids, pp, nil)
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
	return &((*pps)[0]), nil
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
	return &((*pps)[0]), nil
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
	return c.Search(ProjectProjectModel, criteria, options)
}

// FindProjectProjectId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
