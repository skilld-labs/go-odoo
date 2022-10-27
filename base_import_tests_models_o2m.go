package odoo

import (
	"fmt"
)

// BaseImportTestsModelsO2M represents base_import.tests.models.o2m model.
type BaseImportTestsModelsO2M struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Value       *Relation `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsO2Ms represents array of base_import.tests.models.o2m model.
type BaseImportTestsModelsO2Ms []BaseImportTestsModelsO2M

// BaseImportTestsModelsO2MModel is the odoo model name.
const BaseImportTestsModelsO2MModel = "base_import.tests.models.o2m"

// Many2One convert BaseImportTestsModelsO2M to *Many2One.
func (btmo *BaseImportTestsModelsO2M) Many2One() *Many2One {
	return NewMany2One(btmo.Id.Get(), "")
}

// CreateBaseImportTestsModelsO2M creates a new base_import.tests.models.o2m model and returns its id.
func (c *Client) CreateBaseImportTestsModelsO2M(btmo *BaseImportTestsModelsO2M) (int64, error) {
	return c.Create(BaseImportTestsModelsO2MModel, btmo)
}

// UpdateBaseImportTestsModelsO2M updates an existing base_import.tests.models.o2m record.
func (c *Client) UpdateBaseImportTestsModelsO2M(btmo *BaseImportTestsModelsO2M) error {
	return c.UpdateBaseImportTestsModelsO2Ms([]int64{btmo.Id.Get()}, btmo)
}

// UpdateBaseImportTestsModelsO2Ms updates existing base_import.tests.models.o2m records.
// All records (represented by ids) will be updated by btmo values.
func (c *Client) UpdateBaseImportTestsModelsO2Ms(ids []int64, btmo *BaseImportTestsModelsO2M) error {
	return c.Update(BaseImportTestsModelsO2MModel, ids, btmo)
}

// DeleteBaseImportTestsModelsO2M deletes an existing base_import.tests.models.o2m record.
func (c *Client) DeleteBaseImportTestsModelsO2M(id int64) error {
	return c.DeleteBaseImportTestsModelsO2Ms([]int64{id})
}

// DeleteBaseImportTestsModelsO2Ms deletes existing base_import.tests.models.o2m records.
func (c *Client) DeleteBaseImportTestsModelsO2Ms(ids []int64) error {
	return c.Delete(BaseImportTestsModelsO2MModel, ids)
}

// GetBaseImportTestsModelsO2M gets base_import.tests.models.o2m existing record.
func (c *Client) GetBaseImportTestsModelsO2M(id int64) (*BaseImportTestsModelsO2M, error) {
	btmos, err := c.GetBaseImportTestsModelsO2Ms([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmos != nil && len(*btmos) > 0 {
		return &((*btmos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.o2m not found", id)
}

// GetBaseImportTestsModelsO2Ms gets base_import.tests.models.o2m existing records.
func (c *Client) GetBaseImportTestsModelsO2Ms(ids []int64) (*BaseImportTestsModelsO2Ms, error) {
	btmos := &BaseImportTestsModelsO2Ms{}
	if err := c.Read(BaseImportTestsModelsO2MModel, ids, nil, btmos); err != nil {
		return nil, err
	}
	return btmos, nil
}

// FindBaseImportTestsModelsO2M finds base_import.tests.models.o2m record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsO2M(criteria *Criteria) (*BaseImportTestsModelsO2M, error) {
	btmos := &BaseImportTestsModelsO2Ms{}
	if err := c.SearchRead(BaseImportTestsModelsO2MModel, criteria, NewOptions().Limit(1), btmos); err != nil {
		return nil, err
	}
	if btmos != nil && len(*btmos) > 0 {
		return &((*btmos)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.o2m was not found")
}

// FindBaseImportTestsModelsO2Ms finds base_import.tests.models.o2m records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsO2Ms(criteria *Criteria, options *Options) (*BaseImportTestsModelsO2Ms, error) {
	btmos := &BaseImportTestsModelsO2Ms{}
	if err := c.SearchRead(BaseImportTestsModelsO2MModel, criteria, options, btmos); err != nil {
		return nil, err
	}
	return btmos, nil
}

// FindBaseImportTestsModelsO2MIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsO2MIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsO2MModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsO2MId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsO2MId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsO2MModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.o2m was not found")
}
