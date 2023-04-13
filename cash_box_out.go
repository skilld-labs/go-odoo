package odoo

import (
	"fmt"
)

// CashBoxOut represents cash.box.out model.
type CashBoxOut struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Amount      *Float    `xmlrpc:"amount,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CashBoxOuts represents array of cash.box.out model.
type CashBoxOuts []CashBoxOut

// CashBoxOutModel is the odoo model name.
const CashBoxOutModel = "cash.box.out"

// Many2One convert CashBoxOut to *Many2One.
func (cbo *CashBoxOut) Many2One() *Many2One {
	return NewMany2One(cbo.Id.Get(), "")
}

// CreateCashBoxOut creates a new cash.box.out model and returns its id.
func (c *Client) CreateCashBoxOut(cbo *CashBoxOut) (int64, error) {
	ids, err := c.CreateCashBoxOuts([]*CashBoxOut{cbo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCashBoxOut creates a new cash.box.out model and returns its id.
func (c *Client) CreateCashBoxOuts(cbos []*CashBoxOut) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbos {
		vv = append(vv, v)
	}
	return c.Create(CashBoxOutModel, vv)
}

// UpdateCashBoxOut updates an existing cash.box.out record.
func (c *Client) UpdateCashBoxOut(cbo *CashBoxOut) error {
	return c.UpdateCashBoxOuts([]int64{cbo.Id.Get()}, cbo)
}

// UpdateCashBoxOuts updates existing cash.box.out records.
// All records (represented by ids) will be updated by cbo values.
func (c *Client) UpdateCashBoxOuts(ids []int64, cbo *CashBoxOut) error {
	return c.Update(CashBoxOutModel, ids, cbo)
}

// DeleteCashBoxOut deletes an existing cash.box.out record.
func (c *Client) DeleteCashBoxOut(id int64) error {
	return c.DeleteCashBoxOuts([]int64{id})
}

// DeleteCashBoxOuts deletes existing cash.box.out records.
func (c *Client) DeleteCashBoxOuts(ids []int64) error {
	return c.Delete(CashBoxOutModel, ids)
}

// GetCashBoxOut gets cash.box.out existing record.
func (c *Client) GetCashBoxOut(id int64) (*CashBoxOut, error) {
	cbos, err := c.GetCashBoxOuts([]int64{id})
	if err != nil {
		return nil, err
	}
	if cbos != nil && len(*cbos) > 0 {
		return &((*cbos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cash.box.out not found", id)
}

// GetCashBoxOuts gets cash.box.out existing records.
func (c *Client) GetCashBoxOuts(ids []int64) (*CashBoxOuts, error) {
	cbos := &CashBoxOuts{}
	if err := c.Read(CashBoxOutModel, ids, nil, cbos); err != nil {
		return nil, err
	}
	return cbos, nil
}

// FindCashBoxOut finds cash.box.out record by querying it with criteria.
func (c *Client) FindCashBoxOut(criteria *Criteria) (*CashBoxOut, error) {
	cbos := &CashBoxOuts{}
	if err := c.SearchRead(CashBoxOutModel, criteria, NewOptions().Limit(1), cbos); err != nil {
		return nil, err
	}
	if cbos != nil && len(*cbos) > 0 {
		return &((*cbos)[0]), nil
	}
	return nil, fmt.Errorf("cash.box.out was not found with criteria %v", criteria)
}

// FindCashBoxOuts finds cash.box.out records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxOuts(criteria *Criteria, options *Options) (*CashBoxOuts, error) {
	cbos := &CashBoxOuts{}
	if err := c.SearchRead(CashBoxOutModel, criteria, options, cbos); err != nil {
		return nil, err
	}
	return cbos, nil
}

// FindCashBoxOutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxOutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CashBoxOutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCashBoxOutId finds record id by querying it with criteria.
func (c *Client) FindCashBoxOutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CashBoxOutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cash.box.out was not found with criteria %v and options %v", criteria, options)
}
