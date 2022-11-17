package odoo

import (
	"fmt"
)

// BaseImportTestsModelsChar represents base_import.tests.models.char model.
type BaseImportTestsModelsChar struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsChars represents array of base_import.tests.models.char model.
type BaseImportTestsModelsChars []BaseImportTestsModelsChar

// BaseImportTestsModelsCharModel is the odoo model name.
const BaseImportTestsModelsCharModel = "base_import.tests.models.char"

// Many2One convert BaseImportTestsModelsChar to *Many2One.
func (btmc *BaseImportTestsModelsChar) Many2One() *Many2One {
	return NewMany2One(btmc.Id.Get(), "")
}

// CreateBaseImportTestsModelsChar creates a new base_import.tests.models.char model and returns its id.
func (c *Client) CreateBaseImportTestsModelsChar(btmc *BaseImportTestsModelsChar) (int64, error) {
	return c.Create(BaseImportTestsModelsCharModel, btmc)
}

// UpdateBaseImportTestsModelsChar updates an existing base_import.tests.models.char record.
func (c *Client) UpdateBaseImportTestsModelsChar(btmc *BaseImportTestsModelsChar) error {
	return c.UpdateBaseImportTestsModelsChars([]int64{btmc.Id.Get()}, btmc)
}

// UpdateBaseImportTestsModelsChars updates existing base_import.tests.models.char records.
// All records (represented by ids) will be updated by btmc values.
func (c *Client) UpdateBaseImportTestsModelsChars(ids []int64, btmc *BaseImportTestsModelsChar) error {
	return c.Update(BaseImportTestsModelsCharModel, ids, btmc)
}

// DeleteBaseImportTestsModelsChar deletes an existing base_import.tests.models.char record.
func (c *Client) DeleteBaseImportTestsModelsChar(id int64) error {
	return c.DeleteBaseImportTestsModelsChars([]int64{id})
}

// DeleteBaseImportTestsModelsChars deletes existing base_import.tests.models.char records.
func (c *Client) DeleteBaseImportTestsModelsChars(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharModel, ids)
}

// GetBaseImportTestsModelsChar gets base_import.tests.models.char existing record.
func (c *Client) GetBaseImportTestsModelsChar(id int64) (*BaseImportTestsModelsChar, error) {
	btmcs, err := c.GetBaseImportTestsModelsChars([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmcs != nil && len(*btmcs) > 0 {
		return &((*btmcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.char not found", id)
}

// GetBaseImportTestsModelsChars gets base_import.tests.models.char existing records.
func (c *Client) GetBaseImportTestsModelsChars(ids []int64) (*BaseImportTestsModelsChars, error) {
	btmcs := &BaseImportTestsModelsChars{}
	if err := c.Read(BaseImportTestsModelsCharModel, ids, nil, btmcs); err != nil {
		return nil, err
	}
	return btmcs, nil
}

// FindBaseImportTestsModelsChar finds base_import.tests.models.char record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsChar(criteria *Criteria) (*BaseImportTestsModelsChar, error) {
	btmcs := &BaseImportTestsModelsChars{}
	if err := c.SearchRead(BaseImportTestsModelsCharModel, criteria, NewOptions().Limit(1), btmcs); err != nil {
		return nil, err
	}
	if btmcs != nil && len(*btmcs) > 0 {
		return &((*btmcs)[0]), nil
	}
	return nil, fmt.Errorf("no base_import.tests.models.char was found with criteria %v", criteria)
}

// FindBaseImportTestsModelsChars finds base_import.tests.models.char records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsChars(criteria *Criteria, options *Options) (*BaseImportTestsModelsChars, error) {
	btmcs := &BaseImportTestsModelsChars{}
	if err := c.SearchRead(BaseImportTestsModelsCharModel, criteria, options, btmcs); err != nil {
		return nil, err
	}
	return btmcs, nil
}

// FindBaseImportTestsModelsCharIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsCharId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no base_import.tests.models.char was found with criteria %v and options %v", criteria, options)
}
