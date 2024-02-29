package odoo

// MailMassMailingContact represents mail.mass_mailing.contact model.
type MailMassMailingContact struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyName              *String   `xmlrpc:"company_name,omitempty"`
	CountryId                *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Email                    *String   `xmlrpc:"email,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	ListIds                  *Relation `xmlrpc:"list_ids,omitempty"`
	MessageBounce            *Int      `xmlrpc:"message_bounce,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	OptOut                   *Bool     `xmlrpc:"opt_out,omitempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omitempty"`
	TitleId                  *Many2One `xmlrpc:"title_id,omitempty"`
	UnsubscriptionDate       *Time     `xmlrpc:"unsubscription_date,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMassMailingContacts represents array of mail.mass_mailing.contact model.
type MailMassMailingContacts []MailMassMailingContact

// MailMassMailingContactModel is the odoo model name.
const MailMassMailingContactModel = "mail.mass_mailing.contact"

// Many2One convert MailMassMailingContact to *Many2One.
func (mmc *MailMassMailingContact) Many2One() *Many2One {
	return NewMany2One(mmc.Id.Get(), "")
}

// CreateMailMassMailingContact creates a new mail.mass_mailing.contact model and returns its id.
func (c *Client) CreateMailMassMailingContact(mmc *MailMassMailingContact) (int64, error) {
	ids, err := c.CreateMailMassMailingContacts([]*MailMassMailingContact{mmc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingContacts creates a new mail.mass_mailing.contact model and returns its id.
func (c *Client) CreateMailMassMailingContacts(mmcs []*MailMassMailingContact) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmcs {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingContactModel, vv, nil)
}

// UpdateMailMassMailingContact updates an existing mail.mass_mailing.contact record.
func (c *Client) UpdateMailMassMailingContact(mmc *MailMassMailingContact) error {
	return c.UpdateMailMassMailingContacts([]int64{mmc.Id.Get()}, mmc)
}

// UpdateMailMassMailingContacts updates existing mail.mass_mailing.contact records.
// All records (represented by ids) will be updated by mmc values.
func (c *Client) UpdateMailMassMailingContacts(ids []int64, mmc *MailMassMailingContact) error {
	return c.Update(MailMassMailingContactModel, ids, mmc, nil)
}

// DeleteMailMassMailingContact deletes an existing mail.mass_mailing.contact record.
func (c *Client) DeleteMailMassMailingContact(id int64) error {
	return c.DeleteMailMassMailingContacts([]int64{id})
}

// DeleteMailMassMailingContacts deletes existing mail.mass_mailing.contact records.
func (c *Client) DeleteMailMassMailingContacts(ids []int64) error {
	return c.Delete(MailMassMailingContactModel, ids)
}

// GetMailMassMailingContact gets mail.mass_mailing.contact existing record.
func (c *Client) GetMailMassMailingContact(id int64) (*MailMassMailingContact, error) {
	mmcs, err := c.GetMailMassMailingContacts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmcs)[0]), nil
}

// GetMailMassMailingContacts gets mail.mass_mailing.contact existing records.
func (c *Client) GetMailMassMailingContacts(ids []int64) (*MailMassMailingContacts, error) {
	mmcs := &MailMassMailingContacts{}
	if err := c.Read(MailMassMailingContactModel, ids, nil, mmcs); err != nil {
		return nil, err
	}
	return mmcs, nil
}

// FindMailMassMailingContact finds mail.mass_mailing.contact record by querying it with criteria.
func (c *Client) FindMailMassMailingContact(criteria *Criteria) (*MailMassMailingContact, error) {
	mmcs := &MailMassMailingContacts{}
	if err := c.SearchRead(MailMassMailingContactModel, criteria, NewOptions().Limit(1), mmcs); err != nil {
		return nil, err
	}
	return &((*mmcs)[0]), nil
}

// FindMailMassMailingContacts finds mail.mass_mailing.contact records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingContacts(criteria *Criteria, options *Options) (*MailMassMailingContacts, error) {
	mmcs := &MailMassMailingContacts{}
	if err := c.SearchRead(MailMassMailingContactModel, criteria, options, mmcs); err != nil {
		return nil, err
	}
	return mmcs, nil
}

// FindMailMassMailingContactIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingContactIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMassMailingContactModel, criteria, options)
}

// FindMailMassMailingContactId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingContactId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingContactModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
