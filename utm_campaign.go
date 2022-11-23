package odoo

import (
	"fmt"
)

// UtmCampaign represents utm.campaign model.
type UtmCampaign struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	Color          *Int      `xmlrpc:"color,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId     *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	InvoicedAmount *Int      `xmlrpc:"invoiced_amount,omitempty"`
	IsWebsite      *Bool     `xmlrpc:"is_website,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	QuotationCount *Int      `xmlrpc:"quotation_count,omitempty"`
	StageId        *Many2One `xmlrpc:"stage_id,omitempty"`
	TagIds         *Relation `xmlrpc:"tag_ids,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
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
// All records (represented by IDs) will be updated by uc values.
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
	return nil, fmt.Errorf("utm.campaign was not found")
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

// FindUtmCampaignIds finds records IDs by querying it
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
	return -1, fmt.Errorf("utm.campaign was not found")
}
