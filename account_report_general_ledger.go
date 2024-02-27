package odoo

// AccountReportGeneralLedger represents account.report.general.ledger model.
type AccountReportGeneralLedger struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omitempty"`
	DateTo         *Time      `xmlrpc:"date_to,omitempty"`
	DisplayAccount *Selection `xmlrpc:"display_account,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	InitialBalance *Bool      `xmlrpc:"initial_balance,omitempty"`
	JournalIds     *Relation  `xmlrpc:"journal_ids,omitempty"`
	Sortby         *Selection `xmlrpc:"sortby,omitempty"`
	TargetMove     *Selection `xmlrpc:"target_move,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReportGeneralLedgers represents array of account.report.general.ledger model.
type AccountReportGeneralLedgers []AccountReportGeneralLedger

// AccountReportGeneralLedgerModel is the odoo model name.
const AccountReportGeneralLedgerModel = "account.report.general.ledger"

// Many2One convert AccountReportGeneralLedger to *Many2One.
func (argl *AccountReportGeneralLedger) Many2One() *Many2One {
	return NewMany2One(argl.Id.Get(), "")
}

// CreateAccountReportGeneralLedger creates a new account.report.general.ledger model and returns its id.
func (c *Client) CreateAccountReportGeneralLedger(argl *AccountReportGeneralLedger) (int64, error) {
	ids, err := c.CreateAccountReportGeneralLedgers([]*AccountReportGeneralLedger{argl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportGeneralLedgers creates a new account.report.general.ledger model and returns its id.
func (c *Client) CreateAccountReportGeneralLedgers(argls []*AccountReportGeneralLedger) ([]int64, error) {
	var vv []interface{}
	for _, v := range argls {
		vv = append(vv, v)
	}
	return c.Create(AccountReportGeneralLedgerModel, vv, nil)
}

// UpdateAccountReportGeneralLedger updates an existing account.report.general.ledger record.
func (c *Client) UpdateAccountReportGeneralLedger(argl *AccountReportGeneralLedger) error {
	return c.UpdateAccountReportGeneralLedgers([]int64{argl.Id.Get()}, argl)
}

// UpdateAccountReportGeneralLedgers updates existing account.report.general.ledger records.
// All records (represented by ids) will be updated by argl values.
func (c *Client) UpdateAccountReportGeneralLedgers(ids []int64, argl *AccountReportGeneralLedger) error {
	return c.Update(AccountReportGeneralLedgerModel, ids, argl, nil)
}

// DeleteAccountReportGeneralLedger deletes an existing account.report.general.ledger record.
func (c *Client) DeleteAccountReportGeneralLedger(id int64) error {
	return c.DeleteAccountReportGeneralLedgers([]int64{id})
}

// DeleteAccountReportGeneralLedgers deletes existing account.report.general.ledger records.
func (c *Client) DeleteAccountReportGeneralLedgers(ids []int64) error {
	return c.Delete(AccountReportGeneralLedgerModel, ids)
}

// GetAccountReportGeneralLedger gets account.report.general.ledger existing record.
func (c *Client) GetAccountReportGeneralLedger(id int64) (*AccountReportGeneralLedger, error) {
	argls, err := c.GetAccountReportGeneralLedgers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*argls)[0]), nil
}

// GetAccountReportGeneralLedgers gets account.report.general.ledger existing records.
func (c *Client) GetAccountReportGeneralLedgers(ids []int64) (*AccountReportGeneralLedgers, error) {
	argls := &AccountReportGeneralLedgers{}
	if err := c.Read(AccountReportGeneralLedgerModel, ids, nil, argls); err != nil {
		return nil, err
	}
	return argls, nil
}

// FindAccountReportGeneralLedger finds account.report.general.ledger record by querying it with criteria.
func (c *Client) FindAccountReportGeneralLedger(criteria *Criteria) (*AccountReportGeneralLedger, error) {
	argls := &AccountReportGeneralLedgers{}
	if err := c.SearchRead(AccountReportGeneralLedgerModel, criteria, NewOptions().Limit(1), argls); err != nil {
		return nil, err
	}
	return &((*argls)[0]), nil
}

// FindAccountReportGeneralLedgers finds account.report.general.ledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportGeneralLedgers(criteria *Criteria, options *Options) (*AccountReportGeneralLedgers, error) {
	argls := &AccountReportGeneralLedgers{}
	if err := c.SearchRead(AccountReportGeneralLedgerModel, criteria, options, argls); err != nil {
		return nil, err
	}
	return argls, nil
}

// FindAccountReportGeneralLedgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportGeneralLedgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportGeneralLedgerModel, criteria, options)
}

// FindAccountReportGeneralLedgerId finds record id by querying it with criteria.
func (c *Client) FindAccountReportGeneralLedgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportGeneralLedgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
