package odoo

import (
	"fmt"
)

// UtmMixin represents utm.mixin model.
type UtmMixin struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CampaignId  *Many2One `xmlrpc:"campaign_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MediumId    *Many2One `xmlrpc:"medium_id,omptempty"`
	SourceId    *Many2One `xmlrpc:"source_id,omptempty"`
}

// UtmMixins represents array of utm.mixin model.
type UtmMixins []UtmMixin

// UtmMixinModel is the odoo model name.
const UtmMixinModel = "utm.mixin"

// Many2One convert UtmMixin to *Many2One.
func (um *UtmMixin) Many2One() *Many2One {
	return NewMany2One(um.Id.Get(), "")
}

// CreateUtmMixin creates a new utm.mixin model and returns its id.
func (c *Client) CreateUtmMixin(um *UtmMixin) (int64, error) {
	return c.Create(UtmMixinModel, um)
}

// UpdateUtmMixin updates an existing utm.mixin record.
func (c *Client) UpdateUtmMixin(um *UtmMixin) error {
	return c.UpdateUtmMixins([]int64{um.Id.Get()}, um)
}

// UpdateUtmMixins updates existing utm.mixin records.
// All records (represented by ids) will be updated by um values.
func (c *Client) UpdateUtmMixins(ids []int64, um *UtmMixin) error {
	return c.Update(UtmMixinModel, ids, um)
}

// DeleteUtmMixin deletes an existing utm.mixin record.
func (c *Client) DeleteUtmMixin(id int64) error {
	return c.DeleteUtmMixins([]int64{id})
}

// DeleteUtmMixins deletes existing utm.mixin records.
func (c *Client) DeleteUtmMixins(ids []int64) error {
	return c.Delete(UtmMixinModel, ids)
}

// GetUtmMixin gets utm.mixin existing record.
func (c *Client) GetUtmMixin(id int64) (*UtmMixin, error) {
	ums, err := c.GetUtmMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if ums != nil && len(*ums) > 0 {
		return &((*ums)[0]), nil
	}
	return nil, fmt.Errorf("id %v of utm.mixin not found", id)
}

// GetUtmMixins gets utm.mixin existing records.
func (c *Client) GetUtmMixins(ids []int64) (*UtmMixins, error) {
	ums := &UtmMixins{}
	if err := c.Read(UtmMixinModel, ids, nil, ums); err != nil {
		return nil, err
	}
	return ums, nil
}

// FindUtmMixin finds utm.mixin record by querying it with criteria.
func (c *Client) FindUtmMixin(criteria *Criteria) (*UtmMixin, error) {
	ums := &UtmMixins{}
	if err := c.SearchRead(UtmMixinModel, criteria, NewOptions().Limit(1), ums); err != nil {
		return nil, err
	}
	if ums != nil && len(*ums) > 0 {
		return &((*ums)[0]), nil
	}
	return nil, fmt.Errorf("utm.mixin was not found with criteria %v", criteria)
}

// FindUtmMixins finds utm.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmMixins(criteria *Criteria, options *Options) (*UtmMixins, error) {
	ums := &UtmMixins{}
	if err := c.SearchRead(UtmMixinModel, criteria, options, ums); err != nil {
		return nil, err
	}
	return ums, nil
}

// FindUtmMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UtmMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUtmMixinId finds record id by querying it with criteria.
func (c *Client) FindUtmMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("utm.mixin was not found with criteria %v and options %v", criteria, options)
}
