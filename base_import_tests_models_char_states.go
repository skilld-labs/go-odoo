package odoo

import (
	"fmt"
)

// BaseImportTestsModelsCharStates represents base_import.tests.models.char.states model.
type BaseImportTestsModelsCharStates struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsCharStatess represents array of base_import.tests.models.char.states model.
type BaseImportTestsModelsCharStatess []BaseImportTestsModelsCharStates

// BaseImportTestsModelsCharStatesModel is the odoo model name.
const BaseImportTestsModelsCharStatesModel = "base_import.tests.models.char.states"

// Many2One convert BaseImportTestsModelsCharStates to *Many2One.
func (btmcs *BaseImportTestsModelsCharStates) Many2One() *Many2One {
	return NewMany2One(btmcs.Id.Get(), "")
}

// CreateBaseImportTestsModelsCharStates creates a new base_import.tests.models.char.states model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharStates(btmcs *BaseImportTestsModelsCharStates) (int64, error) {
	return c.Create(BaseImportTestsModelsCharStatesModel, btmcs)
}

// UpdateBaseImportTestsModelsCharStates updates an existing base_import.tests.models.char.states record.
func (c *Client) UpdateBaseImportTestsModelsCharStates(btmcs *BaseImportTestsModelsCharStates) error {
	return c.UpdateBaseImportTestsModelsCharStatess([]int64{btmcs.Id.Get()}, btmcs)
}

// UpdateBaseImportTestsModelsCharStatess updates existing base_import.tests.models.char.states records.
// All records (represented by ids) will be updated by btmcs values.
func (c *Client) UpdateBaseImportTestsModelsCharStatess(ids []int64, btmcs *BaseImportTestsModelsCharStates) error {
	return c.Update(BaseImportTestsModelsCharStatesModel, ids, btmcs)
}

// DeleteBaseImportTestsModelsCharStates deletes an existing base_import.tests.models.char.states record.
func (c *Client) DeleteBaseImportTestsModelsCharStates(id int64) error {
	return c.DeleteBaseImportTestsModelsCharStatess([]int64{id})
}

// DeleteBaseImportTestsModelsCharStatess deletes existing base_import.tests.models.char.states records.
func (c *Client) DeleteBaseImportTestsModelsCharStatess(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharStatesModel, ids)
}

// GetBaseImportTestsModelsCharStates gets base_import.tests.models.char.states existing record.
func (c *Client) GetBaseImportTestsModelsCharStates(id int64) (*BaseImportTestsModelsCharStates, error) {
	btmcss, err := c.GetBaseImportTestsModelsCharStatess([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmcss != nil && len(*btmcss) > 0 {
		return &((*btmcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.char.states not found", id)
}

// GetBaseImportTestsModelsCharStatess gets base_import.tests.models.char.states existing records.
func (c *Client) GetBaseImportTestsModelsCharStatess(ids []int64) (*BaseImportTestsModelsCharStatess, error) {
	btmcss := &BaseImportTestsModelsCharStatess{}
	if err := c.Read(BaseImportTestsModelsCharStatesModel, ids, nil, btmcss); err != nil {
		return nil, err
	}
	return btmcss, nil
}

// FindBaseImportTestsModelsCharStates finds base_import.tests.models.char.states record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharStates(criteria *Criteria) (*BaseImportTestsModelsCharStates, error) {
	btmcss := &BaseImportTestsModelsCharStatess{}
	if err := c.SearchRead(BaseImportTestsModelsCharStatesModel, criteria, NewOptions().Limit(1), btmcss); err != nil {
		return nil, err
	}
	if btmcss != nil && len(*btmcss) > 0 {
		return &((*btmcss)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.char.states was not found")
}

// FindBaseImportTestsModelsCharStatess finds base_import.tests.models.char.states records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharStatess(criteria *Criteria, options *Options) (*BaseImportTestsModelsCharStatess, error) {
	btmcss := &BaseImportTestsModelsCharStatess{}
	if err := c.SearchRead(BaseImportTestsModelsCharStatesModel, criteria, options, btmcss); err != nil {
		return nil, err
	}
	return btmcss, nil
}

// FindBaseImportTestsModelsCharStatesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharStatesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharStatesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsCharStatesId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharStatesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharStatesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.char.states was not found")
}
