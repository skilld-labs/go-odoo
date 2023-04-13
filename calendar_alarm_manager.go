package odoo

import (
	"fmt"
)

// CalendarAlarmManager represents calendar.alarm_manager model.
type CalendarAlarmManager struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// CalendarAlarmManagers represents array of calendar.alarm_manager model.
type CalendarAlarmManagers []CalendarAlarmManager

// CalendarAlarmManagerModel is the odoo model name.
const CalendarAlarmManagerModel = "calendar.alarm_manager"

// Many2One convert CalendarAlarmManager to *Many2One.
func (ca *CalendarAlarmManager) Many2One() *Many2One {
	return NewMany2One(ca.Id.Get(), "")
}

// CreateCalendarAlarmManager creates a new calendar.alarm_manager model and returns its id.
func (c *Client) CreateCalendarAlarmManager(ca *CalendarAlarmManager) (int64, error) {
	ids, err := c.Create(CalendarAlarmManagerModel, []interface{}{ca})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarAlarmManager creates a new calendar.alarm_manager model and returns its id.
func (c *Client) CreateCalendarAlarmManagers(cas []*CalendarAlarmManager) ([]int64, error) {
	var vv []interface{}
	for _, v := range cas {
		vv = append(vv, v)
	}
	return c.Create(CalendarAlarmManagerModel, vv)
}

// UpdateCalendarAlarmManager updates an existing calendar.alarm_manager record.
func (c *Client) UpdateCalendarAlarmManager(ca *CalendarAlarmManager) error {
	return c.UpdateCalendarAlarmManagers([]int64{ca.Id.Get()}, ca)
}

// UpdateCalendarAlarmManagers updates existing calendar.alarm_manager records.
// All records (represented by ids) will be updated by ca values.
func (c *Client) UpdateCalendarAlarmManagers(ids []int64, ca *CalendarAlarmManager) error {
	return c.Update(CalendarAlarmManagerModel, ids, ca)
}

// DeleteCalendarAlarmManager deletes an existing calendar.alarm_manager record.
func (c *Client) DeleteCalendarAlarmManager(id int64) error {
	return c.DeleteCalendarAlarmManagers([]int64{id})
}

// DeleteCalendarAlarmManagers deletes existing calendar.alarm_manager records.
func (c *Client) DeleteCalendarAlarmManagers(ids []int64) error {
	return c.Delete(CalendarAlarmManagerModel, ids)
}

// GetCalendarAlarmManager gets calendar.alarm_manager existing record.
func (c *Client) GetCalendarAlarmManager(id int64) (*CalendarAlarmManager, error) {
	cas, err := c.GetCalendarAlarmManagers([]int64{id})
	if err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.alarm_manager not found", id)
}

// GetCalendarAlarmManagers gets calendar.alarm_manager existing records.
func (c *Client) GetCalendarAlarmManagers(ids []int64) (*CalendarAlarmManagers, error) {
	cas := &CalendarAlarmManagers{}
	if err := c.Read(CalendarAlarmManagerModel, ids, nil, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAlarmManager finds calendar.alarm_manager record by querying it with criteria.
func (c *Client) FindCalendarAlarmManager(criteria *Criteria) (*CalendarAlarmManager, error) {
	cas := &CalendarAlarmManagers{}
	if err := c.SearchRead(CalendarAlarmManagerModel, criteria, NewOptions().Limit(1), cas); err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("calendar.alarm_manager was not found with criteria %v", criteria)
}

// FindCalendarAlarmManagers finds calendar.alarm_manager records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAlarmManagers(criteria *Criteria, options *Options) (*CalendarAlarmManagers, error) {
	cas := &CalendarAlarmManagers{}
	if err := c.SearchRead(CalendarAlarmManagerModel, criteria, options, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCalendarAlarmManagerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarAlarmManagerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarAlarmManagerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarAlarmManagerId finds record id by querying it with criteria.
func (c *Client) FindCalendarAlarmManagerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarAlarmManagerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.alarm_manager was not found with criteria %v and options %v", criteria, options)
}
