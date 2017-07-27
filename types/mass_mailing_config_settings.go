package types

import (
	"time"
)

type MassMailingConfigSettings struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DisplayName              string    `xmlrpc:"display_name"`
	GroupMassMailingCampaign string    `xmlrpc:"group_mass_mailing_campaign"`
	Id                       int64     `xmlrpc:"id"`
	ModuleMassMailingThemes  bool      `xmlrpc:"module_mass_mailing_themes"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type MassMailingConfigSettingsNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	GroupMassMailingCampaign interface{} `xmlrpc:"group_mass_mailing_campaign"`
	Id                       interface{} `xmlrpc:"id"`
	ModuleMassMailingThemes  bool        `xmlrpc:"module_mass_mailing_themes"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var MassMailingConfigSettingsModel string = "mass.mailing.config.settings"

type MassMailingConfigSettingss []MassMailingConfigSettings

type MassMailingConfigSettingssNil []MassMailingConfigSettingsNil

func (s *MassMailingConfigSettings) NilableType_() interface{} {
	return &MassMailingConfigSettingsNil{}
}

func (ns *MassMailingConfigSettingsNil) Type_() interface{} {
	s := &MassMailingConfigSettings{}
	return load(ns, s)
}

func (s *MassMailingConfigSettingss) NilableType_() interface{} {
	return &MassMailingConfigSettingssNil{}
}

func (ns *MassMailingConfigSettingssNil) Type_() interface{} {
	s := &MassMailingConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MassMailingConfigSettings))
	}
	return s
}
