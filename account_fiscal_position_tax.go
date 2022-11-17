package odoo

import (
	"fmt"
)

// AccountFiscalPositionTax represents account.fiscal.position.tax model.
type AccountFiscalPositionTax struct {
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

// AccountFiscalPositionTaxs represents array of account.fiscal.position.tax model.
type AccountFiscalPositionTaxs []AccountFiscalPositionTax

// AccountFiscalPositionTaxModel is the odoo model name.
const AccountFiscalPositionTaxModel = "account.fiscal.position.tax"

// Many2One convert AccountFiscalPositionTax to *Many2One.
func (afpt *AccountFiscalPositionTax) Many2One() *Many2One {
	return NewMany2One(afpt.Id.Get(), "")
}

// CreateAccountFiscalPositionTax creates a new account.fiscal.position.tax model and returns its id.
func (c *Client) CreateAccountFiscalPositionTax(afpt *AccountFiscalPositionTax) (int64, error) {
	return c.Create(AccountFiscalPositionTaxModel, afpt)
}

// UpdateAccountFiscalPositionTax updates an existing account.fiscal.position.tax record.
func (c *Client) UpdateAccountFiscalPositionTax(afpt *AccountFiscalPositionTax) error {
	return c.UpdateAccountFiscalPositionTaxs([]int64{afpt.Id.Get()}, afpt)
}

// UpdateAccountFiscalPositionTaxs updates existing account.fiscal.position.tax records.
// All records (represented by ids) will be updated by afpt values.
func (c *Client) UpdateAccountFiscalPositionTaxs(ids []int64, afpt *AccountFiscalPositionTax) error {
	return c.Update(AccountFiscalPositionTaxModel, ids, afpt)
}

// DeleteAccountFiscalPositionTax deletes an existing account.fiscal.position.tax record.
func (c *Client) DeleteAccountFiscalPositionTax(id int64) error {
	return c.DeleteAccountFiscalPositionTaxs([]int64{id})
}

// DeleteAccountFiscalPositionTaxs deletes existing account.fiscal.position.tax records.
func (c *Client) DeleteAccountFiscalPositionTaxs(ids []int64) error {
	return c.Delete(AccountFiscalPositionTaxModel, ids)
}

// GetAccountFiscalPositionTax gets account.fiscal.position.tax existing record.
func (c *Client) GetAccountFiscalPositionTax(id int64) (*AccountFiscalPositionTax, error) {
	afpts, err := c.GetAccountFiscalPositionTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if afpts != nil && len(*afpts) > 0 {
		return &((*afpts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.position.tax not found", id)
}

// GetAccountFiscalPositionTaxs gets account.fiscal.position.tax existing records.
func (c *Client) GetAccountFiscalPositionTaxs(ids []int64) (*AccountFiscalPositionTaxs, error) {
	afpts := &AccountFiscalPositionTaxs{}
	if err := c.Read(AccountFiscalPositionTaxModel, ids, nil, afpts); err != nil {
		return nil, err
	}
	return afpts, nil
}

// FindAccountFiscalPositionTax finds account.fiscal.position.tax record by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTax(criteria *Criteria) (*AccountFiscalPositionTax, error) {
	afpts := &AccountFiscalPositionTaxs{}
	if err := c.SearchRead(AccountFiscalPositionTaxModel, criteria, NewOptions().Limit(1), afpts); err != nil {
		return nil, err
	}
	if afpts != nil && len(*afpts) > 0 {
		return &((*afpts)[0]), nil
	}
	return nil, fmt.Errorf("no account.fiscal.position.tax was found with criteria %v", criteria)
}

// FindAccountFiscalPositionTaxs finds account.fiscal.position.tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTaxs(criteria *Criteria, options *Options) (*AccountFiscalPositionTaxs, error) {
	afpts := &AccountFiscalPositionTaxs{}
	if err := c.SearchRead(AccountFiscalPositionTaxModel, criteria, options, afpts); err != nil {
		return nil, err
	}
	return afpts, nil
}

// FindAccountFiscalPositionTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalPositionTaxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalPositionTaxId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.fiscal.position.tax was found with criteria %v and options %v", criteria, options)
}
