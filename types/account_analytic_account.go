package types

import (
	"time"
)

type AccountAnalyticAccount struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	Active                   bool      `xmlrpc:"active"`
	Balance                  float64   `xmlrpc:"balance"`
	Code                     string    `xmlrpc:"code"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CompanyUomId             Many2One  `xmlrpc:"company_uom_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	Credit                   float64   `xmlrpc:"credit"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	Debit                    float64   `xmlrpc:"debit"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
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
	ProjectCount             int64     `xmlrpc:"project_count"`
	ProjectIds               []int64   `xmlrpc:"project_ids"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	UseTasks                 bool      `xmlrpc:"use_tasks"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type AccountAnalyticAccountNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	Active                   bool        `xmlrpc:"active"`
	Balance                  interface{} `xmlrpc:"balance"`
	Code                     interface{} `xmlrpc:"code"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CompanyUomId             interface{} `xmlrpc:"company_uom_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Credit                   interface{} `xmlrpc:"credit"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	Debit                    interface{} `xmlrpc:"debit"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
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
	ProjectCount             interface{} `xmlrpc:"project_count"`
	ProjectIds               interface{} `xmlrpc:"project_ids"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	UseTasks                 bool        `xmlrpc:"use_tasks"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var AccountAnalyticAccountModel string = "account.analytic.account"

type AccountAnalyticAccounts []AccountAnalyticAccount

type AccountAnalyticAccountsNil []AccountAnalyticAccountNil

func (s *AccountAnalyticAccount) NilableType_() interface{} {
	return &AccountAnalyticAccountNil{}
}

func (ns *AccountAnalyticAccountNil) Type_() interface{} {
	s := &AccountAnalyticAccount{}
	return load(ns, s)
}

func (s *AccountAnalyticAccounts) NilableType_() interface{} {
	return &AccountAnalyticAccountsNil{}
}

func (ns *AccountAnalyticAccountsNil) Type_() interface{} {
	s := &AccountAnalyticAccounts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAnalyticAccount))
	}
	return s
}
