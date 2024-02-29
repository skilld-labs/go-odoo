package odoo

// MailFollowers represents mail.followers model.
type MailFollowers struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	ChannelId   *Many2One `xmlrpc:"channel_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	ResId       *Int      `xmlrpc:"res_id,omitempty"`
	ResModel    *String   `xmlrpc:"res_model,omitempty"`
	SubtypeIds  *Relation `xmlrpc:"subtype_ids,omitempty"`
}

// MailFollowerss represents array of mail.followers model.
type MailFollowerss []MailFollowers

// MailFollowersModel is the odoo model name.
const MailFollowersModel = "mail.followers"

// Many2One convert MailFollowers to *Many2One.
func (mf *MailFollowers) Many2One() *Many2One {
	return NewMany2One(mf.Id.Get(), "")
}

// CreateMailFollowers creates a new mail.followers model and returns its id.
func (c *Client) CreateMailFollowers(mf *MailFollowers) (int64, error) {
	ids, err := c.CreateMailFollowerss([]*MailFollowers{mf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailFollowerss creates a new mail.followers model and returns its id.
func (c *Client) CreateMailFollowerss(mfs []*MailFollowers) ([]int64, error) {
	var vv []interface{}
	for _, v := range mfs {
		vv = append(vv, v)
	}
	return c.Create(MailFollowersModel, vv, nil)
}

// UpdateMailFollowers updates an existing mail.followers record.
func (c *Client) UpdateMailFollowers(mf *MailFollowers) error {
	return c.UpdateMailFollowerss([]int64{mf.Id.Get()}, mf)
}

// UpdateMailFollowerss updates existing mail.followers records.
// All records (represented by ids) will be updated by mf values.
func (c *Client) UpdateMailFollowerss(ids []int64, mf *MailFollowers) error {
	return c.Update(MailFollowersModel, ids, mf, nil)
}

// DeleteMailFollowers deletes an existing mail.followers record.
func (c *Client) DeleteMailFollowers(id int64) error {
	return c.DeleteMailFollowerss([]int64{id})
}

// DeleteMailFollowerss deletes existing mail.followers records.
func (c *Client) DeleteMailFollowerss(ids []int64) error {
	return c.Delete(MailFollowersModel, ids)
}

// GetMailFollowers gets mail.followers existing record.
func (c *Client) GetMailFollowers(id int64) (*MailFollowers, error) {
	mfs, err := c.GetMailFollowerss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mfs)[0]), nil
}

// GetMailFollowerss gets mail.followers existing records.
func (c *Client) GetMailFollowerss(ids []int64) (*MailFollowerss, error) {
	mfs := &MailFollowerss{}
	if err := c.Read(MailFollowersModel, ids, nil, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailFollowers finds mail.followers record by querying it with criteria.
func (c *Client) FindMailFollowers(criteria *Criteria) (*MailFollowers, error) {
	mfs := &MailFollowerss{}
	if err := c.SearchRead(MailFollowersModel, criteria, NewOptions().Limit(1), mfs); err != nil {
		return nil, err
	}
	return &((*mfs)[0]), nil
}

// FindMailFollowerss finds mail.followers records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowerss(criteria *Criteria, options *Options) (*MailFollowerss, error) {
	mfs := &MailFollowerss{}
	if err := c.SearchRead(MailFollowersModel, criteria, options, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailFollowersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowersIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailFollowersModel, criteria, options)
}

// FindMailFollowersId finds record id by querying it with criteria.
func (c *Client) FindMailFollowersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailFollowersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
