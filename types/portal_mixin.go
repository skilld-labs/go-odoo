package types

import (
	"time"
)

type PortalMixin struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	PortalUrl   string    `xmlrpc:"portal_url"`
}

type PortalMixinNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	PortalUrl   interface{} `xmlrpc:"portal_url"`
}

var PortalMixinModel string = "portal.mixin"

type PortalMixins []PortalMixin

type PortalMixinsNil []PortalMixinNil

func (s *PortalMixin) NilableType_() interface{} {
	return &PortalMixinNil{}
}

func (ns *PortalMixinNil) Type_() interface{} {
	s := &PortalMixin{}
	return load(ns, s)
}

func (s *PortalMixins) NilableType_() interface{} {
	return &PortalMixinsNil{}
}

func (ns *PortalMixinsNil) Type_() interface{} {
	s := &PortalMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PortalMixin))
	}
	return s
}
