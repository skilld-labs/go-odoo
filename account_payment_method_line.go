package odoo

// AccountPaymentMethodLine represents account.payment.method.line model.
type AccountPaymentMethodLine struct {
	AvailablePaymentMethodIds *Relation  `xmlrpc:"available_payment_method_ids,omitempty"`
	Code                      *String    `xmlrpc:"code,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultAccountId          *Many2One  `xmlrpc:"default_account_id,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PaymentAccountId          *Many2One  `xmlrpc:"payment_account_id,omitempty"`
	PaymentMethodId           *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PaymentProviderId         *Many2One  `xmlrpc:"payment_provider_id,omitempty"`
	PaymentProviderState      *Selection `xmlrpc:"payment_provider_state,omitempty"`
	PaymentType               *Selection `xmlrpc:"payment_type,omitempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountPaymentMethodLines represents array of account.payment.method.line model.
type AccountPaymentMethodLines []AccountPaymentMethodLine

// AccountPaymentMethodLineModel is the odoo model name.
const AccountPaymentMethodLineModel = "account.payment.method.line"

// Many2One convert AccountPaymentMethodLine to *Many2One.
func (apml *AccountPaymentMethodLine) Many2One() *Many2One {
	return NewMany2One(apml.Id.Get(), "")
}

// CreateAccountPaymentMethodLine creates a new account.payment.method.line model and returns its id.
func (c *Client) CreateAccountPaymentMethodLine(apml *AccountPaymentMethodLine) (int64, error) {
	ids, err := c.CreateAccountPaymentMethodLines([]*AccountPaymentMethodLine{apml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPaymentMethodLine creates a new account.payment.method.line model and returns its id.
func (c *Client) CreateAccountPaymentMethodLines(apmls []*AccountPaymentMethodLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range apmls {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentMethodLineModel, vv, nil)
}

// UpdateAccountPaymentMethodLine updates an existing account.payment.method.line record.
func (c *Client) UpdateAccountPaymentMethodLine(apml *AccountPaymentMethodLine) error {
	return c.UpdateAccountPaymentMethodLines([]int64{apml.Id.Get()}, apml)
}

// UpdateAccountPaymentMethodLines updates existing account.payment.method.line records.
// All records (represented by ids) will be updated by apml values.
func (c *Client) UpdateAccountPaymentMethodLines(ids []int64, apml *AccountPaymentMethodLine) error {
	return c.Update(AccountPaymentMethodLineModel, ids, apml, nil)
}

// DeleteAccountPaymentMethodLine deletes an existing account.payment.method.line record.
func (c *Client) DeleteAccountPaymentMethodLine(id int64) error {
	return c.DeleteAccountPaymentMethodLines([]int64{id})
}

// DeleteAccountPaymentMethodLines deletes existing account.payment.method.line records.
func (c *Client) DeleteAccountPaymentMethodLines(ids []int64) error {
	return c.Delete(AccountPaymentMethodLineModel, ids)
}

// GetAccountPaymentMethodLine gets account.payment.method.line existing record.
func (c *Client) GetAccountPaymentMethodLine(id int64) (*AccountPaymentMethodLine, error) {
	apmls, err := c.GetAccountPaymentMethodLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apmls)[0]), nil
}

// GetAccountPaymentMethodLines gets account.payment.method.line existing records.
func (c *Client) GetAccountPaymentMethodLines(ids []int64) (*AccountPaymentMethodLines, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.Read(AccountPaymentMethodLineModel, ids, nil, apmls); err != nil {
		return nil, err
	}
	return apmls, nil
}

// FindAccountPaymentMethodLine finds account.payment.method.line record by querying it with criteria.
func (c *Client) FindAccountPaymentMethodLine(criteria *Criteria) (*AccountPaymentMethodLine, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.SearchRead(AccountPaymentMethodLineModel, criteria, NewOptions().Limit(1), apmls); err != nil {
		return nil, err
	}
	return &((*apmls)[0]), nil
}

// FindAccountPaymentMethodLines finds account.payment.method.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethodLines(criteria *Criteria, options *Options) (*AccountPaymentMethodLines, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.SearchRead(AccountPaymentMethodLineModel, criteria, options, apmls); err != nil {
		return nil, err
	}
	return apmls, nil
}

// FindAccountPaymentMethodLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethodLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentMethodLineModel, criteria, options)
}

// FindAccountPaymentMethodLineId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentMethodLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentMethodLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
