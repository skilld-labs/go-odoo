package odoo

// AccountReportPartnerLedger represents account.report.partner.ledger model.
type AccountReportPartnerLedger struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	AmountCurrency  *Bool      `xmlrpc:"amount_currency,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omptempty"`
	DateTo          *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	JournalIds      *Relation  `xmlrpc:"journal_ids,omptempty"`
	Reconciled      *Bool      `xmlrpc:"reconciled,omptempty"`
	ResultSelection *Selection `xmlrpc:"result_selection,omptempty"`
	TargetMove      *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReportPartnerLedgers represents array of account.report.partner.ledger model.
type AccountReportPartnerLedgers []AccountReportPartnerLedger

// AccountReportPartnerLedgerModel is the odoo model name.
const AccountReportPartnerLedgerModel = "account.report.partner.ledger"

// Many2One convert AccountReportPartnerLedger to *Many2One.
func (arpl *AccountReportPartnerLedger) Many2One() *Many2One {
	return NewMany2One(arpl.Id.Get(), "")
}

// CreateAccountReportPartnerLedger creates a new account.report.partner.ledger model and returns its id.
func (c *Client) CreateAccountReportPartnerLedger(arpl *AccountReportPartnerLedger) (int64, error) {
	ids, err := c.CreateAccountReportPartnerLedgers([]*AccountReportPartnerLedger{arpl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportPartnerLedger creates a new account.report.partner.ledger model and returns its id.
func (c *Client) CreateAccountReportPartnerLedgers(arpls []*AccountReportPartnerLedger) ([]int64, error) {
	var vv []interface{}
	for _, v := range arpls {
		vv = append(vv, v)
	}
	return c.Create(AccountReportPartnerLedgerModel, vv, nil)
}

// UpdateAccountReportPartnerLedger updates an existing account.report.partner.ledger record.
func (c *Client) UpdateAccountReportPartnerLedger(arpl *AccountReportPartnerLedger) error {
	return c.UpdateAccountReportPartnerLedgers([]int64{arpl.Id.Get()}, arpl)
}

// UpdateAccountReportPartnerLedgers updates existing account.report.partner.ledger records.
// All records (represented by ids) will be updated by arpl values.
func (c *Client) UpdateAccountReportPartnerLedgers(ids []int64, arpl *AccountReportPartnerLedger) error {
	return c.Update(AccountReportPartnerLedgerModel, ids, arpl, nil)
}

// DeleteAccountReportPartnerLedger deletes an existing account.report.partner.ledger record.
func (c *Client) DeleteAccountReportPartnerLedger(id int64) error {
	return c.DeleteAccountReportPartnerLedgers([]int64{id})
}

// DeleteAccountReportPartnerLedgers deletes existing account.report.partner.ledger records.
func (c *Client) DeleteAccountReportPartnerLedgers(ids []int64) error {
	return c.Delete(AccountReportPartnerLedgerModel, ids)
}

// GetAccountReportPartnerLedger gets account.report.partner.ledger existing record.
func (c *Client) GetAccountReportPartnerLedger(id int64) (*AccountReportPartnerLedger, error) {
	arpls, err := c.GetAccountReportPartnerLedgers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arpls)[0]), nil
}

// GetAccountReportPartnerLedgers gets account.report.partner.ledger existing records.
func (c *Client) GetAccountReportPartnerLedgers(ids []int64) (*AccountReportPartnerLedgers, error) {
	arpls := &AccountReportPartnerLedgers{}
	if err := c.Read(AccountReportPartnerLedgerModel, ids, nil, arpls); err != nil {
		return nil, err
	}
	return arpls, nil
}

// FindAccountReportPartnerLedger finds account.report.partner.ledger record by querying it with criteria.
func (c *Client) FindAccountReportPartnerLedger(criteria *Criteria) (*AccountReportPartnerLedger, error) {
	arpls := &AccountReportPartnerLedgers{}
	if err := c.SearchRead(AccountReportPartnerLedgerModel, criteria, NewOptions().Limit(1), arpls); err != nil {
		return nil, err
	}
	return &((*arpls)[0]), nil
}

// FindAccountReportPartnerLedgers finds account.report.partner.ledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportPartnerLedgers(criteria *Criteria, options *Options) (*AccountReportPartnerLedgers, error) {
	arpls := &AccountReportPartnerLedgers{}
	if err := c.SearchRead(AccountReportPartnerLedgerModel, criteria, options, arpls); err != nil {
		return nil, err
	}
	return arpls, nil
}

// FindAccountReportPartnerLedgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportPartnerLedgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportPartnerLedgerModel, criteria, options)
}

// FindAccountReportPartnerLedgerId finds record id by querying it with criteria.
func (c *Client) FindAccountReportPartnerLedgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportPartnerLedgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
