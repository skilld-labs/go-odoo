package odoo

// AccountPaymentMethod represents account.payment.method model.
type AccountPaymentMethod struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	Code        *String    `xmlrpc:"code,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	PaymentType *Selection `xmlrpc:"payment_type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountPaymentMethods represents array of account.payment.method model.
type AccountPaymentMethods []AccountPaymentMethod

// AccountPaymentMethodModel is the odoo model name.
const AccountPaymentMethodModel = "account.payment.method"

// Many2One convert AccountPaymentMethod to *Many2One.
func (apm *AccountPaymentMethod) Many2One() *Many2One {
	return NewMany2One(apm.Id.Get(), "")
}

// CreateAccountPaymentMethod creates a new account.payment.method model and returns its id.
func (c *Client) CreateAccountPaymentMethod(apm *AccountPaymentMethod) (int64, error) {
	ids, err := c.CreateAccountPaymentMethods([]*AccountPaymentMethod{apm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPaymentMethod creates a new account.payment.method model and returns its id.
func (c *Client) CreateAccountPaymentMethods(apms []*AccountPaymentMethod) ([]int64, error) {
	var vv []interface{}
	for _, v := range apms {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentMethodModel, vv, nil)
}

// UpdateAccountPaymentMethod updates an existing account.payment.method record.
func (c *Client) UpdateAccountPaymentMethod(apm *AccountPaymentMethod) error {
	return c.UpdateAccountPaymentMethods([]int64{apm.Id.Get()}, apm)
}

// UpdateAccountPaymentMethods updates existing account.payment.method records.
// All records (represented by ids) will be updated by apm values.
func (c *Client) UpdateAccountPaymentMethods(ids []int64, apm *AccountPaymentMethod) error {
	return c.Update(AccountPaymentMethodModel, ids, apm, nil)
}

// DeleteAccountPaymentMethod deletes an existing account.payment.method record.
func (c *Client) DeleteAccountPaymentMethod(id int64) error {
	return c.DeleteAccountPaymentMethods([]int64{id})
}

// DeleteAccountPaymentMethods deletes existing account.payment.method records.
func (c *Client) DeleteAccountPaymentMethods(ids []int64) error {
	return c.Delete(AccountPaymentMethodModel, ids)
}

// GetAccountPaymentMethod gets account.payment.method existing record.
func (c *Client) GetAccountPaymentMethod(id int64) (*AccountPaymentMethod, error) {
	apms, err := c.GetAccountPaymentMethods([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apms)[0]), nil
}

// GetAccountPaymentMethods gets account.payment.method existing records.
func (c *Client) GetAccountPaymentMethods(ids []int64) (*AccountPaymentMethods, error) {
	apms := &AccountPaymentMethods{}
	if err := c.Read(AccountPaymentMethodModel, ids, nil, apms); err != nil {
		return nil, err
	}
	return apms, nil
}

// FindAccountPaymentMethod finds account.payment.method record by querying it with criteria.
func (c *Client) FindAccountPaymentMethod(criteria *Criteria) (*AccountPaymentMethod, error) {
	apms := &AccountPaymentMethods{}
	if err := c.SearchRead(AccountPaymentMethodModel, criteria, NewOptions().Limit(1), apms); err != nil {
		return nil, err
	}
	return &((*apms)[0]), nil
}

// FindAccountPaymentMethods finds account.payment.method records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethods(criteria *Criteria, options *Options) (*AccountPaymentMethods, error) {
	apms := &AccountPaymentMethods{}
	if err := c.SearchRead(AccountPaymentMethodModel, criteria, options, apms); err != nil {
		return nil, err
	}
	return apms, nil
}

// FindAccountPaymentMethodIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethodIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentMethodModel, criteria, options)
}

// FindAccountPaymentMethodId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentMethodId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentMethodModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
