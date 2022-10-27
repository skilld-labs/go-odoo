package odoo

import (
	"fmt"
)

// AccountTaxReport represents account.tax.report model.
type AccountTaxReport struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LineIds     *Relation `xmlrpc:"line_ids,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	RootLineIds *Relation `xmlrpc:"root_line_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxReports represents array of account.tax.report model.
type AccountTaxReports []AccountTaxReport

// AccountTaxReportModel is the odoo model name.
const AccountTaxReportModel = "account.tax.report"

// Many2One convert AccountTaxReport to *Many2One.
func (atr *AccountTaxReport) Many2One() *Many2One {
	return NewMany2One(atr.Id.Get(), "")
}

// CreateAccountTaxReport creates a new account.tax.report model and returns its id.
func (c *Client) CreateAccountTaxReport(atr *AccountTaxReport) (int64, error) {
	return c.Create(AccountTaxReportModel, atr)
}

// UpdateAccountTaxReport updates an existing account.tax.report record.
func (c *Client) UpdateAccountTaxReport(atr *AccountTaxReport) error {
	return c.UpdateAccountTaxReports([]int64{atr.Id.Get()}, atr)
}

// UpdateAccountTaxReports updates existing account.tax.report records.
// All records (represented by ids) will be updated by atr values.
func (c *Client) UpdateAccountTaxReports(ids []int64, atr *AccountTaxReport) error {
	return c.Update(AccountTaxReportModel, ids, atr)
}

// DeleteAccountTaxReport deletes an existing account.tax.report record.
func (c *Client) DeleteAccountTaxReport(id int64) error {
	return c.DeleteAccountTaxReports([]int64{id})
}

// DeleteAccountTaxReports deletes existing account.tax.report records.
func (c *Client) DeleteAccountTaxReports(ids []int64) error {
	return c.Delete(AccountTaxReportModel, ids)
}

// GetAccountTaxReport gets account.tax.report existing record.
func (c *Client) GetAccountTaxReport(id int64) (*AccountTaxReport, error) {
	atrs, err := c.GetAccountTaxReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if atrs != nil && len(*atrs) > 0 {
		return &((*atrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.report not found", id)
}

// GetAccountTaxReports gets account.tax.report existing records.
func (c *Client) GetAccountTaxReports(ids []int64) (*AccountTaxReports, error) {
	atrs := &AccountTaxReports{}
	if err := c.Read(AccountTaxReportModel, ids, nil, atrs); err != nil {
		return nil, err
	}
	return atrs, nil
}

// FindAccountTaxReport finds account.tax.report record by querying it with criteria.
func (c *Client) FindAccountTaxReport(criteria *Criteria) (*AccountTaxReport, error) {
	atrs := &AccountTaxReports{}
	if err := c.SearchRead(AccountTaxReportModel, criteria, NewOptions().Limit(1), atrs); err != nil {
		return nil, err
	}
	if atrs != nil && len(*atrs) > 0 {
		return &((*atrs)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.report was not found")
}

// FindAccountTaxReports finds account.tax.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReports(criteria *Criteria, options *Options) (*AccountTaxReports, error) {
	atrs := &AccountTaxReports{}
	if err := c.SearchRead(AccountTaxReportModel, criteria, options, atrs); err != nil {
		return nil, err
	}
	return atrs, nil
}

// FindAccountTaxReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxReportId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.report was not found")
}
