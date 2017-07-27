package types

import (
	"time"
)

type ProductPutaway struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	FixedLocationIds []int64   `xmlrpc:"fixed_location_ids"`
	Id               int64     `xmlrpc:"id"`
	Method           string    `xmlrpc:"method"`
	Name             string    `xmlrpc:"name"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type ProductPutawayNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	FixedLocationIds interface{} `xmlrpc:"fixed_location_ids"`
	Id               interface{} `xmlrpc:"id"`
	Method           interface{} `xmlrpc:"method"`
	Name             interface{} `xmlrpc:"name"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var ProductPutawayModel string = "product.putaway"

type ProductPutaways []ProductPutaway

type ProductPutawaysNil []ProductPutawayNil

func (s *ProductPutaway) NilableType_() interface{} {
	return &ProductPutawayNil{}
}

func (ns *ProductPutawayNil) Type_() interface{} {
	s := &ProductPutaway{}
	return load(ns, s)
}

func (s *ProductPutaways) NilableType_() interface{} {
	return &ProductPutawaysNil{}
}

func (ns *ProductPutawaysNil) Type_() interface{} {
	s := &ProductPutaways{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPutaway))
	}
	return s
}
