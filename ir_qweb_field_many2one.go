package odoo

import (
	"fmt"
)

// IrQwebFieldMany2One represents ir.qweb.field.many2one model.
type IrQwebFieldMany2One struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrQwebFieldMany2Ones represents array of ir.qweb.field.many2one model.
type IrQwebFieldMany2Ones []IrQwebFieldMany2One

// IrQwebFieldMany2OneModel is the odoo model name.
const IrQwebFieldMany2OneModel = "ir.qweb.field.many2one"

// Many2One convert IrQwebFieldMany2One to *Many2One.
func (iqfm *IrQwebFieldMany2One) Many2One() *Many2One {
	return NewMany2One(iqfm.Id.Get(), "")
}

// CreateIrQwebFieldMany2One creates a new ir.qweb.field.many2one model and returns its id.
func (c *Client) CreateIrQwebFieldMany2One(iqfm *IrQwebFieldMany2One) (int64, error) {
	ids, err := c.Create(IrQwebFieldMany2OneModel, []interface{}{iqfm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldMany2One creates a new ir.qweb.field.many2one model and returns its id.
func (c *Client) CreateIrQwebFieldMany2Ones(iqfms []*IrQwebFieldMany2One) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfms {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldMany2OneModel, vv)
}

// UpdateIrQwebFieldMany2One updates an existing ir.qweb.field.many2one record.
func (c *Client) UpdateIrQwebFieldMany2One(iqfm *IrQwebFieldMany2One) error {
	return c.UpdateIrQwebFieldMany2Ones([]int64{iqfm.Id.Get()}, iqfm)
}

// UpdateIrQwebFieldMany2Ones updates existing ir.qweb.field.many2one records.
// All records (represented by ids) will be updated by iqfm values.
func (c *Client) UpdateIrQwebFieldMany2Ones(ids []int64, iqfm *IrQwebFieldMany2One) error {
	return c.Update(IrQwebFieldMany2OneModel, ids, iqfm)
}

// DeleteIrQwebFieldMany2One deletes an existing ir.qweb.field.many2one record.
func (c *Client) DeleteIrQwebFieldMany2One(id int64) error {
	return c.DeleteIrQwebFieldMany2Ones([]int64{id})
}

// DeleteIrQwebFieldMany2Ones deletes existing ir.qweb.field.many2one records.
func (c *Client) DeleteIrQwebFieldMany2Ones(ids []int64) error {
	return c.Delete(IrQwebFieldMany2OneModel, ids)
}

// GetIrQwebFieldMany2One gets ir.qweb.field.many2one existing record.
func (c *Client) GetIrQwebFieldMany2One(id int64) (*IrQwebFieldMany2One, error) {
	iqfms, err := c.GetIrQwebFieldMany2Ones([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.many2one not found", id)
}

// GetIrQwebFieldMany2Ones gets ir.qweb.field.many2one existing records.
func (c *Client) GetIrQwebFieldMany2Ones(ids []int64) (*IrQwebFieldMany2Ones, error) {
	iqfms := &IrQwebFieldMany2Ones{}
	if err := c.Read(IrQwebFieldMany2OneModel, ids, nil, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMany2One finds ir.qweb.field.many2one record by querying it with criteria.
func (c *Client) FindIrQwebFieldMany2One(criteria *Criteria) (*IrQwebFieldMany2One, error) {
	iqfms := &IrQwebFieldMany2Ones{}
	if err := c.SearchRead(IrQwebFieldMany2OneModel, criteria, NewOptions().Limit(1), iqfms); err != nil {
		return nil, err
	}
	if iqfms != nil && len(*iqfms) > 0 {
		return &((*iqfms)[0]), nil
	}
	return nil, fmt.Errorf("ir.qweb.field.many2one was not found with criteria %v", criteria)
}

// FindIrQwebFieldMany2Ones finds ir.qweb.field.many2one records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMany2Ones(criteria *Criteria, options *Options) (*IrQwebFieldMany2Ones, error) {
	iqfms := &IrQwebFieldMany2Ones{}
	if err := c.SearchRead(IrQwebFieldMany2OneModel, criteria, options, iqfms); err != nil {
		return nil, err
	}
	return iqfms, nil
}

// FindIrQwebFieldMany2OneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldMany2OneIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldMany2OneModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldMany2OneId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldMany2OneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldMany2OneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.qweb.field.many2one was not found with criteria %v and options %v", criteria, options)
}
