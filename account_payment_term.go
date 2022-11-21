package odoo

import (
	"fmt"
)

// AccountPaymentTerm represents account.payment.term model.
type AccountPaymentTerm struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LineIds     *Relation `xmlrpc:"line_ids,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Note        *String   `xmlrpc:"note,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountPaymentTerms represents array of account.payment.term model.
type AccountPaymentTerms []AccountPaymentTerm

// AccountPaymentTermModel is the odoo model name.
const AccountPaymentTermModel = "account.payment.term"

// Many2One convert AccountPaymentTerm to *Many2One.
func (apt *AccountPaymentTerm) Many2One() *Many2One {
	return NewMany2One(apt.Id.Get(), "")
}

// CreateAccountPaymentTerm creates a new account.payment.term model and returns its id.
func (c *Client) CreateAccountPaymentTerm(apt *AccountPaymentTerm) (int64, error) {
	return c.Create(AccountPaymentTermModel, apt)
}

// UpdateAccountPaymentTerm updates an existing account.payment.term record.
func (c *Client) UpdateAccountPaymentTerm(apt *AccountPaymentTerm) error {
	return c.UpdateAccountPaymentTerms([]int64{apt.Id.Get()}, apt)
}

// UpdateAccountPaymentTerms updates existing account.payment.term records.
// All records (represented by ids) will be updated by apt values.
func (c *Client) UpdateAccountPaymentTerms(ids []int64, apt *AccountPaymentTerm) error {
	return c.Update(AccountPaymentTermModel, ids, apt)
}

// DeleteAccountPaymentTerm deletes an existing account.payment.term record.
func (c *Client) DeleteAccountPaymentTerm(id int64) error {
	return c.DeleteAccountPaymentTerms([]int64{id})
}

// DeleteAccountPaymentTerms deletes existing account.payment.term records.
func (c *Client) DeleteAccountPaymentTerms(ids []int64) error {
	return c.Delete(AccountPaymentTermModel, ids)
}

// GetAccountPaymentTerm gets account.payment.term existing record.
func (c *Client) GetAccountPaymentTerm(id int64) (*AccountPaymentTerm, error) {
	apts, err := c.GetAccountPaymentTerms([]int64{id})
	if err != nil {
		return nil, err
	}
	if apts != nil && len(*apts) > 0 {
		return &((*apts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment.term not found", id)
}

// GetAccountPaymentTerms gets account.payment.term existing records.
func (c *Client) GetAccountPaymentTerms(ids []int64) (*AccountPaymentTerms, error) {
	apts := &AccountPaymentTerms{}
	if err := c.Read(AccountPaymentTermModel, ids, nil, apts); err != nil {
		return nil, err
	}
	return apts, nil
}

// FindAccountPaymentTerm finds account.payment.term record by querying it with criteria.
func (c *Client) FindAccountPaymentTerm(criteria *Criteria) (*AccountPaymentTerm, error) {
	apts := &AccountPaymentTerms{}
	if err := c.SearchRead(AccountPaymentTermModel, criteria, NewOptions().Limit(1), apts); err != nil {
		return nil, err
	}
	if apts != nil && len(*apts) > 0 {
		return &((*apts)[0]), nil
	}
	return nil, fmt.Errorf("no account.payment.term was found with criteria %v", criteria)
}

// FindAccountPaymentTerms finds account.payment.term records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentTerms(criteria *Criteria, options *Options) (*AccountPaymentTerms, error) {
	apts := &AccountPaymentTerms{}
	if err := c.SearchRead(AccountPaymentTermModel, criteria, options, apts); err != nil {
		return nil, err
	}
	return apts, nil
}

// FindAccountPaymentTermIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentTermIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentTermModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentTermId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentTermId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentTermModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.payment.term was found with criteria %v and options %v", criteria, options)
}
