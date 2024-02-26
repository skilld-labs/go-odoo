package odoo

// MailMassMailingCampaign represents mail.mass_mailing.campaign model.
type MailMassMailingCampaign struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Bounced         *Int      `xmlrpc:"bounced,omptempty"`
	BouncedRatio    *Int      `xmlrpc:"bounced_ratio,omptempty"`
	CampaignId      *Many2One `xmlrpc:"campaign_id,omptempty"`
	ClicksRatio     *Int      `xmlrpc:"clicks_ratio,omptempty"`
	Color           *Int      `xmlrpc:"color,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	Delivered       *Int      `xmlrpc:"delivered,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Failed          *Int      `xmlrpc:"failed,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	MassMailingIds  *Relation `xmlrpc:"mass_mailing_ids,omptempty"`
	MediumId        *Many2One `xmlrpc:"medium_id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Opened          *Int      `xmlrpc:"opened,omptempty"`
	OpenedRatio     *Int      `xmlrpc:"opened_ratio,omptempty"`
	ReceivedRatio   *Int      `xmlrpc:"received_ratio,omptempty"`
	Replied         *Int      `xmlrpc:"replied,omptempty"`
	RepliedRatio    *Int      `xmlrpc:"replied_ratio,omptempty"`
	Scheduled       *Int      `xmlrpc:"scheduled,omptempty"`
	Sent            *Int      `xmlrpc:"sent,omptempty"`
	SourceId        *Many2One `xmlrpc:"source_id,omptempty"`
	StageId         *Many2One `xmlrpc:"stage_id,omptempty"`
	TagIds          *Relation `xmlrpc:"tag_ids,omptempty"`
	Total           *Int      `xmlrpc:"total,omptempty"`
	TotalMailings   *Int      `xmlrpc:"total_mailings,omptempty"`
	UniqueAbTesting *Bool     `xmlrpc:"unique_ab_testing,omptempty"`
	UserId          *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMassMailingCampaigns represents array of mail.mass_mailing.campaign model.
type MailMassMailingCampaigns []MailMassMailingCampaign

// MailMassMailingCampaignModel is the odoo model name.
const MailMassMailingCampaignModel = "mail.mass_mailing.campaign"

// Many2One convert MailMassMailingCampaign to *Many2One.
func (mmc *MailMassMailingCampaign) Many2One() *Many2One {
	return NewMany2One(mmc.Id.Get(), "")
}

// CreateMailMassMailingCampaign creates a new mail.mass_mailing.campaign model and returns its id.
func (c *Client) CreateMailMassMailingCampaign(mmc *MailMassMailingCampaign) (int64, error) {
	ids, err := c.CreateMailMassMailingCampaigns([]*MailMassMailingCampaign{mmc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingCampaign creates a new mail.mass_mailing.campaign model and returns its id.
func (c *Client) CreateMailMassMailingCampaigns(mmcs []*MailMassMailingCampaign) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmcs {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingCampaignModel, vv, nil)
}

// UpdateMailMassMailingCampaign updates an existing mail.mass_mailing.campaign record.
func (c *Client) UpdateMailMassMailingCampaign(mmc *MailMassMailingCampaign) error {
	return c.UpdateMailMassMailingCampaigns([]int64{mmc.Id.Get()}, mmc)
}

// UpdateMailMassMailingCampaigns updates existing mail.mass_mailing.campaign records.
// All records (represented by ids) will be updated by mmc values.
func (c *Client) UpdateMailMassMailingCampaigns(ids []int64, mmc *MailMassMailingCampaign) error {
	return c.Update(MailMassMailingCampaignModel, ids, mmc, nil)
}

// DeleteMailMassMailingCampaign deletes an existing mail.mass_mailing.campaign record.
func (c *Client) DeleteMailMassMailingCampaign(id int64) error {
	return c.DeleteMailMassMailingCampaigns([]int64{id})
}

// DeleteMailMassMailingCampaigns deletes existing mail.mass_mailing.campaign records.
func (c *Client) DeleteMailMassMailingCampaigns(ids []int64) error {
	return c.Delete(MailMassMailingCampaignModel, ids)
}

// GetMailMassMailingCampaign gets mail.mass_mailing.campaign existing record.
func (c *Client) GetMailMassMailingCampaign(id int64) (*MailMassMailingCampaign, error) {
	mmcs, err := c.GetMailMassMailingCampaigns([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmcs)[0]), nil
}

// GetMailMassMailingCampaigns gets mail.mass_mailing.campaign existing records.
func (c *Client) GetMailMassMailingCampaigns(ids []int64) (*MailMassMailingCampaigns, error) {
	mmcs := &MailMassMailingCampaigns{}
	if err := c.Read(MailMassMailingCampaignModel, ids, nil, mmcs); err != nil {
		return nil, err
	}
	return mmcs, nil
}

// FindMailMassMailingCampaign finds mail.mass_mailing.campaign record by querying it with criteria.
func (c *Client) FindMailMassMailingCampaign(criteria *Criteria) (*MailMassMailingCampaign, error) {
	mmcs := &MailMassMailingCampaigns{}
	if err := c.SearchRead(MailMassMailingCampaignModel, criteria, NewOptions().Limit(1), mmcs); err != nil {
		return nil, err
	}
	return &((*mmcs)[0]), nil
}

// FindMailMassMailingCampaigns finds mail.mass_mailing.campaign records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingCampaigns(criteria *Criteria, options *Options) (*MailMassMailingCampaigns, error) {
	mmcs := &MailMassMailingCampaigns{}
	if err := c.SearchRead(MailMassMailingCampaignModel, criteria, options, mmcs); err != nil {
		return nil, err
	}
	return mmcs, nil
}

// FindMailMassMailingCampaignIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingCampaignIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMassMailingCampaignModel, criteria, options)
}

// FindMailMassMailingCampaignId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingCampaignId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingCampaignModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
