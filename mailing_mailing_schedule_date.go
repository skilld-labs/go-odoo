package odoo

// MailingMailingScheduleDate represents mailing.mailing.schedule.date model.
type MailingMailingScheduleDate struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	MassMailingId *Many2One `xmlrpc:"mass_mailing_id,omitempty"`
	ScheduleDate  *Time     `xmlrpc:"schedule_date,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingMailingScheduleDates represents array of mailing.mailing.schedule.date model.
type MailingMailingScheduleDates []MailingMailingScheduleDate

// MailingMailingScheduleDateModel is the odoo model name.
const MailingMailingScheduleDateModel = "mailing.mailing.schedule.date"

// Many2One convert MailingMailingScheduleDate to *Many2One.
func (mmsd *MailingMailingScheduleDate) Many2One() *Many2One {
	return NewMany2One(mmsd.Id.Get(), "")
}

// CreateMailingMailingScheduleDate creates a new mailing.mailing.schedule.date model and returns its id.
func (c *Client) CreateMailingMailingScheduleDate(mmsd *MailingMailingScheduleDate) (int64, error) {
	ids, err := c.CreateMailingMailingScheduleDates([]*MailingMailingScheduleDate{mmsd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingMailingScheduleDate creates a new mailing.mailing.schedule.date model and returns its id.
func (c *Client) CreateMailingMailingScheduleDates(mmsds []*MailingMailingScheduleDate) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmsds {
		vv = append(vv, v)
	}
	return c.Create(MailingMailingScheduleDateModel, vv, nil)
}

// UpdateMailingMailingScheduleDate updates an existing mailing.mailing.schedule.date record.
func (c *Client) UpdateMailingMailingScheduleDate(mmsd *MailingMailingScheduleDate) error {
	return c.UpdateMailingMailingScheduleDates([]int64{mmsd.Id.Get()}, mmsd)
}

// UpdateMailingMailingScheduleDates updates existing mailing.mailing.schedule.date records.
// All records (represented by ids) will be updated by mmsd values.
func (c *Client) UpdateMailingMailingScheduleDates(ids []int64, mmsd *MailingMailingScheduleDate) error {
	return c.Update(MailingMailingScheduleDateModel, ids, mmsd, nil)
}

// DeleteMailingMailingScheduleDate deletes an existing mailing.mailing.schedule.date record.
func (c *Client) DeleteMailingMailingScheduleDate(id int64) error {
	return c.DeleteMailingMailingScheduleDates([]int64{id})
}

// DeleteMailingMailingScheduleDates deletes existing mailing.mailing.schedule.date records.
func (c *Client) DeleteMailingMailingScheduleDates(ids []int64) error {
	return c.Delete(MailingMailingScheduleDateModel, ids)
}

// GetMailingMailingScheduleDate gets mailing.mailing.schedule.date existing record.
func (c *Client) GetMailingMailingScheduleDate(id int64) (*MailingMailingScheduleDate, error) {
	mmsds, err := c.GetMailingMailingScheduleDates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmsds)[0]), nil
}

// GetMailingMailingScheduleDates gets mailing.mailing.schedule.date existing records.
func (c *Client) GetMailingMailingScheduleDates(ids []int64) (*MailingMailingScheduleDates, error) {
	mmsds := &MailingMailingScheduleDates{}
	if err := c.Read(MailingMailingScheduleDateModel, ids, nil, mmsds); err != nil {
		return nil, err
	}
	return mmsds, nil
}

// FindMailingMailingScheduleDate finds mailing.mailing.schedule.date record by querying it with criteria.
func (c *Client) FindMailingMailingScheduleDate(criteria *Criteria) (*MailingMailingScheduleDate, error) {
	mmsds := &MailingMailingScheduleDates{}
	if err := c.SearchRead(MailingMailingScheduleDateModel, criteria, NewOptions().Limit(1), mmsds); err != nil {
		return nil, err
	}
	return &((*mmsds)[0]), nil
}

// FindMailingMailingScheduleDates finds mailing.mailing.schedule.date records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingScheduleDates(criteria *Criteria, options *Options) (*MailingMailingScheduleDates, error) {
	mmsds := &MailingMailingScheduleDates{}
	if err := c.SearchRead(MailingMailingScheduleDateModel, criteria, options, mmsds); err != nil {
		return nil, err
	}
	return mmsds, nil
}

// FindMailingMailingScheduleDateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingScheduleDateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingMailingScheduleDateModel, criteria, options)
}

// FindMailingMailingScheduleDateId finds record id by querying it with criteria.
func (c *Client) FindMailingMailingScheduleDateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingMailingScheduleDateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
