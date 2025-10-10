package odoo

// MailingSubscription represents mailing.subscription model.
type MailingSubscription struct {
	ContactId      *Many2One `xmlrpc:"contact_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	IsBlacklisted  *Bool     `xmlrpc:"is_blacklisted,omitempty"`
	ListId         *Many2One `xmlrpc:"list_id,omitempty"`
	MessageBounce  *Int      `xmlrpc:"message_bounce,omitempty"`
	OptOut         *Bool     `xmlrpc:"opt_out,omitempty"`
	OptOutDatetime *Time     `xmlrpc:"opt_out_datetime,omitempty"`
	OptOutReasonId *Many2One `xmlrpc:"opt_out_reason_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingSubscriptions represents array of mailing.subscription model.
type MailingSubscriptions []MailingSubscription

// MailingSubscriptionModel is the odoo model name.
const MailingSubscriptionModel = "mailing.subscription"

// Many2One convert MailingSubscription to *Many2One.
func (ms *MailingSubscription) Many2One() *Many2One {
	return NewMany2One(ms.Id.Get(), "")
}

// CreateMailingSubscription creates a new mailing.subscription model and returns its id.
func (c *Client) CreateMailingSubscription(ms *MailingSubscription) (int64, error) {
	ids, err := c.CreateMailingSubscriptions([]*MailingSubscription{ms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingSubscription creates a new mailing.subscription model and returns its id.
func (c *Client) CreateMailingSubscriptions(mss []*MailingSubscription) ([]int64, error) {
	var vv []interface{}
	for _, v := range mss {
		vv = append(vv, v)
	}
	return c.Create(MailingSubscriptionModel, vv, nil)
}

// UpdateMailingSubscription updates an existing mailing.subscription record.
func (c *Client) UpdateMailingSubscription(ms *MailingSubscription) error {
	return c.UpdateMailingSubscriptions([]int64{ms.Id.Get()}, ms)
}

// UpdateMailingSubscriptions updates existing mailing.subscription records.
// All records (represented by ids) will be updated by ms values.
func (c *Client) UpdateMailingSubscriptions(ids []int64, ms *MailingSubscription) error {
	return c.Update(MailingSubscriptionModel, ids, ms, nil)
}

// DeleteMailingSubscription deletes an existing mailing.subscription record.
func (c *Client) DeleteMailingSubscription(id int64) error {
	return c.DeleteMailingSubscriptions([]int64{id})
}

// DeleteMailingSubscriptions deletes existing mailing.subscription records.
func (c *Client) DeleteMailingSubscriptions(ids []int64) error {
	return c.Delete(MailingSubscriptionModel, ids)
}

// GetMailingSubscription gets mailing.subscription existing record.
func (c *Client) GetMailingSubscription(id int64) (*MailingSubscription, error) {
	mss, err := c.GetMailingSubscriptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mss)[0]), nil
}

// GetMailingSubscriptions gets mailing.subscription existing records.
func (c *Client) GetMailingSubscriptions(ids []int64) (*MailingSubscriptions, error) {
	mss := &MailingSubscriptions{}
	if err := c.Read(MailingSubscriptionModel, ids, nil, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailingSubscription finds mailing.subscription record by querying it with criteria.
func (c *Client) FindMailingSubscription(criteria *Criteria) (*MailingSubscription, error) {
	mss := &MailingSubscriptions{}
	if err := c.SearchRead(MailingSubscriptionModel, criteria, NewOptions().Limit(1), mss); err != nil {
		return nil, err
	}
	return &((*mss)[0]), nil
}

// FindMailingSubscriptions finds mailing.subscription records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSubscriptions(criteria *Criteria, options *Options) (*MailingSubscriptions, error) {
	mss := &MailingSubscriptions{}
	if err := c.SearchRead(MailingSubscriptionModel, criteria, options, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailingSubscriptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSubscriptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingSubscriptionModel, criteria, options)
}

// FindMailingSubscriptionId finds record id by querying it with criteria.
func (c *Client) FindMailingSubscriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingSubscriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
