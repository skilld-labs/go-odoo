package odoo

// AccountInvoiceRefund represents account.invoice.refund model.
type AccountInvoiceRefund struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date         *Time      `xmlrpc:"date,omitempty"`
	DateInvoice  *Time      `xmlrpc:"date_invoice,omitempty"`
	Description  *String    `xmlrpc:"description,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	FilterRefund *Selection `xmlrpc:"filter_refund,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	RefundOnly   *Bool      `xmlrpc:"refund_only,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountInvoiceRefunds represents array of account.invoice.refund model.
type AccountInvoiceRefunds []AccountInvoiceRefund

// AccountInvoiceRefundModel is the odoo model name.
const AccountInvoiceRefundModel = "account.invoice.refund"

// Many2One convert AccountInvoiceRefund to *Many2One.
func (air *AccountInvoiceRefund) Many2One() *Many2One {
	return NewMany2One(air.Id.Get(), "")
}

// CreateAccountInvoiceRefund creates a new account.invoice.refund model and returns its id.
func (c *Client) CreateAccountInvoiceRefund(air *AccountInvoiceRefund) (int64, error) {
	ids, err := c.CreateAccountInvoiceRefunds([]*AccountInvoiceRefund{air})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceRefund creates a new account.invoice.refund model and returns its id.
func (c *Client) CreateAccountInvoiceRefunds(airs []*AccountInvoiceRefund) ([]int64, error) {
	var vv []interface{}
	for _, v := range airs {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceRefundModel, vv, nil)
}

// UpdateAccountInvoiceRefund updates an existing account.invoice.refund record.
func (c *Client) UpdateAccountInvoiceRefund(air *AccountInvoiceRefund) error {
	return c.UpdateAccountInvoiceRefunds([]int64{air.Id.Get()}, air)
}

// UpdateAccountInvoiceRefunds updates existing account.invoice.refund records.
// All records (represented by ids) will be updated by air values.
func (c *Client) UpdateAccountInvoiceRefunds(ids []int64, air *AccountInvoiceRefund) error {
	return c.Update(AccountInvoiceRefundModel, ids, air, nil)
}

// DeleteAccountInvoiceRefund deletes an existing account.invoice.refund record.
func (c *Client) DeleteAccountInvoiceRefund(id int64) error {
	return c.DeleteAccountInvoiceRefunds([]int64{id})
}

// DeleteAccountInvoiceRefunds deletes existing account.invoice.refund records.
func (c *Client) DeleteAccountInvoiceRefunds(ids []int64) error {
	return c.Delete(AccountInvoiceRefundModel, ids)
}

// GetAccountInvoiceRefund gets account.invoice.refund existing record.
func (c *Client) GetAccountInvoiceRefund(id int64) (*AccountInvoiceRefund, error) {
	airs, err := c.GetAccountInvoiceRefunds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// GetAccountInvoiceRefunds gets account.invoice.refund existing records.
func (c *Client) GetAccountInvoiceRefunds(ids []int64) (*AccountInvoiceRefunds, error) {
	airs := &AccountInvoiceRefunds{}
	if err := c.Read(AccountInvoiceRefundModel, ids, nil, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceRefund finds account.invoice.refund record by querying it with criteria.
func (c *Client) FindAccountInvoiceRefund(criteria *Criteria) (*AccountInvoiceRefund, error) {
	airs := &AccountInvoiceRefunds{}
	if err := c.SearchRead(AccountInvoiceRefundModel, criteria, NewOptions().Limit(1), airs); err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// FindAccountInvoiceRefunds finds account.invoice.refund records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceRefunds(criteria *Criteria, options *Options) (*AccountInvoiceRefunds, error) {
	airs := &AccountInvoiceRefunds{}
	if err := c.SearchRead(AccountInvoiceRefundModel, criteria, options, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceRefundIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceRefundIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceRefundModel, criteria, options)
}

// FindAccountInvoiceRefundId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceRefundId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceRefundModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
