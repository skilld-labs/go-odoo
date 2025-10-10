package odoo

// MailingMailing represents mailing.mailing model.
type MailingMailing struct {
	AbTestingCompleted          *Bool      `xmlrpc:"ab_testing_completed,omitempty"`
	AbTestingDescription        *String    `xmlrpc:"ab_testing_description,omitempty"`
	AbTestingEnabled            *Bool      `xmlrpc:"ab_testing_enabled,omitempty"`
	AbTestingIsWinnerMailing    *Bool      `xmlrpc:"ab_testing_is_winner_mailing,omitempty"`
	AbTestingMailingsCount      *Int       `xmlrpc:"ab_testing_mailings_count,omitempty"`
	AbTestingPc                 *Int       `xmlrpc:"ab_testing_pc,omitempty"`
	AbTestingScheduleDatetime   *Time      `xmlrpc:"ab_testing_schedule_datetime,omitempty"`
	AbTestingWinnerSelection    *Selection `xmlrpc:"ab_testing_winner_selection,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omitempty"`
	BodyArch                    *String    `xmlrpc:"body_arch,omitempty"`
	BodyHtml                    *String    `xmlrpc:"body_html,omitempty"`
	Bounced                     *Int       `xmlrpc:"bounced,omitempty"`
	BouncedRatio                *Float     `xmlrpc:"bounced_ratio,omitempty"`
	CalendarDate                *Time      `xmlrpc:"calendar_date,omitempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omitempty"`
	Canceled                    *Int       `xmlrpc:"canceled,omitempty"`
	Clicked                     *Int       `xmlrpc:"clicked,omitempty"`
	ClicksRatio                 *Float     `xmlrpc:"clicks_ratio,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	ContactListIds              *Relation  `xmlrpc:"contact_list_ids,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmLeadCount                *Int       `xmlrpc:"crm_lead_count,omitempty"`
	Delivered                   *Int       `xmlrpc:"delivered,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omitempty"`
	Expected                    *Int       `xmlrpc:"expected,omitempty"`
	Failed                      *Int       `xmlrpc:"failed,omitempty"`
	Favorite                    *Bool      `xmlrpc:"favorite,omitempty"`
	FavoriteDate                *Time      `xmlrpc:"favorite_date,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsAbTestSent                *Bool      `xmlrpc:"is_ab_test_sent,omitempty"`
	IsBodyEmpty                 *Bool      `xmlrpc:"is_body_empty,omitempty"`
	KeepArchives                *Bool      `xmlrpc:"keep_archives,omitempty"`
	KpiMailRequired             *Bool      `xmlrpc:"kpi_mail_required,omitempty"`
	Lang                        *String    `xmlrpc:"lang,omitempty"`
	LinkTrackersCount           *Int       `xmlrpc:"link_trackers_count,omitempty"`
	MailServerAvailable         *Bool      `xmlrpc:"mail_server_available,omitempty"`
	MailServerId                *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MailingDomain               *String    `xmlrpc:"mailing_domain,omitempty"`
	MailingFilterCount          *Int       `xmlrpc:"mailing_filter_count,omitempty"`
	MailingFilterDomain         *String    `xmlrpc:"mailing_filter_domain,omitempty"`
	MailingFilterId             *Many2One  `xmlrpc:"mailing_filter_id,omitempty"`
	MailingModelId              *Many2One  `xmlrpc:"mailing_model_id,omitempty"`
	MailingModelName            *String    `xmlrpc:"mailing_model_name,omitempty"`
	MailingModelReal            *String    `xmlrpc:"mailing_model_real,omitempty"`
	MailingOnMailingList        *Bool      `xmlrpc:"mailing_on_mailing_list,omitempty"`
	MailingTraceIds             *Relation  `xmlrpc:"mailing_trace_ids,omitempty"`
	MailingType                 *Selection `xmlrpc:"mailing_type,omitempty"`
	MailingTypeDescription      *String    `xmlrpc:"mailing_type_description,omitempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NextDeparture               *Time      `xmlrpc:"next_departure,omitempty"`
	NextDepartureIsPast         *Bool      `xmlrpc:"next_departure_is_past,omitempty"`
	Opened                      *Int       `xmlrpc:"opened,omitempty"`
	OpenedRatio                 *Float     `xmlrpc:"opened_ratio,omitempty"`
	Pending                     *Int       `xmlrpc:"pending,omitempty"`
	Preview                     *String    `xmlrpc:"preview,omitempty"`
	Process                     *Int       `xmlrpc:"process,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReceivedRatio               *Float     `xmlrpc:"received_ratio,omitempty"`
	RenderModel                 *String    `xmlrpc:"render_model,omitempty"`
	Replied                     *Int       `xmlrpc:"replied,omitempty"`
	RepliedRatio                *Float     `xmlrpc:"replied_ratio,omitempty"`
	ReplyTo                     *String    `xmlrpc:"reply_to,omitempty"`
	ReplyToMode                 *Selection `xmlrpc:"reply_to_mode,omitempty"`
	SaleInvoicedAmount          *Int       `xmlrpc:"sale_invoiced_amount,omitempty"`
	SaleQuotationCount          *Int       `xmlrpc:"sale_quotation_count,omitempty"`
	ScheduleDate                *Time      `xmlrpc:"schedule_date,omitempty"`
	ScheduleType                *Selection `xmlrpc:"schedule_type,omitempty"`
	Scheduled                   *Int       `xmlrpc:"scheduled,omitempty"`
	Sent                        *Int       `xmlrpc:"sent,omitempty"`
	SentDate                    *Time      `xmlrpc:"sent_date,omitempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	Subject                     *String    `xmlrpc:"subject,omitempty"`
	Total                       *Int       `xmlrpc:"total,omitempty"`
	UseLeads                    *Bool      `xmlrpc:"use_leads,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WarningMessage              *String    `xmlrpc:"warning_message,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailingMailings represents array of mailing.mailing model.
type MailingMailings []MailingMailing

// MailingMailingModel is the odoo model name.
const MailingMailingModel = "mailing.mailing"

// Many2One convert MailingMailing to *Many2One.
func (mm *MailingMailing) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailingMailing creates a new mailing.mailing model and returns its id.
func (c *Client) CreateMailingMailing(mm *MailingMailing) (int64, error) {
	ids, err := c.CreateMailingMailings([]*MailingMailing{mm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingMailing creates a new mailing.mailing model and returns its id.
func (c *Client) CreateMailingMailings(mms []*MailingMailing) ([]int64, error) {
	var vv []interface{}
	for _, v := range mms {
		vv = append(vv, v)
	}
	return c.Create(MailingMailingModel, vv, nil)
}

// UpdateMailingMailing updates an existing mailing.mailing record.
func (c *Client) UpdateMailingMailing(mm *MailingMailing) error {
	return c.UpdateMailingMailings([]int64{mm.Id.Get()}, mm)
}

// UpdateMailingMailings updates existing mailing.mailing records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailingMailings(ids []int64, mm *MailingMailing) error {
	return c.Update(MailingMailingModel, ids, mm, nil)
}

// DeleteMailingMailing deletes an existing mailing.mailing record.
func (c *Client) DeleteMailingMailing(id int64) error {
	return c.DeleteMailingMailings([]int64{id})
}

// DeleteMailingMailings deletes existing mailing.mailing records.
func (c *Client) DeleteMailingMailings(ids []int64) error {
	return c.Delete(MailingMailingModel, ids)
}

// GetMailingMailing gets mailing.mailing existing record.
func (c *Client) GetMailingMailing(id int64) (*MailingMailing, error) {
	mms, err := c.GetMailingMailings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// GetMailingMailings gets mailing.mailing existing records.
func (c *Client) GetMailingMailings(ids []int64) (*MailingMailings, error) {
	mms := &MailingMailings{}
	if err := c.Read(MailingMailingModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailingMailing finds mailing.mailing record by querying it with criteria.
func (c *Client) FindMailingMailing(criteria *Criteria) (*MailingMailing, error) {
	mms := &MailingMailings{}
	if err := c.SearchRead(MailingMailingModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// FindMailingMailings finds mailing.mailing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailings(criteria *Criteria, options *Options) (*MailingMailings, error) {
	mms := &MailingMailings{}
	if err := c.SearchRead(MailingMailingModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailingMailingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingMailingModel, criteria, options)
}

// FindMailingMailingId finds record id by querying it with criteria.
func (c *Client) FindMailingMailingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingMailingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
