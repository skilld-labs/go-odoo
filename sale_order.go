package odoo

// SaleOrder represents sale.order model.
type SaleOrder struct {
	AccessToken                   *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                     *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                 *String     `xmlrpc:"access_warning,omitempty"`
	ActivityCalendarEventId       *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline          *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration   *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon         *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                   *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                 *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary               *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon              *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AmountInvoiced                *Float      `xmlrpc:"amount_invoiced,omitempty"`
	AmountPaid                    *Float      `xmlrpc:"amount_paid,omitempty"`
	AmountTax                     *Float      `xmlrpc:"amount_tax,omitempty"`
	AmountToInvoice               *Float      `xmlrpc:"amount_to_invoice,omitempty"`
	AmountTotal                   *Float      `xmlrpc:"amount_total,omitempty"`
	AmountUndiscounted            *Float      `xmlrpc:"amount_undiscounted,omitempty"`
	AmountUntaxed                 *Float      `xmlrpc:"amount_untaxed,omitempty"`
	AuthorizedTransactionIds      *Relation   `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutosalesConfigId             *Many2One   `xmlrpc:"autosales_config_id,omitempty"`
	AvailableQuotationDocumentIds *Relation   `xmlrpc:"available_quotation_document_ids,omitempty"`
	CampaignId                    *Many2One   `xmlrpc:"campaign_id,omitempty"`
	ClientOrderRef                *String     `xmlrpc:"client_order_ref,omitempty"`
	ClosedTaskCount               *Int        `xmlrpc:"closed_task_count,omitempty"`
	CommitmentDate                *Time       `xmlrpc:"commitment_date,omitempty"`
	CompanyId                     *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyPriceInclude           *Selection  `xmlrpc:"company_price_include,omitempty"`
	CompletedTaskPercentage       *Float      `xmlrpc:"completed_task_percentage,omitempty"`
	ContributionsCreated          *Bool       `xmlrpc:"contributions_created,omitempty"`
	CountryCode                   *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                    *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                    *Many2One   `xmlrpc:"currency_id,omitempty"`
	CurrencyRate                  *Float      `xmlrpc:"currency_rate,omitempty"`
	CustomizablePdfFormFields     interface{} `xmlrpc:"customizable_pdf_form_fields,omitempty"`
	DateOrder                     *Time       `xmlrpc:"date_order,omitempty"`
	DeliveryCount                 *Int        `xmlrpc:"delivery_count,omitempty"`
	DeliveryStatus                *Selection  `xmlrpc:"delivery_status,omitempty"`
	DisplayName                   *String     `xmlrpc:"display_name,omitempty"`
	DuplicatedOrderIds            *Relation   `xmlrpc:"duplicated_order_ids,omitempty"`
	EffectiveDate                 *Time       `xmlrpc:"effective_date,omitempty"`
	ExpectedDate                  *Time       `xmlrpc:"expected_date,omitempty"`
	FiscalPositionId              *Many2One   `xmlrpc:"fiscal_position_id,omitempty"`
	HasActivePricelist            *Bool       `xmlrpc:"has_active_pricelist,omitempty"`
	HasArchivedProducts           *Bool       `xmlrpc:"has_archived_products,omitempty"`
	HasAuthorizedTransactionIds   *Bool       `xmlrpc:"has_authorized_transaction_ids,omitempty"`
	HasMessage                    *Bool       `xmlrpc:"has_message,omitempty"`
	Id                            *Int        `xmlrpc:"id,omitempty"`
	Incoterm                      *Many2One   `xmlrpc:"incoterm,omitempty"`
	IncotermLocation              *String     `xmlrpc:"incoterm_location,omitempty"`
	InvoiceCount                  *Int        `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds                    *Relation   `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus                 *Selection  `xmlrpc:"invoice_status,omitempty"`
	IsExpired                     *Bool       `xmlrpc:"is_expired,omitempty"`
	IsPdfQuoteBuilderAvailable    *Bool       `xmlrpc:"is_pdf_quote_builder_available,omitempty"`
	IsProductMilestone            *Bool       `xmlrpc:"is_product_milestone,omitempty"`
	JournalId                     *Many2One   `xmlrpc:"journal_id,omitempty"`
	JsonPopover                   *String     `xmlrpc:"json_popover,omitempty"`
	LateAvailability              *Bool       `xmlrpc:"late_availability,omitempty"`
	Locked                        *Bool       `xmlrpc:"locked,omitempty"`
	MediumId                      *Many2One   `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount        *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds            *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError               *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower             *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction             *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneCount                *Int        `xmlrpc:"milestone_count,omitempty"`
	MyActivityDateDeadline        *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                          *String     `xmlrpc:"name,omitempty"`
	Note                          *String     `xmlrpc:"note,omitempty"`
	OpportunityId                 *Many2One   `xmlrpc:"opportunity_id,omitempty"`
	OrderLine                     *Relation   `xmlrpc:"order_line,omitempty"`
	Origin                        *String     `xmlrpc:"origin,omitempty"`
	PartnerCreditWarning          *String     `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                     *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerInvoiceId              *Many2One   `xmlrpc:"partner_invoice_id,omitempty"`
	PartnerShippingId             *Many2One   `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentTermId                 *Many2One   `xmlrpc:"payment_term_id,omitempty"`
	PendingEmailTemplateId        *Many2One   `xmlrpc:"pending_email_template_id,omitempty"`
	PickingIds                    *Relation   `xmlrpc:"picking_ids,omitempty"`
	PickingPolicy                 *Selection  `xmlrpc:"picking_policy,omitempty"`
	PreferredPaymentMethodLineId  *Many2One   `xmlrpc:"preferred_payment_method_line_id,omitempty"`
	PrepaymentPercent             *Float      `xmlrpc:"prepayment_percent,omitempty"`
	PricelistId                   *Many2One   `xmlrpc:"pricelist_id,omitempty"`
	ProjectAccountId              *Many2One   `xmlrpc:"project_account_id,omitempty"`
	ProjectCount                  *Int        `xmlrpc:"project_count,omitempty"`
	ProjectId                     *Many2One   `xmlrpc:"project_id,omitempty"`
	ProjectIds                    *Relation   `xmlrpc:"project_ids,omitempty"`
	PurchaseOrderCount            *Int        `xmlrpc:"purchase_order_count,omitempty"`
	QuotationDocumentIds          *Relation   `xmlrpc:"quotation_document_ids,omitempty"`
	RatingIds                     *Relation   `xmlrpc:"rating_ids,omitempty"`
	Reference                     *String     `xmlrpc:"reference,omitempty"`
	RequirePayment                *Bool       `xmlrpc:"require_payment,omitempty"`
	RequireSignature              *Bool       `xmlrpc:"require_signature,omitempty"`
	SaleOrderTemplateId           *Many2One   `xmlrpc:"sale_order_template_id,omitempty"`
	SaleWarningText               *String     `xmlrpc:"sale_warning_text,omitempty"`
	ShowCreateProjectButton       *Bool       `xmlrpc:"show_create_project_button,omitempty"`
	ShowHoursRecordedButton       *Bool       `xmlrpc:"show_hours_recorded_button,omitempty"`
	ShowJsonPopover               *Bool       `xmlrpc:"show_json_popover,omitempty"`
	ShowProjectButton             *Bool       `xmlrpc:"show_project_button,omitempty"`
	ShowUpdateFpos                *Bool       `xmlrpc:"show_update_fpos,omitempty"`
	ShowUpdatePricelist           *Bool       `xmlrpc:"show_update_pricelist,omitempty"`
	Signature                     *String     `xmlrpc:"signature,omitempty"`
	SignedBy                      *String     `xmlrpc:"signed_by,omitempty"`
	SignedOn                      *Time       `xmlrpc:"signed_on,omitempty"`
	SourceId                      *Many2One   `xmlrpc:"source_id,omitempty"`
	State                         *Selection  `xmlrpc:"state,omitempty"`
	StockReferenceIds             *Relation   `xmlrpc:"stock_reference_ids,omitempty"`
	TagIds                        *Relation   `xmlrpc:"tag_ids,omitempty"`
	TasksCount                    *Int        `xmlrpc:"tasks_count,omitempty"`
	TasksIds                      *Relation   `xmlrpc:"tasks_ids,omitempty"`
	TaxCalculationRoundingMethod  *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCountryId                  *Many2One   `xmlrpc:"tax_country_id,omitempty"`
	TaxTotals                     *String     `xmlrpc:"tax_totals,omitempty"`
	TeamId                        *Many2One   `xmlrpc:"team_id,omitempty"`
	TermsType                     *Selection  `xmlrpc:"terms_type,omitempty"`
	TimesheetCount                *Float      `xmlrpc:"timesheet_count,omitempty"`
	TimesheetEncodeUomId          *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetTotalDuration        *Int        `xmlrpc:"timesheet_total_duration,omitempty"`
	TransactionIds                *Relation   `xmlrpc:"transaction_ids,omitempty"`
	TypeName                      *String     `xmlrpc:"type_name,omitempty"`
	UserId                        *Many2One   `xmlrpc:"user_id,omitempty"`
	ValidityDate                  *Time       `xmlrpc:"validity_date,omitempty"`
	VisibleProject                *Bool       `xmlrpc:"visible_project,omitempty"`
	WarehouseId                   *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds             *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                     *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One   `xmlrpc:"write_uid,omitempty"`
	XTitle                        *String     `xmlrpc:"x_title,omitempty"`
	XTitleReadOnly                *String     `xmlrpc:"x_title_read_only,omitempty"`
}

