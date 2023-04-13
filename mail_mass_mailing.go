package odoo

import (
	"fmt"
)

// MailMassMailing represents mail.mass_mailing model.
type MailMassMailing struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	AttachmentIds         *Relation  `xmlrpc:"attachment_ids,omptempty"`
	BodyHtml              *String    `xmlrpc:"body_html,omptempty"`
	Bounced               *Int       `xmlrpc:"bounced,omptempty"`
	BouncedRatio          *Int       `xmlrpc:"bounced_ratio,omptempty"`
	CampaignId            *Many2One  `xmlrpc:"campaign_id,omptempty"`
	ClicksRatio           *Int       `xmlrpc:"clicks_ratio,omptempty"`
	Color                 *Int       `xmlrpc:"color,omptempty"`
	ContactAbPc           *Int       `xmlrpc:"contact_ab_pc,omptempty"`
	ContactListIds        *Relation  `xmlrpc:"contact_list_ids,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	Delivered             *Int       `xmlrpc:"delivered,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom             *String    `xmlrpc:"email_from,omptempty"`
	Failed                *Int       `xmlrpc:"failed,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	KeepArchives          *Bool      `xmlrpc:"keep_archives,omptempty"`
	MailingDomain         *String    `xmlrpc:"mailing_domain,omptempty"`
	MailingModelId        *Many2One  `xmlrpc:"mailing_model_id,omptempty"`
	MailingModelName      *String    `xmlrpc:"mailing_model_name,omptempty"`
	MailingModelReal      *String    `xmlrpc:"mailing_model_real,omptempty"`
	MassMailingCampaignId *Many2One  `xmlrpc:"mass_mailing_campaign_id,omptempty"`
	MediumId              *Many2One  `xmlrpc:"medium_id,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	NextDeparture         *Time      `xmlrpc:"next_departure,omptempty"`
	Opened                *Int       `xmlrpc:"opened,omptempty"`
	OpenedRatio           *Int       `xmlrpc:"opened_ratio,omptempty"`
	ReceivedRatio         *Int       `xmlrpc:"received_ratio,omptempty"`
	Replied               *Int       `xmlrpc:"replied,omptempty"`
	RepliedRatio          *Int       `xmlrpc:"replied_ratio,omptempty"`
	ReplyTo               *String    `xmlrpc:"reply_to,omptempty"`
	ReplyToMode           *Selection `xmlrpc:"reply_to_mode,omptempty"`
	ScheduleDate          *Time      `xmlrpc:"schedule_date,omptempty"`
	Scheduled             *Int       `xmlrpc:"scheduled,omptempty"`
	Sent                  *Int       `xmlrpc:"sent,omptempty"`
	SentDate              *Time      `xmlrpc:"sent_date,omptempty"`
	SourceId              *Many2One  `xmlrpc:"source_id,omptempty"`
	State                 *Selection `xmlrpc:"state,omptempty"`
	StatisticsIds         *Relation  `xmlrpc:"statistics_ids,omptempty"`
	Total                 *Int       `xmlrpc:"total,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(MailMassMailingModel, vv)
}

// UpdateMailMassMailing updates an existing mail.mass_mailing record.
func (c *Client) UpdateMailMassMailing(mm *MailMassMailing) error {
	return c.UpdateMailMassMailings([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMassMailings updates existing mail.mass_mailing records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMassMailings(ids []int64, mm *MailMassMailing) error {
	return c.Update(MailMassMailingModel, ids, mm)
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
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mass_mailing not found", id)
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
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mail.mass_mailing was not found with criteria %v", criteria)
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
	ids, err := c.Search(MailMassMailingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMassMailingId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.mass_mailing was not found with criteria %v and options %v", criteria, options)
}
