package odoo

import (
	"fmt"
)

// BaseModuleUpgrade represents base.module.upgrade model.
type BaseModuleUpgrade struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModuleInfo  *String   `xmlrpc:"module_info,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseModuleUpgrades represents array of base.module.upgrade model.
type BaseModuleUpgrades []BaseModuleUpgrade

// BaseModuleUpgradeModel is the odoo model name.
const BaseModuleUpgradeModel = "base.module.upgrade"

// Many2One convert BaseModuleUpgrade to *Many2One.
func (bmu *BaseModuleUpgrade) Many2One() *Many2One {
	return NewMany2One(bmu.Id.Get(), "")
}

// CreateBaseModuleUpgrade creates a new base.module.upgrade model and returns its id.
func (c *Client) CreateBaseModuleUpgrade(bmu *BaseModuleUpgrade) (int64, error) {
	ids, err := c.Create(BaseModuleUpgradeModel, []interface{}{bmu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseModuleUpgrade creates a new base.module.upgrade model and returns its id.
func (c *Client) CreateBaseModuleUpgrades(bmus []*BaseModuleUpgrade) ([]int64, error) {
	var vv []interface{}
	for _, v := range bmus {
		vv = append(vv, v)
	}
	return c.Create(BaseModuleUpgradeModel, vv)
}

// UpdateBaseModuleUpgrade updates an existing base.module.upgrade record.
func (c *Client) UpdateBaseModuleUpgrade(bmu *BaseModuleUpgrade) error {
	return c.UpdateBaseModuleUpgrades([]int64{bmu.Id.Get()}, bmu)
}

// UpdateBaseModuleUpgrades updates existing base.module.upgrade records.
// All records (represented by ids) will be updated by bmu values.
func (c *Client) UpdateBaseModuleUpgrades(ids []int64, bmu *BaseModuleUpgrade) error {
	return c.Update(BaseModuleUpgradeModel, ids, bmu)
}

// DeleteBaseModuleUpgrade deletes an existing base.module.upgrade record.
func (c *Client) DeleteBaseModuleUpgrade(id int64) error {
	return c.DeleteBaseModuleUpgrades([]int64{id})
}

// DeleteBaseModuleUpgrades deletes existing base.module.upgrade records.
func (c *Client) DeleteBaseModuleUpgrades(ids []int64) error {
	return c.Delete(BaseModuleUpgradeModel, ids)
}

// GetBaseModuleUpgrade gets base.module.upgrade existing record.
func (c *Client) GetBaseModuleUpgrade(id int64) (*BaseModuleUpgrade, error) {
	bmus, err := c.GetBaseModuleUpgrades([]int64{id})
	if err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.module.upgrade not found", id)
}

// GetBaseModuleUpgrades gets base.module.upgrade existing records.
func (c *Client) GetBaseModuleUpgrades(ids []int64) (*BaseModuleUpgrades, error) {
	bmus := &BaseModuleUpgrades{}
	if err := c.Read(BaseModuleUpgradeModel, ids, nil, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUpgrade finds base.module.upgrade record by querying it with criteria.
func (c *Client) FindBaseModuleUpgrade(criteria *Criteria) (*BaseModuleUpgrade, error) {
	bmus := &BaseModuleUpgrades{}
	if err := c.SearchRead(BaseModuleUpgradeModel, criteria, NewOptions().Limit(1), bmus); err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("base.module.upgrade was not found with criteria %v", criteria)
}

// FindBaseModuleUpgrades finds base.module.upgrade records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUpgrades(criteria *Criteria, options *Options) (*BaseModuleUpgrades, error) {
	bmus := &BaseModuleUpgrades{}
	if err := c.SearchRead(BaseModuleUpgradeModel, criteria, options, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUpgradeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUpgradeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModuleUpgradeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseModuleUpgradeId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleUpgradeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleUpgradeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.module.upgrade was not found with criteria %v and options %v", criteria, options)
}
