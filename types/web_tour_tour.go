package types

import (
	"time"
)

type WebTourTour struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	UserId      Many2One  `xmlrpc:"user_id"`
}

type WebTourTourNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	UserId      interface{} `xmlrpc:"user_id"`
}

var WebTourTourModel string = "web_tour.tour"

type WebTourTours []WebTourTour

type WebTourToursNil []WebTourTourNil

func (s *WebTourTour) NilableType_() interface{} {
	return &WebTourTourNil{}
}

func (ns *WebTourTourNil) Type_() interface{} {
	s := &WebTourTour{}
	return load(ns, s)
}

func (s *WebTourTours) NilableType_() interface{} {
	return &WebTourToursNil{}
}

func (ns *WebTourToursNil) Type_() interface{} {
	s := &WebTourTours{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WebTourTour))
	}
	return s
}
