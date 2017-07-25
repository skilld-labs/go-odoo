package types

import (
	"time"
)

type StockChangeProductQty struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	DisplayName         string    `xmlrpc:"display_name"`
	Id                  int64     `xmlrpc:"id"`
	LocationId          Many2One  `xmlrpc:"location_id"`
	LotId               Many2One  `xmlrpc:"lot_id"`
	NewQuantity         float64   `xmlrpc:"new_quantity"`
	ProductId           Many2One  `xmlrpc:"product_id"`
	ProductTmplId       Many2One  `xmlrpc:"product_tmpl_id"`
	ProductVariantCount int64     `xmlrpc:"product_variant_count"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
}

type StockChangeProductQtyNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	Id                  interface{} `xmlrpc:"id"`
	LocationId          interface{} `xmlrpc:"location_id"`
	LotId               interface{} `xmlrpc:"lot_id"`
	NewQuantity         interface{} `xmlrpc:"new_quantity"`
	ProductId           interface{} `xmlrpc:"product_id"`
	ProductTmplId       interface{} `xmlrpc:"product_tmpl_id"`
	ProductVariantCount interface{} `xmlrpc:"product_variant_count"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
}

var StockChangeProductQtyModel string = "stock.change.product.qty"

type StockChangeProductQtys []StockChangeProductQty

type StockChangeProductQtysNil []StockChangeProductQtyNil

func (s *StockChangeProductQty) NilableType_() interface{} {
	return &StockChangeProductQtyNil{}
}

func (ns *StockChangeProductQtyNil) Type_() interface{} {
	s := &StockChangeProductQty{}
	return load(ns, s)
}

func (s *StockChangeProductQtys) NilableType_() interface{} {
	return &StockChangeProductQtysNil{}
}

func (ns *StockChangeProductQtysNil) Type_() interface{} {
	s := &StockChangeProductQtys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockChangeProductQty))
	}
	return s
}
