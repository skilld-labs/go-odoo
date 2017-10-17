package types

import (
	"time"
)

type ResourceMixin struct {
	CompanyId          Many2One  `xmlrpc:"company_id"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	ResourceCalendarId Many2One  `xmlrpc:"resource_calendar_id"`
	ResourceId         Many2One  `xmlrpc:"resource_id"`
}

type ResourceMixinNil struct {
	CompanyId          interface{} `xmlrpc:"company_id"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	ResourceCalendarId interface{} `xmlrpc:"resource_calendar_id"`
	ResourceId         interface{} `xmlrpc:"resource_id"`
}

var ResourceMixinModel string = "resource.mixin"

type ResourceMixins []ResourceMixin

type ResourceMixinsNil []ResourceMixinNil

func (s *ResourceMixin) NilableType_() interface{} {
	return &ResourceMixinNil{}
}

func (ns *ResourceMixinNil) Type_() interface{} {
	s := &ResourceMixin{}
	return load(ns, s)
}

func (s *ResourceMixins) NilableType_() interface{} {
	return &ResourceMixinsNil{}
}

func (ns *ResourceMixinsNil) Type_() interface{} {
	s := &ResourceMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceMixin))
	}
	return s
}
