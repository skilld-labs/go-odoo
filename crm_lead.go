package odoo

import (
	"fmt"
)

// CrmLead represents crm.lead model.
type CrmLead struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline     *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState            *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary          *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omptempty"`
	City                     *String    `xmlrpc:"city,omptempty"`
	Color                    *Int       `xmlrpc:"color,omptempty"`
	CompanyCurrency          *Many2One  `xmlrpc:"company_currency,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	ContactName              *String    `xmlrpc:"contact_name,omptempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateActionLast           *Time      `xmlrpc:"date_action_last,omptempty"`
	DateClosed               *Time      `xmlrpc:"date_closed,omptempty"`
	DateConversion           *Time      `xmlrpc:"date_conversion,omptempty"`
	DateDeadline             *Time      `xmlrpc:"date_deadline,omptempty"`
	DateLastStageUpdate      *Time      `xmlrpc:"date_last_stage_update,omptempty"`
	DateOpen                 *Time      `xmlrpc:"date_open,omptempty"`
	DayClose                 *Float     `xmlrpc:"day_close,omptempty"`
	DayOpen                  *Float     `xmlrpc:"day_open,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	EmailCc                  *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom                *String    `xmlrpc:"email_from,omptempty"`
	Function                 *String    `xmlrpc:"function,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	KanbanState              *Selection `xmlrpc:"kanban_state,omptempty"`
	LostReason               *Many2One  `xmlrpc:"lost_reason,omptempty"`
	MachineLeadName          *String    `xmlrpc:"machine_lead_name,omptempty"`
	MediumId                 *Many2One  `xmlrpc:"medium_id,omptempty"`
	MeetingCount             *Int       `xmlrpc:"meeting_count,omptempty"`
	MessageBounce            *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mobile                   *String    `xmlrpc:"mobile,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OptOut                   *Bool      `xmlrpc:"opt_out,omptempty"`
	OrderIds                 *Relation  `xmlrpc:"order_ids,omptempty"`
	PartnerAddressEmail      *String    `xmlrpc:"partner_address_email,omptempty"`
	PartnerAddressName       *String    `xmlrpc:"partner_address_name,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerName              *String    `xmlrpc:"partner_name,omptempty"`
	Phone                    *String    `xmlrpc:"phone,omptempty"`
	PlannedRevenue           *Float     `xmlrpc:"planned_revenue,omptempty"`
	Priority                 *Selection `xmlrpc:"priority,omptempty"`
	Probability              *Float     `xmlrpc:"probability,omptempty"`
	Referred                 *String    `xmlrpc:"referred,omptempty"`
	SaleAmountTotal          *Float     `xmlrpc:"sale_amount_total,omptempty"`
	SaleNumber               *Int       `xmlrpc:"sale_number,omptempty"`
	SourceId                 *Many2One  `xmlrpc:"source_id,omptempty"`
	StageId                  *Many2One  `xmlrpc:"stage_id,omptempty"`
	StateId                  *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                   *String    `xmlrpc:"street,omptempty"`
	Street2                  *String    `xmlrpc:"street2,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	TeamId                   *Many2One  `xmlrpc:"team_id,omptempty"`
	Title                    *Many2One  `xmlrpc:"title,omptempty"`
	Type                     *Selection `xmlrpc:"type,omptempty"`
	UserEmail                *String    `xmlrpc:"user_email,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	UserLogin                *String    `xmlrpc:"user_login,omptempty"`
	Website                  *String    `xmlrpc:"website,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                      *String    `xmlrpc:"zip,omptempty"`
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
	ids, err := c.CreateCrmLeads([]*CrmLead{cl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLead creates a new crm.lead model and returns its id.
func (c *Client) CreateCrmLeads(cls []*CrmLead) ([]int64, error) {
	var vv []interface{}
	for _, v := range cls {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadModel, vv)
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
	return nil, fmt.Errorf("crm.lead was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("crm.lead was not found with criteria %v and options %v", criteria, options)
}
