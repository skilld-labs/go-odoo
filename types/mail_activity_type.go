package types

import (
	"time"
)

type MailActivityType struct {
	Category        string    `xmlrpc:"category"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	Days            int64     `xmlrpc:"days"`
	DisplayName     string    `xmlrpc:"display_name"`
	Icon            string    `xmlrpc:"icon"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Name            string    `xmlrpc:"name"`
	NextTypeIds     []int64   `xmlrpc:"next_type_ids"`
	PreviousTypeIds []int64   `xmlrpc:"previous_type_ids"`
	ResModelId      Many2One  `xmlrpc:"res_model_id"`
	Sequence        int64     `xmlrpc:"sequence"`
	Summary         string    `xmlrpc:"summary"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type MailActivityTypeNil struct {
	Category        interface{} `xmlrpc:"category"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	Days            interface{} `xmlrpc:"days"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Icon            interface{} `xmlrpc:"icon"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Name            interface{} `xmlrpc:"name"`
	NextTypeIds     interface{} `xmlrpc:"next_type_ids"`
	PreviousTypeIds interface{} `xmlrpc:"previous_type_ids"`
	ResModelId      interface{} `xmlrpc:"res_model_id"`
	Sequence        interface{} `xmlrpc:"sequence"`
	Summary         interface{} `xmlrpc:"summary"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var MailActivityTypeModel string = "mail.activity.type"

type MailActivityTypes []MailActivityType

type MailActivityTypesNil []MailActivityTypeNil

func (s *MailActivityType) NilableType_() interface{} {
	return &MailActivityTypeNil{}
}

func (ns *MailActivityTypeNil) Type_() interface{} {
	s := &MailActivityType{}
	return load(ns, s)
}

func (s *MailActivityTypes) NilableType_() interface{} {
	return &MailActivityTypesNil{}
}

func (ns *MailActivityTypesNil) Type_() interface{} {
	s := &MailActivityTypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailActivityType))
	}
	return s
}
