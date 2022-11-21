package odoo

import (
	"fmt"
)

// PortalMixin represents portal.mixin model.
type PortalMixin struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
	PortalUrl   *String `xmlrpc:"portal_url,omptempty"`
}

// PortalMixins represents array of portal.mixin model.
type PortalMixins []PortalMixin

// PortalMixinModel is the odoo model name.
const PortalMixinModel = "portal.mixin"

// Many2One convert PortalMixin to *Many2One.
func (pm *PortalMixin) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreatePortalMixin creates a new portal.mixin model and returns its id.
func (c *Client) CreatePortalMixin(pm *PortalMixin) (int64, error) {
	return c.Create(PortalMixinModel, pm)
}

// UpdatePortalMixin updates an existing portal.mixin record.
func (c *Client) UpdatePortalMixin(pm *PortalMixin) error {
	return c.UpdatePortalMixins([]int64{pm.Id.Get()}, pm)
}

// UpdatePortalMixins updates existing portal.mixin records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdatePortalMixins(ids []int64, pm *PortalMixin) error {
	return c.Update(PortalMixinModel, ids, pm)
}

// DeletePortalMixin deletes an existing portal.mixin record.
func (c *Client) DeletePortalMixin(id int64) error {
	return c.DeletePortalMixins([]int64{id})
}

// DeletePortalMixins deletes existing portal.mixin records.
func (c *Client) DeletePortalMixins(ids []int64) error {
	return c.Delete(PortalMixinModel, ids)
}

// GetPortalMixin gets portal.mixin existing record.
func (c *Client) GetPortalMixin(id int64) (*PortalMixin, error) {
	pms, err := c.GetPortalMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of portal.mixin not found", id)
}

// GetPortalMixins gets portal.mixin existing records.
func (c *Client) GetPortalMixins(ids []int64) (*PortalMixins, error) {
	pms := &PortalMixins{}
	if err := c.Read(PortalMixinModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPortalMixin finds portal.mixin record by querying it with criteria.
func (c *Client) FindPortalMixin(criteria *Criteria) (*PortalMixin, error) {
	pms := &PortalMixins{}
	if err := c.SearchRead(PortalMixinModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("no portal.mixin was found with criteria %v", criteria)
}

// FindPortalMixins finds portal.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalMixins(criteria *Criteria, options *Options) (*PortalMixins, error) {
	pms := &PortalMixins{}
	if err := c.SearchRead(PortalMixinModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPortalMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PortalMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPortalMixinId finds record id by querying it with criteria.
func (c *Client) FindPortalMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PortalMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no portal.mixin was found with criteria %v and options %v", criteria, options)
}
