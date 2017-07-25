package types

import (
	"time"
)

type AccountFiscalPosition struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	AccountIds     []int64   `xmlrpc:"account_ids"`
	Active         bool      `xmlrpc:"active"`
	AutoApply      bool      `xmlrpc:"auto_apply"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CountryGroupId Many2One  `xmlrpc:"country_group_id"`
	CountryId      Many2One  `xmlrpc:"country_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	Name           string    `xmlrpc:"name"`
	Note           string    `xmlrpc:"note"`
	Sequence       int64     `xmlrpc:"sequence"`
	StateIds       []int64   `xmlrpc:"state_ids"`
	StatesCount    int64     `xmlrpc:"states_count"`
	TaxIds         []int64   `xmlrpc:"tax_ids"`
	VatRequired    bool      `xmlrpc:"vat_required"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
	ZipFrom        int64     `xmlrpc:"zip_from"`
	ZipTo          int64     `xmlrpc:"zip_to"`
}

type AccountFiscalPositionNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	AccountIds     interface{} `xmlrpc:"account_ids"`
	Active         bool        `xmlrpc:"active"`
	AutoApply      bool        `xmlrpc:"auto_apply"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CountryGroupId interface{} `xmlrpc:"country_group_id"`
	CountryId      interface{} `xmlrpc:"country_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	Name           interface{} `xmlrpc:"name"`
	Note           interface{} `xmlrpc:"note"`
	Sequence       interface{} `xmlrpc:"sequence"`
	StateIds       interface{} `xmlrpc:"state_ids"`
	StatesCount    interface{} `xmlrpc:"states_count"`
	TaxIds         interface{} `xmlrpc:"tax_ids"`
	VatRequired    bool        `xmlrpc:"vat_required"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
	ZipFrom        interface{} `xmlrpc:"zip_from"`
	ZipTo          interface{} `xmlrpc:"zip_to"`
}

var AccountFiscalPositionModel string = "account.fiscal.position"

type AccountFiscalPositions []AccountFiscalPosition

type AccountFiscalPositionsNil []AccountFiscalPositionNil

func (s *AccountFiscalPosition) NilableType_() interface{} {
	return &AccountFiscalPositionNil{}
}

func (ns *AccountFiscalPositionNil) Type_() interface{} {
	s := &AccountFiscalPosition{}
	return load(ns, s)
}

func (s *AccountFiscalPositions) NilableType_() interface{} {
	return &AccountFiscalPositionsNil{}
}

func (ns *AccountFiscalPositionsNil) Type_() interface{} {
	s := &AccountFiscalPositions{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPosition))
	}
	return s
}
