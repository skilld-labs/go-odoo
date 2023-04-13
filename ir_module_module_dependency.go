package odoo

import (
	"fmt"
)

// IrModuleModuleDependency represents ir.module.module.dependency model.
type IrModuleModuleDependency struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DependId    *Many2One  `xmlrpc:"depend_id,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	ModuleId    *Many2One  `xmlrpc:"module_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrModuleModuleDependencys represents array of ir.module.module.dependency model.
type IrModuleModuleDependencys []IrModuleModuleDependency

// IrModuleModuleDependencyModel is the odoo model name.
const IrModuleModuleDependencyModel = "ir.module.module.dependency"

// Many2One convert IrModuleModuleDependency to *Many2One.
func (immd *IrModuleModuleDependency) Many2One() *Many2One {
	return NewMany2One(immd.Id.Get(), "")
}

// CreateIrModuleModuleDependency creates a new ir.module.module.dependency model and returns its id.
func (c *Client) CreateIrModuleModuleDependency(immd *IrModuleModuleDependency) (int64, error) {
	ids, err := c.Create(IrModuleModuleDependencyModel, []interface{}{immd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModuleModuleDependency creates a new ir.module.module.dependency model and returns its id.
func (c *Client) CreateIrModuleModuleDependencys(immds []*IrModuleModuleDependency) ([]int64, error) {
	var vv []interface{}
	for _, v := range immds {
		vv = append(vv, v)
	}
	return c.Create(IrModuleModuleDependencyModel, vv)
}

// UpdateIrModuleModuleDependency updates an existing ir.module.module.dependency record.
func (c *Client) UpdateIrModuleModuleDependency(immd *IrModuleModuleDependency) error {
	return c.UpdateIrModuleModuleDependencys([]int64{immd.Id.Get()}, immd)
}

// UpdateIrModuleModuleDependencys updates existing ir.module.module.dependency records.
// All records (represented by ids) will be updated by immd values.
func (c *Client) UpdateIrModuleModuleDependencys(ids []int64, immd *IrModuleModuleDependency) error {
	return c.Update(IrModuleModuleDependencyModel, ids, immd)
}

// DeleteIrModuleModuleDependency deletes an existing ir.module.module.dependency record.
func (c *Client) DeleteIrModuleModuleDependency(id int64) error {
	return c.DeleteIrModuleModuleDependencys([]int64{id})
}

// DeleteIrModuleModuleDependencys deletes existing ir.module.module.dependency records.
func (c *Client) DeleteIrModuleModuleDependencys(ids []int64) error {
	return c.Delete(IrModuleModuleDependencyModel, ids)
}

// GetIrModuleModuleDependency gets ir.module.module.dependency existing record.
func (c *Client) GetIrModuleModuleDependency(id int64) (*IrModuleModuleDependency, error) {
	immds, err := c.GetIrModuleModuleDependencys([]int64{id})
	if err != nil {
		return nil, err
	}
	if immds != nil && len(*immds) > 0 {
		return &((*immds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.module.module.dependency not found", id)
}

// GetIrModuleModuleDependencys gets ir.module.module.dependency existing records.
func (c *Client) GetIrModuleModuleDependencys(ids []int64) (*IrModuleModuleDependencys, error) {
	immds := &IrModuleModuleDependencys{}
	if err := c.Read(IrModuleModuleDependencyModel, ids, nil, immds); err != nil {
		return nil, err
	}
	return immds, nil
}

// FindIrModuleModuleDependency finds ir.module.module.dependency record by querying it with criteria.
func (c *Client) FindIrModuleModuleDependency(criteria *Criteria) (*IrModuleModuleDependency, error) {
	immds := &IrModuleModuleDependencys{}
	if err := c.SearchRead(IrModuleModuleDependencyModel, criteria, NewOptions().Limit(1), immds); err != nil {
		return nil, err
	}
	if immds != nil && len(*immds) > 0 {
		return &((*immds)[0]), nil
	}
	return nil, fmt.Errorf("ir.module.module.dependency was not found with criteria %v", criteria)
}

// FindIrModuleModuleDependencys finds ir.module.module.dependency records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModuleDependencys(criteria *Criteria, options *Options) (*IrModuleModuleDependencys, error) {
	immds := &IrModuleModuleDependencys{}
	if err := c.SearchRead(IrModuleModuleDependencyModel, criteria, options, immds); err != nil {
		return nil, err
	}
	return immds, nil
}

// FindIrModuleModuleDependencyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModuleDependencyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModuleModuleDependencyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModuleModuleDependencyId finds record id by querying it with criteria.
func (c *Client) FindIrModuleModuleDependencyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModuleModuleDependencyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.module.module.dependency was not found with criteria %v and options %v", criteria, options)
}
