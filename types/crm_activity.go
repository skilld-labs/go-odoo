package types

import (
	"time"
)

type CrmActivity struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	Days                   int64     `xmlrpc:"days"`
	Default                bool      `xmlrpc:"default"`
	Description            string    `xmlrpc:"description"`
	DisplayName            string    `xmlrpc:"display_name"`
	Hidden                 bool      `xmlrpc:"hidden"`
	Id                     int64     `xmlrpc:"id"`
	Internal               bool      `xmlrpc:"internal"`
	Name                   string    `xmlrpc:"name"`
	ParentId               Many2One  `xmlrpc:"parent_id"`
	PrecedingActivityIds   []int64   `xmlrpc:"preceding_activity_ids"`
	RecommendedActivityIds []int64   `xmlrpc:"recommended_activity_ids"`
	RelationField          string    `xmlrpc:"relation_field"`
	ResModel               string    `xmlrpc:"res_model"`
	Sequence               int64     `xmlrpc:"sequence"`
	SubtypeId              Many2One  `xmlrpc:"subtype_id"`
	TeamId                 Many2One  `xmlrpc:"team_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type CrmActivityNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	Days                   interface{} `xmlrpc:"days"`
	Default                bool        `xmlrpc:"default"`
	Description            interface{} `xmlrpc:"description"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Hidden                 bool        `xmlrpc:"hidden"`
	Id                     interface{} `xmlrpc:"id"`
	Internal               bool        `xmlrpc:"internal"`
	Name                   interface{} `xmlrpc:"name"`
	ParentId               interface{} `xmlrpc:"parent_id"`
	PrecedingActivityIds   interface{} `xmlrpc:"preceding_activity_ids"`
	RecommendedActivityIds interface{} `xmlrpc:"recommended_activity_ids"`
	RelationField          interface{} `xmlrpc:"relation_field"`
	ResModel               interface{} `xmlrpc:"res_model"`
	Sequence               interface{} `xmlrpc:"sequence"`
	SubtypeId              interface{} `xmlrpc:"subtype_id"`
	TeamId                 interface{} `xmlrpc:"team_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var CrmActivityModel string = "crm.activity"

type CrmActivitys []CrmActivity

type CrmActivitysNil []CrmActivityNil

func (s *CrmActivity) NilableType_() interface{} {
	return &CrmActivityNil{}
}

func (ns *CrmActivityNil) Type_() interface{} {
	s := &CrmActivity{}
	return load(ns, s)
}

func (s *CrmActivitys) NilableType_() interface{} {
	return &CrmActivitysNil{}
}

func (ns *CrmActivitysNil) Type_() interface{} {
	s := &CrmActivitys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmActivity))
	}
	return s
}
