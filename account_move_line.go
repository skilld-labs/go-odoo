package odoo

// AccountMoveLine represents account.move.line model.
type AccountMoveLine struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	AccountId                *Many2One `xmlrpc:"account_id,omitempty"`
	AmountCurrency           *Float    `xmlrpc:"amount_currency,omitempty"`
	AmountResidual           *Float    `xmlrpc:"amount_residual,omitempty"`
	AmountResidualCurrency   *Float    `xmlrpc:"amount_residual_currency,omitempty"`
	AnalyticAccountId        *Many2One `xmlrpc:"analytic_account_id,omitempty"`
	AnalyticLineIds          *Relation `xmlrpc:"analytic_line_ids,omitempty"`
	AnalyticTagIds           *Relation `xmlrpc:"analytic_tag_ids,omitempty"`
	Balance                  *Float    `xmlrpc:"balance,omitempty"`
	BalanceCashBasis         *Float    `xmlrpc:"balance_cash_basis,omitempty"`
	Blocked                  *Bool     `xmlrpc:"blocked,omitempty"`
	CompanyCurrencyId        *Many2One `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	Counterpart              *String   `xmlrpc:"counterpart,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	Credit                   *Float    `xmlrpc:"credit,omitempty"`
	CreditCashBasis          *Float    `xmlrpc:"credit_cash_basis,omitempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omitempty"`
	Date                     *Time     `xmlrpc:"date,omitempty"`
	DateMaturity             *Time     `xmlrpc:"date_maturity,omitempty"`
	Debit                    *Float    `xmlrpc:"debit,omitempty"`
	DebitCashBasis           *Float    `xmlrpc:"debit_cash_basis,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	FullReconcileId          *Many2One `xmlrpc:"full_reconcile_id,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	InvoiceId                *Many2One `xmlrpc:"invoice_id,omitempty"`
	IsUnaffectedEarningsLine *Bool     `xmlrpc:"is_unaffected_earnings_line,omitempty"`
	JournalId                *Many2One `xmlrpc:"journal_id,omitempty"`
	MatchedCreditIds         *Relation `xmlrpc:"matched_credit_ids,omitempty"`
	MatchedDebitIds          *Relation `xmlrpc:"matched_debit_ids,omitempty"`
	MoveId                   *Many2One `xmlrpc:"move_id,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	Narration                *String   `xmlrpc:"narration,omitempty"`
	ParentState              *String   `xmlrpc:"parent_state,omitempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omitempty"`
	PaymentId                *Many2One `xmlrpc:"payment_id,omitempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomId             *Many2One `xmlrpc:"product_uom_id,omitempty"`
	Quantity                 *Float    `xmlrpc:"quantity,omitempty"`
	Reconciled               *Bool     `xmlrpc:"reconciled,omitempty"`
	Ref                      *String   `xmlrpc:"ref,omitempty"`
	StatementId              *Many2One `xmlrpc:"statement_id,omitempty"`
	StatementLineId          *Many2One `xmlrpc:"statement_line_id,omitempty"`
	TaxBaseAmount            *Float    `xmlrpc:"tax_base_amount,omitempty"`
	TaxExigible              *Bool     `xmlrpc:"tax_exigible,omitempty"`
	TaxIds                   *Relation `xmlrpc:"tax_ids,omitempty"`
	TaxLineId                *Many2One `xmlrpc:"tax_line_id,omitempty"`
	UserTypeId               *Many2One `xmlrpc:"user_type_id,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
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
