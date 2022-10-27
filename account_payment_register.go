package odoo

import (
	"fmt"
)

// AccountPaymentRegister represents account.payment.register model.
type AccountPaymentRegister struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omitempty"`
	Amount                         *Float     `xmlrpc:"amount,omitempty"`
	AvailablePaymentMethodIds      *Relation  `xmlrpc:"available_payment_method_ids,omitempty"`
	CanEditWizard                  *Bool      `xmlrpc:"can_edit_wizard,omitempty"`
	CanGroupPayments               *Bool      `xmlrpc:"can_group_payments,omitempty"`
	Communication                  *String    `xmlrpc:"communication,omitempty"`
	CompanyCurrencyId              *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                    *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	GroupPayment                   *Bool      `xmlrpc:"group_payment,omitempty"`
	HidePaymentMethod              *Bool      `xmlrpc:"hide_payment_method,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	JournalId                      *Many2One  `xmlrpc:"journal_id,omitempty"`
	LineIds                        *Relation  `xmlrpc:"line_ids,omitempty"`
	PartnerBankId                  *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                      *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerType                    *Selection `xmlrpc:"partner_type,omitempty"`
	PaymentDate                    *Time      `xmlrpc:"payment_date,omitempty"`
	PaymentDifference              *Float     `xmlrpc:"payment_difference,omitempty"`
	PaymentDifferenceHandling      *Selection `xmlrpc:"payment_difference_handling,omitempty"`
	PaymentMethodCode              *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId                *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PaymentTokenId                 *Many2One  `xmlrpc:"payment_token_id,omitempty"`
	PaymentType                    *Selection `xmlrpc:"payment_type,omitempty"`
	RequirePartnerBankAccount      *Bool      `xmlrpc:"require_partner_bank_account,omitempty"`
	ShowPartnerBankAccount         *Bool      `xmlrpc:"show_partner_bank_account,omitempty"`
	SourceAmount                   *Float     `xmlrpc:"source_amount,omitempty"`
	SourceAmountCurrency           *Float     `xmlrpc:"source_amount_currency,omitempty"`
	SourceCurrencyId               *Many2One  `xmlrpc:"source_currency_id,omitempty"`
	SuitablePaymentTokenPartnerIds *Relation  `xmlrpc:"suitable_payment_token_partner_ids,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
	WriteoffAccountId              *Many2One  `xmlrpc:"writeoff_account_id,omitempty"`
	WriteoffLabel                  *String    `xmlrpc:"writeoff_label,omitempty"`
}

// AccountPaymentRegisters represents array of account.payment.register model.
type AccountPaymentRegisters []AccountPaymentRegister

// AccountPaymentRegisterModel is the odoo model name.
const AccountPaymentRegisterModel = "account.payment.register"

// Many2One convert AccountPaymentRegister to *Many2One.
func (apr *AccountPaymentRegister) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegister(apr *AccountPaymentRegister) (int64, error) {
	return c.Create(AccountPaymentRegisterModel, apr)
}

// UpdateAccountPaymentRegister updates an existing account.payment.register record.
func (c *Client) UpdateAccountPaymentRegister(apr *AccountPaymentRegister) error {
	return c.UpdateAccountPaymentRegisters([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPaymentRegisters updates existing account.payment.register records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPaymentRegisters(ids []int64, apr *AccountPaymentRegister) error {
	return c.Update(AccountPaymentRegisterModel, ids, apr)
}

// DeleteAccountPaymentRegister deletes an existing account.payment.register record.
func (c *Client) DeleteAccountPaymentRegister(id int64) error {
	return c.DeleteAccountPaymentRegisters([]int64{id})
}

// DeleteAccountPaymentRegisters deletes existing account.payment.register records.
func (c *Client) DeleteAccountPaymentRegisters(ids []int64) error {
	return c.Delete(AccountPaymentRegisterModel, ids)
}

// GetAccountPaymentRegister gets account.payment.register existing record.
func (c *Client) GetAccountPaymentRegister(id int64) (*AccountPaymentRegister, error) {
	aprs, err := c.GetAccountPaymentRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	if aprs != nil && len(*aprs) > 0 {
		return &((*aprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment.register not found", id)
}

// GetAccountPaymentRegisters gets account.payment.register existing records.
func (c *Client) GetAccountPaymentRegisters(ids []int64) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.Read(AccountPaymentRegisterModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegister finds account.payment.register record by querying it with criteria.
func (c *Client) FindAccountPaymentRegister(criteria *Criteria) (*AccountPaymentRegister, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	if aprs != nil && len(*aprs) > 0 {
		return &((*aprs)[0]), nil
	}
	return nil, fmt.Errorf("account.payment.register was not found")
}

// FindAccountPaymentRegisters finds account.payment.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisters(criteria *Criteria, options *Options) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentRegisterId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.payment.register was not found")
}
