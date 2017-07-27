package types

import (
	"time"
)

type ProjectProject struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	Active                   bool      `xmlrpc:"active"`
	AliasContact             string    `xmlrpc:"alias_contact"`
	AliasDefaults            string    `xmlrpc:"alias_defaults"`
	AliasDomain              string    `xmlrpc:"alias_domain"`
	AliasForceThreadId       int64     `xmlrpc:"alias_force_thread_id"`
	AliasId                  Many2One  `xmlrpc:"alias_id"`
	AliasModel               string    `xmlrpc:"alias_model"`
	AliasModelId             Many2One  `xmlrpc:"alias_model_id"`
	AliasName                string    `xmlrpc:"alias_name"`
	AliasParentModelId       Many2One  `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId      int64     `xmlrpc:"alias_parent_thread_id"`
	AliasUserId              Many2One  `xmlrpc:"alias_user_id"`
	AllowTimesheets          bool      `xmlrpc:"allow_timesheets"`
	AnalyticAccountId        Many2One  `xmlrpc:"analytic_account_id"`
	Balance                  float64   `xmlrpc:"balance"`
	Code                     string    `xmlrpc:"code"`
	Color                    int64     `xmlrpc:"color"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CompanyUomId             Many2One  `xmlrpc:"company_uom_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	Credit                   float64   `xmlrpc:"credit"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	Date                     time.Time `xmlrpc:"date"`
	DateStart                time.Time `xmlrpc:"date_start"`
	Debit                    float64   `xmlrpc:"debit"`
	DisplayName              string    `xmlrpc:"display_name"`
	DocCount                 int64     `xmlrpc:"doc_count"`
	FavoriteUserIds          []int64   `xmlrpc:"favorite_user_ids"`
	Id                       int64     `xmlrpc:"id"`
	IsFavorite               bool      `xmlrpc:"is_favorite"`
	LabelTasks               string    `xmlrpc:"label_tasks"`
	LineIds                  []int64   `xmlrpc:"line_ids"`
	MachineProjectName       string    `xmlrpc:"machine_project_name"`
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
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PrivacyVisibility        string    `xmlrpc:"privacy_visibility"`
	ProjectCount             int64     `xmlrpc:"project_count"`
	ProjectIds               []int64   `xmlrpc:"project_ids"`
	ResourceCalendarId       Many2One  `xmlrpc:"resource_calendar_id"`
	Sequence                 int64     `xmlrpc:"sequence"`
	SubtaskProjectId         Many2One  `xmlrpc:"subtask_project_id"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	TaskCount                int64     `xmlrpc:"task_count"`
	TaskIds                  []int64   `xmlrpc:"task_ids"`
	TaskNeedactionCount      int64     `xmlrpc:"task_needaction_count"`
	Tasks                    []int64   `xmlrpc:"tasks"`
	TypeIds                  []int64   `xmlrpc:"type_ids"`
	UseTasks                 bool      `xmlrpc:"use_tasks"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type ProjectProjectNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	Active                   bool        `xmlrpc:"active"`
	AliasContact             interface{} `xmlrpc:"alias_contact"`
	AliasDefaults            interface{} `xmlrpc:"alias_defaults"`
	AliasDomain              interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId       interface{} `xmlrpc:"alias_force_thread_id"`
	AliasId                  interface{} `xmlrpc:"alias_id"`
	AliasModel               interface{} `xmlrpc:"alias_model"`
	AliasModelId             interface{} `xmlrpc:"alias_model_id"`
	AliasName                interface{} `xmlrpc:"alias_name"`
	AliasParentModelId       interface{} `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId      interface{} `xmlrpc:"alias_parent_thread_id"`
	AliasUserId              interface{} `xmlrpc:"alias_user_id"`
	AllowTimesheets          bool        `xmlrpc:"allow_timesheets"`
	AnalyticAccountId        interface{} `xmlrpc:"analytic_account_id"`
	Balance                  interface{} `xmlrpc:"balance"`
	Code                     interface{} `xmlrpc:"code"`
	Color                    interface{} `xmlrpc:"color"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CompanyUomId             interface{} `xmlrpc:"company_uom_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Credit                   interface{} `xmlrpc:"credit"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	Date                     interface{} `xmlrpc:"date"`
	DateStart                interface{} `xmlrpc:"date_start"`
	Debit                    interface{} `xmlrpc:"debit"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	DocCount                 interface{} `xmlrpc:"doc_count"`
	FavoriteUserIds          interface{} `xmlrpc:"favorite_user_ids"`
	Id                       interface{} `xmlrpc:"id"`
	IsFavorite               bool        `xmlrpc:"is_favorite"`
	LabelTasks               interface{} `xmlrpc:"label_tasks"`
	LineIds                  interface{} `xmlrpc:"line_ids"`
	MachineProjectName       interface{} `xmlrpc:"machine_project_name"`
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
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PrivacyVisibility        interface{} `xmlrpc:"privacy_visibility"`
	ProjectCount             interface{} `xmlrpc:"project_count"`
	ProjectIds               interface{} `xmlrpc:"project_ids"`
	ResourceCalendarId       interface{} `xmlrpc:"resource_calendar_id"`
	Sequence                 interface{} `xmlrpc:"sequence"`
	SubtaskProjectId         interface{} `xmlrpc:"subtask_project_id"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TaskCount                interface{} `xmlrpc:"task_count"`
	TaskIds                  interface{} `xmlrpc:"task_ids"`
	TaskNeedactionCount      interface{} `xmlrpc:"task_needaction_count"`
	Tasks                    interface{} `xmlrpc:"tasks"`
	TypeIds                  interface{} `xmlrpc:"type_ids"`
	UseTasks                 bool        `xmlrpc:"use_tasks"`
	UserId                   interface{} `xmlrpc:"user_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var ProjectProjectModel string = "project.project"

type ProjectProjects []ProjectProject

type ProjectProjectsNil []ProjectProjectNil

func (s *ProjectProject) NilableType_() interface{} {
	return &ProjectProjectNil{}
}

func (ns *ProjectProjectNil) Type_() interface{} {
	s := &ProjectProject{}
	return load(ns, s)
}

func (s *ProjectProjects) NilableType_() interface{} {
	return &ProjectProjectsNil{}
}

func (ns *ProjectProjectsNil) Type_() interface{} {
	s := &ProjectProjects{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectProject))
	}
	return s
}
