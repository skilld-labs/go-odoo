package types

import (
	"time"
)

type ProjectConfigSettings struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	CompanyId             Many2One  `xmlrpc:"company_id"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	GenerateProjectAlias  string    `xmlrpc:"generate_project_alias"`
	Id                    int64     `xmlrpc:"id"`
	ModulePad             string    `xmlrpc:"module_pad"`
	ModuleProjectForecast bool      `xmlrpc:"module_project_forecast"`
	ModuleRatingProject   string    `xmlrpc:"module_rating_project"`
	ProjectTimeModeId     Many2One  `xmlrpc:"project_time_mode_id"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type ProjectConfigSettingsNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	CompanyId             interface{} `xmlrpc:"company_id"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	GenerateProjectAlias  interface{} `xmlrpc:"generate_project_alias"`
	Id                    interface{} `xmlrpc:"id"`
	ModulePad             interface{} `xmlrpc:"module_pad"`
	ModuleProjectForecast bool        `xmlrpc:"module_project_forecast"`
	ModuleRatingProject   interface{} `xmlrpc:"module_rating_project"`
	ProjectTimeModeId     interface{} `xmlrpc:"project_time_mode_id"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var ProjectConfigSettingsModel string = "project.config.settings"

type ProjectConfigSettingss []ProjectConfigSettings

type ProjectConfigSettingssNil []ProjectConfigSettingsNil

func (s *ProjectConfigSettings) NilableType_() interface{} {
	return &ProjectConfigSettingsNil{}
}

func (ns *ProjectConfigSettingsNil) Type_() interface{} {
	s := &ProjectConfigSettings{}
	return load(ns, s)
}

func (s *ProjectConfigSettingss) NilableType_() interface{} {
	return &ProjectConfigSettingssNil{}
}

func (ns *ProjectConfigSettingssNil) Type_() interface{} {
	s := &ProjectConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectConfigSettings))
	}
	return s
}
