package types

import (
	"time"
)

type MailActivityMixin struct {
	ActivityDateDeadline time.Time   `xmlrpc:"activity_date_deadline"`
	ActivityIds          []int64     `xmlrpc:"activity_ids"`
	ActivityState        interface{} `xmlrpc:"activity_state"`
	ActivitySummary      string      `xmlrpc:"activity_summary"`
	ActivityTypeId       Many2One    `xmlrpc:"activity_type_id"`
	ActivityUserId       Many2One    `xmlrpc:"activity_user_id"`
	DisplayName          string      `xmlrpc:"display_name"`
	Id                   int64       `xmlrpc:"id"`
	LastUpdate           time.Time   `xmlrpc:"__last_update"`
}

type MailActivityMixinNil struct {
	ActivityDateDeadline interface{} `xmlrpc:"activity_date_deadline"`
	ActivityIds          interface{} `xmlrpc:"activity_ids"`
	ActivityState        interface{} `xmlrpc:"activity_state"`
	ActivitySummary      interface{} `xmlrpc:"activity_summary"`
	ActivityTypeId       interface{} `xmlrpc:"activity_type_id"`
	ActivityUserId       interface{} `xmlrpc:"activity_user_id"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	Id                   interface{} `xmlrpc:"id"`
	LastUpdate           interface{} `xmlrpc:"__last_update"`
}

var MailActivityMixinModel string = "mail.activity.mixin"

type MailActivityMixins []MailActivityMixin

type MailActivityMixinsNil []MailActivityMixinNil

func (s *MailActivityMixin) NilableType_() interface{} {
	return &MailActivityMixinNil{}
}

func (ns *MailActivityMixinNil) Type_() interface{} {
	s := &MailActivityMixin{}
	return load(ns, s)
}

func (s *MailActivityMixins) NilableType_() interface{} {
	return &MailActivityMixinsNil{}
}

func (ns *MailActivityMixinsNil) Type_() interface{} {
	s := &MailActivityMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailActivityMixin))
	}
	return s
}
