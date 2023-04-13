package odoo

import (
	"fmt"
)

// ResCountry represents res.country model.
type ResCountry struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	AddressFormat   *String    `xmlrpc:"address_format,omptempty"`
	AddressViewId   *Many2One  `xmlrpc:"address_view_id,omptempty"`
	Code            *String    `xmlrpc:"code,omptempty"`
	CountryGroupIds *Relation  `xmlrpc:"country_group_ids,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Image           *String    `xmlrpc:"image,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	NamePosition    *Selection `xmlrpc:"name_position,omptempty"`
	PhoneCode       *Int       `xmlrpc:"phone_code,omptempty"`
	StateIds        *Relation  `xmlrpc:"state_ids,omptempty"`
	VatLabel        *String    `xmlrpc:"vat_label,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResCountrys represents array of res.country model.
type ResCountrys []ResCountry

// ResCountryModel is the odoo model name.
const ResCountryModel = "res.country"

// Many2One convert ResCountry to *Many2One.
func (rc *ResCountry) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCountry creates a new res.country model and returns its id.
func (c *Client) CreateResCountry(rc *ResCountry) (int64, error) {
	ids, err := c.Create(ResCountryModel, []interface{}{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCountry creates a new res.country model and returns its id.
func (c *Client) CreateResCountrys(rcs []*ResCountry) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCountryModel, vv)
}

// UpdateResCountry updates an existing res.country record.
func (c *Client) UpdateResCountry(rc *ResCountry) error {
	return c.UpdateResCountrys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCountrys updates existing res.country records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCountrys(ids []int64, rc *ResCountry) error {
	return c.Update(ResCountryModel, ids, rc)
}

// DeleteResCountry deletes an existing res.country record.
func (c *Client) DeleteResCountry(id int64) error {
	return c.DeleteResCountrys([]int64{id})
}

// DeleteResCountrys deletes existing res.country records.
func (c *Client) DeleteResCountrys(ids []int64) error {
	return c.Delete(ResCountryModel, ids)
}

// GetResCountry gets res.country existing record.
func (c *Client) GetResCountry(id int64) (*ResCountry, error) {
	rcs, err := c.GetResCountrys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.country not found", id)
}

// GetResCountrys gets res.country existing records.
func (c *Client) GetResCountrys(ids []int64) (*ResCountrys, error) {
	rcs := &ResCountrys{}
	if err := c.Read(ResCountryModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCountry finds res.country record by querying it with criteria.
func (c *Client) FindResCountry(criteria *Criteria) (*ResCountry, error) {
	rcs := &ResCountrys{}
	if err := c.SearchRead(ResCountryModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.country was not found with criteria %v", criteria)
}

// FindResCountrys finds res.country records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountrys(criteria *Criteria, options *Options) (*ResCountrys, error) {
	rcs := &ResCountrys{}
	if err := c.SearchRead(ResCountryModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCountryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCountryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCountryId finds record id by querying it with criteria.
func (c *Client) FindResCountryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCountryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.country was not found with criteria %v and options %v", criteria, options)
}
