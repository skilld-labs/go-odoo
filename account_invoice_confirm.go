package odoo

// AccountInvoiceConfirm represents account.invoice.confirm model.
type AccountInvoiceConfirm struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountInvoiceConfirms represents array of account.invoice.confirm model.
type AccountInvoiceConfirms []AccountInvoiceConfirm

// AccountInvoiceConfirmModel is the odoo model name.
const AccountInvoiceConfirmModel = "account.invoice.confirm"

// Many2One convert AccountInvoiceConfirm to *Many2One.
func (aic *AccountInvoiceConfirm) Many2One() *Many2One {
	return NewMany2One(aic.Id.Get(), "")
}

// CreateAccountInvoiceConfirm creates a new account.invoice.confirm model and returns its id.
func (c *Client) CreateAccountInvoiceConfirm(aic *AccountInvoiceConfirm) (int64, error) {
	ids, err := c.CreateAccountInvoiceConfirms([]*AccountInvoiceConfirm{aic})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceConfirms creates a new account.invoice.confirm model and returns its id.
func (c *Client) CreateAccountInvoiceConfirms(aics []*AccountInvoiceConfirm) ([]int64, error) {
	var vv []interface{}
	for _, v := range aics {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceConfirmModel, vv, nil)
}

// UpdateAccountInvoiceConfirm updates an existing account.invoice.confirm record.
func (c *Client) UpdateAccountInvoiceConfirm(aic *AccountInvoiceConfirm) error {
	return c.UpdateAccountInvoiceConfirms([]int64{aic.Id.Get()}, aic)
}

// UpdateAccountInvoiceConfirms updates existing account.invoice.confirm records.
// All records (represented by ids) will be updated by aic values.
func (c *Client) UpdateAccountInvoiceConfirms(ids []int64, aic *AccountInvoiceConfirm) error {
	return c.Update(AccountInvoiceConfirmModel, ids, aic, nil)
}

// DeleteAccountInvoiceConfirm deletes an existing account.invoice.confirm record.
func (c *Client) DeleteAccountInvoiceConfirm(id int64) error {
	return c.DeleteAccountInvoiceConfirms([]int64{id})
}

// DeleteAccountInvoiceConfirms deletes existing account.invoice.confirm records.
func (c *Client) DeleteAccountInvoiceConfirms(ids []int64) error {
	return c.Delete(AccountInvoiceConfirmModel, ids)
}

// GetAccountInvoiceConfirm gets account.invoice.confirm existing record.
func (c *Client) GetAccountInvoiceConfirm(id int64) (*AccountInvoiceConfirm, error) {
	aics, err := c.GetAccountInvoiceConfirms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aics)[0]), nil
}

// GetAccountInvoiceConfirms gets account.invoice.confirm existing records.
func (c *Client) GetAccountInvoiceConfirms(ids []int64) (*AccountInvoiceConfirms, error) {
	aics := &AccountInvoiceConfirms{}
	if err := c.Read(AccountInvoiceConfirmModel, ids, nil, aics); err != nil {
		return nil, err
	}
	return aics, nil
}

// FindAccountInvoiceConfirm finds account.invoice.confirm record by querying it with criteria.
func (c *Client) FindAccountInvoiceConfirm(criteria *Criteria) (*AccountInvoiceConfirm, error) {
	aics := &AccountInvoiceConfirms{}
	if err := c.SearchRead(AccountInvoiceConfirmModel, criteria, NewOptions().Limit(1), aics); err != nil {
		return nil, err
	}
	return &((*aics)[0]), nil
}

// FindAccountInvoiceConfirms finds account.invoice.confirm records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceConfirms(criteria *Criteria, options *Options) (*AccountInvoiceConfirms, error) {
	aics := &AccountInvoiceConfirms{}
	if err := c.SearchRead(AccountInvoiceConfirmModel, criteria, options, aics); err != nil {
		return nil, err
	}
	return aics, nil
}

// FindAccountInvoiceConfirmIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceConfirmIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceConfirmModel, criteria, options)
}

// FindAccountInvoiceConfirmId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceConfirmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceConfirmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
