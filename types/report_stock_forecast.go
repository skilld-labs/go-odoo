package types

import (
	"time"
)

type ReportStockForecast struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	CumulativeQuantity float64   `xmlrpc:"cumulative_quantity"`
	Date               time.Time `xmlrpc:"date"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	ProductId          Many2One  `xmlrpc:"product_id"`
	ProductTmplId      Many2One  `xmlrpc:"product_tmpl_id"`
	Quantity           float64   `xmlrpc:"quantity"`
}

type ReportStockForecastNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	CumulativeQuantity interface{} `xmlrpc:"cumulative_quantity"`
	Date               interface{} `xmlrpc:"date"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	ProductId          interface{} `xmlrpc:"product_id"`
	ProductTmplId      interface{} `xmlrpc:"product_tmpl_id"`
	Quantity           interface{} `xmlrpc:"quantity"`
}

var ReportStockForecastModel string = "report.stock.forecast"

type ReportStockForecasts []ReportStockForecast

type ReportStockForecastsNil []ReportStockForecastNil

func (s *ReportStockForecast) NilableType_() interface{} {
	return &ReportStockForecastNil{}
}

func (ns *ReportStockForecastNil) Type_() interface{} {
	s := &ReportStockForecast{}
	return load(ns, s)
}

func (s *ReportStockForecasts) NilableType_() interface{} {
	return &ReportStockForecastsNil{}
}

func (ns *ReportStockForecastsNil) Type_() interface{} {
	s := &ReportStockForecasts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportStockForecast))
	}
	return s
}
