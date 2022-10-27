package odoo

import (
	"fmt"
)

// CrmTeam represents crm.team model.
type CrmTeam struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	Color                    *Int      `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omitempty"`
	DashboardButtonName      *String   `xmlrpc:"dashboard_button_name,omitempty"`
	DashboardGraphData       *String   `xmlrpc:"dashboard_graph_data,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	FavoriteUserIds          *Relation `xmlrpc:"favorite_user_ids,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	Invoiced                 *Float    `xmlrpc:"invoiced,omitempty"`
	InvoicedTarget           *Float    `xmlrpc:"invoiced_target,omitempty"`
	IsFavorite               *Bool     `xmlrpc:"is_favorite,omitempty"`
	MemberIds                *Relation `xmlrpc:"member_ids,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	QuotationsAmount         *Float    `xmlrpc:"quotations_amount,omitempty"`
	QuotationsCount          *Int      `xmlrpc:"quotations_count,omitempty"`
	SalesToInvoiceCount      *Int      `xmlrpc:"sales_to_invoice_count,omitempty"`
	Sequence                 *Int      `xmlrpc:"sequence,omitempty"`
	UseQuotations            *Bool     `xmlrpc:"use_quotations,omitempty"`
	UserId                   *Many2One `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return nil, fmt.Errorf("crm.team was not found")
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
	return -1, fmt.Errorf("crm.team was not found")
}
