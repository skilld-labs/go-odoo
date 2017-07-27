package types

import (
	"time"
)

type IrModuleModule struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	Application      bool      `xmlrpc:"application"`
	Author           string    `xmlrpc:"author"`
	AutoInstall      bool      `xmlrpc:"auto_install"`
	CategoryId       Many2One  `xmlrpc:"category_id"`
	Contributors     string    `xmlrpc:"contributors"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	Demo             bool      `xmlrpc:"demo"`
	DependenciesId   []int64   `xmlrpc:"dependencies_id"`
	Description      string    `xmlrpc:"description"`
	DescriptionHtml  string    `xmlrpc:"description_html"`
	DisplayName      string    `xmlrpc:"display_name"`
	Icon             string    `xmlrpc:"icon"`
	IconImage        string    `xmlrpc:"icon_image"`
	Id               int64     `xmlrpc:"id"`
	InstalledVersion string    `xmlrpc:"installed_version"`
	LatestVersion    string    `xmlrpc:"latest_version"`
	License          string    `xmlrpc:"license"`
	Maintainer       string    `xmlrpc:"maintainer"`
	MenusByModule    string    `xmlrpc:"menus_by_module"`
	Name             string    `xmlrpc:"name"`
	PublishedVersion string    `xmlrpc:"published_version"`
	ReportsByModule  string    `xmlrpc:"reports_by_module"`
	Sequence         int64     `xmlrpc:"sequence"`
	Shortdesc        string    `xmlrpc:"shortdesc"`
	State            string    `xmlrpc:"state"`
	Summary          string    `xmlrpc:"summary"`
	Url              string    `xmlrpc:"url"`
	ViewsByModule    string    `xmlrpc:"views_by_module"`
	Website          string    `xmlrpc:"website"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type IrModuleModuleNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	Application      bool        `xmlrpc:"application"`
	Author           interface{} `xmlrpc:"author"`
	AutoInstall      bool        `xmlrpc:"auto_install"`
	CategoryId       interface{} `xmlrpc:"category_id"`
	Contributors     interface{} `xmlrpc:"contributors"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	Demo             bool        `xmlrpc:"demo"`
	DependenciesId   interface{} `xmlrpc:"dependencies_id"`
	Description      interface{} `xmlrpc:"description"`
	DescriptionHtml  interface{} `xmlrpc:"description_html"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Icon             interface{} `xmlrpc:"icon"`
	IconImage        interface{} `xmlrpc:"icon_image"`
	Id               interface{} `xmlrpc:"id"`
	InstalledVersion interface{} `xmlrpc:"installed_version"`
	LatestVersion    interface{} `xmlrpc:"latest_version"`
	License          interface{} `xmlrpc:"license"`
	Maintainer       interface{} `xmlrpc:"maintainer"`
	MenusByModule    interface{} `xmlrpc:"menus_by_module"`
	Name             interface{} `xmlrpc:"name"`
	PublishedVersion interface{} `xmlrpc:"published_version"`
	ReportsByModule  interface{} `xmlrpc:"reports_by_module"`
	Sequence         interface{} `xmlrpc:"sequence"`
	Shortdesc        interface{} `xmlrpc:"shortdesc"`
	State            interface{} `xmlrpc:"state"`
	Summary          interface{} `xmlrpc:"summary"`
	Url              interface{} `xmlrpc:"url"`
	ViewsByModule    interface{} `xmlrpc:"views_by_module"`
	Website          interface{} `xmlrpc:"website"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var IrModuleModuleModel string = "ir.module.module"

type IrModuleModules []IrModuleModule

type IrModuleModulesNil []IrModuleModuleNil

func (s *IrModuleModule) NilableType_() interface{} {
	return &IrModuleModuleNil{}
}

func (ns *IrModuleModuleNil) Type_() interface{} {
	s := &IrModuleModule{}
	return load(ns, s)
}

func (s *IrModuleModules) NilableType_() interface{} {
	return &IrModuleModulesNil{}
}

func (ns *IrModuleModulesNil) Type_() interface{} {
	s := &IrModuleModules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModuleModule))
	}
	return s
}
