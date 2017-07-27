package types

import (
	"time"
)

type IrSequence struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	Active           bool      `xmlrpc:"active"`
	Code             string    `xmlrpc:"code"`
	CompanyId        Many2One  `xmlrpc:"company_id"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DateRangeIds     []int64   `xmlrpc:"date_range_ids"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	Implementation   string    `xmlrpc:"implementation"`
	Name             string    `xmlrpc:"name"`
	NumberIncrement  int64     `xmlrpc:"number_increment"`
	NumberNext       int64     `xmlrpc:"number_next"`
	NumberNextActual int64     `xmlrpc:"number_next_actual"`
	Padding          int64     `xmlrpc:"padding"`
	Prefix           string    `xmlrpc:"prefix"`
	Suffix           string    `xmlrpc:"suffix"`
	UseDateRange     bool      `xmlrpc:"use_date_range"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type IrSequenceNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	Active           bool        `xmlrpc:"active"`
	Code             interface{} `xmlrpc:"code"`
	CompanyId        interface{} `xmlrpc:"company_id"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DateRangeIds     interface{} `xmlrpc:"date_range_ids"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	Implementation   interface{} `xmlrpc:"implementation"`
	Name             interface{} `xmlrpc:"name"`
	NumberIncrement  interface{} `xmlrpc:"number_increment"`
	NumberNext       interface{} `xmlrpc:"number_next"`
	NumberNextActual interface{} `xmlrpc:"number_next_actual"`
	Padding          interface{} `xmlrpc:"padding"`
	Prefix           interface{} `xmlrpc:"prefix"`
	Suffix           interface{} `xmlrpc:"suffix"`
	UseDateRange     bool        `xmlrpc:"use_date_range"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var IrSequenceModel string = "ir.sequence"

type IrSequences []IrSequence

type IrSequencesNil []IrSequenceNil

func (s *IrSequence) NilableType_() interface{} {
	return &IrSequenceNil{}
}

func (ns *IrSequenceNil) Type_() interface{} {
	s := &IrSequence{}
	return load(ns, s)
}

func (s *IrSequences) NilableType_() interface{} {
	return &IrSequencesNil{}
}

func (ns *IrSequencesNil) Type_() interface{} {
	s := &IrSequences{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrSequence))
	}
	return s
}
