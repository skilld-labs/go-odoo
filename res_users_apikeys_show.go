package odoo

import (
	"fmt"
)

// ResUsersApikeysShow represents res.users.apikeys.show model.
type ResUsersApikeysShow struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	Key         *String `xmlrpc:"key,omitempty"`
}

// ResUsersApikeysShows represents array of res.users.apikeys.show model.
type ResUsersApikeysShows []ResUsersApikeysShow

// ResUsersApikeysShowModel is the odoo model name.
const ResUsersApikeysShowModel = "res.users.apikeys.show"

// Many2One convert ResUsersApikeysShow to *Many2One.
func (ruas *ResUsersApikeysShow) Many2One() *Many2One {
	return NewMany2One(ruas.Id.Get(), "")
}

// CreateResUsersApikeysShow creates a new res.users.apikeys.show model and returns its id.
func (c *Client) CreateResUsersApikeysShow(ruas *ResUsersApikeysShow) (int64, error) {
	return c.Create(ResUsersApikeysShowModel, ruas)
}

// UpdateResUsersApikeysShow updates an existing res.users.apikeys.show record.
func (c *Client) UpdateResUsersApikeysShow(ruas *ResUsersApikeysShow) error {
	return c.UpdateResUsersApikeysShows([]int64{ruas.Id.Get()}, ruas)
}

// UpdateResUsersApikeysShows updates existing res.users.apikeys.show records.
// All records (represented by ids) will be updated by ruas values.
func (c *Client) UpdateResUsersApikeysShows(ids []int64, ruas *ResUsersApikeysShow) error {
	return c.Update(ResUsersApikeysShowModel, ids, ruas)
}

// DeleteResUsersApikeysShow deletes an existing res.users.apikeys.show record.
func (c *Client) DeleteResUsersApikeysShow(id int64) error {
	return c.DeleteResUsersApikeysShows([]int64{id})
}

// DeleteResUsersApikeysShows deletes existing res.users.apikeys.show records.
func (c *Client) DeleteResUsersApikeysShows(ids []int64) error {
	return c.Delete(ResUsersApikeysShowModel, ids)
}

// GetResUsersApikeysShow gets res.users.apikeys.show existing record.
func (c *Client) GetResUsersApikeysShow(id int64) (*ResUsersApikeysShow, error) {
	ruass, err := c.GetResUsersApikeysShows([]int64{id})
	if err != nil {
		return nil, err
	}
	if ruass != nil && len(*ruass) > 0 {
		return &((*ruass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.apikeys.show not found", id)
}

// GetResUsersApikeysShows gets res.users.apikeys.show existing records.
func (c *Client) GetResUsersApikeysShows(ids []int64) (*ResUsersApikeysShows, error) {
	ruass := &ResUsersApikeysShows{}
	if err := c.Read(ResUsersApikeysShowModel, ids, nil, ruass); err != nil {
		return nil, err
	}
	return ruass, nil
}

// FindResUsersApikeysShow finds res.users.apikeys.show record by querying it with criteria.
func (c *Client) FindResUsersApikeysShow(criteria *Criteria) (*ResUsersApikeysShow, error) {
	ruass := &ResUsersApikeysShows{}
	if err := c.SearchRead(ResUsersApikeysShowModel, criteria, NewOptions().Limit(1), ruass); err != nil {
		return nil, err
	}
	if ruass != nil && len(*ruass) > 0 {
		return &((*ruass)[0]), nil
	}
	return nil, fmt.Errorf("res.users.apikeys.show was not found")
}

// FindResUsersApikeysShows finds res.users.apikeys.show records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeysShows(criteria *Criteria, options *Options) (*ResUsersApikeysShows, error) {
	ruass := &ResUsersApikeysShows{}
	if err := c.SearchRead(ResUsersApikeysShowModel, criteria, options, ruass); err != nil {
		return nil, err
	}
	return ruass, nil
}

// FindResUsersApikeysShowIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeysShowIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersApikeysShowModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersApikeysShowId finds record id by querying it with criteria.
func (c *Client) FindResUsersApikeysShowId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersApikeysShowModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.apikeys.show was not found")
}
