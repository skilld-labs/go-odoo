package odoo

import (
	"fmt"
)

// CrmTeam represents crm.team model.
type CrmTeam struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	AliasContact                 *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults                *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId           *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId                 *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId           *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId          *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId                  *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	Color                        *Int       `xmlrpc:"color,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	DashboardButtonName          *String    `xmlrpc:"dashboard_button_name,omptempty"`
	DashboardGraphData           *String    `xmlrpc:"dashboard_graph_data,omptempty"`
	DashboardGraphGroup          *Selection `xmlrpc:"dashboard_graph_group,omptempty"`
	DashboardGraphGroupPipeline  *Selection `xmlrpc:"dashboard_graph_group_pipeline,omptempty"`
	DashboardGraphModel          *Selection `xmlrpc:"dashboard_graph_model,omptempty"`
	DashboardGraphPeriod         *Selection `xmlrpc:"dashboard_graph_period,omptempty"`
	DashboardGraphPeriodPipeline *Selection `xmlrpc:"dashboard_graph_period_pipeline,omptempty"`
	DashboardGraphType           *Selection `xmlrpc:"dashboard_graph_type,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	FavoriteUserIds              *Relation  `xmlrpc:"favorite_user_ids,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	Invoiced                     *Int       `xmlrpc:"invoiced,omptempty"`
	InvoicedTarget               *Int       `xmlrpc:"invoiced_target,omptempty"`
	IsFavorite                   *Bool      `xmlrpc:"is_favorite,omptempty"`
	MemberIds                    *Relation  `xmlrpc:"member_ids,omptempty"`
	MessageChannelIds            *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost              *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter         *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	OpportunitiesAmount          *Int       `xmlrpc:"opportunities_amount,omptempty"`
	OpportunitiesCount           *Int       `xmlrpc:"opportunities_count,omptempty"`
	QuotationsAmount             *Int       `xmlrpc:"quotations_amount,omptempty"`
	QuotationsCount              *Int       `xmlrpc:"quotations_count,omptempty"`
	ReplyTo                      *String    `xmlrpc:"reply_to,omptempty"`
	SalesToInvoiceCount          *Int       `xmlrpc:"sales_to_invoice_count,omptempty"`
	TeamType                     *Selection `xmlrpc:"team_type,omptempty"`
	UnassignedLeadsCount         *Int       `xmlrpc:"unassigned_leads_count,omptempty"`
	UseInvoices                  *Bool      `xmlrpc:"use_invoices,omptempty"`
	UseLeads                     *Bool      `xmlrpc:"use_leads,omptempty"`
	UseOpportunities             *Bool      `xmlrpc:"use_opportunities,omptempty"`
	UseQuotations                *Bool      `xmlrpc:"use_quotations,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CrmTeams represents array of crm.team model.
type CrmTeams []CrmTeam

// CrmTeamModel is the odoo model name.
const CrmTeamModel = "crm.team"

// Many2One convert CrmTeam to *Many2One.
func (ct *CrmTeam) Many2One() *Many2One {
	return NewMany2One(ct.Id.Get(), "")
}

// CreateCrmTeam creates a new crm.team model and returns its id.
func (c *Client) CreateCrmTeam(ct *CrmTeam) (int64, error) {
	return c.Create(CrmTeamModel, ct)
}

// UpdateCrmTeam updates an existing crm.team record.
func (c *Client) UpdateCrmTeam(ct *CrmTeam) error {
	return c.UpdateCrmTeams([]int64{ct.Id.Get()}, ct)
}

// UpdateCrmTeams updates existing crm.team records.
// All records (represented by ids) will be updated by ct values.
func (c *Client) UpdateCrmTeams(ids []int64, ct *CrmTeam) error {
	return c.Update(CrmTeamModel, ids, ct)
}

// DeleteCrmTeam deletes an existing crm.team record.
func (c *Client) DeleteCrmTeam(id int64) error {
	return c.DeleteCrmTeams([]int64{id})
}

// DeleteCrmTeams deletes existing crm.team records.
func (c *Client) DeleteCrmTeams(ids []int64) error {
	return c.Delete(CrmTeamModel, ids)
}

// GetCrmTeam gets crm.team existing record.
func (c *Client) GetCrmTeam(id int64) (*CrmTeam, error) {
	cts, err := c.GetCrmTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	if cts != nil && len(*cts) > 0 {
		return &((*cts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.team not found", id)
}

// GetCrmTeams gets crm.team existing records.
func (c *Client) GetCrmTeams(ids []int64) (*CrmTeams, error) {
	cts := &CrmTeams{}
	if err := c.Read(CrmTeamModel, ids, nil, cts); err != nil {
		return nil, err
	}
	return cts, nil
}

// FindCrmTeam finds crm.team record by querying it with criteria.
func (c *Client) FindCrmTeam(criteria *Criteria) (*CrmTeam, error) {
	cts := &CrmTeams{}
	if err := c.SearchRead(CrmTeamModel, criteria, NewOptions().Limit(1), cts); err != nil {
		return nil, err
	}
	if cts != nil && len(*cts) > 0 {
		return &((*cts)[0]), nil
	}
	return nil, fmt.Errorf("no crm.team was found with criteria %v", criteria)
}

// FindCrmTeams finds crm.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeams(criteria *Criteria, options *Options) (*CrmTeams, error) {
	cts := &CrmTeams{}
	if err := c.SearchRead(CrmTeamModel, criteria, options, cts); err != nil {
		return nil, err
	}
	return cts, nil
}

// FindCrmTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmTeamModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmTeamId finds record id by querying it with criteria.
func (c *Client) FindCrmTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no crm.team was found with criteria %v and options %v", criteria, options)
}
