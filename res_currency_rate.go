package odoo

import (
	"fmt"
)

// ResCurrencyRate represents res.currency.rate model.
type ResCurrencyRate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *Time     `xmlrpc:"name,omptempty"`
	Rate        *Float    `xmlrpc:"rate,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResCurrencyRates represents array of res.currency.rate model.
type ResCurrencyRates []ResCurrencyRate

// ResCurrencyRateModel is the odoo model name.
const ResCurrencyRateModel = "res.currency.rate"

// Many2One convert ResCurrencyRate to *Many2One.
func (rcr *ResCurrencyRate) Many2One() *Many2One {
	return NewMany2One(rcr.Id.Get(), "")
}

// CreateResCurrencyRate creates a new res.currency.rate model and returns its id.
func (c *Client) CreateResCurrencyRate(rcr *ResCurrencyRate) (int64, error) {
	return c.Create(ResCurrencyRateModel, rcr)
}

// UpdateResCurrencyRate updates an existing res.currency.rate record.
func (c *Client) UpdateResCurrencyRate(rcr *ResCurrencyRate) error {
	return c.UpdateResCurrencyRates([]int64{rcr.Id.Get()}, rcr)
}

// UpdateResCurrencyRates updates existing res.currency.rate records.
// All records (represented by ids) will be updated by rcr values.
func (c *Client) UpdateResCurrencyRates(ids []int64, rcr *ResCurrencyRate) error {
	return c.Update(ResCurrencyRateModel, ids, rcr)
}

// DeleteResCurrencyRate deletes an existing res.currency.rate record.
func (c *Client) DeleteResCurrencyRate(id int64) error {
	return c.DeleteResCurrencyRates([]int64{id})
}

// DeleteResCurrencyRates deletes existing res.currency.rate records.
func (c *Client) DeleteResCurrencyRates(ids []int64) error {
	return c.Delete(ResCurrencyRateModel, ids)
}

// GetResCurrencyRate gets res.currency.rate existing record.
func (c *Client) GetResCurrencyRate(id int64) (*ResCurrencyRate, error) {
	rcrs, err := c.GetResCurrencyRates([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcrs != nil && len(*rcrs) > 0 {
		return &((*rcrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.currency.rate not found", id)
}

// GetResCurrencyRates gets res.currency.rate existing records.
func (c *Client) GetResCurrencyRates(ids []int64) (*ResCurrencyRates, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.Read(ResCurrencyRateModel, ids, nil, rcrs); err != nil {
		return nil, err
	}
	return rcrs, nil
}

// FindResCurrencyRate finds res.currency.rate record by querying it with criteria.
func (c *Client) FindResCurrencyRate(criteria *Criteria) (*ResCurrencyRate, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.SearchRead(ResCurrencyRateModel, criteria, NewOptions().Limit(1), rcrs); err != nil {
		return nil, err
	}
	if rcrs != nil && len(*rcrs) > 0 {
		return &((*rcrs)[0]), nil
	}
	return nil, fmt.Errorf("no res.currency.rate was found with criteria %v", criteria)
}

// FindResCurrencyRates finds res.currency.rate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencyRates(criteria *Criteria, options *Options) (*ResCurrencyRates, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.SearchRead(ResCurrencyRateModel, criteria, options, rcrs); err != nil {
		return nil, err
	}
	return rcrs, nil
}

// FindResCurrencyRateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencyRateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCurrencyRateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCurrencyRateId finds record id by querying it with criteria.
func (c *Client) FindResCurrencyRateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCurrencyRateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no res.currency.rate was found with criteria %v and options %v", criteria, options)
}
