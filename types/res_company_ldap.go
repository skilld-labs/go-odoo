package types

import (
	"time"
)

type ResCompanyLdap struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Company        Many2One  `xmlrpc:"company"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	CreateUser     bool      `xmlrpc:"create_user"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LdapBase       string    `xmlrpc:"ldap_base"`
	LdapBinddn     string    `xmlrpc:"ldap_binddn"`
	LdapFilter     string    `xmlrpc:"ldap_filter"`
	LdapPassword   string    `xmlrpc:"ldap_password"`
	LdapServer     string    `xmlrpc:"ldap_server"`
	LdapServerPort int64     `xmlrpc:"ldap_server_port"`
	LdapTls        bool      `xmlrpc:"ldap_tls"`
	Sequence       int64     `xmlrpc:"sequence"`
	User           Many2One  `xmlrpc:"user"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ResCompanyLdapNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Company        interface{} `xmlrpc:"company"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	CreateUser     bool        `xmlrpc:"create_user"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LdapBase       interface{} `xmlrpc:"ldap_base"`
	LdapBinddn     interface{} `xmlrpc:"ldap_binddn"`
	LdapFilter     interface{} `xmlrpc:"ldap_filter"`
	LdapPassword   interface{} `xmlrpc:"ldap_password"`
	LdapServer     interface{} `xmlrpc:"ldap_server"`
	LdapServerPort interface{} `xmlrpc:"ldap_server_port"`
	LdapTls        bool        `xmlrpc:"ldap_tls"`
	Sequence       interface{} `xmlrpc:"sequence"`
	User           interface{} `xmlrpc:"user"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var ResCompanyLdapModel string = "res.company.ldap"

type ResCompanyLdaps []ResCompanyLdap

type ResCompanyLdapsNil []ResCompanyLdapNil

func (s *ResCompanyLdap) NilableType_() interface{} {
	return &ResCompanyLdapNil{}
}

func (ns *ResCompanyLdapNil) Type_() interface{} {
	s := &ResCompanyLdap{}
	return load(ns, s)
}

func (s *ResCompanyLdaps) NilableType_() interface{} {
	return &ResCompanyLdapsNil{}
}

func (ns *ResCompanyLdapsNil) Type_() interface{} {
	s := &ResCompanyLdaps{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCompanyLdap))
	}
	return s
}
