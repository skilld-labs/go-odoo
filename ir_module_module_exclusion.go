package odoo

import (
	"fmt"
)

// IrModuleModuleExclusion represents ir.module.module.exclusion model.
type IrModuleModuleExclusion struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	ExclusionId *Many2One  `xmlrpc:"exclusion_id,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	ModuleId    *Many2One  `xmlrpc:"module_id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	State       *Selection `xmlrpc:"state,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrModuleModuleExclusions represents array of ir.module.module.exclusion model.
type IrModuleModuleExclusions []IrModuleModuleExclusion

// IrModuleModuleExclusionModel is the odoo model name.
const IrModuleModuleExclusionModel = "ir.module.module.exclusion"

// Many2One convert IrModuleModuleExclusion to *Many2One.
func (imme *IrModuleModuleExclusion) Many2One() *Many2One {
	return NewMany2One(imme.Id.Get(), "")
}

// CreateIrModuleModuleExclusion creates a new ir.module.module.exclusion model and returns its id.
func (c *Client) CreateIrModuleModuleExclusion(imme *IrModuleModuleExclusion) (int64, error) {
	return c.Create(IrModuleModuleExclusionModel, imme)
}

// UpdateIrModuleModuleExclusion updates an existing ir.module.module.exclusion record.
func (c *Client) UpdateIrModuleModuleExclusion(imme *IrModuleModuleExclusion) error {
	return c.UpdateIrModuleModuleExclusions([]int64{imme.Id.Get()}, imme)
}

// UpdateIrModuleModuleExclusions updates existing ir.module.module.exclusion records.
// All records (represented by ids) will be updated by imme values.
func (c *Client) UpdateIrModuleModuleExclusions(ids []int64, imme *IrModuleModuleExclusion) error {
	return c.Update(IrModuleModuleExclusionModel, ids, imme)
}

// DeleteIrModuleModuleExclusion deletes an existing ir.module.module.exclusion record.
func (c *Client) DeleteIrModuleModuleExclusion(id int64) error {
	return c.DeleteIrModuleModuleExclusions([]int64{id})
}

// DeleteIrModuleModuleExclusions deletes existing ir.module.module.exclusion records.
func (c *Client) DeleteIrModuleModuleExclusions(ids []int64) error {
	return c.Delete(IrModuleModuleExclusionModel, ids)
}

// GetIrModuleModuleExclusion gets ir.module.module.exclusion existing record.
func (c *Client) GetIrModuleModuleExclusion(id int64) (*IrModuleModuleExclusion, error) {
	immes, err := c.GetIrModuleModuleExclusions([]int64{id})
	if err != nil {
		return nil, err
	}
	if immes != nil && len(*immes) > 0 {
		return &((*immes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.module.module.exclusion not found", id)
}

// GetIrModuleModuleExclusions gets ir.module.module.exclusion existing records.
func (c *Client) GetIrModuleModuleExclusions(ids []int64) (*IrModuleModuleExclusions, error) {
	immes := &IrModuleModuleExclusions{}
	if err := c.Read(IrModuleModuleExclusionModel, ids, nil, immes); err != nil {
		return nil, err
	}
	return immes, nil
}

// FindIrModuleModuleExclusion finds ir.module.module.exclusion record by querying it with criteria.
func (c *Client) FindIrModuleModuleExclusion(criteria *Criteria) (*IrModuleModuleExclusion, error) {
	immes := &IrModuleModuleExclusions{}
	if err := c.SearchRead(IrModuleModuleExclusionModel, criteria, NewOptions().Limit(1), immes); err != nil {
		return nil, err
	}
	if immes != nil && len(*immes) > 0 {
		return &((*immes)[0]), nil
	}
	return nil, fmt.Errorf("ir.module.module.exclusion was not found")
}

// FindIrModuleModuleExclusions finds ir.module.module.exclusion records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModuleExclusions(criteria *Criteria, options *Options) (*IrModuleModuleExclusions, error) {
	immes := &IrModuleModuleExclusions{}
	if err := c.SearchRead(IrModuleModuleExclusionModel, criteria, options, immes); err != nil {
		return nil, err
	}
	return immes, nil
}

// FindIrModuleModuleExclusionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModuleExclusionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModuleModuleExclusionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModuleModuleExclusionId finds record id by querying it with criteria.
func (c *Client) FindIrModuleModuleExclusionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModuleModuleExclusionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.module.module.exclusion was not found")
}
