package odoo

import (
	"fmt"
)

// ResPartnerIndustry represents res.partner.industry model.
type ResPartnerIndustry struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FullName    *String   `xmlrpc:"full_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResPartnerIndustrys represents array of res.partner.industry model.
type ResPartnerIndustrys []ResPartnerIndustry

// ResPartnerIndustryModel is the odoo model name.
const ResPartnerIndustryModel = "res.partner.industry"

// Many2One convert ResPartnerIndustry to *Many2One.
func (rpi *ResPartnerIndustry) Many2One() *Many2One {
	return NewMany2One(rpi.Id.Get(), "")
}

// CreateResPartnerIndustry creates a new res.partner.industry model and returns its id.
func (c *Client) CreateResPartnerIndustry(rpi *ResPartnerIndustry) (int64, error) {
	return c.Create(ResPartnerIndustryModel, rpi)
}

// UpdateResPartnerIndustry updates an existing res.partner.industry record.
func (c *Client) UpdateResPartnerIndustry(rpi *ResPartnerIndustry) error {
	return c.UpdateResPartnerIndustrys([]int64{rpi.Id.Get()}, rpi)
}

// UpdateResPartnerIndustrys updates existing res.partner.industry records.
// All records (represented by ids) will be updated by rpi values.
func (c *Client) UpdateResPartnerIndustrys(ids []int64, rpi *ResPartnerIndustry) error {
	return c.Update(ResPartnerIndustryModel, ids, rpi)
}

// DeleteResPartnerIndustry deletes an existing res.partner.industry record.
func (c *Client) DeleteResPartnerIndustry(id int64) error {
	return c.DeleteResPartnerIndustrys([]int64{id})
}

// DeleteResPartnerIndustrys deletes existing res.partner.industry records.
func (c *Client) DeleteResPartnerIndustrys(ids []int64) error {
	return c.Delete(ResPartnerIndustryModel, ids)
}

// GetResPartnerIndustry gets res.partner.industry existing record.
func (c *Client) GetResPartnerIndustry(id int64) (*ResPartnerIndustry, error) {
	rpis, err := c.GetResPartnerIndustrys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpis != nil && len(*rpis) > 0 {
		return &((*rpis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.industry not found", id)
}

// GetResPartnerIndustrys gets res.partner.industry existing records.
func (c *Client) GetResPartnerIndustrys(ids []int64) (*ResPartnerIndustrys, error) {
	rpis := &ResPartnerIndustrys{}
	if err := c.Read(ResPartnerIndustryModel, ids, nil, rpis); err != nil {
		return nil, err
	}
	return rpis, nil
}

// FindResPartnerIndustry finds res.partner.industry record by querying it with criteria.
func (c *Client) FindResPartnerIndustry(criteria *Criteria) (*ResPartnerIndustry, error) {
	rpis := &ResPartnerIndustrys{}
	if err := c.SearchRead(ResPartnerIndustryModel, criteria, NewOptions().Limit(1), rpis); err != nil {
		return nil, err
	}
	if rpis != nil && len(*rpis) > 0 {
		return &((*rpis)[0]), nil
	}
	return nil, fmt.Errorf("no res.partner.industry was found with criteria %v", criteria)
}

// FindResPartnerIndustrys finds res.partner.industry records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIndustrys(criteria *Criteria, options *Options) (*ResPartnerIndustrys, error) {
	rpis := &ResPartnerIndustrys{}
	if err := c.SearchRead(ResPartnerIndustryModel, criteria, options, rpis); err != nil {
		return nil, err
	}
	return rpis, nil
}

// FindResPartnerIndustryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIndustryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerIndustryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerIndustryId finds record id by querying it with criteria.
func (c *Client) FindResPartnerIndustryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerIndustryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no res.partner.industry was found with criteria %v and options %v", criteria, options)
}
