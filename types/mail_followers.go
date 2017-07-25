package types

import (
	"time"
)

type MailFollowers struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ChannelId   Many2One  `xmlrpc:"channel_id"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	ResId       int64     `xmlrpc:"res_id"`
	ResModel    string    `xmlrpc:"res_model"`
	SubtypeIds  []int64   `xmlrpc:"subtype_ids"`
}

type MailFollowersNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ChannelId   interface{} `xmlrpc:"channel_id"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	ResId       interface{} `xmlrpc:"res_id"`
	ResModel    interface{} `xmlrpc:"res_model"`
	SubtypeIds  interface{} `xmlrpc:"subtype_ids"`
}

var MailFollowersModel string = "mail.followers"

type MailFollowerss []MailFollowers

type MailFollowerssNil []MailFollowersNil

func (s *MailFollowers) NilableType_() interface{} {
	return &MailFollowersNil{}
}

func (ns *MailFollowersNil) Type_() interface{} {
	s := &MailFollowers{}
	return load(ns, s)
}

func (s *MailFollowerss) NilableType_() interface{} {
	return &MailFollowerssNil{}
}

func (ns *MailFollowerssNil) Type_() interface{} {
	s := &MailFollowerss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailFollowers))
	}
	return s
}
