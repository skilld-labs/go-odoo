package odoo

// MailMassMailingList represents mail.mass_mailing.list model.
type MailMassMailingList struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	ContactNbr  *Int      `xmlrpc:"contact_nbr,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMassMailingLists represents array of mail.mass_mailing.list model.
type MailMassMailingLists []MailMassMailingList

// MailMassMailingListModel is the odoo model name.
const MailMassMailingListModel = "mail.mass_mailing.list"

// Many2One convert MailMassMailingList to *Many2One.
func (mml *MailMassMailingList) Many2One() *Many2One {
	return NewMany2One(mml.Id.Get(), "")
}

// CreateMailMassMailingList creates a new mail.mass_mailing.list model and returns its id.
func (c *Client) CreateMailMassMailingList(mml *MailMassMailingList) (int64, error) {
	ids, err := c.CreateMailMassMailingLists([]*MailMassMailingList{mml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingList creates a new mail.mass_mailing.list model and returns its id.
func (c *Client) CreateMailMassMailingLists(mmls []*MailMassMailingList) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmls {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingListModel, vv, nil)
}

// UpdateMailMassMailingList updates an existing mail.mass_mailing.list record.
func (c *Client) UpdateMailMassMailingList(mml *MailMassMailingList) error {
	return c.UpdateMailMassMailingLists([]int64{mml.Id.Get()}, mml)
}

// UpdateMailMassMailingLists updates existing mail.mass_mailing.list records.
// All records (represented by ids) will be updated by mml values.
func (c *Client) UpdateMailMassMailingLists(ids []int64, mml *MailMassMailingList) error {
	return c.Update(MailMassMailingListModel, ids, mml, nil)
}

// DeleteMailMassMailingList deletes an existing mail.mass_mailing.list record.
func (c *Client) DeleteMailMassMailingList(id int64) error {
	return c.DeleteMailMassMailingLists([]int64{id})
}

// DeleteMailMassMailingLists deletes existing mail.mass_mailing.list records.
func (c *Client) DeleteMailMassMailingLists(ids []int64) error {
	return c.Delete(MailMassMailingListModel, ids)
}

// GetMailMassMailingList gets mail.mass_mailing.list existing record.
func (c *Client) GetMailMassMailingList(id int64) (*MailMassMailingList, error) {
	mmls, err := c.GetMailMassMailingLists([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmls)[0]), nil
}

// GetMailMassMailingLists gets mail.mass_mailing.list existing records.
func (c *Client) GetMailMassMailingLists(ids []int64) (*MailMassMailingLists, error) {
	mmls := &MailMassMailingLists{}
	if err := c.Read(MailMassMailingListModel, ids, nil, mmls); err != nil {
		return nil, err
	}
	return mmls, nil
}

// FindMailMassMailingList finds mail.mass_mailing.list record by querying it with criteria.
func (c *Client) FindMailMassMailingList(criteria *Criteria) (*MailMassMailingList, error) {
	mmls := &MailMassMailingLists{}
	if err := c.SearchRead(MailMassMailingListModel, criteria, NewOptions().Limit(1), mmls); err != nil {
		return nil, err
	}
	return &((*mmls)[0]), nil
}

// FindMailMassMailingLists finds mail.mass_mailing.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingLists(criteria *Criteria, options *Options) (*MailMassMailingLists, error) {
	mmls := &MailMassMailingLists{}
	if err := c.SearchRead(MailMassMailingListModel, criteria, options, mmls); err != nil {
		return nil, err
	}
	return mmls, nil
}

// FindMailMassMailingListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingListIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMassMailingListModel, criteria, options)
}

// FindMailMassMailingListId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
