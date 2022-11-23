package odoo

import (
	"fmt"
)

// PaymentToken represents payment.token model.
type PaymentToken struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AcquirerId  *Many2One `xmlrpc:"acquirer_id,omptempty"`
	AcquirerRef *String   `xmlrpc:"acquirer_ref,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omptempty"`
	PaymentIds  *Relation `xmlrpc:"payment_ids,omptempty"`
	ShortName   *String   `xmlrpc:"short_name,omptempty"`
	Verified    *Bool     `xmlrpc:"verified,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PaymentTokens represents array of payment.token model.
type PaymentTokens []PaymentToken

// PaymentTokenModel is the odoo model name.
const PaymentTokenModel = "payment.token"

// Many2One convert PaymentToken to *Many2One.
func (pt *PaymentToken) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreatePaymentToken creates a new payment.token model and returns its id.
func (c *Client) CreatePaymentToken(pt *PaymentToken) (int64, error) {
	return c.Create(PaymentTokenModel, pt)
}

// UpdatePaymentToken updates an existing payment.token record.
func (c *Client) UpdatePaymentToken(pt *PaymentToken) error {
	return c.UpdatePaymentTokens([]int64{pt.Id.Get()}, pt)
}

// UpdatePaymentTokens updates existing payment.token records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdatePaymentTokens(ids []int64, pt *PaymentToken) error {
	return c.Update(PaymentTokenModel, ids, pt)
}

// DeletePaymentToken deletes an existing payment.token record.
func (c *Client) DeletePaymentToken(id int64) error {
	return c.DeletePaymentTokens([]int64{id})
}

// DeletePaymentTokens deletes existing payment.token records.
func (c *Client) DeletePaymentTokens(ids []int64) error {
	return c.Delete(PaymentTokenModel, ids)
}

// GetPaymentToken gets payment.token existing record.
func (c *Client) GetPaymentToken(id int64) (*PaymentToken, error) {
	pts, err := c.GetPaymentTokens([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.token not found", id)
}

// GetPaymentTokens gets payment.token existing records.
func (c *Client) GetPaymentTokens(ids []int64) (*PaymentTokens, error) {
	pts := &PaymentTokens{}
	if err := c.Read(PaymentTokenModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentToken finds payment.token record by querying it with criteria.
func (c *Client) FindPaymentToken(criteria *Criteria) (*PaymentToken, error) {
	pts := &PaymentTokens{}
	if err := c.SearchRead(PaymentTokenModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("payment.token was not found with criteria %v", criteria)
}

// FindPaymentTokens finds payment.token records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTokens(criteria *Criteria, options *Options) (*PaymentTokens, error) {
	pts := &PaymentTokens{}
	if err := c.SearchRead(PaymentTokenModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTokenIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTokenIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentTokenModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentTokenId finds record id by querying it with criteria.
func (c *Client) FindPaymentTokenId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentTokenModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.token was not found with criteria %v and options %v", criteria, options)
}
