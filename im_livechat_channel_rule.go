package odoo

// ImLivechatChannelRule represents im_livechat.channel.rule model.
type ImLivechatChannelRule struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	Action         *Selection `xmlrpc:"action,omitempty"`
	AutoPopupTimer *Int       `xmlrpc:"auto_popup_timer,omitempty"`
	ChannelId      *Many2One  `xmlrpc:"channel_id,omitempty"`
	CountryIds     *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	RegexUrl       *String    `xmlrpc:"regex_url,omitempty"`
	Sequence       *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ImLivechatChannelRules represents array of im_livechat.channel.rule model.
type ImLivechatChannelRules []ImLivechatChannelRule

// ImLivechatChannelRuleModel is the odoo model name.
const ImLivechatChannelRuleModel = "im_livechat.channel.rule"

// Many2One convert ImLivechatChannelRule to *Many2One.
func (icr *ImLivechatChannelRule) Many2One() *Many2One {
	return NewMany2One(icr.Id.Get(), "")
}

// CreateImLivechatChannelRule creates a new im_livechat.channel.rule model and returns its id.
func (c *Client) CreateImLivechatChannelRule(icr *ImLivechatChannelRule) (int64, error) {
	ids, err := c.CreateImLivechatChannelRules([]*ImLivechatChannelRule{icr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatChannelRule creates a new im_livechat.channel.rule model and returns its id.
func (c *Client) CreateImLivechatChannelRules(icrs []*ImLivechatChannelRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range icrs {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatChannelRuleModel, vv, nil)
}

// UpdateImLivechatChannelRule updates an existing im_livechat.channel.rule record.
func (c *Client) UpdateImLivechatChannelRule(icr *ImLivechatChannelRule) error {
	return c.UpdateImLivechatChannelRules([]int64{icr.Id.Get()}, icr)
}

// UpdateImLivechatChannelRules updates existing im_livechat.channel.rule records.
// All records (represented by ids) will be updated by icr values.
func (c *Client) UpdateImLivechatChannelRules(ids []int64, icr *ImLivechatChannelRule) error {
	return c.Update(ImLivechatChannelRuleModel, ids, icr, nil)
}

// DeleteImLivechatChannelRule deletes an existing im_livechat.channel.rule record.
func (c *Client) DeleteImLivechatChannelRule(id int64) error {
	return c.DeleteImLivechatChannelRules([]int64{id})
}

// DeleteImLivechatChannelRules deletes existing im_livechat.channel.rule records.
func (c *Client) DeleteImLivechatChannelRules(ids []int64) error {
	return c.Delete(ImLivechatChannelRuleModel, ids)
}

// GetImLivechatChannelRule gets im_livechat.channel.rule existing record.
func (c *Client) GetImLivechatChannelRule(id int64) (*ImLivechatChannelRule, error) {
	icrs, err := c.GetImLivechatChannelRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*icrs)[0]), nil
}

// GetImLivechatChannelRules gets im_livechat.channel.rule existing records.
func (c *Client) GetImLivechatChannelRules(ids []int64) (*ImLivechatChannelRules, error) {
	icrs := &ImLivechatChannelRules{}
	if err := c.Read(ImLivechatChannelRuleModel, ids, nil, icrs); err != nil {
		return nil, err
	}
	return icrs, nil
}

// FindImLivechatChannelRule finds im_livechat.channel.rule record by querying it with criteria.
func (c *Client) FindImLivechatChannelRule(criteria *Criteria) (*ImLivechatChannelRule, error) {
	icrs := &ImLivechatChannelRules{}
	if err := c.SearchRead(ImLivechatChannelRuleModel, criteria, NewOptions().Limit(1), icrs); err != nil {
		return nil, err
	}
	return &((*icrs)[0]), nil
}

// FindImLivechatChannelRules finds im_livechat.channel.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelRules(criteria *Criteria, options *Options) (*ImLivechatChannelRules, error) {
	icrs := &ImLivechatChannelRules{}
	if err := c.SearchRead(ImLivechatChannelRuleModel, criteria, options, icrs); err != nil {
		return nil, err
	}
	return icrs, nil
}

// FindImLivechatChannelRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatChannelRuleModel, criteria, options)
}

// FindImLivechatChannelRuleId finds record id by querying it with criteria.
func (c *Client) FindImLivechatChannelRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatChannelRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
