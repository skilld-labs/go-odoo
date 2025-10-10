package odoo

// MailingList represents mailing.list model.
type MailingList struct {
	Active                  *Bool     `xmlrpc:"active,omitempty"`
	ContactCount            *Int      `xmlrpc:"contact_count,omitempty"`
	ContactCountBlacklisted *Int      `xmlrpc:"contact_count_blacklisted,omitempty"`
	ContactCountEmail       *Int      `xmlrpc:"contact_count_email,omitempty"`
	ContactCountOptOut      *Int      `xmlrpc:"contact_count_opt_out,omitempty"`
	ContactIds              *Relation `xmlrpc:"contact_ids,omitempty"`
	ContactPctBlacklisted   *Float    `xmlrpc:"contact_pct_blacklisted,omitempty"`
	ContactPctBounce        *Float    `xmlrpc:"contact_pct_bounce,omitempty"`
	ContactPctOptOut        *Float    `xmlrpc:"contact_pct_opt_out,omitempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	IsPublic                *Bool     `xmlrpc:"is_public,omitempty"`
	MailingCount            *Int      `xmlrpc:"mailing_count,omitempty"`
	MailingIds              *Relation `xmlrpc:"mailing_ids,omitempty"`
	Name                    *String   `xmlrpc:"name,omitempty"`
	SubscriptionIds         *Relation `xmlrpc:"subscription_ids,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingLists represents array of mailing.list model.
type MailingLists []MailingList

// MailingListModel is the odoo model name.
const MailingListModel = "mailing.list"

// Many2One convert MailingList to *Many2One.
func (ml *MailingList) Many2One() *Many2One {
	return NewMany2One(ml.Id.Get(), "")
}

// CreateMailingList creates a new mailing.list model and returns its id.
func (c *Client) CreateMailingList(ml *MailingList) (int64, error) {
	ids, err := c.CreateMailingLists([]*MailingList{ml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingList creates a new mailing.list model and returns its id.
func (c *Client) CreateMailingLists(mls []*MailingList) ([]int64, error) {
	var vv []interface{}
	for _, v := range mls {
		vv = append(vv, v)
	}
	return c.Create(MailingListModel, vv, nil)
}

// UpdateMailingList updates an existing mailing.list record.
func (c *Client) UpdateMailingList(ml *MailingList) error {
	return c.UpdateMailingLists([]int64{ml.Id.Get()}, ml)
}

// UpdateMailingLists updates existing mailing.list records.
// All records (represented by ids) will be updated by ml values.
func (c *Client) UpdateMailingLists(ids []int64, ml *MailingList) error {
	return c.Update(MailingListModel, ids, ml, nil)
}

// DeleteMailingList deletes an existing mailing.list record.
func (c *Client) DeleteMailingList(id int64) error {
	return c.DeleteMailingLists([]int64{id})
}

// DeleteMailingLists deletes existing mailing.list records.
func (c *Client) DeleteMailingLists(ids []int64) error {
	return c.Delete(MailingListModel, ids)
}

// GetMailingList gets mailing.list existing record.
func (c *Client) GetMailingList(id int64) (*MailingList, error) {
	mls, err := c.GetMailingLists([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mls)[0]), nil
}

// GetMailingLists gets mailing.list existing records.
func (c *Client) GetMailingLists(ids []int64) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.Read(MailingListModel, ids, nil, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingList finds mailing.list record by querying it with criteria.
func (c *Client) FindMailingList(criteria *Criteria) (*MailingList, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, NewOptions().Limit(1), mls); err != nil {
		return nil, err
	}
	return &((*mls)[0]), nil
}

// FindMailingLists finds mailing.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingLists(criteria *Criteria, options *Options) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, options, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingListIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingListModel, criteria, options)
}

// FindMailingListId finds record id by querying it with criteria.
func (c *Client) FindMailingListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
