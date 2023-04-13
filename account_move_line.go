package odoo

import (
	"fmt"
)

// AccountMoveLine represents account.move.line model.
type AccountMoveLine struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId                *Many2One `xmlrpc:"account_id,omptempty"`
	AmountCurrency           *Float    `xmlrpc:"amount_currency,omptempty"`
	AmountResidual           *Float    `xmlrpc:"amount_residual,omptempty"`
	AmountResidualCurrency   *Float    `xmlrpc:"amount_residual_currency,omptempty"`
	AnalyticAccountId        *Many2One `xmlrpc:"analytic_account_id,omptempty"`
	AnalyticLineIds          *Relation `xmlrpc:"analytic_line_ids,omptempty"`
	AnalyticTagIds           *Relation `xmlrpc:"analytic_tag_ids,omptempty"`
	Balance                  *Float    `xmlrpc:"balance,omptempty"`
	BalanceCashBasis         *Float    `xmlrpc:"balance_cash_basis,omptempty"`
	Blocked                  *Bool     `xmlrpc:"blocked,omptempty"`
	CompanyCurrencyId        *Many2One `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omptempty"`
	Counterpart              *String   `xmlrpc:"counterpart,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	Credit                   *Float    `xmlrpc:"credit,omptempty"`
	CreditCashBasis          *Float    `xmlrpc:"credit_cash_basis,omptempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omptempty"`
	Date                     *Time     `xmlrpc:"date,omptempty"`
	DateMaturity             *Time     `xmlrpc:"date_maturity,omptempty"`
	Debit                    *Float    `xmlrpc:"debit,omptempty"`
	DebitCashBasis           *Float    `xmlrpc:"debit_cash_basis,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	FullReconcileId          *Many2One `xmlrpc:"full_reconcile_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	InvoiceId                *Many2One `xmlrpc:"invoice_id,omptempty"`
	IsUnaffectedEarningsLine *Bool     `xmlrpc:"is_unaffected_earnings_line,omptempty"`
	JournalId                *Many2One `xmlrpc:"journal_id,omptempty"`
	MatchedCreditIds         *Relation `xmlrpc:"matched_credit_ids,omptempty"`
	MatchedDebitIds          *Relation `xmlrpc:"matched_debit_ids,omptempty"`
	MoveId                   *Many2One `xmlrpc:"move_id,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	Narration                *String   `xmlrpc:"narration,omptempty"`
	ParentState              *String   `xmlrpc:"parent_state,omptempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omptempty"`
	PaymentId                *Many2One `xmlrpc:"payment_id,omptempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomId             *Many2One `xmlrpc:"product_uom_id,omptempty"`
	Quantity                 *Float    `xmlrpc:"quantity,omptempty"`
	Reconciled               *Bool     `xmlrpc:"reconciled,omptempty"`
	Ref                      *String   `xmlrpc:"ref,omptempty"`
	StatementId              *Many2One `xmlrpc:"statement_id,omptempty"`
	StatementLineId          *Many2One `xmlrpc:"statement_line_id,omptempty"`
	TaxBaseAmount            *Float    `xmlrpc:"tax_base_amount,omptempty"`
	TaxExigible              *Bool     `xmlrpc:"tax_exigible,omptempty"`
	TaxIds                   *Relation `xmlrpc:"tax_ids,omptempty"`
	TaxLineId                *Many2One `xmlrpc:"tax_line_id,omptempty"`
	UserTypeId               *Many2One `xmlrpc:"user_type_id,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(AccountMoveLineModel, []interface{}{aml})
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
	return c.Create(AccountMoveLineModel, vv)
}

// UpdateAccountMoveLine updates an existing account.move.line record.
func (c *Client) UpdateAccountMoveLine(aml *AccountMoveLine) error {
	return c.UpdateAccountMoveLines([]int64{aml.Id.Get()}, aml)
}

// UpdateAccountMoveLines updates existing account.move.line records.
// All records (represented by ids) will be updated by aml values.
func (c *Client) UpdateAccountMoveLines(ids []int64, aml *AccountMoveLine) error {
	return c.Update(AccountMoveLineModel, ids, aml)
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
	if amls != nil && len(*amls) > 0 {
		return &((*amls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move.line not found", id)
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
	if amls != nil && len(*amls) > 0 {
		return &((*amls)[0]), nil
	}
	return nil, fmt.Errorf("account.move.line was not found with criteria %v", criteria)
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
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveLineId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move.line was not found with criteria %v and options %v", criteria, options)
}
