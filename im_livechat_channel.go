package odoo

// ImLivechatChannel represents im_livechat.channel model.
type ImLivechatChannel struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	AreYouInside                 *Bool     `xmlrpc:"are_you_inside,omptempty"`
	ButtonText                   *String   `xmlrpc:"button_text,omptempty"`
	ChannelIds                   *Relation `xmlrpc:"channel_ids,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	DefaultMessage               *String   `xmlrpc:"default_message,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	Image                        *String   `xmlrpc:"image,omptempty"`
	ImageMedium                  *String   `xmlrpc:"image_medium,omptempty"`
	ImageSmall                   *String   `xmlrpc:"image_small,omptempty"`
	InputPlaceholder             *String   `xmlrpc:"input_placeholder,omptempty"`
	Name                         *String   `xmlrpc:"name,omptempty"`
	NbrChannel                   *Int      `xmlrpc:"nbr_channel,omptempty"`
	RatingPercentageSatisfaction *Int      `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RuleIds                      *Relation `xmlrpc:"rule_ids,omptempty"`
	ScriptExternal               *String   `xmlrpc:"script_external,omptempty"`
	UserIds                      *Relation `xmlrpc:"user_ids,omptempty"`
	WebPage                      *String   `xmlrpc:"web_page,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ImLivechatChannels represents array of im_livechat.channel model.
type ImLivechatChannels []ImLivechatChannel

// ImLivechatChannelModel is the odoo model name.
const ImLivechatChannelModel = "im_livechat.channel"

// Many2One convert ImLivechatChannel to *Many2One.
func (ic *ImLivechatChannel) Many2One() *Many2One {
	return NewMany2One(ic.Id.Get(), "")
}

// CreateImLivechatChannel creates a new im_livechat.channel model and returns its id.
func (c *Client) CreateImLivechatChannel(ic *ImLivechatChannel) (int64, error) {
	ids, err := c.CreateImLivechatChannels([]*ImLivechatChannel{ic})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatChannel creates a new im_livechat.channel model and returns its id.
func (c *Client) CreateImLivechatChannels(ics []*ImLivechatChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range ics {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatChannelModel, vv, nil)
}

// UpdateImLivechatChannel updates an existing im_livechat.channel record.
func (c *Client) UpdateImLivechatChannel(ic *ImLivechatChannel) error {
	return c.UpdateImLivechatChannels([]int64{ic.Id.Get()}, ic)
}

// UpdateImLivechatChannels updates existing im_livechat.channel records.
// All records (represented by ids) will be updated by ic values.
func (c *Client) UpdateImLivechatChannels(ids []int64, ic *ImLivechatChannel) error {
	return c.Update(ImLivechatChannelModel, ids, ic, nil)
}

// DeleteImLivechatChannel deletes an existing im_livechat.channel record.
func (c *Client) DeleteImLivechatChannel(id int64) error {
	return c.DeleteImLivechatChannels([]int64{id})
}

// DeleteImLivechatChannels deletes existing im_livechat.channel records.
func (c *Client) DeleteImLivechatChannels(ids []int64) error {
	return c.Delete(ImLivechatChannelModel, ids)
}

// GetImLivechatChannel gets im_livechat.channel existing record.
func (c *Client) GetImLivechatChannel(id int64) (*ImLivechatChannel, error) {
	ics, err := c.GetImLivechatChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// GetImLivechatChannels gets im_livechat.channel existing records.
func (c *Client) GetImLivechatChannels(ids []int64) (*ImLivechatChannels, error) {
	ics := &ImLivechatChannels{}
	if err := c.Read(ImLivechatChannelModel, ids, nil, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindImLivechatChannel finds im_livechat.channel record by querying it with criteria.
func (c *Client) FindImLivechatChannel(criteria *Criteria) (*ImLivechatChannel, error) {
	ics := &ImLivechatChannels{}
	if err := c.SearchRead(ImLivechatChannelModel, criteria, NewOptions().Limit(1), ics); err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// FindImLivechatChannels finds im_livechat.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannels(criteria *Criteria, options *Options) (*ImLivechatChannels, error) {
	ics := &ImLivechatChannels{}
	if err := c.SearchRead(ImLivechatChannelModel, criteria, options, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindImLivechatChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatChannelModel, criteria, options)
}

// FindImLivechatChannelId finds record id by querying it with criteria.
func (c *Client) FindImLivechatChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
