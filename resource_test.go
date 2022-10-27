package odoo

import (
	"fmt"
)

// ResourceTest represents resource.test model.
type ResourceTest struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	ResourceCalendarId *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId         *Many2One  `xmlrpc:"resource_id,omitempty"`
	Tz                 *Selection `xmlrpc:"tz,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResourceTests represents array of resource.test model.
type ResourceTests []ResourceTest

// ResourceTestModel is the odoo model name.
const ResourceTestModel = "resource.test"

// Many2One convert ResourceTest to *Many2One.
func (rt *ResourceTest) Many2One() *Many2One {
	return NewMany2One(rt.Id.Get(), "")
}

// CreateResourceTest creates a new resource.test model and returns its id.
func (c *Client) CreateResourceTest(rt *ResourceTest) (int64, error) {
	return c.Create(ResourceTestModel, rt)
}

// UpdateResourceTest updates an existing resource.test record.
func (c *Client) UpdateResourceTest(rt *ResourceTest) error {
	return c.UpdateResourceTests([]int64{rt.Id.Get()}, rt)
}

// UpdateResourceTests updates existing resource.test records.
// All records (represented by ids) will be updated by rt values.
func (c *Client) UpdateResourceTests(ids []int64, rt *ResourceTest) error {
	return c.Update(ResourceTestModel, ids, rt)
}

// DeleteResourceTest deletes an existing resource.test record.
func (c *Client) DeleteResourceTest(id int64) error {
	return c.DeleteResourceTests([]int64{id})
}

// DeleteResourceTests deletes existing resource.test records.
func (c *Client) DeleteResourceTests(ids []int64) error {
	return c.Delete(ResourceTestModel, ids)
}

// GetResourceTest gets resource.test existing record.
func (c *Client) GetResourceTest(id int64) (*ResourceTest, error) {
	rts, err := c.GetResourceTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if rts != nil && len(*rts) > 0 {
		return &((*rts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of resource.test not found", id)
}

// GetResourceTests gets resource.test existing records.
func (c *Client) GetResourceTests(ids []int64) (*ResourceTests, error) {
	rts := &ResourceTests{}
	if err := c.Read(ResourceTestModel, ids, nil, rts); err != nil {
		return nil, err
	}
	return rts, nil
}

// FindResourceTest finds resource.test record by querying it with criteria.
func (c *Client) FindResourceTest(criteria *Criteria) (*ResourceTest, error) {
	rts := &ResourceTests{}
	if err := c.SearchRead(ResourceTestModel, criteria, NewOptions().Limit(1), rts); err != nil {
		return nil, err
	}
	if rts != nil && len(*rts) > 0 {
		return &((*rts)[0]), nil
	}
	return nil, fmt.Errorf("resource.test was not found")
}

// FindResourceTests finds resource.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceTests(criteria *Criteria, options *Options) (*ResourceTests, error) {
	rts := &ResourceTests{}
	if err := c.SearchRead(ResourceTestModel, criteria, options, rts); err != nil {
		return nil, err
	}
	return rts, nil
}

// FindResourceTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResourceTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResourceTestId finds record id by querying it with criteria.
func (c *Client) FindResourceTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("resource.test was not found")
}
