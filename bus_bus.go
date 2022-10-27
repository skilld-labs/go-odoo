package odoo

import (
	"fmt"
)

// BusBus represents bus.bus model.
type BusBus struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Channel     *String   `xmlrpc:"channel,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Message     *String   `xmlrpc:"message,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BusBuss represents array of bus.bus model.
type BusBuss []BusBus

// BusBusModel is the odoo model name.
const BusBusModel = "bus.bus"

// Many2One convert BusBus to *Many2One.
func (bb *BusBus) Many2One() *Many2One {
	return NewMany2One(bb.Id.Get(), "")
}

// CreateBusBus creates a new bus.bus model and returns its id.
func (c *Client) CreateBusBus(bb *BusBus) (int64, error) {
	return c.Create(BusBusModel, bb)
}

// UpdateBusBus updates an existing bus.bus record.
func (c *Client) UpdateBusBus(bb *BusBus) error {
	return c.UpdateBusBuss([]int64{bb.Id.Get()}, bb)
}

// UpdateBusBuss updates existing bus.bus records.
// All records (represented by ids) will be updated by bb values.
func (c *Client) UpdateBusBuss(ids []int64, bb *BusBus) error {
	return c.Update(BusBusModel, ids, bb)
}

// DeleteBusBus deletes an existing bus.bus record.
func (c *Client) DeleteBusBus(id int64) error {
	return c.DeleteBusBuss([]int64{id})
}

// DeleteBusBuss deletes existing bus.bus records.
func (c *Client) DeleteBusBuss(ids []int64) error {
	return c.Delete(BusBusModel, ids)
}

// GetBusBus gets bus.bus existing record.
func (c *Client) GetBusBus(id int64) (*BusBus, error) {
	bbs, err := c.GetBusBuss([]int64{id})
	if err != nil {
		return nil, err
	}
	if bbs != nil && len(*bbs) > 0 {
		return &((*bbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of bus.bus not found", id)
}

// GetBusBuss gets bus.bus existing records.
func (c *Client) GetBusBuss(ids []int64) (*BusBuss, error) {
	bbs := &BusBuss{}
	if err := c.Read(BusBusModel, ids, nil, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBusBus finds bus.bus record by querying it with criteria.
func (c *Client) FindBusBus(criteria *Criteria) (*BusBus, error) {
	bbs := &BusBuss{}
	if err := c.SearchRead(BusBusModel, criteria, NewOptions().Limit(1), bbs); err != nil {
		return nil, err
	}
	if bbs != nil && len(*bbs) > 0 {
		return &((*bbs)[0]), nil
	}
	return nil, fmt.Errorf("bus.bus was not found")
}

// FindBusBuss finds bus.bus records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusBuss(criteria *Criteria, options *Options) (*BusBuss, error) {
	bbs := &BusBuss{}
	if err := c.SearchRead(BusBusModel, criteria, options, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBusBusIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusBusIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BusBusModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBusBusId finds record id by querying it with criteria.
func (c *Client) FindBusBusId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BusBusModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("bus.bus was not found")
}
