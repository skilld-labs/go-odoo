package odoo

// MailActivityScheduleLine represents mail.activity.schedule.line model.
type MailActivityScheduleLine struct {
	ActivityScheduleId *Many2One `xmlrpc:"activity_schedule_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	LineDateDeadline   *Time     `xmlrpc:"line_date_deadline,omitempty"`
	LineDescription    *String   `xmlrpc:"line_description,omitempty"`
	ResponsibleUserId  *Many2One `xmlrpc:"responsible_user_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailActivityScheduleLines represents array of mail.activity.schedule.line model.
type MailActivityScheduleLines []MailActivityScheduleLine

// MailActivityScheduleLineModel is the odoo model name.
const MailActivityScheduleLineModel = "mail.activity.schedule.line"

// Many2One convert MailActivityScheduleLine to *Many2One.
func (masl *MailActivityScheduleLine) Many2One() *Many2One {
	return NewMany2One(masl.Id.Get(), "")
}

// CreateMailActivityScheduleLine creates a new mail.activity.schedule.line model and returns its id.
func (c *Client) CreateMailActivityScheduleLine(masl *MailActivityScheduleLine) (int64, error) {
	ids, err := c.CreateMailActivityScheduleLines([]*MailActivityScheduleLine{masl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityScheduleLine creates a new mail.activity.schedule.line model and returns its id.
func (c *Client) CreateMailActivityScheduleLines(masls []*MailActivityScheduleLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range masls {
		vv = append(vv, v)
	}
	return c.Create(MailActivityScheduleLineModel, vv, nil)
}

// UpdateMailActivityScheduleLine updates an existing mail.activity.schedule.line record.
func (c *Client) UpdateMailActivityScheduleLine(masl *MailActivityScheduleLine) error {
	return c.UpdateMailActivityScheduleLines([]int64{masl.Id.Get()}, masl)
}

// UpdateMailActivityScheduleLines updates existing mail.activity.schedule.line records.
// All records (represented by ids) will be updated by masl values.
func (c *Client) UpdateMailActivityScheduleLines(ids []int64, masl *MailActivityScheduleLine) error {
	return c.Update(MailActivityScheduleLineModel, ids, masl, nil)
}

// DeleteMailActivityScheduleLine deletes an existing mail.activity.schedule.line record.
func (c *Client) DeleteMailActivityScheduleLine(id int64) error {
	return c.DeleteMailActivityScheduleLines([]int64{id})
}

// DeleteMailActivityScheduleLines deletes existing mail.activity.schedule.line records.
func (c *Client) DeleteMailActivityScheduleLines(ids []int64) error {
	return c.Delete(MailActivityScheduleLineModel, ids)
}

// GetMailActivityScheduleLine gets mail.activity.schedule.line existing record.
func (c *Client) GetMailActivityScheduleLine(id int64) (*MailActivityScheduleLine, error) {
	masls, err := c.GetMailActivityScheduleLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*masls)[0]), nil
}

// GetMailActivityScheduleLines gets mail.activity.schedule.line existing records.
func (c *Client) GetMailActivityScheduleLines(ids []int64) (*MailActivityScheduleLines, error) {
	masls := &MailActivityScheduleLines{}
	if err := c.Read(MailActivityScheduleLineModel, ids, nil, masls); err != nil {
		return nil, err
	}
	return masls, nil
}

// FindMailActivityScheduleLine finds mail.activity.schedule.line record by querying it with criteria.
func (c *Client) FindMailActivityScheduleLine(criteria *Criteria) (*MailActivityScheduleLine, error) {
	masls := &MailActivityScheduleLines{}
	if err := c.SearchRead(MailActivityScheduleLineModel, criteria, NewOptions().Limit(1), masls); err != nil {
		return nil, err
	}
	return &((*masls)[0]), nil
}

// FindMailActivityScheduleLines finds mail.activity.schedule.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityScheduleLines(criteria *Criteria, options *Options) (*MailActivityScheduleLines, error) {
	masls := &MailActivityScheduleLines{}
	if err := c.SearchRead(MailActivityScheduleLineModel, criteria, options, masls); err != nil {
		return nil, err
	}
	return masls, nil
}

// FindMailActivityScheduleLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityScheduleLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityScheduleLineModel, criteria, options)
}

// FindMailActivityScheduleLineId finds record id by querying it with criteria.
func (c *Client) FindMailActivityScheduleLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityScheduleLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
