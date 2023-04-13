package odoo

import (
	"fmt"
)

// ImLivechatChannelRule represents im_livechat.channel.rule model.
type ImLivechatChannelRule struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Action         *Selection `xmlrpc:"action,omptempty"`
	AutoPopupTimer *Int       `xmlrpc:"auto_popup_timer,omptempty"`
	ChannelId      *Many2One  `xmlrpc:"channel_id,omptempty"`
	CountryIds     *Relation  `xmlrpc:"country_ids,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	RegexUrl       *String    `xmlrpc:"regex_url,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(ImLivechatChannelRuleModel, []interface{}{icr})
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
	return c.Create(ImLivechatChannelRuleModel, vv)
}

// UpdateImLivechatChannelRule updates an existing im_livechat.channel.rule record.
func (c *Client) UpdateImLivechatChannelRule(icr *ImLivechatChannelRule) error {
	return c.UpdateImLivechatChannelRules([]int64{icr.Id.Get()}, icr)
}

// UpdateImLivechatChannelRules updates existing im_livechat.channel.rule records.
// All records (represented by ids) will be updated by icr values.
func (c *Client) UpdateImLivechatChannelRules(ids []int64, icr *ImLivechatChannelRule) error {
	return c.Update(ImLivechatChannelRuleModel, ids, icr)
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
	if icrs != nil && len(*icrs) > 0 {
		return &((*icrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of im_livechat.channel.rule not found", id)
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
	if icrs != nil && len(*icrs) > 0 {
		return &((*icrs)[0]), nil
	}
	return nil, fmt.Errorf("im_livechat.channel.rule was not found with criteria %v", criteria)
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
	ids, err := c.Search(ImLivechatChannelRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindImLivechatChannelRuleId finds record id by querying it with criteria.
func (c *Client) FindImLivechatChannelRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatChannelRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("im_livechat.channel.rule was not found with criteria %v and options %v", criteria, options)
}
