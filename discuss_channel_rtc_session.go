package odoo

// DiscussChannelRtcSession represents discuss.channel.rtc.session model.
type DiscussChannelRtcSession struct {
	ChannelId         *Many2One `xmlrpc:"channel_id,omitempty"`
	ChannelMemberId   *Many2One `xmlrpc:"channel_member_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	GuestId           *Many2One `xmlrpc:"guest_id,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	IsCameraOn        *Bool     `xmlrpc:"is_camera_on,omitempty"`
	IsDeaf            *Bool     `xmlrpc:"is_deaf,omitempty"`
	IsMuted           *Bool     `xmlrpc:"is_muted,omitempty"`
	IsScreenSharingOn *Bool     `xmlrpc:"is_screen_sharing_on,omitempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DiscussChannelRtcSessions represents array of discuss.channel.rtc.session model.
type DiscussChannelRtcSessions []DiscussChannelRtcSession

// DiscussChannelRtcSessionModel is the odoo model name.
const DiscussChannelRtcSessionModel = "discuss.channel.rtc.session"

// Many2One convert DiscussChannelRtcSession to *Many2One.
func (dcrs *DiscussChannelRtcSession) Many2One() *Many2One {
	return NewMany2One(dcrs.Id.Get(), "")
}

// CreateDiscussChannelRtcSession creates a new discuss.channel.rtc.session model and returns its id.
func (c *Client) CreateDiscussChannelRtcSession(dcrs *DiscussChannelRtcSession) (int64, error) {
	ids, err := c.CreateDiscussChannelRtcSessions([]*DiscussChannelRtcSession{dcrs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussChannelRtcSession creates a new discuss.channel.rtc.session model and returns its id.
func (c *Client) CreateDiscussChannelRtcSessions(dcrss []*DiscussChannelRtcSession) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcrss {
		vv = append(vv, v)
	}
	return c.Create(DiscussChannelRtcSessionModel, vv, nil)
}

// UpdateDiscussChannelRtcSession updates an existing discuss.channel.rtc.session record.
func (c *Client) UpdateDiscussChannelRtcSession(dcrs *DiscussChannelRtcSession) error {
	return c.UpdateDiscussChannelRtcSessions([]int64{dcrs.Id.Get()}, dcrs)
}

// UpdateDiscussChannelRtcSessions updates existing discuss.channel.rtc.session records.
// All records (represented by ids) will be updated by dcrs values.
func (c *Client) UpdateDiscussChannelRtcSessions(ids []int64, dcrs *DiscussChannelRtcSession) error {
	return c.Update(DiscussChannelRtcSessionModel, ids, dcrs, nil)
}

// DeleteDiscussChannelRtcSession deletes an existing discuss.channel.rtc.session record.
func (c *Client) DeleteDiscussChannelRtcSession(id int64) error {
	return c.DeleteDiscussChannelRtcSessions([]int64{id})
}

// DeleteDiscussChannelRtcSessions deletes existing discuss.channel.rtc.session records.
func (c *Client) DeleteDiscussChannelRtcSessions(ids []int64) error {
	return c.Delete(DiscussChannelRtcSessionModel, ids)
}

// GetDiscussChannelRtcSession gets discuss.channel.rtc.session existing record.
func (c *Client) GetDiscussChannelRtcSession(id int64) (*DiscussChannelRtcSession, error) {
	dcrss, err := c.GetDiscussChannelRtcSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dcrss)[0]), nil
}

// GetDiscussChannelRtcSessions gets discuss.channel.rtc.session existing records.
func (c *Client) GetDiscussChannelRtcSessions(ids []int64) (*DiscussChannelRtcSessions, error) {
	dcrss := &DiscussChannelRtcSessions{}
	if err := c.Read(DiscussChannelRtcSessionModel, ids, nil, dcrss); err != nil {
		return nil, err
	}
	return dcrss, nil
}

// FindDiscussChannelRtcSession finds discuss.channel.rtc.session record by querying it with criteria.
func (c *Client) FindDiscussChannelRtcSession(criteria *Criteria) (*DiscussChannelRtcSession, error) {
	dcrss := &DiscussChannelRtcSessions{}
	if err := c.SearchRead(DiscussChannelRtcSessionModel, criteria, NewOptions().Limit(1), dcrss); err != nil {
		return nil, err
	}
	return &((*dcrss)[0]), nil
}

// FindDiscussChannelRtcSessions finds discuss.channel.rtc.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelRtcSessions(criteria *Criteria, options *Options) (*DiscussChannelRtcSessions, error) {
	dcrss := &DiscussChannelRtcSessions{}
	if err := c.SearchRead(DiscussChannelRtcSessionModel, criteria, options, dcrss); err != nil {
		return nil, err
	}
	return dcrss, nil
}

// FindDiscussChannelRtcSessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelRtcSessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussChannelRtcSessionModel, criteria, options)
}

// FindDiscussChannelRtcSessionId finds record id by querying it with criteria.
func (c *Client) FindDiscussChannelRtcSessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussChannelRtcSessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
