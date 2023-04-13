package odoo

import (
	"fmt"
)

// AccountInvoice represents account.invoice model.
type AccountInvoice struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                    *String    `xmlrpc:"access_token,omptempty"`
	AccountId                      *Many2One  `xmlrpc:"account_id,omptempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountTax                      *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTotal                    *Float     `xmlrpc:"amount_total,omptempty"`
	AmountTotalCompanySigned       *Float     `xmlrpc:"amount_total_company_signed,omptempty"`
	AmountTotalSigned              *Float     `xmlrpc:"amount_total_signed,omptempty"`
	AmountUntaxed                  *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedSigned            *Float     `xmlrpc:"amount_untaxed_signed,omptempty"`
	CampaignId                     *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CashRoundingId                 *Many2One  `xmlrpc:"cash_rounding_id,omptempty"`
	Comment                        *String    `xmlrpc:"comment,omptempty"`
	CommercialPartnerId            *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId              *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                           *Time      `xmlrpc:"date,omptempty"`
	DateDue                        *Time      `xmlrpc:"date_due,omptempty"`
	DateInvoice                    *Time      `xmlrpc:"date_invoice,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId               *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	HasOutstanding                 *Bool      `xmlrpc:"has_outstanding,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	IncotermsId                    *Many2One  `xmlrpc:"incoterms_id,omptempty"`
	InvoiceLineIds                 *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	JournalId                      *Many2One  `xmlrpc:"journal_id,omptempty"`
	MachineInvoice                 *Bool      `xmlrpc:"machine_invoice,omptempty"`
	MachineInvoiceTitle            *String    `xmlrpc:"machine_invoice_title,omptempty"`
	MachinePurchaseOrder           *String    `xmlrpc:"machine_purchase_order,omptempty"`
	MediumId                       *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageChannelIds              *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost                *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                  *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter           *Int       `xmlrpc:"message_unread_counter,omptempty"`
	MoveId                         *Many2One  `xmlrpc:"move_id,omptempty"`
	MoveName                       *String    `xmlrpc:"move_name,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	Number                         *String    `xmlrpc:"number,omptempty"`
	Origin                         *String    `xmlrpc:"origin,omptempty"`
	OutstandingCreditsDebitsWidget *String    `xmlrpc:"outstanding_credits_debits_widget,omptempty"`
	PartnerBankId                  *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerId                      *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerShippingId              *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentIds                     *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentMoveLineIds             *Relation  `xmlrpc:"payment_move_line_ids,omptempty"`
	PaymentTermId                  *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PaymentsWidget                 *String    `xmlrpc:"payments_widget,omptempty"`
	PortalUrl                      *String    `xmlrpc:"portal_url,omptempty"`
	PurchaseId                     *Many2One  `xmlrpc:"purchase_id,omptempty"`
	QuantityTotal                  *Float     `xmlrpc:"quantity_total,omptempty"`
	Reconciled                     *Bool      `xmlrpc:"reconciled,omptempty"`
	Reference                      *String    `xmlrpc:"reference,omptempty"`
	ReferenceType                  *Selection `xmlrpc:"reference_type,omptempty"`
	RefundInvoiceId                *Many2One  `xmlrpc:"refund_invoice_id,omptempty"`
	RefundInvoiceIds               *Relation  `xmlrpc:"refund_invoice_ids,omptempty"`
	Residual                       *Float     `xmlrpc:"residual,omptempty"`
	ResidualCompanySigned          *Float     `xmlrpc:"residual_company_signed,omptempty"`
	ResidualSigned                 *Float     `xmlrpc:"residual_signed,omptempty"`
	Sent                           *Bool      `xmlrpc:"sent,omptempty"`
	SequenceNumberNext             *String    `xmlrpc:"sequence_number_next,omptempty"`
	SequenceNumberNextPrefix       *String    `xmlrpc:"sequence_number_next_prefix,omptempty"`
	SourceId                       *Many2One  `xmlrpc:"source_id,omptempty"`
	State                          *Selection `xmlrpc:"state,omptempty"`
	TaxLineIds                     *Relation  `xmlrpc:"tax_line_ids,omptempty"`
	TeamId                         *Many2One  `xmlrpc:"team_id,omptempty"`
	TimesheetCount                 *Int       `xmlrpc:"timesheet_count,omptempty"`
	TimesheetIds                   *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	Type                           *Selection `xmlrpc:"type,omptempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountInvoices represents array of account.invoice model.
type AccountInvoices []AccountInvoice

// AccountInvoiceModel is the odoo model name.
const AccountInvoiceModel = "account.invoice"

// Many2One convert AccountInvoice to *Many2One.
func (ai *AccountInvoice) Many2One() *Many2One {
	return NewMany2One(ai.Id.Get(), "")
}

// CreateAccountInvoice creates a new account.invoice model and returns its id.
func (c *Client) CreateAccountInvoice(ai *AccountInvoice) (int64, error) {
	ids, err := c.CreateAccountInvoices([]*AccountInvoice{ai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoice creates a new account.invoice model and returns its id.
func (c *Client) CreateAccountInvoices(ais []*AccountInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range ais {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceModel, vv)
}

// UpdateAccountInvoice updates an existing account.invoice record.
func (c *Client) UpdateAccountInvoice(ai *AccountInvoice) error {
	return c.UpdateAccountInvoices([]int64{ai.Id.Get()}, ai)
}

// UpdateAccountInvoices updates existing account.invoice records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAccountInvoices(ids []int64, ai *AccountInvoice) error {
	return c.Update(AccountInvoiceModel, ids, ai)
}

// DeleteAccountInvoice deletes an existing account.invoice record.
func (c *Client) DeleteAccountInvoice(id int64) error {
	return c.DeleteAccountInvoices([]int64{id})
}

// DeleteAccountInvoices deletes existing account.invoice records.
func (c *Client) DeleteAccountInvoices(ids []int64) error {
	return c.Delete(AccountInvoiceModel, ids)
}

// GetAccountInvoice gets account.invoice existing record.
func (c *Client) GetAccountInvoice(id int64) (*AccountInvoice, error) {
	ais, err := c.GetAccountInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice not found", id)
}

// GetAccountInvoices gets account.invoice existing records.
func (c *Client) GetAccountInvoices(ids []int64) (*AccountInvoices, error) {
	ais := &AccountInvoices{}
	if err := c.Read(AccountInvoiceModel, ids, nil, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountInvoice finds account.invoice record by querying it with criteria.
func (c *Client) FindAccountInvoice(criteria *Criteria) (*AccountInvoice, error) {
	ais := &AccountInvoices{}
	if err := c.SearchRead(AccountInvoiceModel, criteria, NewOptions().Limit(1), ais); err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("account.invoice was not found with criteria %v", criteria)
}

// FindAccountInvoices finds account.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoices(criteria *Criteria, options *Options) (*AccountInvoices, error) {
	ais := &AccountInvoices{}
	if err := c.SearchRead(AccountInvoiceModel, criteria, options, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.invoice was not found with criteria %v and options %v", criteria, options)
}
