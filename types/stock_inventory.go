package types

import (
	"time"
)

type StockInventory struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	AccountingDate time.Time `xmlrpc:"accounting_date"`
	CategoryId     Many2One  `xmlrpc:"category_id"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	Date           time.Time `xmlrpc:"date"`
	DisplayName    string    `xmlrpc:"display_name"`
	Exhausted      bool      `xmlrpc:"exhausted"`
	Filter         string    `xmlrpc:"filter"`
	Id             int64     `xmlrpc:"id"`
	LineIds        []int64   `xmlrpc:"line_ids"`
	LocationId     Many2One  `xmlrpc:"location_id"`
	LotId          Many2One  `xmlrpc:"lot_id"`
	MoveIds        []int64   `xmlrpc:"move_ids"`
	Name           string    `xmlrpc:"name"`
	PackageId      Many2One  `xmlrpc:"package_id"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	ProductId      Many2One  `xmlrpc:"product_id"`
	State          string    `xmlrpc:"state"`
	TotalQty       float64   `xmlrpc:"total_qty"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type StockInventoryNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	AccountingDate interface{} `xmlrpc:"accounting_date"`
	CategoryId     interface{} `xmlrpc:"category_id"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	Date           interface{} `xmlrpc:"date"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Exhausted      bool        `xmlrpc:"exhausted"`
	Filter         interface{} `xmlrpc:"filter"`
	Id             interface{} `xmlrpc:"id"`
	LineIds        interface{} `xmlrpc:"line_ids"`
	LocationId     interface{} `xmlrpc:"location_id"`
	LotId          interface{} `xmlrpc:"lot_id"`
	MoveIds        interface{} `xmlrpc:"move_ids"`
	Name           interface{} `xmlrpc:"name"`
	PackageId      interface{} `xmlrpc:"package_id"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	ProductId      interface{} `xmlrpc:"product_id"`
	State          interface{} `xmlrpc:"state"`
	TotalQty       interface{} `xmlrpc:"total_qty"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var StockInventoryModel string = "stock.inventory"

type StockInventorys []StockInventory

type StockInventorysNil []StockInventoryNil

func (s *StockInventory) NilableType_() interface{} {
	return &StockInventoryNil{}
}

func (ns *StockInventoryNil) Type_() interface{} {
	s := &StockInventory{}
	return load(ns, s)
}

func (s *StockInventorys) NilableType_() interface{} {
	return &StockInventorysNil{}
}

func (ns *StockInventorysNil) Type_() interface{} {
	s := &StockInventorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockInventory))
	}
	return s
}
