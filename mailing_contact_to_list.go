package odoo

// MailingContactToList represents mailing.contact.to.list model.
type MailingContactToList struct {
	ContactIds    *Relation `xmlrpc:"contact_ids,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	MailingListId *Many2One `xmlrpc:"mailing_list_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingContactToLists represents array of mailing.contact.to.list model.
type MailingContactToLists []MailingContactToList

// MailingContactToListModel is the odoo model name.
const MailingContactToListModel = "mailing.contact.to.list"

// Many2One convert MailingContactToList to *Many2One.
func (mctl *MailingContactToList) Many2One() *Many2One {
	return NewMany2One(mctl.Id.Get(), "")
}

// CreateMailingContactToList creates a new mailing.contact.to.list model and returns its id.
func (c *Client) CreateMailingContactToList(mctl *MailingContactToList) (int64, error) {
	ids, err := c.CreateMailingContactToLists([]*MailingContactToList{mctl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingContactToList creates a new mailing.contact.to.list model and returns its id.
func (c *Client) CreateMailingContactToLists(mctls []*MailingContactToList) ([]int64, error) {
	var vv []interface{}
	for _, v := range mctls {
		vv = append(vv, v)
	}
	return c.Create(MailingContactToListModel, vv, nil)
}

// UpdateMailingContactToList updates an existing mailing.contact.to.list record.
func (c *Client) UpdateMailingContactToList(mctl *MailingContactToList) error {
	return c.UpdateMailingContactToLists([]int64{mctl.Id.Get()}, mctl)
}

// UpdateMailingContactToLists updates existing mailing.contact.to.list records.
// All records (represented by ids) will be updated by mctl values.
func (c *Client) UpdateMailingContactToLists(ids []int64, mctl *MailingContactToList) error {
	return c.Update(MailingContactToListModel, ids, mctl, nil)
}

// DeleteMailingContactToList deletes an existing mailing.contact.to.list record.
func (c *Client) DeleteMailingContactToList(id int64) error {
	return c.DeleteMailingContactToLists([]int64{id})
}

// DeleteMailingContactToLists deletes existing mailing.contact.to.list records.
func (c *Client) DeleteMailingContactToLists(ids []int64) error {
	return c.Delete(MailingContactToListModel, ids)
}

// GetMailingContactToList gets mailing.contact.to.list existing record.
func (c *Client) GetMailingContactToList(id int64) (*MailingContactToList, error) {
	mctls, err := c.GetMailingContactToLists([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mctls)[0]), nil
}

// GetMailingContactToLists gets mailing.contact.to.list existing records.
func (c *Client) GetMailingContactToLists(ids []int64) (*MailingContactToLists, error) {
	mctls := &MailingContactToLists{}
	if err := c.Read(MailingContactToListModel, ids, nil, mctls); err != nil {
		return nil, err
	}
	return mctls, nil
}

// FindMailingContactToList finds mailing.contact.to.list record by querying it with criteria.
func (c *Client) FindMailingContactToList(criteria *Criteria) (*MailingContactToList, error) {
	mctls := &MailingContactToLists{}
	if err := c.SearchRead(MailingContactToListModel, criteria, NewOptions().Limit(1), mctls); err != nil {
		return nil, err
	}
	return &((*mctls)[0]), nil
}

// FindMailingContactToLists finds mailing.contact.to.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactToLists(criteria *Criteria, options *Options) (*MailingContactToLists, error) {
	mctls := &MailingContactToLists{}
	if err := c.SearchRead(MailingContactToListModel, criteria, options, mctls); err != nil {
		return nil, err
	}
	return mctls, nil
}

// FindMailingContactToListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactToListIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingContactToListModel, criteria, options)
}

// FindMailingContactToListId finds record id by querying it with criteria.
func (c *Client) FindMailingContactToListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingContactToListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
