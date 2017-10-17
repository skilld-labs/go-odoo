package types

import (
	"time"
)

type FormatAddressMixin struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type FormatAddressMixinNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var FormatAddressMixinModel string = "format.address.mixin"

type FormatAddressMixins []FormatAddressMixin

type FormatAddressMixinsNil []FormatAddressMixinNil

func (s *FormatAddressMixin) NilableType_() interface{} {
	return &FormatAddressMixinNil{}
}

func (ns *FormatAddressMixinNil) Type_() interface{} {
	s := &FormatAddressMixin{}
	return load(ns, s)
}

func (s *FormatAddressMixins) NilableType_() interface{} {
	return &FormatAddressMixinsNil{}
}

func (ns *FormatAddressMixinsNil) Type_() interface{} {
	s := &FormatAddressMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*FormatAddressMixin))
	}
	return s
}
