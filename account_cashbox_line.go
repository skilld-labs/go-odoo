package odoo

// AccountCashboxLine represents account.cashbox.line model.
type AccountCashboxLine struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CashboxId   *Many2One `xmlrpc:"cashbox_id,omitempty"`
	CoinValue   *Float    `xmlrpc:"coin_value,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Number      *Int      `xmlrpc:"number,omitempty"`
	Subtotal    *Float    `xmlrpc:"subtotal,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountCashboxLines represents array of account.cashbox.line model.
type AccountCashboxLines []AccountCashboxLine

// AccountCashboxLineModel is the odoo model name.
const AccountCashboxLineModel = "account.cashbox.line"

// Many2One convert AccountCashboxLine to *Many2One.
func (acl *AccountCashboxLine) Many2One() *Many2One {
	return NewMany2One(acl.Id.Get(), "")
}

// CreateAccountCashboxLine creates a new account.cashbox.line model and returns its id.
func (c *Client) CreateAccountCashboxLine(acl *AccountCashboxLine) (int64, error) {
	ids, err := c.CreateAccountCashboxLines([]*AccountCashboxLine{acl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCashboxLines creates a new account.cashbox.line model and returns its id.
func (c *Client) CreateAccountCashboxLines(acls []*AccountCashboxLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range acls {
		vv = append(vv, v)
	}
	return c.Create(AccountCashboxLineModel, vv, nil)
}

// UpdateAccountCashboxLine updates an existing account.cashbox.line record.
func (c *Client) UpdateAccountCashboxLine(acl *AccountCashboxLine) error {
	return c.UpdateAccountCashboxLines([]int64{acl.Id.Get()}, acl)
}

// UpdateAccountCashboxLines updates existing account.cashbox.line records.
// All records (represented by ids) will be updated by acl values.
func (c *Client) UpdateAccountCashboxLines(ids []int64, acl *AccountCashboxLine) error {
	return c.Update(AccountCashboxLineModel, ids, acl, nil)
}

// DeleteAccountCashboxLine deletes an existing account.cashbox.line record.
func (c *Client) DeleteAccountCashboxLine(id int64) error {
	return c.DeleteAccountCashboxLines([]int64{id})
}

// DeleteAccountCashboxLines deletes existing account.cashbox.line records.
func (c *Client) DeleteAccountCashboxLines(ids []int64) error {
	return c.Delete(AccountCashboxLineModel, ids)
}

// GetAccountCashboxLine gets account.cashbox.line existing record.
func (c *Client) GetAccountCashboxLine(id int64) (*AccountCashboxLine, error) {
	acls, err := c.GetAccountCashboxLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acls)[0]), nil
}

// GetAccountCashboxLines gets account.cashbox.line existing records.
func (c *Client) GetAccountCashboxLines(ids []int64) (*AccountCashboxLines, error) {
	acls := &AccountCashboxLines{}
	if err := c.Read(AccountCashboxLineModel, ids, nil, acls); err != nil {
		return nil, err
	}
	return acls, nil
}

// FindAccountCashboxLine finds account.cashbox.line record by querying it with criteria.
func (c *Client) FindAccountCashboxLine(criteria *Criteria) (*AccountCashboxLine, error) {
	acls := &AccountCashboxLines{}
	if err := c.SearchRead(AccountCashboxLineModel, criteria, NewOptions().Limit(1), acls); err != nil {
		return nil, err
	}
	return &((*acls)[0]), nil
}

// FindAccountCashboxLines finds account.cashbox.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashboxLines(criteria *Criteria, options *Options) (*AccountCashboxLines, error) {
	acls := &AccountCashboxLines{}
	if err := c.SearchRead(AccountCashboxLineModel, criteria, options, acls); err != nil {
		return nil, err
	}
	return acls, nil
}

// FindAccountCashboxLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashboxLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCashboxLineModel, criteria, options)
}

// FindAccountCashboxLineId finds record id by querying it with criteria.
func (c *Client) FindAccountCashboxLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCashboxLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
