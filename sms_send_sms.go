package odoo

// SmsSendSms represents sms.send_sms model.
type SmsSendSms struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Message     *String   `xmlrpc:"message,omitempty"`
	Recipients  *String   `xmlrpc:"recipients,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsSendSmss represents array of sms.send_sms model.
type SmsSendSmss []SmsSendSms

// SmsSendSmsModel is the odoo model name.
const SmsSendSmsModel = "sms.send_sms"

// Many2One convert SmsSendSms to *Many2One.
func (ss *SmsSendSms) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSmsSendSms creates a new sms.send_sms model and returns its id.
func (c *Client) CreateSmsSendSms(ss *SmsSendSms) (int64, error) {
	ids, err := c.CreateSmsSendSmss([]*SmsSendSms{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsSendSmss creates a new sms.send_sms model and returns its id.
func (c *Client) CreateSmsSendSmss(sss []*SmsSendSms) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SmsSendSmsModel, vv, nil)
}

// UpdateSmsSendSms updates an existing sms.send_sms record.
func (c *Client) UpdateSmsSendSms(ss *SmsSendSms) error {
	return c.UpdateSmsSendSmss([]int64{ss.Id.Get()}, ss)
}

// UpdateSmsSendSmss updates existing sms.send_sms records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSmsSendSmss(ids []int64, ss *SmsSendSms) error {
	return c.Update(SmsSendSmsModel, ids, ss, nil)
}

// DeleteSmsSendSms deletes an existing sms.send_sms record.
func (c *Client) DeleteSmsSendSms(id int64) error {
	return c.DeleteSmsSendSmss([]int64{id})
}

// DeleteSmsSendSmss deletes existing sms.send_sms records.
func (c *Client) DeleteSmsSendSmss(ids []int64) error {
	return c.Delete(SmsSendSmsModel, ids)
}

// GetSmsSendSms gets sms.send_sms existing record.
func (c *Client) GetSmsSendSms(id int64) (*SmsSendSms, error) {
	sss, err := c.GetSmsSendSmss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// GetSmsSendSmss gets sms.send_sms existing records.
func (c *Client) GetSmsSendSmss(ids []int64) (*SmsSendSmss, error) {
	sss := &SmsSendSmss{}
	if err := c.Read(SmsSendSmsModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSendSms finds sms.send_sms record by querying it with criteria.
func (c *Client) FindSmsSendSms(criteria *Criteria) (*SmsSendSms, error) {
	sss := &SmsSendSmss{}
	if err := c.SearchRead(SmsSendSmsModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// FindSmsSendSmss finds sms.send_sms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSendSmss(criteria *Criteria, options *Options) (*SmsSendSmss, error) {
	sss := &SmsSendSmss{}
	if err := c.SearchRead(SmsSendSmsModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSendSmsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSendSmsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsSendSmsModel, criteria, options)
}

// FindSmsSendSmsId finds record id by querying it with criteria.
func (c *Client) FindSmsSendSmsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsSendSmsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
