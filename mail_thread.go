package odoo

// MailThread represents mail.thread model.
type MailThread struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
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
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
}

// MailThreads represents array of mail.thread model.
type MailThreads []MailThread

// MailThreadModel is the odoo model name.
const MailThreadModel = "mail.thread"

// Many2One convert MailThread to *Many2One.
func (mt *MailThread) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMailThread creates a new mail.thread model and returns its id.
func (c *Client) CreateMailThread(mt *MailThread) (int64, error) {
	ids, err := c.CreateMailThreads([]*MailThread{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailThreads creates a new mail.thread model and returns its id.
func (c *Client) CreateMailThreads(mts []*MailThread) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MailThreadModel, vv, nil)
}

// UpdateMailThread updates an existing mail.thread record.
func (c *Client) UpdateMailThread(mt *MailThread) error {
	return c.UpdateMailThreads([]int64{mt.Id.Get()}, mt)
}

// UpdateMailThreads updates existing mail.thread records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailThreads(ids []int64, mt *MailThread) error {
	return c.Update(MailThreadModel, ids, mt, nil)
}

// DeleteMailThread deletes an existing mail.thread record.
func (c *Client) DeleteMailThread(id int64) error {
	return c.DeleteMailThreads([]int64{id})
}

// DeleteMailThreads deletes existing mail.thread records.
func (c *Client) DeleteMailThreads(ids []int64) error {
	return c.Delete(MailThreadModel, ids)
}

// GetMailThread gets mail.thread existing record.
func (c *Client) GetMailThread(id int64) (*MailThread, error) {
	mts, err := c.GetMailThreads([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// GetMailThreads gets mail.thread existing records.
func (c *Client) GetMailThreads(ids []int64) (*MailThreads, error) {
	mts := &MailThreads{}
	if err := c.Read(MailThreadModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailThread finds mail.thread record by querying it with criteria.
func (c *Client) FindMailThread(criteria *Criteria) (*MailThread, error) {
	mts := &MailThreads{}
	if err := c.SearchRead(MailThreadModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	return &((*mts)[0]), nil
}

// FindMailThreads finds mail.thread records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreads(criteria *Criteria, options *Options) (*MailThreads, error) {
	mts := &MailThreads{}
	if err := c.SearchRead(MailThreadModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailThreadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailThreadModel, criteria, options)
}

// FindMailThreadId finds record id by querying it with criteria.
func (c *Client) FindMailThreadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
