package odoo

import (
	"fmt"
)

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omitempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountByGroup                         *String    `xmlrpc:"amount_by_group,omitempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omitempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omitempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omitempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omitempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Bool      `xmlrpc:"auto_post,omitempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time      `xmlrpc:"date,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omitempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omitempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omitempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omitempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omitempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omitempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omitempty"`
	HighestName                           *String    `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omitempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasMatchingSuspenseAmount      *Bool      `xmlrpc:"invoice_has_matching_suspense_amount,omitempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omitempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omitempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omitempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds                     *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                         *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter                  *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	Narration                             *String    `xmlrpc:"narration,omitempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omitempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omitempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omitempty"`
	PreferredPaymentMethodId              *Many2One  `xmlrpc:"preferred_payment_method_id,omitempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omitempty"`
	Ref                                   *String    `xmlrpc:"ref,omitempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omitempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omitempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omitempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omitempty"`
	ShowResetToDraftButton                *Bool      `xmlrpc:"show_reset_to_draft_button,omitempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omitempty"`
	State                                 *Selection `xmlrpc:"state,omitempty"`
	StatementLineId                       *Many2One  `xmlrpc:"statement_line_id,omitempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omitempty"`
	SuitableJournalIds                    *Relation  `xmlrpc:"suitable_journal_ids,omitempty"`
	TaxCashBasisMoveId                    *Many2One  `xmlrpc:"tax_cash_basis_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omitempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omitempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omitempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TypeName                              *String    `xmlrpc:"type_name,omitempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountMoves represents array of account.move model.
type AccountMoves []AccountMove

// AccountMoveModel is the odoo model name.
const AccountMoveModel = "account.move"

// Many2One convert AccountMove to *Many2One.
func (am *AccountMove) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMove(am *AccountMove) (int64, error) {
	return c.Create(AccountMoveModel, am)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by IDs) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am)
}

// DeleteAccountMove deletes an existing account.move record.
func (c *Client) DeleteAccountMove(id int64) error {
	return c.DeleteAccountMoves([]int64{id})
}

// DeleteAccountMoves deletes existing account.move records.
func (c *Client) DeleteAccountMoves(ids []int64) error {
	return c.Delete(AccountMoveModel, ids)
}

// GetAccountMove gets account.move existing record.
func (c *Client) GetAccountMove(id int64) (*AccountMove, error) {
	ams, err := c.GetAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move not found", id)
}

// GetAccountMoves gets account.move existing records.
func (c *Client) GetAccountMoves(ids []int64) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.Read(AccountMoveModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMove finds account.move record by querying it with criteria.
func (c *Client) FindAccountMove(criteria *Criteria) (*AccountMove, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("account.move was not found")
}

// FindAccountMoves finds account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoves(criteria *Criteria, options *Options) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMoveIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move was not found")
}
