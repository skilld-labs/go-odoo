package odoo

import (
	"fmt"
)

// IrModuleModule represents ir.module.module model.
type IrModuleModule struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	Application      *Bool      `xmlrpc:"application,omitempty"`
	Author           *String    `xmlrpc:"author,omitempty"`
	AutoInstall      *Bool      `xmlrpc:"auto_install,omitempty"`
	CategoryId       *Many2One  `xmlrpc:"category_id,omitempty"`
	Contributors     *String    `xmlrpc:"contributors,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Demo             *Bool      `xmlrpc:"demo,omitempty"`
	DependenciesId   *Relation  `xmlrpc:"dependencies_id,omitempty"`
	Description      *String    `xmlrpc:"description,omitempty"`
	DescriptionHtml  *String    `xmlrpc:"description_html,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	ExclusionIds     *Relation  `xmlrpc:"exclusion_ids,omitempty"`
	Icon             *String    `xmlrpc:"icon,omitempty"`
	IconImage        *String    `xmlrpc:"icon_image,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	InstalledVersion *String    `xmlrpc:"installed_version,omitempty"`
	LatestVersion    *String    `xmlrpc:"latest_version,omitempty"`
	License          *Selection `xmlrpc:"license,omitempty"`
	Maintainer       *String    `xmlrpc:"maintainer,omitempty"`
	MenusByModule    *String    `xmlrpc:"menus_by_module,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	PublishedVersion *String    `xmlrpc:"published_version,omitempty"`
	ReportsByModule  *String    `xmlrpc:"reports_by_module,omitempty"`
	Sequence         *Int       `xmlrpc:"sequence,omitempty"`
	Shortdesc        *String    `xmlrpc:"shortdesc,omitempty"`
	State            *Selection `xmlrpc:"state,omitempty"`
	Summary          *String    `xmlrpc:"summary,omitempty"`
	Url              *String    `xmlrpc:"url,omitempty"`
	ViewsByModule    *String    `xmlrpc:"views_by_module,omitempty"`
	Website          *String    `xmlrpc:"website,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrModuleModules represents array of ir.module.module model.
type IrModuleModules []IrModuleModule

// IrModuleModuleModel is the odoo model name.
const IrModuleModuleModel = "ir.module.module"

// Many2One convert IrModuleModule to *Many2One.
func (imm *IrModuleModule) Many2One() *Many2One {
	return NewMany2One(imm.Id.Get(), "")
}

// CreateIrModuleModule creates a new ir.module.module model and returns its id.
func (c *Client) CreateIrModuleModule(imm *IrModuleModule) (int64, error) {
	return c.Create(IrModuleModuleModel, imm)
}

// UpdateIrModuleModule updates an existing ir.module.module record.
func (c *Client) UpdateIrModuleModule(imm *IrModuleModule) error {
	return c.UpdateIrModuleModules([]int64{imm.Id.Get()}, imm)
}

// UpdateIrModuleModules updates existing ir.module.module records.
// All records (represented by ids) will be updated by imm values.
func (c *Client) UpdateIrModuleModules(ids []int64, imm *IrModuleModule) error {
	return c.Update(IrModuleModuleModel, ids, imm)
}

// DeleteIrModuleModule deletes an existing ir.module.module record.
func (c *Client) DeleteIrModuleModule(id int64) error {
	return c.DeleteIrModuleModules([]int64{id})
}

// DeleteIrModuleModules deletes existing ir.module.module records.
func (c *Client) DeleteIrModuleModules(ids []int64) error {
	return c.Delete(IrModuleModuleModel, ids)
}

// GetIrModuleModule gets ir.module.module existing record.
func (c *Client) GetIrModuleModule(id int64) (*IrModuleModule, error) {
	imms, err := c.GetIrModuleModules([]int64{id})
	if err != nil {
		return nil, err
	}
	if imms != nil && len(*imms) > 0 {
		return &((*imms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.module.module not found", id)
}

// GetIrModuleModules gets ir.module.module existing records.
func (c *Client) GetIrModuleModules(ids []int64) (*IrModuleModules, error) {
	imms := &IrModuleModules{}
	if err := c.Read(IrModuleModuleModel, ids, nil, imms); err != nil {
		return nil, err
	}
	return imms, nil
}

// FindIrModuleModule finds ir.module.module record by querying it with criteria.
func (c *Client) FindIrModuleModule(criteria *Criteria) (*IrModuleModule, error) {
	imms := &IrModuleModules{}
	if err := c.SearchRead(IrModuleModuleModel, criteria, NewOptions().Limit(1), imms); err != nil {
		return nil, err
	}
	if imms != nil && len(*imms) > 0 {
		return &((*imms)[0]), nil
	}
	return nil, fmt.Errorf("no ir.module.module was found with criteria %v", criteria)
}

// FindIrModuleModules finds ir.module.module records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModules(criteria *Criteria, options *Options) (*IrModuleModules, error) {
	imms := &IrModuleModules{}
	if err := c.SearchRead(IrModuleModuleModel, criteria, options, imms); err != nil {
		return nil, err
	}
	return imms, nil
}

// FindIrModuleModuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleModuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModuleModuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModuleModuleId finds record id by querying it with criteria.
func (c *Client) FindIrModuleModuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModuleModuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.module.module was found with criteria %v and options %v", criteria, options)
}
