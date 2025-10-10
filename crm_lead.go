package odoo

// CrmLead represents crm.lead model.
type CrmLead struct {
	Active                          *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId         *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline            *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration     *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon           *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                     *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                   *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                 *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                  *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                  *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AutomatedProbability            *Float      `xmlrpc:"automated_probability,omitempty"`
	CalendarEventIds                *Relation   `xmlrpc:"calendar_event_ids,omitempty"`
	CampaignId                      *Many2One   `xmlrpc:"campaign_id,omitempty"`
	City                            *String     `xmlrpc:"city,omitempty"`
	Color                           *Int        `xmlrpc:"color,omitempty"`
	CompanyCurrency                 *Many2One   `xmlrpc:"company_currency,omitempty"`
	CompanyId                       *Many2One   `xmlrpc:"company_id,omitempty"`
	ContactName                     *String     `xmlrpc:"contact_name,omitempty"`
	CountryId                       *Many2One   `xmlrpc:"country_id,omitempty"`
	CreateDate                      *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                       *Many2One   `xmlrpc:"create_uid,omitempty"`
	DateAutomationLast              *Time       `xmlrpc:"date_automation_last,omitempty"`
	DateClosed                      *Time       `xmlrpc:"date_closed,omitempty"`
	DateConversion                  *Time       `xmlrpc:"date_conversion,omitempty"`
	DateDeadline                    *Time       `xmlrpc:"date_deadline,omitempty"`
	DateLastStageUpdate             *Time       `xmlrpc:"date_last_stage_update,omitempty"`
	DateOpen                        *Time       `xmlrpc:"date_open,omitempty"`
	DayClose                        *Float      `xmlrpc:"day_close,omitempty"`
	DayOpen                         *Float      `xmlrpc:"day_open,omitempty"`
	Description                     *String     `xmlrpc:"description,omitempty"`
	DisplayName                     *String     `xmlrpc:"display_name,omitempty"`
	DuplicateLeadCount              *Int        `xmlrpc:"duplicate_lead_count,omitempty"`
	DuplicateLeadIds                *Relation   `xmlrpc:"duplicate_lead_ids,omitempty"`
	DurationTracking                interface{} `xmlrpc:"duration_tracking,omitempty"`
	EmailCc                         *String     `xmlrpc:"email_cc,omitempty"`
	EmailDomainCriterion            *String     `xmlrpc:"email_domain_criterion,omitempty"`
	EmailFrom                       *String     `xmlrpc:"email_from,omitempty"`
	EmailNormalized                 *String     `xmlrpc:"email_normalized,omitempty"`
	EmailState                      *Selection  `xmlrpc:"email_state,omitempty"`
	ExpectedRevenue                 *Float      `xmlrpc:"expected_revenue,omitempty"`
	Function                        *String     `xmlrpc:"function,omitempty"`
	HasMessage                      *Bool       `xmlrpc:"has_message,omitempty"`
	IapEnrichDone                   *Bool       `xmlrpc:"iap_enrich_done,omitempty"`
	Id                              *Int        `xmlrpc:"id,omitempty"`
	IsAutomatedProbability          *Bool       `xmlrpc:"is_automated_probability,omitempty"`
	IsBlacklisted                   *Bool       `xmlrpc:"is_blacklisted,omitempty"`
	IsPartnerVisible                *Bool       `xmlrpc:"is_partner_visible,omitempty"`
	LangActiveCount                 *Int        `xmlrpc:"lang_active_count,omitempty"`
	LangCode                        *String     `xmlrpc:"lang_code,omitempty"`
	LangId                          *Many2One   `xmlrpc:"lang_id,omitempty"`
	LeadMiningRequestId             *Many2One   `xmlrpc:"lead_mining_request_id,omitempty"`
	LeadProperties                  interface{} `xmlrpc:"lead_properties,omitempty"`
	LostReasonId                    *Many2One   `xmlrpc:"lost_reason_id,omitempty"`
	MachineLeadName                 *String     `xmlrpc:"machine_lead_name,omitempty"`
	MediumId                        *Many2One   `xmlrpc:"medium_id,omitempty"`
	MeetingDisplayDate              *Time       `xmlrpc:"meeting_display_date,omitempty"`
	MeetingDisplayLabel             *String     `xmlrpc:"meeting_display_label,omitempty"`
	MessageAttachmentCount          *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce                   *Int        `xmlrpc:"message_bounce,omitempty"`
	MessageFollowerIds              *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                 *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter          *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError              *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                      *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower               *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction               *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter        *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds               *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                          *String     `xmlrpc:"mobile,omitempty"`
	MobileBlacklisted               *Bool       `xmlrpc:"mobile_blacklisted,omitempty"`
	MyActivityDateDeadline          *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                            *String     `xmlrpc:"name,omitempty"`
	OrderIds                        *Relation   `xmlrpc:"order_ids,omitempty"`
	PartnerEmailUpdate              *Bool       `xmlrpc:"partner_email_update,omitempty"`
	PartnerId                       *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerIsBlacklisted            *Bool       `xmlrpc:"partner_is_blacklisted,omitempty"`
	PartnerName                     *String     `xmlrpc:"partner_name,omitempty"`
	PartnerPhoneUpdate              *Bool       `xmlrpc:"partner_phone_update,omitempty"`
	Phone                           *String     `xmlrpc:"phone,omitempty"`
	PhoneBlacklisted                *Bool       `xmlrpc:"phone_blacklisted,omitempty"`
	PhoneMobileSearch               *String     `xmlrpc:"phone_mobile_search,omitempty"`
	PhoneSanitized                  *String     `xmlrpc:"phone_sanitized,omitempty"`
	PhoneSanitizedBlacklisted       *Bool       `xmlrpc:"phone_sanitized_blacklisted,omitempty"`
	PhoneState                      *Selection  `xmlrpc:"phone_state,omitempty"`
	Priority                        *Selection  `xmlrpc:"priority,omitempty"`
	Probability                     *Float      `xmlrpc:"probability,omitempty"`
	ProratedRevenue                 *Float      `xmlrpc:"prorated_revenue,omitempty"`
	QuotationCount                  *Int        `xmlrpc:"quotation_count,omitempty"`
	RatingIds                       *Relation   `xmlrpc:"rating_ids,omitempty"`
	RecurringPlan                   *Many2One   `xmlrpc:"recurring_plan,omitempty"`
	RecurringRevenue                *Float      `xmlrpc:"recurring_revenue,omitempty"`
	RecurringRevenueMonthly         *Float      `xmlrpc:"recurring_revenue_monthly,omitempty"`
	RecurringRevenueMonthlyProrated *Float      `xmlrpc:"recurring_revenue_monthly_prorated,omitempty"`
	RecurringRevenueProrated        *Float      `xmlrpc:"recurring_revenue_prorated,omitempty"`
	Referred                        *String     `xmlrpc:"referred,omitempty"`
	RevealId                        *String     `xmlrpc:"reveal_id,omitempty"`
	SaleAmountTotal                 *Float      `xmlrpc:"sale_amount_total,omitempty"`
	SaleOrderCount                  *Int        `xmlrpc:"sale_order_count,omitempty"`
	ShowEnrichButton                *Bool       `xmlrpc:"show_enrich_button,omitempty"`
	SourceId                        *Many2One   `xmlrpc:"source_id,omitempty"`
	StageId                         *Many2One   `xmlrpc:"stage_id,omitempty"`
	StateId                         *Many2One   `xmlrpc:"state_id,omitempty"`
	Street                          *String     `xmlrpc:"street,omitempty"`
	Street2                         *String     `xmlrpc:"street2,omitempty"`
	TagIds                          *Relation   `xmlrpc:"tag_ids,omitempty"`
	TeamId                          *Many2One   `xmlrpc:"team_id,omitempty"`
	Title                           *Many2One   `xmlrpc:"title,omitempty"`
	Type                            *Selection  `xmlrpc:"type,omitempty"`
	UserCompanyIds                  *Relation   `xmlrpc:"user_company_ids,omitempty"`
	UserId                          *Many2One   `xmlrpc:"user_id,omitempty"`
	Website                         *String     `xmlrpc:"website,omitempty"`
	WebsiteMessageIds               *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                       *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One   `xmlrpc:"write_uid,omitempty"`
	Zip                             *String     `xmlrpc:"zip,omitempty"`
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
	return c.Create(CrmLeadModel, vv, nil)
}

// UpdateCrmLead updates an existing crm.lead record.
func (c *Client) UpdateCrmLead(cl *CrmLead) error {
	return c.UpdateCrmLeads([]int64{cl.Id.Get()}, cl)
}

// UpdateCrmLeads updates existing crm.lead records.
// All records (represented by ids) will be updated by cl values.
func (c *Client) UpdateCrmLeads(ids []int64, cl *CrmLead) error {
	return c.Update(CrmLeadModel, ids, cl, nil)
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
	return &((*cls)[0]), nil
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
	return &((*cls)[0]), nil
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
	return c.Search(CrmLeadModel, criteria, options)
}

// FindCrmLeadId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
