package types

import (
	"time"
)

type Unknown struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type UnknownNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var UnknownModel string = "_unknown"

type Unknowns []Unknown

type UnknownsNil []UnknownNil

func (s *Unknown) NilableType_() interface{} {
	return &UnknownNil{}
}

func (ns *UnknownNil) Type_() interface{} {
	s := &Unknown{}
	return load(ns, s)
}

func (s *Unknowns) NilableType_() interface{} {
	return &UnknownsNil{}
}

func (ns *UnknownsNil) Type_() interface{} {
	s := &Unknowns{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*Unknown))
	}
	return s
}
