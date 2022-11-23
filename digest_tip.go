package odoo

import (
	"fmt"
)

// DigestTip represents digest.tip model.
type DigestTip struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	GroupId        *Many2One `xmlrpc:"group_id,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	TipDescription *String   `xmlrpc:"tip_description,omitempty"`
	UserIds        *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DigestTips represents array of digest.tip model.
type DigestTips []DigestTip

// DigestTipModel is the odoo model name.
const DigestTipModel = "digest.tip"

// Many2One convert DigestTip to *Many2One.
func (dt *DigestTip) Many2One() *Many2One {
	return NewMany2One(dt.Id.Get(), "")
}

// CreateDigestTip creates a new digest.tip model and returns its id.
func (c *Client) CreateDigestTip(dt *DigestTip) (int64, error) {
	return c.Create(DigestTipModel, dt)
}

// UpdateDigestTip updates an existing digest.tip record.
func (c *Client) UpdateDigestTip(dt *DigestTip) error {
	return c.UpdateDigestTips([]int64{dt.Id.Get()}, dt)
}

// UpdateDigestTips updates existing digest.tip records.
// All records (represented by IDs) will be updated by dt values.
func (c *Client) UpdateDigestTips(ids []int64, dt *DigestTip) error {
	return c.Update(DigestTipModel, ids, dt)
}

// DeleteDigestTip deletes an existing digest.tip record.
func (c *Client) DeleteDigestTip(id int64) error {
	return c.DeleteDigestTips([]int64{id})
}

// DeleteDigestTips deletes existing digest.tip records.
func (c *Client) DeleteDigestTips(ids []int64) error {
	return c.Delete(DigestTipModel, ids)
}

// GetDigestTip gets digest.tip existing record.
func (c *Client) GetDigestTip(id int64) (*DigestTip, error) {
	dts, err := c.GetDigestTips([]int64{id})
	if err != nil {
		return nil, err
	}
	if dts != nil && len(*dts) > 0 {
		return &((*dts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of digest.tip not found", id)
}

// GetDigestTips gets digest.tip existing records.
func (c *Client) GetDigestTips(ids []int64) (*DigestTips, error) {
	dts := &DigestTips{}
	if err := c.Read(DigestTipModel, ids, nil, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDigestTip finds digest.tip record by querying it with criteria.
func (c *Client) FindDigestTip(criteria *Criteria) (*DigestTip, error) {
	dts := &DigestTips{}
	if err := c.SearchRead(DigestTipModel, criteria, NewOptions().Limit(1), dts); err != nil {
		return nil, err
	}
	if dts != nil && len(*dts) > 0 {
		return &((*dts)[0]), nil
	}
	return nil, fmt.Errorf("digest.tip was not found")
}

// FindDigestTips finds digest.tip records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestTips(criteria *Criteria, options *Options) (*DigestTips, error) {
	dts := &DigestTips{}
	if err := c.SearchRead(DigestTipModel, criteria, options, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDigestTipIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestTipIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DigestTipModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDigestTipId finds record id by querying it with criteria.
func (c *Client) FindDigestTipId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DigestTipModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("digest.tip was not found")
}
