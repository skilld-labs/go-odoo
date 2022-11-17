package odoo

import (
	"fmt"
)

// AccountInvoiceLine represents account.invoice.line model.
type AccountInvoiceLine struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omitempty"`
	AccountAnalyticId      *Many2One `xmlrpc:"account_analytic_id,omitempty"`
	AccountId              *Many2One `xmlrpc:"account_id,omitempty"`
	AnalyticTagIds         *Relation `xmlrpc:"analytic_tag_ids,omitempty"`
	CompanyCurrencyId      *Many2One `xmlrpc:"company_currency_id,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId             *Many2One `xmlrpc:"currency_id,omitempty"`
	Discount               *Float    `xmlrpc:"discount,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	InvoiceId              *Many2One `xmlrpc:"invoice_id,omitempty"`
	InvoiceLineTaxIds      *Relation `xmlrpc:"invoice_line_tax_ids,omitempty"`
	IsRoundingLine         *Bool     `xmlrpc:"is_rounding_line,omitempty"`
	LayoutCategoryId       *Many2One `xmlrpc:"layout_category_id,omitempty"`
	LayoutCategorySequence *Int      `xmlrpc:"layout_category_sequence,omitempty"`
	Name                   *String   `xmlrpc:"name,omitempty"`
	Origin                 *String   `xmlrpc:"origin,omitempty"`
	PartnerId              *Many2One `xmlrpc:"partner_id,omitempty"`
	PriceSubtotal          *Float    `xmlrpc:"price_subtotal,omitempty"`
	PriceSubtotalSigned    *Float    `xmlrpc:"price_subtotal_signed,omitempty"`
	PriceTotal             *Float    `xmlrpc:"price_total,omitempty"`
	PriceUnit              *Float    `xmlrpc:"price_unit,omitempty"`
	ProductId              *Many2One `xmlrpc:"product_id,omitempty"`
	ProductImage           *String   `xmlrpc:"product_image,omitempty"`
	PurchaseId             *Many2One `xmlrpc:"purchase_id,omitempty"`
	PurchaseLineId         *Many2One `xmlrpc:"purchase_line_id,omitempty"`
	Quantity               *Float    `xmlrpc:"quantity,omitempty"`
	SaleLineIds            *Relation `xmlrpc:"sale_line_ids,omitempty"`
	Sequence               *Int      `xmlrpc:"sequence,omitempty"`
	UomId                  *Many2One `xmlrpc:"uom_id,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountInvoiceLines represents array of account.invoice.line model.
type AccountInvoiceLines []AccountInvoiceLine

// AccountInvoiceLineModel is the odoo model name.
const AccountInvoiceLineModel = "account.invoice.line"

// Many2One convert AccountInvoiceLine to *Many2One.
func (ail *AccountInvoiceLine) Many2One() *Many2One {
	return NewMany2One(ail.Id.Get(), "")
}

// CreateAccountInvoiceLine creates a new account.invoice.line model and returns its id.
func (c *Client) CreateAccountInvoiceLine(ail *AccountInvoiceLine) (int64, error) {
	return c.Create(AccountInvoiceLineModel, ail)
}

// UpdateAccountInvoiceLine updates an existing account.invoice.line record.
func (c *Client) UpdateAccountInvoiceLine(ail *AccountInvoiceLine) error {
	return c.UpdateAccountInvoiceLines([]int64{ail.Id.Get()}, ail)
}

// UpdateAccountInvoiceLines updates existing account.invoice.line records.
// All records (represented by ids) will be updated by ail values.
func (c *Client) UpdateAccountInvoiceLines(ids []int64, ail *AccountInvoiceLine) error {
	return c.Update(AccountInvoiceLineModel, ids, ail)
}

// DeleteAccountInvoiceLine deletes an existing account.invoice.line record.
func (c *Client) DeleteAccountInvoiceLine(id int64) error {
	return c.DeleteAccountInvoiceLines([]int64{id})
}

// DeleteAccountInvoiceLines deletes existing account.invoice.line records.
func (c *Client) DeleteAccountInvoiceLines(ids []int64) error {
	return c.Delete(AccountInvoiceLineModel, ids)
}

// GetAccountInvoiceLine gets account.invoice.line existing record.
func (c *Client) GetAccountInvoiceLine(id int64) (*AccountInvoiceLine, error) {
	ails, err := c.GetAccountInvoiceLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if ails != nil && len(*ails) > 0 {
		return &((*ails)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice.line not found", id)
}

// GetAccountInvoiceLines gets account.invoice.line existing records.
func (c *Client) GetAccountInvoiceLines(ids []int64) (*AccountInvoiceLines, error) {
	ails := &AccountInvoiceLines{}
	if err := c.Read(AccountInvoiceLineModel, ids, nil, ails); err != nil {
		return nil, err
	}
	return ails, nil
}

// FindAccountInvoiceLine finds account.invoice.line record by querying it with criteria.
func (c *Client) FindAccountInvoiceLine(criteria *Criteria) (*AccountInvoiceLine, error) {
	ails := &AccountInvoiceLines{}
	if err := c.SearchRead(AccountInvoiceLineModel, criteria, NewOptions().Limit(1), ails); err != nil {
		return nil, err
	}
	if ails != nil && len(*ails) > 0 {
		return &((*ails)[0]), nil
	}
	return nil, fmt.Errorf("no account.invoice.line was found with criteria %v", criteria)
}

// FindAccountInvoiceLines finds account.invoice.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceLines(criteria *Criteria, options *Options) (*AccountInvoiceLines, error) {
	ails := &AccountInvoiceLines{}
	if err := c.SearchRead(AccountInvoiceLineModel, criteria, options, ails); err != nil {
		return nil, err
	}
	return ails, nil
}

// FindAccountInvoiceLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceLineId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.invoice.line was found with criteria %v and options %v", criteria, options)
}
