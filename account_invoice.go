package odoo

// AccountInvoice represents account.invoice model.
type AccountInvoice struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                    *String    `xmlrpc:"access_token,omitempty"`
	AccountId                      *Many2One  `xmlrpc:"account_id,omitempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountTax                      *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal                    *Float     `xmlrpc:"amount_total,omitempty"`
	AmountTotalCompanySigned       *Float     `xmlrpc:"amount_total_company_signed,omitempty"`
	AmountTotalSigned              *Float     `xmlrpc:"amount_total_signed,omitempty"`
	AmountUntaxed                  *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned            *Float     `xmlrpc:"amount_untaxed_signed,omitempty"`
	CampaignId                     *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CashRoundingId                 *Many2One  `xmlrpc:"cash_rounding_id,omitempty"`
	Comment                        *String    `xmlrpc:"comment,omitempty"`
	CommercialPartnerId            *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId              *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                           *Time      `xmlrpc:"date,omitempty"`
	DateDue                        *Time      `xmlrpc:"date_due,omitempty"`
	DateInvoice                    *Time      `xmlrpc:"date_invoice,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId               *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasOutstanding                 *Bool      `xmlrpc:"has_outstanding,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	IncotermsId                    *Many2One  `xmlrpc:"incoterms_id,omitempty"`
	InvoiceLineIds                 *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	JournalId                      *Many2One  `xmlrpc:"journal_id,omitempty"`
	MachineInvoice                 *Bool      `xmlrpc:"machine_invoice,omitempty"`
	MachineInvoiceTitle            *String    `xmlrpc:"machine_invoice_title,omitempty"`
	MachinePurchaseOrder           *String    `xmlrpc:"machine_purchase_order,omitempty"`
	MediumId                       *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageChannelIds              *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost                *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                  *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter           *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveId                         *Many2One  `xmlrpc:"move_id,omitempty"`
	MoveName                       *String    `xmlrpc:"move_name,omitempty"`
	Name                           *String    `xmlrpc:"name,omitempty"`
	Number                         *String    `xmlrpc:"number,omitempty"`
	Origin                         *String    `xmlrpc:"origin,omitempty"`
	OutstandingCreditsDebitsWidget *String    `xmlrpc:"outstanding_credits_debits_widget,omitempty"`
	PartnerBankId                  *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                      *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId              *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentIds                     *Relation  `xmlrpc:"payment_ids,omitempty"`
	PaymentMoveLineIds             *Relation  `xmlrpc:"payment_move_line_ids,omitempty"`
	PaymentTermId                  *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PaymentsWidget                 *String    `xmlrpc:"payments_widget,omitempty"`
	PortalUrl                      *String    `xmlrpc:"portal_url,omitempty"`
	PurchaseId                     *Many2One  `xmlrpc:"purchase_id,omitempty"`
	QuantityTotal                  *Float     `xmlrpc:"quantity_total,omitempty"`
	Reconciled                     *Bool      `xmlrpc:"reconciled,omitempty"`
	Reference                      *String    `xmlrpc:"reference,omitempty"`
	ReferenceType                  *Selection `xmlrpc:"reference_type,omitempty"`
	RefundInvoiceId                *Many2One  `xmlrpc:"refund_invoice_id,omitempty"`
	RefundInvoiceIds               *Relation  `xmlrpc:"refund_invoice_ids,omitempty"`
	Residual                       *Float     `xmlrpc:"residual,omitempty"`
	ResidualCompanySigned          *Float     `xmlrpc:"residual_company_signed,omitempty"`
	ResidualSigned                 *Float     `xmlrpc:"residual_signed,omitempty"`
	Sent                           *Bool      `xmlrpc:"sent,omitempty"`
	SequenceNumberNext             *String    `xmlrpc:"sequence_number_next,omitempty"`
	SequenceNumberNextPrefix       *String    `xmlrpc:"sequence_number_next_prefix,omitempty"`
	SourceId                       *Many2One  `xmlrpc:"source_id,omitempty"`
	State                          *Selection `xmlrpc:"state,omitempty"`
	TaxLineIds                     *Relation  `xmlrpc:"tax_line_ids,omitempty"`
	TeamId                         *Many2One  `xmlrpc:"team_id,omitempty"`
	TimesheetCount                 *Int       `xmlrpc:"timesheet_count,omitempty"`
	TimesheetIds                   *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	Type                           *Selection `xmlrpc:"type,omitempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(AccountInvoiceModel, vv, nil)
}

// UpdateAccountInvoice updates an existing account.invoice record.
func (c *Client) UpdateAccountInvoice(ai *AccountInvoice) error {
	return c.UpdateAccountInvoices([]int64{ai.Id.Get()}, ai)
}

// UpdateAccountInvoices updates existing account.invoice records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAccountInvoices(ids []int64, ai *AccountInvoice) error {
	return c.Update(AccountInvoiceModel, ids, ai, nil)
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
	return &((*ais)[0]), nil
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
	return &((*ais)[0]), nil
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
	return c.Search(AccountInvoiceModel, criteria, options)
}

// FindAccountInvoiceId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
