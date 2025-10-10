package odoo

// AccountBankStatement represents account.bank.statement model.
type AccountBankStatement struct {
	AttachmentIds      *Relation `xmlrpc:"attachment_ids,omitempty"`
	BalanceEnd         *Float    `xmlrpc:"balance_end,omitempty"`
	BalanceEndReal     *Float    `xmlrpc:"balance_end_real,omitempty"`
	BalanceStart       *Float    `xmlrpc:"balance_start,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId         *Many2One `xmlrpc:"currency_id,omitempty"`
	Date               *Time     `xmlrpc:"date,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	FirstLineIndex     *String   `xmlrpc:"first_line_index,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	IsComplete         *Bool     `xmlrpc:"is_complete,omitempty"`
	IsValid            *Bool     `xmlrpc:"is_valid,omitempty"`
	JournalId          *Many2One `xmlrpc:"journal_id,omitempty"`
	LineIds            *Relation `xmlrpc:"line_ids,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	ProblemDescription *String   `xmlrpc:"problem_description,omitempty"`
	Reference          *String   `xmlrpc:"reference,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatements represents array of account.bank.statement model.
type AccountBankStatements []AccountBankStatement

// AccountBankStatementModel is the odoo model name.
const AccountBankStatementModel = "account.bank.statement"

// Many2One convert AccountBankStatement to *Many2One.
func (abs *AccountBankStatement) Many2One() *Many2One {
	return NewMany2One(abs.Id.Get(), "")
}

// CreateAccountBankStatement creates a new account.bank.statement model and returns its id.
func (c *Client) CreateAccountBankStatement(abs *AccountBankStatement) (int64, error) {
	ids, err := c.CreateAccountBankStatements([]*AccountBankStatement{abs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatement creates a new account.bank.statement model and returns its id.
func (c *Client) CreateAccountBankStatements(abss []*AccountBankStatement) ([]int64, error) {
	var vv []interface{}
	for _, v := range abss {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementModel, vv, nil)
}

// UpdateAccountBankStatement updates an existing account.bank.statement record.
func (c *Client) UpdateAccountBankStatement(abs *AccountBankStatement) error {
	return c.UpdateAccountBankStatements([]int64{abs.Id.Get()}, abs)
}

// UpdateAccountBankStatements updates existing account.bank.statement records.
// All records (represented by ids) will be updated by abs values.
func (c *Client) UpdateAccountBankStatements(ids []int64, abs *AccountBankStatement) error {
	return c.Update(AccountBankStatementModel, ids, abs, nil)
}

// DeleteAccountBankStatement deletes an existing account.bank.statement record.
func (c *Client) DeleteAccountBankStatement(id int64) error {
	return c.DeleteAccountBankStatements([]int64{id})
}

// DeleteAccountBankStatements deletes existing account.bank.statement records.
func (c *Client) DeleteAccountBankStatements(ids []int64) error {
	return c.Delete(AccountBankStatementModel, ids)
}

// GetAccountBankStatement gets account.bank.statement existing record.
func (c *Client) GetAccountBankStatement(id int64) (*AccountBankStatement, error) {
	abss, err := c.GetAccountBankStatements([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abss)[0]), nil
}

// GetAccountBankStatements gets account.bank.statement existing records.
func (c *Client) GetAccountBankStatements(ids []int64) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.Read(AccountBankStatementModel, ids, nil, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatement finds account.bank.statement record by querying it with criteria.
func (c *Client) FindAccountBankStatement(criteria *Criteria) (*AccountBankStatement, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, NewOptions().Limit(1), abss); err != nil {
		return nil, err
	}
	return &((*abss)[0]), nil
}

// FindAccountBankStatements finds account.bank.statement records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatements(criteria *Criteria, options *Options) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, options, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatementIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankStatementModel, criteria, options)
}

// FindAccountBankStatementId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
