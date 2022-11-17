package odoo

import (
	"fmt"
)

// PaymentTransaction represents payment.transaction model.
type PaymentTransaction struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	AcquirerId        *Many2One  `xmlrpc:"acquirer_id,omitempty"`
	AcquirerReference *String    `xmlrpc:"acquirer_reference,omitempty"`
	Amount            *Float     `xmlrpc:"amount,omitempty"`
	CallbackHash      *String    `xmlrpc:"callback_hash,omitempty"`
	CallbackMethod    *String    `xmlrpc:"callback_method,omitempty"`
	CallbackModelId   *Many2One  `xmlrpc:"callback_model_id,omitempty"`
	CallbackResId     *Int       `xmlrpc:"callback_res_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateValidate      *Time      `xmlrpc:"date_validate,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Fees              *Float     `xmlrpc:"fees,omitempty"`
	Html3ds           *String    `xmlrpc:"html_3ds,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	PartnerAddress    *String    `xmlrpc:"partner_address,omitempty"`
	PartnerCity       *String    `xmlrpc:"partner_city,omitempty"`
	PartnerCountryId  *Many2One  `xmlrpc:"partner_country_id,omitempty"`
	PartnerEmail      *String    `xmlrpc:"partner_email,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerLang       *Selection `xmlrpc:"partner_lang,omitempty"`
	PartnerName       *String    `xmlrpc:"partner_name,omitempty"`
	PartnerPhone      *String    `xmlrpc:"partner_phone,omitempty"`
	PartnerZip        *String    `xmlrpc:"partner_zip,omitempty"`
	PaymentTokenId    *Many2One  `xmlrpc:"payment_token_id,omitempty"`
	Provider          *Selection `xmlrpc:"provider,omitempty"`
	Reference         *String    `xmlrpc:"reference,omitempty"`
	State             *Selection `xmlrpc:"state,omitempty"`
	StateMessage      *String    `xmlrpc:"state_message,omitempty"`
	Type              *Selection `xmlrpc:"type,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PaymentTransactions represents array of payment.transaction model.
type PaymentTransactions []PaymentTransaction

// PaymentTransactionModel is the odoo model name.
const PaymentTransactionModel = "payment.transaction"

// Many2One convert PaymentTransaction to *Many2One.
func (pt *PaymentTransaction) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreatePaymentTransaction creates a new payment.transaction model and returns its id.
func (c *Client) CreatePaymentTransaction(pt *PaymentTransaction) (int64, error) {
	return c.Create(PaymentTransactionModel, pt)
}

// UpdatePaymentTransaction updates an existing payment.transaction record.
func (c *Client) UpdatePaymentTransaction(pt *PaymentTransaction) error {
	return c.UpdatePaymentTransactions([]int64{pt.Id.Get()}, pt)
}

// UpdatePaymentTransactions updates existing payment.transaction records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdatePaymentTransactions(ids []int64, pt *PaymentTransaction) error {
	return c.Update(PaymentTransactionModel, ids, pt)
}

// DeletePaymentTransaction deletes an existing payment.transaction record.
func (c *Client) DeletePaymentTransaction(id int64) error {
	return c.DeletePaymentTransactions([]int64{id})
}

// DeletePaymentTransactions deletes existing payment.transaction records.
func (c *Client) DeletePaymentTransactions(ids []int64) error {
	return c.Delete(PaymentTransactionModel, ids)
}

// GetPaymentTransaction gets payment.transaction existing record.
func (c *Client) GetPaymentTransaction(id int64) (*PaymentTransaction, error) {
	pts, err := c.GetPaymentTransactions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.transaction not found", id)
}

// GetPaymentTransactions gets payment.transaction existing records.
func (c *Client) GetPaymentTransactions(ids []int64) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.Read(PaymentTransactionModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransaction finds payment.transaction record by querying it with criteria.
func (c *Client) FindPaymentTransaction(criteria *Criteria) (*PaymentTransaction, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("no payment.transaction was found with criteria %v", criteria)
}

// FindPaymentTransactions finds payment.transaction records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactions(criteria *Criteria, options *Options) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransactionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentTransactionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentTransactionId finds record id by querying it with criteria.
func (c *Client) FindPaymentTransactionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentTransactionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no payment.transaction was found with criteria %v and options %v", criteria, options)
}
