package odoo

import (
	"fmt"
)

// MailTestSimple represents mail.test.simple model.
type MailTestSimple struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	Description              *String   `xmlrpc:"description,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	EmailFrom                *String   `xmlrpc:"email_from,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(MailTestSimpleModel, mts)
}

// UpdateMailTestSimple updates an existing mail.test.simple record.
func (c *Client) UpdateMailTestSimple(mts *MailTestSimple) error {
	return c.UpdateMailTestSimples([]int64{mts.Id.Get()}, mts)
}

// UpdateMailTestSimples updates existing mail.test.simple records.
// All records (represented by ids) will be updated by mts values.
func (c *Client) UpdateMailTestSimples(ids []int64, mts *MailTestSimple) error {
	return c.Update(MailTestSimpleModel, ids, mts)
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
	if mtss != nil && len(*mtss) > 0 {
		return &((*mtss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.test.simple not found", id)
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
	if mtss != nil && len(*mtss) > 0 {
		return &((*mtss)[0]), nil
	}
	return nil, fmt.Errorf("mail.test.simple was not found with criteria %v", criteria)
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
	ids, err := c.Search(MailTestSimpleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTestSimpleId finds record id by querying it with criteria.
func (c *Client) FindMailTestSimpleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTestSimpleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.test.simple was not found with criteria %v and options %v", criteria, options)
}
