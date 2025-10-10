package odoo

// DiscussChannel represents discuss.channel model.
type DiscussChannel struct {
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	AllowPublicUpload            *Bool      `xmlrpc:"allow_public_upload,omitempty"`
	AnonymousName                *String    `xmlrpc:"anonymous_name,omitempty"`
	Avatar128                    *String    `xmlrpc:"avatar_128,omitempty"`
	AvatarCacheKey               *String    `xmlrpc:"avatar_cache_key,omitempty"`
	ChannelMemberIds             *Relation  `xmlrpc:"channel_member_ids,omitempty"`
	ChannelPartnerIds            *Relation  `xmlrpc:"channel_partner_ids,omitempty"`
	ChannelType                  *Selection `xmlrpc:"channel_type,omitempty"`
	ChatbotCurrentStepId         *Many2One  `xmlrpc:"chatbot_current_step_id,omitempty"`
	ChatbotMessageIds            *Relation  `xmlrpc:"chatbot_message_ids,omitempty"`
	CountryId                    *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultDisplayMode           *Selection `xmlrpc:"default_display_mode,omitempty"`
	Description                  *String    `xmlrpc:"description,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	Duration                     *Float     `xmlrpc:"duration,omitempty"`
	FromMessageId                *Many2One  `xmlrpc:"from_message_id,omitempty"`
	GroupIds                     *Relation  `xmlrpc:"group_ids,omitempty"`
	GroupPublicId                *Many2One  `xmlrpc:"group_public_id,omitempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	Image128                     *String    `xmlrpc:"image_128,omitempty"`
	InvitationUrl                *String    `xmlrpc:"invitation_url,omitempty"`
	IsEditable                   *Bool      `xmlrpc:"is_editable,omitempty"`
	IsMember                     *Bool      `xmlrpc:"is_member,omitempty"`
	LastInterestDt               *Time      `xmlrpc:"last_interest_dt,omitempty"`
	LivechatActive               *Bool      `xmlrpc:"livechat_active,omitempty"`
	LivechatChannelId            *Many2One  `xmlrpc:"livechat_channel_id,omitempty"`
	LivechatOperatorId           *Many2One  `xmlrpc:"livechat_operator_id,omitempty"`
	MemberCount                  *Int       `xmlrpc:"member_count,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	ParentChannelId              *Many2One  `xmlrpc:"parent_channel_id,omitempty"`
	PinnedMessageIds             *Relation  `xmlrpc:"pinned_message_ids,omitempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omitempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omitempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omitempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	RtcSessionIds                *Relation  `xmlrpc:"rtc_session_ids,omitempty"`
	SfuChannelUuid               *String    `xmlrpc:"sfu_channel_uuid,omitempty"`
	SfuServerUrl                 *String    `xmlrpc:"sfu_server_url,omitempty"`
	SubChannelIds                *Relation  `xmlrpc:"sub_channel_ids,omitempty"`
	SubscriptionDepartmentIds    *Relation  `xmlrpc:"subscription_department_ids,omitempty"`
	Uuid                         *String    `xmlrpc:"uuid,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DiscussChannels represents array of discuss.channel model.
type DiscussChannels []DiscussChannel

// DiscussChannelModel is the odoo model name.
const DiscussChannelModel = "discuss.channel"

// Many2One convert DiscussChannel to *Many2One.
func (dc *DiscussChannel) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannel(dc *DiscussChannel) (int64, error) {
	ids, err := c.CreateDiscussChannels([]*DiscussChannel{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannels(dcs []*DiscussChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DiscussChannelModel, vv, nil)
}

// UpdateDiscussChannel updates an existing discuss.channel record.
func (c *Client) UpdateDiscussChannel(dc *DiscussChannel) error {
	return c.UpdateDiscussChannels([]int64{dc.Id.Get()}, dc)
}

// UpdateDiscussChannels updates existing discuss.channel records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDiscussChannels(ids []int64, dc *DiscussChannel) error {
	return c.Update(DiscussChannelModel, ids, dc, nil)
}

// DeleteDiscussChannel deletes an existing discuss.channel record.
func (c *Client) DeleteDiscussChannel(id int64) error {
	return c.DeleteDiscussChannels([]int64{id})
}

// DeleteDiscussChannels deletes existing discuss.channel records.
func (c *Client) DeleteDiscussChannels(ids []int64) error {
	return c.Delete(DiscussChannelModel, ids)
}

// GetDiscussChannel gets discuss.channel existing record.
func (c *Client) GetDiscussChannel(id int64) (*DiscussChannel, error) {
	dcs, err := c.GetDiscussChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// GetDiscussChannels gets discuss.channel existing records.
func (c *Client) GetDiscussChannels(ids []int64) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.Read(DiscussChannelModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannel finds discuss.channel record by querying it with criteria.
func (c *Client) FindDiscussChannel(criteria *Criteria) (*DiscussChannel, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// FindDiscussChannels finds discuss.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannels(criteria *Criteria, options *Options) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussChannelModel, criteria, options)
}

// FindDiscussChannelId finds record id by querying it with criteria.
func (c *Client) FindDiscussChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
