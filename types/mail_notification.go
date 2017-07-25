package types

import (
	"time"
)

type MailNotification struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	DisplayName   string    `xmlrpc:"display_name"`
	EmailStatus   string    `xmlrpc:"email_status"`
	Id            int64     `xmlrpc:"id"`
	IsEmail       bool      `xmlrpc:"is_email"`
	IsRead        bool      `xmlrpc:"is_read"`
	MailMessageId Many2One  `xmlrpc:"mail_message_id"`
	ResPartnerId  Many2One  `xmlrpc:"res_partner_id"`
}

type MailNotificationNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	EmailStatus   interface{} `xmlrpc:"email_status"`
	Id            interface{} `xmlrpc:"id"`
	IsEmail       bool        `xmlrpc:"is_email"`
	IsRead        bool        `xmlrpc:"is_read"`
	MailMessageId interface{} `xmlrpc:"mail_message_id"`
	ResPartnerId  interface{} `xmlrpc:"res_partner_id"`
}

var MailNotificationModel string = "mail.notification"

type MailNotifications []MailNotification

type MailNotificationsNil []MailNotificationNil

func (s *MailNotification) NilableType_() interface{} {
	return &MailNotificationNil{}
}

func (ns *MailNotificationNil) Type_() interface{} {
	s := &MailNotification{}
	return load(ns, s)
}

func (s *MailNotifications) NilableType_() interface{} {
	return &MailNotificationsNil{}
}

func (ns *MailNotificationsNil) Type_() interface{} {
	s := &MailNotifications{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailNotification))
	}
	return s
}
