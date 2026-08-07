package odoo

// ImLivechatReportChannel represents im_livechat.report.channel model.
type ImLivechatReportChannel struct {
	AgentProvidingHelpHistory  *Many2One  `xmlrpc:"agent_providing_help_history,omitempty"`
	AgentRequestingHelpHistory *Many2One  `xmlrpc:"agent_requesting_help_history,omitempty"`
	CallDurationHour           *Float     `xmlrpc:"call_duration_hour,omitempty"`
	ChannelId                  *Many2One  `xmlrpc:"channel_id,omitempty"`
	ChannelName                *String    `xmlrpc:"channel_name,omitempty"`
	ChatbotAnswersPath         *String    `xmlrpc:"chatbot_answers_path,omitempty"`
	ChatbotAnswersPathStr      *String    `xmlrpc:"chatbot_answers_path_str,omitempty"`
	ChatbotScriptId            *Many2One  `xmlrpc:"chatbot_script_id,omitempty"`
	ConversationTagIds         *Relation  `xmlrpc:"conversation_tag_ids,omitempty"`
	CountryId                  *Many2One  `xmlrpc:"country_id,omitempty"`
	DayNumber                  *Selection `xmlrpc:"day_number,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Duration                   *Float     `xmlrpc:"duration,omitempty"`
	HandledByAgent             *Int       `xmlrpc:"handled_by_agent,omitempty"`
	HandledByBot               *Int       `xmlrpc:"handled_by_bot,omitempty"`
	HasCall                    *Float     `xmlrpc:"has_call,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	LangId                     *Many2One  `xmlrpc:"lang_id,omitempty"`
	LeadsCreated               *Int       `xmlrpc:"leads_created,omitempty"`
	LivechatChannelId          *Many2One  `xmlrpc:"livechat_channel_id,omitempty"`
	NbrMessage                 *Int       `xmlrpc:"nbr_message,omitempty"`
	NumberOfCalls              *Float     `xmlrpc:"number_of_calls,omitempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omitempty"`
	PercentageOfCalls          *Float     `xmlrpc:"percentage_of_calls,omitempty"`
	Rating                     *Int       `xmlrpc:"rating,omitempty"`
	RatingText                 *String    `xmlrpc:"rating_text,omitempty"`
	SessionExpertiseIds        *Relation  `xmlrpc:"session_expertise_ids,omitempty"`
	SessionExpertises          *String    `xmlrpc:"session_expertises,omitempty"`
	SessionOutcome             *Selection `xmlrpc:"session_outcome,omitempty"`
	StartDate                  *Time      `xmlrpc:"start_date,omitempty"`
	StartDateHour              *String    `xmlrpc:"start_date_hour,omitempty"`
	StartDateMinutes           *String    `xmlrpc:"start_date_minutes,omitempty"`
	StartHour                  *String    `xmlrpc:"start_hour,omitempty"`
	TimeToAnswer               *Float     `xmlrpc:"time_to_answer,omitempty"`
	Uuid                       *String    `xmlrpc:"uuid,omitempty"`
	VisitorPartnerId           *Many2One  `xmlrpc:"visitor_partner_id,omitempty"`
}

// ImLivechatReportChannels represents array of im_livechat.report.channel model.
type ImLivechatReportChannels []ImLivechatReportChannel

// ImLivechatReportChannelModel is the odoo model name.
const ImLivechatReportChannelModel = "im_livechat.report.channel"

// Many2One convert ImLivechatReportChannel to *Many2One.
func (irc *ImLivechatReportChannel) Many2One() *Many2One {
	return NewMany2One(irc.Id.Get(), "")
}

// CreateImLivechatReportChannel creates a new im_livechat.report.channel model and returns its id.
func (c *Client) CreateImLivechatReportChannel(irc *ImLivechatReportChannel) (int64, error) {
	ids, err := c.CreateImLivechatReportChannels([]*ImLivechatReportChannel{irc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatReportChannel creates a new im_livechat.report.channel model and returns its id.
func (c *Client) CreateImLivechatReportChannels(ircs []*ImLivechatReportChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range ircs {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatReportChannelModel, vv, nil)
}

// UpdateImLivechatReportChannel updates an existing im_livechat.report.channel record.
func (c *Client) UpdateImLivechatReportChannel(irc *ImLivechatReportChannel) error {
	return c.UpdateImLivechatReportChannels([]int64{irc.Id.Get()}, irc)
}

// UpdateImLivechatReportChannels updates existing im_livechat.report.channel records.
// All records (represented by ids) will be updated by irc values.
func (c *Client) UpdateImLivechatReportChannels(ids []int64, irc *ImLivechatReportChannel) error {
	return c.Update(ImLivechatReportChannelModel, ids, irc, nil)
}

// DeleteImLivechatReportChannel deletes an existing im_livechat.report.channel record.
func (c *Client) DeleteImLivechatReportChannel(id int64) error {
	return c.DeleteImLivechatReportChannels([]int64{id})
}

// DeleteImLivechatReportChannels deletes existing im_livechat.report.channel records.
func (c *Client) DeleteImLivechatReportChannels(ids []int64) error {
	return c.Delete(ImLivechatReportChannelModel, ids)
}

// GetImLivechatReportChannel gets im_livechat.report.channel existing record.
func (c *Client) GetImLivechatReportChannel(id int64) (*ImLivechatReportChannel, error) {
	ircs, err := c.GetImLivechatReportChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ircs)[0]), nil
}

// GetImLivechatReportChannels gets im_livechat.report.channel existing records.
func (c *Client) GetImLivechatReportChannels(ids []int64) (*ImLivechatReportChannels, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.Read(ImLivechatReportChannelModel, ids, nil, ircs); err != nil {
		return nil, err
	}
	return ircs, nil
}

// FindImLivechatReportChannel finds im_livechat.report.channel record by querying it with criteria.
func (c *Client) FindImLivechatReportChannel(criteria *Criteria) (*ImLivechatReportChannel, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.SearchRead(ImLivechatReportChannelModel, criteria, NewOptions().Limit(1), ircs); err != nil {
		return nil, err
	}
	return &((*ircs)[0]), nil
}

// FindImLivechatReportChannels finds im_livechat.report.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportChannels(criteria *Criteria, options *Options) (*ImLivechatReportChannels, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.SearchRead(ImLivechatReportChannelModel, criteria, options, ircs); err != nil {
		return nil, err
	}
	return ircs, nil
}

// FindImLivechatReportChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatReportChannelModel, criteria, options)
}

// FindImLivechatReportChannelId finds record id by querying it with criteria.
func (c *Client) FindImLivechatReportChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatReportChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
