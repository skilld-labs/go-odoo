package odoo

import (
	"fmt"
)

// AccountRegisterPayments represents account.register.payments model.
type AccountRegisterPayments struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Amount            *Float     `xmlrpc:"amount,omptempty"`
	Communication     *String    `xmlrpc:"communication,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	HidePaymentMethod *Bool      `xmlrpc:"hide_payment_method,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	InvoiceIds        *Relation  `xmlrpc:"invoice_ids,omptempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omptempty"`
	Multi             *Bool      `xmlrpc:"multi,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerType       *Selection `xmlrpc:"partner_type,omptempty"`
	PaymentDate       *Time      `xmlrpc:"payment_date,omptempty"`
	PaymentMethodCode *String    `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodId   *Many2One  `xmlrpc:"payment_method_id,omptempty"`
	PaymentType       *Selection `xmlrpc:"payment_type,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountRegisterPaymentss represents array of account.register.payments model.
type AccountRegisterPaymentss []AccountRegisterPayments

// AccountRegisterPaymentsModel is the odoo model name.
const AccountRegisterPaymentsModel = "account.register.payments"

// Many2One convert AccountRegisterPayments to *Many2One.
func (arp *AccountRegisterPayments) Many2One() *Many2One {
	return NewMany2One(arp.Id.Get(), "")
}

// CreateAccountRegisterPayments creates a new account.register.payments model and returns its id.
func (c *Client) CreateAccountRegisterPayments(arp *AccountRegisterPayments) (int64, error) {
	return c.Create(AccountRegisterPaymentsModel, arp)
}

// UpdateAccountRegisterPayments updates an existing account.register.payments record.
func (c *Client) UpdateAccountRegisterPayments(arp *AccountRegisterPayments) error {
	return c.UpdateAccountRegisterPaymentss([]int64{arp.Id.Get()}, arp)
}

// UpdateAccountRegisterPaymentss updates existing account.register.payments records.
// All records (represented by ids) will be updated by arp values.
func (c *Client) UpdateAccountRegisterPaymentss(ids []int64, arp *AccountRegisterPayments) error {
	return c.Update(AccountRegisterPaymentsModel, ids, arp)
}

// DeleteAccountRegisterPayments deletes an existing account.register.payments record.
func (c *Client) DeleteAccountRegisterPayments(id int64) error {
	return c.DeleteAccountRegisterPaymentss([]int64{id})
}

// DeleteAccountRegisterPaymentss deletes existing account.register.payments records.
func (c *Client) DeleteAccountRegisterPaymentss(ids []int64) error {
	return c.Delete(AccountRegisterPaymentsModel, ids)
}

// GetAccountRegisterPayments gets account.register.payments existing record.
func (c *Client) GetAccountRegisterPayments(id int64) (*AccountRegisterPayments, error) {
	arps, err := c.GetAccountRegisterPaymentss([]int64{id})
	if err != nil {
		return nil, err
	}
	if arps != nil && len(*arps) > 0 {
		return &((*arps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.register.payments not found", id)
}

// GetAccountRegisterPaymentss gets account.register.payments existing records.
func (c *Client) GetAccountRegisterPaymentss(ids []int64) (*AccountRegisterPaymentss, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.Read(AccountRegisterPaymentsModel, ids, nil, arps); err != nil {
		return nil, err
	}
	return arps, nil
}

// FindAccountRegisterPayments finds account.register.payments record by querying it with criteria.
func (c *Client) FindAccountRegisterPayments(criteria *Criteria) (*AccountRegisterPayments, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.SearchRead(AccountRegisterPaymentsModel, criteria, NewOptions().Limit(1), arps); err != nil {
		return nil, err
	}
	if arps != nil && len(*arps) > 0 {
		return &((*arps)[0]), nil
	}
	return nil, fmt.Errorf("account.register.payments was not found with criteria %v", criteria)
}

// FindAccountRegisterPaymentss finds account.register.payments records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRegisterPaymentss(criteria *Criteria, options *Options) (*AccountRegisterPaymentss, error) {
	arps := &AccountRegisterPaymentss{}
	if err := c.SearchRead(AccountRegisterPaymentsModel, criteria, options, arps); err != nil {
		return nil, err
	}
	return arps, nil
}

// FindAccountRegisterPaymentsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountRegisterPaymentsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountRegisterPaymentsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountRegisterPaymentsId finds record id by querying it with criteria.
func (c *Client) FindAccountRegisterPaymentsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountRegisterPaymentsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.register.payments was not found with criteria %v and options %v", criteria, options)
}
