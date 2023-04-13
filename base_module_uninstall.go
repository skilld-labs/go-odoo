package odoo

import (
	"fmt"
)

// BaseModuleUninstall represents base.module.uninstall model.
type BaseModuleUninstall struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModelIds    *Relation `xmlrpc:"model_ids,omptempty"`
	ModuleId    *Many2One `xmlrpc:"module_id,omptempty"`
	ModuleIds   *Relation `xmlrpc:"module_ids,omptempty"`
	ShowAll     *Bool     `xmlrpc:"show_all,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseModuleUninstalls represents array of base.module.uninstall model.
type BaseModuleUninstalls []BaseModuleUninstall

// BaseModuleUninstallModel is the odoo model name.
const BaseModuleUninstallModel = "base.module.uninstall"

// Many2One convert BaseModuleUninstall to *Many2One.
func (bmu *BaseModuleUninstall) Many2One() *Many2One {
	return NewMany2One(bmu.Id.Get(), "")
}

// CreateBaseModuleUninstall creates a new base.module.uninstall model and returns its id.
func (c *Client) CreateBaseModuleUninstall(bmu *BaseModuleUninstall) (int64, error) {
	ids, err := c.Create(BaseModuleUninstallModel, []interface{}{bmu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseModuleUninstall creates a new base.module.uninstall model and returns its id.
func (c *Client) CreateBaseModuleUninstalls(bmus []*BaseModuleUninstall) ([]int64, error) {
	var vv []interface{}
	for _, v := range bmus {
		vv = append(vv, v)
	}
	return c.Create(BaseModuleUninstallModel, vv)
}

// UpdateBaseModuleUninstall updates an existing base.module.uninstall record.
func (c *Client) UpdateBaseModuleUninstall(bmu *BaseModuleUninstall) error {
	return c.UpdateBaseModuleUninstalls([]int64{bmu.Id.Get()}, bmu)
}

// UpdateBaseModuleUninstalls updates existing base.module.uninstall records.
// All records (represented by ids) will be updated by bmu values.
func (c *Client) UpdateBaseModuleUninstalls(ids []int64, bmu *BaseModuleUninstall) error {
	return c.Update(BaseModuleUninstallModel, ids, bmu)
}

// DeleteBaseModuleUninstall deletes an existing base.module.uninstall record.
func (c *Client) DeleteBaseModuleUninstall(id int64) error {
	return c.DeleteBaseModuleUninstalls([]int64{id})
}

// DeleteBaseModuleUninstalls deletes existing base.module.uninstall records.
func (c *Client) DeleteBaseModuleUninstalls(ids []int64) error {
	return c.Delete(BaseModuleUninstallModel, ids)
}

// GetBaseModuleUninstall gets base.module.uninstall existing record.
func (c *Client) GetBaseModuleUninstall(id int64) (*BaseModuleUninstall, error) {
	bmus, err := c.GetBaseModuleUninstalls([]int64{id})
	if err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.module.uninstall not found", id)
}

// GetBaseModuleUninstalls gets base.module.uninstall existing records.
func (c *Client) GetBaseModuleUninstalls(ids []int64) (*BaseModuleUninstalls, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.Read(BaseModuleUninstallModel, ids, nil, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUninstall finds base.module.uninstall record by querying it with criteria.
func (c *Client) FindBaseModuleUninstall(criteria *Criteria) (*BaseModuleUninstall, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.SearchRead(BaseModuleUninstallModel, criteria, NewOptions().Limit(1), bmus); err != nil {
		return nil, err
	}
	if bmus != nil && len(*bmus) > 0 {
		return &((*bmus)[0]), nil
	}
	return nil, fmt.Errorf("base.module.uninstall was not found with criteria %v", criteria)
}

// FindBaseModuleUninstalls finds base.module.uninstall records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUninstalls(criteria *Criteria, options *Options) (*BaseModuleUninstalls, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.SearchRead(BaseModuleUninstallModel, criteria, options, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUninstallIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUninstallIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModuleUninstallModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseModuleUninstallId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleUninstallId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleUninstallModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.module.uninstall was not found with criteria %v and options %v", criteria, options)
}
