package odoo

// CalendarContacts represents calendar.contacts model.
type CalendarContacts struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CalendarContactss represents array of calendar.contacts model.
type CalendarContactss []CalendarContacts

// CalendarContactsModel is the odoo model name.
const CalendarContactsModel = "calendar.contacts"

// Many2One convert CalendarContacts to *Many2One.
func (cc *CalendarContacts) Many2One() *Many2One {
	return NewMany2One(cc.Id.Get(), "")
}

// CreateCalendarContacts creates a new calendar.contacts model and returns its id.
func (c *Client) CreateCalendarContacts(cc *CalendarContacts) (int64, error) {
	ids, err := c.CreateCalendarContactss([]*CalendarContacts{cc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarContacts creates a new calendar.contacts model and returns its id.
func (c *Client) CreateCalendarContactss(ccs []*CalendarContacts) ([]int64, error) {
	var vv []interface{}
	for _, v := range ccs {
		vv = append(vv, v)
	}
	return c.Create(CalendarContactsModel, vv, nil)
}

// UpdateCalendarContacts updates an existing calendar.contacts record.
func (c *Client) UpdateCalendarContacts(cc *CalendarContacts) error {
	return c.UpdateCalendarContactss([]int64{cc.Id.Get()}, cc)
}

// UpdateCalendarContactss updates existing calendar.contacts records.
// All records (represented by ids) will be updated by cc values.
func (c *Client) UpdateCalendarContactss(ids []int64, cc *CalendarContacts) error {
	return c.Update(CalendarContactsModel, ids, cc, nil)
}

// DeleteCalendarContacts deletes an existing calendar.contacts record.
func (c *Client) DeleteCalendarContacts(id int64) error {
	return c.DeleteCalendarContactss([]int64{id})
}

// DeleteCalendarContactss deletes existing calendar.contacts records.
func (c *Client) DeleteCalendarContactss(ids []int64) error {
	return c.Delete(CalendarContactsModel, ids)
}

// GetCalendarContacts gets calendar.contacts existing record.
func (c *Client) GetCalendarContacts(id int64) (*CalendarContacts, error) {
	ccs, err := c.GetCalendarContactss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ccs)[0]), nil
}

// GetCalendarContactss gets calendar.contacts existing records.
func (c *Client) GetCalendarContactss(ids []int64) (*CalendarContactss, error) {
	ccs := &CalendarContactss{}
	if err := c.Read(CalendarContactsModel, ids, nil, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCalendarContacts finds calendar.contacts record by querying it with criteria.
func (c *Client) FindCalendarContacts(criteria *Criteria) (*CalendarContacts, error) {
	ccs := &CalendarContactss{}
	if err := c.SearchRead(CalendarContactsModel, criteria, NewOptions().Limit(1), ccs); err != nil {
		return nil, err
	}
	return &((*ccs)[0]), nil
}

// FindCalendarContactss finds calendar.contacts records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarContactss(criteria *Criteria, options *Options) (*CalendarContactss, error) {
	ccs := &CalendarContactss{}
	if err := c.SearchRead(CalendarContactsModel, criteria, options, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCalendarContactsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarContactsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarContactsModel, criteria, options)
}

// FindCalendarContactsId finds record id by querying it with criteria.
func (c *Client) FindCalendarContactsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarContactsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
