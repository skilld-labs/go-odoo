package odoo

import (
	"fmt"
)

// IrFilters represents ir.filters model.
type IrFilters struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId    *Many2One  `xmlrpc:"action_id,omptempty"`
	Active      *Bool      `xmlrpc:"active,omptempty"`
	Context     *String    `xmlrpc:"context,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Domain      *String    `xmlrpc:"domain,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	IsDefault   *Bool      `xmlrpc:"is_default,omptempty"`
	ModelId     *Selection `xmlrpc:"model_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Sort        *String    `xmlrpc:"sort,omptempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrFilterss represents array of ir.filters model.
type IrFilterss []IrFilters

// IrFiltersModel is the odoo model name.
const IrFiltersModel = "ir.filters"

// Many2One convert IrFilters to *Many2One.
func (IF *IrFilters) Many2One() *Many2One {
	return NewMany2One(IF.Id.Get(), "")
}

// CreateIrFilters creates a new ir.filters model and returns its id.
func (c *Client) CreateIrFilters(IF *IrFilters) (int64, error) {
	ids, err := c.Create(IrFiltersModel, []interface{}{IF})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrFilters creates a new ir.filters model and returns its id.
func (c *Client) CreateIrFilterss(IFs []*IrFilters) ([]int64, error) {
	var vv []interface{}
	for _, v := range IFs {
		vv = append(vv, v)
	}
	return c.Create(IrFiltersModel, vv)
}

// UpdateIrFilters updates an existing ir.filters record.
func (c *Client) UpdateIrFilters(IF *IrFilters) error {
	return c.UpdateIrFilterss([]int64{IF.Id.Get()}, IF)
}

// UpdateIrFilterss updates existing ir.filters records.
// All records (represented by ids) will be updated by IF values.
func (c *Client) UpdateIrFilterss(ids []int64, IF *IrFilters) error {
	return c.Update(IrFiltersModel, ids, IF)
}

// DeleteIrFilters deletes an existing ir.filters record.
func (c *Client) DeleteIrFilters(id int64) error {
	return c.DeleteIrFilterss([]int64{id})
}

// DeleteIrFilterss deletes existing ir.filters records.
func (c *Client) DeleteIrFilterss(ids []int64) error {
	return c.Delete(IrFiltersModel, ids)
}

// GetIrFilters gets ir.filters existing record.
func (c *Client) GetIrFilters(id int64) (*IrFilters, error) {
	IFs, err := c.GetIrFilterss([]int64{id})
	if err != nil {
		return nil, err
	}
	if IFs != nil && len(*IFs) > 0 {
		return &((*IFs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.filters not found", id)
}

// GetIrFilterss gets ir.filters existing records.
func (c *Client) GetIrFilterss(ids []int64) (*IrFilterss, error) {
	IFs := &IrFilterss{}
	if err := c.Read(IrFiltersModel, ids, nil, IFs); err != nil {
		return nil, err
	}
	return IFs, nil
}

// FindIrFilters finds ir.filters record by querying it with criteria.
func (c *Client) FindIrFilters(criteria *Criteria) (*IrFilters, error) {
	IFs := &IrFilterss{}
	if err := c.SearchRead(IrFiltersModel, criteria, NewOptions().Limit(1), IFs); err != nil {
		return nil, err
	}
	if IFs != nil && len(*IFs) > 0 {
		return &((*IFs)[0]), nil
	}
	return nil, fmt.Errorf("ir.filters was not found with criteria %v", criteria)
}

// FindIrFilterss finds ir.filters records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrFilterss(criteria *Criteria, options *Options) (*IrFilterss, error) {
	IFs := &IrFilterss{}
	if err := c.SearchRead(IrFiltersModel, criteria, options, IFs); err != nil {
		return nil, err
	}
	return IFs, nil
}

// FindIrFiltersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrFiltersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrFiltersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrFiltersId finds record id by querying it with criteria.
func (c *Client) FindIrFiltersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrFiltersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.filters was not found with criteria %v and options %v", criteria, options)
}
