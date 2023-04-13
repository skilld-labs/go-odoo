package odoo

import (
	"fmt"
)

// AccountMoveLineReconcile represents account.move.line.reconcile model.
type AccountMoveLineReconcile struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Credit      *Float    `xmlrpc:"credit,omptempty"`
	Debit       *Float    `xmlrpc:"debit,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	TransNbr    *Int      `xmlrpc:"trans_nbr,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	Writeoff    *Float    `xmlrpc:"writeoff,omptempty"`
}

// AccountMoveLineReconciles represents array of account.move.line.reconcile model.
type AccountMoveLineReconciles []AccountMoveLineReconcile

// AccountMoveLineReconcileModel is the odoo model name.
const AccountMoveLineReconcileModel = "account.move.line.reconcile"

// Many2One convert AccountMoveLineReconcile to *Many2One.
func (amlr *AccountMoveLineReconcile) Many2One() *Many2One {
	return NewMany2One(amlr.Id.Get(), "")
}

// CreateAccountMoveLineReconcile creates a new account.move.line.reconcile model and returns its id.
func (c *Client) CreateAccountMoveLineReconcile(amlr *AccountMoveLineReconcile) (int64, error) {
	ids, err := c.Create(AccountMoveLineReconcileModel, []interface{}{amlr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveLineReconcile creates a new account.move.line.reconcile model and returns its id.
func (c *Client) CreateAccountMoveLineReconciles(amlrs []*AccountMoveLineReconcile) ([]int64, error) {
	var vv []interface{}
	for _, v := range amlrs {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveLineReconcileModel, vv)
}

// UpdateAccountMoveLineReconcile updates an existing account.move.line.reconcile record.
func (c *Client) UpdateAccountMoveLineReconcile(amlr *AccountMoveLineReconcile) error {
	return c.UpdateAccountMoveLineReconciles([]int64{amlr.Id.Get()}, amlr)
}

// UpdateAccountMoveLineReconciles updates existing account.move.line.reconcile records.
// All records (represented by ids) will be updated by amlr values.
func (c *Client) UpdateAccountMoveLineReconciles(ids []int64, amlr *AccountMoveLineReconcile) error {
	return c.Update(AccountMoveLineReconcileModel, ids, amlr)
}

// DeleteAccountMoveLineReconcile deletes an existing account.move.line.reconcile record.
func (c *Client) DeleteAccountMoveLineReconcile(id int64) error {
	return c.DeleteAccountMoveLineReconciles([]int64{id})
}

// DeleteAccountMoveLineReconciles deletes existing account.move.line.reconcile records.
func (c *Client) DeleteAccountMoveLineReconciles(ids []int64) error {
	return c.Delete(AccountMoveLineReconcileModel, ids)
}

// GetAccountMoveLineReconcile gets account.move.line.reconcile existing record.
func (c *Client) GetAccountMoveLineReconcile(id int64) (*AccountMoveLineReconcile, error) {
	amlrs, err := c.GetAccountMoveLineReconciles([]int64{id})
	if err != nil {
		return nil, err
	}
	if amlrs != nil && len(*amlrs) > 0 {
		return &((*amlrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move.line.reconcile not found", id)
}

// GetAccountMoveLineReconciles gets account.move.line.reconcile existing records.
func (c *Client) GetAccountMoveLineReconciles(ids []int64) (*AccountMoveLineReconciles, error) {
	amlrs := &AccountMoveLineReconciles{}
	if err := c.Read(AccountMoveLineReconcileModel, ids, nil, amlrs); err != nil {
		return nil, err
	}
	return amlrs, nil
}

// FindAccountMoveLineReconcile finds account.move.line.reconcile record by querying it with criteria.
func (c *Client) FindAccountMoveLineReconcile(criteria *Criteria) (*AccountMoveLineReconcile, error) {
	amlrs := &AccountMoveLineReconciles{}
	if err := c.SearchRead(AccountMoveLineReconcileModel, criteria, NewOptions().Limit(1), amlrs); err != nil {
		return nil, err
	}
	if amlrs != nil && len(*amlrs) > 0 {
		return &((*amlrs)[0]), nil
	}
	return nil, fmt.Errorf("account.move.line.reconcile was not found with criteria %v", criteria)
}

// FindAccountMoveLineReconciles finds account.move.line.reconcile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineReconciles(criteria *Criteria, options *Options) (*AccountMoveLineReconciles, error) {
	amlrs := &AccountMoveLineReconciles{}
	if err := c.SearchRead(AccountMoveLineReconcileModel, criteria, options, amlrs); err != nil {
		return nil, err
	}
	return amlrs, nil
}

// FindAccountMoveLineReconcileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineReconcileIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveLineReconcileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveLineReconcileId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineReconcileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineReconcileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move.line.reconcile was not found with criteria %v and options %v", criteria, options)
}
