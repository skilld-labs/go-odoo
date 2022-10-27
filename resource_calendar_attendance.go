package odoo

import (
	"fmt"
)

// ResourceCalendarAttendance represents resource.calendar.attendance model.
type ResourceCalendarAttendance struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	CalendarId       *Many2One  `xmlrpc:"calendar_id,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom         *Time      `xmlrpc:"date_from,omitempty"`
	DateTo           *Time      `xmlrpc:"date_to,omitempty"`
	DayPeriod        *Selection `xmlrpc:"day_period,omitempty"`
	Dayofweek        *Selection `xmlrpc:"dayofweek,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	DisplayType      *Selection `xmlrpc:"display_type,omitempty"`
	HourFrom         *Float     `xmlrpc:"hour_from,omitempty"`
	HourTo           *Float     `xmlrpc:"hour_to,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	ResourceId       *Many2One  `xmlrpc:"resource_id,omitempty"`
	Sequence         *Int       `xmlrpc:"sequence,omitempty"`
	TwoWeeksCalendar *Bool      `xmlrpc:"two_weeks_calendar,omitempty"`
	WeekType         *Selection `xmlrpc:"week_type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResourceCalendarAttendances represents array of resource.calendar.attendance model.
type ResourceCalendarAttendances []ResourceCalendarAttendance

// ResourceCalendarAttendanceModel is the odoo model name.
const ResourceCalendarAttendanceModel = "resource.calendar.attendance"

// Many2One convert ResourceCalendarAttendance to *Many2One.
func (rca *ResourceCalendarAttendance) Many2One() *Many2One {
	return NewMany2One(rca.Id.Get(), "")
}

// CreateResourceCalendarAttendance creates a new resource.calendar.attendance model and returns its id.
func (c *Client) CreateResourceCalendarAttendance(rca *ResourceCalendarAttendance) (int64, error) {
	return c.Create(ResourceCalendarAttendanceModel, rca)
}

// UpdateResourceCalendarAttendance updates an existing resource.calendar.attendance record.
func (c *Client) UpdateResourceCalendarAttendance(rca *ResourceCalendarAttendance) error {
	return c.UpdateResourceCalendarAttendances([]int64{rca.Id.Get()}, rca)
}

// UpdateResourceCalendarAttendances updates existing resource.calendar.attendance records.
// All records (represented by ids) will be updated by rca values.
func (c *Client) UpdateResourceCalendarAttendances(ids []int64, rca *ResourceCalendarAttendance) error {
	return c.Update(ResourceCalendarAttendanceModel, ids, rca)
}

// DeleteResourceCalendarAttendance deletes an existing resource.calendar.attendance record.
func (c *Client) DeleteResourceCalendarAttendance(id int64) error {
	return c.DeleteResourceCalendarAttendances([]int64{id})
}

// DeleteResourceCalendarAttendances deletes existing resource.calendar.attendance records.
func (c *Client) DeleteResourceCalendarAttendances(ids []int64) error {
	return c.Delete(ResourceCalendarAttendanceModel, ids)
}

// GetResourceCalendarAttendance gets resource.calendar.attendance existing record.
func (c *Client) GetResourceCalendarAttendance(id int64) (*ResourceCalendarAttendance, error) {
	rcas, err := c.GetResourceCalendarAttendances([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcas != nil && len(*rcas) > 0 {
		return &((*rcas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of resource.calendar.attendance not found", id)
}

// GetResourceCalendarAttendances gets resource.calendar.attendance existing records.
func (c *Client) GetResourceCalendarAttendances(ids []int64) (*ResourceCalendarAttendances, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.Read(ResourceCalendarAttendanceModel, ids, nil, rcas); err != nil {
		return nil, err
	}
	return rcas, nil
}

// FindResourceCalendarAttendance finds resource.calendar.attendance record by querying it with criteria.
func (c *Client) FindResourceCalendarAttendance(criteria *Criteria) (*ResourceCalendarAttendance, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.SearchRead(ResourceCalendarAttendanceModel, criteria, NewOptions().Limit(1), rcas); err != nil {
		return nil, err
	}
	if rcas != nil && len(*rcas) > 0 {
		return &((*rcas)[0]), nil
	}
	return nil, fmt.Errorf("resource.calendar.attendance was not found")
}

// FindResourceCalendarAttendances finds resource.calendar.attendance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarAttendances(criteria *Criteria, options *Options) (*ResourceCalendarAttendances, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.SearchRead(ResourceCalendarAttendanceModel, criteria, options, rcas); err != nil {
		return nil, err
	}
	return rcas, nil
}

// FindResourceCalendarAttendanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarAttendanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResourceCalendarAttendanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResourceCalendarAttendanceId finds record id by querying it with criteria.
func (c *Client) FindResourceCalendarAttendanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceCalendarAttendanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("resource.calendar.attendance was not found")
}
