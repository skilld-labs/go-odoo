package types

import (
	"time"
)

type StockLocation struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	Active                bool      `xmlrpc:"active"`
	Barcode               string    `xmlrpc:"barcode"`
	ChildIds              []int64   `xmlrpc:"child_ids"`
	Comment               string    `xmlrpc:"comment"`
	CompanyId             Many2One  `xmlrpc:"company_id"`
	CompleteName          string    `xmlrpc:"complete_name"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Id                    int64     `xmlrpc:"id"`
	LocationId            Many2One  `xmlrpc:"location_id"`
	Name                  string    `xmlrpc:"name"`
	ParentLeft            int64     `xmlrpc:"parent_left"`
	ParentRight           int64     `xmlrpc:"parent_right"`
	PartnerId             Many2One  `xmlrpc:"partner_id"`
	Posx                  int64     `xmlrpc:"posx"`
	Posy                  int64     `xmlrpc:"posy"`
	Posz                  int64     `xmlrpc:"posz"`
	PutawayStrategyId     Many2One  `xmlrpc:"putaway_strategy_id"`
	RemovalStrategyId     Many2One  `xmlrpc:"removal_strategy_id"`
	ReturnLocation        bool      `xmlrpc:"return_location"`
	ScrapLocation         bool      `xmlrpc:"scrap_location"`
	Usage                 string    `xmlrpc:"usage"`
	ValuationInAccountId  Many2One  `xmlrpc:"valuation_in_account_id"`
	ValuationOutAccountId Many2One  `xmlrpc:"valuation_out_account_id"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type StockLocationNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	Active                bool        `xmlrpc:"active"`
	Barcode               interface{} `xmlrpc:"barcode"`
	ChildIds              interface{} `xmlrpc:"child_ids"`
	Comment               interface{} `xmlrpc:"comment"`
	CompanyId             interface{} `xmlrpc:"company_id"`
	CompleteName          interface{} `xmlrpc:"complete_name"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Id                    interface{} `xmlrpc:"id"`
	LocationId            interface{} `xmlrpc:"location_id"`
	Name                  interface{} `xmlrpc:"name"`
	ParentLeft            interface{} `xmlrpc:"parent_left"`
	ParentRight           interface{} `xmlrpc:"parent_right"`
	PartnerId             interface{} `xmlrpc:"partner_id"`
	Posx                  interface{} `xmlrpc:"posx"`
	Posy                  interface{} `xmlrpc:"posy"`
	Posz                  interface{} `xmlrpc:"posz"`
	PutawayStrategyId     interface{} `xmlrpc:"putaway_strategy_id"`
	RemovalStrategyId     interface{} `xmlrpc:"removal_strategy_id"`
	ReturnLocation        bool        `xmlrpc:"return_location"`
	ScrapLocation         bool        `xmlrpc:"scrap_location"`
	Usage                 interface{} `xmlrpc:"usage"`
	ValuationInAccountId  interface{} `xmlrpc:"valuation_in_account_id"`
	ValuationOutAccountId interface{} `xmlrpc:"valuation_out_account_id"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var StockLocationModel string = "stock.location"

type StockLocations []StockLocation

type StockLocationsNil []StockLocationNil

func (s *StockLocation) NilableType_() interface{} {
	return &StockLocationNil{}
}

func (ns *StockLocationNil) Type_() interface{} {
	s := &StockLocation{}
	return load(ns, s)
}

func (s *StockLocations) NilableType_() interface{} {
	return &StockLocationsNil{}
}

func (ns *StockLocationsNil) Type_() interface{} {
	s := &StockLocations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockLocation))
	}
	return s
}
