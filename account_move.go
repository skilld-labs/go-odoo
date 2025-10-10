package odoo

// AccountMove represents account.move model.
type AccountMove struct {
	AbnormalAmountWarning                 *String     `xmlrpc:"abnormal_amount_warning,omitempty"`
	AbnormalDateWarning                   *String     `xmlrpc:"abnormal_date_warning,omitempty"`
	AccessToken                           *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String     `xmlrpc:"access_warning,omitempty"`
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
	AssetDepreciatedValue                 *Float      `xmlrpc:"asset_depreciated_value,omitempty"`
	AssetDepreciationBeginningDate        *Time       `xmlrpc:"asset_depreciation_beginning_date,omitempty"`
	AssetId                               *Many2One   `xmlrpc:"asset_id,omitempty"`
	AssetIdDisplayName                    *String     `xmlrpc:"asset_id_display_name,omitempty"`
	AssetIds                              *Relation   `xmlrpc:"asset_ids,omitempty"`
	AssetMoveType                         *Selection  `xmlrpc:"asset_move_type,omitempty"`
	AssetNumberDays                       *Int        `xmlrpc:"asset_number_days,omitempty"`
	AssetRemainingValue                   *Float      `xmlrpc:"asset_remaining_value,omitempty"`
	AssetValueChange                      *Bool       `xmlrpc:"asset_value_change,omitempty"`
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
	CountAsset                            *Int        `xmlrpc:"count_asset,omitempty"`
	CountryCode                           *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time       `xmlrpc:"date,omitempty"`
	DeferredEntryType                     *Selection  `xmlrpc:"deferred_entry_type,omitempty"`
	DeferredMoveIds                       *Relation   `xmlrpc:"deferred_move_ids,omitempty"`
	DeferredOriginalMoveIds               *Relation   `xmlrpc:"deferred_original_move_ids,omitempty"`
	DeliveryDate                          *Time       `xmlrpc:"delivery_date,omitempty"`
	DepreciationValue                     *Float      `xmlrpc:"depreciation_value,omitempty"`
	DirectionSign                         *Int        `xmlrpc:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool       `xmlrpc:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool       `xmlrpc:"display_qr_code,omitempty"`
	DraftAssetExists                      *Bool       `xmlrpc:"draft_asset_exists,omitempty"`
	DuplicatedRefIds                      *Relation   `xmlrpc:"duplicated_ref_ids,omitempty"`
	EdiBlockingLevel                      *Selection  `xmlrpc:"edi_blocking_level,omitempty"`
	EdiDocumentIds                        *Relation   `xmlrpc:"edi_document_ids,omitempty"`
	EdiErrorCount                         *Int        `xmlrpc:"edi_error_count,omitempty"`
	EdiErrorMessage                       *String     `xmlrpc:"edi_error_message,omitempty"`
	EdiShowAbandonCancelButton            *Bool       `xmlrpc:"edi_show_abandon_cancel_button,omitempty"`
	EdiShowCancelButton                   *Bool       `xmlrpc:"edi_show_cancel_button,omitempty"`
	EdiShowForceCancelButton              *Bool       `xmlrpc:"edi_show_force_cancel_button,omitempty"`
	EdiState                              *Selection  `xmlrpc:"edi_state,omitempty"`
	EdiWebServicesToProcess               *String     `xmlrpc:"edi_web_services_to_process,omitempty"`
	ExpectedCurrencyRate                  *Float      `xmlrpc:"expected_currency_rate,omitempty"`
	ExtractAttachmentId                   *Many2One   `xmlrpc:"extract_attachment_id,omitempty"`
	ExtractCanShowBanners                 *Bool       `xmlrpc:"extract_can_show_banners,omitempty"`
	ExtractCanShowSendButton              *Bool       `xmlrpc:"extract_can_show_send_button,omitempty"`
	ExtractDetectedLayout                 *Int        `xmlrpc:"extract_detected_layout,omitempty"`
	ExtractDocumentUuid                   *String     `xmlrpc:"extract_document_uuid,omitempty"`
	ExtractErrorMessage                   *String     `xmlrpc:"extract_error_message,omitempty"`
	ExtractPartnerName                    *String     `xmlrpc:"extract_partner_name,omitempty"`
	ExtractPrefillData                    interface{} `xmlrpc:"extract_prefill_data,omitempty"`
	ExtractState                          *Selection  `xmlrpc:"extract_state,omitempty"`
	ExtractStateProcessed                 *Bool       `xmlrpc:"extract_state_processed,omitempty"`
	ExtractStatus                         *String     `xmlrpc:"extract_status,omitempty"`
	ExtractWordIds                        *Relation   `xmlrpc:"extract_word_ids,omitempty"`
	FiscalPositionId                      *Many2One   `xmlrpc:"fiscal_position_id,omitempty"`
	GeneratingLoanLineId                  *Many2One   `xmlrpc:"generating_loan_line_id,omitempty"`
	HasDocuments                          *Bool       `xmlrpc:"has_documents,omitempty"`
	HasMessage                            *Bool       `xmlrpc:"has_message,omitempty"`
	HasReconciledEntries                  *Bool       `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool       `xmlrpc:"hide_post_button,omitempty"`
	HighestName                           *String     `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int        `xmlrpc:"id,omitempty"`
	InalterableHash                       *String     `xmlrpc:"inalterable_hash,omitempty"`
	IncotermLocation                      *String     `xmlrpc:"incoterm_location,omitempty"`
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
	IsInExtractableState                  *Bool       `xmlrpc:"is_in_extractable_state,omitempty"`
	IsLoanPaymentMove                     *Bool       `xmlrpc:"is_loan_payment_move,omitempty"`
	IsManuallyModified                    *Bool       `xmlrpc:"is_manually_modified,omitempty"`
	IsMoveSent                            *Bool       `xmlrpc:"is_move_sent,omitempty"`
	IsPurchaseMatched                     *Bool       `xmlrpc:"is_purchase_matched,omitempty"`
	IsStorno                              *Bool       `xmlrpc:"is_storno,omitempty"`
	JournalGroupId                        *Many2One   `xmlrpc:"journal_group_id,omitempty"`
	JournalId                             *Many2One   `xmlrpc:"journal_id,omitempty"`
	L10NFrIsCompanyFrench                 *Bool       `xmlrpc:"l10n_fr_is_company_french,omitempty"`
	LineIds                               *Relation   `xmlrpc:"line_ids,omitempty"`
	LoanId                                *Many2One   `xmlrpc:"loan_id,omitempty"`
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
	PartnerShippingId                     *Many2One   `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentCount                          *Int        `xmlrpc:"payment_count,omitempty"`
	PaymentIds                            *Relation   `xmlrpc:"payment_ids,omitempty"`
	PaymentReference                      *String     `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection  `xmlrpc:"payment_state,omitempty"`
	PaymentStateBeforeSwitch              *String     `xmlrpc:"payment_state_before_switch,omitempty"`
	PaymentTermDetails                    *String     `xmlrpc:"payment_term_details,omitempty"`
	PosOrderIds                           *Relation   `xmlrpc:"pos_order_ids,omitempty"`
	PosPaymentIds                         *Relation   `xmlrpc:"pos_payment_ids,omitempty"`
	PosRefundedInvoiceIds                 *Relation   `xmlrpc:"pos_refunded_invoice_ids,omitempty"`
	PosSessionIds                         *Relation   `xmlrpc:"pos_session_ids,omitempty"`
	PostedBefore                          *Bool       `xmlrpc:"posted_before,omitempty"`
	PreferredPaymentMethodLineId          *Many2One   `xmlrpc:"preferred_payment_method_line_id,omitempty"`
	PurchaseId                            *Many2One   `xmlrpc:"purchase_id,omitempty"`
	PurchaseOrderCount                    *Int        `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseOrderName                     *String     `xmlrpc:"purchase_order_name,omitempty"`
	PurchaseVendorBillId                  *Many2One   `xmlrpc:"purchase_vendor_bill_id,omitempty"`
	QrCodeMethod                          *Selection  `xmlrpc:"qr_code_method,omitempty"`
	QuickEditMode                         *Bool       `xmlrpc:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float      `xmlrpc:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     interface{} `xmlrpc:"quick_encoding_vals,omitempty"`
	RatingIds                             *Relation   `xmlrpc:"rating_ids,omitempty"`
	Ref                                   *String     `xmlrpc:"ref,omitempty"`
	RestrictModeHashTable                 *Bool       `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveIds                       *Relation   `xmlrpc:"reversal_move_ids,omitempty"`
	ReversedEntryId                       *Many2One   `xmlrpc:"reversed_entry_id,omitempty"`
	ReversedPosOrderId                    *Many2One   `xmlrpc:"reversed_pos_order_id,omitempty"`
	SaleOrderCount                        *Int        `xmlrpc:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int        `xmlrpc:"secure_sequence_number,omitempty"`
	Secured                               *Bool       `xmlrpc:"secured,omitempty"`
	SendingData                           interface{} `xmlrpc:"sending_data,omitempty"`
	SequenceNumber                        *Int        `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String     `xmlrpc:"sequence_prefix,omitempty"`
	ShowDeliveryDate                      *Bool       `xmlrpc:"show_delivery_date,omitempty"`
	ShowDiscountDetails                   *Bool       `xmlrpc:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool       `xmlrpc:"show_name_warning,omitempty"`
	ShowPaymentTermDetails                *Bool       `xmlrpc:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool       `xmlrpc:"show_reset_to_draft_button,omitempty"`
	ShowSignatureArea                     *Bool       `xmlrpc:"show_signature_area,omitempty"`
	ShowUpdateFpos                        *Bool       `xmlrpc:"show_update_fpos,omitempty"`
	Signature                             *String     `xmlrpc:"signature,omitempty"`
	SigningUser                           *Many2One   `xmlrpc:"signing_user,omitempty"`
	SourceId                              *Many2One   `xmlrpc:"source_id,omitempty"`
	State                                 *Selection  `xmlrpc:"state,omitempty"`
	StatementId                           *Many2One   `xmlrpc:"statement_id,omitempty"`
	StatementLineId                       *Many2One   `xmlrpc:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation   `xmlrpc:"statement_line_ids,omitempty"`
	StatusInPayment                       *Selection  `xmlrpc:"status_in_payment,omitempty"`
	StockMoveId                           *Many2One   `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	SuitableJournalIds                    *Relation   `xmlrpc:"suitable_journal_ids,omitempty"`
	SuspenseStatementLineId               *Many2One   `xmlrpc:"suspense_statement_line_id,omitempty"`
	TaxCalculationRoundingMethod          *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation   `xmlrpc:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One   `xmlrpc:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One   `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxClosingAlert                       *Bool       `xmlrpc:"tax_closing_alert,omitempty"`
	TaxClosingReportId                    *Many2One   `xmlrpc:"tax_closing_report_id,omitempty"`
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
	TransactionIds                        *Relation   `xmlrpc:"transaction_ids,omitempty"`
	TransferModelId                       *Many2One   `xmlrpc:"transfer_model_id,omitempty"`
	TypeName                              *String     `xmlrpc:"type_name,omitempty"`
	UblCiiXmlFile                         *String     `xmlrpc:"ubl_cii_xml_file,omitempty"`
	UblCiiXmlId                           *Many2One   `xmlrpc:"ubl_cii_xml_id,omitempty"`
	UserId                                *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteId                             *Many2One   `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                     *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omitempty"`
	XStudioDateDeDbut                     *Time       `xmlrpc:"x_studio_date_de_dbut,omitempty"`
	XStudioDateDeFin                      *Time       `xmlrpc:"x_studio_date_de_fin,omitempty"`
	XStudioDateDeLevenement               *Time       `xmlrpc:"x_studio_date_de_levenement,omitempty"`
	XStudioEtatFacture                    *Selection  `xmlrpc:"x_studio_etat_facture,omitempty"`
	XStudioFactureYooz                    *String     `xmlrpc:"x_studio_facture_yooz,omitempty"`
	XStudioInformationsComplmentaires     *String     `xmlrpc:"x_studio_informations_complmentaires,omitempty"`
	XStudioJournalType                    *Selection  `xmlrpc:"x_studio_journal_type,omitempty"`
	XStudioNPiceYooz                      *String     `xmlrpc:"x_studio_n_pice_yooz,omitempty"`
	XStudioNfacture                       *String     `xmlrpc:"x_studio_nfacture,omitempty"`
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
	ids, err := c.CreateAccountMoves([]*AccountMove{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMoves(ams []*AccountMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveModel, vv, nil)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am, nil)
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
	return &((*ams)[0]), nil
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
	return &((*ams)[0]), nil
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

// FindAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveModel, criteria, options)
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
