package odoo

import (
	"fmt"
)

// BaseImportTestsModelsComplex represents base_import.tests.models.complex model.
type BaseImportTestsModelsComplex struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	C           *String   `xmlrpc:"c,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	D           *Time     `xmlrpc:"d,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Dt          *Time     `xmlrpc:"dt,omitempty"`
	F           *Float    `xmlrpc:"f,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	M           *Float    `xmlrpc:"m,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsComplexs represents array of base_import.tests.models.complex model.
type BaseImportTestsModelsComplexs []BaseImportTestsModelsComplex

// BaseImportTestsModelsComplexModel is the odoo model name.
const BaseImportTestsModelsComplexModel = "base_import.tests.models.complex"

// Many2One convert BaseImportTestsModelsComplex to *Many2One.
func (btmc *BaseImportTestsModelsComplex) Many2One() *Many2One {
	return NewMany2One(btmc.Id.Get(), "")
}

// CreateBaseImportTestsModelsComplex creates a new base_import.tests.models.complex model and returns its id.
func (c *Client) CreateBaseImportTestsModelsComplex(btmc *BaseImportTestsModelsComplex) (int64, error) {
	return c.Create(BaseImportTestsModelsComplexModel, btmc)
}

// UpdateBaseImportTestsModelsComplex updates an existing base_import.tests.models.complex record.
func (c *Client) UpdateBaseImportTestsModelsComplex(btmc *BaseImportTestsModelsComplex) error {
	return c.UpdateBaseImportTestsModelsComplexs([]int64{btmc.Id.Get()}, btmc)
}

// UpdateBaseImportTestsModelsComplexs updates existing base_import.tests.models.complex records.
// All records (represented by IDs) will be updated by btmc values.
func (c *Client) UpdateBaseImportTestsModelsComplexs(ids []int64, btmc *BaseImportTestsModelsComplex) error {
	return c.Update(BaseImportTestsModelsComplexModel, ids, btmc)
}

// DeleteBaseImportTestsModelsComplex deletes an existing base_import.tests.models.complex record.
func (c *Client) DeleteBaseImportTestsModelsComplex(id int64) error {
	return c.DeleteBaseImportTestsModelsComplexs([]int64{id})
}

// DeleteBaseImportTestsModelsComplexs deletes existing base_import.tests.models.complex records.
func (c *Client) DeleteBaseImportTestsModelsComplexs(ids []int64) error {
	return c.Delete(BaseImportTestsModelsComplexModel, ids)
}

// GetBaseImportTestsModelsComplex gets base_import.tests.models.complex existing record.
func (c *Client) GetBaseImportTestsModelsComplex(id int64) (*BaseImportTestsModelsComplex, error) {
	btmcs, err := c.GetBaseImportTestsModelsComplexs([]int64{id})
	if err != nil {
		return nil, err
	}
	if btmcs != nil && len(*btmcs) > 0 {
		return &((*btmcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base_import.tests.models.complex not found", id)
}

// GetBaseImportTestsModelsComplexs gets base_import.tests.models.complex existing records.
func (c *Client) GetBaseImportTestsModelsComplexs(ids []int64) (*BaseImportTestsModelsComplexs, error) {
	btmcs := &BaseImportTestsModelsComplexs{}
	if err := c.Read(BaseImportTestsModelsComplexModel, ids, nil, btmcs); err != nil {
		return nil, err
	}
	return btmcs, nil
}

// FindBaseImportTestsModelsComplex finds base_import.tests.models.complex record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsComplex(criteria *Criteria) (*BaseImportTestsModelsComplex, error) {
	btmcs := &BaseImportTestsModelsComplexs{}
	if err := c.SearchRead(BaseImportTestsModelsComplexModel, criteria, NewOptions().Limit(1), btmcs); err != nil {
		return nil, err
	}
	if btmcs != nil && len(*btmcs) > 0 {
		return &((*btmcs)[0]), nil
	}
	return nil, fmt.Errorf("base_import.tests.models.complex was not found")
}

// FindBaseImportTestsModelsComplexs finds base_import.tests.models.complex records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsComplexs(criteria *Criteria, options *Options) (*BaseImportTestsModelsComplexs, error) {
	btmcs := &BaseImportTestsModelsComplexs{}
	if err := c.SearchRead(BaseImportTestsModelsComplexModel, criteria, options, btmcs); err != nil {
		return nil, err
	}
	return btmcs, nil
}

// FindBaseImportTestsModelsComplexIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsComplexIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportTestsModelsComplexModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportTestsModelsComplexId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsComplexId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsComplexModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base_import.tests.models.complex was not found")
}
