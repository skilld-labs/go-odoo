package odoo

// SmsAccountCode represents sms.account.code model.
type SmsAccountCode struct {
	AccountId        *Many2One `xmlrpc:"account_id,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	VerificationCode *String   `xmlrpc:"verification_code,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsAccountCodes represents array of sms.account.code model.
type SmsAccountCodes []SmsAccountCode

// SmsAccountCodeModel is the odoo model name.
const SmsAccountCodeModel = "sms.account.code"

// Many2One convert SmsAccountCode to *Many2One.
func (sac *SmsAccountCode) Many2One() *Many2One {
	return NewMany2One(sac.Id.Get(), "")
}

// CreateSmsAccountCode creates a new sms.account.code model and returns its id.
func (c *Client) CreateSmsAccountCode(sac *SmsAccountCode) (int64, error) {
	ids, err := c.CreateSmsAccountCodes([]*SmsAccountCode{sac})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsAccountCode creates a new sms.account.code model and returns its id.
func (c *Client) CreateSmsAccountCodes(sacs []*SmsAccountCode) ([]int64, error) {
	var vv []interface{}
	for _, v := range sacs {
		vv = append(vv, v)
	}
	return c.Create(SmsAccountCodeModel, vv, nil)
}

// UpdateSmsAccountCode updates an existing sms.account.code record.
func (c *Client) UpdateSmsAccountCode(sac *SmsAccountCode) error {
	return c.UpdateSmsAccountCodes([]int64{sac.Id.Get()}, sac)
}

// UpdateSmsAccountCodes updates existing sms.account.code records.
// All records (represented by ids) will be updated by sac values.
func (c *Client) UpdateSmsAccountCodes(ids []int64, sac *SmsAccountCode) error {
	return c.Update(SmsAccountCodeModel, ids, sac, nil)
}

// DeleteSmsAccountCode deletes an existing sms.account.code record.
func (c *Client) DeleteSmsAccountCode(id int64) error {
	return c.DeleteSmsAccountCodes([]int64{id})
}

// DeleteSmsAccountCodes deletes existing sms.account.code records.
func (c *Client) DeleteSmsAccountCodes(ids []int64) error {
	return c.Delete(SmsAccountCodeModel, ids)
}

// GetSmsAccountCode gets sms.account.code existing record.
func (c *Client) GetSmsAccountCode(id int64) (*SmsAccountCode, error) {
	sacs, err := c.GetSmsAccountCodes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sacs)[0]), nil
}

// GetSmsAccountCodes gets sms.account.code existing records.
func (c *Client) GetSmsAccountCodes(ids []int64) (*SmsAccountCodes, error) {
	sacs := &SmsAccountCodes{}
	if err := c.Read(SmsAccountCodeModel, ids, nil, sacs); err != nil {
		return nil, err
	}
	return sacs, nil
}

// FindSmsAccountCode finds sms.account.code record by querying it with criteria.
func (c *Client) FindSmsAccountCode(criteria *Criteria) (*SmsAccountCode, error) {
	sacs := &SmsAccountCodes{}
	if err := c.SearchRead(SmsAccountCodeModel, criteria, NewOptions().Limit(1), sacs); err != nil {
		return nil, err
	}
	return &((*sacs)[0]), nil
}

// FindSmsAccountCodes finds sms.account.code records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountCodes(criteria *Criteria, options *Options) (*SmsAccountCodes, error) {
	sacs := &SmsAccountCodes{}
	if err := c.SearchRead(SmsAccountCodeModel, criteria, options, sacs); err != nil {
		return nil, err
	}
	return sacs, nil
}

// FindSmsAccountCodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsAccountCodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsAccountCodeModel, criteria, options)
}

// FindSmsAccountCodeId finds record id by querying it with criteria.
func (c *Client) FindSmsAccountCodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsAccountCodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
