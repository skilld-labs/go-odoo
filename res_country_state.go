package odoo

import (
	"fmt"
)

// ResCountryState represents res.country.state model.
type ResCountryState struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Code        *String   `xmlrpc:"code,omptempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResCountryStates represents array of res.country.state model.
type ResCountryStates []ResCountryState

// ResCountryStateModel is the odoo model name.
const ResCountryStateModel = "res.country.state"

// Many2One convert ResCountryState to *Many2One.
func (rcs *ResCountryState) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResCountryState creates a new res.country.state model and returns its id.
func (c *Client) CreateResCountryState(rcs *ResCountryState) (int64, error) {
	return c.Create(ResCountryStateModel, rcs)
}

// UpdateResCountryState updates an existing res.country.state record.
func (c *Client) UpdateResCountryState(rcs *ResCountryState) error {
	return c.UpdateResCountryStates([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResCountryStates updates existing res.country.state records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResCountryStates(ids []int64, rcs *ResCountryState) error {
	return c.Update(ResCountryStateModel, ids, rcs)
}

// DeleteResCountryState deletes an existing res.country.state record.
func (c *Client) DeleteResCountryState(id int64) error {
	return c.DeleteResCountryStates([]int64{id})
}

// DeleteResCountryStates deletes existing res.country.state records.
func (c *Client) DeleteResCountryStates(ids []int64) error {
	return c.Delete(ResCountryStateModel, ids)
}

// GetResCountryState gets res.country.state existing record.
func (c *Client) GetResCountryState(id int64) (*ResCountryState, error) {
	rcss, err := c.GetResCountryStates([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.country.state not found", id)
}

// GetResCountryStates gets res.country.state existing records.
func (c *Client) GetResCountryStates(ids []int64) (*ResCountryStates, error) {
	rcss := &ResCountryStates{}
	if err := c.Read(ResCountryStateModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResCountryState finds res.country.state record by querying it with criteria.
func (c *Client) FindResCountryState(criteria *Criteria) (*ResCountryState, error) {
	rcss := &ResCountryStates{}
	if err := c.SearchRead(ResCountryStateModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.country.state was not found with criteria %v", criteria)
}

// FindResCountryStates finds res.country.state records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryStates(criteria *Criteria, options *Options) (*ResCountryStates, error) {
	rcss := &ResCountryStates{}
	if err := c.SearchRead(ResCountryStateModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResCountryStateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryStateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCountryStateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCountryStateId finds record id by querying it with criteria.
func (c *Client) FindResCountryStateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCountryStateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.country.state was not found with criteria %v and options %v", criteria, options)
}
