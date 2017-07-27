package types

import (
	"time"
)

type IrSequenceDateRange struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DateFrom         time.Time `xmlrpc:"date_from"`
	DateTo           time.Time `xmlrpc:"date_to"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	NumberNext       int64     `xmlrpc:"number_next"`
	NumberNextActual int64     `xmlrpc:"number_next_actual"`
	SequenceId       Many2One  `xmlrpc:"sequence_id"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type IrSequenceDateRangeNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DateFrom         interface{} `xmlrpc:"date_from"`
	DateTo           interface{} `xmlrpc:"date_to"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	NumberNext       interface{} `xmlrpc:"number_next"`
	NumberNextActual interface{} `xmlrpc:"number_next_actual"`
	SequenceId       interface{} `xmlrpc:"sequence_id"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var IrSequenceDateRangeModel string = "ir.sequence.date_range"

type IrSequenceDateRanges []IrSequenceDateRange

type IrSequenceDateRangesNil []IrSequenceDateRangeNil

func (s *IrSequenceDateRange) NilableType_() interface{} {
	return &IrSequenceDateRangeNil{}
}

func (ns *IrSequenceDateRangeNil) Type_() interface{} {
	s := &IrSequenceDateRange{}
	return load(ns, s)
}

func (s *IrSequenceDateRanges) NilableType_() interface{} {
	return &IrSequenceDateRangesNil{}
}

func (ns *IrSequenceDateRangesNil) Type_() interface{} {
	s := &IrSequenceDateRanges{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrSequenceDateRange))
	}
	return s
}
