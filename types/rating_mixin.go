package types

import (
	"time"
)

type RatingMixin struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	RatingCount     int64     `xmlrpc:"rating_count"`
	RatingIds       []int64   `xmlrpc:"rating_ids"`
	RatingLastValue float64   `xmlrpc:"rating_last_value"`
}

type RatingMixinNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	RatingCount     interface{} `xmlrpc:"rating_count"`
	RatingIds       interface{} `xmlrpc:"rating_ids"`
	RatingLastValue interface{} `xmlrpc:"rating_last_value"`
}

var RatingMixinModel string = "rating.mixin"

type RatingMixins []RatingMixin

type RatingMixinsNil []RatingMixinNil

func (s *RatingMixin) NilableType_() interface{} {
	return &RatingMixinNil{}
}

func (ns *RatingMixinNil) Type_() interface{} {
	s := &RatingMixin{}
	return load(ns, s)
}

func (s *RatingMixins) NilableType_() interface{} {
	return &RatingMixinsNil{}
}

func (ns *RatingMixinsNil) Type_() interface{} {
	s := &RatingMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*RatingMixin))
	}
	return s
}
