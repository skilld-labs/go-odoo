package odoo

import (
	"fmt"
)

// IrModuleCategory represents ir.module.category model.
type IrModuleCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ChildIds    *Relation `xmlrpc:"child_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Exclusive   *Bool     `xmlrpc:"exclusive,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModuleIds   *Relation `xmlrpc:"module_ids,omptempty"`
	ModuleNr    *Int      `xmlrpc:"module_nr,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Visible     *Bool     `xmlrpc:"visible,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XmlId       *String   `xmlrpc:"xml_id,omptempty"`
}

// IrModuleCategorys represents array of ir.module.category model.
type IrModuleCategorys []IrModuleCategory

// IrModuleCategoryModel is the odoo model name.
const IrModuleCategoryModel = "ir.module.category"

// Many2One convert IrModuleCategory to *Many2One.
func (imc *IrModuleCategory) Many2One() *Many2One {
	return NewMany2One(imc.Id.Get(), "")
}

// CreateIrModuleCategory creates a new ir.module.category model and returns its id.
func (c *Client) CreateIrModuleCategory(imc *IrModuleCategory) (int64, error) {
	ids, err := c.CreateIrModuleCategorys([]*IrModuleCategory{imc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModuleCategory creates a new ir.module.category model and returns its id.
func (c *Client) CreateIrModuleCategorys(imcs []*IrModuleCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range imcs {
		vv = append(vv, v)
	}
	return c.Create(IrModuleCategoryModel, vv)
}

// UpdateIrModuleCategory updates an existing ir.module.category record.
func (c *Client) UpdateIrModuleCategory(imc *IrModuleCategory) error {
	return c.UpdateIrModuleCategorys([]int64{imc.Id.Get()}, imc)
}

// UpdateIrModuleCategorys updates existing ir.module.category records.
// All records (represented by ids) will be updated by imc values.
func (c *Client) UpdateIrModuleCategorys(ids []int64, imc *IrModuleCategory) error {
	return c.Update(IrModuleCategoryModel, ids, imc)
}

// DeleteIrModuleCategory deletes an existing ir.module.category record.
func (c *Client) DeleteIrModuleCategory(id int64) error {
	return c.DeleteIrModuleCategorys([]int64{id})
}

// DeleteIrModuleCategorys deletes existing ir.module.category records.
func (c *Client) DeleteIrModuleCategorys(ids []int64) error {
	return c.Delete(IrModuleCategoryModel, ids)
}

// GetIrModuleCategory gets ir.module.category existing record.
func (c *Client) GetIrModuleCategory(id int64) (*IrModuleCategory, error) {
	imcs, err := c.GetIrModuleCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if imcs != nil && len(*imcs) > 0 {
		return &((*imcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.module.category not found", id)
}

// GetIrModuleCategorys gets ir.module.category existing records.
func (c *Client) GetIrModuleCategorys(ids []int64) (*IrModuleCategorys, error) {
	imcs := &IrModuleCategorys{}
	if err := c.Read(IrModuleCategoryModel, ids, nil, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModuleCategory finds ir.module.category record by querying it with criteria.
func (c *Client) FindIrModuleCategory(criteria *Criteria) (*IrModuleCategory, error) {
	imcs := &IrModuleCategorys{}
	if err := c.SearchRead(IrModuleCategoryModel, criteria, NewOptions().Limit(1), imcs); err != nil {
		return nil, err
	}
	if imcs != nil && len(*imcs) > 0 {
		return &((*imcs)[0]), nil
	}
	return nil, fmt.Errorf("ir.module.category was not found with criteria %v", criteria)
}

// FindIrModuleCategorys finds ir.module.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleCategorys(criteria *Criteria, options *Options) (*IrModuleCategorys, error) {
	imcs := &IrModuleCategorys{}
	if err := c.SearchRead(IrModuleCategoryModel, criteria, options, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModuleCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModuleCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModuleCategoryId finds record id by querying it with criteria.
func (c *Client) FindIrModuleCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModuleCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.module.category was not found with criteria %v and options %v", criteria, options)
}
