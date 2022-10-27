package odoo

import (
	"fmt"
)

// ResPartnerTitle represents res.partner.title model.
type ResPartnerTitle struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Shortcut    *String   `xmlrpc:"shortcut,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResPartnerTitles represents array of res.partner.title model.
type ResPartnerTitles []ResPartnerTitle

// ResPartnerTitleModel is the odoo model name.
const ResPartnerTitleModel = "res.partner.title"

// Many2One convert ResPartnerTitle to *Many2One.
func (rpt *ResPartnerTitle) Many2One() *Many2One {
	return NewMany2One(rpt.Id.Get(), "")
}

// CreateResPartnerTitle creates a new res.partner.title model and returns its id.
func (c *Client) CreateResPartnerTitle(rpt *ResPartnerTitle) (int64, error) {
	return c.Create(ResPartnerTitleModel, rpt)
}

// UpdateResPartnerTitle updates an existing res.partner.title record.
func (c *Client) UpdateResPartnerTitle(rpt *ResPartnerTitle) error {
	return c.UpdateResPartnerTitles([]int64{rpt.Id.Get()}, rpt)
}

// UpdateResPartnerTitles updates existing res.partner.title records.
// All records (represented by ids) will be updated by rpt values.
func (c *Client) UpdateResPartnerTitles(ids []int64, rpt *ResPartnerTitle) error {
	return c.Update(ResPartnerTitleModel, ids, rpt)
}

// DeleteResPartnerTitle deletes an existing res.partner.title record.
func (c *Client) DeleteResPartnerTitle(id int64) error {
	return c.DeleteResPartnerTitles([]int64{id})
}

// DeleteResPartnerTitles deletes existing res.partner.title records.
func (c *Client) DeleteResPartnerTitles(ids []int64) error {
	return c.Delete(ResPartnerTitleModel, ids)
}

// GetResPartnerTitle gets res.partner.title existing record.
func (c *Client) GetResPartnerTitle(id int64) (*ResPartnerTitle, error) {
	rpts, err := c.GetResPartnerTitles([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpts != nil && len(*rpts) > 0 {
		return &((*rpts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.title not found", id)
}

// GetResPartnerTitles gets res.partner.title existing records.
func (c *Client) GetResPartnerTitles(ids []int64) (*ResPartnerTitles, error) {
	rpts := &ResPartnerTitles{}
	if err := c.Read(ResPartnerTitleModel, ids, nil, rpts); err != nil {
		return nil, err
	}
	return rpts, nil
}

// FindResPartnerTitle finds res.partner.title record by querying it with criteria.
func (c *Client) FindResPartnerTitle(criteria *Criteria) (*ResPartnerTitle, error) {
	rpts := &ResPartnerTitles{}
	if err := c.SearchRead(ResPartnerTitleModel, criteria, NewOptions().Limit(1), rpts); err != nil {
		return nil, err
	}
	if rpts != nil && len(*rpts) > 0 {
		return &((*rpts)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.title was not found")
}

// FindResPartnerTitles finds res.partner.title records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerTitles(criteria *Criteria, options *Options) (*ResPartnerTitles, error) {
	rpts := &ResPartnerTitles{}
	if err := c.SearchRead(ResPartnerTitleModel, criteria, options, rpts); err != nil {
		return nil, err
	}
	return rpts, nil
}

// FindResPartnerTitleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerTitleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerTitleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerTitleId finds record id by querying it with criteria.
func (c *Client) FindResPartnerTitleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerTitleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.title was not found")
}
