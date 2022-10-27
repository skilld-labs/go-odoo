package odoo

import (
	"fmt"
)

// ResourceCalendarLeaves represents resource.calendar.leaves model.
type ResourceCalendarLeaves struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CalendarId  *Many2One  `xmlrpc:"calendar_id,omitempty"`
	CompanyId   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo      *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	ResourceId  *Many2One  `xmlrpc:"resource_id,omitempty"`
	TimeType    *Selection `xmlrpc:"time_type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResourceCalendarLeavess represents array of resource.calendar.leaves model.
type ResourceCalendarLeavess []ResourceCalendarLeaves

// ResourceCalendarLeavesModel is the odoo model name.
const ResourceCalendarLeavesModel = "resource.calendar.leaves"

// Many2One convert ResourceCalendarLeaves to *Many2One.
func (rcl *ResourceCalendarLeaves) Many2One() *Many2One {
	return NewMany2One(rcl.Id.Get(), "")
}

// CreateResourceCalendarLeaves creates a new resource.calendar.leaves model and returns its id.
func (c *Client) CreateResourceCalendarLeaves(rcl *ResourceCalendarLeaves) (int64, error) {
	return c.Create(ResourceCalendarLeavesModel, rcl)
}

// UpdateResourceCalendarLeaves updates an existing resource.calendar.leaves record.
func (c *Client) UpdateResourceCalendarLeaves(rcl *ResourceCalendarLeaves) error {
	return c.UpdateResourceCalendarLeavess([]int64{rcl.Id.Get()}, rcl)
}

// UpdateResourceCalendarLeavess updates existing resource.calendar.leaves records.
// All records (represented by ids) will be updated by rcl values.
func (c *Client) UpdateResourceCalendarLeavess(ids []int64, rcl *ResourceCalendarLeaves) error {
	return c.Update(ResourceCalendarLeavesModel, ids, rcl)
}

// DeleteResourceCalendarLeaves deletes an existing resource.calendar.leaves record.
func (c *Client) DeleteResourceCalendarLeaves(id int64) error {
	return c.DeleteResourceCalendarLeavess([]int64{id})
}

// DeleteResourceCalendarLeavess deletes existing resource.calendar.leaves records.
func (c *Client) DeleteResourceCalendarLeavess(ids []int64) error {
	return c.Delete(ResourceCalendarLeavesModel, ids)
}

// GetResourceCalendarLeaves gets resource.calendar.leaves existing record.
func (c *Client) GetResourceCalendarLeaves(id int64) (*ResourceCalendarLeaves, error) {
	rcls, err := c.GetResourceCalendarLeavess([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcls != nil && len(*rcls) > 0 {
		return &((*rcls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of resource.calendar.leaves not found", id)
}

// GetResourceCalendarLeavess gets resource.calendar.leaves existing records.
func (c *Client) GetResourceCalendarLeavess(ids []int64) (*ResourceCalendarLeavess, error) {
	rcls := &ResourceCalendarLeavess{}
	if err := c.Read(ResourceCalendarLeavesModel, ids, nil, rcls); err != nil {
		return nil, err
	}
	return rcls, nil
}

// FindResourceCalendarLeaves finds resource.calendar.leaves record by querying it with criteria.
func (c *Client) FindResourceCalendarLeaves(criteria *Criteria) (*ResourceCalendarLeaves, error) {
	rcls := &ResourceCalendarLeavess{}
	if err := c.SearchRead(ResourceCalendarLeavesModel, criteria, NewOptions().Limit(1), rcls); err != nil {
		return nil, err
	}
	if rcls != nil && len(*rcls) > 0 {
		return &((*rcls)[0]), nil
	}
	return nil, fmt.Errorf("resource.calendar.leaves was not found")
}

// FindResourceCalendarLeavess finds resource.calendar.leaves records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarLeavess(criteria *Criteria, options *Options) (*ResourceCalendarLeavess, error) {
	rcls := &ResourceCalendarLeavess{}
	if err := c.SearchRead(ResourceCalendarLeavesModel, criteria, options, rcls); err != nil {
		return nil, err
	}
	return rcls, nil
}

// FindResourceCalendarLeavesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarLeavesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResourceCalendarLeavesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResourceCalendarLeavesId finds record id by querying it with criteria.
func (c *Client) FindResourceCalendarLeavesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceCalendarLeavesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("resource.calendar.leaves was not found")
}
