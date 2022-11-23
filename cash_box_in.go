package odoo

import (
	"fmt"
)

// CashBoxIn represents cash.box.in model.
type CashBoxIn struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Amount      *Float    `xmlrpc:"amount,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Ref         *String   `xmlrpc:"ref,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CashBoxIns represents array of cash.box.in model.
type CashBoxIns []CashBoxIn

// CashBoxInModel is the odoo model name.
const CashBoxInModel = "cash.box.in"

// Many2One convert CashBoxIn to *Many2One.
func (cbi *CashBoxIn) Many2One() *Many2One {
	return NewMany2One(cbi.Id.Get(), "")
}

// CreateCashBoxIn creates a new cash.box.in model and returns its id.
func (c *Client) CreateCashBoxIn(cbi *CashBoxIn) (int64, error) {
	return c.Create(CashBoxInModel, cbi)
}

// UpdateCashBoxIn updates an existing cash.box.in record.
func (c *Client) UpdateCashBoxIn(cbi *CashBoxIn) error {
	return c.UpdateCashBoxIns([]int64{cbi.Id.Get()}, cbi)
}

// UpdateCashBoxIns updates existing cash.box.in records.
// All records (represented by ids) will be updated by cbi values.
func (c *Client) UpdateCashBoxIns(ids []int64, cbi *CashBoxIn) error {
	return c.Update(CashBoxInModel, ids, cbi)
}

// DeleteCashBoxIn deletes an existing cash.box.in record.
func (c *Client) DeleteCashBoxIn(id int64) error {
	return c.DeleteCashBoxIns([]int64{id})
}

// DeleteCashBoxIns deletes existing cash.box.in records.
func (c *Client) DeleteCashBoxIns(ids []int64) error {
	return c.Delete(CashBoxInModel, ids)
}

// GetCashBoxIn gets cash.box.in existing record.
func (c *Client) GetCashBoxIn(id int64) (*CashBoxIn, error) {
	cbis, err := c.GetCashBoxIns([]int64{id})
	if err != nil {
		return nil, err
	}
	if cbis != nil && len(*cbis) > 0 {
		return &((*cbis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cash.box.in not found", id)
}

// GetCashBoxIns gets cash.box.in existing records.
func (c *Client) GetCashBoxIns(ids []int64) (*CashBoxIns, error) {
	cbis := &CashBoxIns{}
	if err := c.Read(CashBoxInModel, ids, nil, cbis); err != nil {
		return nil, err
	}
	return cbis, nil
}

// FindCashBoxIn finds cash.box.in record by querying it with criteria.
func (c *Client) FindCashBoxIn(criteria *Criteria) (*CashBoxIn, error) {
	cbis := &CashBoxIns{}
	if err := c.SearchRead(CashBoxInModel, criteria, NewOptions().Limit(1), cbis); err != nil {
		return nil, err
	}
	if cbis != nil && len(*cbis) > 0 {
		return &((*cbis)[0]), nil
	}
	return nil, fmt.Errorf("cash.box.in was not found with criteria %v", criteria)
}

// FindCashBoxIns finds cash.box.in records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxIns(criteria *Criteria, options *Options) (*CashBoxIns, error) {
	cbis := &CashBoxIns{}
	if err := c.SearchRead(CashBoxInModel, criteria, options, cbis); err != nil {
		return nil, err
	}
	return cbis, nil
}

// FindCashBoxInIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxInIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CashBoxInModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCashBoxInId finds record id by querying it with criteria.
func (c *Client) FindCashBoxInId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CashBoxInModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cash.box.in was not found with criteria %v and options %v", criteria, options)
}
