package odoo

import (
	"fmt"
)

// AccountMoveReversal represents account.move.reversal model.
type AccountMoveReversal struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Date        *Time     `xmlrpc:"date,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountMoveReversals represents array of account.move.reversal model.
type AccountMoveReversals []AccountMoveReversal

// AccountMoveReversalModel is the odoo model name.
const AccountMoveReversalModel = "account.move.reversal"

// Many2One convert AccountMoveReversal to *Many2One.
func (amr *AccountMoveReversal) Many2One() *Many2One {
	return NewMany2One(amr.Id.Get(), "")
}

// CreateAccountMoveReversal creates a new account.move.reversal model and returns its id.
func (c *Client) CreateAccountMoveReversal(amr *AccountMoveReversal) (int64, error) {
	ids, err := c.Create(AccountMoveReversalModel, []interface{}{amr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveReversal creates a new account.move.reversal model and returns its id.
func (c *Client) CreateAccountMoveReversals(amrs []*AccountMoveReversal) ([]int64, error) {
	var vv []interface{}
	for _, v := range amrs {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveReversalModel, vv)
}

// UpdateAccountMoveReversal updates an existing account.move.reversal record.
func (c *Client) UpdateAccountMoveReversal(amr *AccountMoveReversal) error {
	return c.UpdateAccountMoveReversals([]int64{amr.Id.Get()}, amr)
}

// UpdateAccountMoveReversals updates existing account.move.reversal records.
// All records (represented by ids) will be updated by amr values.
func (c *Client) UpdateAccountMoveReversals(ids []int64, amr *AccountMoveReversal) error {
	return c.Update(AccountMoveReversalModel, ids, amr)
}

// DeleteAccountMoveReversal deletes an existing account.move.reversal record.
func (c *Client) DeleteAccountMoveReversal(id int64) error {
	return c.DeleteAccountMoveReversals([]int64{id})
}

// DeleteAccountMoveReversals deletes existing account.move.reversal records.
func (c *Client) DeleteAccountMoveReversals(ids []int64) error {
	return c.Delete(AccountMoveReversalModel, ids)
}

// GetAccountMoveReversal gets account.move.reversal existing record.
func (c *Client) GetAccountMoveReversal(id int64) (*AccountMoveReversal, error) {
	amrs, err := c.GetAccountMoveReversals([]int64{id})
	if err != nil {
		return nil, err
	}
	if amrs != nil && len(*amrs) > 0 {
		return &((*amrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move.reversal not found", id)
}

// GetAccountMoveReversals gets account.move.reversal existing records.
func (c *Client) GetAccountMoveReversals(ids []int64) (*AccountMoveReversals, error) {
	amrs := &AccountMoveReversals{}
	if err := c.Read(AccountMoveReversalModel, ids, nil, amrs); err != nil {
		return nil, err
	}
	return amrs, nil
}

// FindAccountMoveReversal finds account.move.reversal record by querying it with criteria.
func (c *Client) FindAccountMoveReversal(criteria *Criteria) (*AccountMoveReversal, error) {
	amrs := &AccountMoveReversals{}
	if err := c.SearchRead(AccountMoveReversalModel, criteria, NewOptions().Limit(1), amrs); err != nil {
		return nil, err
	}
	if amrs != nil && len(*amrs) > 0 {
		return &((*amrs)[0]), nil
	}
	return nil, fmt.Errorf("account.move.reversal was not found with criteria %v", criteria)
}

// FindAccountMoveReversals finds account.move.reversal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveReversals(criteria *Criteria, options *Options) (*AccountMoveReversals, error) {
	amrs := &AccountMoveReversals{}
	if err := c.SearchRead(AccountMoveReversalModel, criteria, options, amrs); err != nil {
		return nil, err
	}
	return amrs, nil
}

// FindAccountMoveReversalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveReversalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveReversalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveReversalId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveReversalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveReversalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move.reversal was not found with criteria %v and options %v", criteria, options)
}
