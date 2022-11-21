package odoo

import (
	"fmt"
)

// MailTest represents mail.test model.
type MailTest struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailTests represents array of mail.test model.
type MailTests []MailTest

// MailTestModel is the odoo model name.
const MailTestModel = "mail.test"

// Many2One convert MailTest to *Many2One.
func (mt *MailTest) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMailTest creates a new mail.test model and returns its id.
func (c *Client) CreateMailTest(mt *MailTest) (int64, error) {
	return c.Create(MailTestModel, mt)
}

// UpdateMailTest updates an existing mail.test record.
func (c *Client) UpdateMailTest(mt *MailTest) error {
	return c.UpdateMailTests([]int64{mt.Id.Get()}, mt)
}

// UpdateMailTests updates existing mail.test records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailTests(ids []int64, mt *MailTest) error {
	return c.Update(MailTestModel, ids, mt)
}

// DeleteMailTest deletes an existing mail.test record.
func (c *Client) DeleteMailTest(id int64) error {
	return c.DeleteMailTests([]int64{id})
}

// DeleteMailTests deletes existing mail.test records.
func (c *Client) DeleteMailTests(ids []int64) error {
	return c.Delete(MailTestModel, ids)
}

// GetMailTest gets mail.test existing record.
func (c *Client) GetMailTest(id int64) (*MailTest, error) {
	mts, err := c.GetMailTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.test not found", id)
}

// GetMailTests gets mail.test existing records.
func (c *Client) GetMailTests(ids []int64) (*MailTests, error) {
	mts := &MailTests{}
	if err := c.Read(MailTestModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailTest finds mail.test record by querying it with criteria.
func (c *Client) FindMailTest(criteria *Criteria) (*MailTest, error) {
	mts := &MailTests{}
	if err := c.SearchRead(MailTestModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("no mail.test was found with criteria %v", criteria)
}

// FindMailTests finds mail.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTests(criteria *Criteria, options *Options) (*MailTests, error) {
	mts := &MailTests{}
	if err := c.SearchRead(MailTestModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTestId finds record id by querying it with criteria.
func (c *Client) FindMailTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no mail.test was found with criteria %v and options %v", criteria, options)
}
