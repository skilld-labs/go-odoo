package odoo

import (
	"fmt"
)

// BaseImportMapping represents base_import.mapping model.
type BaseImportMapping struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	ColumnName  *String   `xmlrpc:"column_name,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FieldName   *String   `xmlrpc:"field_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ResModel    *String   `xmlrpc:"res_model,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportMappings represents array of base_import.mapping model.
type BaseImportMappings []BaseImportMapping

// BaseImportMappingModel is the odoo model name.
const BaseImportMappingModel = "base_import.mapping"

// Many2One convert BaseImportMapping to *Many2One.
func (bm *BaseImportMapping) Many2One() *Many2One {
	return NewMany2One(bm.Id.Get(), "")
}

// CreateBaseImportMapping creates a new base_import.mapping model and returns its id.
func (c *Client) CreateBaseImportMapping(bm *BaseImportMapping) (int64, error) {
	return c.Create(BaseImportMappingModel, bm)
}

// UpdateBaseImportMapping updates an existing base_import.mapping record.
func (c *Client) UpdateBaseImportMapping(bm *BaseImportMapping) error {
	return c.UpdateBaseImportMappings([]int64{bm.Id.Get()}, bm)
}

// UpdateBaseImportMappings updates existing base_import.mapping records.
// All records (represented by IDs) will be updated by bm values.
func (c *Client) UpdateBaseImportMappings(ids []int64, bm *BaseImportMapping) error {
	return c.Update(BaseImportMappingModel, ids, bm)
}

// DeleteBaseImportMapping deletes an existing base_import.mapping record.
func (c *Client) DeleteBaseImportMapping(id int64) error {
	return c.DeleteBaseImportMappings([]int64{id})
}

// DeleteBaseImportMappings deletes existing base_import.mapping records.
func (c *Client) DeleteBaseImportMappings(ids []int64) error {
	return c.Delete(BaseImportMappingModel, ids)
}

// GetBaseImportMapping gets base_import.mapping existing record.
func (c *Client) GetBaseImportMapping(id int64) (*BaseImportMapping, error) {
	bms, err := c.GetBaseImportMappings([]int64{id})
	if err != nil {
		return nil, err
	}
	if bms != nil && len(*bms) > 0 {
		return &((*bms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.mapping not found", id)
}

// GetBaseImportMappings gets base_import.mapping existing records.
func (c *Client) GetBaseImportMappings(ids []int64) (*BaseImportMappings, error) {
	bms := &BaseImportMappings{}
	if err := c.Read(BaseImportMappingModel, ids, nil, bms); err != nil {
		return nil, err
	}
	return bms, nil
}

// FindBaseImportMapping finds base_import.mapping record by querying it with criteria.
func (c *Client) FindBaseImportMapping(criteria *Criteria) (*BaseImportMapping, error) {
	bms := &BaseImportMappings{}
	if err := c.SearchRead(BaseImportMappingModel, criteria, NewOptions().Limit(1), bms); err != nil {
		return nil, err
	}
	if bms != nil && len(*bms) > 0 {
		return &((*bms)[0]), nil
	}
	return nil, fmt.Errorf("base_import.mapping was not found")
}

// FindBaseImportMappings finds base_import.mapping records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportMappings(criteria *Criteria, options *Options) (*BaseImportMappings, error) {
	bms := &BaseImportMappings{}
	if err := c.SearchRead(BaseImportMappingModel, criteria, options, bms); err != nil {
		return nil, err
	}
	return bms, nil
}

// FindBaseImportMappingIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportMappingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportMappingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportMappingId finds record id by querying it with criteria.
func (c *Client) FindBaseImportMappingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportMappingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.mapping was not found")
}
