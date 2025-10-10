package odoo

// MailMessageSchedule represents mail.message.schedule model.
type MailMessageSchedule struct {
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	MailMessageId          *Many2One `xmlrpc:"mail_message_id,omitempty"`
	NotificationParameters *String   `xmlrpc:"notification_parameters,omitempty"`
	ScheduledDatetime      *Time     `xmlrpc:"scheduled_datetime,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMessageSchedules represents array of mail.message.schedule model.
type MailMessageSchedules []MailMessageSchedule

// MailMessageScheduleModel is the odoo model name.
const MailMessageScheduleModel = "mail.message.schedule"

// Many2One convert MailMessageSchedule to *Many2One.
func (mms *MailMessageSchedule) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMessageSchedule creates a new mail.message.schedule model and returns its id.
func (c *Client) CreateMailMessageSchedule(mms *MailMessageSchedule) (int64, error) {
	ids, err := c.CreateMailMessageSchedules([]*MailMessageSchedule{mms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessageSchedule creates a new mail.message.schedule model and returns its id.
func (c *Client) CreateMailMessageSchedules(mmss []*MailMessageSchedule) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmss {
		vv = append(vv, v)
	}
	return c.Create(MailMessageScheduleModel, vv, nil)
}

// UpdateMailMessageSchedule updates an existing mail.message.schedule record.
func (c *Client) UpdateMailMessageSchedule(mms *MailMessageSchedule) error {
	return c.UpdateMailMessageSchedules([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMessageSchedules updates existing mail.message.schedule records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMessageSchedules(ids []int64, mms *MailMessageSchedule) error {
	return c.Update(MailMessageScheduleModel, ids, mms, nil)
}

// DeleteMailMessageSchedule deletes an existing mail.message.schedule record.
func (c *Client) DeleteMailMessageSchedule(id int64) error {
	return c.DeleteMailMessageSchedules([]int64{id})
}

// DeleteMailMessageSchedules deletes existing mail.message.schedule records.
func (c *Client) DeleteMailMessageSchedules(ids []int64) error {
	return c.Delete(MailMessageScheduleModel, ids)
}

// GetMailMessageSchedule gets mail.message.schedule existing record.
func (c *Client) GetMailMessageSchedule(id int64) (*MailMessageSchedule, error) {
	mmss, err := c.GetMailMessageSchedules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmss)[0]), nil
}

// GetMailMessageSchedules gets mail.message.schedule existing records.
func (c *Client) GetMailMessageSchedules(ids []int64) (*MailMessageSchedules, error) {
	mmss := &MailMessageSchedules{}
	if err := c.Read(MailMessageScheduleModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSchedule finds mail.message.schedule record by querying it with criteria.
func (c *Client) FindMailMessageSchedule(criteria *Criteria) (*MailMessageSchedule, error) {
	mmss := &MailMessageSchedules{}
	if err := c.SearchRead(MailMessageScheduleModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	return &((*mmss)[0]), nil
}

// FindMailMessageSchedules finds mail.message.schedule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSchedules(criteria *Criteria, options *Options) (*MailMessageSchedules, error) {
	mmss := &MailMessageSchedules{}
	if err := c.SearchRead(MailMessageScheduleModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageScheduleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageScheduleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageScheduleModel, criteria, options)
}

// FindMailMessageScheduleId finds record id by querying it with criteria.
func (c *Client) FindMailMessageScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
