package odoo

// MailPresence represents mail.presence model.
type MailPresence struct {
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	GuestId      *Many2One  `xmlrpc:"guest_id,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	LastPoll     *Time      `xmlrpc:"last_poll,omitempty"`
	LastPresence *Time      `xmlrpc:"last_presence,omitempty"`
	Status       *Selection `xmlrpc:"status,omitempty"`
	UserId       *Many2One  `xmlrpc:"user_id,omitempty"`
}

// MailPresences represents array of mail.presence model.
type MailPresences []MailPresence

// MailPresenceModel is the odoo model name.
const MailPresenceModel = "mail.presence"

// Many2One convert MailPresence to *Many2One.
func (mp *MailPresence) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMailPresence creates a new mail.presence model and returns its id.
func (c *Client) CreateMailPresence(mp *MailPresence) (int64, error) {
	ids, err := c.CreateMailPresences([]*MailPresence{mp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailPresence creates a new mail.presence model and returns its id.
func (c *Client) CreateMailPresences(mps []*MailPresence) ([]int64, error) {
	var vv []interface{}
	for _, v := range mps {
		vv = append(vv, v)
	}
	return c.Create(MailPresenceModel, vv, nil)
}

// UpdateMailPresence updates an existing mail.presence record.
func (c *Client) UpdateMailPresence(mp *MailPresence) error {
	return c.UpdateMailPresences([]int64{mp.Id.Get()}, mp)
}

// UpdateMailPresences updates existing mail.presence records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMailPresences(ids []int64, mp *MailPresence) error {
	return c.Update(MailPresenceModel, ids, mp, nil)
}

// DeleteMailPresence deletes an existing mail.presence record.
func (c *Client) DeleteMailPresence(id int64) error {
	return c.DeleteMailPresences([]int64{id})
}

// DeleteMailPresences deletes existing mail.presence records.
func (c *Client) DeleteMailPresences(ids []int64) error {
	return c.Delete(MailPresenceModel, ids)
}

// GetMailPresence gets mail.presence existing record.
func (c *Client) GetMailPresence(id int64) (*MailPresence, error) {
	mps, err := c.GetMailPresences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// GetMailPresences gets mail.presence existing records.
func (c *Client) GetMailPresences(ids []int64) (*MailPresences, error) {
	mps := &MailPresences{}
	if err := c.Read(MailPresenceModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMailPresence finds mail.presence record by querying it with criteria.
func (c *Client) FindMailPresence(criteria *Criteria) (*MailPresence, error) {
	mps := &MailPresences{}
	if err := c.SearchRead(MailPresenceModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// FindMailPresences finds mail.presence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPresences(criteria *Criteria, options *Options) (*MailPresences, error) {
	mps := &MailPresences{}
	if err := c.SearchRead(MailPresenceModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMailPresenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPresenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailPresenceModel, criteria, options)
}

// FindMailPresenceId finds record id by querying it with criteria.
func (c *Client) FindMailPresenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailPresenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
