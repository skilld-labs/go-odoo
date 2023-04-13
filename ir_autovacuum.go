package odoo

import (
	"fmt"
)

// IrAutovacuum represents ir.autovacuum model.
type IrAutovacuum struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrAutovacuums represents array of ir.autovacuum model.
type IrAutovacuums []IrAutovacuum

// IrAutovacuumModel is the odoo model name.
const IrAutovacuumModel = "ir.autovacuum"

// Many2One convert IrAutovacuum to *Many2One.
func (ia *IrAutovacuum) Many2One() *Many2One {
	return NewMany2One(ia.Id.Get(), "")
}

// CreateIrAutovacuum creates a new ir.autovacuum model and returns its id.
func (c *Client) CreateIrAutovacuum(ia *IrAutovacuum) (int64, error) {
	ids, err := c.CreateIrAutovacuums([]*IrAutovacuum{ia})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrAutovacuum creates a new ir.autovacuum model and returns its id.
func (c *Client) CreateIrAutovacuums(ias []*IrAutovacuum) ([]int64, error) {
	var vv []interface{}
	for _, v := range ias {
		vv = append(vv, v)
	}
	return c.Create(IrAutovacuumModel, vv)
}

// UpdateIrAutovacuum updates an existing ir.autovacuum record.
func (c *Client) UpdateIrAutovacuum(ia *IrAutovacuum) error {
	return c.UpdateIrAutovacuums([]int64{ia.Id.Get()}, ia)
}

// UpdateIrAutovacuums updates existing ir.autovacuum records.
// All records (represented by ids) will be updated by ia values.
func (c *Client) UpdateIrAutovacuums(ids []int64, ia *IrAutovacuum) error {
	return c.Update(IrAutovacuumModel, ids, ia)
}

// DeleteIrAutovacuum deletes an existing ir.autovacuum record.
func (c *Client) DeleteIrAutovacuum(id int64) error {
	return c.DeleteIrAutovacuums([]int64{id})
}

// DeleteIrAutovacuums deletes existing ir.autovacuum records.
func (c *Client) DeleteIrAutovacuums(ids []int64) error {
	return c.Delete(IrAutovacuumModel, ids)
}

// GetIrAutovacuum gets ir.autovacuum existing record.
func (c *Client) GetIrAutovacuum(id int64) (*IrAutovacuum, error) {
	ias, err := c.GetIrAutovacuums([]int64{id})
	if err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.autovacuum not found", id)
}

// GetIrAutovacuums gets ir.autovacuum existing records.
func (c *Client) GetIrAutovacuums(ids []int64) (*IrAutovacuums, error) {
	ias := &IrAutovacuums{}
	if err := c.Read(IrAutovacuumModel, ids, nil, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAutovacuum finds ir.autovacuum record by querying it with criteria.
func (c *Client) FindIrAutovacuum(criteria *Criteria) (*IrAutovacuum, error) {
	ias := &IrAutovacuums{}
	if err := c.SearchRead(IrAutovacuumModel, criteria, NewOptions().Limit(1), ias); err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("ir.autovacuum was not found with criteria %v", criteria)
}

// FindIrAutovacuums finds ir.autovacuum records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAutovacuums(criteria *Criteria, options *Options) (*IrAutovacuums, error) {
	ias := &IrAutovacuums{}
	if err := c.SearchRead(IrAutovacuumModel, criteria, options, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAutovacuumIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAutovacuumIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrAutovacuumModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrAutovacuumId finds record id by querying it with criteria.
func (c *Client) FindIrAutovacuumId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrAutovacuumModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.autovacuum was not found with criteria %v and options %v", criteria, options)
}
