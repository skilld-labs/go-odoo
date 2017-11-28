package types

import (
	"time"
)

type CrmTeam struct {
	Active                       bool        `xmlrpc:"active"`
	AliasContact                 interface{} `xmlrpc:"alias_contact"`
	AliasDefaults                string      `xmlrpc:"alias_defaults"`
	AliasDomain                  string      `xmlrpc:"alias_domain"`
	AliasForceThreadId           int64       `xmlrpc:"alias_force_thread_id"`
	AliasId                      Many2One    `xmlrpc:"alias_id"`
	AliasModelId                 Many2One    `xmlrpc:"alias_model_id"`
	AliasName                    string      `xmlrpc:"alias_name"`
	AliasParentModelId           Many2One    `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId          int64       `xmlrpc:"alias_parent_thread_id"`
	AliasUserId                  Many2One    `xmlrpc:"alias_user_id"`
	Color                        int64       `xmlrpc:"color"`
	CompanyId                    Many2One    `xmlrpc:"company_id"`
	CreateDate                   time.Time   `xmlrpc:"create_date"`
	CreateUid                    Many2One    `xmlrpc:"create_uid"`
	CurrencyId                   Many2One    `xmlrpc:"currency_id"`
	DashboardButtonName          string      `xmlrpc:"dashboard_button_name"`
	DashboardGraphData           string      `xmlrpc:"dashboard_graph_data"`
	DashboardGraphGroup          interface{} `xmlrpc:"dashboard_graph_group"`
	DashboardGraphGroupPipeline  interface{} `xmlrpc:"dashboard_graph_group_pipeline"`
	DashboardGraphModel          interface{} `xmlrpc:"dashboard_graph_model"`
	DashboardGraphPeriod         interface{} `xmlrpc:"dashboard_graph_period"`
	DashboardGraphPeriodPipeline interface{} `xmlrpc:"dashboard_graph_period_pipeline"`
	DashboardGraphType           interface{} `xmlrpc:"dashboard_graph_type"`
	DisplayName                  string      `xmlrpc:"display_name"`
	FavoriteUserIds              []int64     `xmlrpc:"favorite_user_ids"`
	Id                           int64       `xmlrpc:"id"`
	Invoiced                     int64       `xmlrpc:"invoiced"`
	InvoicedTarget               int64       `xmlrpc:"invoiced_target"`
	IsFavorite                   bool        `xmlrpc:"is_favorite"`
	LastUpdate                   time.Time   `xmlrpc:"__last_update"`
	MemberIds                    []int64     `xmlrpc:"member_ids"`
	MessageChannelIds            []int64     `xmlrpc:"message_channel_ids"`
	MessageFollowerIds           []int64     `xmlrpc:"message_follower_ids"`
	MessageIds                   []int64     `xmlrpc:"message_ids"`
	MessageIsFollower            bool        `xmlrpc:"message_is_follower"`
	MessageLastPost              time.Time   `xmlrpc:"message_last_post"`
	MessageNeedaction            bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter     int64       `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds            []int64     `xmlrpc:"message_partner_ids"`
	MessageUnread                bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter         int64       `xmlrpc:"message_unread_counter"`
	Name                         string      `xmlrpc:"name"`
	OpportunitiesAmount          int64       `xmlrpc:"opportunities_amount"`
	OpportunitiesCount           int64       `xmlrpc:"opportunities_count"`
	QuotationsAmount             int64       `xmlrpc:"quotations_amount"`
	QuotationsCount              int64       `xmlrpc:"quotations_count"`
	ReplyTo                      string      `xmlrpc:"reply_to"`
	SalesToInvoiceCount          int64       `xmlrpc:"sales_to_invoice_count"`
	TeamType                     interface{} `xmlrpc:"team_type"`
	UnassignedLeadsCount         int64       `xmlrpc:"unassigned_leads_count"`
	UseInvoices                  bool        `xmlrpc:"use_invoices"`
	UseLeads                     bool        `xmlrpc:"use_leads"`
	UseOpportunities             bool        `xmlrpc:"use_opportunities"`
	UseQuotations                bool        `xmlrpc:"use_quotations"`
	UserId                       Many2One    `xmlrpc:"user_id"`
	WebsiteMessageIds            []int64     `xmlrpc:"website_message_ids"`
	WriteDate                    time.Time   `xmlrpc:"write_date"`
	WriteUid                     Many2One    `xmlrpc:"write_uid"`
}

