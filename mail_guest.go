package odoo

// MailGuest represents mail.guest model.
type MailGuest struct {
	AccessToken *String    `xmlrpc:"access_token,omitempty"`
	Avatar1024  *String    `xmlrpc:"avatar_1024,omitempty"`
	Avatar128   *String    `xmlrpc:"avatar_128,omitempty"`
	Avatar1920  *String    `xmlrpc:"avatar_1920,omitempty"`
	Avatar256   *String    `xmlrpc:"avatar_256,omitempty"`
	Avatar512   *String    `xmlrpc:"avatar_512,omitempty"`
	ChannelIds  *Relation  `xmlrpc:"channel_ids,omitempty"`
	CountryId   *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	ImStatus    *String    `xmlrpc:"im_status,omitempty"`
	Image1024   *String    `xmlrpc:"image_1024,omitempty"`
	Image128    *String    `xmlrpc:"image_128,omitempty"`
	Image1920   *String    `xmlrpc:"image_1920,omitempty"`
	Image256    *String    `xmlrpc:"image_256,omitempty"`
	Image512    *String    `xmlrpc:"image_512,omitempty"`
	Lang        *Selection `xmlrpc:"lang,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Timezone    *Selection `xmlrpc:"timezone,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailGuests represents array of mail.guest model.
type MailGuests []MailGuest

// MailGuestModel is the odoo model name.
const MailGuestModel = "mail.guest"

// Many2One convert MailGuest to *Many2One.
func (mg *MailGuest) Many2One() *Many2One {
	return NewMany2One(mg.Id.Get(), "")
}

// CreateMailGuest creates a new mail.guest model and returns its id.
func (c *Client) CreateMailGuest(mg *MailGuest) (int64, error) {
	ids, err := c.CreateMailGuests([]*MailGuest{mg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailGuest creates a new mail.guest model and returns its id.
func (c *Client) CreateMailGuests(mgs []*MailGuest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mgs {
		vv = append(vv, v)
	}
	return c.Create(MailGuestModel, vv, nil)
}

// UpdateMailGuest updates an existing mail.guest record.
func (c *Client) UpdateMailGuest(mg *MailGuest) error {
	return c.UpdateMailGuests([]int64{mg.Id.Get()}, mg)
}

// UpdateMailGuests updates existing mail.guest records.
// All records (represented by ids) will be updated by mg values.
func (c *Client) UpdateMailGuests(ids []int64, mg *MailGuest) error {
	return c.Update(MailGuestModel, ids, mg, nil)
}

// DeleteMailGuest deletes an existing mail.guest record.
func (c *Client) DeleteMailGuest(id int64) error {
	return c.DeleteMailGuests([]int64{id})
}

// DeleteMailGuests deletes existing mail.guest records.
func (c *Client) DeleteMailGuests(ids []int64) error {
	return c.Delete(MailGuestModel, ids)
}

// GetMailGuest gets mail.guest existing record.
func (c *Client) GetMailGuest(id int64) (*MailGuest, error) {
	mgs, err := c.GetMailGuests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mgs)[0]), nil
}

// GetMailGuests gets mail.guest existing records.
func (c *Client) GetMailGuests(ids []int64) (*MailGuests, error) {
	mgs := &MailGuests{}
	if err := c.Read(MailGuestModel, ids, nil, mgs); err != nil {
		return nil, err
	}
	return mgs, nil
}

// FindMailGuest finds mail.guest record by querying it with criteria.
func (c *Client) FindMailGuest(criteria *Criteria) (*MailGuest, error) {
	mgs := &MailGuests{}
	if err := c.SearchRead(MailGuestModel, criteria, NewOptions().Limit(1), mgs); err != nil {
		return nil, err
	}
	return &((*mgs)[0]), nil
}

// FindMailGuests finds mail.guest records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGuests(criteria *Criteria, options *Options) (*MailGuests, error) {
	mgs := &MailGuests{}
	if err := c.SearchRead(MailGuestModel, criteria, options, mgs); err != nil {
		return nil, err
	}
	return mgs, nil
}

// FindMailGuestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGuestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailGuestModel, criteria, options)
}

// FindMailGuestId finds record id by querying it with criteria.
func (c *Client) FindMailGuestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailGuestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
