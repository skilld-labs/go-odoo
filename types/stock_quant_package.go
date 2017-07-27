package types

import (
	"time"
)

type StockQuantPackage struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	AncestorIds      []int64   `xmlrpc:"ancestor_ids"`
	ChildrenIds      []int64   `xmlrpc:"children_ids"`
	ChildrenQuantIds []int64   `xmlrpc:"children_quant_ids"`
	CompanyId        Many2One  `xmlrpc:"company_id"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	LocationId       Many2One  `xmlrpc:"location_id"`
	Name             string    `xmlrpc:"name"`
	OwnerId          Many2One  `xmlrpc:"owner_id"`
	PackagingId      Many2One  `xmlrpc:"packaging_id"`
	ParentId         Many2One  `xmlrpc:"parent_id"`
	ParentLeft       int64     `xmlrpc:"parent_left"`
	ParentRight      int64     `xmlrpc:"parent_right"`
	QuantIds         []int64   `xmlrpc:"quant_ids"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type StockQuantPackageNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	AncestorIds      interface{} `xmlrpc:"ancestor_ids"`
	ChildrenIds      interface{} `xmlrpc:"children_ids"`
	ChildrenQuantIds interface{} `xmlrpc:"children_quant_ids"`
	CompanyId        interface{} `xmlrpc:"company_id"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	LocationId       interface{} `xmlrpc:"location_id"`
	Name             interface{} `xmlrpc:"name"`
	OwnerId          interface{} `xmlrpc:"owner_id"`
	PackagingId      interface{} `xmlrpc:"packaging_id"`
	ParentId         interface{} `xmlrpc:"parent_id"`
	ParentLeft       interface{} `xmlrpc:"parent_left"`
	ParentRight      interface{} `xmlrpc:"parent_right"`
	QuantIds         interface{} `xmlrpc:"quant_ids"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var StockQuantPackageModel string = "stock.quant.package"

type StockQuantPackages []StockQuantPackage

type StockQuantPackagesNil []StockQuantPackageNil

func (s *StockQuantPackage) NilableType_() interface{} {
	return &StockQuantPackageNil{}
}

func (ns *StockQuantPackageNil) Type_() interface{} {
	s := &StockQuantPackage{}
	return load(ns, s)
}

func (s *StockQuantPackages) NilableType_() interface{} {
	return &StockQuantPackagesNil{}
}

func (ns *StockQuantPackagesNil) Type_() interface{} {
	s := &StockQuantPackages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockQuantPackage))
	}
	return s
}
