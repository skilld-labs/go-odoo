package odoo

import (
	"fmt"
)

// AccountInvoiceReport represents account.invoice.report model.
type AccountInvoiceReport struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId        *Many2One  `xmlrpc:"account_analytic_id,omptempty"`
	AccountId                *Many2One  `xmlrpc:"account_id,omptempty"`
	AccountLineId            *Many2One  `xmlrpc:"account_line_id,omptempty"`
	CategId                  *Many2One  `xmlrpc:"categ_id,omptempty"`
	CommercialPartnerId      *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyRate             *Float     `xmlrpc:"currency_rate,omptempty"`
	Date                     *Time      `xmlrpc:"date,omptempty"`
	DateDue                  *Time      `xmlrpc:"date_due,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId         *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	JournalId                *Many2One  `xmlrpc:"journal_id,omptempty"`
	Nbr                      *Int       `xmlrpc:"nbr,omptempty"`
	PartnerBankId            *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentTermId            *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PriceAverage             *Float     `xmlrpc:"price_average,omptempty"`
	PriceTotal               *Float     `xmlrpc:"price_total,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductQty               *Float     `xmlrpc:"product_qty,omptempty"`
	Residual                 *Float     `xmlrpc:"residual,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	TeamId                   *Many2One  `xmlrpc:"team_id,omptempty"`
	Type                     *Selection `xmlrpc:"type,omptempty"`
	UomName                  *String    `xmlrpc:"uom_name,omptempty"`
	UserCurrencyPriceAverage *Float     `xmlrpc:"user_currency_price_average,omptempty"`
	UserCurrencyPriceTotal   *Float     `xmlrpc:"user_currency_price_total,omptempty"`
	UserCurrencyResidual     *Float     `xmlrpc:"user_currency_residual,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
}

// AccountInvoiceReports represents array of account.invoice.report model.
type AccountInvoiceReports []AccountInvoiceReport

// AccountInvoiceReportModel is the odoo model name.
const AccountInvoiceReportModel = "account.invoice.report"

// Many2One convert AccountInvoiceReport to *Many2One.
func (air *AccountInvoiceReport) Many2One() *Many2One {
	return NewMany2One(air.Id.Get(), "")
}

// CreateAccountInvoiceReport creates a new account.invoice.report model and returns its id.
func (c *Client) CreateAccountInvoiceReport(air *AccountInvoiceReport) (int64, error) {
	return c.Create(AccountInvoiceReportModel, air)
}

// UpdateAccountInvoiceReport updates an existing account.invoice.report record.
func (c *Client) UpdateAccountInvoiceReport(air *AccountInvoiceReport) error {
	return c.UpdateAccountInvoiceReports([]int64{air.Id.Get()}, air)
}

// UpdateAccountInvoiceReports updates existing account.invoice.report records.
// All records (represented by ids) will be updated by air values.
func (c *Client) UpdateAccountInvoiceReports(ids []int64, air *AccountInvoiceReport) error {
	return c.Update(AccountInvoiceReportModel, ids, air)
}

// DeleteAccountInvoiceReport deletes an existing account.invoice.report record.
func (c *Client) DeleteAccountInvoiceReport(id int64) error {
	return c.DeleteAccountInvoiceReports([]int64{id})
}

// DeleteAccountInvoiceReports deletes existing account.invoice.report records.
func (c *Client) DeleteAccountInvoiceReports(ids []int64) error {
	return c.Delete(AccountInvoiceReportModel, ids)
}

// GetAccountInvoiceReport gets account.invoice.report existing record.
func (c *Client) GetAccountInvoiceReport(id int64) (*AccountInvoiceReport, error) {
	airs, err := c.GetAccountInvoiceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if airs != nil && len(*airs) > 0 {
		return &((*airs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice.report not found", id)
}

// GetAccountInvoiceReports gets account.invoice.report existing records.
func (c *Client) GetAccountInvoiceReports(ids []int64) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.Read(AccountInvoiceReportModel, ids, nil, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReport finds account.invoice.report record by querying it with criteria.
func (c *Client) FindAccountInvoiceReport(criteria *Criteria) (*AccountInvoiceReport, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, NewOptions().Limit(1), airs); err != nil {
		return nil, err
	}
	if airs != nil && len(*airs) > 0 {
		return &((*airs)[0]), nil
	}
	return nil, fmt.Errorf("no account.invoice.report was found with criteria %v", criteria)
}

// FindAccountInvoiceReports finds account.invoice.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReports(criteria *Criteria, options *Options) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, options, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceReportId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.invoice.report was found with criteria %v and options %v", criteria, options)
}
