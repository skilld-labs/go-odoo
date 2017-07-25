package types

import (
	"time"
)

type IrNeedactionMixin struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrNeedactionMixinNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrNeedactionMixinModel string = "ir.needaction_mixin"

type IrNeedactionMixins []IrNeedactionMixin

type IrNeedactionMixinsNil []IrNeedactionMixinNil

func (s *IrNeedactionMixin) NilableType_() interface{} {
	return &IrNeedactionMixinNil{}
}

func (ns *IrNeedactionMixinNil) Type_() interface{} {
	s := &IrNeedactionMixin{}
	return load(ns, s)
}

func (s *IrNeedactionMixins) NilableType_() interface{} {
	return &IrNeedactionMixinsNil{}
}

func (ns *IrNeedactionMixinsNil) Type_() interface{} {
	s := &IrNeedactionMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrNeedactionMixin))
	}
	return s
}
