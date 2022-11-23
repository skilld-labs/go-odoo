package odoo

import (
	"fmt"
)

// AccountAnalyticDefault represents account.analytic.default model.
type AccountAnalyticDefault struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	AccountId      *Many2One `xmlrpc:"account_id,omitempty"`
	AnalyticId     *Many2One `xmlrpc:"analytic_id,omitempty"`
	AnalyticTagIds *Relation `xmlrpc:"analytic_tag_ids,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DateStart      *Time     `xmlrpc:"date_start,omitempty"`
	DateStop       *Time     `xmlrpc:"date_stop,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omitempty"`
	ProductId      *Many2One `xmlrpc:"product_id,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticDefaults represents array of account.analytic.default model.
type AccountAnalyticDefaults []AccountAnalyticDefault

// AccountAnalyticDefaultModel is the odoo model name.
const AccountAnalyticDefaultModel = "account.analytic.default"

// Many2One convert AccountAnalyticDefault to *Many2One.
func (aad *AccountAnalyticDefault) Many2One() *Many2One {
	return NewMany2One(aad.Id.Get(), "")
}

// CreateAccountAnalyticDefault creates a new account.analytic.default model and returns its id.
func (c *Client) CreateAccountAnalyticDefault(aad *AccountAnalyticDefault) (int64, error) {
	return c.Create(AccountAnalyticDefaultModel, aad)
}

// UpdateAccountAnalyticDefault updates an existing account.analytic.default record.
func (c *Client) UpdateAccountAnalyticDefault(aad *AccountAnalyticDefault) error {
	return c.UpdateAccountAnalyticDefaults([]int64{aad.Id.Get()}, aad)
}

// UpdateAccountAnalyticDefaults updates existing account.analytic.default records.
// All records (represented by IDs) will be updated by aad values.
func (c *Client) UpdateAccountAnalyticDefaults(ids []int64, aad *AccountAnalyticDefault) error {
	return c.Update(AccountAnalyticDefaultModel, ids, aad)
}

// DeleteAccountAnalyticDefault deletes an existing account.analytic.default record.
func (c *Client) DeleteAccountAnalyticDefault(id int64) error {
	return c.DeleteAccountAnalyticDefaults([]int64{id})
}

// DeleteAccountAnalyticDefaults deletes existing account.analytic.default records.
func (c *Client) DeleteAccountAnalyticDefaults(ids []int64) error {
	return c.Delete(AccountAnalyticDefaultModel, ids)
}

// GetAccountAnalyticDefault gets account.analytic.default existing record.
func (c *Client) GetAccountAnalyticDefault(id int64) (*AccountAnalyticDefault, error) {
	aads, err := c.GetAccountAnalyticDefaults([]int64{id})
	if err != nil {
		return nil, err
	}
	if aads != nil && len(*aads) > 0 {
		return &((*aads)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.default not found", id)
}

// GetAccountAnalyticDefaults gets account.analytic.default existing records.
func (c *Client) GetAccountAnalyticDefaults(ids []int64) (*AccountAnalyticDefaults, error) {
	aads := &AccountAnalyticDefaults{}
	if err := c.Read(AccountAnalyticDefaultModel, ids, nil, aads); err != nil {
		return nil, err
	}
	return aads, nil
}

// FindAccountAnalyticDefault finds account.analytic.default record by querying it with criteria.
func (c *Client) FindAccountAnalyticDefault(criteria *Criteria) (*AccountAnalyticDefault, error) {
	aads := &AccountAnalyticDefaults{}
	if err := c.SearchRead(AccountAnalyticDefaultModel, criteria, NewOptions().Limit(1), aads); err != nil {
		return nil, err
	}
	if aads != nil && len(*aads) > 0 {
		return &((*aads)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.default was not found")
}

// FindAccountAnalyticDefaults finds account.analytic.default records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDefaults(criteria *Criteria, options *Options) (*AccountAnalyticDefaults, error) {
	aads := &AccountAnalyticDefaults{}
	if err := c.SearchRead(AccountAnalyticDefaultModel, criteria, options, aads); err != nil {
		return nil, err
	}
	return aads, nil
}

// FindAccountAnalyticDefaultIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDefaultIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticDefaultModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticDefaultId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticDefaultId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticDefaultModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.default was not found")
}
