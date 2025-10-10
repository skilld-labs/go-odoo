package odoo

// SmsResend represents sms.resend model.
type SmsResend struct {
	CanCancel              *Bool     `xmlrpc:"can_cancel,omitempty"`
	CanResend              *Bool     `xmlrpc:"can_resend,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	HasInsufficientCredit  *Bool     `xmlrpc:"has_insufficient_credit,omitempty"`
	HasUnregisteredAccount *Bool     `xmlrpc:"has_unregistered_account,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	MailMessageId          *Many2One `xmlrpc:"mail_message_id,omitempty"`
	RecipientIds           *Relation `xmlrpc:"recipient_ids,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SmsResends represents array of sms.resend model.
type SmsResends []SmsResend

// SmsResendModel is the odoo model name.
const SmsResendModel = "sms.resend"

// Many2One convert SmsResend to *Many2One.
func (sr *SmsResend) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSmsResend creates a new sms.resend model and returns its id.
func (c *Client) CreateSmsResend(sr *SmsResend) (int64, error) {
	ids, err := c.CreateSmsResends([]*SmsResend{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSmsResend creates a new sms.resend model and returns its id.
func (c *Client) CreateSmsResends(srs []*SmsResend) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(SmsResendModel, vv, nil)
}

// UpdateSmsResend updates an existing sms.resend record.
func (c *Client) UpdateSmsResend(sr *SmsResend) error {
	return c.UpdateSmsResends([]int64{sr.Id.Get()}, sr)
}

// UpdateSmsResends updates existing sms.resend records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSmsResends(ids []int64, sr *SmsResend) error {
	return c.Update(SmsResendModel, ids, sr, nil)
}

// DeleteSmsResend deletes an existing sms.resend record.
func (c *Client) DeleteSmsResend(id int64) error {
	return c.DeleteSmsResends([]int64{id})
}

// DeleteSmsResends deletes existing sms.resend records.
func (c *Client) DeleteSmsResends(ids []int64) error {
	return c.Delete(SmsResendModel, ids)
}

// GetSmsResend gets sms.resend existing record.
func (c *Client) GetSmsResend(id int64) (*SmsResend, error) {
	srs, err := c.GetSmsResends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetSmsResends gets sms.resend existing records.
func (c *Client) GetSmsResends(ids []int64) (*SmsResends, error) {
	srs := &SmsResends{}
	if err := c.Read(SmsResendModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSmsResend finds sms.resend record by querying it with criteria.
func (c *Client) FindSmsResend(criteria *Criteria) (*SmsResend, error) {
	srs := &SmsResends{}
	if err := c.SearchRead(SmsResendModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindSmsResends finds sms.resend records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResends(criteria *Criteria, options *Options) (*SmsResends, error) {
	srs := &SmsResends{}
	if err := c.SearchRead(SmsResendModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSmsResendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SmsResendModel, criteria, options)
}

// FindSmsResendId finds record id by querying it with criteria.
func (c *Client) FindSmsResendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsResendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
