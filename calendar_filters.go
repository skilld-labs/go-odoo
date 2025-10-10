package odoo

// CalendarFilters represents calendar.filters model.
type CalendarFilters struct {
	Active         *Bool     `xmlrpc:"active,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	PartnerChecked *Bool     `xmlrpc:"partner_checked,omitempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CalendarFilterss represents array of calendar.filters model.
type CalendarFilterss []CalendarFilters

// CalendarFiltersModel is the odoo model name.
const CalendarFiltersModel = "calendar.filters"

// Many2One convert CalendarFilters to *Many2One.
func (cf *CalendarFilters) Many2One() *Many2One {
	return NewMany2One(cf.Id.Get(), "")
}

// CreateCalendarFilters creates a new calendar.filters model and returns its id.
func (c *Client) CreateCalendarFilters(cf *CalendarFilters) (int64, error) {
	ids, err := c.CreateCalendarFilterss([]*CalendarFilters{cf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarFilters creates a new calendar.filters model and returns its id.
func (c *Client) CreateCalendarFilterss(cfs []*CalendarFilters) ([]int64, error) {
	var vv []interface{}
	for _, v := range cfs {
		vv = append(vv, v)
	}
	return c.Create(CalendarFiltersModel, vv, nil)
}

// UpdateCalendarFilters updates an existing calendar.filters record.
func (c *Client) UpdateCalendarFilters(cf *CalendarFilters) error {
	return c.UpdateCalendarFilterss([]int64{cf.Id.Get()}, cf)
}

// UpdateCalendarFilterss updates existing calendar.filters records.
// All records (represented by ids) will be updated by cf values.
func (c *Client) UpdateCalendarFilterss(ids []int64, cf *CalendarFilters) error {
	return c.Update(CalendarFiltersModel, ids, cf, nil)
}

// DeleteCalendarFilters deletes an existing calendar.filters record.
func (c *Client) DeleteCalendarFilters(id int64) error {
	return c.DeleteCalendarFilterss([]int64{id})
}

// DeleteCalendarFilterss deletes existing calendar.filters records.
func (c *Client) DeleteCalendarFilterss(ids []int64) error {
	return c.Delete(CalendarFiltersModel, ids)
}

// GetCalendarFilters gets calendar.filters existing record.
func (c *Client) GetCalendarFilters(id int64) (*CalendarFilters, error) {
	cfs, err := c.GetCalendarFilterss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cfs)[0]), nil
}

// GetCalendarFilterss gets calendar.filters existing records.
func (c *Client) GetCalendarFilterss(ids []int64) (*CalendarFilterss, error) {
	cfs := &CalendarFilterss{}
	if err := c.Read(CalendarFiltersModel, ids, nil, cfs); err != nil {
		return nil, err
	}
	return cfs, nil
}

// FindCalendarFilters finds calendar.filters record by querying it with criteria.
func (c *Client) FindCalendarFilters(criteria *Criteria) (*CalendarFilters, error) {
	cfs := &CalendarFilterss{}
	if err := c.SearchRead(CalendarFiltersModel, criteria, NewOptions().Limit(1), cfs); err != nil {
		return nil, err
	}
	return &((*cfs)[0]), nil
}

// FindCalendarFilterss finds calendar.filters records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarFilterss(criteria *Criteria, options *Options) (*CalendarFilterss, error) {
	cfs := &CalendarFilterss{}
	if err := c.SearchRead(CalendarFiltersModel, criteria, options, cfs); err != nil {
		return nil, err
	}
	return cfs, nil
}

// FindCalendarFiltersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarFiltersIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarFiltersModel, criteria, options)
}

// FindCalendarFiltersId finds record id by querying it with criteria.
func (c *Client) FindCalendarFiltersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarFiltersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
