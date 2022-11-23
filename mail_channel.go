package odoo

import (
	"fmt"
)

// MailChannel represents mail.channel model.
type MailChannel struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	AliasBouncedContent       *String    `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact              *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults             *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain               *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId        *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                   *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId              *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                 *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId        *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId       *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId               *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	ChannelLastSeenPartnerIds *Relation  `xmlrpc:"channel_last_seen_partner_ids,omitempty"`
	ChannelMessageIds         *Relation  `xmlrpc:"channel_message_ids,omitempty"`
	ChannelPartnerIds         *Relation  `xmlrpc:"channel_partner_ids,omitempty"`
	ChannelType               *Selection `xmlrpc:"channel_type,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description               *String    `xmlrpc:"description,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	EmailSend                 *Bool      `xmlrpc:"email_send,omitempty"`
	GroupIds                  *Relation  `xmlrpc:"group_ids,omitempty"`
	GroupPublicId             *Many2One  `xmlrpc:"group_public_id,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	Image128                  *String    `xmlrpc:"image_128,omitempty"`
	IsChat                    *Bool      `xmlrpc:"is_chat,omitempty"`
	IsMember                  *Bool      `xmlrpc:"is_member,omitempty"`
	IsModerator               *Bool      `xmlrpc:"is_moderator,omitempty"`
	IsSubscribed              *Bool      `xmlrpc:"is_subscribed,omitempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Moderation                *Bool      `xmlrpc:"moderation,omitempty"`
	ModerationCount           *Int       `xmlrpc:"moderation_count,omitempty"`
	ModerationGuidelines      *Bool      `xmlrpc:"moderation_guidelines,omitempty"`
	ModerationGuidelinesMsg   *String    `xmlrpc:"moderation_guidelines_msg,omitempty"`
	ModerationIds             *Relation  `xmlrpc:"moderation_ids,omitempty"`
	ModerationNotify          *Bool      `xmlrpc:"moderation_notify,omitempty"`
	ModerationNotifyMsg       *String    `xmlrpc:"moderation_notify_msg,omitempty"`
	ModeratorIds              *Relation  `xmlrpc:"moderator_ids,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	Public                    *Selection `xmlrpc:"public,omitempty"`
	Uuid                      *String    `xmlrpc:"uuid,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(MailChannelModel, mc)
}

// UpdateMailChannel updates an existing mail.channel record.
func (c *Client) UpdateMailChannel(mc *MailChannel) error {
	return c.UpdateMailChannels([]int64{mc.Id.Get()}, mc)
}

// UpdateMailChannels updates existing mail.channel records.
// All records (represented by IDs) will be updated by mc values.
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
	return nil, fmt.Errorf("mail.channel was not found")
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

// FindMailChannelIds finds records IDs by querying it
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
	return -1, fmt.Errorf("mail.channel was not found")
}
