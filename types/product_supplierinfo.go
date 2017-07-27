package types

import (
	"time"
)

type ProductSupplierinfo struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	CompanyId     Many2One  `xmlrpc:"company_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	CurrencyId    Many2One  `xmlrpc:"currency_id"`
	DateEnd       time.Time `xmlrpc:"date_end"`
	DateStart     time.Time `xmlrpc:"date_start"`
	Delay         int64     `xmlrpc:"delay"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	MinQty        float64   `xmlrpc:"min_qty"`
	Name          Many2One  `xmlrpc:"name"`
	Price         float64   `xmlrpc:"price"`
	ProductCode   string    `xmlrpc:"product_code"`
	ProductId     Many2One  `xmlrpc:"product_id"`
	ProductName   string    `xmlrpc:"product_name"`
	ProductTmplId Many2One  `xmlrpc:"product_tmpl_id"`
	ProductUom    Many2One  `xmlrpc:"product_uom"`
	Sequence      int64     `xmlrpc:"sequence"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type ProductSupplierinfoNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	CompanyId     interface{} `xmlrpc:"company_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	CurrencyId    interface{} `xmlrpc:"currency_id"`
	DateEnd       interface{} `xmlrpc:"date_end"`
	DateStart     interface{} `xmlrpc:"date_start"`
	Delay         interface{} `xmlrpc:"delay"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	MinQty        interface{} `xmlrpc:"min_qty"`
	Name          interface{} `xmlrpc:"name"`
	Price         interface{} `xmlrpc:"price"`
	ProductCode   interface{} `xmlrpc:"product_code"`
	ProductId     interface{} `xmlrpc:"product_id"`
	ProductName   interface{} `xmlrpc:"product_name"`
	ProductTmplId interface{} `xmlrpc:"product_tmpl_id"`
	ProductUom    interface{} `xmlrpc:"product_uom"`
	Sequence      interface{} `xmlrpc:"sequence"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var ProductSupplierinfoModel string = "product.supplierinfo"

type ProductSupplierinfos []ProductSupplierinfo

type ProductSupplierinfosNil []ProductSupplierinfoNil

func (s *ProductSupplierinfo) NilableType_() interface{} {
	return &ProductSupplierinfoNil{}
}

func (ns *ProductSupplierinfoNil) Type_() interface{} {
	s := &ProductSupplierinfo{}
	return load(ns, s)
}

func (s *ProductSupplierinfos) NilableType_() interface{} {
	return &ProductSupplierinfosNil{}
}

func (ns *ProductSupplierinfosNil) Type_() interface{} {
	s := &ProductSupplierinfos{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductSupplierinfo))
	}
	return s
}
