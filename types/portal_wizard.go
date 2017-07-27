package types

import (
	"time"
)

type PortalWizard struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	PortalId       Many2One  `xmlrpc:"portal_id"`
	UserIds        []int64   `xmlrpc:"user_ids"`
	WelcomeMessage string    `xmlrpc:"welcome_message"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type PortalWizardNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	PortalId       interface{} `xmlrpc:"portal_id"`
	UserIds        interface{} `xmlrpc:"user_ids"`
	WelcomeMessage interface{} `xmlrpc:"welcome_message"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var PortalWizardModel string = "portal.wizard"

type PortalWizards []PortalWizard

type PortalWizardsNil []PortalWizardNil

func (s *PortalWizard) NilableType_() interface{} {
	return &PortalWizardNil{}
}

func (ns *PortalWizardNil) Type_() interface{} {
	s := &PortalWizard{}
	return load(ns, s)
}

func (s *PortalWizards) NilableType_() interface{} {
	return &PortalWizardsNil{}
}

func (ns *PortalWizardsNil) Type_() interface{} {
	s := &PortalWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PortalWizard))
	}
	return s
}
