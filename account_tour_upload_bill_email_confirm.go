package odoo

import (
	"fmt"
)

// AccountTourUploadBillEmailConfirm represents account.tour.upload.bill.email.confirm model.
type AccountTourUploadBillEmailConfirm struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	EmailAlias  *String   `xmlrpc:"email_alias,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTourUploadBillEmailConfirms represents array of account.tour.upload.bill.email.confirm model.
type AccountTourUploadBillEmailConfirms []AccountTourUploadBillEmailConfirm

// AccountTourUploadBillEmailConfirmModel is the odoo model name.
const AccountTourUploadBillEmailConfirmModel = "account.tour.upload.bill.email.confirm"

// Many2One convert AccountTourUploadBillEmailConfirm to *Many2One.
func (atubec *AccountTourUploadBillEmailConfirm) Many2One() *Many2One {
	return NewMany2One(atubec.Id.Get(), "")
}

// CreateAccountTourUploadBillEmailConfirm creates a new account.tour.upload.bill.email.confirm model and returns its id.
func (c *Client) CreateAccountTourUploadBillEmailConfirm(atubec *AccountTourUploadBillEmailConfirm) (int64, error) {
	return c.Create(AccountTourUploadBillEmailConfirmModel, atubec)
}

// UpdateAccountTourUploadBillEmailConfirm updates an existing account.tour.upload.bill.email.confirm record.
func (c *Client) UpdateAccountTourUploadBillEmailConfirm(atubec *AccountTourUploadBillEmailConfirm) error {
	return c.UpdateAccountTourUploadBillEmailConfirms([]int64{atubec.Id.Get()}, atubec)
}

// UpdateAccountTourUploadBillEmailConfirms updates existing account.tour.upload.bill.email.confirm records.
// All records (represented by ids) will be updated by atubec values.
func (c *Client) UpdateAccountTourUploadBillEmailConfirms(ids []int64, atubec *AccountTourUploadBillEmailConfirm) error {
	return c.Update(AccountTourUploadBillEmailConfirmModel, ids, atubec)
}

// DeleteAccountTourUploadBillEmailConfirm deletes an existing account.tour.upload.bill.email.confirm record.
func (c *Client) DeleteAccountTourUploadBillEmailConfirm(id int64) error {
	return c.DeleteAccountTourUploadBillEmailConfirms([]int64{id})
}

// DeleteAccountTourUploadBillEmailConfirms deletes existing account.tour.upload.bill.email.confirm records.
func (c *Client) DeleteAccountTourUploadBillEmailConfirms(ids []int64) error {
	return c.Delete(AccountTourUploadBillEmailConfirmModel, ids)
}

// GetAccountTourUploadBillEmailConfirm gets account.tour.upload.bill.email.confirm existing record.
func (c *Client) GetAccountTourUploadBillEmailConfirm(id int64) (*AccountTourUploadBillEmailConfirm, error) {
	atubecs, err := c.GetAccountTourUploadBillEmailConfirms([]int64{id})
	if err != nil {
		return nil, err
	}
	if atubecs != nil && len(*atubecs) > 0 {
		return &((*atubecs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tour.upload.bill.email.confirm not found", id)
}

// GetAccountTourUploadBillEmailConfirms gets account.tour.upload.bill.email.confirm existing records.
func (c *Client) GetAccountTourUploadBillEmailConfirms(ids []int64) (*AccountTourUploadBillEmailConfirms, error) {
	atubecs := &AccountTourUploadBillEmailConfirms{}
	if err := c.Read(AccountTourUploadBillEmailConfirmModel, ids, nil, atubecs); err != nil {
		return nil, err
	}
	return atubecs, nil
}

// FindAccountTourUploadBillEmailConfirm finds account.tour.upload.bill.email.confirm record by querying it with criteria.
func (c *Client) FindAccountTourUploadBillEmailConfirm(criteria *Criteria) (*AccountTourUploadBillEmailConfirm, error) {
	atubecs := &AccountTourUploadBillEmailConfirms{}
	if err := c.SearchRead(AccountTourUploadBillEmailConfirmModel, criteria, NewOptions().Limit(1), atubecs); err != nil {
		return nil, err
	}
	if atubecs != nil && len(*atubecs) > 0 {
		return &((*atubecs)[0]), nil
	}
	return nil, fmt.Errorf("account.tour.upload.bill.email.confirm was not found")
}

// FindAccountTourUploadBillEmailConfirms finds account.tour.upload.bill.email.confirm records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBillEmailConfirms(criteria *Criteria, options *Options) (*AccountTourUploadBillEmailConfirms, error) {
	atubecs := &AccountTourUploadBillEmailConfirms{}
	if err := c.SearchRead(AccountTourUploadBillEmailConfirmModel, criteria, options, atubecs); err != nil {
		return nil, err
	}
	return atubecs, nil
}

// FindAccountTourUploadBillEmailConfirmIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBillEmailConfirmIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTourUploadBillEmailConfirmModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTourUploadBillEmailConfirmId finds record id by querying it with criteria.
func (c *Client) FindAccountTourUploadBillEmailConfirmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTourUploadBillEmailConfirmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tour.upload.bill.email.confirm was not found")
}
