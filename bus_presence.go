package odoo

import (
	"fmt"
)

// BusPresence represents bus.presence model.
type BusPresence struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	LastPoll     *Time      `xmlrpc:"last_poll,omptempty"`
	LastPresence *Time      `xmlrpc:"last_presence,omptempty"`
	Status       *Selection `xmlrpc:"status,omptempty"`
	UserId       *Many2One  `xmlrpc:"user_id,omptempty"`
}

// BusPresences represents array of bus.presence model.
type BusPresences []BusPresence

// BusPresenceModel is the odoo model name.
const BusPresenceModel = "bus.presence"

// Many2One convert BusPresence to *Many2One.
func (bp *BusPresence) Many2One() *Many2One {
	return NewMany2One(bp.Id.Get(), "")
}

// CreateBusPresence creates a new bus.presence model and returns its id.
func (c *Client) CreateBusPresence(bp *BusPresence) (int64, error) {
	ids, err := c.Create(BusPresenceModel, []interface{}{bp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBusPresence creates a new bus.presence model and returns its id.
func (c *Client) CreateBusPresences(bps []*BusPresence) ([]int64, error) {
	var vv []interface{}
	for _, v := range bps {
		vv = append(vv, v)
	}
	return c.Create(BusPresenceModel, vv)
}

// UpdateBusPresence updates an existing bus.presence record.
func (c *Client) UpdateBusPresence(bp *BusPresence) error {
	return c.UpdateBusPresences([]int64{bp.Id.Get()}, bp)
}

// UpdateBusPresences updates existing bus.presence records.
// All records (represented by ids) will be updated by bp values.
func (c *Client) UpdateBusPresences(ids []int64, bp *BusPresence) error {
	return c.Update(BusPresenceModel, ids, bp)
}

// DeleteBusPresence deletes an existing bus.presence record.
func (c *Client) DeleteBusPresence(id int64) error {
	return c.DeleteBusPresences([]int64{id})
}

// DeleteBusPresences deletes existing bus.presence records.
func (c *Client) DeleteBusPresences(ids []int64) error {
	return c.Delete(BusPresenceModel, ids)
}

// GetBusPresence gets bus.presence existing record.
func (c *Client) GetBusPresence(id int64) (*BusPresence, error) {
	bps, err := c.GetBusPresences([]int64{id})
	if err != nil {
		return nil, err
	}
	if bps != nil && len(*bps) > 0 {
		return &((*bps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of bus.presence not found", id)
}

// GetBusPresences gets bus.presence existing records.
func (c *Client) GetBusPresences(ids []int64) (*BusPresences, error) {
	bps := &BusPresences{}
	if err := c.Read(BusPresenceModel, ids, nil, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBusPresence finds bus.presence record by querying it with criteria.
func (c *Client) FindBusPresence(criteria *Criteria) (*BusPresence, error) {
	bps := &BusPresences{}
	if err := c.SearchRead(BusPresenceModel, criteria, NewOptions().Limit(1), bps); err != nil {
		return nil, err
	}
	if bps != nil && len(*bps) > 0 {
		return &((*bps)[0]), nil
	}
	return nil, fmt.Errorf("bus.presence was not found with criteria %v", criteria)
}

// FindBusPresences finds bus.presence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusPresences(criteria *Criteria, options *Options) (*BusPresences, error) {
	bps := &BusPresences{}
	if err := c.SearchRead(BusPresenceModel, criteria, options, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBusPresenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusPresenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BusPresenceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBusPresenceId finds record id by querying it with criteria.
func (c *Client) FindBusPresenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BusPresenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("bus.presence was not found with criteria %v and options %v", criteria, options)
}
