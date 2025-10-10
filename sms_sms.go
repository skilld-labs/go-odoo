package odoo

// SmsSms represents sms.sms model.
type SmsSms struct {
	Body          *String    `xmlrpc:"body,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	FailureType   *Selection `xmlrpc:"failure_type,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	MailMessageId *Many2One  `xmlrpc:"mail_message_id,omitempty"`
	Number        *String    `xmlrpc:"number,omitempty"`
	PartnerId     *Many2One  `xmlrpc:"partner_id,omitempty"`
	SmsTrackerId  *Many2One  `xmlrpc:"sms_tracker_id,omitempty"`
	State         *Selection `xmlrpc:"state,omitempty"`
	ToDelete      *Bool      `xmlrpc:"to_delete,omitempty"`
	Uuid          *String    `xmlrpc:"uuid,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SmsSmss represents array of sms.sms model.
type SmsSmss []SmsSms

// SmsSmsModel is the odoo model name.
const SmsSmsModel = "sms.sms"

// Many2One convert SmsSms to *Many2One.
func (ss *SmsSms) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSmsSms creates a new sms.sms model and returns its id.
func (c *Client) CreateSmsSms(ss *SmsSms) (int64, error) {
	ids, err := c.CreateSmsSmss([]*SmsSms{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsSms creates a new sms.sms model and returns its id.
func (c *Client) CreateSmsSmss(sss []*SmsSms) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SmsSmsModel, vv, nil)
}

// UpdateSmsSms updates an existing sms.sms record.
func (c *Client) UpdateSmsSms(ss *SmsSms) error {
	return c.UpdateSmsSmss([]int64{ss.Id.Get()}, ss)
}

// UpdateSmsSmss updates existing sms.sms records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSmsSmss(ids []int64, ss *SmsSms) error {
	return c.Update(SmsSmsModel, ids, ss, nil)
}

// DeleteSmsSms deletes an existing sms.sms record.
func (c *Client) DeleteSmsSms(id int64) error {
	return c.DeleteSmsSmss([]int64{id})
}

// DeleteSmsSmss deletes existing sms.sms records.
func (c *Client) DeleteSmsSmss(ids []int64) error {
	return c.Delete(SmsSmsModel, ids)
}

// GetSmsSms gets sms.sms existing record.
func (c *Client) GetSmsSms(id int64) (*SmsSms, error) {
	sss, err := c.GetSmsSmss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// GetSmsSmss gets sms.sms existing records.
func (c *Client) GetSmsSmss(ids []int64) (*SmsSmss, error) {
	sss := &SmsSmss{}
	if err := c.Read(SmsSmsModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSms finds sms.sms record by querying it with criteria.
func (c *Client) FindSmsSms(criteria *Criteria) (*SmsSms, error) {
	sss := &SmsSmss{}
	if err := c.SearchRead(SmsSmsModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// FindSmsSmss finds sms.sms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSmss(criteria *Criteria, options *Options) (*SmsSmss, error) {
	sss := &SmsSmss{}
	if err := c.SearchRead(SmsSmsModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSmsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSmsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsSmsModel, criteria, options)
}

// FindSmsSmsId finds record id by querying it with criteria.
func (c *Client) FindSmsSmsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsSmsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
