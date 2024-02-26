package odoo

// AccountBankAccountsWizard represents account.bank.accounts.wizard model.
type AccountBankAccountsWizard struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	AccName       *String    `xmlrpc:"acc_name,omptempty"`
	AccountType   *Selection `xmlrpc:"account_type,omptempty"`
	BankAccountId *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId    *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBankAccountsWizards represents array of account.bank.accounts.wizard model.
type AccountBankAccountsWizards []AccountBankAccountsWizard

// AccountBankAccountsWizardModel is the odoo model name.
const AccountBankAccountsWizardModel = "account.bank.accounts.wizard"

// Many2One convert AccountBankAccountsWizard to *Many2One.
func (abaw *AccountBankAccountsWizard) Many2One() *Many2One {
	return NewMany2One(abaw.Id.Get(), "")
}

// CreateAccountBankAccountsWizard creates a new account.bank.accounts.wizard model and returns its id.
func (c *Client) CreateAccountBankAccountsWizard(abaw *AccountBankAccountsWizard) (int64, error) {
	ids, err := c.CreateAccountBankAccountsWizards([]*AccountBankAccountsWizard{abaw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankAccountsWizard creates a new account.bank.accounts.wizard model and returns its id.
func (c *Client) CreateAccountBankAccountsWizards(abaws []*AccountBankAccountsWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range abaws {
		vv = append(vv, v)
	}
	return c.Create(AccountBankAccountsWizardModel, vv, nil)
}

// UpdateAccountBankAccountsWizard updates an existing account.bank.accounts.wizard record.
func (c *Client) UpdateAccountBankAccountsWizard(abaw *AccountBankAccountsWizard) error {
	return c.UpdateAccountBankAccountsWizards([]int64{abaw.Id.Get()}, abaw)
}

// UpdateAccountBankAccountsWizards updates existing account.bank.accounts.wizard records.
// All records (represented by ids) will be updated by abaw values.
func (c *Client) UpdateAccountBankAccountsWizards(ids []int64, abaw *AccountBankAccountsWizard) error {
	return c.Update(AccountBankAccountsWizardModel, ids, abaw, nil)
}

// DeleteAccountBankAccountsWizard deletes an existing account.bank.accounts.wizard record.
func (c *Client) DeleteAccountBankAccountsWizard(id int64) error {
	return c.DeleteAccountBankAccountsWizards([]int64{id})
}

// DeleteAccountBankAccountsWizards deletes existing account.bank.accounts.wizard records.
func (c *Client) DeleteAccountBankAccountsWizards(ids []int64) error {
	return c.Delete(AccountBankAccountsWizardModel, ids)
}

// GetAccountBankAccountsWizard gets account.bank.accounts.wizard existing record.
func (c *Client) GetAccountBankAccountsWizard(id int64) (*AccountBankAccountsWizard, error) {
	abaws, err := c.GetAccountBankAccountsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abaws)[0]), nil
}

// GetAccountBankAccountsWizards gets account.bank.accounts.wizard existing records.
func (c *Client) GetAccountBankAccountsWizards(ids []int64) (*AccountBankAccountsWizards, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.Read(AccountBankAccountsWizardModel, ids, nil, abaws); err != nil {
		return nil, err
	}
	return abaws, nil
}

// FindAccountBankAccountsWizard finds account.bank.accounts.wizard record by querying it with criteria.
func (c *Client) FindAccountBankAccountsWizard(criteria *Criteria) (*AccountBankAccountsWizard, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.SearchRead(AccountBankAccountsWizardModel, criteria, NewOptions().Limit(1), abaws); err != nil {
		return nil, err
	}
	return &((*abaws)[0]), nil
}

// FindAccountBankAccountsWizards finds account.bank.accounts.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankAccountsWizards(criteria *Criteria, options *Options) (*AccountBankAccountsWizards, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.SearchRead(AccountBankAccountsWizardModel, criteria, options, abaws); err != nil {
		return nil, err
	}
	return abaws, nil
}

// FindAccountBankAccountsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankAccountsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankAccountsWizardModel, criteria, options)
}

// FindAccountBankAccountsWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountBankAccountsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankAccountsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
