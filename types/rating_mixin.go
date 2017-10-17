package types

import (
	"time"
)

type RatingMixin struct {
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	RatingCount        int64     `xmlrpc:"rating_count"`
	RatingIds          []int64   `xmlrpc:"rating_ids"`
	RatingLastFeedback string    `xmlrpc:"rating_last_feedback"`
	RatingLastImage    string    `xmlrpc:"rating_last_image"`
	RatingLastValue    float64   `xmlrpc:"rating_last_value"`
}

type RatingMixinNil struct {
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	RatingCount        interface{} `xmlrpc:"rating_count"`
	RatingIds          interface{} `xmlrpc:"rating_ids"`
	RatingLastFeedback interface{} `xmlrpc:"rating_last_feedback"`
	RatingLastImage    interface{} `xmlrpc:"rating_last_image"`
	RatingLastValue    interface{} `xmlrpc:"rating_last_value"`
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
