package odoo

import (
	"fmt"
)

// AccountTaxReportLine represents account.tax.report.line model.
type AccountTaxReportLine struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	ChildrenLineIds *Relation `xmlrpc:"children_line_ids,omitempty"`
	Code            *String   `xmlrpc:"code,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Formula         *String   `xmlrpc:"formula,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	ParentId        *Many2One `xmlrpc:"parent_id,omitempty"`
	ParentPath      *String   `xmlrpc:"parent_path,omitempty"`
	ReportActionId  *Many2One `xmlrpc:"report_action_id,omitempty"`
	ReportId        *Many2One `xmlrpc:"report_id,omitempty"`
	Sequence        *Int      `xmlrpc:"sequence,omitempty"`
	TagIds          *Relation `xmlrpc:"tag_ids,omitempty"`
	TagName         *String   `xmlrpc:"tag_name,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxReportLines represents array of account.tax.report.line model.
type AccountTaxReportLines []AccountTaxReportLine

// AccountTaxReportLineModel is the odoo model name.
const AccountTaxReportLineModel = "account.tax.report.line"

// Many2One convert AccountTaxReportLine to *Many2One.
func (atrl *AccountTaxReportLine) Many2One() *Many2One {
	return NewMany2One(atrl.Id.Get(), "")
}

// CreateAccountTaxReportLine creates a new account.tax.report.line model and returns its id.
func (c *Client) CreateAccountTaxReportLine(atrl *AccountTaxReportLine) (int64, error) {
	return c.Create(AccountTaxReportLineModel, atrl)
}

// UpdateAccountTaxReportLine updates an existing account.tax.report.line record.
func (c *Client) UpdateAccountTaxReportLine(atrl *AccountTaxReportLine) error {
	return c.UpdateAccountTaxReportLines([]int64{atrl.Id.Get()}, atrl)
}

// UpdateAccountTaxReportLines updates existing account.tax.report.line records.
// All records (represented by IDs) will be updated by atrl values.
func (c *Client) UpdateAccountTaxReportLines(ids []int64, atrl *AccountTaxReportLine) error {
	return c.Update(AccountTaxReportLineModel, ids, atrl)
}

// DeleteAccountTaxReportLine deletes an existing account.tax.report.line record.
func (c *Client) DeleteAccountTaxReportLine(id int64) error {
	return c.DeleteAccountTaxReportLines([]int64{id})
}

// DeleteAccountTaxReportLines deletes existing account.tax.report.line records.
func (c *Client) DeleteAccountTaxReportLines(ids []int64) error {
	return c.Delete(AccountTaxReportLineModel, ids)
}

// GetAccountTaxReportLine gets account.tax.report.line existing record.
func (c *Client) GetAccountTaxReportLine(id int64) (*AccountTaxReportLine, error) {
	atrls, err := c.GetAccountTaxReportLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if atrls != nil && len(*atrls) > 0 {
		return &((*atrls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.report.line not found", id)
}

// GetAccountTaxReportLines gets account.tax.report.line existing records.
func (c *Client) GetAccountTaxReportLines(ids []int64) (*AccountTaxReportLines, error) {
	atrls := &AccountTaxReportLines{}
	if err := c.Read(AccountTaxReportLineModel, ids, nil, atrls); err != nil {
		return nil, err
	}
	return atrls, nil
}

// FindAccountTaxReportLine finds account.tax.report.line record by querying it with criteria.
func (c *Client) FindAccountTaxReportLine(criteria *Criteria) (*AccountTaxReportLine, error) {
	atrls := &AccountTaxReportLines{}
	if err := c.SearchRead(AccountTaxReportLineModel, criteria, NewOptions().Limit(1), atrls); err != nil {
		return nil, err
	}
	if atrls != nil && len(*atrls) > 0 {
		return &((*atrls)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.report.line was not found")
}

// FindAccountTaxReportLines finds account.tax.report.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReportLines(criteria *Criteria, options *Options) (*AccountTaxReportLines, error) {
	atrls := &AccountTaxReportLines{}
	if err := c.SearchRead(AccountTaxReportLineModel, criteria, options, atrls); err != nil {
		return nil, err
	}
	return atrls, nil
}

// FindAccountTaxReportLineIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReportLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxReportLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxReportLineId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxReportLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxReportLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.report.line was not found")
}
