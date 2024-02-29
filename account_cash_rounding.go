package odoo

// AccountCashRounding represents account.cash.rounding model.
type AccountCashRounding struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId      *Many2One  `xmlrpc:"account_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	Rounding       *Float     `xmlrpc:"rounding,omitempty"`
	RoundingMethod *Selection `xmlrpc:"rounding_method,omitempty"`
	Strategy       *Selection `xmlrpc:"strategy,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountCashRoundings represents array of account.cash.rounding model.
type AccountCashRoundings []AccountCashRounding

// AccountCashRoundingModel is the odoo model name.
const AccountCashRoundingModel = "account.cash.rounding"

// Many2One convert AccountCashRounding to *Many2One.
func (acr *AccountCashRounding) Many2One() *Many2One {
	return NewMany2One(acr.Id.Get(), "")
}

// CreateAccountCashRounding creates a new account.cash.rounding model and returns its id.
func (c *Client) CreateAccountCashRounding(acr *AccountCashRounding) (int64, error) {
	ids, err := c.CreateAccountCashRoundings([]*AccountCashRounding{acr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCashRoundings creates a new account.cash.rounding model and returns its id.
func (c *Client) CreateAccountCashRoundings(acrs []*AccountCashRounding) ([]int64, error) {
	var vv []interface{}
	for _, v := range acrs {
		vv = append(vv, v)
	}
	return c.Create(AccountCashRoundingModel, vv, nil)
}

// UpdateAccountCashRounding updates an existing account.cash.rounding record.
func (c *Client) UpdateAccountCashRounding(acr *AccountCashRounding) error {
	return c.UpdateAccountCashRoundings([]int64{acr.Id.Get()}, acr)
}

// UpdateAccountCashRoundings updates existing account.cash.rounding records.
// All records (represented by ids) will be updated by acr values.
func (c *Client) UpdateAccountCashRoundings(ids []int64, acr *AccountCashRounding) error {
	return c.Update(AccountCashRoundingModel, ids, acr, nil)
}

// DeleteAccountCashRounding deletes an existing account.cash.rounding record.
func (c *Client) DeleteAccountCashRounding(id int64) error {
	return c.DeleteAccountCashRoundings([]int64{id})
}

// DeleteAccountCashRoundings deletes existing account.cash.rounding records.
func (c *Client) DeleteAccountCashRoundings(ids []int64) error {
	return c.Delete(AccountCashRoundingModel, ids)
}

// GetAccountCashRounding gets account.cash.rounding existing record.
func (c *Client) GetAccountCashRounding(id int64) (*AccountCashRounding, error) {
	acrs, err := c.GetAccountCashRoundings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acrs)[0]), nil
}

// GetAccountCashRoundings gets account.cash.rounding existing records.
func (c *Client) GetAccountCashRoundings(ids []int64) (*AccountCashRoundings, error) {
	acrs := &AccountCashRoundings{}
	if err := c.Read(AccountCashRoundingModel, ids, nil, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCashRounding finds account.cash.rounding record by querying it with criteria.
func (c *Client) FindAccountCashRounding(criteria *Criteria) (*AccountCashRounding, error) {
	acrs := &AccountCashRoundings{}
	if err := c.SearchRead(AccountCashRoundingModel, criteria, NewOptions().Limit(1), acrs); err != nil {
		return nil, err
	}
	return &((*acrs)[0]), nil
}

// FindAccountCashRoundings finds account.cash.rounding records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashRoundings(criteria *Criteria, options *Options) (*AccountCashRoundings, error) {
	acrs := &AccountCashRoundings{}
	if err := c.SearchRead(AccountCashRoundingModel, criteria, options, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCashRoundingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashRoundingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCashRoundingModel, criteria, options)
}

// FindAccountCashRoundingId finds record id by querying it with criteria.
func (c *Client) FindAccountCashRoundingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCashRoundingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
