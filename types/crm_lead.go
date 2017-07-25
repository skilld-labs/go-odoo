package types

import (
	"time"
)

type CrmLead struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	Active                   bool      `xmlrpc:"active"`
	CampaignId               Many2One  `xmlrpc:"campaign_id"`
	City                     string    `xmlrpc:"city"`
	Color                    int64     `xmlrpc:"color"`
	CompanyCurrency          Many2One  `xmlrpc:"company_currency"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	ContactName              string    `xmlrpc:"contact_name"`
	CountryId                Many2One  `xmlrpc:"country_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DateAction               time.Time `xmlrpc:"date_action"`
	DateActionLast           time.Time `xmlrpc:"date_action_last"`
	DateActionNext           time.Time `xmlrpc:"date_action_next"`
	DateClosed               time.Time `xmlrpc:"date_closed"`
	DateConversion           time.Time `xmlrpc:"date_conversion"`
	DateDeadline             time.Time `xmlrpc:"date_deadline"`
	DateLastStageUpdate      time.Time `xmlrpc:"date_last_stage_update"`
	DateOpen                 time.Time `xmlrpc:"date_open"`
	DayClose                 float64   `xmlrpc:"day_close"`
	DayOpen                  float64   `xmlrpc:"day_open"`
	Description              string    `xmlrpc:"description"`
	DisplayName              string    `xmlrpc:"display_name"`
	EmailCc                  string    `xmlrpc:"email_cc"`
	EmailFrom                string    `xmlrpc:"email_from"`
	Fax                      string    `xmlrpc:"fax"`
	Function                 string    `xmlrpc:"function"`
	Id                       int64     `xmlrpc:"id"`
	KanbanState              string    `xmlrpc:"kanban_state"`
	LostReason               Many2One  `xmlrpc:"lost_reason"`
	MachineLeadName          string    `xmlrpc:"machine_lead_name"`
	MediumId                 Many2One  `xmlrpc:"medium_id"`
	MeetingCount             int64     `xmlrpc:"meeting_count"`
	MessageBounce            int64     `xmlrpc:"message_bounce"`
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
	Mobile                   string    `xmlrpc:"mobile"`
	Name                     string    `xmlrpc:"name"`
	NextActivityId           Many2One  `xmlrpc:"next_activity_id"`
	OptOut                   bool      `xmlrpc:"opt_out"`
	OrderIds                 []int64   `xmlrpc:"order_ids"`
	PartnerAddressEmail      string    `xmlrpc:"partner_address_email"`
	PartnerAddressName       string    `xmlrpc:"partner_address_name"`
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PartnerName              string    `xmlrpc:"partner_name"`
	Phone                    string    `xmlrpc:"phone"`
	PlannedRevenue           float64   `xmlrpc:"planned_revenue"`
	Priority                 string    `xmlrpc:"priority"`
	Probability              float64   `xmlrpc:"probability"`
	Referred                 string    `xmlrpc:"referred"`
	SaleAmountTotal          float64   `xmlrpc:"sale_amount_total"`
	SaleNumber               int64     `xmlrpc:"sale_number"`
	SourceId                 Many2One  `xmlrpc:"source_id"`
	StageId                  Many2One  `xmlrpc:"stage_id"`
	StateId                  Many2One  `xmlrpc:"state_id"`
	Street                   string    `xmlrpc:"street"`
	Street2                  string    `xmlrpc:"street2"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	TeamId                   Many2One  `xmlrpc:"team_id"`
	Title                    Many2One  `xmlrpc:"title"`
	TitleAction              string    `xmlrpc:"title_action"`
	Type                     string    `xmlrpc:"type"`
	UserEmail                string    `xmlrpc:"user_email"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	UserLogin                string    `xmlrpc:"user_login"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
	Zip                      string    `xmlrpc:"zip"`
}

type CrmLeadNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	Active                   bool        `xmlrpc:"active"`
	CampaignId               interface{} `xmlrpc:"campaign_id"`
	City                     interface{} `xmlrpc:"city"`
	Color                    interface{} `xmlrpc:"color"`
	CompanyCurrency          interface{} `xmlrpc:"company_currency"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	ContactName              interface{} `xmlrpc:"contact_name"`
	CountryId                interface{} `xmlrpc:"country_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DateAction               interface{} `xmlrpc:"date_action"`
	DateActionLast           interface{} `xmlrpc:"date_action_last"`
	DateActionNext           interface{} `xmlrpc:"date_action_next"`
	DateClosed               interface{} `xmlrpc:"date_closed"`
	DateConversion           interface{} `xmlrpc:"date_conversion"`
	DateDeadline             interface{} `xmlrpc:"date_deadline"`
	DateLastStageUpdate      interface{} `xmlrpc:"date_last_stage_update"`
	DateOpen                 interface{} `xmlrpc:"date_open"`
	DayClose                 interface{} `xmlrpc:"day_close"`
	DayOpen                  interface{} `xmlrpc:"day_open"`
	Description              interface{} `xmlrpc:"description"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	EmailCc                  interface{} `xmlrpc:"email_cc"`
	EmailFrom                interface{} `xmlrpc:"email_from"`
	Fax                      interface{} `xmlrpc:"fax"`
	Function                 interface{} `xmlrpc:"function"`
	Id                       interface{} `xmlrpc:"id"`
	KanbanState              interface{} `xmlrpc:"kanban_state"`
	LostReason               interface{} `xmlrpc:"lost_reason"`
	MachineLeadName          interface{} `xmlrpc:"machine_lead_name"`
	MediumId                 interface{} `xmlrpc:"medium_id"`
	MeetingCount             interface{} `xmlrpc:"meeting_count"`
	MessageBounce            interface{} `xmlrpc:"message_bounce"`
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
	Mobile                   interface{} `xmlrpc:"mobile"`
	Name                     interface{} `xmlrpc:"name"`
	NextActivityId           interface{} `xmlrpc:"next_activity_id"`
	OptOut                   bool        `xmlrpc:"opt_out"`
	OrderIds                 interface{} `xmlrpc:"order_ids"`
	PartnerAddressEmail      interface{} `xmlrpc:"partner_address_email"`
	PartnerAddressName       interface{} `xmlrpc:"partner_address_name"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PartnerName              interface{} `xmlrpc:"partner_name"`
	Phone                    interface{} `xmlrpc:"phone"`
	PlannedRevenue           interface{} `xmlrpc:"planned_revenue"`
	Priority                 interface{} `xmlrpc:"priority"`
	Probability              interface{} `xmlrpc:"probability"`
	Referred                 interface{} `xmlrpc:"referred"`
	SaleAmountTotal          interface{} `xmlrpc:"sale_amount_total"`
	SaleNumber               interface{} `xmlrpc:"sale_number"`
	SourceId                 interface{} `xmlrpc:"source_id"`
	StageId                  interface{} `xmlrpc:"stage_id"`
	StateId                  interface{} `xmlrpc:"state_id"`
	Street                   interface{} `xmlrpc:"street"`
	Street2                  interface{} `xmlrpc:"street2"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TeamId                   interface{} `xmlrpc:"team_id"`
	Title                    interface{} `xmlrpc:"title"`
	TitleAction              interface{} `xmlrpc:"title_action"`
	Type                     interface{} `xmlrpc:"type"`
	UserEmail                interface{} `xmlrpc:"user_email"`
	UserId                   interface{} `xmlrpc:"user_id"`
	UserLogin                interface{} `xmlrpc:"user_login"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
	Zip                      interface{} `xmlrpc:"zip"`
}

var CrmLeadModel string = "crm.lead"

type CrmLeads []CrmLead

type CrmLeadsNil []CrmLeadNil

func (s *CrmLead) NilableType_() interface{} {
	return &CrmLeadNil{}
}

func (ns *CrmLeadNil) Type_() interface{} {
	s := &CrmLead{}
	return load(ns, s)
}

func (s *CrmLeads) NilableType_() interface{} {
	return &CrmLeadsNil{}
}

func (ns *CrmLeadsNil) Type_() interface{} {
	s := &CrmLeads{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLead))
	}
	return s
}
