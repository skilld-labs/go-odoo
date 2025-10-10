package odoo

// MailingContact represents mailing.contact model.
type MailingContact struct {
	CompanyName              *String   `xmlrpc:"company_name,omitempty"`
	CountryId                *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Email                    *String   `xmlrpc:"email,omitempty"`
	EmailNormalized          *String   `xmlrpc:"email_normalized,omitempty"`
	FirstName                *String   `xmlrpc:"first_name,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsBlacklisted            *Bool     `xmlrpc:"is_blacklisted,omitempty"`
	LastName                 *String   `xmlrpc:"last_name,omitempty"`
	ListIds                  *Relation `xmlrpc:"list_ids,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce            *Int      `xmlrpc:"message_bounce,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	OptOut                   *Bool     `xmlrpc:"opt_out,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	SubscriptionIds          *Relation `xmlrpc:"subscription_ids,omitempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omitempty"`
	TitleId                  *Many2One `xmlrpc:"title_id,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingContacts represents array of mailing.contact model.
type MailingContacts []MailingContact

// MailingContactModel is the odoo model name.
const MailingContactModel = "mailing.contact"

// Many2One convert MailingContact to *Many2One.
func (mc *MailingContact) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMailingContact creates a new mailing.contact model and returns its id.
func (c *Client) CreateMailingContact(mc *MailingContact) (int64, error) {
	ids, err := c.CreateMailingContacts([]*MailingContact{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingContact creates a new mailing.contact model and returns its id.
func (c *Client) CreateMailingContacts(mcs []*MailingContact) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MailingContactModel, vv, nil)
}

// UpdateMailingContact updates an existing mailing.contact record.
func (c *Client) UpdateMailingContact(mc *MailingContact) error {
	return c.UpdateMailingContacts([]int64{mc.Id.Get()}, mc)
}

// UpdateMailingContacts updates existing mailing.contact records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMailingContacts(ids []int64, mc *MailingContact) error {
	return c.Update(MailingContactModel, ids, mc, nil)
}

// DeleteMailingContact deletes an existing mailing.contact record.
func (c *Client) DeleteMailingContact(id int64) error {
	return c.DeleteMailingContacts([]int64{id})
}

// DeleteMailingContacts deletes existing mailing.contact records.
func (c *Client) DeleteMailingContacts(ids []int64) error {
	return c.Delete(MailingContactModel, ids)
}

// GetMailingContact gets mailing.contact existing record.
func (c *Client) GetMailingContact(id int64) (*MailingContact, error) {
	mcs, err := c.GetMailingContacts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// GetMailingContacts gets mailing.contact existing records.
func (c *Client) GetMailingContacts(ids []int64) (*MailingContacts, error) {
	mcs := &MailingContacts{}
	if err := c.Read(MailingContactModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailingContact finds mailing.contact record by querying it with criteria.
func (c *Client) FindMailingContact(criteria *Criteria) (*MailingContact, error) {
	mcs := &MailingContacts{}
	if err := c.SearchRead(MailingContactModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// FindMailingContacts finds mailing.contact records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContacts(criteria *Criteria, options *Options) (*MailingContacts, error) {
	mcs := &MailingContacts{}
	if err := c.SearchRead(MailingContactModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailingContactIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingContactModel, criteria, options)
}

// FindMailingContactId finds record id by querying it with criteria.
func (c *Client) FindMailingContactId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingContactModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
