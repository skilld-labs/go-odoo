package odoo

import (
	"fmt"
)

// AccountCommonJournalReport represents account.common.journal.report model.
type AccountCommonJournalReport struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AmountCurrency *Bool      `xmlrpc:"amount_currency,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omptempty"`
	DateTo         *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	JournalIds     *Relation  `xmlrpc:"journal_ids,omptempty"`
	TargetMove     *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountCommonJournalReports represents array of account.common.journal.report model.
type AccountCommonJournalReports []AccountCommonJournalReport

// AccountCommonJournalReportModel is the odoo model name.
const AccountCommonJournalReportModel = "account.common.journal.report"

// Many2One convert AccountCommonJournalReport to *Many2One.
func (acjr *AccountCommonJournalReport) Many2One() *Many2One {
	return NewMany2One(acjr.Id.Get(), "")
}

// CreateAccountCommonJournalReport creates a new account.common.journal.report model and returns its id.
func (c *Client) CreateAccountCommonJournalReport(acjr *AccountCommonJournalReport) (int64, error) {
	ids, err := c.CreateAccountCommonJournalReports([]*AccountCommonJournalReport{acjr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCommonJournalReport creates a new account.common.journal.report model and returns its id.
func (c *Client) CreateAccountCommonJournalReports(acjrs []*AccountCommonJournalReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range acjrs {
		vv = append(vv, v)
	}
	return c.Create(AccountCommonJournalReportModel, vv)
}

// UpdateAccountCommonJournalReport updates an existing account.common.journal.report record.
func (c *Client) UpdateAccountCommonJournalReport(acjr *AccountCommonJournalReport) error {
	return c.UpdateAccountCommonJournalReports([]int64{acjr.Id.Get()}, acjr)
}

// UpdateAccountCommonJournalReports updates existing account.common.journal.report records.
// All records (represented by ids) will be updated by acjr values.
func (c *Client) UpdateAccountCommonJournalReports(ids []int64, acjr *AccountCommonJournalReport) error {
	return c.Update(AccountCommonJournalReportModel, ids, acjr)
}

// DeleteAccountCommonJournalReport deletes an existing account.common.journal.report record.
func (c *Client) DeleteAccountCommonJournalReport(id int64) error {
	return c.DeleteAccountCommonJournalReports([]int64{id})
}

// DeleteAccountCommonJournalReports deletes existing account.common.journal.report records.
func (c *Client) DeleteAccountCommonJournalReports(ids []int64) error {
	return c.Delete(AccountCommonJournalReportModel, ids)
}

// GetAccountCommonJournalReport gets account.common.journal.report existing record.
func (c *Client) GetAccountCommonJournalReport(id int64) (*AccountCommonJournalReport, error) {
	acjrs, err := c.GetAccountCommonJournalReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if acjrs != nil && len(*acjrs) > 0 {
		return &((*acjrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.common.journal.report not found", id)
}

// GetAccountCommonJournalReports gets account.common.journal.report existing records.
func (c *Client) GetAccountCommonJournalReports(ids []int64) (*AccountCommonJournalReports, error) {
	acjrs := &AccountCommonJournalReports{}
	if err := c.Read(AccountCommonJournalReportModel, ids, nil, acjrs); err != nil {
		return nil, err
	}
	return acjrs, nil
}

// FindAccountCommonJournalReport finds account.common.journal.report record by querying it with criteria.
func (c *Client) FindAccountCommonJournalReport(criteria *Criteria) (*AccountCommonJournalReport, error) {
	acjrs := &AccountCommonJournalReports{}
	if err := c.SearchRead(AccountCommonJournalReportModel, criteria, NewOptions().Limit(1), acjrs); err != nil {
		return nil, err
	}
	if acjrs != nil && len(*acjrs) > 0 {
		return &((*acjrs)[0]), nil
	}
	return nil, fmt.Errorf("account.common.journal.report was not found with criteria %v", criteria)
}

// FindAccountCommonJournalReports finds account.common.journal.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonJournalReports(criteria *Criteria, options *Options) (*AccountCommonJournalReports, error) {
	acjrs := &AccountCommonJournalReports{}
	if err := c.SearchRead(AccountCommonJournalReportModel, criteria, options, acjrs); err != nil {
		return nil, err
	}
	return acjrs, nil
}

// FindAccountCommonJournalReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonJournalReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountCommonJournalReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountCommonJournalReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCommonJournalReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCommonJournalReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.common.journal.report was not found with criteria %v and options %v", criteria, options)
}
