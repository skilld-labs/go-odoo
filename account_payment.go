package odoo

import (
	"fmt"
)

// AccountPayment represents account.payment model.
type AccountPayment struct {
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
	Amount                                *Float     `xmlrpc:"amount,omitempty"`
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
	AvailablePaymentMethodIds             *Relation  `xmlrpc:"available_payment_method_ids,omitempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CheckAmountInWords                    *String    `xmlrpc:"check_amount_in_words,omitempty"`
	CheckManualSequencing                 *Bool      `xmlrpc:"check_manual_sequencing,omitempty"`
	CheckNumber                           *String    `xmlrpc:"check_number,omitempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time      `xmlrpc:"date,omitempty"`
	DestinationAccountId                  *Many2One  `xmlrpc:"destination_account_id,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omitempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omitempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omitempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omitempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omitempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omitempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePaymentMethod                     *Bool      `xmlrpc:"hide_payment_method,omitempty"`
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
	IsInternalTransfer                    *Bool      `xmlrpc:"is_internal_transfer,omitempty"`
	IsMatched                             *Bool      `xmlrpc:"is_matched,omitempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omitempty"`
	IsReconciled                          *Bool      `xmlrpc:"is_reconciled,omitempty"`
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
	MoveId                                *Many2One  `xmlrpc:"move_id,omitempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	Narration                             *String    `xmlrpc:"narration,omitempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PartnerType                           *Selection `xmlrpc:"partner_type,omitempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omitempty"`
	PaymentMethodCode                     *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId                       *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omitempty"`
	PaymentTokenId                        *Many2One  `xmlrpc:"payment_token_id,omitempty"`
	PaymentTransactionId                  *Many2One  `xmlrpc:"payment_transaction_id,omitempty"`
	PaymentType                           *Selection `xmlrpc:"payment_type,omitempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omitempty"`
	PreferredPaymentMethodId              *Many2One  `xmlrpc:"preferred_payment_method_id,omitempty"`
	QrCode                                *String    `xmlrpc:"qr_code,omitempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omitempty"`
	ReconciledBillIds                     *Relation  `xmlrpc:"reconciled_bill_ids,omitempty"`
	ReconciledBillsCount                  *Int       `xmlrpc:"reconciled_bills_count,omitempty"`
	ReconciledInvoiceIds                  *Relation  `xmlrpc:"reconciled_invoice_ids,omitempty"`
	ReconciledInvoicesCount               *Int       `xmlrpc:"reconciled_invoices_count,omitempty"`
	ReconciledStatementIds                *Relation  `xmlrpc:"reconciled_statement_ids,omitempty"`
	ReconciledStatementsCount             *Int       `xmlrpc:"reconciled_statements_count,omitempty"`
	Ref                                   *String    `xmlrpc:"ref,omitempty"`
	RelatedPartnerIds                     *Relation  `xmlrpc:"related_partner_ids,omitempty"`
	RequirePartnerBankAccount             *Bool      `xmlrpc:"require_partner_bank_account,omitempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omitempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omitempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omitempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omitempty"`
	ShowPartnerBankAccount                *Bool      `xmlrpc:"show_partner_bank_account,omitempty"`
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
	return nil, fmt.Errorf("account.payment was not found")
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
	return -1, fmt.Errorf("account.payment was not found")
}
