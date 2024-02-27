package odoo

// MailMassMailing represents mail.mass_mailing model.
type MailMassMailing struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	Active                *Bool      `xmlrpc:"active,omitempty"`
	AttachmentIds         *Relation  `xmlrpc:"attachment_ids,omitempty"`
	BodyHtml              *String    `xmlrpc:"body_html,omitempty"`
	Bounced               *Int       `xmlrpc:"bounced,omitempty"`
	BouncedRatio          *Int       `xmlrpc:"bounced_ratio,omitempty"`
	CampaignId            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	ClicksRatio           *Int       `xmlrpc:"clicks_ratio,omitempty"`
	Color                 *Int       `xmlrpc:"color,omitempty"`
	ContactAbPc           *Int       `xmlrpc:"contact_ab_pc,omitempty"`
	ContactListIds        *Relation  `xmlrpc:"contact_list_ids,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delivered             *Int       `xmlrpc:"delivered,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom             *String    `xmlrpc:"email_from,omitempty"`
	Failed                *Int       `xmlrpc:"failed,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	KeepArchives          *Bool      `xmlrpc:"keep_archives,omitempty"`
	MailingDomain         *String    `xmlrpc:"mailing_domain,omitempty"`
	MailingModelId        *Many2One  `xmlrpc:"mailing_model_id,omitempty"`
	MailingModelName      *String    `xmlrpc:"mailing_model_name,omitempty"`
	MailingModelReal      *String    `xmlrpc:"mailing_model_real,omitempty"`
	MassMailingCampaignId *Many2One  `xmlrpc:"mass_mailing_campaign_id,omitempty"`
	MediumId              *Many2One  `xmlrpc:"medium_id,omitempty"`
	Name                  *String    `xmlrpc:"name,omitempty"`
	NextDeparture         *Time      `xmlrpc:"next_departure,omitempty"`
	Opened                *Int       `xmlrpc:"opened,omitempty"`
	OpenedRatio           *Int       `xmlrpc:"opened_ratio,omitempty"`
	ReceivedRatio         *Int       `xmlrpc:"received_ratio,omitempty"`
	Replied               *Int       `xmlrpc:"replied,omitempty"`
	RepliedRatio          *Int       `xmlrpc:"replied_ratio,omitempty"`
	ReplyTo               *String    `xmlrpc:"reply_to,omitempty"`
	ReplyToMode           *Selection `xmlrpc:"reply_to_mode,omitempty"`
	ScheduleDate          *Time      `xmlrpc:"schedule_date,omitempty"`
	Scheduled             *Int       `xmlrpc:"scheduled,omitempty"`
	Sent                  *Int       `xmlrpc:"sent,omitempty"`
	SentDate              *Time      `xmlrpc:"sent_date,omitempty"`
	SourceId              *Many2One  `xmlrpc:"source_id,omitempty"`
	State                 *Selection `xmlrpc:"state,omitempty"`
	StatisticsIds         *Relation  `xmlrpc:"statistics_ids,omitempty"`
	Total                 *Int       `xmlrpc:"total,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailMassMailings represents array of mail.mass_mailing model.
type MailMassMailings []MailMassMailing

// MailMassMailingModel is the odoo model name.
const MailMassMailingModel = "mail.mass_mailing"

// Many2One convert MailMassMailing to *Many2One.
func (mm *MailMassMailing) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMassMailing creates a new mail.mass_mailing model and returns its id.
func (c *Client) CreateMailMassMailing(mm *MailMassMailing) (int64, error) {
	ids, err := c.CreateMailMassMailings([]*MailMassMailing{mm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailing creates a new mail.mass_mailing model and returns its id.
func (c *Client) CreateMailMassMailings(mms []*MailMassMailing) ([]int64, error) {
	var vv []interface{}
	for _, v := range mms {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingModel, vv, nil)
}

// UpdateMailMassMailing updates an existing mail.mass_mailing record.
func (c *Client) UpdateMailMassMailing(mm *MailMassMailing) error {
	return c.UpdateMailMassMailings([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMassMailings updates existing mail.mass_mailing records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMassMailings(ids []int64, mm *MailMassMailing) error {
	return c.Update(MailMassMailingModel, ids, mm, nil)
}

// DeleteMailMassMailing deletes an existing mail.mass_mailing record.
func (c *Client) DeleteMailMassMailing(id int64) error {
	return c.DeleteMailMassMailings([]int64{id})
}

// DeleteMailMassMailings deletes existing mail.mass_mailing records.
func (c *Client) DeleteMailMassMailings(ids []int64) error {
	return c.Delete(MailMassMailingModel, ids)
}

// GetMailMassMailing gets mail.mass_mailing existing record.
func (c *Client) GetMailMassMailing(id int64) (*MailMassMailing, error) {
	mms, err := c.GetMailMassMailings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// GetMailMassMailings gets mail.mass_mailing existing records.
func (c *Client) GetMailMassMailings(ids []int64) (*MailMassMailings, error) {
	mms := &MailMassMailings{}
	if err := c.Read(MailMassMailingModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMassMailing finds mail.mass_mailing record by querying it with criteria.
func (c *Client) FindMailMassMailing(criteria *Criteria) (*MailMassMailing, error) {
	mms := &MailMassMailings{}
	if err := c.SearchRead(MailMassMailingModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// FindMailMassMailings finds mail.mass_mailing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailings(criteria *Criteria, options *Options) (*MailMassMailings, error) {
	mms := &MailMassMailings{}
	if err := c.SearchRead(MailMassMailingModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMassMailingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMassMailingModel, criteria, options)
}

// FindMailMassMailingId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
