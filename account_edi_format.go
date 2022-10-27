package odoo

import (
	"fmt"
)

// AccountEdiFormat represents account.edi.format model.
type AccountEdiFormat struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountEdiFormats represents array of account.edi.format model.
type AccountEdiFormats []AccountEdiFormat

// AccountEdiFormatModel is the odoo model name.
const AccountEdiFormatModel = "account.edi.format"

// Many2One convert AccountEdiFormat to *Many2One.
func (aef *AccountEdiFormat) Many2One() *Many2One {
	return NewMany2One(aef.Id.Get(), "")
}

// CreateAccountEdiFormat creates a new account.edi.format model and returns its id.
func (c *Client) CreateAccountEdiFormat(aef *AccountEdiFormat) (int64, error) {
	return c.Create(AccountEdiFormatModel, aef)
}

// UpdateAccountEdiFormat updates an existing account.edi.format record.
func (c *Client) UpdateAccountEdiFormat(aef *AccountEdiFormat) error {
	return c.UpdateAccountEdiFormats([]int64{aef.Id.Get()}, aef)
}

// UpdateAccountEdiFormats updates existing account.edi.format records.
// All records (represented by ids) will be updated by aef values.
func (c *Client) UpdateAccountEdiFormats(ids []int64, aef *AccountEdiFormat) error {
	return c.Update(AccountEdiFormatModel, ids, aef)
}

// DeleteAccountEdiFormat deletes an existing account.edi.format record.
func (c *Client) DeleteAccountEdiFormat(id int64) error {
	return c.DeleteAccountEdiFormats([]int64{id})
}

// DeleteAccountEdiFormats deletes existing account.edi.format records.
func (c *Client) DeleteAccountEdiFormats(ids []int64) error {
	return c.Delete(AccountEdiFormatModel, ids)
}

// GetAccountEdiFormat gets account.edi.format existing record.
func (c *Client) GetAccountEdiFormat(id int64) (*AccountEdiFormat, error) {
	aefs, err := c.GetAccountEdiFormats([]int64{id})
	if err != nil {
		return nil, err
	}
	if aefs != nil && len(*aefs) > 0 {
		return &((*aefs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.edi.format not found", id)
}

// GetAccountEdiFormats gets account.edi.format existing records.
func (c *Client) GetAccountEdiFormats(ids []int64) (*AccountEdiFormats, error) {
	aefs := &AccountEdiFormats{}
	if err := c.Read(AccountEdiFormatModel, ids, nil, aefs); err != nil {
		return nil, err
	}
	return aefs, nil
}

// FindAccountEdiFormat finds account.edi.format record by querying it with criteria.
func (c *Client) FindAccountEdiFormat(criteria *Criteria) (*AccountEdiFormat, error) {
	aefs := &AccountEdiFormats{}
	if err := c.SearchRead(AccountEdiFormatModel, criteria, NewOptions().Limit(1), aefs); err != nil {
		return nil, err
	}
	if aefs != nil && len(*aefs) > 0 {
		return &((*aefs)[0]), nil
	}
	return nil, fmt.Errorf("account.edi.format was not found")
}

// FindAccountEdiFormats finds account.edi.format records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiFormats(criteria *Criteria, options *Options) (*AccountEdiFormats, error) {
	aefs := &AccountEdiFormats{}
	if err := c.SearchRead(AccountEdiFormatModel, criteria, options, aefs); err != nil {
		return nil, err
	}
	return aefs, nil
}

// FindAccountEdiFormatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiFormatIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountEdiFormatModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountEdiFormatId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiFormatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiFormatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.edi.format was not found")
}
