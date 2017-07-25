package types

import (
	"time"
)

type PortalWizardUser struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Email       string    `xmlrpc:"email"`
	Id          int64     `xmlrpc:"id"`
	InPortal    bool      `xmlrpc:"in_portal"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WizardId    Many2One  `xmlrpc:"wizard_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type PortalWizardUserNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Email       interface{} `xmlrpc:"email"`
	Id          interface{} `xmlrpc:"id"`
	InPortal    bool        `xmlrpc:"in_portal"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	UserId      interface{} `xmlrpc:"user_id"`
	WizardId    interface{} `xmlrpc:"wizard_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var PortalWizardUserModel string = "portal.wizard.user"

type PortalWizardUsers []PortalWizardUser

type PortalWizardUsersNil []PortalWizardUserNil

func (s *PortalWizardUser) NilableType_() interface{} {
	return &PortalWizardUserNil{}
}

func (ns *PortalWizardUserNil) Type_() interface{} {
	s := &PortalWizardUser{}
	return load(ns, s)
}

func (s *PortalWizardUsers) NilableType_() interface{} {
	return &PortalWizardUsersNil{}
}

func (ns *PortalWizardUsersNil) Type_() interface{} {
	s := &PortalWizardUsers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PortalWizardUser))
	}
	return s
}
