package odoo

// CalendarPopoverDeleteWizard represents calendar.popover.delete.wizard model.
type CalendarPopoverDeleteWizard struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delete      *Selection `xmlrpc:"delete,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Record      *Many2One  `xmlrpc:"record,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CalendarPopoverDeleteWizards represents array of calendar.popover.delete.wizard model.
type CalendarPopoverDeleteWizards []CalendarPopoverDeleteWizard

// CalendarPopoverDeleteWizardModel is the odoo model name.
const CalendarPopoverDeleteWizardModel = "calendar.popover.delete.wizard"

// Many2One convert CalendarPopoverDeleteWizard to *Many2One.
func (cpdw *CalendarPopoverDeleteWizard) Many2One() *Many2One {
	return NewMany2One(cpdw.Id.Get(), "")
}

// CreateCalendarPopoverDeleteWizard creates a new calendar.popover.delete.wizard model and returns its id.
func (c *Client) CreateCalendarPopoverDeleteWizard(cpdw *CalendarPopoverDeleteWizard) (int64, error) {
	ids, err := c.CreateCalendarPopoverDeleteWizards([]*CalendarPopoverDeleteWizard{cpdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarPopoverDeleteWizard creates a new calendar.popover.delete.wizard model and returns its id.
func (c *Client) CreateCalendarPopoverDeleteWizards(cpdws []*CalendarPopoverDeleteWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpdws {
		vv = append(vv, v)
	}
	return c.Create(CalendarPopoverDeleteWizardModel, vv, nil)
}

// UpdateCalendarPopoverDeleteWizard updates an existing calendar.popover.delete.wizard record.
func (c *Client) UpdateCalendarPopoverDeleteWizard(cpdw *CalendarPopoverDeleteWizard) error {
	return c.UpdateCalendarPopoverDeleteWizards([]int64{cpdw.Id.Get()}, cpdw)
}

// UpdateCalendarPopoverDeleteWizards updates existing calendar.popover.delete.wizard records.
// All records (represented by ids) will be updated by cpdw values.
func (c *Client) UpdateCalendarPopoverDeleteWizards(ids []int64, cpdw *CalendarPopoverDeleteWizard) error {
	return c.Update(CalendarPopoverDeleteWizardModel, ids, cpdw, nil)
}

// DeleteCalendarPopoverDeleteWizard deletes an existing calendar.popover.delete.wizard record.
func (c *Client) DeleteCalendarPopoverDeleteWizard(id int64) error {
	return c.DeleteCalendarPopoverDeleteWizards([]int64{id})
}

// DeleteCalendarPopoverDeleteWizards deletes existing calendar.popover.delete.wizard records.
func (c *Client) DeleteCalendarPopoverDeleteWizards(ids []int64) error {
	return c.Delete(CalendarPopoverDeleteWizardModel, ids)
}

// GetCalendarPopoverDeleteWizard gets calendar.popover.delete.wizard existing record.
func (c *Client) GetCalendarPopoverDeleteWizard(id int64) (*CalendarPopoverDeleteWizard, error) {
	cpdws, err := c.GetCalendarPopoverDeleteWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cpdws)[0]), nil
}

// GetCalendarPopoverDeleteWizards gets calendar.popover.delete.wizard existing records.
func (c *Client) GetCalendarPopoverDeleteWizards(ids []int64) (*CalendarPopoverDeleteWizards, error) {
	cpdws := &CalendarPopoverDeleteWizards{}
	if err := c.Read(CalendarPopoverDeleteWizardModel, ids, nil, cpdws); err != nil {
		return nil, err
	}
	return cpdws, nil
}

// FindCalendarPopoverDeleteWizard finds calendar.popover.delete.wizard record by querying it with criteria.
func (c *Client) FindCalendarPopoverDeleteWizard(criteria *Criteria) (*CalendarPopoverDeleteWizard, error) {
	cpdws := &CalendarPopoverDeleteWizards{}
	if err := c.SearchRead(CalendarPopoverDeleteWizardModel, criteria, NewOptions().Limit(1), cpdws); err != nil {
		return nil, err
	}
	return &((*cpdws)[0]), nil
}

// FindCalendarPopoverDeleteWizards finds calendar.popover.delete.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarPopoverDeleteWizards(criteria *Criteria, options *Options) (*CalendarPopoverDeleteWizards, error) {
	cpdws := &CalendarPopoverDeleteWizards{}
	if err := c.SearchRead(CalendarPopoverDeleteWizardModel, criteria, options, cpdws); err != nil {
		return nil, err
	}
	return cpdws, nil
}

// FindCalendarPopoverDeleteWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarPopoverDeleteWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarPopoverDeleteWizardModel, criteria, options)
}

// FindCalendarPopoverDeleteWizardId finds record id by querying it with criteria.
func (c *Client) FindCalendarPopoverDeleteWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarPopoverDeleteWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
