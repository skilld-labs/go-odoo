package odoo

import (
	"fmt"
)

// BaseImportTestsModelsM2ORequiredRelated represents base_import.tests.models.m2o.required.related model.
type BaseImportTestsModelsM2ORequiredRelated struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *Int      `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsM2ORequiredRelateds represents array of base_import.tests.models.m2o.required.related model.
type BaseImportTestsModelsM2ORequiredRelateds []BaseImportTestsModelsM2ORequiredRelated

// BaseImportTestsModelsM2ORequiredRelatedModel is the odoo model name.
const BaseImportTestsModelsM2ORequiredRelatedModel = "base_import.tests.models.m2o.required.related"

// Many2One convert BaseImportTestsModelsM2ORequiredRelated to *Many2One.
func (btmmrr *BaseImportTestsModelsM2ORequiredRelated) Many2One() *Many2One {
	return NewMany2One(btmmrr.Id.Get(), "")
}

// CreateBaseImportTestsModelsM2ORequiredRelated creates a new base_import.tests.models.m2o.required.related model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORequiredRelated(btmmrr *BaseImportTestsModelsM2ORequiredRelated) (int64, error) {
	return c.Create(BaseImportTestsModelsM2ORequiredRelatedModel, btmmrr)
}

// UpdateBaseImportTestsModelsM2ORequiredRelated updates an existing base_import.tests.models.m2o.required.related record.
func (c *Client) UpdateBaseImportTestsModelsM2ORequiredRelated(btmmrr *BaseImportTestsModelsM2ORequiredRelated) error {
	return c.UpdateBaseImportTestsModelsM2ORequiredRelateds([]int64{btmmrr.Id.Get()}, btmmrr)
}

// UpdateBaseImportTestsModelsM2ORequiredRelateds updates existing base_import.tests.models.m2o.required.related records.
// All records (represented by IDs) will be updated by btmmrr values.
func (c *Client) UpdateBaseImportTestsModelsM2ORequiredRelateds(ids []int64, btmmrr *BaseImportTestsModelsM2ORequiredRelated) error {
	return c.Update(BaseImportTestsModelsM2ORequiredRelatedModel, ids, btmmrr)
}

// DeleteBaseImportTestsModelsM2ORequiredRelated deletes an existing base_import.tests.models.m2o.required.related record.
func (c *Client) DeleteBaseImportTestsModelsM2ORequiredRelated(id int64) error {
	return c.DeleteBaseImportTestsModelsM2ORequiredRelateds([]int64{id})
}

// DeleteBaseImportTestsModelsM2ORequiredRelateds deletes existing base_import.tests.models.m2o.required.related records.
func (c *Client) DeleteBaseImportTestsModelsM2ORequiredRelateds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsM2ORequiredRelatedModel, ids)
}

// GetBaseImportTestsModelsM2ORequiredRelated gets base_import.tests.models.m2o.required.related existing record.
func (c *Client) GetBaseImportTestsModelsM2ORequiredRelated(id int64) (*BaseImportTestsModelsM2ORequiredRelated, error) {
	btmmrrs, err := c.GetBaseImportTestsModelsM2ORequiredRelateds([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmmrrs != nil && len(*btmmrrs) > 0 {
		return &((*btmmrrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.m2o.required.related not found", id)
}

// GetBaseImportTestsModelsM2ORequiredRelateds gets base_import.tests.models.m2o.required.related existing records.
func (c *Client) GetBaseImportTestsModelsM2ORequiredRelateds(ids []int64) (*BaseImportTestsModelsM2ORequiredRelateds, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.Read(BaseImportTestsModelsM2ORequiredRelatedModel, ids, nil, btmmrrs); err != nil {
		return nil, err
	}
	return btmmrrs, nil
}

// FindBaseImportTestsModelsM2ORequiredRelated finds base_import.tests.models.m2o.required.related record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelated(criteria *Criteria) (*BaseImportTestsModelsM2ORequiredRelated, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, NewOptions().Limit(1), btmmrrs); err != nil {
		return nil, err
	}
	if btmmrrs != nil && len(*btmmrrs) > 0 {
		return &((*btmmrrs)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.m2o.required.related was not found")
}

// FindBaseImportTestsModelsM2ORequiredRelateds finds base_import.tests.models.m2o.required.related records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelateds(criteria *Criteria, options *Options) (*BaseImportTestsModelsM2ORequiredRelateds, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options, btmmrrs); err != nil {
		return nil, err
	}
	return btmmrrs, nil
}

// FindBaseImportTestsModelsM2ORequiredRelatedIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelatedIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsM2ORequiredRelatedId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelatedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.m2o.required.related was not found")
}
