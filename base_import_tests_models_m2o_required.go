package odoo

import (
	"fmt"
)

// BaseImportTestsModelsM2ORequired represents base_import.tests.models.m2o.required model.
type BaseImportTestsModelsM2ORequired struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *Many2One `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsM2ORequireds represents array of base_import.tests.models.m2o.required model.
type BaseImportTestsModelsM2ORequireds []BaseImportTestsModelsM2ORequired

// BaseImportTestsModelsM2ORequiredModel is the odoo model name.
const BaseImportTestsModelsM2ORequiredModel = "base_import.tests.models.m2o.required"

// Many2One convert BaseImportTestsModelsM2ORequired to *Many2One.
func (btmmr *BaseImportTestsModelsM2ORequired) Many2One() *Many2One {
	return NewMany2One(btmmr.Id.Get(), "")
}

// CreateBaseImportTestsModelsM2ORequired creates a new base_import.tests.models.m2o.required model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORequired(btmmr *BaseImportTestsModelsM2ORequired) (int64, error) {
	return c.Create(BaseImportTestsModelsM2ORequiredModel, btmmr)
}

// UpdateBaseImportTestsModelsM2ORequired updates an existing base_import.tests.models.m2o.required record.
func (c *Client) UpdateBaseImportTestsModelsM2ORequired(btmmr *BaseImportTestsModelsM2ORequired) error {
	return c.UpdateBaseImportTestsModelsM2ORequireds([]int64{btmmr.Id.Get()}, btmmr)
}

// UpdateBaseImportTestsModelsM2ORequireds updates existing base_import.tests.models.m2o.required records.
// All records (represented by ids) will be updated by btmmr values.
func (c *Client) UpdateBaseImportTestsModelsM2ORequireds(ids []int64, btmmr *BaseImportTestsModelsM2ORequired) error {
	return c.Update(BaseImportTestsModelsM2ORequiredModel, ids, btmmr)
}

// DeleteBaseImportTestsModelsM2ORequired deletes an existing base_import.tests.models.m2o.required record.
func (c *Client) DeleteBaseImportTestsModelsM2ORequired(id int64) error {
	return c.DeleteBaseImportTestsModelsM2ORequireds([]int64{id})
}

// DeleteBaseImportTestsModelsM2ORequireds deletes existing base_import.tests.models.m2o.required records.
func (c *Client) DeleteBaseImportTestsModelsM2ORequireds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsM2ORequiredModel, ids)
}

// GetBaseImportTestsModelsM2ORequired gets base_import.tests.models.m2o.required existing record.
func (c *Client) GetBaseImportTestsModelsM2ORequired(id int64) (*BaseImportTestsModelsM2ORequired, error) {
	btmmrs, err := c.GetBaseImportTestsModelsM2ORequireds([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmmrs != nil && len(*btmmrs) > 0 {
		return &((*btmmrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.m2o.required not found", id)
}

// GetBaseImportTestsModelsM2ORequireds gets base_import.tests.models.m2o.required existing records.
func (c *Client) GetBaseImportTestsModelsM2ORequireds(ids []int64) (*BaseImportTestsModelsM2ORequireds, error) {
	btmmrs := &BaseImportTestsModelsM2ORequireds{}
	if err := c.Read(BaseImportTestsModelsM2ORequiredModel, ids, nil, btmmrs); err != nil {
		return nil, err
	}
	return btmmrs, nil
}

// FindBaseImportTestsModelsM2ORequired finds base_import.tests.models.m2o.required record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequired(criteria *Criteria) (*BaseImportTestsModelsM2ORequired, error) {
	btmmrs := &BaseImportTestsModelsM2ORequireds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredModel, criteria, NewOptions().Limit(1), btmmrs); err != nil {
		return nil, err
	}
	if btmmrs != nil && len(*btmmrs) > 0 {
		return &((*btmmrs)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.m2o.required was not found")
}

// FindBaseImportTestsModelsM2ORequireds finds base_import.tests.models.m2o.required records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequireds(criteria *Criteria, options *Options) (*BaseImportTestsModelsM2ORequireds, error) {
	btmmrs := &BaseImportTestsModelsM2ORequireds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredModel, criteria, options, btmmrs); err != nil {
		return nil, err
	}
	return btmmrs, nil
}

// FindBaseImportTestsModelsM2ORequiredIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequiredIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORequiredModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsM2ORequiredId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequiredId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORequiredModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.m2o.required was not found")
}
