package types

import (
	"time"
)

type PublisherWarrantyContract struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type PublisherWarrantyContractNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var PublisherWarrantyContractModel string = "publisher_warranty.contract"

type PublisherWarrantyContracts []PublisherWarrantyContract

type PublisherWarrantyContractsNil []PublisherWarrantyContractNil

func (s *PublisherWarrantyContract) NilableType_() interface{} {
	return &PublisherWarrantyContractNil{}
}

func (ns *PublisherWarrantyContractNil) Type_() interface{} {
	s := &PublisherWarrantyContract{}
	return load(ns, s)
}

func (s *PublisherWarrantyContracts) NilableType_() interface{} {
	return &PublisherWarrantyContractsNil{}
}

func (ns *PublisherWarrantyContractsNil) Type_() interface{} {
	s := &PublisherWarrantyContracts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PublisherWarrantyContract))
	}
	return s
}
