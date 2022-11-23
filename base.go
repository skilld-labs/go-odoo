package odoo

import (
	"fmt"
)

// Base represents base model.
type Base struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// Bases represents array of base model.
type Bases []Base

// BaseModel is the odoo model name.
const BaseModel = "base"

// Many2One convert Base to *Many2One.
func (b *Base) Many2One() *Many2One {
	return NewMany2One(b.Id.Get(), "")
}

// CreateBase creates a new base model and returns its id.
func (c *Client) CreateBase(b *Base) (int64, error) {
	return c.Create(BaseModel, b)
}

// UpdateBase updates an existing base record.
func (c *Client) UpdateBase(b *Base) error {
	return c.UpdateBases([]int64{b.Id.Get()}, b)
}

// UpdateBases updates existing base records.
// All records (represented by ids) will be updated by b values.
func (c *Client) UpdateBases(ids []int64, b *Base) error {
	return c.Update(BaseModel, ids, b)
}

// DeleteBase deletes an existing base record.
func (c *Client) DeleteBase(id int64) error {
	return c.DeleteBases([]int64{id})
}

// DeleteBases deletes existing base records.
func (c *Client) DeleteBases(ids []int64) error {
	return c.Delete(BaseModel, ids)
}

// GetBase gets base existing record.
func (c *Client) GetBase(id int64) (*Base, error) {
	bs, err := c.GetBases([]int64{id})
	if err != nil {
		return nil, err
	}
	if bs != nil && len(*bs) > 0 {
		return &((*bs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base not found", id)
}

// GetBases gets base existing records.
func (c *Client) GetBases(ids []int64) (*Bases, error) {
	bs := &Bases{}
	if err := c.Read(BaseModel, ids, nil, bs); err != nil {
		return nil, err
	}
	return bs, nil
}

// FindBase finds base record by querying it with criteria.
func (c *Client) FindBase(criteria *Criteria) (*Base, error) {
	bs := &Bases{}
	if err := c.SearchRead(BaseModel, criteria, NewOptions().Limit(1), bs); err != nil {
		return nil, err
	}
	if bs != nil && len(*bs) > 0 {
		return &((*bs)[0]), nil
	}
	return nil, fmt.Errorf("base was not found with criteria %v", criteria)
}

// FindBases finds base records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBases(criteria *Criteria, options *Options) (*Bases, error) {
	bs := &Bases{}
	if err := c.SearchRead(BaseModel, criteria, options, bs); err != nil {
		return nil, err
	}
	return bs, nil
}

// FindBaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseId finds record id by querying it with criteria.
func (c *Client) FindBaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base was not found with criteria %v and options %v", criteria, options)
}
