package odoo

// DiscussChannelMember represents discuss.channel.member model.
type DiscussChannelMember struct {
	ChannelId            *Many2One  `xmlrpc:"channel_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomChannelName    *String    `xmlrpc:"custom_channel_name,omitempty"`
	CustomNotifications  *Selection `xmlrpc:"custom_notifications,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	FetchedMessageId     *Many2One  `xmlrpc:"fetched_message_id,omitempty"`
	FoldState            *Selection `xmlrpc:"fold_state,omitempty"`
	GuestId              *Many2One  `xmlrpc:"guest_id,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	IsPinned             *Bool      `xmlrpc:"is_pinned,omitempty"`
	IsSelf               *Bool      `xmlrpc:"is_self,omitempty"`
	LastInterestDt       *Time      `xmlrpc:"last_interest_dt,omitempty"`
	LastSeenDt           *Time      `xmlrpc:"last_seen_dt,omitempty"`
	MessageUnreadCounter *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MuteUntilDt          *Time      `xmlrpc:"mute_until_dt,omitempty"`
	NewMessageSeparator  *Int       `xmlrpc:"new_message_separator,omitempty"`
	PartnerId            *Many2One  `xmlrpc:"partner_id,omitempty"`
	RtcInvitingSessionId *Many2One  `xmlrpc:"rtc_inviting_session_id,omitempty"`
	RtcSessionIds        *Relation  `xmlrpc:"rtc_session_ids,omitempty"`
	SeenMessageId        *Many2One  `xmlrpc:"seen_message_id,omitempty"`
	UnpinDt              *Time      `xmlrpc:"unpin_dt,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DiscussChannelMembers represents array of discuss.channel.member model.
type DiscussChannelMembers []DiscussChannelMember

// DiscussChannelMemberModel is the odoo model name.
const DiscussChannelMemberModel = "discuss.channel.member"

// Many2One convert DiscussChannelMember to *Many2One.
func (dcm *DiscussChannelMember) Many2One() *Many2One {
	return NewMany2One(dcm.Id.Get(), "")
}

// CreateDiscussChannelMember creates a new discuss.channel.member model and returns its id.
func (c *Client) CreateDiscussChannelMember(dcm *DiscussChannelMember) (int64, error) {
	ids, err := c.CreateDiscussChannelMembers([]*DiscussChannelMember{dcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussChannelMember creates a new discuss.channel.member model and returns its id.
func (c *Client) CreateDiscussChannelMembers(dcms []*DiscussChannelMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcms {
		vv = append(vv, v)
	}
	return c.Create(DiscussChannelMemberModel, vv, nil)
}

// UpdateDiscussChannelMember updates an existing discuss.channel.member record.
func (c *Client) UpdateDiscussChannelMember(dcm *DiscussChannelMember) error {
	return c.UpdateDiscussChannelMembers([]int64{dcm.Id.Get()}, dcm)
}

// UpdateDiscussChannelMembers updates existing discuss.channel.member records.
// All records (represented by ids) will be updated by dcm values.
func (c *Client) UpdateDiscussChannelMembers(ids []int64, dcm *DiscussChannelMember) error {
	return c.Update(DiscussChannelMemberModel, ids, dcm, nil)
}

// DeleteDiscussChannelMember deletes an existing discuss.channel.member record.
func (c *Client) DeleteDiscussChannelMember(id int64) error {
	return c.DeleteDiscussChannelMembers([]int64{id})
}

// DeleteDiscussChannelMembers deletes existing discuss.channel.member records.
func (c *Client) DeleteDiscussChannelMembers(ids []int64) error {
	return c.Delete(DiscussChannelMemberModel, ids)
}

// GetDiscussChannelMember gets discuss.channel.member existing record.
func (c *Client) GetDiscussChannelMember(id int64) (*DiscussChannelMember, error) {
	dcms, err := c.GetDiscussChannelMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dcms)[0]), nil
}

// GetDiscussChannelMembers gets discuss.channel.member existing records.
func (c *Client) GetDiscussChannelMembers(ids []int64) (*DiscussChannelMembers, error) {
	dcms := &DiscussChannelMembers{}
	if err := c.Read(DiscussChannelMemberModel, ids, nil, dcms); err != nil {
		return nil, err
	}
	return dcms, nil
}

// FindDiscussChannelMember finds discuss.channel.member record by querying it with criteria.
func (c *Client) FindDiscussChannelMember(criteria *Criteria) (*DiscussChannelMember, error) {
	dcms := &DiscussChannelMembers{}
	if err := c.SearchRead(DiscussChannelMemberModel, criteria, NewOptions().Limit(1), dcms); err != nil {
		return nil, err
	}
	return &((*dcms)[0]), nil
}

// FindDiscussChannelMembers finds discuss.channel.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelMembers(criteria *Criteria, options *Options) (*DiscussChannelMembers, error) {
	dcms := &DiscussChannelMembers{}
	if err := c.SearchRead(DiscussChannelMemberModel, criteria, options, dcms); err != nil {
		return nil, err
	}
	return dcms, nil
}

// FindDiscussChannelMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussChannelMemberModel, criteria, options)
}

// FindDiscussChannelMemberId finds record id by querying it with criteria.
func (c *Client) FindDiscussChannelMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussChannelMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
