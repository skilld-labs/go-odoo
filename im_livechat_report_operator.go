package odoo

import (
	"fmt"
)

// ImLivechatReportOperator represents im_livechat.report.operator model.
type ImLivechatReportOperator struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	ChannelId         *Many2One `xmlrpc:"channel_id,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Duration          *Float    `xmlrpc:"duration,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LivechatChannelId *Many2One `xmlrpc:"livechat_channel_id,omptempty"`
	NbrChannel        *Int      `xmlrpc:"nbr_channel,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	StartDate         *Time     `xmlrpc:"start_date,omptempty"`
	TimeToAnswer      *Float    `xmlrpc:"time_to_answer,omptempty"`
}

// ImLivechatReportOperators represents array of im_livechat.report.operator model.
type ImLivechatReportOperators []ImLivechatReportOperator

// ImLivechatReportOperatorModel is the odoo model name.
const ImLivechatReportOperatorModel = "im_livechat.report.operator"

// Many2One convert ImLivechatReportOperator to *Many2One.
func (iro *ImLivechatReportOperator) Many2One() *Many2One {
	return NewMany2One(iro.Id.Get(), "")
}

// CreateImLivechatReportOperator creates a new im_livechat.report.operator model and returns its id.
func (c *Client) CreateImLivechatReportOperator(iro *ImLivechatReportOperator) (int64, error) {
	ids, err := c.CreateImLivechatReportOperators([]*ImLivechatReportOperator{iro})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatReportOperator creates a new im_livechat.report.operator model and returns its id.
func (c *Client) CreateImLivechatReportOperators(iros []*ImLivechatReportOperator) ([]int64, error) {
	var vv []interface{}
	for _, v := range iros {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatReportOperatorModel, vv)
}

// UpdateImLivechatReportOperator updates an existing im_livechat.report.operator record.
func (c *Client) UpdateImLivechatReportOperator(iro *ImLivechatReportOperator) error {
	return c.UpdateImLivechatReportOperators([]int64{iro.Id.Get()}, iro)
}

// UpdateImLivechatReportOperators updates existing im_livechat.report.operator records.
// All records (represented by ids) will be updated by iro values.
func (c *Client) UpdateImLivechatReportOperators(ids []int64, iro *ImLivechatReportOperator) error {
	return c.Update(ImLivechatReportOperatorModel, ids, iro)
}

// DeleteImLivechatReportOperator deletes an existing im_livechat.report.operator record.
func (c *Client) DeleteImLivechatReportOperator(id int64) error {
	return c.DeleteImLivechatReportOperators([]int64{id})
}

// DeleteImLivechatReportOperators deletes existing im_livechat.report.operator records.
func (c *Client) DeleteImLivechatReportOperators(ids []int64) error {
	return c.Delete(ImLivechatReportOperatorModel, ids)
}

// GetImLivechatReportOperator gets im_livechat.report.operator existing record.
func (c *Client) GetImLivechatReportOperator(id int64) (*ImLivechatReportOperator, error) {
	iros, err := c.GetImLivechatReportOperators([]int64{id})
	if err != nil {
		return nil, err
	}
	if iros != nil && len(*iros) > 0 {
		return &((*iros)[0]), nil
	}
	return nil, fmt.Errorf("id %v of im_livechat.report.operator not found", id)
}

// GetImLivechatReportOperators gets im_livechat.report.operator existing records.
func (c *Client) GetImLivechatReportOperators(ids []int64) (*ImLivechatReportOperators, error) {
	iros := &ImLivechatReportOperators{}
	if err := c.Read(ImLivechatReportOperatorModel, ids, nil, iros); err != nil {
		return nil, err
	}
	return iros, nil
}

// FindImLivechatReportOperator finds im_livechat.report.operator record by querying it with criteria.
func (c *Client) FindImLivechatReportOperator(criteria *Criteria) (*ImLivechatReportOperator, error) {
	iros := &ImLivechatReportOperators{}
	if err := c.SearchRead(ImLivechatReportOperatorModel, criteria, NewOptions().Limit(1), iros); err != nil {
		return nil, err
	}
	if iros != nil && len(*iros) > 0 {
		return &((*iros)[0]), nil
	}
	return nil, fmt.Errorf("im_livechat.report.operator was not found with criteria %v", criteria)
}

// FindImLivechatReportOperators finds im_livechat.report.operator records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportOperators(criteria *Criteria, options *Options) (*ImLivechatReportOperators, error) {
	iros := &ImLivechatReportOperators{}
	if err := c.SearchRead(ImLivechatReportOperatorModel, criteria, options, iros); err != nil {
		return nil, err
	}
	return iros, nil
}

// FindImLivechatReportOperatorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportOperatorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ImLivechatReportOperatorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindImLivechatReportOperatorId finds record id by querying it with criteria.
func (c *Client) FindImLivechatReportOperatorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatReportOperatorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("im_livechat.report.operator was not found with criteria %v and options %v", criteria, options)
}
