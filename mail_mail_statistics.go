package odoo

// MailMailStatistics represents mail.mail.statistics model.
type MailMailStatistics struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Bounced               *Time      `xmlrpc:"bounced,omptempty"`
	Clicked               *Time      `xmlrpc:"clicked,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Exception             *Time      `xmlrpc:"exception,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	LinksClickIds         *Relation  `xmlrpc:"links_click_ids,omptempty"`
	MailMailId            *Many2One  `xmlrpc:"mail_mail_id,omptempty"`
	MailMailIdInt         *Int       `xmlrpc:"mail_mail_id_int,omptempty"`
	MassMailingCampaignId *Many2One  `xmlrpc:"mass_mailing_campaign_id,omptempty"`
	MassMailingId         *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MessageId             *String    `xmlrpc:"message_id,omptempty"`
	Model                 *String    `xmlrpc:"model,omptempty"`
	Opened                *Time      `xmlrpc:"opened,omptempty"`
	Recipient             *String    `xmlrpc:"recipient,omptempty"`
	Replied               *Time      `xmlrpc:"replied,omptempty"`
	ResId                 *Int       `xmlrpc:"res_id,omptempty"`
	Scheduled             *Time      `xmlrpc:"scheduled,omptempty"`
	Sent                  *Time      `xmlrpc:"sent,omptempty"`
	State                 *Selection `xmlrpc:"state,omptempty"`
	StateUpdate           *Time      `xmlrpc:"state_update,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.CreateMailMailStatisticss([]*MailMailStatistics{mms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMailStatistics creates a new mail.mail.statistics model and returns its id.
func (c *Client) CreateMailMailStatisticss(mmss []*MailMailStatistics) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmss {
		vv = append(vv, v)
	}
	return c.Create(MailMailStatisticsModel, vv, nil)
}

// UpdateMailMailStatistics updates an existing mail.mail.statistics record.
func (c *Client) UpdateMailMailStatistics(mms *MailMailStatistics) error {
	return c.UpdateMailMailStatisticss([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMailStatisticss updates existing mail.mail.statistics records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMailStatisticss(ids []int64, mms *MailMailStatistics) error {
	return c.Update(MailMailStatisticsModel, ids, mms, nil)
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
	return &((*mmss)[0]), nil
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
	return &((*mmss)[0]), nil
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
	return c.Search(MailMailStatisticsModel, criteria, options)
}

// FindMailMailStatisticsId finds record id by querying it with criteria.
func (c *Client) FindMailMailStatisticsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailStatisticsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
