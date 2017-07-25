package types

import (
	"time"
)

type MailMail struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	AttachmentIds        []int64   `xmlrpc:"attachment_ids"`
	AuthorAvatar         string    `xmlrpc:"author_avatar"`
	AuthorId             Many2One  `xmlrpc:"author_id"`
	AutoDelete           bool      `xmlrpc:"auto_delete"`
	Body                 string    `xmlrpc:"body"`
	BodyHtml             string    `xmlrpc:"body_html"`
	ChannelIds           []int64   `xmlrpc:"channel_ids"`
	ChildIds             []int64   `xmlrpc:"child_ids"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	Date                 time.Time `xmlrpc:"date"`
	DisplayName          string    `xmlrpc:"display_name"`
	EmailCc              string    `xmlrpc:"email_cc"`
	EmailFrom            string    `xmlrpc:"email_from"`
	EmailTo              string    `xmlrpc:"email_to"`
	FailureReason        string    `xmlrpc:"failure_reason"`
	FetchmailServerId    Many2One  `xmlrpc:"fetchmail_server_id"`
	Headers              string    `xmlrpc:"headers"`
	Id                   int64     `xmlrpc:"id"`
	MailMessageId        Many2One  `xmlrpc:"mail_message_id"`
	MailServerId         Many2One  `xmlrpc:"mail_server_id"`
	MailingId            Many2One  `xmlrpc:"mailing_id"`
	MessageId            string    `xmlrpc:"message_id"`
	MessageType          string    `xmlrpc:"message_type"`
	Model                string    `xmlrpc:"model"`
	Needaction           bool      `xmlrpc:"needaction"`
	NeedactionPartnerIds []int64   `xmlrpc:"needaction_partner_ids"`
	NoAutoThread         bool      `xmlrpc:"no_auto_thread"`
	Notification         bool      `xmlrpc:"notification"`
	NotificationIds      []int64   `xmlrpc:"notification_ids"`
	ParentId             Many2One  `xmlrpc:"parent_id"`
	PartnerIds           []int64   `xmlrpc:"partner_ids"`
	RecipientIds         []int64   `xmlrpc:"recipient_ids"`
	RecordName           string    `xmlrpc:"record_name"`
	References           string    `xmlrpc:"references"`
	ReplyTo              string    `xmlrpc:"reply_to"`
	ResId                int64     `xmlrpc:"res_id"`
	ScheduledDate        string    `xmlrpc:"scheduled_date"`
	Starred              bool      `xmlrpc:"starred"`
	StarredPartnerIds    []int64   `xmlrpc:"starred_partner_ids"`
	State                string    `xmlrpc:"state"`
	StatisticsIds        []int64   `xmlrpc:"statistics_ids"`
	Subject              string    `xmlrpc:"subject"`
	SubtypeId            Many2One  `xmlrpc:"subtype_id"`
	TrackingValueIds     []int64   `xmlrpc:"tracking_value_ids"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type MailMailNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	AttachmentIds        interface{} `xmlrpc:"attachment_ids"`
	AuthorAvatar         interface{} `xmlrpc:"author_avatar"`
	AuthorId             interface{} `xmlrpc:"author_id"`
	AutoDelete           bool        `xmlrpc:"auto_delete"`
	Body                 interface{} `xmlrpc:"body"`
	BodyHtml             interface{} `xmlrpc:"body_html"`
	ChannelIds           interface{} `xmlrpc:"channel_ids"`
	ChildIds             interface{} `xmlrpc:"child_ids"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	Date                 interface{} `xmlrpc:"date"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	EmailCc              interface{} `xmlrpc:"email_cc"`
	EmailFrom            interface{} `xmlrpc:"email_from"`
	EmailTo              interface{} `xmlrpc:"email_to"`
	FailureReason        interface{} `xmlrpc:"failure_reason"`
	FetchmailServerId    interface{} `xmlrpc:"fetchmail_server_id"`
	Headers              interface{} `xmlrpc:"headers"`
	Id                   interface{} `xmlrpc:"id"`
	MailMessageId        interface{} `xmlrpc:"mail_message_id"`
	MailServerId         interface{} `xmlrpc:"mail_server_id"`
	MailingId            interface{} `xmlrpc:"mailing_id"`
	MessageId            interface{} `xmlrpc:"message_id"`
	MessageType          interface{} `xmlrpc:"message_type"`
	Model                interface{} `xmlrpc:"model"`
	Needaction           bool        `xmlrpc:"needaction"`
	NeedactionPartnerIds interface{} `xmlrpc:"needaction_partner_ids"`
	NoAutoThread         bool        `xmlrpc:"no_auto_thread"`
	Notification         bool        `xmlrpc:"notification"`
	NotificationIds      interface{} `xmlrpc:"notification_ids"`
	ParentId             interface{} `xmlrpc:"parent_id"`
	PartnerIds           interface{} `xmlrpc:"partner_ids"`
	RecipientIds         interface{} `xmlrpc:"recipient_ids"`
	RecordName           interface{} `xmlrpc:"record_name"`
	References           interface{} `xmlrpc:"references"`
	ReplyTo              interface{} `xmlrpc:"reply_to"`
	ResId                interface{} `xmlrpc:"res_id"`
	ScheduledDate        interface{} `xmlrpc:"scheduled_date"`
	Starred              bool        `xmlrpc:"starred"`
	StarredPartnerIds    interface{} `xmlrpc:"starred_partner_ids"`
	State                interface{} `xmlrpc:"state"`
	StatisticsIds        interface{} `xmlrpc:"statistics_ids"`
	Subject              interface{} `xmlrpc:"subject"`
	SubtypeId            interface{} `xmlrpc:"subtype_id"`
	TrackingValueIds     interface{} `xmlrpc:"tracking_value_ids"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var MailMailModel string = "mail.mail"

type MailMails []MailMail

type MailMailsNil []MailMailNil

func (s *MailMail) NilableType_() interface{} {
	return &MailMailNil{}
}

func (ns *MailMailNil) Type_() interface{} {
	s := &MailMail{}
	return load(ns, s)
}

func (s *MailMails) NilableType_() interface{} {
	return &MailMailsNil{}
}

func (ns *MailMailsNil) Type_() interface{} {
	s := &MailMails{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMail))
	}
	return s
}
