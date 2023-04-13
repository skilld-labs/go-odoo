package odoo

import (
	"fmt"
)

// StockIncoterms represents stock.incoterms model.
type StockIncoterms struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	Code        *String   `xmlrpc:"code,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockIncotermss represents array of stock.incoterms model.
type StockIncotermss []StockIncoterms

// StockIncotermsModel is the odoo model name.
const StockIncotermsModel = "stock.incoterms"

// Many2One convert StockIncoterms to *Many2One.
func (si *StockIncoterms) Many2One() *Many2One {
	return NewMany2One(si.Id.Get(), "")
}

// CreateStockIncoterms creates a new stock.incoterms model and returns its id.
func (c *Client) CreateStockIncoterms(si *StockIncoterms) (int64, error) {
	ids, err := c.Create(StockIncotermsModel, []interface{}{si})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockIncoterms creates a new stock.incoterms model and returns its id.
func (c *Client) CreateStockIncotermss(sis []*StockIncoterms) ([]int64, error) {
	var vv []interface{}
	for _, v := range sis {
		vv = append(vv, v)
	}
	return c.Create(StockIncotermsModel, vv)
}

// UpdateStockIncoterms updates an existing stock.incoterms record.
func (c *Client) UpdateStockIncoterms(si *StockIncoterms) error {
	return c.UpdateStockIncotermss([]int64{si.Id.Get()}, si)
}

// UpdateStockIncotermss updates existing stock.incoterms records.
// All records (represented by ids) will be updated by si values.
func (c *Client) UpdateStockIncotermss(ids []int64, si *StockIncoterms) error {
	return c.Update(StockIncotermsModel, ids, si)
}

// DeleteStockIncoterms deletes an existing stock.incoterms record.
func (c *Client) DeleteStockIncoterms(id int64) error {
	return c.DeleteStockIncotermss([]int64{id})
}

// DeleteStockIncotermss deletes existing stock.incoterms records.
func (c *Client) DeleteStockIncotermss(ids []int64) error {
	return c.Delete(StockIncotermsModel, ids)
}

// GetStockIncoterms gets stock.incoterms existing record.
func (c *Client) GetStockIncoterms(id int64) (*StockIncoterms, error) {
	sis, err := c.GetStockIncotermss([]int64{id})
	if err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.incoterms not found", id)
}

// GetStockIncotermss gets stock.incoterms existing records.
func (c *Client) GetStockIncotermss(ids []int64) (*StockIncotermss, error) {
	sis := &StockIncotermss{}
	if err := c.Read(StockIncotermsModel, ids, nil, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindStockIncoterms finds stock.incoterms record by querying it with criteria.
func (c *Client) FindStockIncoterms(criteria *Criteria) (*StockIncoterms, error) {
	sis := &StockIncotermss{}
	if err := c.SearchRead(StockIncotermsModel, criteria, NewOptions().Limit(1), sis); err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("stock.incoterms was not found with criteria %v", criteria)
}

// FindStockIncotermss finds stock.incoterms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockIncotermss(criteria *Criteria, options *Options) (*StockIncotermss, error) {
	sis := &StockIncotermss{}
	if err := c.SearchRead(StockIncotermsModel, criteria, options, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindStockIncotermsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockIncotermsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockIncotermsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockIncotermsId finds record id by querying it with criteria.
func (c *Client) FindStockIncotermsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockIncotermsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.incoterms was not found with criteria %v and options %v", criteria, options)
}
