package types

import (
	"time"
)

type AccountMoveLine struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	AccountId              Many2One  `xmlrpc:"account_id"`
	AmountCurrency         float64   `xmlrpc:"amount_currency"`
	AmountResidual         float64   `xmlrpc:"amount_residual"`
	AmountResidualCurrency float64   `xmlrpc:"amount_residual_currency"`
	AnalyticAccountId      Many2One  `xmlrpc:"analytic_account_id"`
	AnalyticLineIds        []int64   `xmlrpc:"analytic_line_ids"`
	AnalyticTagIds         []int64   `xmlrpc:"analytic_tag_ids"`
	Balance                float64   `xmlrpc:"balance"`
	BalanceCashBasis       float64   `xmlrpc:"balance_cash_basis"`
	Blocked                bool      `xmlrpc:"blocked"`
	CompanyCurrencyId      Many2One  `xmlrpc:"company_currency_id"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	Counterpart            string    `xmlrpc:"counterpart"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	Credit                 float64   `xmlrpc:"credit"`
	CreditCashBasis        float64   `xmlrpc:"credit_cash_basis"`
	CurrencyId             Many2One  `xmlrpc:"currency_id"`
	Date                   time.Time `xmlrpc:"date"`
	DateMaturity           time.Time `xmlrpc:"date_maturity"`
	Debit                  float64   `xmlrpc:"debit"`
	DebitCashBasis         float64   `xmlrpc:"debit_cash_basis"`
	DisplayName            string    `xmlrpc:"display_name"`
	FullReconcileId        Many2One  `xmlrpc:"full_reconcile_id"`
	Id                     int64     `xmlrpc:"id"`
	InvoiceId              Many2One  `xmlrpc:"invoice_id"`
	JournalId              Many2One  `xmlrpc:"journal_id"`
	MatchedCreditIds       []int64   `xmlrpc:"matched_credit_ids"`
	MatchedDebitIds        []int64   `xmlrpc:"matched_debit_ids"`
	MoveId                 Many2One  `xmlrpc:"move_id"`
	Name                   string    `xmlrpc:"name"`
	Narration              string    `xmlrpc:"narration"`
	PartnerId              Many2One  `xmlrpc:"partner_id"`
	PaymentId              Many2One  `xmlrpc:"payment_id"`
	ProductId              Many2One  `xmlrpc:"product_id"`
	ProductUomId           Many2One  `xmlrpc:"product_uom_id"`
	Quantity               float64   `xmlrpc:"quantity"`
	Reconciled             bool      `xmlrpc:"reconciled"`
	Ref                    string    `xmlrpc:"ref"`
	StatementId            Many2One  `xmlrpc:"statement_id"`
	TaxExigible            bool      `xmlrpc:"tax_exigible"`
	TaxIds                 []int64   `xmlrpc:"tax_ids"`
	TaxLineId              Many2One  `xmlrpc:"tax_line_id"`
	UserTypeId             Many2One  `xmlrpc:"user_type_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type AccountMoveLineNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	AccountId              interface{} `xmlrpc:"account_id"`
	AmountCurrency         interface{} `xmlrpc:"amount_currency"`
	AmountResidual         interface{} `xmlrpc:"amount_residual"`
	AmountResidualCurrency interface{} `xmlrpc:"amount_residual_currency"`
	AnalyticAccountId      interface{} `xmlrpc:"analytic_account_id"`
	AnalyticLineIds        interface{} `xmlrpc:"analytic_line_ids"`
	AnalyticTagIds         interface{} `xmlrpc:"analytic_tag_ids"`
	Balance                interface{} `xmlrpc:"balance"`
	BalanceCashBasis       interface{} `xmlrpc:"balance_cash_basis"`
	Blocked                bool        `xmlrpc:"blocked"`
	CompanyCurrencyId      interface{} `xmlrpc:"company_currency_id"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	Counterpart            interface{} `xmlrpc:"counterpart"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	Credit                 interface{} `xmlrpc:"credit"`
	CreditCashBasis        interface{} `xmlrpc:"credit_cash_basis"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	Date                   interface{} `xmlrpc:"date"`
	DateMaturity           interface{} `xmlrpc:"date_maturity"`
	Debit                  interface{} `xmlrpc:"debit"`
	DebitCashBasis         interface{} `xmlrpc:"debit_cash_basis"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	FullReconcileId        interface{} `xmlrpc:"full_reconcile_id"`
	Id                     interface{} `xmlrpc:"id"`
	InvoiceId              interface{} `xmlrpc:"invoice_id"`
	JournalId              interface{} `xmlrpc:"journal_id"`
	MatchedCreditIds       interface{} `xmlrpc:"matched_credit_ids"`
	MatchedDebitIds        interface{} `xmlrpc:"matched_debit_ids"`
	MoveId                 interface{} `xmlrpc:"move_id"`
	Name                   interface{} `xmlrpc:"name"`
	Narration              interface{} `xmlrpc:"narration"`
	PartnerId              interface{} `xmlrpc:"partner_id"`
	PaymentId              interface{} `xmlrpc:"payment_id"`
	ProductId              interface{} `xmlrpc:"product_id"`
	ProductUomId           interface{} `xmlrpc:"product_uom_id"`
	Quantity               interface{} `xmlrpc:"quantity"`
	Reconciled             bool        `xmlrpc:"reconciled"`
	Ref                    interface{} `xmlrpc:"ref"`
	StatementId            interface{} `xmlrpc:"statement_id"`
	TaxExigible            bool        `xmlrpc:"tax_exigible"`
	TaxIds                 interface{} `xmlrpc:"tax_ids"`
	TaxLineId              interface{} `xmlrpc:"tax_line_id"`
	UserTypeId             interface{} `xmlrpc:"user_type_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var AccountMoveLineModel string = "account.move.line"

type AccountMoveLines []AccountMoveLine

type AccountMoveLinesNil []AccountMoveLineNil

func (s *AccountMoveLine) NilableType_() interface{} {
	return &AccountMoveLineNil{}
}

func (ns *AccountMoveLineNil) Type_() interface{} {
	s := &AccountMoveLine{}
	return load(ns, s)
}

func (s *AccountMoveLines) NilableType_() interface{} {
	return &AccountMoveLinesNil{}
}

func (ns *AccountMoveLinesNil) Type_() interface{} {
	s := &AccountMoveLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountMoveLine))
	}
	return s
}
