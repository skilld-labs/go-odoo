package odoo

// SmsAccountSender represents sms.account.sender model.
type SmsAccountSender struct {
	AccountId   *Many2One `xmlrpc:"account_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	SenderName  *String   `xmlrpc:"sender_name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsAccountSenders represents array of sms.account.sender model.
type SmsAccountSenders []SmsAccountSender

// SmsAccountSenderModel is the odoo model name.
const SmsAccountSenderModel = "sms.account.sender"

// Many2One convert SmsAccountSender to *Many2One.
func (sas *SmsAccountSender) Many2One() *Many2One {
	return NewMany2One(sas.Id.Get(), "")
}

// CreateSmsAccountSender creates a new sms.account.sender model and returns its id.
func (c *Client) CreateSmsAccountSender(sas *SmsAccountSender) (int64, error) {
	ids, err := c.CreateSmsAccountSenders([]*SmsAccountSender{sas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsAccountSender creates a new sms.account.sender model and returns its id.
func (c *Client) CreateSmsAccountSenders(sass []*SmsAccountSender) ([]int64, error) {
	var vv []interface{}
	for _, v := range sass {
		vv = append(vv, v)
	}
	return c.Create(SmsAccountSenderModel, vv, nil)
}

// UpdateSmsAccountSender updates an existing sms.account.sender record.
func (c *Client) UpdateSmsAccountSender(sas *SmsAccountSender) error {
	return c.UpdateSmsAccountSenders([]int64{sas.Id.Get()}, sas)
}

// UpdateSmsAccountSenders updates existing sms.account.sender records.
// All records (represented by ids) will be updated by sas values.
func (c *Client) UpdateSmsAccountSenders(ids []int64, sas *SmsAccountSender) error {
	return c.Update(SmsAccountSenderModel, ids, sas, nil)
}

// DeleteSmsAccountSender deletes an existing sms.account.sender record.
func (c *Client) DeleteSmsAccountSender(id int64) error {
	return c.DeleteSmsAccountSenders([]int64{id})
}

// DeleteSmsAccountSenders deletes existing sms.account.sender records.
func (c *Client) DeleteSmsAccountSenders(ids []int64) error {
	return c.Delete(SmsAccountSenderModel, ids)
}

// GetSmsAccountSender gets sms.account.sender existing record.
func (c *Client) GetSmsAccountSender(id int64) (*SmsAccountSender, error) {
	sass, err := c.GetSmsAccountSenders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sass)[0]), nil
}

// GetSmsAccountSenders gets sms.account.sender existing records.
func (c *Client) GetSmsAccountSenders(ids []int64) (*SmsAccountSenders, error) {
	sass := &SmsAccountSenders{}
	if err := c.Read(SmsAccountSenderModel, ids, nil, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindSmsAccountSender finds sms.account.sender record by querying it with criteria.
func (c *Client) FindSmsAccountSender(criteria *Criteria) (*SmsAccountSender, error) {
	sass := &SmsAccountSenders{}
	if err := c.SearchRead(SmsAccountSenderModel, criteria, NewOptions().Limit(1), sass); err != nil {
		return nil, err
	}
	return &((*sass)[0]), nil
}

// FindSmsAccountSenders finds sms.account.sender records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountSenders(criteria *Criteria, options *Options) (*SmsAccountSenders, error) {
	sass := &SmsAccountSenders{}
	if err := c.SearchRead(SmsAccountSenderModel, criteria, options, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindSmsAccountSenderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountSenderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsAccountSenderModel, criteria, options)
}

// FindSmsAccountSenderId finds record id by querying it with criteria.
func (c *Client) FindSmsAccountSenderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsAccountSenderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
