package odoo

import (
	"fmt"
)

// UtmSource represents utm.source model.
type UtmSource struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UtmSources represents array of utm.source model.
type UtmSources []UtmSource

// UtmSourceModel is the odoo model name.
const UtmSourceModel = "utm.source"

// Many2One convert UtmSource to *Many2One.
func (us *UtmSource) Many2One() *Many2One {
	return NewMany2One(us.Id.Get(), "")
}

// CreateUtmSource creates a new utm.source model and returns its id.
func (c *Client) CreateUtmSource(us *UtmSource) (int64, error) {
	return c.Create(UtmSourceModel, us)
}

// UpdateUtmSource updates an existing utm.source record.
func (c *Client) UpdateUtmSource(us *UtmSource) error {
	return c.UpdateUtmSources([]int64{us.Id.Get()}, us)
}

// UpdateUtmSources updates existing utm.source records.
// All records (represented by ids) will be updated by us values.
func (c *Client) UpdateUtmSources(ids []int64, us *UtmSource) error {
	return c.Update(UtmSourceModel, ids, us)
}

// DeleteUtmSource deletes an existing utm.source record.
func (c *Client) DeleteUtmSource(id int64) error {
	return c.DeleteUtmSources([]int64{id})
}

// DeleteUtmSources deletes existing utm.source records.
func (c *Client) DeleteUtmSources(ids []int64) error {
	return c.Delete(UtmSourceModel, ids)
}

// GetUtmSource gets utm.source existing record.
func (c *Client) GetUtmSource(id int64) (*UtmSource, error) {
	uss, err := c.GetUtmSources([]int64{id})
	if err != nil {
		return nil, err
	}
	if uss != nil && len(*uss) > 0 {
		return &((*uss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of utm.source not found", id)
}

// GetUtmSources gets utm.source existing records.
func (c *Client) GetUtmSources(ids []int64) (*UtmSources, error) {
	uss := &UtmSources{}
	if err := c.Read(UtmSourceModel, ids, nil, uss); err != nil {
		return nil, err
	}
	return uss, nil
}

// FindUtmSource finds utm.source record by querying it with criteria.
func (c *Client) FindUtmSource(criteria *Criteria) (*UtmSource, error) {
	uss := &UtmSources{}
	if err := c.SearchRead(UtmSourceModel, criteria, NewOptions().Limit(1), uss); err != nil {
		return nil, err
	}
	if uss != nil && len(*uss) > 0 {
		return &((*uss)[0]), nil
	}
	return nil, fmt.Errorf("no utm.source was found with criteria %v", criteria)
}

// FindUtmSources finds utm.source records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmSources(criteria *Criteria, options *Options) (*UtmSources, error) {
	uss := &UtmSources{}
	if err := c.SearchRead(UtmSourceModel, criteria, options, uss); err != nil {
		return nil, err
	}
	return uss, nil
}

// FindUtmSourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmSourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UtmSourceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUtmSourceId finds record id by querying it with criteria.
func (c *Client) FindUtmSourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmSourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no utm.source was found with criteria %v and options %v", criteria, options)
}
