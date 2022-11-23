package odoo

import (
	"fmt"
)

// BaseModuleUpdate represents base.module.update model.
type BaseModuleUpdate struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Added       *Int       `xmlrpc:"added,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	Updated     *Int       `xmlrpc:"updated,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseModuleUpdates represents array of base.module.update model.
type BaseModuleUpdates []BaseModuleUpdate

// BaseModuleUpdateModel is the odoo model name.
const BaseModuleUpdateModel = "base.module.update"

// Many2One convert BaseModuleUpdate to *Many2One.
func (bmu *BaseModuleUpdate) Many2One() *Many2One {
	return NewMany2One(bmu.Id.Get(), "")
}

// CreateBaseModuleUpdate creates a new base.module.update model and returns its id.
func (c *Client) CreateBaseModuleUpdate(bmu *BaseModuleUpdate) (int64, error) {
	return c.Create(BaseModuleUpdateModel, bmu)
}

// UpdateBaseModuleUpdate updates an existing base.module.update record.
func (c *Client) UpdateBaseModuleUpdate(bmu *BaseModuleUpdate) error {
	return c.UpdateBaseModuleUpdates([]int64{bmu.Id.Get()}, bmu)
}

// UpdateBaseModuleUpdates updates existing base.module.update records.
// All records (represented by ids) will be updated by bmu values.
func (c *Client) UpdateBaseModuleUpdates(ids []int64, bmu *BaseModuleUpdate) error {
	return c.Update(BaseModuleUpdateModel, ids, bmu)
}

// DeleteBaseModuleUpdate deletes an existing base.module.update record.
func (c *Client) DeleteBaseModuleUpdate(id int64) error {
	return c.DeleteBaseModuleUpdates([]int64{id})
}

// DeleteBaseModuleUpdates deletes existing base.module.update records.
func (c *Client) DeleteBaseModuleUpdates(ids []int64) error {
	return c.Delete(BaseModuleUpdateModel, ids)
}

// GetBaseModuleUpdate gets base.module.update existing record.
func (c *Client) GetBaseModuleUpdate(id int64) (*BaseModuleUpdate, error) {
	bmus, err := c.GetBaseModuleUpdates([]int64{id})
	if err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.module.update not found", id)
}

// GetBaseModuleUpdates gets base.module.update existing records.
func (c *Client) GetBaseModuleUpdates(ids []int64) (*BaseModuleUpdates, error) {
	bmus := &BaseModuleUpdates{}
	if err := c.Read(BaseModuleUpdateModel, ids, nil, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUpdate finds base.module.update record by querying it with criteria.
func (c *Client) FindBaseModuleUpdate(criteria *Criteria) (*BaseModuleUpdate, error) {
	bmus := &BaseModuleUpdates{}
	if err := c.SearchRead(BaseModuleUpdateModel, criteria, NewOptions().Limit(1), bmus); err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("base.module.update was not found with criteria %v", criteria)
}

// FindBaseModuleUpdates finds base.module.update records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUpdates(criteria *Criteria, options *Options) (*BaseModuleUpdates, error) {
	bmus := &BaseModuleUpdates{}
	if err := c.SearchRead(BaseModuleUpdateModel, criteria, options, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUpdateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUpdateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModuleUpdateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseModuleUpdateId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleUpdateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleUpdateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.module.update was not found with criteria %v and options %v", criteria, options)
}
