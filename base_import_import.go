package odoo

import (
	"fmt"
)

// BaseImportImport represents base_import.import model.
type BaseImportImport struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	File        *String   `xmlrpc:"file,omptempty"`
	FileName    *String   `xmlrpc:"file_name,omptempty"`
	FileType    *String   `xmlrpc:"file_type,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ResModel    *String   `xmlrpc:"res_model,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseImportImports represents array of base_import.import model.
type BaseImportImports []BaseImportImport

// BaseImportImportModel is the odoo model name.
const BaseImportImportModel = "base_import.import"

// Many2One convert BaseImportImport to *Many2One.
func (bi *BaseImportImport) Many2One() *Many2One {
	return NewMany2One(bi.Id.Get(), "")
}

// CreateBaseImportImport creates a new base_import.import model and returns its id.
func (c *Client) CreateBaseImportImport(bi *BaseImportImport) (int64, error) {
	ids, err := c.CreateBaseImportImports([]*BaseImportImport{bi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportImport creates a new base_import.import model and returns its id.
func (c *Client) CreateBaseImportImports(bis []*BaseImportImport) ([]int64, error) {
	var vv []interface{}
	for _, v := range bis {
		vv = append(vv, v)
	}
	return c.Create(BaseImportImportModel, vv)
}

// UpdateBaseImportImport updates an existing base_import.import record.
func (c *Client) UpdateBaseImportImport(bi *BaseImportImport) error {
	return c.UpdateBaseImportImports([]int64{bi.Id.Get()}, bi)
}

// UpdateBaseImportImports updates existing base_import.import records.
// All records (represented by ids) will be updated by bi values.
func (c *Client) UpdateBaseImportImports(ids []int64, bi *BaseImportImport) error {
	return c.Update(BaseImportImportModel, ids, bi)
}

// DeleteBaseImportImport deletes an existing base_import.import record.
func (c *Client) DeleteBaseImportImport(id int64) error {
	return c.DeleteBaseImportImports([]int64{id})
}

// DeleteBaseImportImports deletes existing base_import.import records.
func (c *Client) DeleteBaseImportImports(ids []int64) error {
	return c.Delete(BaseImportImportModel, ids)
}

// GetBaseImportImport gets base_import.import existing record.
func (c *Client) GetBaseImportImport(id int64) (*BaseImportImport, error) {
	bis, err := c.GetBaseImportImports([]int64{id})
	if err != nil {
		return nil, err
	}
	if bis != nil && len(*bis) > 0 {
		return &((*bis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.import not found", id)
}

// GetBaseImportImports gets base_import.import existing records.
func (c *Client) GetBaseImportImports(ids []int64) (*BaseImportImports, error) {
	bis := &BaseImportImports{}
	if err := c.Read(BaseImportImportModel, ids, nil, bis); err != nil {
		return nil, err
	}
	return bis, nil
}

// FindBaseImportImport finds base_import.import record by querying it with criteria.
func (c *Client) FindBaseImportImport(criteria *Criteria) (*BaseImportImport, error) {
	bis := &BaseImportImports{}
	if err := c.SearchRead(BaseImportImportModel, criteria, NewOptions().Limit(1), bis); err != nil {
		return nil, err
	}
	if bis != nil && len(*bis) > 0 {
		return &((*bis)[0]), nil
	}
	return nil, fmt.Errorf("base_import.import was not found with criteria %v", criteria)
}

// FindBaseImportImports finds base_import.import records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportImports(criteria *Criteria, options *Options) (*BaseImportImports, error) {
	bis := &BaseImportImports{}
	if err := c.SearchRead(BaseImportImportModel, criteria, options, bis); err != nil {
		return nil, err
	}
	return bis, nil
}

// FindBaseImportImportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportImportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportImportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportImportId finds record id by querying it with criteria.
func (c *Client) FindBaseImportImportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportImportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.import was not found with criteria %v and options %v", criteria, options)
}
