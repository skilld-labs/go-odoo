package types

import (
	"time"
)

type StockInventoryLine struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	InventoryId    Many2One  `xmlrpc:"inventory_id"`
	LocationId     Many2One  `xmlrpc:"location_id"`
	LocationName   string    `xmlrpc:"location_name"`
	PackageId      Many2One  `xmlrpc:"package_id"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	ProdLotId      Many2One  `xmlrpc:"prod_lot_id"`
	ProdlotName    string    `xmlrpc:"prodlot_name"`
	ProductCode    string    `xmlrpc:"product_code"`
	ProductId      Many2One  `xmlrpc:"product_id"`
	ProductName    string    `xmlrpc:"product_name"`
	ProductQty     float64   `xmlrpc:"product_qty"`
	ProductUomId   Many2One  `xmlrpc:"product_uom_id"`
	State          string    `xmlrpc:"state"`
	TheoreticalQty float64   `xmlrpc:"theoretical_qty"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type StockInventoryLineNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	InventoryId    interface{} `xmlrpc:"inventory_id"`
	LocationId     interface{} `xmlrpc:"location_id"`
	LocationName   interface{} `xmlrpc:"location_name"`
	PackageId      interface{} `xmlrpc:"package_id"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	ProdLotId      interface{} `xmlrpc:"prod_lot_id"`
	ProdlotName    interface{} `xmlrpc:"prodlot_name"`
	ProductCode    interface{} `xmlrpc:"product_code"`
	ProductId      interface{} `xmlrpc:"product_id"`
	ProductName    interface{} `xmlrpc:"product_name"`
	ProductQty     interface{} `xmlrpc:"product_qty"`
	ProductUomId   interface{} `xmlrpc:"product_uom_id"`
	State          interface{} `xmlrpc:"state"`
	TheoreticalQty interface{} `xmlrpc:"theoretical_qty"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var StockInventoryLineModel string = "stock.inventory.line"

type StockInventoryLines []StockInventoryLine

type StockInventoryLinesNil []StockInventoryLineNil

func (s *StockInventoryLine) NilableType_() interface{} {
	return &StockInventoryLineNil{}
}

func (ns *StockInventoryLineNil) Type_() interface{} {
	s := &StockInventoryLine{}
	return load(ns, s)
}

func (s *StockInventoryLines) NilableType_() interface{} {
	return &StockInventoryLinesNil{}
}

func (ns *StockInventoryLinesNil) Type_() interface{} {
	s := &StockInventoryLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockInventoryLine))
	}
	return s
}
