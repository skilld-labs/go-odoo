package odoo

import (
	"fmt"
)

// AccountFiscalPositionTaxTemplate represents account.fiscal.position.tax.template model.
type AccountFiscalPositionTaxTemplate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PositionId  *Many2One `xmlrpc:"position_id,omitempty"`
	TaxDestId   *Many2One `xmlrpc:"tax_dest_id,omitempty"`
	TaxSrcId    *Many2One `xmlrpc:"tax_src_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFiscalPositionTaxTemplates represents array of account.fiscal.position.tax.template model.
type AccountFiscalPositionTaxTemplates []AccountFiscalPositionTaxTemplate

// AccountFiscalPositionTaxTemplateModel is the odoo model name.
const AccountFiscalPositionTaxTemplateModel = "account.fiscal.position.tax.template"

// Many2One convert AccountFiscalPositionTaxTemplate to *Many2One.
func (afptt *AccountFiscalPositionTaxTemplate) Many2One() *Many2One {
	return NewMany2One(afptt.Id.Get(), "")
}

// CreateAccountFiscalPositionTaxTemplate creates a new account.fiscal.position.tax.template model and returns its id.
func (c *Client) CreateAccountFiscalPositionTaxTemplate(afptt *AccountFiscalPositionTaxTemplate) (int64, error) {
	return c.Create(AccountFiscalPositionTaxTemplateModel, afptt)
}

// UpdateAccountFiscalPositionTaxTemplate updates an existing account.fiscal.position.tax.template record.
func (c *Client) UpdateAccountFiscalPositionTaxTemplate(afptt *AccountFiscalPositionTaxTemplate) error {
	return c.UpdateAccountFiscalPositionTaxTemplates([]int64{afptt.Id.Get()}, afptt)
}

// UpdateAccountFiscalPositionTaxTemplates updates existing account.fiscal.position.tax.template records.
// All records (represented by IDs) will be updated by afptt values.
func (c *Client) UpdateAccountFiscalPositionTaxTemplates(ids []int64, afptt *AccountFiscalPositionTaxTemplate) error {
	return c.Update(AccountFiscalPositionTaxTemplateModel, ids, afptt)
}

// DeleteAccountFiscalPositionTaxTemplate deletes an existing account.fiscal.position.tax.template record.
func (c *Client) DeleteAccountFiscalPositionTaxTemplate(id int64) error {
	return c.DeleteAccountFiscalPositionTaxTemplates([]int64{id})
}

// DeleteAccountFiscalPositionTaxTemplates deletes existing account.fiscal.position.tax.template records.
func (c *Client) DeleteAccountFiscalPositionTaxTemplates(ids []int64) error {
	return c.Delete(AccountFiscalPositionTaxTemplateModel, ids)
}

// GetAccountFiscalPositionTaxTemplate gets account.fiscal.position.tax.template existing record.
func (c *Client) GetAccountFiscalPositionTaxTemplate(id int64) (*AccountFiscalPositionTaxTemplate, error) {
	afptts, err := c.GetAccountFiscalPositionTaxTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if afptts != nil && len(*afptts) > 0 {
		return &((*afptts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.position.tax.template not found", id)
}

// GetAccountFiscalPositionTaxTemplates gets account.fiscal.position.tax.template existing records.
func (c *Client) GetAccountFiscalPositionTaxTemplates(ids []int64) (*AccountFiscalPositionTaxTemplates, error) {
	afptts := &AccountFiscalPositionTaxTemplates{}
	if err := c.Read(AccountFiscalPositionTaxTemplateModel, ids, nil, afptts); err != nil {
		return nil, err
	}
	return afptts, nil
}

// FindAccountFiscalPositionTaxTemplate finds account.fiscal.position.tax.template record by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTaxTemplate(criteria *Criteria) (*AccountFiscalPositionTaxTemplate, error) {
	afptts := &AccountFiscalPositionTaxTemplates{}
	if err := c.SearchRead(AccountFiscalPositionTaxTemplateModel, criteria, NewOptions().Limit(1), afptts); err != nil {
		return nil, err
	}
	if afptts != nil && len(*afptts) > 0 {
		return &((*afptts)[0]), nil
	}
	return nil, fmt.Errorf("account.fiscal.position.tax.template was not found")
}

// FindAccountFiscalPositionTaxTemplates finds account.fiscal.position.tax.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTaxTemplates(criteria *Criteria, options *Options) (*AccountFiscalPositionTaxTemplates, error) {
	afptts := &AccountFiscalPositionTaxTemplates{}
	if err := c.SearchRead(AccountFiscalPositionTaxTemplateModel, criteria, options, afptts); err != nil {
		return nil, err
	}
	return afptts, nil
}

// FindAccountFiscalPositionTaxTemplateIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTaxTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalPositionTaxTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalPositionTaxTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTaxTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionTaxTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.fiscal.position.tax.template was not found")
}