type CrmTeamNil struct {
	Active                       bool        `xmlrpc:"active"`
	AliasContact                 interface{} `xmlrpc:"alias_contact"`
	AliasDefaults                interface{} `xmlrpc:"alias_defaults"`
	AliasDomain                  interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId           interface{} `xmlrpc:"alias_force_thread_id"`
	AliasId                      interface{} `xmlrpc:"alias_id"`
	AliasModelId                 interface{} `xmlrpc:"alias_model_id"`
	AliasName                    interface{} `xmlrpc:"alias_name"`
	AliasParentModelId           interface{} `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId          interface{} `xmlrpc:"alias_parent_thread_id"`
	AliasUserId                  interface{} `xmlrpc:"alias_user_id"`
	Color                        interface{} `xmlrpc:"color"`
	CompanyId                    interface{} `xmlrpc:"company_id"`
	CreateDate                   interface{} `xmlrpc:"create_date"`
	CreateUid                    interface{} `xmlrpc:"create_uid"`
	CurrencyId                   interface{} `xmlrpc:"currency_id"`
	DashboardButtonName          interface{} `xmlrpc:"dashboard_button_name"`
	DashboardGraphData           interface{} `xmlrpc:"dashboard_graph_data"`
	DashboardGraphGroup          interface{} `xmlrpc:"dashboard_graph_group"`
	DashboardGraphGroupPipeline  interface{} `xmlrpc:"dashboard_graph_group_pipeline"`
	DashboardGraphModel          interface{} `xmlrpc:"dashboard_graph_model"`
	DashboardGraphPeriod         interface{} `xmlrpc:"dashboard_graph_period"`
	DashboardGraphPeriodPipeline interface{} `xmlrpc:"dashboard_graph_period_pipeline"`
	DashboardGraphType           interface{} `xmlrpc:"dashboard_graph_type"`
	DisplayName                  interface{} `xmlrpc:"display_name"`
	FavoriteUserIds              interface{} `xmlrpc:"favorite_user_ids"`
	Id                           interface{} `xmlrpc:"id"`
	Invoiced                     interface{} `xmlrpc:"invoiced"`
	InvoicedTarget               interface{} `xmlrpc:"invoiced_target"`
	IsFavorite                   bool        `xmlrpc:"is_favorite"`
	LastUpdate                   interface{} `xmlrpc:"__last_update"`
	MemberIds                    interface{} `xmlrpc:"member_ids"`
	MessageChannelIds            interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds           interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                   interface{} `xmlrpc:"message_ids"`
	MessageIsFollower            bool        `xmlrpc:"message_is_follower"`
	MessageLastPost              interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction            bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter     interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds            interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread                bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter         interface{} `xmlrpc:"message_unread_counter"`
	Name                         interface{} `xmlrpc:"name"`
	OpportunitiesAmount          interface{} `xmlrpc:"opportunities_amount"`
	OpportunitiesCount           interface{} `xmlrpc:"opportunities_count"`
	QuotationsAmount             interface{} `xmlrpc:"quotations_amount"`
	QuotationsCount              interface{} `xmlrpc:"quotations_count"`
	ReplyTo                      interface{} `xmlrpc:"reply_to"`
	SalesToInvoiceCount          interface{} `xmlrpc:"sales_to_invoice_count"`
	TeamType                     interface{} `xmlrpc:"team_type"`
	UnassignedLeadsCount         interface{} `xmlrpc:"unassigned_leads_count"`
	UseInvoices                  bool        `xmlrpc:"use_invoices"`
	UseLeads                     bool        `xmlrpc:"use_leads"`
	UseOpportunities             bool        `xmlrpc:"use_opportunities"`
	UseQuotations                bool        `xmlrpc:"use_quotations"`
	UserId                       interface{} `xmlrpc:"user_id"`
	WebsiteMessageIds            interface{} `xmlrpc:"website_message_ids"`
	WriteDate                    interface{} `xmlrpc:"write_date"`
	WriteUid                     interface{} `xmlrpc:"write_uid"`
}

var CrmTeamModel string = "crm.team"

type CrmTeams []CrmTeam

type CrmTeamsNil []CrmTeamNil

func (s *CrmTeam) NilableType_() interface{} {
	return &CrmTeamNil{}
}

func (ns *CrmTeamNil) Type_() interface{} {
	s := &CrmTeam{}
	return load(ns, s)
}

func (s *CrmTeams) NilableType_() interface{} {
	return &CrmTeamsNil{}
}

func (ns *CrmTeamsNil) Type_() interface{} {
	s := &CrmTeams{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmTeam))
	}
	return s
}
