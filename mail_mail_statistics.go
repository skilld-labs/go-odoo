package odoo

import (
	"fmt"
)

// MailMailStatistics represents mail.mail.statistics model.
type MailMailStatistics struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	Bounced               *Time      `xmlrpc:"bounced,omitempty"`
	Clicked               *Time      `xmlrpc:"clicked,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	Exception             *Time      `xmlrpc:"exception,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	LinksClickIds         *Relation  `xmlrpc:"links_click_ids,omitempty"`
	MailMailId            *Many2One  `xmlrpc:"mail_mail_id,omitempty"`
	MailMailIdInt         *Int       `xmlrpc:"mail_mail_id_int,omitempty"`
	MassMailingCampaignId *Many2One  `xmlrpc:"mass_mailing_campaign_id,omitempty"`
	MassMailingId         *Many2One  `xmlrpc:"mass_mailing_id,omitempty"`
	MessageId             *String    `xmlrpc:"message_id,omitempty"`
	Model                 *String    `xmlrpc:"model,omitempty"`
	Opened                *Time      `xmlrpc:"opened,omitempty"`
	Recipient             *String    `xmlrpc:"recipient,omitempty"`
	Replied               *Time      `xmlrpc:"replied,omitempty"`
	ResId                 *Int       `xmlrpc:"res_id,omitempty"`
	Scheduled             *Time      `xmlrpc:"scheduled,omitempty"`
	Sent                  *Time      `xmlrpc:"sent,omitempty"`
	State                 *Selection `xmlrpc:"state,omitempty"`
	StateUpdate           *Time      `xmlrpc:"state_update,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailMailStatisticss represents array of mail.mail.statistics model.
type MailMailStatisticss []MailMailStatistics

// MailMailStatisticsModel is the odoo model name.
const MailMailStatisticsModel = "mail.mail.statistics"

// Many2One convert MailMailStatistics to *Many2One.
func (mms *MailMailStatistics) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMailStatistics creates a new mail.mail.statistics model and returns its id.
func (c *Client) CreateMailMailStatistics(mms *MailMailStatistics) (int64, error) {
	return c.Create(MailMailStatisticsModel, mms)
}

// UpdateMailMailStatistics updates an existing mail.mail.statistics record.
func (c *Client) UpdateMailMailStatistics(mms *MailMailStatistics) error {
	return c.UpdateMailMailStatisticss([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMailStatisticss updates existing mail.mail.statistics records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMailStatisticss(ids []int64, mms *MailMailStatistics) error {
	return c.Update(MailMailStatisticsModel, ids, mms)
}

// DeleteMailMailStatistics deletes an existing mail.mail.statistics record.
func (c *Client) DeleteMailMailStatistics(id int64) error {
	return c.DeleteMailMailStatisticss([]int64{id})
}

// DeleteMailMailStatisticss deletes existing mail.mail.statistics records.
func (c *Client) DeleteMailMailStatisticss(ids []int64) error {
	return c.Delete(MailMailStatisticsModel, ids)
}

// GetMailMailStatistics gets mail.mail.statistics existing record.
func (c *Client) GetMailMailStatistics(id int64) (*MailMailStatistics, error) {
	mmss, err := c.GetMailMailStatisticss([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mail.statistics not found", id)
}

// GetMailMailStatisticss gets mail.mail.statistics existing records.
func (c *Client) GetMailMailStatisticss(ids []int64) (*MailMailStatisticss, error) {
	mmss := &MailMailStatisticss{}
	if err := c.Read(MailMailStatisticsModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMailStatistics finds mail.mail.statistics record by querying it with criteria.
func (c *Client) FindMailMailStatistics(criteria *Criteria) (*MailMailStatistics, error) {
	mmss := &MailMailStatisticss{}
	if err := c.SearchRead(MailMailStatisticsModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("no mail.mail.statistics was found with criteria %v", criteria)
}

// FindMailMailStatisticss finds mail.mail.statistics records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailStatisticss(criteria *Criteria, options *Options) (*MailMailStatisticss, error) {
	mmss := &MailMailStatisticss{}
	if err := c.SearchRead(MailMailStatisticsModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMailStatisticsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailStatisticsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMailStatisticsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMailStatisticsId finds record id by querying it with criteria.
func (c *Client) FindMailMailStatisticsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailStatisticsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no mail.mail.statistics was found with criteria %v and options %v", criteria, options)
}
