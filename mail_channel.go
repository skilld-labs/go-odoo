package odoo

import (
	"fmt"
)

// MailChannel represents mail.channel model.
type MailChannel struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AliasContact              *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults             *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain               *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId        *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                   *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId              *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                 *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId        *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId       *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId               *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	AnonymousName             *String    `xmlrpc:"anonymous_name,omptempty"`
	ChannelLastSeenPartnerIds *Relation  `xmlrpc:"channel_last_seen_partner_ids,omptempty"`
	ChannelMessageIds         *Relation  `xmlrpc:"channel_message_ids,omptempty"`
	ChannelPartnerIds         *Relation  `xmlrpc:"channel_partner_ids,omptempty"`
	ChannelType               *Selection `xmlrpc:"channel_type,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description               *String    `xmlrpc:"description,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	EmailSend                 *Bool      `xmlrpc:"email_send,omptempty"`
	GroupIds                  *Relation  `xmlrpc:"group_ids,omptempty"`
	GroupPublicId             *Many2One  `xmlrpc:"group_public_id,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	Image                     *String    `xmlrpc:"image,omptempty"`
	ImageMedium               *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                *String    `xmlrpc:"image_small,omptempty"`
	IsMember                  *Bool      `xmlrpc:"is_member,omptempty"`
	IsSubscribed              *Bool      `xmlrpc:"is_subscribed,omptempty"`
	LivechatChannelId         *Many2One  `xmlrpc:"livechat_channel_id,omptempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost           *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	Public                    *Selection `xmlrpc:"public,omptempty"`
	RatingCount               *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                 *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback        *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage           *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastValue           *Float     `xmlrpc:"rating_last_value,omptempty"`
	Uuid                      *String    `xmlrpc:"uuid,omptempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailChannels represents array of mail.channel model.
type MailChannels []MailChannel

// MailChannelModel is the odoo model name.
const MailChannelModel = "mail.channel"

// Many2One convert MailChannel to *Many2One.
func (mc *MailChannel) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMailChannel creates a new mail.channel model and returns its id.
func (c *Client) CreateMailChannel(mc *MailChannel) (int64, error) {
	ids, err := c.CreateMailChannels([]*MailChannel{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailChannel creates a new mail.channel model and returns its id.
func (c *Client) CreateMailChannels(mcs []*MailChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MailChannelModel, vv)
}

// UpdateMailChannel updates an existing mail.channel record.
func (c *Client) UpdateMailChannel(mc *MailChannel) error {
	return c.UpdateMailChannels([]int64{mc.Id.Get()}, mc)
}

// UpdateMailChannels updates existing mail.channel records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMailChannels(ids []int64, mc *MailChannel) error {
	return c.Update(MailChannelModel, ids, mc)
}

// DeleteMailChannel deletes an existing mail.channel record.
func (c *Client) DeleteMailChannel(id int64) error {
	return c.DeleteMailChannels([]int64{id})
}

// DeleteMailChannels deletes existing mail.channel records.
func (c *Client) DeleteMailChannels(ids []int64) error {
	return c.Delete(MailChannelModel, ids)
}

// GetMailChannel gets mail.channel existing record.
func (c *Client) GetMailChannel(id int64) (*MailChannel, error) {
	mcs, err := c.GetMailChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.channel not found", id)
}

// GetMailChannels gets mail.channel existing records.
func (c *Client) GetMailChannels(ids []int64) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.Read(MailChannelModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannel finds mail.channel record by querying it with criteria.
func (c *Client) FindMailChannel(criteria *Criteria) (*MailChannel, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("mail.channel was not found with criteria %v", criteria)
}

// FindMailChannels finds mail.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannels(criteria *Criteria, options *Options) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailChannelId finds record id by querying it with criteria.
func (c *Client) FindMailChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.channel was not found with criteria %v and options %v", criteria, options)
}
