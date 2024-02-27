package odoo

// AccountRegisterPayments represents account.register.payments model.
type AccountRegisterPayments struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	Amount            *Float     `xmlrpc:"amount,omitempty"`
	Communication     *String    `xmlrpc:"communication,omitempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	HidePaymentMethod *Bool      `xmlrpc:"hide_payment_method,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	InvoiceIds        *Relation  `xmlrpc:"invoice_ids,omitempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omitempty"`
	Multi             *Bool      `xmlrpc:"multi,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerType       *Selection `xmlrpc:"partner_type,omitempty"`
	PaymentDate       *Time      `xmlrpc:"payment_date,omitempty"`
	PaymentMethodCode *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId   *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PaymentType       *Selection `xmlrpc:"payment_type,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountRegisterPaymentss represents array of account.register.payments model.
type AccountRegisterPaymentss []AccountRegisterPayments

// AccountRegisterPaymentsModel is the odoo model name.
const AccountRegisterPaymentsModel = "account.register.payments"

// Many2One convert AccountRegisterPayments to *Many2One.
func (arp *AccountRegisterPayments) Many2One() *Many2One {
	return NewMany2One(arp.Id.Get(), "")
}

// CreateAccountRegisterPayments creates a new account.register.payments model and returns its id.
func (c *Client) CreateAccountRegisterPayments(arp *AccountRegisterPayments) (int64, error) {
	ids, err := c.CreateAccountRegisterPaymentss([]*AccountRegisterPayments{arp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountRegisterPayments creates a new account.register.payments model and returns its id.
func (c *Client) CreateAccountRegisterPaymentss(arps []*AccountRegisterPayments) ([]int64, error) {
	var vv []interface{}
	for _, v := range arps {
		vv = append(vv, v)
	}
	return c.Create(AccountRegisterPaymentsModel, vv, nil)
}

// UpdateAccountRegisterPayments updates an existing account.register.payments record.
func (c *Client) UpdateAccountRegisterPayments(arp *AccountRegisterPayments) error {
	return c.UpdateAccountRegisterPaymentss([]int64{arp.Id.Get()}, arp)
}

// UpdateAccountRegisterPaymentss updates existing account.register.payments records.
// All records (represented by ids) will be updated by arp values.
func (c *Client) UpdateAccountRegisterPaymentss(ids []int64, arp *AccountRegisterPayments) error {
	return c.Update(AccountRegisterPaymentsModel, ids, arp, nil)
}

// DeleteAccountRegisterPayments deletes an existing account.register.payments record.
func (c *Client) DeleteAccountRegisterPayments(id int64) error {
	return c.DeleteAccountRegisterPaymentss([]int64{id})
}

// DeleteAccountRegisterPaymentss deletes existing account.register.payments records.
func (c *Client) DeleteAccountRegisterPaymentss(ids []int64) error {
	return c.Delete(AccountRegisterPaymentsModel, ids)
}

// GetAccountRegisterPayments gets account.register.payments existing record.
func (c *Client) GetAccountRegisterPayments(id int64) (*AccountRegisterPayments, error) {
	arps, err := c.GetAccountRegisterPaymentss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arps)[0]), nil
}

// GetAccountRegisterPaymentss gets account.register.payments existing records.
func (c *Client) GetAccountRegisterPaymentss(ids []int64) (*AccountRegisterPaymentss, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.Read(AccountRegisterPaymentsModel, ids, nil, arps); err != nil {
		return nil, err
	}
	return arps, nil
}

// FindAccountRegisterPayments finds account.register.payments record by querying it with criteria.
func (c *Client) FindAccountRegisterPayments(criteria *Criteria) (*AccountRegisterPayments, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.SearchRead(AccountRegisterPaymentsModel, criteria, NewOptions().Limit(1), arps); err != nil {
		return nil, err
	}
	return &((*arps)[0]), nil
}

// FindAccountRegisterPaymentss finds account.register.payments records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRegisterPaymentss(criteria *Criteria, options *Options) (*AccountRegisterPaymentss, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.SearchRead(AccountRegisterPaymentsModel, criteria, options, arps); err != nil {
		return nil, err
	}
	return arps, nil
}

// FindAccountRegisterPaymentsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRegisterPaymentsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountRegisterPaymentsModel, criteria, options)
}

// FindAccountRegisterPaymentsId finds record id by querying it with criteria.
func (c *Client) FindAccountRegisterPaymentsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountRegisterPaymentsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
