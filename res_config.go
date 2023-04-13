package odoo

import (
	"fmt"
)

// ResConfig represents res.config model.
type ResConfig struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResConfigs represents array of res.config model.
type ResConfigs []ResConfig

// ResConfigModel is the odoo model name.
const ResConfigModel = "res.config"

// Many2One convert ResConfig to *Many2One.
func (rc *ResConfig) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResConfig creates a new res.config model and returns its id.
func (c *Client) CreateResConfig(rc *ResConfig) (int64, error) {
	ids, err := c.CreateResConfigs([]*ResConfig{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfig creates a new res.config model and returns its id.
func (c *Client) CreateResConfigs(rcs []*ResConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResConfigModel, vv)
}

// UpdateResConfig updates an existing res.config record.
func (c *Client) UpdateResConfig(rc *ResConfig) error {
	return c.UpdateResConfigs([]int64{rc.Id.Get()}, rc)
}

// UpdateResConfigs updates existing res.config records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResConfigs(ids []int64, rc *ResConfig) error {
	return c.Update(ResConfigModel, ids, rc)
}

// DeleteResConfig deletes an existing res.config record.
func (c *Client) DeleteResConfig(id int64) error {
	return c.DeleteResConfigs([]int64{id})
}

// DeleteResConfigs deletes existing res.config records.
func (c *Client) DeleteResConfigs(ids []int64) error {
	return c.Delete(ResConfigModel, ids)
}

// GetResConfig gets res.config existing record.
func (c *Client) GetResConfig(id int64) (*ResConfig, error) {
	rcs, err := c.GetResConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config not found", id)
}

// GetResConfigs gets res.config existing records.
func (c *Client) GetResConfigs(ids []int64) (*ResConfigs, error) {
	rcs := &ResConfigs{}
	if err := c.Read(ResConfigModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResConfig finds res.config record by querying it with criteria.
func (c *Client) FindResConfig(criteria *Criteria) (*ResConfig, error) {
	rcs := &ResConfigs{}
	if err := c.SearchRead(ResConfigModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.config was not found with criteria %v", criteria)
}

// FindResConfigs finds res.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigs(criteria *Criteria, options *Options) (*ResConfigs, error) {
	rcs := &ResConfigs{}
	if err := c.SearchRead(ResConfigModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigId finds record id by querying it with criteria.
func (c *Client) FindResConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config was not found with criteria %v and options %v", criteria, options)
}
