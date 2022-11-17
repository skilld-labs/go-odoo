package odoo

import (
	"fmt"
)

// AccountAccount represents account.account model.
type AccountAccount struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	Code                   *String    `xmlrpc:"code,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omitempty"`
	Deprecated             *Bool      `xmlrpc:"deprecated,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	InternalType           *Selection `xmlrpc:"internal_type,omitempty"`
	LastTimeEntriesChecked *Time      `xmlrpc:"last_time_entries_checked,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	Note                   *String    `xmlrpc:"note,omitempty"`
	OpeningCredit          *Float     `xmlrpc:"opening_credit,omitempty"`
	OpeningDebit           *Float     `xmlrpc:"opening_debit,omitempty"`
	Reconcile              *Bool      `xmlrpc:"reconcile,omitempty"`
	TagIds                 *Relation  `xmlrpc:"tag_ids,omitempty"`
	TaxIds                 *Relation  `xmlrpc:"tax_ids,omitempty"`
	UserTypeId             *Many2One  `xmlrpc:"user_type_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAccounts represents array of account.account model.
type AccountAccounts []AccountAccount

// AccountAccountModel is the odoo model name.
const AccountAccountModel = "account.account"

// Many2One convert AccountAccount to *Many2One.
func (aa *AccountAccount) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAccount creates a new account.account model and returns its id.
func (c *Client) CreateAccountAccount(aa *AccountAccount) (int64, error) {
	return c.Create(AccountAccountModel, aa)
}

// UpdateAccountAccount updates an existing account.account record.
func (c *Client) UpdateAccountAccount(aa *AccountAccount) error {
	return c.UpdateAccountAccounts([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAccounts updates existing account.account records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAccounts(ids []int64, aa *AccountAccount) error {
	return c.Update(AccountAccountModel, ids, aa)
}

// DeleteAccountAccount deletes an existing account.account record.
func (c *Client) DeleteAccountAccount(id int64) error {
	return c.DeleteAccountAccounts([]int64{id})
}

// DeleteAccountAccounts deletes existing account.account records.
func (c *Client) DeleteAccountAccounts(ids []int64) error {
	return c.Delete(AccountAccountModel, ids)
}

// GetAccountAccount gets account.account existing record.
func (c *Client) GetAccountAccount(id int64) (*AccountAccount, error) {
	aas, err := c.GetAccountAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account not found", id)
}

// GetAccountAccounts gets account.account existing records.
func (c *Client) GetAccountAccounts(ids []int64) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.Read(AccountAccountModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccount finds account.account record by querying it with criteria.
func (c *Client) FindAccountAccount(criteria *Criteria) (*AccountAccount, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("no account.account was found with criteria %v", criteria)
}

// FindAccountAccounts finds account.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccounts(criteria *Criteria, options *Options) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.account was found with criteria %v and options %v", criteria, options)
}
