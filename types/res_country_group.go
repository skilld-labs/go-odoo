package types

import (
	"time"
)

type ResCountryGroup struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	CountryIds   []int64   `xmlrpc:"country_ids"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	Id           int64     `xmlrpc:"id"`
	Name         string    `xmlrpc:"name"`
	PricelistIds []int64   `xmlrpc:"pricelist_ids"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type ResCountryGroupNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	CountryIds   interface{} `xmlrpc:"country_ids"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Id           interface{} `xmlrpc:"id"`
	Name         interface{} `xmlrpc:"name"`
	PricelistIds interface{} `xmlrpc:"pricelist_ids"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var ResCountryGroupModel string = "res.country.group"

type ResCountryGroups []ResCountryGroup

type ResCountryGroupsNil []ResCountryGroupNil

func (s *ResCountryGroup) NilableType_() interface{} {
	return &ResCountryGroupNil{}
}

func (ns *ResCountryGroupNil) Type_() interface{} {
	s := &ResCountryGroup{}
	return load(ns, s)
}

func (s *ResCountryGroups) NilableType_() interface{} {
	return &ResCountryGroupsNil{}
}

func (ns *ResCountryGroupsNil) Type_() interface{} {
	s := &ResCountryGroups{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCountryGroup))
	}
	return s
}
