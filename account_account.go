package odoo

import (
	"fmt"
)

// AccountAccount represents account.account model
type AccountAccount struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Code                   *String    `xmlrpc:"code,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omptempty"`
	Deprecated             *Bool      `xmlrpc:"deprecated,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	InternalType           *Selection `xmlrpc:"internal_type,omptempty"`
	LastTimeEntriesChecked *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	Note                   *String    `xmlrpc:"note,omptempty"`
	OpeningCredit          *Float     `xmlrpc:"opening_credit,omptempty"`
	OpeningDebit           *Float     `xmlrpc:"opening_debit,omptempty"`
	Reconcile              *Bool      `xmlrpc:"reconcile,omptempty"`
	TagIds                 *Relation  `xmlrpc:"tag_ids,omptempty"`
	TaxIds                 *Relation  `xmlrpc:"tax_ids,omptempty"`
	UserTypeId             *Many2One  `xmlrpc:"user_type_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAccounts represents array of account.account model
type AccountAccounts []AccountAccount

// AccountAccountModel is the odoo model name
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
	return nil, fmt.Errorf("id %v of %s not found", id, AccountAccountModel)
}

// GetAccountAccounts gets account.account existing records.
func (c *Client) GetAccountAccounts(ids []int64) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.Read(AccountAccountModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccount finds account.account record by querying it with criteria
func (c *Client) FindAccountAccount(criteria *Criteria) (*AccountAccount, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("account.account was not found")
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
