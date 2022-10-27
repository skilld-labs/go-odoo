package odoo

import (
	"fmt"
)

// ResUsersApikeysDescription represents res.users.apikeys.description model.
type ResUsersApikeysDescription struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResUsersApikeysDescriptions represents array of res.users.apikeys.description model.
type ResUsersApikeysDescriptions []ResUsersApikeysDescription

// ResUsersApikeysDescriptionModel is the odoo model name.
const ResUsersApikeysDescriptionModel = "res.users.apikeys.description"

// Many2One convert ResUsersApikeysDescription to *Many2One.
func (ruad *ResUsersApikeysDescription) Many2One() *Many2One {
	return NewMany2One(ruad.Id.Get(), "")
}

// CreateResUsersApikeysDescription creates a new res.users.apikeys.description model and returns its id.
func (c *Client) CreateResUsersApikeysDescription(ruad *ResUsersApikeysDescription) (int64, error) {
	return c.Create(ResUsersApikeysDescriptionModel, ruad)
}

// UpdateResUsersApikeysDescription updates an existing res.users.apikeys.description record.
func (c *Client) UpdateResUsersApikeysDescription(ruad *ResUsersApikeysDescription) error {
	return c.UpdateResUsersApikeysDescriptions([]int64{ruad.Id.Get()}, ruad)
}

// UpdateResUsersApikeysDescriptions updates existing res.users.apikeys.description records.
// All records (represented by ids) will be updated by ruad values.
func (c *Client) UpdateResUsersApikeysDescriptions(ids []int64, ruad *ResUsersApikeysDescription) error {
	return c.Update(ResUsersApikeysDescriptionModel, ids, ruad)
}

// DeleteResUsersApikeysDescription deletes an existing res.users.apikeys.description record.
func (c *Client) DeleteResUsersApikeysDescription(id int64) error {
	return c.DeleteResUsersApikeysDescriptions([]int64{id})
}

// DeleteResUsersApikeysDescriptions deletes existing res.users.apikeys.description records.
func (c *Client) DeleteResUsersApikeysDescriptions(ids []int64) error {
	return c.Delete(ResUsersApikeysDescriptionModel, ids)
}

// GetResUsersApikeysDescription gets res.users.apikeys.description existing record.
func (c *Client) GetResUsersApikeysDescription(id int64) (*ResUsersApikeysDescription, error) {
	ruads, err := c.GetResUsersApikeysDescriptions([]int64{id})
	if err != nil {
		return nil, err
	}
	if ruads != nil && len(*ruads) > 0 {
		return &((*ruads)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.apikeys.description not found", id)
}

// GetResUsersApikeysDescriptions gets res.users.apikeys.description existing records.
func (c *Client) GetResUsersApikeysDescriptions(ids []int64) (*ResUsersApikeysDescriptions, error) {
	ruads := &ResUsersApikeysDescriptions{}
	if err := c.Read(ResUsersApikeysDescriptionModel, ids, nil, ruads); err != nil {
		return nil, err
	}
	return ruads, nil
}

// FindResUsersApikeysDescription finds res.users.apikeys.description record by querying it with criteria.
func (c *Client) FindResUsersApikeysDescription(criteria *Criteria) (*ResUsersApikeysDescription, error) {
	ruads := &ResUsersApikeysDescriptions{}
	if err := c.SearchRead(ResUsersApikeysDescriptionModel, criteria, NewOptions().Limit(1), ruads); err != nil {
		return nil, err
	}
	if ruads != nil && len(*ruads) > 0 {
		return &((*ruads)[0]), nil
	}
	return nil, fmt.Errorf("res.users.apikeys.description was not found")
}

// FindResUsersApikeysDescriptions finds res.users.apikeys.description records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeysDescriptions(criteria *Criteria, options *Options) (*ResUsersApikeysDescriptions, error) {
	ruads := &ResUsersApikeysDescriptions{}
	if err := c.SearchRead(ResUsersApikeysDescriptionModel, criteria, options, ruads); err != nil {
		return nil, err
	}
	return ruads, nil
}

// FindResUsersApikeysDescriptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeysDescriptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersApikeysDescriptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersApikeysDescriptionId finds record id by querying it with criteria.
func (c *Client) FindResUsersApikeysDescriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersApikeysDescriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.apikeys.description was not found")
}
