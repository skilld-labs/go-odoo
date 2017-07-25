package types

import (
	"time"
)

type ProjectTask struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	Active                   bool      `xmlrpc:"active"`
	AttachmentIds            []int64   `xmlrpc:"attachment_ids"`
	ChildIds                 []int64   `xmlrpc:"child_ids"`
	ChildrenHours            float64   `xmlrpc:"children_hours"`
	Color                    int64     `xmlrpc:"color"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DateAssign               time.Time `xmlrpc:"date_assign"`
	DateDeadline             time.Time `xmlrpc:"date_deadline"`
	DateEnd                  time.Time `xmlrpc:"date_end"`
	DateLastStageUpdate      time.Time `xmlrpc:"date_last_stage_update"`
	DateStart                time.Time `xmlrpc:"date_start"`
	DelayHours               float64   `xmlrpc:"delay_hours"`
	Description              string    `xmlrpc:"description"`
	DisplayName              string    `xmlrpc:"display_name"`
	DisplayedImageId         Many2One  `xmlrpc:"displayed_image_id"`
	EffectiveHours           float64   `xmlrpc:"effective_hours"`
	Id                       int64     `xmlrpc:"id"`
	KanbanState              string    `xmlrpc:"kanban_state"`
	LegendBlocked            string    `xmlrpc:"legend_blocked"`
	LegendDone               string    `xmlrpc:"legend_done"`
	LegendNormal             string    `xmlrpc:"legend_normal"`
	ManagerId                Many2One  `xmlrpc:"manager_id"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	Name                     string    `xmlrpc:"name"`
	Notes                    string    `xmlrpc:"notes"`
	ParentId                 Many2One  `xmlrpc:"parent_id"`
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PlannedHours             float64   `xmlrpc:"planned_hours"`
	Priority                 string    `xmlrpc:"priority"`
	ProcurementId            Many2One  `xmlrpc:"procurement_id"`
	Progress                 float64   `xmlrpc:"progress"`
	ProjectId                Many2One  `xmlrpc:"project_id"`
	RemainingHours           float64   `xmlrpc:"remaining_hours"`
	SaleLineId               Many2One  `xmlrpc:"sale_line_id"`
	Sequence                 int64     `xmlrpc:"sequence"`
	StageId                  Many2One  `xmlrpc:"stage_id"`
	SubtaskCount             int64     `xmlrpc:"subtask_count"`
	SubtaskProjectId         Many2One  `xmlrpc:"subtask_project_id"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	TimesheetIds             []int64   `xmlrpc:"timesheet_ids"`
	TotalHours               float64   `xmlrpc:"total_hours"`
	TotalHoursSpent          float64   `xmlrpc:"total_hours_spent"`
	UserEmail                string    `xmlrpc:"user_email"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type ProjectTaskNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	Active                   bool        `xmlrpc:"active"`
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
	DisplayName              interface{} `xmlrpc:"display_name"`
	DisplayedImageId         interface{} `xmlrpc:"displayed_image_id"`
	EffectiveHours           interface{} `xmlrpc:"effective_hours"`
	Id                       interface{} `xmlrpc:"id"`
	KanbanState              interface{} `xmlrpc:"kanban_state"`
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
	Priority                 interface{} `xmlrpc:"priority"`
	ProcurementId            interface{} `xmlrpc:"procurement_id"`
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
