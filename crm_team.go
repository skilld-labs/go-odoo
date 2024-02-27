package odoo

// CrmTeam represents crm.team model.
type CrmTeam struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	AliasContact                 *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId           *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId                 *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId           *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId          *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId                  *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	Color                        *Int       `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty"`
	DashboardButtonName          *String    `xmlrpc:"dashboard_button_name,omitempty"`
	DashboardGraphData           *String    `xmlrpc:"dashboard_graph_data,omitempty"`
	DashboardGraphGroup          *Selection `xmlrpc:"dashboard_graph_group,omitempty"`
	DashboardGraphGroupPipeline  *Selection `xmlrpc:"dashboard_graph_group_pipeline,omitempty"`
	DashboardGraphModel          *Selection `xmlrpc:"dashboard_graph_model,omitempty"`
	DashboardGraphPeriod         *Selection `xmlrpc:"dashboard_graph_period,omitempty"`
	DashboardGraphPeriodPipeline *Selection `xmlrpc:"dashboard_graph_period_pipeline,omitempty"`
	DashboardGraphType           *Selection `xmlrpc:"dashboard_graph_type,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	FavoriteUserIds              *Relation  `xmlrpc:"favorite_user_ids,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	Invoiced                     *Int       `xmlrpc:"invoiced,omitempty"`
	InvoicedTarget               *Int       `xmlrpc:"invoiced_target,omitempty"`
	IsFavorite                   *Bool      `xmlrpc:"is_favorite,omitempty"`
	MemberIds                    *Relation  `xmlrpc:"member_ids,omitempty"`
	MessageChannelIds            *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost              *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter         *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	OpportunitiesAmount          *Int       `xmlrpc:"opportunities_amount,omitempty"`
	OpportunitiesCount           *Int       `xmlrpc:"opportunities_count,omitempty"`
	QuotationsAmount             *Int       `xmlrpc:"quotations_amount,omitempty"`
	QuotationsCount              *Int       `xmlrpc:"quotations_count,omitempty"`
	ReplyTo                      *String    `xmlrpc:"reply_to,omitempty"`
	SalesToInvoiceCount          *Int       `xmlrpc:"sales_to_invoice_count,omitempty"`
	TeamType                     *Selection `xmlrpc:"team_type,omitempty"`
	UnassignedLeadsCount         *Int       `xmlrpc:"unassigned_leads_count,omitempty"`
	UseInvoices                  *Bool      `xmlrpc:"use_invoices,omitempty"`
	UseLeads                     *Bool      `xmlrpc:"use_leads,omitempty"`
	UseOpportunities             *Bool      `xmlrpc:"use_opportunities,omitempty"`
	UseQuotations                *Bool      `xmlrpc:"use_quotations,omitempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateCrmTeams([]*CrmTeam{ct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmTeam creates a new crm.team model and returns its id.
func (c *Client) CreateCrmTeams(cts []*CrmTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range cts {
		vv = append(vv, v)
	}
	return c.Create(CrmTeamModel, vv, nil)
}

// UpdateCrmTeam updates an existing crm.team record.
func (c *Client) UpdateCrmTeam(ct *CrmTeam) error {
	return c.UpdateCrmTeams([]int64{ct.Id.Get()}, ct)
}

// UpdateCrmTeams updates existing crm.team records.
// All records (represented by ids) will be updated by ct values.
func (c *Client) UpdateCrmTeams(ids []int64, ct *CrmTeam) error {
	return c.Update(CrmTeamModel, ids, ct, nil)
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
	return &((*cts)[0]), nil
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
	return &((*cts)[0]), nil
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
	return c.Search(CrmTeamModel, criteria, options)
}

// FindCrmTeamId finds record id by querying it with criteria.
func (c *Client) FindCrmTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
