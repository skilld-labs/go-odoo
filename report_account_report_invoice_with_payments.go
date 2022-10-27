package odoo

import (
	"fmt"
)

// ReportAccountReportInvoiceWithPayments represents report.account.report_invoice_with_payments model.
type ReportAccountReportInvoiceWithPayments struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportInvoiceWithPaymentss represents array of report.account.report_invoice_with_payments model.
type ReportAccountReportInvoiceWithPaymentss []ReportAccountReportInvoiceWithPayments

// ReportAccountReportInvoiceWithPaymentsModel is the odoo model name.
const ReportAccountReportInvoiceWithPaymentsModel = "report.account.report_invoice_with_payments"

// Many2One convert ReportAccountReportInvoiceWithPayments to *Many2One.
func (rar *ReportAccountReportInvoiceWithPayments) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportInvoiceWithPayments creates a new report.account.report_invoice_with_payments model and returns its id.
func (c *Client) CreateReportAccountReportInvoiceWithPayments(rar *ReportAccountReportInvoiceWithPayments) (int64, error) {
	return c.Create(ReportAccountReportInvoiceWithPaymentsModel, rar)
}

// UpdateReportAccountReportInvoiceWithPayments updates an existing report.account.report_invoice_with_payments record.
func (c *Client) UpdateReportAccountReportInvoiceWithPayments(rar *ReportAccountReportInvoiceWithPayments) error {
	return c.UpdateReportAccountReportInvoiceWithPaymentss([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportInvoiceWithPaymentss updates existing report.account.report_invoice_with_payments records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportInvoiceWithPaymentss(ids []int64, rar *ReportAccountReportInvoiceWithPayments) error {
	return c.Update(ReportAccountReportInvoiceWithPaymentsModel, ids, rar)
}

// DeleteReportAccountReportInvoiceWithPayments deletes an existing report.account.report_invoice_with_payments record.
func (c *Client) DeleteReportAccountReportInvoiceWithPayments(id int64) error {
	return c.DeleteReportAccountReportInvoiceWithPaymentss([]int64{id})
}

// DeleteReportAccountReportInvoiceWithPaymentss deletes existing report.account.report_invoice_with_payments records.
func (c *Client) DeleteReportAccountReportInvoiceWithPaymentss(ids []int64) error {
	return c.Delete(ReportAccountReportInvoiceWithPaymentsModel, ids)
}

// GetReportAccountReportInvoiceWithPayments gets report.account.report_invoice_with_payments existing record.
func (c *Client) GetReportAccountReportInvoiceWithPayments(id int64) (*ReportAccountReportInvoiceWithPayments, error) {
	rars, err := c.GetReportAccountReportInvoiceWithPaymentss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_invoice_with_payments not found", id)
}

// GetReportAccountReportInvoiceWithPaymentss gets report.account.report_invoice_with_payments existing records.
func (c *Client) GetReportAccountReportInvoiceWithPaymentss(ids []int64) (*ReportAccountReportInvoiceWithPaymentss, error) {
	rars := &ReportAccountReportInvoiceWithPaymentss{}
	if err := c.Read(ReportAccountReportInvoiceWithPaymentsModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportInvoiceWithPayments finds report.account.report_invoice_with_payments record by querying it with criteria.
func (c *Client) FindReportAccountReportInvoiceWithPayments(criteria *Criteria) (*ReportAccountReportInvoiceWithPayments, error) {
	rars := &ReportAccountReportInvoiceWithPaymentss{}
	if err := c.SearchRead(ReportAccountReportInvoiceWithPaymentsModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("report.account.report_invoice_with_payments was not found")
}

// FindReportAccountReportInvoiceWithPaymentss finds report.account.report_invoice_with_payments records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportInvoiceWithPaymentss(criteria *Criteria, options *Options) (*ReportAccountReportInvoiceWithPaymentss, error) {
	rars := &ReportAccountReportInvoiceWithPaymentss{}
	if err := c.SearchRead(ReportAccountReportInvoiceWithPaymentsModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportInvoiceWithPaymentsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportInvoiceWithPaymentsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportInvoiceWithPaymentsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportInvoiceWithPaymentsId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportInvoiceWithPaymentsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportInvoiceWithPaymentsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.account.report_invoice_with_payments was not found")
}
