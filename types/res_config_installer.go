package types

import (
	"time"
)

type ResConfigInstaller struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResConfigInstallerNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResConfigInstallerModel string = "res.config.installer"

type ResConfigInstallers []ResConfigInstaller

type ResConfigInstallersNil []ResConfigInstallerNil

func (s *ResConfigInstaller) NilableType_() interface{} {
	return &ResConfigInstallerNil{}
}

func (ns *ResConfigInstallerNil) Type_() interface{} {
	s := &ResConfigInstaller{}
	return load(ns, s)
}

func (s *ResConfigInstallers) NilableType_() interface{} {
	return &ResConfigInstallersNil{}
}

func (ns *ResConfigInstallersNil) Type_() interface{} {
	s := &ResConfigInstallers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResConfigInstaller))
	}
	return s
}
