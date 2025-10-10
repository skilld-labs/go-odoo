package odoo

// AccountBankStatementLine represents account.bank.statement.line model.
type AccountBankStatementLine struct {
	AbnormalAmountWarning                 *String     `xmlrpc:"abnormal_amount_warning,omitempty"`
	AbnormalDateWarning                   *String     `xmlrpc:"abnormal_date_warning,omitempty"`
	AccessToken                           *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String     `xmlrpc:"access_warning,omitempty"`
	AccountNumber                         *String     `xmlrpc:"account_number,omitempty"`
	ActivityCalendarEventId               *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AlwaysTaxExigible                     *Bool       `xmlrpc:"always_tax_exigible,omitempty"`
	Amount                                *Float      `xmlrpc:"amount,omitempty"`
	AmountCurrency                        *Float      `xmlrpc:"amount_currency,omitempty"`
	AmountPaid                            *Float      `xmlrpc:"amount_paid,omitempty"`
	AmountResidual                        *Float      `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float      `xmlrpc:"amount_residual_signed,omitempty"`
	AmountTax                             *Float      `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float      `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float      `xmlrpc:"amount_total,omitempty"`
	AmountTotalInCurrencySigned           *Float      `xmlrpc:"amount_total_in_currency_signed,omitempty"`
	AmountTotalSigned                     *Float      `xmlrpc:"amount_total_signed,omitempty"`
	AmountTotalWords                      *String     `xmlrpc:"amount_total_words,omitempty"`
	AmountUntaxed                         *Float      `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedInCurrencySigned         *Float      `xmlrpc:"amount_untaxed_in_currency_signed,omitempty"`
	AmountUntaxedSigned                   *Float      `xmlrpc:"amount_untaxed_signed,omitempty"`
	AttachmentIds                         *Relation   `xmlrpc:"attachment_ids,omitempty"`
	AuditTrailMessageIds                  *Relation   `xmlrpc:"audit_trail_message_ids,omitempty"`
	AuthorizedTransactionIds              *Relation   `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Selection  `xmlrpc:"auto_post,omitempty"`
	AutoPostOriginId                      *Many2One   `xmlrpc:"auto_post_origin_id,omitempty"`
	AutoPostUntil                         *Time       `xmlrpc:"auto_post_until,omitempty"`
	BankPartnerId                         *Many2One   `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One   `xmlrpc:"campaign_id,omitempty"`
	Checked                               *Bool       `xmlrpc:"checked,omitempty"`
	CommercialPartnerId                   *Many2One   `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyPriceInclude                   *Selection  `xmlrpc:"company_price_include,omitempty"`
	CountryCode                           *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time       `xmlrpc:"date,omitempty"`
	DeliveryDate                          *Time       `xmlrpc:"delivery_date,omitempty"`
	DirectionSign                         *Int        `xmlrpc:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool       `xmlrpc:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool       `xmlrpc:"display_qr_code,omitempty"`
	DuplicatedRefIds                      *Relation   `xmlrpc:"duplicated_ref_ids,omitempty"`
	FiscalPositionId                      *Many2One   `xmlrpc:"fiscal_position_id,omitempty"`
	ForeignCurrencyId                     *Many2One   `xmlrpc:"foreign_currency_id,omitempty"`
	HasMessage                            *Bool       `xmlrpc:"has_message,omitempty"`
	HasReconciledEntries                  *Bool       `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool       `xmlrpc:"hide_post_button,omitempty"`
	HighestName                           *String     `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int        `xmlrpc:"id,omitempty"`
	InalterableHash                       *String     `xmlrpc:"inalterable_hash,omitempty"`
	IncotermLocation                      *String     `xmlrpc:"incoterm_location,omitempty"`
	InternalIndex                         *String     `xmlrpc:"internal_index,omitempty"`
	InvoiceCashRoundingId                 *Many2One   `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceCurrencyRate                   *Float      `xmlrpc:"invoice_currency_rate,omitempty"`
	InvoiceDate                           *Time       `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time       `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String     `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasOutstanding                 *Bool       `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One   `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation   `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String     `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String     `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String     `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One   `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String     `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoicePdfReportFile                  *String     `xmlrpc:"invoice_pdf_report_file,omitempty"`
	InvoicePdfReportId                    *Many2One   `xmlrpc:"invoice_pdf_report_id,omitempty"`
	InvoiceSourceEmail                    *String     `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One   `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One   `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsBeingSent                           *Bool       `xmlrpc:"is_being_sent,omitempty"`
	IsManuallyModified                    *Bool       `xmlrpc:"is_manually_modified,omitempty"`
	IsMoveSent                            *Bool       `xmlrpc:"is_move_sent,omitempty"`
	IsPurchaseMatched                     *Bool       `xmlrpc:"is_purchase_matched,omitempty"`
	IsReconciled                          *Bool       `xmlrpc:"is_reconciled,omitempty"`
	IsStorno                              *Bool       `xmlrpc:"is_storno,omitempty"`
	JournalGroupId                        *Many2One   `xmlrpc:"journal_group_id,omitempty"`
	JournalId                             *Many2One   `xmlrpc:"journal_id,omitempty"`
	LineIds                               *Relation   `xmlrpc:"line_ids,omitempty"`
	MachineInvoice                        *Bool       `xmlrpc:"machine_invoice,omitempty"`
	MachineInvoiceTitle                   *String     `xmlrpc:"machine_invoice_title,omitempty"`
	MachinePurchaseOrder                  *String     `xmlrpc:"machine_purchase_order,omitempty"`
	MadeSequenceGap                       *Bool       `xmlrpc:"made_sequence_gap,omitempty"`
	MatchedPaymentIds                     *Relation   `xmlrpc:"matched_payment_ids,omitempty"`
	MediumId                              *Many2One   `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One   `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MoveId                                *Many2One   `xmlrpc:"move_id,omitempty"`
	MoveSentValues                        *Selection  `xmlrpc:"move_sent_values,omitempty"`
	MoveType                              *Selection  `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String     `xmlrpc:"name,omitempty"`
	NamePlaceholder                       *String     `xmlrpc:"name_placeholder,omitempty"`
	Narration                             *String     `xmlrpc:"narration,omitempty"`
	NeedCancelRequest                     *Bool       `xmlrpc:"need_cancel_request,omitempty"`
	NeededTerms                           *String     `xmlrpc:"needed_terms,omitempty"`
	NeededTermsDirty                      *Bool       `xmlrpc:"needed_terms_dirty,omitempty"`
	NextPaymentDate                       *Time       `xmlrpc:"next_payment_date,omitempty"`
	OriginPaymentId                       *Many2One   `xmlrpc:"origin_payment_id,omitempty"`
	PartnerBankId                         *Many2One   `xmlrpc:"partner_bank_id,omitempty"`
	PartnerCredit                         *Float      `xmlrpc:"partner_credit,omitempty"`
	PartnerCreditWarning                  *String     `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                             *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerName                           *String     `xmlrpc:"partner_name,omitempty"`
	PartnerShippingId                     *Many2One   `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentCount                          *Int        `xmlrpc:"payment_count,omitempty"`
	PaymentIds                            *Relation   `xmlrpc:"payment_ids,omitempty"`
	PaymentRef                            *String     `xmlrpc:"payment_ref,omitempty"`
	PaymentReference                      *String     `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection  `xmlrpc:"payment_state,omitempty"`
	PaymentTermDetails                    *String     `xmlrpc:"payment_term_details,omitempty"`
	PostedBefore                          *Bool       `xmlrpc:"posted_before,omitempty"`
	PreferredPaymentMethodLineId          *Many2One   `xmlrpc:"preferred_payment_method_line_id,omitempty"`
	PurchaseId                            *Many2One   `xmlrpc:"purchase_id,omitempty"`
	PurchaseOrderCount                    *Int        `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseOrderName                     *String     `xmlrpc:"purchase_order_name,omitempty"`
	PurchaseVendorBillId                  *Many2One   `xmlrpc:"purchase_vendor_bill_id,omitempty"`
	QrCodeMethod                          *Selection  `xmlrpc:"qr_code_method,omitempty"`
	QuantityTotal                         *Float      `xmlrpc:"quantity_total,omitempty"`
	QuickEditMode                         *Bool       `xmlrpc:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float      `xmlrpc:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     interface{} `xmlrpc:"quick_encoding_vals,omitempty"`
	RatingIds                             *Relation   `xmlrpc:"rating_ids,omitempty"`
	Ref                                   *String     `xmlrpc:"ref,omitempty"`
	RestrictModeHashTable                 *Bool       `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveIds                       *Relation   `xmlrpc:"reversal_move_ids,omitempty"`
	ReversedEntryId                       *Many2One   `xmlrpc:"reversed_entry_id,omitempty"`
	RunningBalance                        *Float      `xmlrpc:"running_balance,omitempty"`
	SaleOrderCount                        *Int        `xmlrpc:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int        `xmlrpc:"secure_sequence_number,omitempty"`
	Secured                               *Bool       `xmlrpc:"secured,omitempty"`
	SendingData                           interface{} `xmlrpc:"sending_data,omitempty"`
	Sequence                              *Int        `xmlrpc:"sequence,omitempty"`
	SequenceNumber                        *Int        `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String     `xmlrpc:"sequence_prefix,omitempty"`
	ShowDeliveryDate                      *Bool       `xmlrpc:"show_delivery_date,omitempty"`
	ShowDiscountDetails                   *Bool       `xmlrpc:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool       `xmlrpc:"show_name_warning,omitempty"`
	ShowPaymentTermDetails                *Bool       `xmlrpc:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool       `xmlrpc:"show_reset_to_draft_button,omitempty"`
	ShowUpdateFpos                        *Bool       `xmlrpc:"show_update_fpos,omitempty"`
	SourceId                              *Many2One   `xmlrpc:"source_id,omitempty"`
	State                                 *Selection  `xmlrpc:"state,omitempty"`
	StatementBalanceEndReal               *Float      `xmlrpc:"statement_balance_end_real,omitempty"`
	StatementComplete                     *Bool       `xmlrpc:"statement_complete,omitempty"`
	StatementId                           *Many2One   `xmlrpc:"statement_id,omitempty"`
	StatementLineId                       *Many2One   `xmlrpc:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation   `xmlrpc:"statement_line_ids,omitempty"`
	StatementName                         *String     `xmlrpc:"statement_name,omitempty"`
	StatementValid                        *Bool       `xmlrpc:"statement_valid,omitempty"`
	StatusInPayment                       *Selection  `xmlrpc:"status_in_payment,omitempty"`
	StockMoveId                           *Many2One   `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	SuitableJournalIds                    *Relation   `xmlrpc:"suitable_journal_ids,omitempty"`
	TaxCalculationRoundingMethod          *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation   `xmlrpc:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One   `xmlrpc:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One   `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxCountryCode                        *String     `xmlrpc:"tax_country_code,omitempty"`
	TaxCountryId                          *Many2One   `xmlrpc:"tax_country_id,omitempty"`
	TaxLockDateMessage                    *String     `xmlrpc:"tax_lock_date_message,omitempty"`
	TaxTotals                             *String     `xmlrpc:"tax_totals,omitempty"`
	TaxesLegalNotes                       *String     `xmlrpc:"taxes_legal_notes,omitempty"`
	TeamId                                *Many2One   `xmlrpc:"team_id,omitempty"`
	TimesheetCount                        *Int        `xmlrpc:"timesheet_count,omitempty"`
	TimesheetEncodeUomId                  *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                          *Relation   `xmlrpc:"timesheet_ids,omitempty"`
	TimesheetTotalDuration                *Int        `xmlrpc:"timesheet_total_duration,omitempty"`
	TransactionCount                      *Int        `xmlrpc:"transaction_count,omitempty"`
	TransactionDetails                    interface{} `xmlrpc:"transaction_details,omitempty"`
	TransactionIds                        *Relation   `xmlrpc:"transaction_ids,omitempty"`
	TransactionType                       *String     `xmlrpc:"transaction_type,omitempty"`
	TypeName                              *String     `xmlrpc:"type_name,omitempty"`
	UblCiiXmlFile                         *String     `xmlrpc:"ubl_cii_xml_file,omitempty"`
	UblCiiXmlId                           *Many2One   `xmlrpc:"ubl_cii_xml_id,omitempty"`
	UserId                                *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatementLines represents array of account.bank.statement.line model.
type AccountBankStatementLines []AccountBankStatementLine

// AccountBankStatementLineModel is the odoo model name.
const AccountBankStatementLineModel = "account.bank.statement.line"

// Many2One convert AccountBankStatementLine to *Many2One.
func (absl *AccountBankStatementLine) Many2One() *Many2One {
	return NewMany2One(absl.Id.Get(), "")
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLine(absl *AccountBankStatementLine) (int64, error) {
	ids, err := c.CreateAccountBankStatementLines([]*AccountBankStatementLine{absl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLines(absls []*AccountBankStatementLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range absls {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementLineModel, vv, nil)
}

// UpdateAccountBankStatementLine updates an existing account.bank.statement.line record.
func (c *Client) UpdateAccountBankStatementLine(absl *AccountBankStatementLine) error {
	return c.UpdateAccountBankStatementLines([]int64{absl.Id.Get()}, absl)
}

// UpdateAccountBankStatementLines updates existing account.bank.statement.line records.
// All records (represented by ids) will be updated by absl values.
func (c *Client) UpdateAccountBankStatementLines(ids []int64, absl *AccountBankStatementLine) error {
	return c.Update(AccountBankStatementLineModel, ids, absl, nil)
}

// DeleteAccountBankStatementLine deletes an existing account.bank.statement.line record.
func (c *Client) DeleteAccountBankStatementLine(id int64) error {
	return c.DeleteAccountBankStatementLines([]int64{id})
}

// DeleteAccountBankStatementLines deletes existing account.bank.statement.line records.
func (c *Client) DeleteAccountBankStatementLines(ids []int64) error {
	return c.Delete(AccountBankStatementLineModel, ids)
}

// GetAccountBankStatementLine gets account.bank.statement.line existing record.
func (c *Client) GetAccountBankStatementLine(id int64) (*AccountBankStatementLine, error) {
	absls, err := c.GetAccountBankStatementLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*absls)[0]), nil
}

// GetAccountBankStatementLines gets account.bank.statement.line existing records.
func (c *Client) GetAccountBankStatementLines(ids []int64) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.Read(AccountBankStatementLineModel, ids, nil, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLine finds account.bank.statement.line record by querying it with criteria.
func (c *Client) FindAccountBankStatementLine(criteria *Criteria) (*AccountBankStatementLine, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, NewOptions().Limit(1), absls); err != nil {
		return nil, err
	}
	return &((*absls)[0]), nil
}

// FindAccountBankStatementLines finds account.bank.statement.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLines(criteria *Criteria, options *Options) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, options, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankStatementLineModel, criteria, options)
}

// FindAccountBankStatementLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
