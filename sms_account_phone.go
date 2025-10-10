package odoo

// SmsAccountPhone represents sms.account.phone model.
type SmsAccountPhone struct {
	AccountId   *Many2One `xmlrpc:"account_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PhoneNumber *String   `xmlrpc:"phone_number,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsAccountPhones represents array of sms.account.phone model.
type SmsAccountPhones []SmsAccountPhone

// SmsAccountPhoneModel is the odoo model name.
const SmsAccountPhoneModel = "sms.account.phone"

// Many2One convert SmsAccountPhone to *Many2One.
func (sap *SmsAccountPhone) Many2One() *Many2One {
	return NewMany2One(sap.Id.Get(), "")
}

// CreateSmsAccountPhone creates a new sms.account.phone model and returns its id.
func (c *Client) CreateSmsAccountPhone(sap *SmsAccountPhone) (int64, error) {
	ids, err := c.CreateSmsAccountPhones([]*SmsAccountPhone{sap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsAccountPhone creates a new sms.account.phone model and returns its id.
func (c *Client) CreateSmsAccountPhones(saps []*SmsAccountPhone) ([]int64, error) {
	var vv []interface{}
	for _, v := range saps {
		vv = append(vv, v)
	}
	return c.Create(SmsAccountPhoneModel, vv, nil)
}

// UpdateSmsAccountPhone updates an existing sms.account.phone record.
func (c *Client) UpdateSmsAccountPhone(sap *SmsAccountPhone) error {
	return c.UpdateSmsAccountPhones([]int64{sap.Id.Get()}, sap)
}

// UpdateSmsAccountPhones updates existing sms.account.phone records.
// All records (represented by ids) will be updated by sap values.
func (c *Client) UpdateSmsAccountPhones(ids []int64, sap *SmsAccountPhone) error {
	return c.Update(SmsAccountPhoneModel, ids, sap, nil)
}

// DeleteSmsAccountPhone deletes an existing sms.account.phone record.
func (c *Client) DeleteSmsAccountPhone(id int64) error {
	return c.DeleteSmsAccountPhones([]int64{id})
}

// DeleteSmsAccountPhones deletes existing sms.account.phone records.
func (c *Client) DeleteSmsAccountPhones(ids []int64) error {
	return c.Delete(SmsAccountPhoneModel, ids)
}

// GetSmsAccountPhone gets sms.account.phone existing record.
func (c *Client) GetSmsAccountPhone(id int64) (*SmsAccountPhone, error) {
	saps, err := c.GetSmsAccountPhones([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*saps)[0]), nil
}

// GetSmsAccountPhones gets sms.account.phone existing records.
func (c *Client) GetSmsAccountPhones(ids []int64) (*SmsAccountPhones, error) {
	saps := &SmsAccountPhones{}
	if err := c.Read(SmsAccountPhoneModel, ids, nil, saps); err != nil {
		return nil, err
	}
	return saps, nil
}

// FindSmsAccountPhone finds sms.account.phone record by querying it with criteria.
func (c *Client) FindSmsAccountPhone(criteria *Criteria) (*SmsAccountPhone, error) {
	saps := &SmsAccountPhones{}
	if err := c.SearchRead(SmsAccountPhoneModel, criteria, NewOptions().Limit(1), saps); err != nil {
		return nil, err
	}
	return &((*saps)[0]), nil
}

// FindSmsAccountPhones finds sms.account.phone records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountPhones(criteria *Criteria, options *Options) (*SmsAccountPhones, error) {
	saps := &SmsAccountPhones{}
	if err := c.SearchRead(SmsAccountPhoneModel, criteria, options, saps); err != nil {
		return nil, err
	}
	return saps, nil
}

// FindSmsAccountPhoneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountPhoneIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsAccountPhoneModel, criteria, options)
}

// FindSmsAccountPhoneId finds record id by querying it with criteria.
func (c *Client) FindSmsAccountPhoneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsAccountPhoneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
