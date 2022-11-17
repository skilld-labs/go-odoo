package odoo

import (
	"fmt"
)

// AccountPayment represents account.payment model.
type AccountPayment struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	Amount                    *Float     `xmlrpc:"amount,omitempty"`
	Communication             *String    `xmlrpc:"communication,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	DestinationAccountId      *Many2One  `xmlrpc:"destination_account_id,omitempty"`
	DestinationJournalId      *Many2One  `xmlrpc:"destination_journal_id,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	HasInvoices               *Bool      `xmlrpc:"has_invoices,omitempty"`
	HidePaymentMethod         *Bool      `xmlrpc:"hide_payment_method,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	InvoiceIds                *Relation  `xmlrpc:"invoice_ids,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	MessageChannelIds         *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost           *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread             *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter      *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveLineIds               *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveName                  *String    `xmlrpc:"move_name,omitempty"`
	MoveReconciled            *Bool      `xmlrpc:"move_reconciled,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PartnerId                 *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerType               *Selection `xmlrpc:"partner_type,omitempty"`
	PaymentDate               *Time      `xmlrpc:"payment_date,omitempty"`
	PaymentDifference         *Float     `xmlrpc:"payment_difference,omitempty"`
	PaymentDifferenceHandling *Selection `xmlrpc:"payment_difference_handling,omitempty"`
	PaymentMethodCode         *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId           *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PaymentReference          *String    `xmlrpc:"payment_reference,omitempty"`
	PaymentTokenId            *Many2One  `xmlrpc:"payment_token_id,omitempty"`
	PaymentTransactionId      *Many2One  `xmlrpc:"payment_transaction_id,omitempty"`
	PaymentType               *Selection `xmlrpc:"payment_type,omitempty"`
	State                     *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
	WriteoffAccountId         *Many2One  `xmlrpc:"writeoff_account_id,omitempty"`
	WriteoffLabel             *String    `xmlrpc:"writeoff_label,omitempty"`
}

// AccountPayments represents array of account.payment model.
type AccountPayments []AccountPayment

// AccountPaymentModel is the odoo model name.
const AccountPaymentModel = "account.payment"

// Many2One convert AccountPayment to *Many2One.
func (ap *AccountPayment) Many2One() *Many2One {
	return NewMany2One(ap.Id.Get(), "")
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayment(ap *AccountPayment) (int64, error) {
	return c.Create(AccountPaymentModel, ap)
}

// UpdateAccountPayment updates an existing account.payment record.
func (c *Client) UpdateAccountPayment(ap *AccountPayment) error {
	return c.UpdateAccountPayments([]int64{ap.Id.Get()}, ap)
}

// UpdateAccountPayments updates existing account.payment records.
// All records (represented by ids) will be updated by ap values.
func (c *Client) UpdateAccountPayments(ids []int64, ap *AccountPayment) error {
	return c.Update(AccountPaymentModel, ids, ap)
}

// DeleteAccountPayment deletes an existing account.payment record.
func (c *Client) DeleteAccountPayment(id int64) error {
	return c.DeleteAccountPayments([]int64{id})
}

// DeleteAccountPayments deletes existing account.payment records.
func (c *Client) DeleteAccountPayments(ids []int64) error {
	return c.Delete(AccountPaymentModel, ids)
}

// GetAccountPayment gets account.payment existing record.
func (c *Client) GetAccountPayment(id int64) (*AccountPayment, error) {
	aps, err := c.GetAccountPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment not found", id)
}

// GetAccountPayments gets account.payment existing records.
func (c *Client) GetAccountPayments(ids []int64) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.Read(AccountPaymentModel, ids, nil, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPayment finds account.payment record by querying it with criteria.
func (c *Client) FindAccountPayment(criteria *Criteria) (*AccountPayment, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, NewOptions().Limit(1), aps); err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("no account.payment was found with criteria %v", criteria)
}

// FindAccountPayments finds account.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPayments(criteria *Criteria, options *Options) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, options, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.payment was found with criteria %v and options %v", criteria, options)
}
