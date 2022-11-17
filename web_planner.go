package odoo

import (
	"fmt"
)

// WebPlanner represents web.planner model.
type WebPlanner struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	Active             *Bool      `xmlrpc:"active,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	Data               *String    `xmlrpc:"data,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	MenuId             *Many2One  `xmlrpc:"menu_id,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	PlannerApplication *Selection `xmlrpc:"planner_application,omitempty"`
	Progress           *Int       `xmlrpc:"progress,omitempty"`
	TooltipPlanner     *String    `xmlrpc:"tooltip_planner,omitempty"`
	ViewId             *Many2One  `xmlrpc:"view_id,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// WebPlanners represents array of web.planner model.
type WebPlanners []WebPlanner

// WebPlannerModel is the odoo model name.
const WebPlannerModel = "web.planner"

// Many2One convert WebPlanner to *Many2One.
func (wp *WebPlanner) Many2One() *Many2One {
	return NewMany2One(wp.Id.Get(), "")
}

// CreateWebPlanner creates a new web.planner model and returns its id.
func (c *Client) CreateWebPlanner(wp *WebPlanner) (int64, error) {
	return c.Create(WebPlannerModel, wp)
}

// UpdateWebPlanner updates an existing web.planner record.
func (c *Client) UpdateWebPlanner(wp *WebPlanner) error {
	return c.UpdateWebPlanners([]int64{wp.Id.Get()}, wp)
}

// UpdateWebPlanners updates existing web.planner records.
// All records (represented by ids) will be updated by wp values.
func (c *Client) UpdateWebPlanners(ids []int64, wp *WebPlanner) error {
	return c.Update(WebPlannerModel, ids, wp)
}

// DeleteWebPlanner deletes an existing web.planner record.
func (c *Client) DeleteWebPlanner(id int64) error {
	return c.DeleteWebPlanners([]int64{id})
}

// DeleteWebPlanners deletes existing web.planner records.
func (c *Client) DeleteWebPlanners(ids []int64) error {
	return c.Delete(WebPlannerModel, ids)
}

// GetWebPlanner gets web.planner existing record.
func (c *Client) GetWebPlanner(id int64) (*WebPlanner, error) {
	wps, err := c.GetWebPlanners([]int64{id})
	if err != nil {
		return nil, err
	}
	if wps != nil && len(*wps) > 0 {
		return &((*wps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of web.planner not found", id)
}

// GetWebPlanners gets web.planner existing records.
func (c *Client) GetWebPlanners(ids []int64) (*WebPlanners, error) {
	wps := &WebPlanners{}
	if err := c.Read(WebPlannerModel, ids, nil, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebPlanner finds web.planner record by querying it with criteria.
func (c *Client) FindWebPlanner(criteria *Criteria) (*WebPlanner, error) {
	wps := &WebPlanners{}
	if err := c.SearchRead(WebPlannerModel, criteria, NewOptions().Limit(1), wps); err != nil {
		return nil, err
	}
	if wps != nil && len(*wps) > 0 {
		return &((*wps)[0]), nil
	}
	return nil, fmt.Errorf("no web.planner was found with criteria %v", criteria)
}

// FindWebPlanners finds web.planner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebPlanners(criteria *Criteria, options *Options) (*WebPlanners, error) {
	wps := &WebPlanners{}
	if err := c.SearchRead(WebPlannerModel, criteria, options, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebPlannerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebPlannerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebPlannerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebPlannerId finds record id by querying it with criteria.
func (c *Client) FindWebPlannerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebPlannerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no web.planner was found with criteria %v and options %v", criteria, options)
}
