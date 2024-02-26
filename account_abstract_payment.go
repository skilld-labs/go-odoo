package odoo

// AccountAbstractPayment represents account.abstract.payment model.
type AccountAbstractPayment struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Amount            *Float     `xmlrpc:"amount,omptempty"`
	Communication     *String    `xmlrpc:"communication,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	HidePaymentMethod *Bool      `xmlrpc:"hide_payment_method,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerType       *Selection `xmlrpc:"partner_type,omptempty"`
	PaymentDate       *Time      `xmlrpc:"payment_date,omptempty"`
	PaymentMethodCode *String    `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodId   *Many2One  `xmlrpc:"payment_method_id,omptempty"`
	PaymentType       *Selection `xmlrpc:"payment_type,omptempty"`
}

// AccountAbstractPayments represents array of account.abstract.payment model.
type AccountAbstractPayments []AccountAbstractPayment

// AccountAbstractPaymentModel is the odoo model name.
const AccountAbstractPaymentModel = "account.abstract.payment"

// Many2One convert AccountAbstractPayment to *Many2One.
func (aap *AccountAbstractPayment) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAbstractPayment creates a new account.abstract.payment model and returns its id.
func (c *Client) CreateAccountAbstractPayment(aap *AccountAbstractPayment) (int64, error) {
	ids, err := c.CreateAccountAbstractPayments([]*AccountAbstractPayment{aap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAbstractPayment creates a new account.abstract.payment model and returns its id.
func (c *Client) CreateAccountAbstractPayments(aaps []*AccountAbstractPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaps {
		vv = append(vv, v)
	}
	return c.Create(AccountAbstractPaymentModel, vv, nil)
}

// UpdateAccountAbstractPayment updates an existing account.abstract.payment record.
func (c *Client) UpdateAccountAbstractPayment(aap *AccountAbstractPayment) error {
	return c.UpdateAccountAbstractPayments([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAbstractPayments updates existing account.abstract.payment records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAbstractPayments(ids []int64, aap *AccountAbstractPayment) error {
	return c.Update(AccountAbstractPaymentModel, ids, aap, nil)
}

// DeleteAccountAbstractPayment deletes an existing account.abstract.payment record.
func (c *Client) DeleteAccountAbstractPayment(id int64) error {
	return c.DeleteAccountAbstractPayments([]int64{id})
}

// DeleteAccountAbstractPayments deletes existing account.abstract.payment records.
func (c *Client) DeleteAccountAbstractPayments(ids []int64) error {
	return c.Delete(AccountAbstractPaymentModel, ids)
}

// GetAccountAbstractPayment gets account.abstract.payment existing record.
func (c *Client) GetAccountAbstractPayment(id int64) (*AccountAbstractPayment, error) {
	aaps, err := c.GetAccountAbstractPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaps)[0]), nil
}

// GetAccountAbstractPayments gets account.abstract.payment existing records.
func (c *Client) GetAccountAbstractPayments(ids []int64) (*AccountAbstractPayments, error) {
	aaps := &AccountAbstractPayments{}
	if err := c.Read(AccountAbstractPaymentModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAbstractPayment finds account.abstract.payment record by querying it with criteria.
func (c *Client) FindAccountAbstractPayment(criteria *Criteria) (*AccountAbstractPayment, error) {
	aaps := &AccountAbstractPayments{}
	if err := c.SearchRead(AccountAbstractPaymentModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	return &((*aaps)[0]), nil
}

// FindAccountAbstractPayments finds account.abstract.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAbstractPayments(criteria *Criteria, options *Options) (*AccountAbstractPayments, error) {
	aaps := &AccountAbstractPayments{}
	if err := c.SearchRead(AccountAbstractPaymentModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAbstractPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAbstractPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAbstractPaymentModel, criteria, options)
}

// FindAccountAbstractPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountAbstractPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAbstractPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
