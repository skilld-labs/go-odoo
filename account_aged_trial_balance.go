package odoo

// AccountAgedTrialBalance represents account.aged.trial.balance model.
type AccountAgedTrialBalance struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omitempty"`
	DateTo          *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	JournalIds      *Relation  `xmlrpc:"journal_ids,omitempty"`
	PeriodLength    *Int       `xmlrpc:"period_length,omitempty"`
	ResultSelection *Selection `xmlrpc:"result_selection,omitempty"`
	TargetMove      *Selection `xmlrpc:"target_move,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAgedTrialBalances represents array of account.aged.trial.balance model.
type AccountAgedTrialBalances []AccountAgedTrialBalance

// AccountAgedTrialBalanceModel is the odoo model name.
const AccountAgedTrialBalanceModel = "account.aged.trial.balance"

// Many2One convert AccountAgedTrialBalance to *Many2One.
func (aatb *AccountAgedTrialBalance) Many2One() *Many2One {
	return NewMany2One(aatb.Id.Get(), "")
}

// CreateAccountAgedTrialBalance creates a new account.aged.trial.balance model and returns its id.
func (c *Client) CreateAccountAgedTrialBalance(aatb *AccountAgedTrialBalance) (int64, error) {
	ids, err := c.CreateAccountAgedTrialBalances([]*AccountAgedTrialBalance{aatb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAgedTrialBalances creates a new account.aged.trial.balance model and returns its id.
func (c *Client) CreateAccountAgedTrialBalances(aatbs []*AccountAgedTrialBalance) ([]int64, error) {
	var vv []interface{}
	for _, v := range aatbs {
		vv = append(vv, v)
	}
	return c.Create(AccountAgedTrialBalanceModel, vv, nil)
}

// UpdateAccountAgedTrialBalance updates an existing account.aged.trial.balance record.
func (c *Client) UpdateAccountAgedTrialBalance(aatb *AccountAgedTrialBalance) error {
	return c.UpdateAccountAgedTrialBalances([]int64{aatb.Id.Get()}, aatb)
}

// UpdateAccountAgedTrialBalances updates existing account.aged.trial.balance records.
// All records (represented by ids) will be updated by aatb values.
func (c *Client) UpdateAccountAgedTrialBalances(ids []int64, aatb *AccountAgedTrialBalance) error {
	return c.Update(AccountAgedTrialBalanceModel, ids, aatb, nil)
}

// DeleteAccountAgedTrialBalance deletes an existing account.aged.trial.balance record.
func (c *Client) DeleteAccountAgedTrialBalance(id int64) error {
	return c.DeleteAccountAgedTrialBalances([]int64{id})
}

// DeleteAccountAgedTrialBalances deletes existing account.aged.trial.balance records.
func (c *Client) DeleteAccountAgedTrialBalances(ids []int64) error {
	return c.Delete(AccountAgedTrialBalanceModel, ids)
}

// GetAccountAgedTrialBalance gets account.aged.trial.balance existing record.
func (c *Client) GetAccountAgedTrialBalance(id int64) (*AccountAgedTrialBalance, error) {
	aatbs, err := c.GetAccountAgedTrialBalances([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aatbs)[0]), nil
}

// GetAccountAgedTrialBalances gets account.aged.trial.balance existing records.
func (c *Client) GetAccountAgedTrialBalances(ids []int64) (*AccountAgedTrialBalances, error) {
	aatbs := &AccountAgedTrialBalances{}
	if err := c.Read(AccountAgedTrialBalanceModel, ids, nil, aatbs); err != nil {
		return nil, err
	}
	return aatbs, nil
}

// FindAccountAgedTrialBalance finds account.aged.trial.balance record by querying it with criteria.
func (c *Client) FindAccountAgedTrialBalance(criteria *Criteria) (*AccountAgedTrialBalance, error) {
	aatbs := &AccountAgedTrialBalances{}
	if err := c.SearchRead(AccountAgedTrialBalanceModel, criteria, NewOptions().Limit(1), aatbs); err != nil {
		return nil, err
	}
	return &((*aatbs)[0]), nil
}

// FindAccountAgedTrialBalances finds account.aged.trial.balance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedTrialBalances(criteria *Criteria, options *Options) (*AccountAgedTrialBalances, error) {
	aatbs := &AccountAgedTrialBalances{}
	if err := c.SearchRead(AccountAgedTrialBalanceModel, criteria, options, aatbs); err != nil {
		return nil, err
	}
	return aatbs, nil
}

// FindAccountAgedTrialBalanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedTrialBalanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAgedTrialBalanceModel, criteria, options)
}

// FindAccountAgedTrialBalanceId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedTrialBalanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedTrialBalanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
