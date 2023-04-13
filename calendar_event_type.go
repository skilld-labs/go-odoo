package odoo

import (
	"fmt"
)

// CalendarEventType represents calendar.event.type model.
type CalendarEventType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CalendarEventTypes represents array of calendar.event.type model.
type CalendarEventTypes []CalendarEventType

// CalendarEventTypeModel is the odoo model name.
const CalendarEventTypeModel = "calendar.event.type"

// Many2One convert CalendarEventType to *Many2One.
func (cet *CalendarEventType) Many2One() *Many2One {
	return NewMany2One(cet.Id.Get(), "")
}

// CreateCalendarEventType creates a new calendar.event.type model and returns its id.
func (c *Client) CreateCalendarEventType(cet *CalendarEventType) (int64, error) {
	ids, err := c.Create(CalendarEventTypeModel, []interface{}{cet})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEventType creates a new calendar.event.type model and returns its id.
func (c *Client) CreateCalendarEventTypes(cets []*CalendarEventType) ([]int64, error) {
	var vv []interface{}
	for _, v := range cets {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventTypeModel, vv)
}

// UpdateCalendarEventType updates an existing calendar.event.type record.
func (c *Client) UpdateCalendarEventType(cet *CalendarEventType) error {
	return c.UpdateCalendarEventTypes([]int64{cet.Id.Get()}, cet)
}

// UpdateCalendarEventTypes updates existing calendar.event.type records.
// All records (represented by ids) will be updated by cet values.
func (c *Client) UpdateCalendarEventTypes(ids []int64, cet *CalendarEventType) error {
	return c.Update(CalendarEventTypeModel, ids, cet)
}

// DeleteCalendarEventType deletes an existing calendar.event.type record.
func (c *Client) DeleteCalendarEventType(id int64) error {
	return c.DeleteCalendarEventTypes([]int64{id})
}

// DeleteCalendarEventTypes deletes existing calendar.event.type records.
func (c *Client) DeleteCalendarEventTypes(ids []int64) error {
	return c.Delete(CalendarEventTypeModel, ids)
}

// GetCalendarEventType gets calendar.event.type existing record.
func (c *Client) GetCalendarEventType(id int64) (*CalendarEventType, error) {
	cets, err := c.GetCalendarEventTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if cets != nil && len(*cets) > 0 {
		return &((*cets)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.event.type not found", id)
}

// GetCalendarEventTypes gets calendar.event.type existing records.
func (c *Client) GetCalendarEventTypes(ids []int64) (*CalendarEventTypes, error) {
	cets := &CalendarEventTypes{}
	if err := c.Read(CalendarEventTypeModel, ids, nil, cets); err != nil {
		return nil, err
	}
	return cets, nil
}

// FindCalendarEventType finds calendar.event.type record by querying it with criteria.
func (c *Client) FindCalendarEventType(criteria *Criteria) (*CalendarEventType, error) {
	cets := &CalendarEventTypes{}
	if err := c.SearchRead(CalendarEventTypeModel, criteria, NewOptions().Limit(1), cets); err != nil {
		return nil, err
	}
	if cets != nil && len(*cets) > 0 {
		return &((*cets)[0]), nil
	}
	return nil, fmt.Errorf("calendar.event.type was not found with criteria %v", criteria)
}

// FindCalendarEventTypes finds calendar.event.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventTypes(criteria *Criteria, options *Options) (*CalendarEventTypes, error) {
	cets := &CalendarEventTypes{}
	if err := c.SearchRead(CalendarEventTypeModel, criteria, options, cets); err != nil {
		return nil, err
	}
	return cets, nil
}

// FindCalendarEventTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarEventTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarEventTypeId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.event.type was not found with criteria %v and options %v", criteria, options)
}
