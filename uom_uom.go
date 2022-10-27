package odoo

import (
	"fmt"
)

// UomUom represents uom.uom model.
type UomUom struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	Active      *Bool      `xmlrpc:"active,omitempty"`
	CategoryId  *Many2One  `xmlrpc:"category_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Factor      *Float     `xmlrpc:"factor,omitempty"`
	FactorInv   *Float     `xmlrpc:"factor_inv,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Rounding    *Float     `xmlrpc:"rounding,omitempty"`
	UomType     *Selection `xmlrpc:"uom_type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// UomUoms represents array of uom.uom model.
type UomUoms []UomUom

// UomUomModel is the odoo model name.
const UomUomModel = "uom.uom"

// Many2One convert UomUom to *Many2One.
func (uu *UomUom) Many2One() *Many2One {
	return NewMany2One(uu.Id.Get(), "")
}

// CreateUomUom creates a new uom.uom model and returns its id.
func (c *Client) CreateUomUom(uu *UomUom) (int64, error) {
	return c.Create(UomUomModel, uu)
}

// UpdateUomUom updates an existing uom.uom record.
func (c *Client) UpdateUomUom(uu *UomUom) error {
	return c.UpdateUomUoms([]int64{uu.Id.Get()}, uu)
}

// UpdateUomUoms updates existing uom.uom records.
// All records (represented by ids) will be updated by uu values.
func (c *Client) UpdateUomUoms(ids []int64, uu *UomUom) error {
	return c.Update(UomUomModel, ids, uu)
}

// DeleteUomUom deletes an existing uom.uom record.
func (c *Client) DeleteUomUom(id int64) error {
	return c.DeleteUomUoms([]int64{id})
}

// DeleteUomUoms deletes existing uom.uom records.
func (c *Client) DeleteUomUoms(ids []int64) error {
	return c.Delete(UomUomModel, ids)
}

// GetUomUom gets uom.uom existing record.
func (c *Client) GetUomUom(id int64) (*UomUom, error) {
	uus, err := c.GetUomUoms([]int64{id})
	if err != nil {
		return nil, err
	}
	if uus != nil && len(*uus) > 0 {
		return &((*uus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of uom.uom not found", id)
}

// GetUomUoms gets uom.uom existing records.
func (c *Client) GetUomUoms(ids []int64) (*UomUoms, error) {
	uus := &UomUoms{}
	if err := c.Read(UomUomModel, ids, nil, uus); err != nil {
		return nil, err
	}
	return uus, nil
}

// FindUomUom finds uom.uom record by querying it with criteria.
func (c *Client) FindUomUom(criteria *Criteria) (*UomUom, error) {
	uus := &UomUoms{}
	if err := c.SearchRead(UomUomModel, criteria, NewOptions().Limit(1), uus); err != nil {
		return nil, err
	}
	if uus != nil && len(*uus) > 0 {
		return &((*uus)[0]), nil
	}
	return nil, fmt.Errorf("uom.uom was not found")
}

// FindUomUoms finds uom.uom records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomUoms(criteria *Criteria, options *Options) (*UomUoms, error) {
	uus := &UomUoms{}
	if err := c.SearchRead(UomUomModel, criteria, options, uus); err != nil {
		return nil, err
	}
	return uus, nil
}

// FindUomUomIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomUomIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UomUomModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUomUomId finds record id by querying it with criteria.
func (c *Client) FindUomUomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UomUomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("uom.uom was not found")
}
