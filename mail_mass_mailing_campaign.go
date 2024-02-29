package odoo

// MailMassMailingCampaign represents mail.mass_mailing.campaign model.
type MailMassMailingCampaign struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	Bounced         *Int      `xmlrpc:"bounced,omitempty"`
	BouncedRatio    *Int      `xmlrpc:"bounced_ratio,omitempty"`
	CampaignId      *Many2One `xmlrpc:"campaign_id,omitempty"`
	ClicksRatio     *Int      `xmlrpc:"clicks_ratio,omitempty"`
	Color           *Int      `xmlrpc:"color,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	Delivered       *Int      `xmlrpc:"delivered,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Failed          *Int      `xmlrpc:"failed,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	MassMailingIds  *Relation `xmlrpc:"mass_mailing_ids,omitempty"`
	MediumId        *Many2One `xmlrpc:"medium_id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	Opened          *Int      `xmlrpc:"opened,omitempty"`
	OpenedRatio     *Int      `xmlrpc:"opened_ratio,omitempty"`
	ReceivedRatio   *Int      `xmlrpc:"received_ratio,omitempty"`
	Replied         *Int      `xmlrpc:"replied,omitempty"`
	RepliedRatio    *Int      `xmlrpc:"replied_ratio,omitempty"`
	Scheduled       *Int      `xmlrpc:"scheduled,omitempty"`
	Sent            *Int      `xmlrpc:"sent,omitempty"`
	SourceId        *Many2One `xmlrpc:"source_id,omitempty"`
	StageId         *Many2One `xmlrpc:"stage_id,omitempty"`
	TagIds          *Relation `xmlrpc:"tag_ids,omitempty"`
	Total           *Int      `xmlrpc:"total,omitempty"`
	TotalMailings   *Int      `xmlrpc:"total_mailings,omitempty"`
	UniqueAbTesting *Bool     `xmlrpc:"unique_ab_testing,omitempty"`
	UserId          *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
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

// CreateMailMassMailingCampaigns creates a new mail.mass_mailing.campaign model and returns its id.
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
