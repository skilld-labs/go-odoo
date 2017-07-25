package types

import (
	"time"
)

type ReportStockLinesDate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	Date        time.Time `xmlrpc:"date"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	MoveDate    time.Time `xmlrpc:"move_date"`
	ProductId   Many2One  `xmlrpc:"product_id"`
}

type ReportStockLinesDateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	Date        interface{} `xmlrpc:"date"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	MoveDate    interface{} `xmlrpc:"move_date"`
	ProductId   interface{} `xmlrpc:"product_id"`
}

var ReportStockLinesDateModel string = "report.stock.lines.date"

type ReportStockLinesDates []ReportStockLinesDate

type ReportStockLinesDatesNil []ReportStockLinesDateNil

func (s *ReportStockLinesDate) NilableType_() interface{} {
	return &ReportStockLinesDateNil{}
}

func (ns *ReportStockLinesDateNil) Type_() interface{} {
	s := &ReportStockLinesDate{}
	return load(ns, s)
}

func (s *ReportStockLinesDates) NilableType_() interface{} {
	return &ReportStockLinesDatesNil{}
}

func (ns *ReportStockLinesDatesNil) Type_() interface{} {
	s := &ReportStockLinesDates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportStockLinesDate))
	}
	return s
}
