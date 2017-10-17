package types

import (
	"time"
)

type ResourceResource struct {
	Active         bool      `xmlrpc:"active"`
	CalendarId     Many2One  `xmlrpc:"calendar_id"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Name           string    `xmlrpc:"name"`
	ResourceType   string    `xmlrpc:"resource_type"`
	TimeEfficiency float64   `xmlrpc:"time_efficiency"`
	UserId         Many2One  `xmlrpc:"user_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ResourceResourceNil struct {
	Active         bool        `xmlrpc:"active"`
	CalendarId     interface{} `xmlrpc:"calendar_id"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Name           interface{} `xmlrpc:"name"`
	ResourceType   interface{} `xmlrpc:"resource_type"`
	TimeEfficiency interface{} `xmlrpc:"time_efficiency"`
	UserId         interface{} `xmlrpc:"user_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var ResourceResourceModel string = "resource.resource"

type ResourceResources []ResourceResource

type ResourceResourcesNil []ResourceResourceNil

func (s *ResourceResource) NilableType_() interface{} {
	return &ResourceResourceNil{}
}

func (ns *ResourceResourceNil) Type_() interface{} {
	s := &ResourceResource{}
	return load(ns, s)
}

func (s *ResourceResources) NilableType_() interface{} {
	return &ResourceResourcesNil{}
}

func (ns *ResourceResourcesNil) Type_() interface{} {
	s := &ResourceResources{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceResource))
	}
	return s
}
