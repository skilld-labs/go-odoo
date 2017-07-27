package types

import (
	"time"
)

type CrmStage struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Fold           bool      `xmlrpc:"fold"`
	Id             int64     `xmlrpc:"id"`
	LegendPriority string    `xmlrpc:"legend_priority"`
	Name           string    `xmlrpc:"name"`
	OnChange       bool      `xmlrpc:"on_change"`
	Probability    float64   `xmlrpc:"probability"`
	Requirements   string    `xmlrpc:"requirements"`
	Sequence       int64     `xmlrpc:"sequence"`
	TeamId         Many2One  `xmlrpc:"team_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type CrmStageNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Fold           bool        `xmlrpc:"fold"`
	Id             interface{} `xmlrpc:"id"`
	LegendPriority interface{} `xmlrpc:"legend_priority"`
	Name           interface{} `xmlrpc:"name"`
	OnChange       bool        `xmlrpc:"on_change"`
	Probability    interface{} `xmlrpc:"probability"`
	Requirements   interface{} `xmlrpc:"requirements"`
	Sequence       interface{} `xmlrpc:"sequence"`
	TeamId         interface{} `xmlrpc:"team_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var CrmStageModel string = "crm.stage"

type CrmStages []CrmStage

type CrmStagesNil []CrmStageNil

func (s *CrmStage) NilableType_() interface{} {
	return &CrmStageNil{}
}

func (ns *CrmStageNil) Type_() interface{} {
	s := &CrmStage{}
	return load(ns, s)
}

func (s *CrmStages) NilableType_() interface{} {
	return &CrmStagesNil{}
}

func (ns *CrmStagesNil) Type_() interface{} {
	s := &CrmStages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmStage))
	}
	return s
}
