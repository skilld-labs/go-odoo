package odoo

// AccountMoveLine represents account.move.line model.
type AccountMoveLine struct {
	AccountCode                    *String     `xmlrpc:"account_code,omitempty"`
	AccountId                      *Many2One   `xmlrpc:"account_id,omitempty"`
	AccountInternalGroup           *Selection  `xmlrpc:"account_internal_group,omitempty"`
	AccountName                    *String     `xmlrpc:"account_name,omitempty"`
	AccountRootId                  *Many2One   `xmlrpc:"account_root_id,omitempty"`
	AccountType                    *Selection  `xmlrpc:"account_type,omitempty"`
	AmountCurrency                 *Float      `xmlrpc:"amount_currency,omitempty"`
	AmountResidual                 *Float      `xmlrpc:"amount_residual,omitempty"`
	AmountResidualCurrency         *Float      `xmlrpc:"amount_residual_currency,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticLineIds                *Relation   `xmlrpc:"analytic_line_ids,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty"`
	AssetIds                       *Relation   `xmlrpc:"asset_ids,omitempty"`
	Balance                        *Float      `xmlrpc:"balance,omitempty"`
	CashpadTicket                  *Many2One   `xmlrpc:"cashpad_ticket,omitempty"`
	CogsOriginId                   *Many2One   `xmlrpc:"cogs_origin_id,omitempty"`
	CommercialPartnerCountry       *Many2One   `xmlrpc:"commercial_partner_country,omitempty"`
	CompanyCurrencyId              *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty"`
	Credit                         *Float      `xmlrpc:"credit,omitempty"`
	CumulatedBalance               *Float      `xmlrpc:"cumulated_balance,omitempty"`
	CurrencyId                     *Many2One   `xmlrpc:"currency_id,omitempty"`
	CurrencyRate                   *Float      `xmlrpc:"currency_rate,omitempty"`
	Date                           *Time       `xmlrpc:"date,omitempty"`
	DateMaturity                   *Time       `xmlrpc:"date_maturity,omitempty"`
	Debit                          *Float      `xmlrpc:"debit,omitempty"`
	DeferredEndDate                *Time       `xmlrpc:"deferred_end_date,omitempty"`
	DeferredStartDate              *Time       `xmlrpc:"deferred_start_date,omitempty"`
	Discount                       *Float      `xmlrpc:"discount,omitempty"`
	DiscountAllocationDirty        *Bool       `xmlrpc:"discount_allocation_dirty,omitempty"`
	DiscountAllocationKey          *String     `xmlrpc:"discount_allocation_key,omitempty"`
	DiscountAllocationNeeded       *String     `xmlrpc:"discount_allocation_needed,omitempty"`
	DiscountAmountCurrency         *Float      `xmlrpc:"discount_amount_currency,omitempty"`
	DiscountBalance                *Float      `xmlrpc:"discount_balance,omitempty"`
	DiscountDate                   *Time       `xmlrpc:"discount_date,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	DisplayType                    *Selection  `xmlrpc:"display_type,omitempty"`
	DistributionAnalyticAccountIds *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	EpdDirty                       *Bool       `xmlrpc:"epd_dirty,omitempty"`
	EpdKey                         *String     `xmlrpc:"epd_key,omitempty"`
	EpdNeeded                      *String     `xmlrpc:"epd_needed,omitempty"`
	ExcludeBankLines               *Bool       `xmlrpc:"exclude_bank_lines,omitempty"`
	FollowupLineId                 *Many2One   `xmlrpc:"followup_line_id,omitempty"`
	FullReconcileId                *Many2One   `xmlrpc:"full_reconcile_id,omitempty"`
	GroupTaxId                     *Many2One   `xmlrpc:"group_tax_id,omitempty"`
	HasAbnormalDeferredDates       *Bool       `xmlrpc:"has_abnormal_deferred_dates,omitempty"`
	HasDeferredMoves               *Bool       `xmlrpc:"has_deferred_moves,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	InvoiceDate                    *Time       `xmlrpc:"invoice_date,omitempty"`
	InvoiceOrigin                  *String     `xmlrpc:"invoice_origin,omitempty"`
	IsAccountReconcile             *Bool       `xmlrpc:"is_account_reconcile,omitempty"`
	IsDownpayment                  *Bool       `xmlrpc:"is_downpayment,omitempty"`
	IsImported                     *Bool       `xmlrpc:"is_imported,omitempty"`
	IsRefund                       *Bool       `xmlrpc:"is_refund,omitempty"`
	IsSameCurrency                 *Bool       `xmlrpc:"is_same_currency,omitempty"`
	IsStorno                       *Bool       `xmlrpc:"is_storno,omitempty"`
	JournalGroupId                 *Many2One   `xmlrpc:"journal_group_id,omitempty"`
	JournalId                      *Many2One   `xmlrpc:"journal_id,omitempty"`
	MatchedCreditIds               *Relation   `xmlrpc:"matched_credit_ids,omitempty"`
	MatchedDebitIds                *Relation   `xmlrpc:"matched_debit_ids,omitempty"`
	MatchingNumber                 *String     `xmlrpc:"matching_number,omitempty"`
	MoveAttachmentIds              *Relation   `xmlrpc:"move_attachment_ids,omitempty"`
	MoveId                         *Many2One   `xmlrpc:"move_id,omitempty"`
	MoveName                       *String     `xmlrpc:"move_name,omitempty"`
	MoveType                       *Selection  `xmlrpc:"move_type,omitempty"`
	Name                           *String     `xmlrpc:"name,omitempty"`
	NonDeductibleTaxValue          *Float      `xmlrpc:"non_deductible_tax_value,omitempty"`
	ParentState                    *Selection  `xmlrpc:"parent_state,omitempty"`
	PartnerId                      *Many2One   `xmlrpc:"partner_id,omitempty"`
	PaymentDate                    *Time       `xmlrpc:"payment_date,omitempty"`
	PaymentId                      *Many2One   `xmlrpc:"payment_id,omitempty"`
	PriceSubtotal                  *Float      `xmlrpc:"price_subtotal,omitempty"`
	PriceTax                       *Float      `xmlrpc:"price_tax,omitempty"`
	PriceTotal                     *Float      `xmlrpc:"price_total,omitempty"`
	PriceUnit                      *Float      `xmlrpc:"price_unit,omitempty"`
	ProductCategoryId              *Many2One   `xmlrpc:"product_category_id,omitempty"`
	ProductId                      *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductUomCategoryId           *Many2One   `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                   *Many2One   `xmlrpc:"product_uom_id,omitempty"`
	PurchaseLineId                 *Many2One   `xmlrpc:"purchase_line_id,omitempty"`
	PurchaseOrderId                *Many2One   `xmlrpc:"purchase_order_id,omitempty"`
	Quantity                       *Float      `xmlrpc:"quantity,omitempty"`
	ReconcileModelId               *Many2One   `xmlrpc:"reconcile_model_id,omitempty"`
	Reconciled                     *Bool       `xmlrpc:"reconciled,omitempty"`
	Ref                            *String     `xmlrpc:"ref,omitempty"`
	SaleLineIds                    *Relation   `xmlrpc:"sale_line_ids,omitempty"`
	Sequence                       *Int        `xmlrpc:"sequence,omitempty"`
	StatementId                    *Many2One   `xmlrpc:"statement_id,omitempty"`
	StatementLineId                *Many2One   `xmlrpc:"statement_line_id,omitempty"`
	StockValuationLayerIds         *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	TaxBaseAmount                  *Float      `xmlrpc:"tax_base_amount,omitempty"`
	TaxCalculationRoundingMethod   *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxGroupId                     *Many2One   `xmlrpc:"tax_group_id,omitempty"`
	TaxIds                         *Relation   `xmlrpc:"tax_ids,omitempty"`
	TaxLineId                      *Many2One   `xmlrpc:"tax_line_id,omitempty"`
	TaxRepartitionLineId           *Many2One   `xmlrpc:"tax_repartition_line_id,omitempty"`
	TaxTagIds                      *Relation   `xmlrpc:"tax_tag_ids,omitempty"`
	TaxTagInvert                   *Bool       `xmlrpc:"tax_tag_invert,omitempty"`
	TermKey                        *String     `xmlrpc:"term_key,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty"`
	XStudioBalanceAbsolue          *Float      `xmlrpc:"x_studio_balance_absolue,omitempty"`
	XStudioDateEchance             *Time       `xmlrpc:"x_studio_date_echance,omitempty"`
	XStudioNPiceYooz               *String     `xmlrpc:"x_studio_n_pice_yooz,omitempty"`
	XStudioSens1                   *String     `xmlrpc:"x_studio_sens_1,omitempty"`
}

// AccountMoveLines represents array of account.move.line model.
type AccountMoveLines []AccountMoveLine

// AccountMoveLineModel is the odoo model name.
const AccountMoveLineModel = "account.move.line"

// Many2One convert AccountMoveLine to *Many2One.
func (aml *AccountMoveLine) Many2One() *Many2One {
	return NewMany2One(aml.Id.Get(), "")
}

// CreateAccountMoveLine creates a new account.move.line model and returns its id.
func (c *Client) CreateAccountMoveLine(aml *AccountMoveLine) (int64, error) {
	ids, err := c.CreateAccountMoveLines([]*AccountMoveLine{aml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveLine creates a new account.move.line model and returns its id.
func (c *Client) CreateAccountMoveLines(amls []*AccountMoveLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range amls {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveLineModel, vv, nil)
}

// UpdateAccountMoveLine updates an existing account.move.line record.
func (c *Client) UpdateAccountMoveLine(aml *AccountMoveLine) error {
	return c.UpdateAccountMoveLines([]int64{aml.Id.Get()}, aml)
}

// UpdateAccountMoveLines updates existing account.move.line records.
// All records (represented by ids) will be updated by aml values.
func (c *Client) UpdateAccountMoveLines(ids []int64, aml *AccountMoveLine) error {
	return c.Update(AccountMoveLineModel, ids, aml, nil)
}

// DeleteAccountMoveLine deletes an existing account.move.line record.
func (c *Client) DeleteAccountMoveLine(id int64) error {
	return c.DeleteAccountMoveLines([]int64{id})
}

// DeleteAccountMoveLines deletes existing account.move.line records.
func (c *Client) DeleteAccountMoveLines(ids []int64) error {
	return c.Delete(AccountMoveLineModel, ids)
}

// GetAccountMoveLine gets account.move.line existing record.
func (c *Client) GetAccountMoveLine(id int64) (*AccountMoveLine, error) {
	amls, err := c.GetAccountMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// GetAccountMoveLines gets account.move.line existing records.
func (c *Client) GetAccountMoveLines(ids []int64) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.Read(AccountMoveLineModel, ids, nil, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLine finds account.move.line record by querying it with criteria.
func (c *Client) FindAccountMoveLine(criteria *Criteria) (*AccountMoveLine, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, NewOptions().Limit(1), amls); err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// FindAccountMoveLines finds account.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLines(criteria *Criteria, options *Options) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, options, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveLineModel, criteria, options)
}

// FindAccountMoveLineId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
