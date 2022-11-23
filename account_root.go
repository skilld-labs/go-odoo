package odoo

import (
	"fmt"
)

// AccountRoot represents account.root model.
type AccountRoot struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
}

// AccountRoots represents array of account.root model.
type AccountRoots []AccountRoot

// AccountRootModel is the odoo model name.
const AccountRootModel = "account.root"

// Many2One convert AccountRoot to *Many2One.
func (ar *AccountRoot) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAccountRoot creates a new account.root model and returns its id.
func (c *Client) CreateAccountRoot(ar *AccountRoot) (int64, error) {
	return c.Create(AccountRootModel, ar)
}

// UpdateAccountRoot updates an existing account.root record.
func (c *Client) UpdateAccountRoot(ar *AccountRoot) error {
	return c.UpdateAccountRoots([]int64{ar.Id.Get()}, ar)
}

// UpdateAccountRoots updates existing account.root records.
// All records (represented by IDs) will be updated by ar values.
func (c *Client) UpdateAccountRoots(ids []int64, ar *AccountRoot) error {
	return c.Update(AccountRootModel, ids, ar)
}

// DeleteAccountRoot deletes an existing account.root record.
func (c *Client) DeleteAccountRoot(id int64) error {
	return c.DeleteAccountRoots([]int64{id})
}

// DeleteAccountRoots deletes existing account.root records.
func (c *Client) DeleteAccountRoots(ids []int64) error {
	return c.Delete(AccountRootModel, ids)
}

// GetAccountRoot gets account.root existing record.
func (c *Client) GetAccountRoot(id int64) (*AccountRoot, error) {
	ars, err := c.GetAccountRoots([]int64{id})
	if err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.root not found", id)
}

// GetAccountRoots gets account.root existing records.
func (c *Client) GetAccountRoots(ids []int64) (*AccountRoots, error) {
	ars := &AccountRoots{}
	if err := c.Read(AccountRootModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountRoot finds account.root record by querying it with criteria.
func (c *Client) FindAccountRoot(criteria *Criteria) (*AccountRoot, error) {
	ars := &AccountRoots{}
	if err := c.SearchRead(AccountRootModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("account.root was not found")
}

// FindAccountRoots finds account.root records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRoots(criteria *Criteria, options *Options) (*AccountRoots, error) {
	ars := &AccountRoots{}
	if err := c.SearchRead(AccountRootModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountRootIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRootIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountRootModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountRootId finds record id by querying it with criteria.
func (c *Client) FindAccountRootId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountRootModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.root was not found")
}
