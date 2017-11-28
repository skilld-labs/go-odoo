package types

import (
	"time"
)

type ProjectTask struct {
	Active                   bool        `xmlrpc:"active"`
	ActivityDateDeadline     time.Time   `xmlrpc:"activity_date_deadline"`
	ActivityIds              []int64     `xmlrpc:"activity_ids"`
	ActivityState            interface{} `xmlrpc:"activity_state"`
	ActivitySummary          string      `xmlrpc:"activity_summary"`
	ActivityTypeId           Many2One    `xmlrpc:"activity_type_id"`
	ActivityUserId           Many2One    `xmlrpc:"activity_user_id"`
	AttachmentIds            []int64     `xmlrpc:"attachment_ids"`
	ChildIds                 []int64     `xmlrpc:"child_ids"`
	ChildrenHours            float64     `xmlrpc:"children_hours"`
	Color                    int64       `xmlrpc:"color"`
	CompanyId                Many2One    `xmlrpc:"company_id"`
	CreateDate               time.Time   `xmlrpc:"create_date"`
	CreateUid                Many2One    `xmlrpc:"create_uid"`
	DateAssign               time.Time   `xmlrpc:"date_assign"`
	DateDeadline             time.Time   `xmlrpc:"date_deadline"`
	DateEnd                  time.Time   `xmlrpc:"date_end"`
	DateLastStageUpdate      time.Time   `xmlrpc:"date_last_stage_update"`
	DateStart                time.Time   `xmlrpc:"date_start"`
	DelayHours               float64     `xmlrpc:"delay_hours"`
	Description              string      `xmlrpc:"description"`
	DisplayedImageId         Many2One    `xmlrpc:"displayed_image_id"`
	DisplayName              string      `xmlrpc:"display_name"`
	EffectiveHours           float64     `xmlrpc:"effective_hours"`
	EmailCc                  string      `xmlrpc:"email_cc"`
	EmailFrom                string      `xmlrpc:"email_from"`
	Id                       int64       `xmlrpc:"id"`
	KanbanState              interface{} `xmlrpc:"kanban_state"`
	KanbanStateLabel         string      `xmlrpc:"kanban_state_label"`
	LastUpdate               time.Time   `xmlrpc:"__last_update"`
	LegendBlocked            string      `xmlrpc:"legend_blocked"`
	LegendDone               string      `xmlrpc:"legend_done"`
	LegendNormal             string      `xmlrpc:"legend_normal"`
	ManagerId                Many2One    `xmlrpc:"manager_id"`
	MessageChannelIds        []int64     `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64     `xmlrpc:"message_follower_ids"`
	MessageIds               []int64     `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time   `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64       `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64     `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64       `xmlrpc:"message_unread_counter"`
	Name                     string      `xmlrpc:"name"`
	Notes                    string      `xmlrpc:"notes"`
	ParentId                 Many2One    `xmlrpc:"parent_id"`
	PartnerId                Many2One    `xmlrpc:"partner_id"`
	PlannedHours             float64     `xmlrpc:"planned_hours"`
	PortalUrl                string      `xmlrpc:"portal_url"`
	Priority                 interface{} `xmlrpc:"priority"`
	Progress                 float64     `xmlrpc:"progress"`
	ProjectId                Many2One    `xmlrpc:"project_id"`
	RemainingHours           float64     `xmlrpc:"remaining_hours"`
	SaleLineId               Many2One    `xmlrpc:"sale_line_id"`
	Sequence                 int64       `xmlrpc:"sequence"`
	StageId                  Many2One    `xmlrpc:"stage_id"`
	SubtaskCount             int64       `xmlrpc:"subtask_count"`
	SubtaskProjectId         Many2One    `xmlrpc:"subtask_project_id"`
	TagIds                   []int64     `xmlrpc:"tag_ids"`
	TimesheetIds             []int64     `xmlrpc:"timesheet_ids"`
	TotalHours               float64     `xmlrpc:"total_hours"`
	TotalHoursSpent          float64     `xmlrpc:"total_hours_spent"`
	UserEmail                string      `xmlrpc:"user_email"`
	UserId                   Many2One    `xmlrpc:"user_id"`
	WebsiteMessageIds        []int64     `xmlrpc:"website_message_ids"`
	WorkingDaysClose         float64     `xmlrpc:"working_days_close"`
	WorkingDaysOpen          float64     `xmlrpc:"working_days_open"`
	WorkingHoursClose        float64     `xmlrpc:"working_hours_close"`
	WorkingHoursOpen         float64     `xmlrpc:"working_hours_open"`
	WriteDate                time.Time   `xmlrpc:"write_date"`
	WriteUid                 Many2One    `xmlrpc:"write_uid"`
}

type ProjectTaskNil struct {
	Active                   bool        `xmlrpc:"active"`
	ActivityDateDeadline     interface{} `xmlrpc:"activity_date_deadline"`
	ActivityIds              interface{} `xmlrpc:"activity_ids"`
	ActivityState            interface{} `xmlrpc:"activity_state"`
	ActivitySummary          interface{} `xmlrpc:"activity_summary"`
	ActivityTypeId           interface{} `xmlrpc:"activity_type_id"`
	ActivityUserId           interface{} `xmlrpc:"activity_user_id"`
	AttachmentIds            interface{} `xmlrpc:"attachment_ids"`
	ChildIds                 interface{} `xmlrpc:"child_ids"`
	ChildrenHours            interface{} `xmlrpc:"children_hours"`
	Color                    interface{} `xmlrpc:"color"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DateAssign               interface{} `xmlrpc:"date_assign"`
	DateDeadline             interface{} `xmlrpc:"date_deadline"`
	DateEnd                  interface{} `xmlrpc:"date_end"`
	DateLastStageUpdate      interface{} `xmlrpc:"date_last_stage_update"`
	DateStart                interface{} `xmlrpc:"date_start"`
	DelayHours               interface{} `xmlrpc:"delay_hours"`
	Description              interface{} `xmlrpc:"description"`
	DisplayedImageId         interface{} `xmlrpc:"displayed_image_id"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	EffectiveHours           interface{} `xmlrpc:"effective_hours"`
	EmailCc                  interface{} `xmlrpc:"email_cc"`
	EmailFrom                interface{} `xmlrpc:"email_from"`
	Id                       interface{} `xmlrpc:"id"`
	KanbanState              interface{} `xmlrpc:"kanban_state"`
	KanbanStateLabel         interface{} `xmlrpc:"kanban_state_label"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	LegendBlocked            interface{} `xmlrpc:"legend_blocked"`
	LegendDone               interface{} `xmlrpc:"legend_done"`
	LegendNormal             interface{} `xmlrpc:"legend_normal"`
	ManagerId                interface{} `xmlrpc:"manager_id"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	Name                     interface{} `xmlrpc:"name"`
	Notes                    interface{} `xmlrpc:"notes"`
	ParentId                 interface{} `xmlrpc:"parent_id"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PlannedHours             interface{} `xmlrpc:"planned_hours"`
	PortalUrl                interface{} `xmlrpc:"portal_url"`
	Priority                 interface{} `xmlrpc:"priority"`
	Progress                 interface{} `xmlrpc:"progress"`
	ProjectId                interface{} `xmlrpc:"project_id"`
	RemainingHours           interface{} `xmlrpc:"remaining_hours"`
	SaleLineId               interface{} `xmlrpc:"sale_line_id"`
	Sequence                 interface{} `xmlrpc:"sequence"`
	StageId                  interface{} `xmlrpc:"stage_id"`
	SubtaskCount             interface{} `xmlrpc:"subtask_count"`
	SubtaskProjectId         interface{} `xmlrpc:"subtask_project_id"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TimesheetIds             interface{} `xmlrpc:"timesheet_ids"`
	TotalHours               interface{} `xmlrpc:"total_hours"`
	TotalHoursSpent          interface{} `xmlrpc:"total_hours_spent"`
	UserEmail                interface{} `xmlrpc:"user_email"`
	UserId                   interface{} `xmlrpc:"user_id"`
	WebsiteMessageIds        interface{} `xmlrpc:"website_message_ids"`
	WorkingDaysClose         interface{} `xmlrpc:"working_days_close"`
	WorkingDaysOpen          interface{} `xmlrpc:"working_days_open"`
	WorkingHoursClose        interface{} `xmlrpc:"working_hours_close"`
	WorkingHoursOpen         interface{} `xmlrpc:"working_hours_open"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var ProjectTaskModel string = "project.task"

type ProjectTasks []ProjectTask

type ProjectTasksNil []ProjectTaskNil

func (s *ProjectTask) NilableType_() interface{} {
	return &ProjectTaskNil{}
}

func (ns *ProjectTaskNil) Type_() interface{} {
	s := &ProjectTask{}
	return load(ns, s)
}

func (s *ProjectTasks) NilableType_() interface{} {
	return &ProjectTasksNil{}
}

func (ns *ProjectTasksNil) Type_() interface{} {
	s := &ProjectTasks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectTask))
	}
	return s
}
