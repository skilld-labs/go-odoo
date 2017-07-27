package types

import (
	"time"
)

type ProjectTags struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProjectTagsNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProjectTagsModel string = "project.tags"

type ProjectTagss []ProjectTags

type ProjectTagssNil []ProjectTagsNil

func (s *ProjectTags) NilableType_() interface{} {
	return &ProjectTagsNil{}
}

func (ns *ProjectTagsNil) Type_() interface{} {
	s := &ProjectTags{}
	return load(ns, s)
}

func (s *ProjectTagss) NilableType_() interface{} {
	return &ProjectTagssNil{}
}

func (ns *ProjectTagssNil) Type_() interface{} {
	s := &ProjectTagss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectTags))
	}
	return s
}
