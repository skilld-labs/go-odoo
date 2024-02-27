package odoo

// AccountPartialReconcile represents account.partial.reconcile model.
type AccountPartialReconcile struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	Amount            *Float    `xmlrpc:"amount,omitempty"`
	AmountCurrency    *Float    `xmlrpc:"amount_currency,omitempty"`
	CompanyCurrencyId *Many2One `xmlrpc:"company_currency_id,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	CreditMoveId      *Many2One `xmlrpc:"credit_move_id,omitempty"`
	CurrencyId        *Many2One `xmlrpc:"currency_id,omitempty"`
	DebitMoveId       *Many2One `xmlrpc:"debit_move_id,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	FullReconcileId   *Many2One `xmlrpc:"full_reconcile_id,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	MaxDate           *Time     `xmlrpc:"max_date,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountPartialReconciles represents array of account.partial.reconcile model.
type AccountPartialReconciles []AccountPartialReconcile

// AccountPartialReconcileModel is the odoo model name.
const AccountPartialReconcileModel = "account.partial.reconcile"

// Many2One convert AccountPartialReconcile to *Many2One.
func (apr *AccountPartialReconcile) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPartialReconcile creates a new account.partial.reconcile model and returns its id.
func (c *Client) CreateAccountPartialReconcile(apr *AccountPartialReconcile) (int64, error) {
	ids, err := c.CreateAccountPartialReconciles([]*AccountPartialReconcile{apr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPartialReconcile creates a new account.partial.reconcile model and returns its id.
func (c *Client) CreateAccountPartialReconciles(aprs []*AccountPartialReconcile) ([]int64, error) {
	var vv []interface{}
	for _, v := range aprs {
		vv = append(vv, v)
	}
	return c.Create(AccountPartialReconcileModel, vv, nil)
}

// UpdateAccountPartialReconcile updates an existing account.partial.reconcile record.
func (c *Client) UpdateAccountPartialReconcile(apr *AccountPartialReconcile) error {
	return c.UpdateAccountPartialReconciles([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPartialReconciles updates existing account.partial.reconcile records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPartialReconciles(ids []int64, apr *AccountPartialReconcile) error {
	return c.Update(AccountPartialReconcileModel, ids, apr, nil)
}

// DeleteAccountPartialReconcile deletes an existing account.partial.reconcile record.
func (c *Client) DeleteAccountPartialReconcile(id int64) error {
	return c.DeleteAccountPartialReconciles([]int64{id})
}

// DeleteAccountPartialReconciles deletes existing account.partial.reconcile records.
func (c *Client) DeleteAccountPartialReconciles(ids []int64) error {
	return c.Delete(AccountPartialReconcileModel, ids)
}

// GetAccountPartialReconcile gets account.partial.reconcile existing record.
func (c *Client) GetAccountPartialReconcile(id int64) (*AccountPartialReconcile, error) {
	aprs, err := c.GetAccountPartialReconciles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// GetAccountPartialReconciles gets account.partial.reconcile existing records.
func (c *Client) GetAccountPartialReconciles(ids []int64) (*AccountPartialReconciles, error) {
	aprs := &AccountPartialReconciles{}
	if err := c.Read(AccountPartialReconcileModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPartialReconcile finds account.partial.reconcile record by querying it with criteria.
func (c *Client) FindAccountPartialReconcile(criteria *Criteria) (*AccountPartialReconcile, error) {
	aprs := &AccountPartialReconciles{}
	if err := c.SearchRead(AccountPartialReconcileModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// FindAccountPartialReconciles finds account.partial.reconcile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartialReconciles(criteria *Criteria, options *Options) (*AccountPartialReconciles, error) {
	aprs := &AccountPartialReconciles{}
	if err := c.SearchRead(AccountPartialReconcileModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPartialReconcileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartialReconcileIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPartialReconcileModel, criteria, options)
}

// FindAccountPartialReconcileId finds record id by querying it with criteria.
func (c *Client) FindAccountPartialReconcileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPartialReconcileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
