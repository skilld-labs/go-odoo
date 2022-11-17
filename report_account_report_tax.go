package odoo

import (
	"fmt"
)

// ReportAccountReportTax represents report.account.report_tax model.
type ReportAccountReportTax struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportTaxs represents array of report.account.report_tax model.
type ReportAccountReportTaxs []ReportAccountReportTax

// ReportAccountReportTaxModel is the odoo model name.
const ReportAccountReportTaxModel = "report.account.report_tax"

// Many2One convert ReportAccountReportTax to *Many2One.
func (rar *ReportAccountReportTax) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportTax creates a new report.account.report_tax model and returns its id.
func (c *Client) CreateReportAccountReportTax(rar *ReportAccountReportTax) (int64, error) {
	return c.Create(ReportAccountReportTaxModel, rar)
}

// UpdateReportAccountReportTax updates an existing report.account.report_tax record.
func (c *Client) UpdateReportAccountReportTax(rar *ReportAccountReportTax) error {
	return c.UpdateReportAccountReportTaxs([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportTaxs updates existing report.account.report_tax records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportTaxs(ids []int64, rar *ReportAccountReportTax) error {
	return c.Update(ReportAccountReportTaxModel, ids, rar)
}

// DeleteReportAccountReportTax deletes an existing report.account.report_tax record.
func (c *Client) DeleteReportAccountReportTax(id int64) error {
	return c.DeleteReportAccountReportTaxs([]int64{id})
}

// DeleteReportAccountReportTaxs deletes existing report.account.report_tax records.
func (c *Client) DeleteReportAccountReportTaxs(ids []int64) error {
	return c.Delete(ReportAccountReportTaxModel, ids)
}

// GetReportAccountReportTax gets report.account.report_tax existing record.
func (c *Client) GetReportAccountReportTax(id int64) (*ReportAccountReportTax, error) {
	rars, err := c.GetReportAccountReportTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_tax not found", id)
}

// GetReportAccountReportTaxs gets report.account.report_tax existing records.
func (c *Client) GetReportAccountReportTaxs(ids []int64) (*ReportAccountReportTaxs, error) {
	rars := &ReportAccountReportTaxs{}
	if err := c.Read(ReportAccountReportTaxModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportTax finds report.account.report_tax record by querying it with criteria.
func (c *Client) FindReportAccountReportTax(criteria *Criteria) (*ReportAccountReportTax, error) {
	rars := &ReportAccountReportTaxs{}
	if err := c.SearchRead(ReportAccountReportTaxModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("no report.account.report_tax was found with criteria %v", criteria)
}

// FindReportAccountReportTaxs finds report.account.report_tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportTaxs(criteria *Criteria, options *Options) (*ReportAccountReportTaxs, error) {
	rars := &ReportAccountReportTaxs{}
	if err := c.SearchRead(ReportAccountReportTaxModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportTaxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportTaxId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no report.account.report_tax was found with criteria %v and options %v", criteria, options)
}
