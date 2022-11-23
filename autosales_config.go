package odoo

import (
	"fmt"
)

// AutosalesConfig represents autosales.config model.
type AutosalesConfig struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ConfigLine  *Relation `xmlrpc:"config_line,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AutosalesConfigs represents array of autosales.config model.
type AutosalesConfigs []AutosalesConfig

// AutosalesConfigModel is the odoo model name.
const AutosalesConfigModel = "autosales.config"

// Many2One convert AutosalesConfig to *Many2One.
func (ac *AutosalesConfig) Many2One() *Many2One {
	return NewMany2One(ac.Id.Get(), "")
}

// CreateAutosalesConfig creates a new autosales.config model and returns its id.
func (c *Client) CreateAutosalesConfig(ac *AutosalesConfig) (int64, error) {
	return c.Create(AutosalesConfigModel, ac)
}

// UpdateAutosalesConfig updates an existing autosales.config record.
func (c *Client) UpdateAutosalesConfig(ac *AutosalesConfig) error {
	return c.UpdateAutosalesConfigs([]int64{ac.Id.Get()}, ac)
}

// UpdateAutosalesConfigs updates existing autosales.config records.
// All records (represented by ids) will be updated by ac values.
func (c *Client) UpdateAutosalesConfigs(ids []int64, ac *AutosalesConfig) error {
	return c.Update(AutosalesConfigModel, ids, ac)
}

// DeleteAutosalesConfig deletes an existing autosales.config record.
func (c *Client) DeleteAutosalesConfig(id int64) error {
	return c.DeleteAutosalesConfigs([]int64{id})
}

// DeleteAutosalesConfigs deletes existing autosales.config records.
func (c *Client) DeleteAutosalesConfigs(ids []int64) error {
	return c.Delete(AutosalesConfigModel, ids)
}

// GetAutosalesConfig gets autosales.config existing record.
func (c *Client) GetAutosalesConfig(id int64) (*AutosalesConfig, error) {
	acs, err := c.GetAutosalesConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if acs != nil && len(*acs) > 0 {
		return &((*acs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of autosales.config not found", id)
}

// GetAutosalesConfigs gets autosales.config existing records.
func (c *Client) GetAutosalesConfigs(ids []int64) (*AutosalesConfigs, error) {
	acs := &AutosalesConfigs{}
	if err := c.Read(AutosalesConfigModel, ids, nil, acs); err != nil {
		return nil, err
	}
	return acs, nil
}

// FindAutosalesConfig finds autosales.config record by querying it with criteria.
func (c *Client) FindAutosalesConfig(criteria *Criteria) (*AutosalesConfig, error) {
	acs := &AutosalesConfigs{}
	if err := c.SearchRead(AutosalesConfigModel, criteria, NewOptions().Limit(1), acs); err != nil {
		return nil, err
	}
	if acs != nil && len(*acs) > 0 {
		return &((*acs)[0]), nil
	}
	return nil, fmt.Errorf("autosales.config was not found with criteria %v", criteria)
}

// FindAutosalesConfigs finds autosales.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAutosalesConfigs(criteria *Criteria, options *Options) (*AutosalesConfigs, error) {
	acs := &AutosalesConfigs{}
	if err := c.SearchRead(AutosalesConfigModel, criteria, options, acs); err != nil {
		return nil, err
	}
	return acs, nil
}

// FindAutosalesConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAutosalesConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AutosalesConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAutosalesConfigId finds record id by querying it with criteria.
func (c *Client) FindAutosalesConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AutosalesConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("autosales.config was not found with criteria %v and options %v", criteria, options)
}