// SaleOrders represents array of sale.order model.
type SaleOrders []SaleOrder

// SaleOrderModel is the odoo model name.
const SaleOrderModel = "sale.order"

// Many2One convert SaleOrder to *Many2One.
func (so *SaleOrder) Many2One() *Many2One {
	return NewMany2One(so.Id.Get(), "")
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrder(so *SaleOrder) (int64, error) {
	ids, err := c.CreateSaleOrders([]*SaleOrder{so})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv, nil)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so, nil)
}

// DeleteSaleOrder deletes an existing sale.order record.
func (c *Client) DeleteSaleOrder(id int64) error {
	return c.DeleteSaleOrders([]int64{id})
}

// DeleteSaleOrders deletes existing sale.order records.
func (c *Client) DeleteSaleOrders(ids []int64) error {
	return c.Delete(SaleOrderModel, ids)
}

// GetSaleOrder gets sale.order existing record.
func (c *Client) GetSaleOrder(id int64) (*SaleOrder, error) {
	sos, err := c.GetSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sos)[0]), nil
}

// GetSaleOrders gets sale.order existing records.
func (c *Client) GetSaleOrders(ids []int64) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.Read(SaleOrderModel, ids, nil, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrder finds sale.order record by querying it with criteria.
func (c *Client) FindSaleOrder(criteria *Criteria) (*SaleOrder, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, NewOptions().Limit(1), sos); err != nil {
		return nil, err
	}
	return &((*sos)[0]), nil
}

// FindSaleOrders finds sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrders(criteria *Criteria, options *Options) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, options, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderModel, criteria, options)
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
