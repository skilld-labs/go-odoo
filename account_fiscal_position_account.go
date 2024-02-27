package odoo

// AccountFiscalPositionAccount represents account.fiscal.position.account model.
type AccountFiscalPositionAccount struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	AccountDestId *Many2One `xmlrpc:"account_dest_id,omitempty"`
	AccountSrcId  *Many2One `xmlrpc:"account_src_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	PositionId    *Many2One `xmlrpc:"position_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFiscalPositionAccounts represents array of account.fiscal.position.account model.
type AccountFiscalPositionAccounts []AccountFiscalPositionAccount

// AccountFiscalPositionAccountModel is the odoo model name.
const AccountFiscalPositionAccountModel = "account.fiscal.position.account"

// Many2One convert AccountFiscalPositionAccount to *Many2One.
func (afpa *AccountFiscalPositionAccount) Many2One() *Many2One {
	return NewMany2One(afpa.Id.Get(), "")
}

// CreateAccountFiscalPositionAccount creates a new account.fiscal.position.account model and returns its id.
func (c *Client) CreateAccountFiscalPositionAccount(afpa *AccountFiscalPositionAccount) (int64, error) {
	ids, err := c.CreateAccountFiscalPositionAccounts([]*AccountFiscalPositionAccount{afpa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFiscalPositionAccount creates a new account.fiscal.position.account model and returns its id.
func (c *Client) CreateAccountFiscalPositionAccounts(afpas []*AccountFiscalPositionAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range afpas {
		vv = append(vv, v)
	}
	return c.Create(AccountFiscalPositionAccountModel, vv, nil)
}

// UpdateAccountFiscalPositionAccount updates an existing account.fiscal.position.account record.
func (c *Client) UpdateAccountFiscalPositionAccount(afpa *AccountFiscalPositionAccount) error {
	return c.UpdateAccountFiscalPositionAccounts([]int64{afpa.Id.Get()}, afpa)
}

// UpdateAccountFiscalPositionAccounts updates existing account.fiscal.position.account records.
// All records (represented by ids) will be updated by afpa values.
func (c *Client) UpdateAccountFiscalPositionAccounts(ids []int64, afpa *AccountFiscalPositionAccount) error {
	return c.Update(AccountFiscalPositionAccountModel, ids, afpa, nil)
}

// DeleteAccountFiscalPositionAccount deletes an existing account.fiscal.position.account record.
func (c *Client) DeleteAccountFiscalPositionAccount(id int64) error {
	return c.DeleteAccountFiscalPositionAccounts([]int64{id})
}

// DeleteAccountFiscalPositionAccounts deletes existing account.fiscal.position.account records.
func (c *Client) DeleteAccountFiscalPositionAccounts(ids []int64) error {
	return c.Delete(AccountFiscalPositionAccountModel, ids)
}

// GetAccountFiscalPositionAccount gets account.fiscal.position.account existing record.
func (c *Client) GetAccountFiscalPositionAccount(id int64) (*AccountFiscalPositionAccount, error) {
	afpas, err := c.GetAccountFiscalPositionAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afpas)[0]), nil
}

// GetAccountFiscalPositionAccounts gets account.fiscal.position.account existing records.
func (c *Client) GetAccountFiscalPositionAccounts(ids []int64) (*AccountFiscalPositionAccounts, error) {
	afpas := &AccountFiscalPositionAccounts{}
	if err := c.Read(AccountFiscalPositionAccountModel, ids, nil, afpas); err != nil {
		return nil, err
	}
	return afpas, nil
}

// FindAccountFiscalPositionAccount finds account.fiscal.position.account record by querying it with criteria.
func (c *Client) FindAccountFiscalPositionAccount(criteria *Criteria) (*AccountFiscalPositionAccount, error) {
	afpas := &AccountFiscalPositionAccounts{}
	if err := c.SearchRead(AccountFiscalPositionAccountModel, criteria, NewOptions().Limit(1), afpas); err != nil {
		return nil, err
	}
	return &((*afpas)[0]), nil
}

// FindAccountFiscalPositionAccounts finds account.fiscal.position.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionAccounts(criteria *Criteria, options *Options) (*AccountFiscalPositionAccounts, error) {
	afpas := &AccountFiscalPositionAccounts{}
	if err := c.SearchRead(AccountFiscalPositionAccountModel, criteria, options, afpas); err != nil {
		return nil, err
	}
	return afpas, nil
}

// FindAccountFiscalPositionAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFiscalPositionAccountModel, criteria, options)
}

// FindAccountFiscalPositionAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
