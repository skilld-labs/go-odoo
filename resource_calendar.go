package odoo

// ResourceCalendar represents resource.calendar model.
type ResourceCalendar struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	AttendanceIds  *Relation `xmlrpc:"attendance_ids,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	GlobalLeaveIds *Relation `xmlrpc:"global_leave_ids,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	LeaveIds       *Relation `xmlrpc:"leave_ids,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResourceCalendars represents array of resource.calendar model.
type ResourceCalendars []ResourceCalendar

// ResourceCalendarModel is the odoo model name.
const ResourceCalendarModel = "resource.calendar"

// Many2One convert ResourceCalendar to *Many2One.
func (rc *ResourceCalendar) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResourceCalendar creates a new resource.calendar model and returns its id.
func (c *Client) CreateResourceCalendar(rc *ResourceCalendar) (int64, error) {
	ids, err := c.CreateResourceCalendars([]*ResourceCalendar{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceCalendars creates a new resource.calendar model and returns its id.
func (c *Client) CreateResourceCalendars(rcs []*ResourceCalendar) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResourceCalendarModel, vv, nil)
}

// UpdateResourceCalendar updates an existing resource.calendar record.
func (c *Client) UpdateResourceCalendar(rc *ResourceCalendar) error {
	return c.UpdateResourceCalendars([]int64{rc.Id.Get()}, rc)
}

// UpdateResourceCalendars updates existing resource.calendar records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResourceCalendars(ids []int64, rc *ResourceCalendar) error {
	return c.Update(ResourceCalendarModel, ids, rc, nil)
}

// DeleteResourceCalendar deletes an existing resource.calendar record.
func (c *Client) DeleteResourceCalendar(id int64) error {
	return c.DeleteResourceCalendars([]int64{id})
}

// DeleteResourceCalendars deletes existing resource.calendar records.
func (c *Client) DeleteResourceCalendars(ids []int64) error {
	return c.Delete(ResourceCalendarModel, ids)
}

// GetResourceCalendar gets resource.calendar existing record.
func (c *Client) GetResourceCalendar(id int64) (*ResourceCalendar, error) {
	rcs, err := c.GetResourceCalendars([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResourceCalendars gets resource.calendar existing records.
func (c *Client) GetResourceCalendars(ids []int64) (*ResourceCalendars, error) {
	rcs := &ResourceCalendars{}
	if err := c.Read(ResourceCalendarModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResourceCalendar finds resource.calendar record by querying it with criteria.
func (c *Client) FindResourceCalendar(criteria *Criteria) (*ResourceCalendar, error) {
	rcs := &ResourceCalendars{}
	if err := c.SearchRead(ResourceCalendarModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResourceCalendars finds resource.calendar records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendars(criteria *Criteria, options *Options) (*ResourceCalendars, error) {
	rcs := &ResourceCalendars{}
	if err := c.SearchRead(ResourceCalendarModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResourceCalendarIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResourceCalendarModel, criteria, options)
}

// FindResourceCalendarId finds record id by querying it with criteria.
func (c *Client) FindResourceCalendarId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceCalendarModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
