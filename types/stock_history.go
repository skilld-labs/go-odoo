package types

import (
	"time"
)

type StockHistory struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	Date              time.Time `xmlrpc:"date"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	InventoryValue    float64   `xmlrpc:"inventory_value"`
	LocationId        Many2One  `xmlrpc:"location_id"`
	MoveId            Many2One  `xmlrpc:"move_id"`
	PriceUnitOnQuant  float64   `xmlrpc:"price_unit_on_quant"`
	ProductCategId    Many2One  `xmlrpc:"product_categ_id"`
	ProductId         Many2One  `xmlrpc:"product_id"`
	ProductTemplateId Many2One  `xmlrpc:"product_template_id"`
	Quantity          float64   `xmlrpc:"quantity"`
	SerialNumber      string    `xmlrpc:"serial_number"`
	Source            string    `xmlrpc:"source"`
}

type StockHistoryNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	Date              interface{} `xmlrpc:"date"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	InventoryValue    interface{} `xmlrpc:"inventory_value"`
	LocationId        interface{} `xmlrpc:"location_id"`
	MoveId            interface{} `xmlrpc:"move_id"`
	PriceUnitOnQuant  interface{} `xmlrpc:"price_unit_on_quant"`
	ProductCategId    interface{} `xmlrpc:"product_categ_id"`
	ProductId         interface{} `xmlrpc:"product_id"`
	ProductTemplateId interface{} `xmlrpc:"product_template_id"`
	Quantity          interface{} `xmlrpc:"quantity"`
	SerialNumber      interface{} `xmlrpc:"serial_number"`
	Source            interface{} `xmlrpc:"source"`
}

var StockHistoryModel string = "stock.history"

type StockHistorys []StockHistory

type StockHistorysNil []StockHistoryNil

func (s *StockHistory) NilableType_() interface{} {
	return &StockHistoryNil{}
}

func (ns *StockHistoryNil) Type_() interface{} {
	s := &StockHistory{}
	return load(ns, s)
}

func (s *StockHistorys) NilableType_() interface{} {
	return &StockHistorysNil{}
}

func (ns *StockHistorysNil) Type_() interface{} {
	s := &StockHistorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockHistory))
	}
	return s
}
