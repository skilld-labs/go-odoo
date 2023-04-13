package odoo

import (
	"fmt"
)

// MailThread represents mail.thread model.
type MailThread struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
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
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
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
	ids, err := c.Create(MailThreadModel, []interface{}{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailThread creates a new mail.thread model and returns its id.
func (c *Client) CreateMailThreads(mts []*MailThread) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MailThreadModel, vv)
}

// UpdateMailThread updates an existing mail.thread record.
func (c *Client) UpdateMailThread(mt *MailThread) error {
	return c.UpdateMailThreads([]int64{mt.Id.Get()}, mt)
}

// UpdateMailThreads updates existing mail.thread records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailThreads(ids []int64, mt *MailThread) error {
	return c.Update(MailThreadModel, ids, mt)
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
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.thread not found", id)
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
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("mail.thread was not found with criteria %v", criteria)
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
	ids, err := c.Search(MailThreadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailThreadId finds record id by querying it with criteria.
func (c *Client) FindMailThreadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.thread was not found with criteria %v and options %v", criteria, options)
}
