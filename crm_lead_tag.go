package odoo

import (
	"fmt"
)

// CrmLeadTag represents crm.lead.tag model.
type CrmLeadTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmLeadTags represents array of crm.lead.tag model.
type CrmLeadTags []CrmLeadTag

// CrmLeadTagModel is the odoo model name.
const CrmLeadTagModel = "crm.lead.tag"

// Many2One convert CrmLeadTag to *Many2One.
func (clt *CrmLeadTag) Many2One() *Many2One {
	return NewMany2One(clt.Id.Get(), "")
}

// CreateCrmLeadTag creates a new crm.lead.tag model and returns its id.
func (c *Client) CreateCrmLeadTag(clt *CrmLeadTag) (int64, error) {
	return c.Create(CrmLeadTagModel, clt)
}

// UpdateCrmLeadTag updates an existing crm.lead.tag record.
func (c *Client) UpdateCrmLeadTag(clt *CrmLeadTag) error {
	return c.UpdateCrmLeadTags([]int64{clt.Id.Get()}, clt)
}

// UpdateCrmLeadTags updates existing crm.lead.tag records.
// All records (represented by ids) will be updated by clt values.
func (c *Client) UpdateCrmLeadTags(ids []int64, clt *CrmLeadTag) error {
	return c.Update(CrmLeadTagModel, ids, clt)
}

// DeleteCrmLeadTag deletes an existing crm.lead.tag record.
func (c *Client) DeleteCrmLeadTag(id int64) error {
	return c.DeleteCrmLeadTags([]int64{id})
}

// DeleteCrmLeadTags deletes existing crm.lead.tag records.
func (c *Client) DeleteCrmLeadTags(ids []int64) error {
	return c.Delete(CrmLeadTagModel, ids)
}

// GetCrmLeadTag gets crm.lead.tag existing record.
func (c *Client) GetCrmLeadTag(id int64) (*CrmLeadTag, error) {
	clts, err := c.GetCrmLeadTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if clts != nil && len(*clts) > 0 {
		return &((*clts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead.tag not found", id)
}

// GetCrmLeadTags gets crm.lead.tag existing records.
func (c *Client) GetCrmLeadTags(ids []int64) (*CrmLeadTags, error) {
	clts := &CrmLeadTags{}
	if err := c.Read(CrmLeadTagModel, ids, nil, clts); err != nil {
		return nil, err
	}
	return clts, nil
}

// FindCrmLeadTag finds crm.lead.tag record by querying it with criteria.
func (c *Client) FindCrmLeadTag(criteria *Criteria) (*CrmLeadTag, error) {
	clts := &CrmLeadTags{}
	if err := c.SearchRead(CrmLeadTagModel, criteria, NewOptions().Limit(1), clts); err != nil {
		return nil, err
	}
	if clts != nil && len(*clts) > 0 {
		return &((*clts)[0]), nil
	}
	return nil, fmt.Errorf("no crm.lead.tag was found with criteria %v", criteria)
}

// FindCrmLeadTags finds crm.lead.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadTags(criteria *Criteria, options *Options) (*CrmLeadTags, error) {
	clts := &CrmLeadTags{}
	if err := c.SearchRead(CrmLeadTagModel, criteria, options, clts); err != nil {
		return nil, err
	}
	return clts, nil
}

// FindCrmLeadTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadTagId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no crm.lead.tag was found with criteria %v and options %v", criteria, options)
}
