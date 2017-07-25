package types

import (
	"time"
)

type RatingRating struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	AccessToken    string    `xmlrpc:"access_token"`
	Consumed       bool      `xmlrpc:"consumed"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Feedback       string    `xmlrpc:"feedback"`
	Id             int64     `xmlrpc:"id"`
	MessageId      Many2One  `xmlrpc:"message_id"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	RatedPartnerId Many2One  `xmlrpc:"rated_partner_id"`
	Rating         float64   `xmlrpc:"rating"`
	ResId          int64     `xmlrpc:"res_id"`
	ResModel       string    `xmlrpc:"res_model"`
	ResName        string    `xmlrpc:"res_name"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type RatingRatingNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	AccessToken    interface{} `xmlrpc:"access_token"`
	Consumed       bool        `xmlrpc:"consumed"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Feedback       interface{} `xmlrpc:"feedback"`
	Id             interface{} `xmlrpc:"id"`
	MessageId      interface{} `xmlrpc:"message_id"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	RatedPartnerId interface{} `xmlrpc:"rated_partner_id"`
	Rating         interface{} `xmlrpc:"rating"`
	ResId          interface{} `xmlrpc:"res_id"`
	ResModel       interface{} `xmlrpc:"res_model"`
	ResName        interface{} `xmlrpc:"res_name"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var RatingRatingModel string = "rating.rating"

type RatingRatings []RatingRating

type RatingRatingsNil []RatingRatingNil

func (s *RatingRating) NilableType_() interface{} {
	return &RatingRatingNil{}
}

func (ns *RatingRatingNil) Type_() interface{} {
	s := &RatingRating{}
	return load(ns, s)
}

func (s *RatingRatings) NilableType_() interface{} {
	return &RatingRatingsNil{}
}

func (ns *RatingRatingsNil) Type_() interface{} {
	s := &RatingRatings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*RatingRating))
	}
	return s
}
