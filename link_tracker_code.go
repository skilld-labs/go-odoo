package odoo

import (
	"fmt"
)

// LinkTrackerCode represents link.tracker.code model.
type LinkTrackerCode struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Code        *String   `xmlrpc:"code,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LinkId      *Many2One `xmlrpc:"link_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LinkTrackerCodes represents array of link.tracker.code model.
type LinkTrackerCodes []LinkTrackerCode

// LinkTrackerCodeModel is the odoo model name.
const LinkTrackerCodeModel = "link.tracker.code"

// Many2One convert LinkTrackerCode to *Many2One.
func (ltc *LinkTrackerCode) Many2One() *Many2One {
	return NewMany2One(ltc.Id.Get(), "")
}

// CreateLinkTrackerCode creates a new link.tracker.code model and returns its id.
func (c *Client) CreateLinkTrackerCode(ltc *LinkTrackerCode) (int64, error) {
	ids, err := c.Create(LinkTrackerCodeModel, []interface{}{ltc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLinkTrackerCode creates a new link.tracker.code model and returns its id.
func (c *Client) CreateLinkTrackerCodes(ltcs []*LinkTrackerCode) ([]int64, error) {
	var vv []interface{}
	for _, v := range ltcs {
		vv = append(vv, v)
	}
	return c.Create(LinkTrackerCodeModel, vv)
}

// UpdateLinkTrackerCode updates an existing link.tracker.code record.
func (c *Client) UpdateLinkTrackerCode(ltc *LinkTrackerCode) error {
	return c.UpdateLinkTrackerCodes([]int64{ltc.Id.Get()}, ltc)
}

// UpdateLinkTrackerCodes updates existing link.tracker.code records.
// All records (represented by ids) will be updated by ltc values.
func (c *Client) UpdateLinkTrackerCodes(ids []int64, ltc *LinkTrackerCode) error {
	return c.Update(LinkTrackerCodeModel, ids, ltc)
}

// DeleteLinkTrackerCode deletes an existing link.tracker.code record.
func (c *Client) DeleteLinkTrackerCode(id int64) error {
	return c.DeleteLinkTrackerCodes([]int64{id})
}

// DeleteLinkTrackerCodes deletes existing link.tracker.code records.
func (c *Client) DeleteLinkTrackerCodes(ids []int64) error {
	return c.Delete(LinkTrackerCodeModel, ids)
}

// GetLinkTrackerCode gets link.tracker.code existing record.
func (c *Client) GetLinkTrackerCode(id int64) (*LinkTrackerCode, error) {
	ltcs, err := c.GetLinkTrackerCodes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ltcs != nil && len(*ltcs) > 0 {
		return &((*ltcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of link.tracker.code not found", id)
}

// GetLinkTrackerCodes gets link.tracker.code existing records.
func (c *Client) GetLinkTrackerCodes(ids []int64) (*LinkTrackerCodes, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.Read(LinkTrackerCodeModel, ids, nil, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerCode finds link.tracker.code record by querying it with criteria.
func (c *Client) FindLinkTrackerCode(criteria *Criteria) (*LinkTrackerCode, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.SearchRead(LinkTrackerCodeModel, criteria, NewOptions().Limit(1), ltcs); err != nil {
		return nil, err
	}
	if ltcs != nil && len(*ltcs) > 0 {
		return &((*ltcs)[0]), nil
	}
	return nil, fmt.Errorf("link.tracker.code was not found with criteria %v", criteria)
}

// FindLinkTrackerCodes finds link.tracker.code records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerCodes(criteria *Criteria, options *Options) (*LinkTrackerCodes, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.SearchRead(LinkTrackerCodeModel, criteria, options, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerCodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerCodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LinkTrackerCodeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLinkTrackerCodeId finds record id by querying it with criteria.
func (c *Client) FindLinkTrackerCodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LinkTrackerCodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("link.tracker.code was not found with criteria %v and options %v", criteria, options)
}
