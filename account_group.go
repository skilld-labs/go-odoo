package odoo

import (
	"fmt"
)

// AccountGroup represents account.group model.
type AccountGroup struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CodePrefix  *String   `xmlrpc:"code_prefix,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	ParentLeft  *Int      `xmlrpc:"parent_left,omptempty"`
	ParentRight *Int      `xmlrpc:"parent_right,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountGroups represents array of account.group model.
type AccountGroups []AccountGroup

// AccountGroupModel is the odoo model name.
const AccountGroupModel = "account.group"

// Many2One convert AccountGroup to *Many2One.
func (ag *AccountGroup) Many2One() *Many2One {
	return NewMany2One(ag.Id.Get(), "")
}

// CreateAccountGroup creates a new account.group model and returns its id.
func (c *Client) CreateAccountGroup(ag *AccountGroup) (int64, error) {
	return c.Create(AccountGroupModel, ag)
}

// UpdateAccountGroup updates an existing account.group record.
func (c *Client) UpdateAccountGroup(ag *AccountGroup) error {
	return c.UpdateAccountGroups([]int64{ag.Id.Get()}, ag)
}

// UpdateAccountGroups updates existing account.group records.
// All records (represented by ids) will be updated by ag values.
func (c *Client) UpdateAccountGroups(ids []int64, ag *AccountGroup) error {
	return c.Update(AccountGroupModel, ids, ag)
}

// DeleteAccountGroup deletes an existing account.group record.
func (c *Client) DeleteAccountGroup(id int64) error {
	return c.DeleteAccountGroups([]int64{id})
}

// DeleteAccountGroups deletes existing account.group records.
func (c *Client) DeleteAccountGroups(ids []int64) error {
	return c.Delete(AccountGroupModel, ids)
}

// GetAccountGroup gets account.group existing record.
func (c *Client) GetAccountGroup(id int64) (*AccountGroup, error) {
	ags, err := c.GetAccountGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if ags != nil && len(*ags) > 0 {
		return &((*ags)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.group not found", id)
}

// GetAccountGroups gets account.group existing records.
func (c *Client) GetAccountGroups(ids []int64) (*AccountGroups, error) {
	ags := &AccountGroups{}
	if err := c.Read(AccountGroupModel, ids, nil, ags); err != nil {
		return nil, err
	}
	return ags, nil
}

// FindAccountGroup finds account.group record by querying it with criteria.
func (c *Client) FindAccountGroup(criteria *Criteria) (*AccountGroup, error) {
	ags := &AccountGroups{}
	if err := c.SearchRead(AccountGroupModel, criteria, NewOptions().Limit(1), ags); err != nil {
		return nil, err
	}
	if ags != nil && len(*ags) > 0 {
		return &((*ags)[0]), nil
	}
	return nil, fmt.Errorf("no account.group was found with criteria %v", criteria)
}

// FindAccountGroups finds account.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroups(criteria *Criteria, options *Options) (*AccountGroups, error) {
	ags := &AccountGroups{}
	if err := c.SearchRead(AccountGroupModel, criteria, options, ags); err != nil {
		return nil, err
	}
	return ags, nil
}

// FindAccountGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.group was found with criteria %v and options %v", criteria, options)
}
