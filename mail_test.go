package odoo

import (
	"fmt"
)

// MailTest represents mail.test model.
type MailTest struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
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
