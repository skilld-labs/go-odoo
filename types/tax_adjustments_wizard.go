package types

import (
	"time"
)

type TaxAdjustmentsWizard struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Amount            float64   `xmlrpc:"amount"`
	CompanyCurrencyId Many2One  `xmlrpc:"company_currency_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CreditAccountId   Many2One  `xmlrpc:"credit_account_id"`
	Date              time.Time `xmlrpc:"date"`
	DebitAccountId    Many2One  `xmlrpc:"debit_account_id"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	JournalId         Many2One  `xmlrpc:"journal_id"`
	Reason            string    `xmlrpc:"reason"`
	TaxId             Many2One  `xmlrpc:"tax_id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type TaxAdjustmentsWizardNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Amount            interface{} `xmlrpc:"amount"`
	CompanyCurrencyId interface{} `xmlrpc:"company_currency_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CreditAccountId   interface{} `xmlrpc:"credit_account_id"`
	Date              interface{} `xmlrpc:"date"`
	DebitAccountId    interface{} `xmlrpc:"debit_account_id"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	JournalId         interface{} `xmlrpc:"journal_id"`
	Reason            interface{} `xmlrpc:"reason"`
	TaxId             interface{} `xmlrpc:"tax_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var TaxAdjustmentsWizardModel string = "tax.adjustments.wizard"

type TaxAdjustmentsWizards []TaxAdjustmentsWizard

type TaxAdjustmentsWizardsNil []TaxAdjustmentsWizardNil

func (s *TaxAdjustmentsWizard) NilableType_() interface{} {
	return &TaxAdjustmentsWizardNil{}
}

func (ns *TaxAdjustmentsWizardNil) Type_() interface{} {
	s := &TaxAdjustmentsWizard{}
	return load(ns, s)
}

func (s *TaxAdjustmentsWizards) NilableType_() interface{} {
	return &TaxAdjustmentsWizardsNil{}
}

func (ns *TaxAdjustmentsWizardsNil) Type_() interface{} {
	s := &TaxAdjustmentsWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*TaxAdjustmentsWizard))
	}
	return s
}
