package odoo

// CalendarEventType represents calendar.event.type model.
type CalendarEventType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateCalendarEventTypes([]*CalendarEventType{cet})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEventTypes creates a new calendar.event.type model and returns its id.
func (c *Client) CreateCalendarEventTypes(cets []*CalendarEventType) ([]int64, error) {
	var vv []interface{}
	for _, v := range cets {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventTypeModel, vv, nil)
}

// UpdateCalendarEventType updates an existing calendar.event.type record.
func (c *Client) UpdateCalendarEventType(cet *CalendarEventType) error {
	return c.UpdateCalendarEventTypes([]int64{cet.Id.Get()}, cet)
}

// UpdateCalendarEventTypes updates existing calendar.event.type records.
// All records (represented by ids) will be updated by cet values.
func (c *Client) UpdateCalendarEventTypes(ids []int64, cet *CalendarEventType) error {
	return c.Update(CalendarEventTypeModel, ids, cet, nil)
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
	return &((*cets)[0]), nil
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
	return &((*cets)[0]), nil
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
	return c.Search(CalendarEventTypeModel, criteria, options)
}

// FindCalendarEventTypeId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
