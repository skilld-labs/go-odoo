package odoo

import (
	"fmt"
)

// BaseImportTestsModelsCharStillreadonly represents base_import.tests.models.char.stillreadonly model.
type BaseImportTestsModelsCharStillreadonly struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Value       *String   `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseImportTestsModelsCharStillreadonlys represents array of base_import.tests.models.char.stillreadonly model.
type BaseImportTestsModelsCharStillreadonlys []BaseImportTestsModelsCharStillreadonly

// BaseImportTestsModelsCharStillreadonlyModel is the odoo model name.
const BaseImportTestsModelsCharStillreadonlyModel = "base_import.tests.models.char.stillreadonly"

// Many2One convert BaseImportTestsModelsCharStillreadonly to *Many2One.
func (btmcs *BaseImportTestsModelsCharStillreadonly) Many2One() *Many2One {
	return NewMany2One(btmcs.Id.Get(), "")
}

// CreateBaseImportTestsModelsCharStillreadonly creates a new base_import.tests.models.char.stillreadonly model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharStillreadonly(btmcs *BaseImportTestsModelsCharStillreadonly) (int64, error) {
	return c.Create(BaseImportTestsModelsCharStillreadonlyModel, btmcs)
}

// UpdateBaseImportTestsModelsCharStillreadonly updates an existing base_import.tests.models.char.stillreadonly record.
func (c *Client) UpdateBaseImportTestsModelsCharStillreadonly(btmcs *BaseImportTestsModelsCharStillreadonly) error {
	return c.UpdateBaseImportTestsModelsCharStillreadonlys([]int64{btmcs.Id.Get()}, btmcs)
}

// UpdateBaseImportTestsModelsCharStillreadonlys updates existing base_import.tests.models.char.stillreadonly records.
// All records (represented by ids) will be updated by btmcs values.
func (c *Client) UpdateBaseImportTestsModelsCharStillreadonlys(ids []int64, btmcs *BaseImportTestsModelsCharStillreadonly) error {
	return c.Update(BaseImportTestsModelsCharStillreadonlyModel, ids, btmcs)
}

// DeleteBaseImportTestsModelsCharStillreadonly deletes an existing base_import.tests.models.char.stillreadonly record.
func (c *Client) DeleteBaseImportTestsModelsCharStillreadonly(id int64) error {
	return c.DeleteBaseImportTestsModelsCharStillreadonlys([]int64{id})
}

// DeleteBaseImportTestsModelsCharStillreadonlys deletes existing base_import.tests.models.char.stillreadonly records.
func (c *Client) DeleteBaseImportTestsModelsCharStillreadonlys(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharStillreadonlyModel, ids)
}

// GetBaseImportTestsModelsCharStillreadonly gets base_import.tests.models.char.stillreadonly existing record.
func (c *Client) GetBaseImportTestsModelsCharStillreadonly(id int64) (*BaseImportTestsModelsCharStillreadonly, error) {
	btmcss, err := c.GetBaseImportTestsModelsCharStillreadonlys([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmcss != nil && len(*btmcss) > 0 {
		return &((*btmcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.char.stillreadonly not found", id)
}

// GetBaseImportTestsModelsCharStillreadonlys gets base_import.tests.models.char.stillreadonly existing records.
func (c *Client) GetBaseImportTestsModelsCharStillreadonlys(ids []int64) (*BaseImportTestsModelsCharStillreadonlys, error) {
	btmcss := &BaseImportTestsModelsCharStillreadonlys{}
	if err := c.Read(BaseImportTestsModelsCharStillreadonlyModel, ids, nil, btmcss); err != nil {
		return nil, err
	}
	return btmcss, nil
}

// FindBaseImportTestsModelsCharStillreadonly finds base_import.tests.models.char.stillreadonly record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharStillreadonly(criteria *Criteria) (*BaseImportTestsModelsCharStillreadonly, error) {
	btmcss := &BaseImportTestsModelsCharStillreadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharStillreadonlyModel, criteria, NewOptions().Limit(1), btmcss); err != nil {
		return nil, err
	}
	if btmcss != nil && len(*btmcss) > 0 {
		return &((*btmcss)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.char.stillreadonly was not found with criteria %v", criteria)
}

// FindBaseImportTestsModelsCharStillreadonlys finds base_import.tests.models.char.stillreadonly records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharStillreadonlys(criteria *Criteria, options *Options) (*BaseImportTestsModelsCharStillreadonlys, error) {
	btmcss := &BaseImportTestsModelsCharStillreadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharStillreadonlyModel, criteria, options, btmcss); err != nil {
		return nil, err
	}
	return btmcss, nil
}

// FindBaseImportTestsModelsCharStillreadonlyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharStillreadonlyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharStillreadonlyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsCharStillreadonlyId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharStillreadonlyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharStillreadonlyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.char.stillreadonly was not found with criteria %v and options %v", criteria, options)
}
