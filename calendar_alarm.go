package odoo

import (
	"fmt"
)

// CalendarAlarm represents calendar.alarm model.
type CalendarAlarm struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Duration        *Int       `xmlrpc:"duration,omptempty"`
	DurationMinutes *Int       `xmlrpc:"duration_minutes,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Interval        *Selection `xmlrpc:"interval,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	Type            *Selection `xmlrpc:"type,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CalendarAlarms represents array of calendar.alarm model.
type CalendarAlarms []CalendarAlarm

// CalendarAlarmModel is the odoo model name.
const CalendarAlarmModel = "calendar.alarm"

// Many2One convert CalendarAlarm to *Many2One.
func (ca *CalendarAlarm) Many2One() *Many2One {
	return NewMany2One(ca.Id.Get(), "")
}

// CreateCalendarAlarm creates a new calendar.alarm model and returns its id.
func (c *Client) CreateCalendarAlarm(ca *CalendarAlarm) (int64, error) {
	return c.Create(CalendarAlarmModel, ca)
}

// UpdateCalendarAlarm updates an existing calendar.alarm record.
func (c *Client) UpdateCalendarAlarm(ca *CalendarAlarm) error {
	return c.UpdateCalendarAlarms([]int64{ca.Id.Get()}, ca)
}

// UpdateCalendarAlarms updates existing calendar.alarm records.
// All records (represented by ids) will be updated by ca values.
func (c *Client) UpdateCalendarAlarms(ids []int64, ca *CalendarAlarm) error {
	return c.Update(CalendarAlarmModel, ids, ca)
}

// DeleteCalendarAlarm deletes an existing calendar.alarm record.
func (c *Client) DeleteCalendarAlarm(id int64) error {
	return c.DeleteCalendarAlarms([]int64{id})
}

// DeleteCalendarAlarms deletes existing calendar.alarm records.
func (c *Client) DeleteCalendarAlarms(ids []int64) error {
	return c.Delete(CalendarAlarmModel, ids)
}

// GetCalendarAlarm gets calendar.alarm existing record.
func (c *Client) GetCalendarAlarm(id int64) (*CalendarAlarm, error) {
	cas, err := c.GetCalendarAlarms([]int64{id})
	if err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.alarm not found", id)
}

// GetCalendarAlarms gets calendar.alarm existing records.
func (c *Client) GetCalendarAlarms(ids []int64) (*CalendarAlarms, error) {
	cas := &CalendarAlarms{}
	if err := c.Read(CalendarAlarmModel, ids, nil, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAlarm finds calendar.alarm record by querying it with criteria.
func (c *Client) FindCalendarAlarm(criteria *Criteria) (*CalendarAlarm, error) {
	cas := &CalendarAlarms{}
	if err := c.SearchRead(CalendarAlarmModel, criteria, NewOptions().Limit(1), cas); err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("no calendar.alarm was found with criteria %v", criteria)
}

// FindCalendarAlarms finds calendar.alarm records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAlarms(criteria *Criteria, options *Options) (*CalendarAlarms, error) {
	cas := &CalendarAlarms{}
	if err := c.SearchRead(CalendarAlarmModel, criteria, options, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAlarmIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAlarmIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarAlarmModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarAlarmId finds record id by querying it with criteria.
func (c *Client) FindCalendarAlarmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarAlarmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no calendar.alarm was found with criteria %v and options %v", criteria, options)
}
