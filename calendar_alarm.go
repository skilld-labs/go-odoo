package odoo

// CalendarAlarm represents calendar.alarm model.
type CalendarAlarm struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Duration        *Int       `xmlrpc:"duration,omitempty"`
	DurationMinutes *Int       `xmlrpc:"duration_minutes,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Interval        *Selection `xmlrpc:"interval,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	Type            *Selection `xmlrpc:"type,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateCalendarAlarms([]*CalendarAlarm{ca})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarAlarms creates a new calendar.alarm model and returns its id.
func (c *Client) CreateCalendarAlarms(cas []*CalendarAlarm) ([]int64, error) {
	var vv []interface{}
	for _, v := range cas {
		vv = append(vv, v)
	}
	return c.Create(CalendarAlarmModel, vv, nil)
}

// UpdateCalendarAlarm updates an existing calendar.alarm record.
func (c *Client) UpdateCalendarAlarm(ca *CalendarAlarm) error {
	return c.UpdateCalendarAlarms([]int64{ca.Id.Get()}, ca)
}

// UpdateCalendarAlarms updates existing calendar.alarm records.
// All records (represented by ids) will be updated by ca values.
func (c *Client) UpdateCalendarAlarms(ids []int64, ca *CalendarAlarm) error {
	return c.Update(CalendarAlarmModel, ids, ca, nil)
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
	return &((*cas)[0]), nil
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
	return &((*cas)[0]), nil
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
	return c.Search(CalendarAlarmModel, criteria, options)
}

// FindCalendarAlarmId finds record id by querying it with criteria.
func (c *Client) FindCalendarAlarmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarAlarmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
