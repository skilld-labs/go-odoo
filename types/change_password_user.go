package types

import (
	"time"
)

type ChangePasswordUser struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	NewPasswd   string    `xmlrpc:"new_passwd"`
	UserId      Many2One  `xmlrpc:"user_id"`
	UserLogin   string    `xmlrpc:"user_login"`
	WizardId    Many2One  `xmlrpc:"wizard_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ChangePasswordUserNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	NewPasswd   interface{} `xmlrpc:"new_passwd"`
	UserId      interface{} `xmlrpc:"user_id"`
	UserLogin   interface{} `xmlrpc:"user_login"`
	WizardId    interface{} `xmlrpc:"wizard_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ChangePasswordUserModel string = "change.password.user"

type ChangePasswordUsers []ChangePasswordUser

type ChangePasswordUsersNil []ChangePasswordUserNil

func (s *ChangePasswordUser) NilableType_() interface{} {
	return &ChangePasswordUserNil{}
}

func (ns *ChangePasswordUserNil) Type_() interface{} {
	s := &ChangePasswordUser{}
	return load(ns, s)
}

func (s *ChangePasswordUsers) NilableType_() interface{} {
	return &ChangePasswordUsersNil{}
}

func (ns *ChangePasswordUsersNil) Type_() interface{} {
	s := &ChangePasswordUsers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ChangePasswordUser))
	}
	return s
}
