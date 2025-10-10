package odoo

// CalendarRecurrence represents calendar.recurrence model.
type CalendarRecurrence struct {
	BaseEventId      *Many2One  `xmlrpc:"base_event_id,omitempty"`
	Byday            *Selection `xmlrpc:"byday,omitempty"`
	CalendarEventIds *Relation  `xmlrpc:"calendar_event_ids,omitempty"`
	Count            *Int       `xmlrpc:"count,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Day              *Int       `xmlrpc:"day,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Dtstart          *Time      `xmlrpc:"dtstart,omitempty"`
	EndType          *Selection `xmlrpc:"end_type,omitempty"`
	EventTz          *Selection `xmlrpc:"event_tz,omitempty"`
	Fri              *Bool      `xmlrpc:"fri,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Interval         *Int       `xmlrpc:"interval,omitempty"`
	Mon              *Bool      `xmlrpc:"mon,omitempty"`
	MonthBy          *Selection `xmlrpc:"month_by,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	Rrule            *String    `xmlrpc:"rrule,omitempty"`
	RruleType        *Selection `xmlrpc:"rrule_type,omitempty"`
	Sat              *Bool      `xmlrpc:"sat,omitempty"`
	Sun              *Bool      `xmlrpc:"sun,omitempty"`
	Thu              *Bool      `xmlrpc:"thu,omitempty"`
	TriggerId        *Many2One  `xmlrpc:"trigger_id,omitempty"`
	Tue              *Bool      `xmlrpc:"tue,omitempty"`
	Until            *Time      `xmlrpc:"until,omitempty"`
	Wed              *Bool      `xmlrpc:"wed,omitempty"`
	Weekday          *Selection `xmlrpc:"weekday,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CalendarRecurrences represents array of calendar.recurrence model.
type CalendarRecurrences []CalendarRecurrence

// CalendarRecurrenceModel is the odoo model name.
const CalendarRecurrenceModel = "calendar.recurrence"

// Many2One convert CalendarRecurrence to *Many2One.
func (cr *CalendarRecurrence) Many2One() *Many2One {
	return NewMany2One(cr.Id.Get(), "")
}

// CreateCalendarRecurrence creates a new calendar.recurrence model and returns its id.
func (c *Client) CreateCalendarRecurrence(cr *CalendarRecurrence) (int64, error) {
	ids, err := c.CreateCalendarRecurrences([]*CalendarRecurrence{cr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarRecurrence creates a new calendar.recurrence model and returns its id.
func (c *Client) CreateCalendarRecurrences(crs []*CalendarRecurrence) ([]int64, error) {
	var vv []interface{}
	for _, v := range crs {
		vv = append(vv, v)
	}
	return c.Create(CalendarRecurrenceModel, vv, nil)
}

// UpdateCalendarRecurrence updates an existing calendar.recurrence record.
func (c *Client) UpdateCalendarRecurrence(cr *CalendarRecurrence) error {
	return c.UpdateCalendarRecurrences([]int64{cr.Id.Get()}, cr)
}

// UpdateCalendarRecurrences updates existing calendar.recurrence records.
// All records (represented by ids) will be updated by cr values.
func (c *Client) UpdateCalendarRecurrences(ids []int64, cr *CalendarRecurrence) error {
	return c.Update(CalendarRecurrenceModel, ids, cr, nil)
}

// DeleteCalendarRecurrence deletes an existing calendar.recurrence record.
func (c *Client) DeleteCalendarRecurrence(id int64) error {
	return c.DeleteCalendarRecurrences([]int64{id})
}

// DeleteCalendarRecurrences deletes existing calendar.recurrence records.
func (c *Client) DeleteCalendarRecurrences(ids []int64) error {
	return c.Delete(CalendarRecurrenceModel, ids)
}

// GetCalendarRecurrence gets calendar.recurrence existing record.
func (c *Client) GetCalendarRecurrence(id int64) (*CalendarRecurrence, error) {
	crs, err := c.GetCalendarRecurrences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*crs)[0]), nil
}

// GetCalendarRecurrences gets calendar.recurrence existing records.
func (c *Client) GetCalendarRecurrences(ids []int64) (*CalendarRecurrences, error) {
	crs := &CalendarRecurrences{}
	if err := c.Read(CalendarRecurrenceModel, ids, nil, crs); err != nil {
		return nil, err
	}
	return crs, nil
}

// FindCalendarRecurrence finds calendar.recurrence record by querying it with criteria.
func (c *Client) FindCalendarRecurrence(criteria *Criteria) (*CalendarRecurrence, error) {
	crs := &CalendarRecurrences{}
	if err := c.SearchRead(CalendarRecurrenceModel, criteria, NewOptions().Limit(1), crs); err != nil {
		return nil, err
	}
	return &((*crs)[0]), nil
}

// FindCalendarRecurrences finds calendar.recurrence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarRecurrences(criteria *Criteria, options *Options) (*CalendarRecurrences, error) {
	crs := &CalendarRecurrences{}
	if err := c.SearchRead(CalendarRecurrenceModel, criteria, options, crs); err != nil {
		return nil, err
	}
	return crs, nil
}

// FindCalendarRecurrenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarRecurrenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarRecurrenceModel, criteria, options)
}

// FindCalendarRecurrenceId finds record id by querying it with criteria.
func (c *Client) FindCalendarRecurrenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarRecurrenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
