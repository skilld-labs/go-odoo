package odoo

import (
	"fmt"
)

// AccountTaxTemplate represents account.tax.template model.
type AccountTaxTemplate struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	Amount                       *Float     `xmlrpc:"amount,omitempty"`
	AmountType                   *Selection `xmlrpc:"amount_type,omitempty"`
	Analytic                     *Bool      `xmlrpc:"analytic,omitempty"`
	CashBasisTransitionAccountId *Many2One  `xmlrpc:"cash_basis_transition_account_id,omitempty"`
	ChartTemplateId              *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	ChildrenTaxIds               *Relation  `xmlrpc:"children_tax_ids,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                  *String    `xmlrpc:"description,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IncludeBaseAmount            *Bool      `xmlrpc:"include_base_amount,omitempty"`
	InvoiceRepartitionLineIds    *Relation  `xmlrpc:"invoice_repartition_line_ids,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	PriceInclude                 *Bool      `xmlrpc:"price_include,omitempty"`
	RefundRepartitionLineIds     *Relation  `xmlrpc:"refund_repartition_line_ids,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	TaxExigibility               *Selection `xmlrpc:"tax_exigibility,omitempty"`
	TaxGroupId                   *Many2One  `xmlrpc:"tax_group_id,omitempty"`
	TaxScope                     *Selection `xmlrpc:"tax_scope,omitempty"`
	TypeTaxUse                   *Selection `xmlrpc:"type_tax_use,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxTemplates represents array of account.tax.template model.
type AccountTaxTemplates []AccountTaxTemplate

// AccountTaxTemplateModel is the odoo model name.
const AccountTaxTemplateModel = "account.tax.template"

// Many2One convert AccountTaxTemplate to *Many2One.
func (att *AccountTaxTemplate) Many2One() *Many2One {
	return NewMany2One(att.Id.Get(), "")
}

// CreateAccountTaxTemplate creates a new account.tax.template model and returns its id.
func (c *Client) CreateAccountTaxTemplate(att *AccountTaxTemplate) (int64, error) {
	return c.Create(AccountTaxTemplateModel, att)
}

// UpdateAccountTaxTemplate updates an existing account.tax.template record.
func (c *Client) UpdateAccountTaxTemplate(att *AccountTaxTemplate) error {
	return c.UpdateAccountTaxTemplates([]int64{att.Id.Get()}, att)
}

// UpdateAccountTaxTemplates updates existing account.tax.template records.
// All records (represented by IDs) will be updated by att values.
func (c *Client) UpdateAccountTaxTemplates(ids []int64, att *AccountTaxTemplate) error {
	return c.Update(AccountTaxTemplateModel, ids, att)
}

// DeleteAccountTaxTemplate deletes an existing account.tax.template record.
func (c *Client) DeleteAccountTaxTemplate(id int64) error {
	return c.DeleteAccountTaxTemplates([]int64{id})
}

// DeleteAccountTaxTemplates deletes existing account.tax.template records.
func (c *Client) DeleteAccountTaxTemplates(ids []int64) error {
	return c.Delete(AccountTaxTemplateModel, ids)
}

// GetAccountTaxTemplate gets account.tax.template existing record.
func (c *Client) GetAccountTaxTemplate(id int64) (*AccountTaxTemplate, error) {
	atts, err := c.GetAccountTaxTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if atts != nil && len(*atts) > 0 {
		return &((*atts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.template not found", id)
}

// GetAccountTaxTemplates gets account.tax.template existing records.
func (c *Client) GetAccountTaxTemplates(ids []int64) (*AccountTaxTemplates, error) {
	atts := &AccountTaxTemplates{}
	if err := c.Read(AccountTaxTemplateModel, ids, nil, atts); err != nil {
		return nil, err
	}
	return atts, nil
}

// FindAccountTaxTemplate finds account.tax.template record by querying it with criteria.
func (c *Client) FindAccountTaxTemplate(criteria *Criteria) (*AccountTaxTemplate, error) {
	atts := &AccountTaxTemplates{}
	if err := c.SearchRead(AccountTaxTemplateModel, criteria, NewOptions().Limit(1), atts); err != nil {
		return nil, err
	}
	if atts != nil && len(*atts) > 0 {
		return &((*atts)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.template was not found")
}

// FindAccountTaxTemplates finds account.tax.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxTemplates(criteria *Criteria, options *Options) (*AccountTaxTemplates, error) {
	atts := &AccountTaxTemplates{}
	if err := c.SearchRead(AccountTaxTemplateModel, criteria, options, atts); err != nil {
		return nil, err
	}
	return atts, nil
}

// FindAccountTaxTemplateIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.template was not found")
}
