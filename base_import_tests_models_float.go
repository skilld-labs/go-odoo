package odoo

import (
	"fmt"
)

// BaseImportTestsModelsFloat represents base_import.tests.models.float model.
type BaseImportTestsModelsFloat struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *Float    `xmlrpc:"value,omitempty"`
	Value2      *Float    `xmlrpc:"value2,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsFloats represents array of base_import.tests.models.float model.
type BaseImportTestsModelsFloats []BaseImportTestsModelsFloat

// BaseImportTestsModelsFloatModel is the odoo model name.
const BaseImportTestsModelsFloatModel = "base_import.tests.models.float"

// Many2One convert BaseImportTestsModelsFloat to *Many2One.
func (btmf *BaseImportTestsModelsFloat) Many2One() *Many2One {
	return NewMany2One(btmf.Id.Get(), "")
}

// CreateBaseImportTestsModelsFloat creates a new base_import.tests.models.float model and returns its id.
func (c *Client) CreateBaseImportTestsModelsFloat(btmf *BaseImportTestsModelsFloat) (int64, error) {
	return c.Create(BaseImportTestsModelsFloatModel, btmf)
}

// UpdateBaseImportTestsModelsFloat updates an existing base_import.tests.models.float record.
func (c *Client) UpdateBaseImportTestsModelsFloat(btmf *BaseImportTestsModelsFloat) error {
	return c.UpdateBaseImportTestsModelsFloats([]int64{btmf.Id.Get()}, btmf)
}

// UpdateBaseImportTestsModelsFloats updates existing base_import.tests.models.float records.
// All records (represented by ids) will be updated by btmf values.
func (c *Client) UpdateBaseImportTestsModelsFloats(ids []int64, btmf *BaseImportTestsModelsFloat) error {
	return c.Update(BaseImportTestsModelsFloatModel, ids, btmf)
}

// DeleteBaseImportTestsModelsFloat deletes an existing base_import.tests.models.float record.
func (c *Client) DeleteBaseImportTestsModelsFloat(id int64) error {
	return c.DeleteBaseImportTestsModelsFloats([]int64{id})
}

// DeleteBaseImportTestsModelsFloats deletes existing base_import.tests.models.float records.
func (c *Client) DeleteBaseImportTestsModelsFloats(ids []int64) error {
	return c.Delete(BaseImportTestsModelsFloatModel, ids)
}

// GetBaseImportTestsModelsFloat gets base_import.tests.models.float existing record.
func (c *Client) GetBaseImportTestsModelsFloat(id int64) (*BaseImportTestsModelsFloat, error) {
	btmfs, err := c.GetBaseImportTestsModelsFloats([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmfs != nil && len(*btmfs) > 0 {
		return &((*btmfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.float not found", id)
}

// GetBaseImportTestsModelsFloats gets base_import.tests.models.float existing records.
func (c *Client) GetBaseImportTestsModelsFloats(ids []int64) (*BaseImportTestsModelsFloats, error) {
	btmfs := &BaseImportTestsModelsFloats{}
	if err := c.Read(BaseImportTestsModelsFloatModel, ids, nil, btmfs); err != nil {
		return nil, err
	}
	return btmfs, nil
}

// FindBaseImportTestsModelsFloat finds base_import.tests.models.float record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsFloat(criteria *Criteria) (*BaseImportTestsModelsFloat, error) {
	btmfs := &BaseImportTestsModelsFloats{}
	if err := c.SearchRead(BaseImportTestsModelsFloatModel, criteria, NewOptions().Limit(1), btmfs); err != nil {
		return nil, err
	}
	if btmfs != nil && len(*btmfs) > 0 {
		return &((*btmfs)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.float was not found")
}

// FindBaseImportTestsModelsFloats finds base_import.tests.models.float records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsFloats(criteria *Criteria, options *Options) (*BaseImportTestsModelsFloats, error) {
	btmfs := &BaseImportTestsModelsFloats{}
	if err := c.SearchRead(BaseImportTestsModelsFloatModel, criteria, options, btmfs); err != nil {
		return nil, err
	}
	return btmfs, nil
}

// FindBaseImportTestsModelsFloatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsFloatIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsFloatModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsFloatId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsFloatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsFloatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.float was not found")
}
