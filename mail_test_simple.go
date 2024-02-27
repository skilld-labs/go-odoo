package odoo

// MailTestSimple represents mail.test.simple model.
type MailTestSimple struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	Description              *String   `xmlrpc:"description,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	EmailFrom                *String   `xmlrpc:"email_from,omitempty"`
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
	Name                     *String   `xmlrpc:"name,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailTestSimples represents array of mail.test.simple model.
type MailTestSimples []MailTestSimple

// MailTestSimpleModel is the odoo model name.
const MailTestSimpleModel = "mail.test.simple"

// Many2One convert MailTestSimple to *Many2One.
func (mts *MailTestSimple) Many2One() *Many2One {
	return NewMany2One(mts.Id.Get(), "")
}

// CreateMailTestSimple creates a new mail.test.simple model and returns its id.
func (c *Client) CreateMailTestSimple(mts *MailTestSimple) (int64, error) {
	ids, err := c.CreateMailTestSimples([]*MailTestSimple{mts})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTestSimple creates a new mail.test.simple model and returns its id.
func (c *Client) CreateMailTestSimples(mtss []*MailTestSimple) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtss {
		vv = append(vv, v)
	}
	return c.Create(MailTestSimpleModel, vv, nil)
}

// UpdateMailTestSimple updates an existing mail.test.simple record.
func (c *Client) UpdateMailTestSimple(mts *MailTestSimple) error {
	return c.UpdateMailTestSimples([]int64{mts.Id.Get()}, mts)
}

// UpdateMailTestSimples updates existing mail.test.simple records.
// All records (represented by ids) will be updated by mts values.
func (c *Client) UpdateMailTestSimples(ids []int64, mts *MailTestSimple) error {
	return c.Update(MailTestSimpleModel, ids, mts, nil)
}

// DeleteMailTestSimple deletes an existing mail.test.simple record.
func (c *Client) DeleteMailTestSimple(id int64) error {
	return c.DeleteMailTestSimples([]int64{id})
}

// DeleteMailTestSimples deletes existing mail.test.simple records.
func (c *Client) DeleteMailTestSimples(ids []int64) error {
	return c.Delete(MailTestSimpleModel, ids)
}

// GetMailTestSimple gets mail.test.simple existing record.
func (c *Client) GetMailTestSimple(id int64) (*MailTestSimple, error) {
	mtss, err := c.GetMailTestSimples([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mtss)[0]), nil
}

// GetMailTestSimples gets mail.test.simple existing records.
func (c *Client) GetMailTestSimples(ids []int64) (*MailTestSimples, error) {
	mtss := &MailTestSimples{}
	if err := c.Read(MailTestSimpleModel, ids, nil, mtss); err != nil {
		return nil, err
	}
	return mtss, nil
}

// FindMailTestSimple finds mail.test.simple record by querying it with criteria.
func (c *Client) FindMailTestSimple(criteria *Criteria) (*MailTestSimple, error) {
	mtss := &MailTestSimples{}
	if err := c.SearchRead(MailTestSimpleModel, criteria, NewOptions().Limit(1), mtss); err != nil {
		return nil, err
	}
	return &((*mtss)[0]), nil
}

// FindMailTestSimples finds mail.test.simple records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTestSimples(criteria *Criteria, options *Options) (*MailTestSimples, error) {
	mtss := &MailTestSimples{}
	if err := c.SearchRead(MailTestSimpleModel, criteria, options, mtss); err != nil {
		return nil, err
	}
	return mtss, nil
}

// FindMailTestSimpleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTestSimpleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailTestSimpleModel, criteria, options)
}

// FindMailTestSimpleId finds record id by querying it with criteria.
func (c *Client) FindMailTestSimpleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTestSimpleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
