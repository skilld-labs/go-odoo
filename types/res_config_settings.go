package types

import (
	"time"
)

type ResConfigSettings struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResConfigSettingsNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResConfigSettingsModel string = "res.config.settings"

type ResConfigSettingss []ResConfigSettings

type ResConfigSettingssNil []ResConfigSettingsNil

func (s *ResConfigSettings) NilableType_() interface{} {
	return &ResConfigSettingsNil{}
}

func (ns *ResConfigSettingsNil) Type_() interface{} {
	s := &ResConfigSettings{}
	return load(ns, s)
}

func (s *ResConfigSettingss) NilableType_() interface{} {
	return &ResConfigSettingssNil{}
}

func (ns *ResConfigSettingssNil) Type_() interface{} {
	s := &ResConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResConfigSettings))
	}
	return s
}
