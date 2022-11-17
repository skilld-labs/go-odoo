package odoo

import (
	"fmt"
)

// CrmLead represents crm.lead model.
type CrmLead struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline     *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState            *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary          *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omitempty"`
	City                     *String    `xmlrpc:"city,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty"`
	CompanyCurrency          *Many2One  `xmlrpc:"company_currency,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	ContactName              *String    `xmlrpc:"contact_name,omitempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateActionLast           *Time      `xmlrpc:"date_action_last,omitempty"`
	DateClosed               *Time      `xmlrpc:"date_closed,omitempty"`
	DateConversion           *Time      `xmlrpc:"date_conversion,omitempty"`
	DateDeadline             *Time      `xmlrpc:"date_deadline,omitempty"`
	DateLastStageUpdate      *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	DateOpen                 *Time      `xmlrpc:"date_open,omitempty"`
	DayClose                 *Float     `xmlrpc:"day_close,omitempty"`
	DayOpen                  *Float     `xmlrpc:"day_open,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	EmailCc                  *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom                *String    `xmlrpc:"email_from,omitempty"`
	Function                 *String    `xmlrpc:"function,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	KanbanState              *Selection `xmlrpc:"kanban_state,omitempty"`
	LostReason               *Many2One  `xmlrpc:"lost_reason,omitempty"`
	MachineLeadName          *String    `xmlrpc:"machine_lead_name,omitempty"`
	MediumId                 *Many2One  `xmlrpc:"medium_id,omitempty"`
	MeetingCount             *Int       `xmlrpc:"meeting_count,omitempty"`
	MessageBounce            *Int       `xmlrpc:"message_bounce,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Mobile                   *String    `xmlrpc:"mobile,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	OptOut                   *Bool      `xmlrpc:"opt_out,omitempty"`
	OrderIds                 *Relation  `xmlrpc:"order_ids,omitempty"`
	PartnerAddressEmail      *String    `xmlrpc:"partner_address_email,omitempty"`
	PartnerAddressName       *String    `xmlrpc:"partner_address_name,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerName              *String    `xmlrpc:"partner_name,omitempty"`
	Phone                    *String    `xmlrpc:"phone,omitempty"`
	PlannedRevenue           *Float     `xmlrpc:"planned_revenue,omitempty"`
	Priority                 *Selection `xmlrpc:"priority,omitempty"`
	Probability              *Float     `xmlrpc:"probability,omitempty"`
	Referred                 *String    `xmlrpc:"referred,omitempty"`
	SaleAmountTotal          *Float     `xmlrpc:"sale_amount_total,omitempty"`
	SaleNumber               *Int       `xmlrpc:"sale_number,omitempty"`
	SourceId                 *Many2One  `xmlrpc:"source_id,omitempty"`
	StageId                  *Many2One  `xmlrpc:"stage_id,omitempty"`
	StateId                  *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                   *String    `xmlrpc:"street,omitempty"`
	Street2                  *String    `xmlrpc:"street2,omitempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId                   *Many2One  `xmlrpc:"team_id,omitempty"`
	Title                    *Many2One  `xmlrpc:"title,omitempty"`
	Type                     *Selection `xmlrpc:"type,omitempty"`
	UserEmail                *String    `xmlrpc:"user_email,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	UserLogin                *String    `xmlrpc:"user_login,omitempty"`
	Website                  *String    `xmlrpc:"website,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
	Zip                      *String    `xmlrpc:"zip,omitempty"`
}

// CrmLeads represents array of crm.lead model.
type CrmLeads []CrmLead

// CrmLeadModel is the odoo model name.
const CrmLeadModel = "crm.lead"

// Many2One convert CrmLead to *Many2One.
func (cl *CrmLead) Many2One() *Many2One {
	return NewMany2One(cl.Id.Get(), "")
}

// CreateCrmLead creates a new crm.lead model and returns its id.
func (c *Client) CreateCrmLead(cl *CrmLead) (int64, error) {
	return c.Create(CrmLeadModel, cl)
}

// UpdateCrmLead updates an existing crm.lead record.
func (c *Client) UpdateCrmLead(cl *CrmLead) error {
	return c.UpdateCrmLeads([]int64{cl.Id.Get()}, cl)
}

// UpdateCrmLeads updates existing crm.lead records.
// All records (represented by ids) will be updated by cl values.
func (c *Client) UpdateCrmLeads(ids []int64, cl *CrmLead) error {
	return c.Update(CrmLeadModel, ids, cl)
}

// DeleteCrmLead deletes an existing crm.lead record.
func (c *Client) DeleteCrmLead(id int64) error {
	return c.DeleteCrmLeads([]int64{id})
}

// DeleteCrmLeads deletes existing crm.lead records.
func (c *Client) DeleteCrmLeads(ids []int64) error {
	return c.Delete(CrmLeadModel, ids)
}

// GetCrmLead gets crm.lead existing record.
func (c *Client) GetCrmLead(id int64) (*CrmLead, error) {
	cls, err := c.GetCrmLeads([]int64{id})
	if err != nil {
		return nil, err
	}
	if cls != nil && len(*cls) > 0 {
		return &((*cls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead not found", id)
}

// GetCrmLeads gets crm.lead existing records.
func (c *Client) GetCrmLeads(ids []int64) (*CrmLeads, error) {
	cls := &CrmLeads{}
	if err := c.Read(CrmLeadModel, ids, nil, cls); err != nil {
		return nil, err
	}
	return cls, nil
}

// FindCrmLead finds crm.lead record by querying it with criteria.
func (c *Client) FindCrmLead(criteria *Criteria) (*CrmLead, error) {
	cls := &CrmLeads{}
	if err := c.SearchRead(CrmLeadModel, criteria, NewOptions().Limit(1), cls); err != nil {
		return nil, err
	}
	if cls != nil && len(*cls) > 0 {
		return &((*cls)[0]), nil
	}
	return nil, fmt.Errorf("no crm.lead was found with criteria %v", criteria)
}

// FindCrmLeads finds crm.lead records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeads(criteria *Criteria, options *Options) (*CrmLeads, error) {
	cls := &CrmLeads{}
	if err := c.SearchRead(CrmLeadModel, criteria, options, cls); err != nil {
		return nil, err
	}
	return cls, nil
}

// FindCrmLeadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no crm.lead was found with criteria %v and options %v", criteria, options)
}
