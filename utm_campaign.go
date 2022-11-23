package odoo

import (
	"fmt"
)

// UtmCampaign represents utm.campaign model.
type UtmCampaign struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UtmCampaigns represents array of utm.campaign model.
type UtmCampaigns []UtmCampaign

// UtmCampaignModel is the odoo model name.
const UtmCampaignModel = "utm.campaign"

// Many2One convert UtmCampaign to *Many2One.
func (uc *UtmCampaign) Many2One() *Many2One {
	return NewMany2One(uc.Id.Get(), "")
}

// CreateUtmCampaign creates a new utm.campaign model and returns its id.
func (c *Client) CreateUtmCampaign(uc *UtmCampaign) (int64, error) {
	return c.Create(UtmCampaignModel, uc)
}

// UpdateUtmCampaign updates an existing utm.campaign record.
func (c *Client) UpdateUtmCampaign(uc *UtmCampaign) error {
	return c.UpdateUtmCampaigns([]int64{uc.Id.Get()}, uc)
}

// UpdateUtmCampaigns updates existing utm.campaign records.
// All records (represented by ids) will be updated by uc values.
func (c *Client) UpdateUtmCampaigns(ids []int64, uc *UtmCampaign) error {
	return c.Update(UtmCampaignModel, ids, uc)
}

// DeleteUtmCampaign deletes an existing utm.campaign record.
func (c *Client) DeleteUtmCampaign(id int64) error {
	return c.DeleteUtmCampaigns([]int64{id})
}

// DeleteUtmCampaigns deletes existing utm.campaign records.
func (c *Client) DeleteUtmCampaigns(ids []int64) error {
	return c.Delete(UtmCampaignModel, ids)
}

// GetUtmCampaign gets utm.campaign existing record.
func (c *Client) GetUtmCampaign(id int64) (*UtmCampaign, error) {
	ucs, err := c.GetUtmCampaigns([]int64{id})
	if err != nil {
		return nil, err
	}
	if ucs != nil && len(*ucs) > 0 {
		return &((*ucs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of utm.campaign not found", id)
}

// GetUtmCampaigns gets utm.campaign existing records.
func (c *Client) GetUtmCampaigns(ids []int64) (*UtmCampaigns, error) {
	ucs := &UtmCampaigns{}
	if err := c.Read(UtmCampaignModel, ids, nil, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUtmCampaign finds utm.campaign record by querying it with criteria.
func (c *Client) FindUtmCampaign(criteria *Criteria) (*UtmCampaign, error) {
	ucs := &UtmCampaigns{}
	if err := c.SearchRead(UtmCampaignModel, criteria, NewOptions().Limit(1), ucs); err != nil {
		return nil, err
	}
	if ucs != nil && len(*ucs) > 0 {
		return &((*ucs)[0]), nil
	}
	return nil, fmt.Errorf("utm.campaign was not found with criteria %v", criteria)
}

// FindUtmCampaigns finds utm.campaign records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmCampaigns(criteria *Criteria, options *Options) (*UtmCampaigns, error) {
	ucs := &UtmCampaigns{}
	if err := c.SearchRead(UtmCampaignModel, criteria, options, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUtmCampaignIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmCampaignIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UtmCampaignModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUtmCampaignId finds record id by querying it with criteria.
func (c *Client) FindUtmCampaignId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmCampaignModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("utm.campaign was not found with criteria %v and options %v", criteria, options)
}
