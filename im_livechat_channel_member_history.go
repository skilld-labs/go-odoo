package odoo

// ImLivechatChannelMemberHistory represents im_livechat.channel.member.history model.
type ImLivechatChannelMemberHistory struct {
	AgentExpertiseIds        *Relation  `xmlrpc:"agent_expertise_ids,omitempty"`
	Avatar128                *String    `xmlrpc:"avatar_128,omitempty"`
	CallCount                *Float     `xmlrpc:"call_count,omitempty"`
	CallDurationHour         *Float     `xmlrpc:"call_duration_hour,omitempty"`
	CallHistoryIds           *Relation  `xmlrpc:"call_history_ids,omitempty"`
	CallPercentage           *Float     `xmlrpc:"call_percentage,omitempty"`
	ChannelId                *Many2One  `xmlrpc:"channel_id,omitempty"`
	ChatbotScriptId          *Many2One  `xmlrpc:"chatbot_script_id,omitempty"`
	ConversationTagIds       *Relation  `xmlrpc:"conversation_tag_ids,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	GuestId                  *Many2One  `xmlrpc:"guest_id,omitempty"`
	HasCall                  *Float     `xmlrpc:"has_call,omitempty"`
	HelpStatus               *Selection `xmlrpc:"help_status,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	LivechatMemberType       *Selection `xmlrpc:"livechat_member_type,omitempty"`
	MemberId                 *Many2One  `xmlrpc:"member_id,omitempty"`
	MessageCount             *Int       `xmlrpc:"message_count,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	Rating                   *Float     `xmlrpc:"rating,omitempty"`
	RatingId                 *Many2One  `xmlrpc:"rating_id,omitempty"`
	RatingText               *Selection `xmlrpc:"rating_text,omitempty"`
	ResponseTimeHour         *Float     `xmlrpc:"response_time_hour,omitempty"`
	SessionCountryId         *Many2One  `xmlrpc:"session_country_id,omitempty"`
	SessionDurationHour      *Float     `xmlrpc:"session_duration_hour,omitempty"`
	SessionLivechatChannelId *Many2One  `xmlrpc:"session_livechat_channel_id,omitempty"`
	SessionOutcome           *Selection `xmlrpc:"session_outcome,omitempty"`
	SessionStartHour         *Float     `xmlrpc:"session_start_hour,omitempty"`
	SessionWeekDay           *Selection `xmlrpc:"session_week_day,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ImLivechatChannelMemberHistorys represents array of im_livechat.channel.member.history model.
type ImLivechatChannelMemberHistorys []ImLivechatChannelMemberHistory

// ImLivechatChannelMemberHistoryModel is the odoo model name.
const ImLivechatChannelMemberHistoryModel = "im_livechat.channel.member.history"

// Many2One convert ImLivechatChannelMemberHistory to *Many2One.
func (icmh *ImLivechatChannelMemberHistory) Many2One() *Many2One {
	return NewMany2One(icmh.Id.Get(), "")
}

// CreateImLivechatChannelMemberHistory creates a new im_livechat.channel.member.history model and returns its id.
func (c *Client) CreateImLivechatChannelMemberHistory(icmh *ImLivechatChannelMemberHistory) (int64, error) {
	ids, err := c.CreateImLivechatChannelMemberHistorys([]*ImLivechatChannelMemberHistory{icmh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatChannelMemberHistory creates a new im_livechat.channel.member.history model and returns its id.
func (c *Client) CreateImLivechatChannelMemberHistorys(icmhs []*ImLivechatChannelMemberHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range icmhs {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatChannelMemberHistoryModel, vv, nil)
}

// UpdateImLivechatChannelMemberHistory updates an existing im_livechat.channel.member.history record.
func (c *Client) UpdateImLivechatChannelMemberHistory(icmh *ImLivechatChannelMemberHistory) error {
	return c.UpdateImLivechatChannelMemberHistorys([]int64{icmh.Id.Get()}, icmh)
}

// UpdateImLivechatChannelMemberHistorys updates existing im_livechat.channel.member.history records.
// All records (represented by ids) will be updated by icmh values.
func (c *Client) UpdateImLivechatChannelMemberHistorys(ids []int64, icmh *ImLivechatChannelMemberHistory) error {
	return c.Update(ImLivechatChannelMemberHistoryModel, ids, icmh, nil)
}

// DeleteImLivechatChannelMemberHistory deletes an existing im_livechat.channel.member.history record.
func (c *Client) DeleteImLivechatChannelMemberHistory(id int64) error {
	return c.DeleteImLivechatChannelMemberHistorys([]int64{id})
}

// DeleteImLivechatChannelMemberHistorys deletes existing im_livechat.channel.member.history records.
func (c *Client) DeleteImLivechatChannelMemberHistorys(ids []int64) error {
	return c.Delete(ImLivechatChannelMemberHistoryModel, ids)
}

// GetImLivechatChannelMemberHistory gets im_livechat.channel.member.history existing record.
func (c *Client) GetImLivechatChannelMemberHistory(id int64) (*ImLivechatChannelMemberHistory, error) {
	icmhs, err := c.GetImLivechatChannelMemberHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*icmhs)[0]), nil
}

// GetImLivechatChannelMemberHistorys gets im_livechat.channel.member.history existing records.
func (c *Client) GetImLivechatChannelMemberHistorys(ids []int64) (*ImLivechatChannelMemberHistorys, error) {
	icmhs := &ImLivechatChannelMemberHistorys{}
	if err := c.Read(ImLivechatChannelMemberHistoryModel, ids, nil, icmhs); err != nil {
		return nil, err
	}
	return icmhs, nil
}

// FindImLivechatChannelMemberHistory finds im_livechat.channel.member.history record by querying it with criteria.
func (c *Client) FindImLivechatChannelMemberHistory(criteria *Criteria) (*ImLivechatChannelMemberHistory, error) {
	icmhs := &ImLivechatChannelMemberHistorys{}
	if err := c.SearchRead(ImLivechatChannelMemberHistoryModel, criteria, NewOptions().Limit(1), icmhs); err != nil {
		return nil, err
	}
	return &((*icmhs)[0]), nil
}

// FindImLivechatChannelMemberHistorys finds im_livechat.channel.member.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelMemberHistorys(criteria *Criteria, options *Options) (*ImLivechatChannelMemberHistorys, error) {
	icmhs := &ImLivechatChannelMemberHistorys{}
	if err := c.SearchRead(ImLivechatChannelMemberHistoryModel, criteria, options, icmhs); err != nil {
		return nil, err
	}
	return icmhs, nil
}

// FindImLivechatChannelMemberHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelMemberHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatChannelMemberHistoryModel, criteria, options)
}

// FindImLivechatChannelMemberHistoryId finds record id by querying it with criteria.
func (c *Client) FindImLivechatChannelMemberHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatChannelMemberHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
