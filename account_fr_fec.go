package odoo

import (
	"fmt"
)

// AccountFrFec represents account.fr.fec model.
type AccountFrFec struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo      *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	ExportType  *Selection `xmlrpc:"export_type,omitempty"`
	FecData     *String    `xmlrpc:"fec_data,omitempty"`
	Filename    *String    `xmlrpc:"filename,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFrFecs represents array of account.fr.fec model.
type AccountFrFecs []AccountFrFec

// AccountFrFecModel is the odoo model name.
const AccountFrFecModel = "account.fr.fec"

// Many2One convert AccountFrFec to *Many2One.
func (aff *AccountFrFec) Many2One() *Many2One {
	return NewMany2One(aff.Id.Get(), "")
}

// CreateAccountFrFec creates a new account.fr.fec model and returns its id.
func (c *Client) CreateAccountFrFec(aff *AccountFrFec) (int64, error) {
	return c.Create(AccountFrFecModel, aff)
}

// UpdateAccountFrFec updates an existing account.fr.fec record.
func (c *Client) UpdateAccountFrFec(aff *AccountFrFec) error {
	return c.UpdateAccountFrFecs([]int64{aff.Id.Get()}, aff)
}

// UpdateAccountFrFecs updates existing account.fr.fec records.
// All records (represented by ids) will be updated by aff values.
func (c *Client) UpdateAccountFrFecs(ids []int64, aff *AccountFrFec) error {
	return c.Update(AccountFrFecModel, ids, aff)
}

// DeleteAccountFrFec deletes an existing account.fr.fec record.
func (c *Client) DeleteAccountFrFec(id int64) error {
	return c.DeleteAccountFrFecs([]int64{id})
}

// DeleteAccountFrFecs deletes existing account.fr.fec records.
func (c *Client) DeleteAccountFrFecs(ids []int64) error {
	return c.Delete(AccountFrFecModel, ids)
}

// GetAccountFrFec gets account.fr.fec existing record.
func (c *Client) GetAccountFrFec(id int64) (*AccountFrFec, error) {
	affs, err := c.GetAccountFrFecs([]int64{id})
	if err != nil {
		return nil, err
	}
	if affs != nil && len(*affs) > 0 {
		return &((*affs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fr.fec not found", id)
}

// GetAccountFrFecs gets account.fr.fec existing records.
func (c *Client) GetAccountFrFecs(ids []int64) (*AccountFrFecs, error) {
	affs := &AccountFrFecs{}
	if err := c.Read(AccountFrFecModel, ids, nil, affs); err != nil {
		return nil, err
	}
	return affs, nil
}

// FindAccountFrFec finds account.fr.fec record by querying it with criteria.
func (c *Client) FindAccountFrFec(criteria *Criteria) (*AccountFrFec, error) {
	affs := &AccountFrFecs{}
	if err := c.SearchRead(AccountFrFecModel, criteria, NewOptions().Limit(1), affs); err != nil {
		return nil, err
	}
	if affs != nil && len(*affs) > 0 {
		return &((*affs)[0]), nil
	}
	return nil, fmt.Errorf("no account.fr.fec was found with criteria %v", criteria)
}

// FindAccountFrFecs finds account.fr.fec records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFrFecs(criteria *Criteria, options *Options) (*AccountFrFecs, error) {
	affs := &AccountFrFecs{}
	if err := c.SearchRead(AccountFrFecModel, criteria, options, affs); err != nil {
		return nil, err
	}
	return affs, nil
}

// FindAccountFrFecIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFrFecIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFrFecModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFrFecId finds record id by querying it with criteria.
func (c *Client) FindAccountFrFecId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFrFecModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.fr.fec was found with criteria %v and options %v", criteria, options)
}
