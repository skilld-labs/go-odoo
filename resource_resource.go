package odoo

import (
	"fmt"
)

// ResourceResource represents resource.resource model.
type ResourceResource struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Active         *Bool      `xmlrpc:"active,omptempty"`
	CalendarId     *Many2One  `xmlrpc:"calendar_id,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	ResourceType   *Selection `xmlrpc:"resource_type,omptempty"`
	TimeEfficiency *Float     `xmlrpc:"time_efficiency,omptempty"`
	UserId         *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResourceResources represents array of resource.resource model.
type ResourceResources []ResourceResource

// ResourceResourceModel is the odoo model name.
const ResourceResourceModel = "resource.resource"

// Many2One convert ResourceResource to *Many2One.
func (rr *ResourceResource) Many2One() *Many2One {
	return NewMany2One(rr.Id.Get(), "")
}

// CreateResourceResource creates a new resource.resource model and returns its id.
func (c *Client) CreateResourceResource(rr *ResourceResource) (int64, error) {
	ids, err := c.Create(ResourceResourceModel, []interface{}{rr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceResource creates a new resource.resource model and returns its id.
func (c *Client) CreateResourceResources(rrs []*ResourceResource) ([]int64, error) {
	var vv []interface{}
	for _, v := range rrs {
		vv = append(vv, v)
	}
	return c.Create(ResourceResourceModel, vv)
}

// UpdateResourceResource updates an existing resource.resource record.
func (c *Client) UpdateResourceResource(rr *ResourceResource) error {
	return c.UpdateResourceResources([]int64{rr.Id.Get()}, rr)
}

// UpdateResourceResources updates existing resource.resource records.
// All records (represented by ids) will be updated by rr values.
func (c *Client) UpdateResourceResources(ids []int64, rr *ResourceResource) error {
	return c.Update(ResourceResourceModel, ids, rr)
}

// DeleteResourceResource deletes an existing resource.resource record.
func (c *Client) DeleteResourceResource(id int64) error {
	return c.DeleteResourceResources([]int64{id})
}

// DeleteResourceResources deletes existing resource.resource records.
func (c *Client) DeleteResourceResources(ids []int64) error {
	return c.Delete(ResourceResourceModel, ids)
}

// GetResourceResource gets resource.resource existing record.
func (c *Client) GetResourceResource(id int64) (*ResourceResource, error) {
	rrs, err := c.GetResourceResources([]int64{id})
	if err != nil {
		return nil, err
	}
	if rrs != nil && len(*rrs) > 0 {
		return &((*rrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of resource.resource not found", id)
}

// GetResourceResources gets resource.resource existing records.
func (c *Client) GetResourceResources(ids []int64) (*ResourceResources, error) {
	rrs := &ResourceResources{}
	if err := c.Read(ResourceResourceModel, ids, nil, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindResourceResource finds resource.resource record by querying it with criteria.
func (c *Client) FindResourceResource(criteria *Criteria) (*ResourceResource, error) {
	rrs := &ResourceResources{}
	if err := c.SearchRead(ResourceResourceModel, criteria, NewOptions().Limit(1), rrs); err != nil {
		return nil, err
	}
	if rrs != nil && len(*rrs) > 0 {
		return &((*rrs)[0]), nil
	}
	return nil, fmt.Errorf("resource.resource was not found with criteria %v", criteria)
}

// FindResourceResources finds resource.resource records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceResources(criteria *Criteria, options *Options) (*ResourceResources, error) {
	rrs := &ResourceResources{}
	if err := c.SearchRead(ResourceResourceModel, criteria, options, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindResourceResourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceResourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResourceResourceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResourceResourceId finds record id by querying it with criteria.
func (c *Client) FindResourceResourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceResourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("resource.resource was not found with criteria %v and options %v", criteria, options)
}
