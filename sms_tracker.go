package odoo

// SmsTracker represents sms.tracker model.
type SmsTracker struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	MailNotificationId *Many2One `xmlrpc:"mail_notification_id,omitempty"`
	SmsUuid            *String   `xmlrpc:"sms_uuid,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsTrackers represents array of sms.tracker model.
type SmsTrackers []SmsTracker

// SmsTrackerModel is the odoo model name.
const SmsTrackerModel = "sms.tracker"

// Many2One convert SmsTracker to *Many2One.
func (st *SmsTracker) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSmsTracker creates a new sms.tracker model and returns its id.
func (c *Client) CreateSmsTracker(st *SmsTracker) (int64, error) {
	ids, err := c.CreateSmsTrackers([]*SmsTracker{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsTracker creates a new sms.tracker model and returns its id.
func (c *Client) CreateSmsTrackers(sts []*SmsTracker) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SmsTrackerModel, vv, nil)
}

// UpdateSmsTracker updates an existing sms.tracker record.
func (c *Client) UpdateSmsTracker(st *SmsTracker) error {
	return c.UpdateSmsTrackers([]int64{st.Id.Get()}, st)
}

// UpdateSmsTrackers updates existing sms.tracker records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSmsTrackers(ids []int64, st *SmsTracker) error {
	return c.Update(SmsTrackerModel, ids, st, nil)
}

// DeleteSmsTracker deletes an existing sms.tracker record.
func (c *Client) DeleteSmsTracker(id int64) error {
	return c.DeleteSmsTrackers([]int64{id})
}

// DeleteSmsTrackers deletes existing sms.tracker records.
func (c *Client) DeleteSmsTrackers(ids []int64) error {
	return c.Delete(SmsTrackerModel, ids)
}

// GetSmsTracker gets sms.tracker existing record.
func (c *Client) GetSmsTracker(id int64) (*SmsTracker, error) {
	sts, err := c.GetSmsTrackers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// GetSmsTrackers gets sms.tracker existing records.
func (c *Client) GetSmsTrackers(ids []int64) (*SmsTrackers, error) {
	sts := &SmsTrackers{}
	if err := c.Read(SmsTrackerModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTracker finds sms.tracker record by querying it with criteria.
func (c *Client) FindSmsTracker(criteria *Criteria) (*SmsTracker, error) {
	sts := &SmsTrackers{}
	if err := c.SearchRead(SmsTrackerModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// FindSmsTrackers finds sms.tracker records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTrackers(criteria *Criteria, options *Options) (*SmsTrackers, error) {
	sts := &SmsTrackers{}
	if err := c.SearchRead(SmsTrackerModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTrackerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTrackerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsTrackerModel, criteria, options)
}

// FindSmsTrackerId finds record id by querying it with criteria.
func (c *Client) FindSmsTrackerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTrackerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
