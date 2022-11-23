package odoo

import (
	"fmt"
)

// ResCountryGroup represents res.country.group model.
type ResCountryGroup struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omitempty"`
	CountryIds   *Relation `xmlrpc:"country_ids,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	PricelistIds *Relation `xmlrpc:"pricelist_ids,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResCountryGroups represents array of res.country.group model.
type ResCountryGroups []ResCountryGroup

// ResCountryGroupModel is the odoo model name.
const ResCountryGroupModel = "res.country.group"

// Many2One convert ResCountryGroup to *Many2One.
func (rcg *ResCountryGroup) Many2One() *Many2One {
	return NewMany2One(rcg.Id.Get(), "")
}

// CreateResCountryGroup creates a new res.country.group model and returns its id.
func (c *Client) CreateResCountryGroup(rcg *ResCountryGroup) (int64, error) {
	return c.Create(ResCountryGroupModel, rcg)
}

// UpdateResCountryGroup updates an existing res.country.group record.
func (c *Client) UpdateResCountryGroup(rcg *ResCountryGroup) error {
	return c.UpdateResCountryGroups([]int64{rcg.Id.Get()}, rcg)
}

// UpdateResCountryGroups updates existing res.country.group records.
// All records (represented by IDs) will be updated by rcg values.
func (c *Client) UpdateResCountryGroups(ids []int64, rcg *ResCountryGroup) error {
	return c.Update(ResCountryGroupModel, ids, rcg)
}

// DeleteResCountryGroup deletes an existing res.country.group record.
func (c *Client) DeleteResCountryGroup(id int64) error {
	return c.DeleteResCountryGroups([]int64{id})
}

// DeleteResCountryGroups deletes existing res.country.group records.
func (c *Client) DeleteResCountryGroups(ids []int64) error {
	return c.Delete(ResCountryGroupModel, ids)
}

// GetResCountryGroup gets res.country.group existing record.
func (c *Client) GetResCountryGroup(id int64) (*ResCountryGroup, error) {
	rcgs, err := c.GetResCountryGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcgs != nil && len(*rcgs) > 0 {
		return &((*rcgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.country.group not found", id)
}

// GetResCountryGroups gets res.country.group existing records.
func (c *Client) GetResCountryGroups(ids []int64) (*ResCountryGroups, error) {
	rcgs := &ResCountryGroups{}
	if err := c.Read(ResCountryGroupModel, ids, nil, rcgs); err != nil {
		return nil, err
	}
	return rcgs, nil
}

// FindResCountryGroup finds res.country.group record by querying it with criteria.
func (c *Client) FindResCountryGroup(criteria *Criteria) (*ResCountryGroup, error) {
	rcgs := &ResCountryGroups{}
	if err := c.SearchRead(ResCountryGroupModel, criteria, NewOptions().Limit(1), rcgs); err != nil {
		return nil, err
	}
	if rcgs != nil && len(*rcgs) > 0 {
		return &((*rcgs)[0]), nil
	}
	return nil, fmt.Errorf("res.country.group was not found")
}

// FindResCountryGroups finds res.country.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryGroups(criteria *Criteria, options *Options) (*ResCountryGroups, error) {
	rcgs := &ResCountryGroups{}
	if err := c.SearchRead(ResCountryGroupModel, criteria, options, rcgs); err != nil {
		return nil, err
	}
	return rcgs, nil
}

// FindResCountryGroupIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCountryGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCountryGroupId finds record id by querying it with criteria.
func (c *Client) FindResCountryGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCountryGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.country.group was not found")
}
