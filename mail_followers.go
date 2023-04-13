package odoo

import (
	"fmt"
)

// MailFollowers represents mail.followers model.
type MailFollowers struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ChannelId   *Many2One `xmlrpc:"channel_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omptempty"`
	ResId       *Int      `xmlrpc:"res_id,omptempty"`
	ResModel    *String   `xmlrpc:"res_model,omptempty"`
	SubtypeIds  *Relation `xmlrpc:"subtype_ids,omptempty"`
}

// MailFollowerss represents array of mail.followers model.
type MailFollowerss []MailFollowers

// MailFollowersModel is the odoo model name.
const MailFollowersModel = "mail.followers"

// Many2One convert MailFollowers to *Many2One.
func (mf *MailFollowers) Many2One() *Many2One {
	return NewMany2One(mf.Id.Get(), "")
}

// CreateMailFollowers creates a new mail.followers model and returns its id.
func (c *Client) CreateMailFollowers(mf *MailFollowers) (int64, error) {
	ids, err := c.Create(MailFollowersModel, []interface{}{mf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailFollowers creates a new mail.followers model and returns its id.
func (c *Client) CreateMailFollowerss(mfs []*MailFollowers) ([]int64, error) {
	var vv []interface{}
	for _, v := range mfs {
		vv = append(vv, v)
	}
	return c.Create(MailFollowersModel, vv)
}

// UpdateMailFollowers updates an existing mail.followers record.
func (c *Client) UpdateMailFollowers(mf *MailFollowers) error {
	return c.UpdateMailFollowerss([]int64{mf.Id.Get()}, mf)
}

// UpdateMailFollowerss updates existing mail.followers records.
// All records (represented by ids) will be updated by mf values.
func (c *Client) UpdateMailFollowerss(ids []int64, mf *MailFollowers) error {
	return c.Update(MailFollowersModel, ids, mf)
}

// DeleteMailFollowers deletes an existing mail.followers record.
func (c *Client) DeleteMailFollowers(id int64) error {
	return c.DeleteMailFollowerss([]int64{id})
}

// DeleteMailFollowerss deletes existing mail.followers records.
func (c *Client) DeleteMailFollowerss(ids []int64) error {
	return c.Delete(MailFollowersModel, ids)
}

// GetMailFollowers gets mail.followers existing record.
func (c *Client) GetMailFollowers(id int64) (*MailFollowers, error) {
	mfs, err := c.GetMailFollowerss([]int64{id})
	if err != nil {
		return nil, err
	}
	if mfs != nil && len(*mfs) > 0 {
		return &((*mfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.followers not found", id)
}

// GetMailFollowerss gets mail.followers existing records.
func (c *Client) GetMailFollowerss(ids []int64) (*MailFollowerss, error) {
	mfs := &MailFollowerss{}
	if err := c.Read(MailFollowersModel, ids, nil, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailFollowers finds mail.followers record by querying it with criteria.
func (c *Client) FindMailFollowers(criteria *Criteria) (*MailFollowers, error) {
	mfs := &MailFollowerss{}
	if err := c.SearchRead(MailFollowersModel, criteria, NewOptions().Limit(1), mfs); err != nil {
		return nil, err
	}
	if mfs != nil && len(*mfs) > 0 {
		return &((*mfs)[0]), nil
	}
	return nil, fmt.Errorf("mail.followers was not found with criteria %v", criteria)
}

// FindMailFollowerss finds mail.followers records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowerss(criteria *Criteria, options *Options) (*MailFollowerss, error) {
	mfs := &MailFollowerss{}
	if err := c.SearchRead(MailFollowersModel, criteria, options, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailFollowersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailFollowersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailFollowersId finds record id by querying it with criteria.
func (c *Client) FindMailFollowersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailFollowersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.followers was not found with criteria %v and options %v", criteria, options)
}
