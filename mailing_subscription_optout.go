package odoo

// MailingSubscriptionOptout represents mailing.subscription.optout model.
type MailingSubscriptionOptout struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	IsFeedback  *Bool     `xmlrpc:"is_feedback,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingSubscriptionOptouts represents array of mailing.subscription.optout model.
type MailingSubscriptionOptouts []MailingSubscriptionOptout

// MailingSubscriptionOptoutModel is the odoo model name.
const MailingSubscriptionOptoutModel = "mailing.subscription.optout"

// Many2One convert MailingSubscriptionOptout to *Many2One.
func (mso *MailingSubscriptionOptout) Many2One() *Many2One {
	return NewMany2One(mso.Id.Get(), "")
}

// CreateMailingSubscriptionOptout creates a new mailing.subscription.optout model and returns its id.
func (c *Client) CreateMailingSubscriptionOptout(mso *MailingSubscriptionOptout) (int64, error) {
	ids, err := c.CreateMailingSubscriptionOptouts([]*MailingSubscriptionOptout{mso})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingSubscriptionOptout creates a new mailing.subscription.optout model and returns its id.
func (c *Client) CreateMailingSubscriptionOptouts(msos []*MailingSubscriptionOptout) ([]int64, error) {
	var vv []interface{}
	for _, v := range msos {
		vv = append(vv, v)
	}
	return c.Create(MailingSubscriptionOptoutModel, vv, nil)
}

// UpdateMailingSubscriptionOptout updates an existing mailing.subscription.optout record.
func (c *Client) UpdateMailingSubscriptionOptout(mso *MailingSubscriptionOptout) error {
	return c.UpdateMailingSubscriptionOptouts([]int64{mso.Id.Get()}, mso)
}

// UpdateMailingSubscriptionOptouts updates existing mailing.subscription.optout records.
// All records (represented by ids) will be updated by mso values.
func (c *Client) UpdateMailingSubscriptionOptouts(ids []int64, mso *MailingSubscriptionOptout) error {
	return c.Update(MailingSubscriptionOptoutModel, ids, mso, nil)
}

// DeleteMailingSubscriptionOptout deletes an existing mailing.subscription.optout record.
func (c *Client) DeleteMailingSubscriptionOptout(id int64) error {
	return c.DeleteMailingSubscriptionOptouts([]int64{id})
}

// DeleteMailingSubscriptionOptouts deletes existing mailing.subscription.optout records.
func (c *Client) DeleteMailingSubscriptionOptouts(ids []int64) error {
	return c.Delete(MailingSubscriptionOptoutModel, ids)
}

// GetMailingSubscriptionOptout gets mailing.subscription.optout existing record.
func (c *Client) GetMailingSubscriptionOptout(id int64) (*MailingSubscriptionOptout, error) {
	msos, err := c.GetMailingSubscriptionOptouts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*msos)[0]), nil
}

// GetMailingSubscriptionOptouts gets mailing.subscription.optout existing records.
func (c *Client) GetMailingSubscriptionOptouts(ids []int64) (*MailingSubscriptionOptouts, error) {
	msos := &MailingSubscriptionOptouts{}
	if err := c.Read(MailingSubscriptionOptoutModel, ids, nil, msos); err != nil {
		return nil, err
	}
	return msos, nil
}

// FindMailingSubscriptionOptout finds mailing.subscription.optout record by querying it with criteria.
func (c *Client) FindMailingSubscriptionOptout(criteria *Criteria) (*MailingSubscriptionOptout, error) {
	msos := &MailingSubscriptionOptouts{}
	if err := c.SearchRead(MailingSubscriptionOptoutModel, criteria, NewOptions().Limit(1), msos); err != nil {
		return nil, err
	}
	return &((*msos)[0]), nil
}

// FindMailingSubscriptionOptouts finds mailing.subscription.optout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSubscriptionOptouts(criteria *Criteria, options *Options) (*MailingSubscriptionOptouts, error) {
	msos := &MailingSubscriptionOptouts{}
	if err := c.SearchRead(MailingSubscriptionOptoutModel, criteria, options, msos); err != nil {
		return nil, err
	}
	return msos, nil
}

// FindMailingSubscriptionOptoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSubscriptionOptoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingSubscriptionOptoutModel, criteria, options)
}

// FindMailingSubscriptionOptoutId finds record id by querying it with criteria.
func (c *Client) FindMailingSubscriptionOptoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingSubscriptionOptoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
