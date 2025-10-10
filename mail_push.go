package odoo

// MailPush represents mail.push model.
type MailPush struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	MailPushDeviceId *Many2One `xmlrpc:"mail_push_device_id,omitempty"`
	Payload          *String   `xmlrpc:"payload,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailPushs represents array of mail.push model.
type MailPushs []MailPush

// MailPushModel is the odoo model name.
const MailPushModel = "mail.push"

// Many2One convert MailPush to *Many2One.
func (mp *MailPush) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMailPush creates a new mail.push model and returns its id.
func (c *Client) CreateMailPush(mp *MailPush) (int64, error) {
	ids, err := c.CreateMailPushs([]*MailPush{mp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailPush creates a new mail.push model and returns its id.
func (c *Client) CreateMailPushs(mps []*MailPush) ([]int64, error) {
	var vv []interface{}
	for _, v := range mps {
		vv = append(vv, v)
	}
	return c.Create(MailPushModel, vv, nil)
}

// UpdateMailPush updates an existing mail.push record.
func (c *Client) UpdateMailPush(mp *MailPush) error {
	return c.UpdateMailPushs([]int64{mp.Id.Get()}, mp)
}

// UpdateMailPushs updates existing mail.push records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMailPushs(ids []int64, mp *MailPush) error {
	return c.Update(MailPushModel, ids, mp, nil)
}

// DeleteMailPush deletes an existing mail.push record.
func (c *Client) DeleteMailPush(id int64) error {
	return c.DeleteMailPushs([]int64{id})
}

// DeleteMailPushs deletes existing mail.push records.
func (c *Client) DeleteMailPushs(ids []int64) error {
	return c.Delete(MailPushModel, ids)
}

// GetMailPush gets mail.push existing record.
func (c *Client) GetMailPush(id int64) (*MailPush, error) {
	mps, err := c.GetMailPushs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// GetMailPushs gets mail.push existing records.
func (c *Client) GetMailPushs(ids []int64) (*MailPushs, error) {
	mps := &MailPushs{}
	if err := c.Read(MailPushModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMailPush finds mail.push record by querying it with criteria.
func (c *Client) FindMailPush(criteria *Criteria) (*MailPush, error) {
	mps := &MailPushs{}
	if err := c.SearchRead(MailPushModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// FindMailPushs finds mail.push records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushs(criteria *Criteria, options *Options) (*MailPushs, error) {
	mps := &MailPushs{}
	if err := c.SearchRead(MailPushModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMailPushIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailPushModel, criteria, options)
}

// FindMailPushId finds record id by querying it with criteria.
func (c *Client) FindMailPushId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailPushModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
