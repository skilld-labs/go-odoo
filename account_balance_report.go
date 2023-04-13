package odoo

import (
	"fmt"
)

// AccountBalanceReport represents account.balance.report model.
type AccountBalanceReport struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omptempty"`
	DateTo         *Time      `xmlrpc:"date_to,omptempty"`
	DisplayAccount *Selection `xmlrpc:"display_account,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	JournalIds     *Relation  `xmlrpc:"journal_ids,omptempty"`
	TargetMove     *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBalanceReports represents array of account.balance.report model.
type AccountBalanceReports []AccountBalanceReport

// AccountBalanceReportModel is the odoo model name.
const AccountBalanceReportModel = "account.balance.report"

// Many2One convert AccountBalanceReport to *Many2One.
func (abr *AccountBalanceReport) Many2One() *Many2One {
	return NewMany2One(abr.Id.Get(), "")
}

// CreateAccountBalanceReport creates a new account.balance.report model and returns its id.
func (c *Client) CreateAccountBalanceReport(abr *AccountBalanceReport) (int64, error) {
	ids, err := c.CreateAccountBalanceReports([]*AccountBalanceReport{abr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBalanceReport creates a new account.balance.report model and returns its id.
func (c *Client) CreateAccountBalanceReports(abrs []*AccountBalanceReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range abrs {
		vv = append(vv, v)
	}
	return c.Create(AccountBalanceReportModel, vv)
}

// UpdateAccountBalanceReport updates an existing account.balance.report record.
func (c *Client) UpdateAccountBalanceReport(abr *AccountBalanceReport) error {
	return c.UpdateAccountBalanceReports([]int64{abr.Id.Get()}, abr)
}

// UpdateAccountBalanceReports updates existing account.balance.report records.
// All records (represented by ids) will be updated by abr values.
func (c *Client) UpdateAccountBalanceReports(ids []int64, abr *AccountBalanceReport) error {
	return c.Update(AccountBalanceReportModel, ids, abr)
}

// DeleteAccountBalanceReport deletes an existing account.balance.report record.
func (c *Client) DeleteAccountBalanceReport(id int64) error {
	return c.DeleteAccountBalanceReports([]int64{id})
}

// DeleteAccountBalanceReports deletes existing account.balance.report records.
func (c *Client) DeleteAccountBalanceReports(ids []int64) error {
	return c.Delete(AccountBalanceReportModel, ids)
}

// GetAccountBalanceReport gets account.balance.report existing record.
func (c *Client) GetAccountBalanceReport(id int64) (*AccountBalanceReport, error) {
	abrs, err := c.GetAccountBalanceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if abrs != nil && len(*abrs) > 0 {
		return &((*abrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.balance.report not found", id)
}

// GetAccountBalanceReports gets account.balance.report existing records.
func (c *Client) GetAccountBalanceReports(ids []int64) (*AccountBalanceReports, error) {
	abrs := &AccountBalanceReports{}
	if err := c.Read(AccountBalanceReportModel, ids, nil, abrs); err != nil {
		return nil, err
	}
	return abrs, nil
}

// FindAccountBalanceReport finds account.balance.report record by querying it with criteria.
func (c *Client) FindAccountBalanceReport(criteria *Criteria) (*AccountBalanceReport, error) {
	abrs := &AccountBalanceReports{}
	if err := c.SearchRead(AccountBalanceReportModel, criteria, NewOptions().Limit(1), abrs); err != nil {
		return nil, err
	}
	if abrs != nil && len(*abrs) > 0 {
		return &((*abrs)[0]), nil
	}
	return nil, fmt.Errorf("account.balance.report was not found with criteria %v", criteria)
}

// FindAccountBalanceReports finds account.balance.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBalanceReports(criteria *Criteria, options *Options) (*AccountBalanceReports, error) {
	abrs := &AccountBalanceReports{}
	if err := c.SearchRead(AccountBalanceReportModel, criteria, options, abrs); err != nil {
		return nil, err
	}
	return abrs, nil
}

// FindAccountBalanceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBalanceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBalanceReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBalanceReportId finds record id by querying it with criteria.
func (c *Client) FindAccountBalanceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBalanceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.balance.report was not found with criteria %v and options %v", criteria, options)
}
