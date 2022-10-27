package odoo

import (
	"fmt"
)

// AccountAccountType represents account.account.type model.
type AccountAccountType struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	IncludeInitialBalance *Bool      `xmlrpc:"include_initial_balance,omitempty"`
	InternalGroup         *Selection `xmlrpc:"internal_group,omitempty"`
	Name                  *String    `xmlrpc:"name,omitempty"`
	Note                  *String    `xmlrpc:"note,omitempty"`
	Type                  *Selection `xmlrpc:"type,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAccountTypes represents array of account.account.type model.
type AccountAccountTypes []AccountAccountType

// AccountAccountTypeModel is the odoo model name.
const AccountAccountTypeModel = "account.account.type"

// Many2One convert AccountAccountType to *Many2One.
func (aat *AccountAccountType) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAccountType creates a new account.account.type model and returns its id.
func (c *Client) CreateAccountAccountType(aat *AccountAccountType) (int64, error) {
	return c.Create(AccountAccountTypeModel, aat)
}

// UpdateAccountAccountType updates an existing account.account.type record.
func (c *Client) UpdateAccountAccountType(aat *AccountAccountType) error {
	return c.UpdateAccountAccountTypes([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAccountTypes updates existing account.account.type records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAccountTypes(ids []int64, aat *AccountAccountType) error {
	return c.Update(AccountAccountTypeModel, ids, aat)
}

// DeleteAccountAccountType deletes an existing account.account.type record.
func (c *Client) DeleteAccountAccountType(id int64) error {
	return c.DeleteAccountAccountTypes([]int64{id})
}

// DeleteAccountAccountTypes deletes existing account.account.type records.
func (c *Client) DeleteAccountAccountTypes(ids []int64) error {
	return c.Delete(AccountAccountTypeModel, ids)
}

// GetAccountAccountType gets account.account.type existing record.
func (c *Client) GetAccountAccountType(id int64) (*AccountAccountType, error) {
	aats, err := c.GetAccountAccountTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account.type not found", id)
}

// GetAccountAccountTypes gets account.account.type existing records.
func (c *Client) GetAccountAccountTypes(ids []int64) (*AccountAccountTypes, error) {
	aats := &AccountAccountTypes{}
	if err := c.Read(AccountAccountTypeModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountType finds account.account.type record by querying it with criteria.
func (c *Client) FindAccountAccountType(criteria *Criteria) (*AccountAccountType, error) {
	aats := &AccountAccountTypes{}
	if err := c.SearchRead(AccountAccountTypeModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("account.account.type was not found")
}

// FindAccountAccountTypes finds account.account.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTypes(criteria *Criteria, options *Options) (*AccountAccountTypes, error) {
	aats := &AccountAccountTypes{}
	if err := c.SearchRead(AccountAccountTypeModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountTypeId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.account.type was not found")
}
