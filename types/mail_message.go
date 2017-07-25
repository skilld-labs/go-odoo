package types

import (
	"time"
)

type MailMessage struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	AttachmentIds        []int64   `xmlrpc:"attachment_ids"`
	AuthorAvatar         string    `xmlrpc:"author_avatar"`
	AuthorId             Many2One  `xmlrpc:"author_id"`
	Body                 string    `xmlrpc:"body"`
	ChannelIds           []int64   `xmlrpc:"channel_ids"`
	ChildIds             []int64   `xmlrpc:"child_ids"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	Date                 time.Time `xmlrpc:"date"`
	DisplayName          string    `xmlrpc:"display_name"`
	EmailFrom            string    `xmlrpc:"email_from"`
	Id                   int64     `xmlrpc:"id"`
	MailServerId         Many2One  `xmlrpc:"mail_server_id"`
	MessageId            string    `xmlrpc:"message_id"`
	MessageType          string    `xmlrpc:"message_type"`
	Model                string    `xmlrpc:"model"`
	Needaction           bool      `xmlrpc:"needaction"`
	NeedactionPartnerIds []int64   `xmlrpc:"needaction_partner_ids"`
	NoAutoThread         bool      `xmlrpc:"no_auto_thread"`
	NotificationIds      []int64   `xmlrpc:"notification_ids"`
	ParentId             Many2One  `xmlrpc:"parent_id"`
	PartnerIds           []int64   `xmlrpc:"partner_ids"`
	RecordName           string    `xmlrpc:"record_name"`
	ReplyTo              string    `xmlrpc:"reply_to"`
	ResId                int64     `xmlrpc:"res_id"`
	Starred              bool      `xmlrpc:"starred"`
	StarredPartnerIds    []int64   `xmlrpc:"starred_partner_ids"`
	Subject              string    `xmlrpc:"subject"`
	SubtypeId            Many2One  `xmlrpc:"subtype_id"`
	TrackingValueIds     []int64   `xmlrpc:"tracking_value_ids"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type MailMessageNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	AttachmentIds        interface{} `xmlrpc:"attachment_ids"`
	AuthorAvatar         interface{} `xmlrpc:"author_avatar"`
	AuthorId             interface{} `xmlrpc:"author_id"`
	Body                 interface{} `xmlrpc:"body"`
	ChannelIds           interface{} `xmlrpc:"channel_ids"`
	ChildIds             interface{} `xmlrpc:"child_ids"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	Date                 interface{} `xmlrpc:"date"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	EmailFrom            interface{} `xmlrpc:"email_from"`
	Id                   interface{} `xmlrpc:"id"`
	MailServerId         interface{} `xmlrpc:"mail_server_id"`
	MessageId            interface{} `xmlrpc:"message_id"`
	MessageType          interface{} `xmlrpc:"message_type"`
	Model                interface{} `xmlrpc:"model"`
	Needaction           bool        `xmlrpc:"needaction"`
	NeedactionPartnerIds interface{} `xmlrpc:"needaction_partner_ids"`
	NoAutoThread         bool        `xmlrpc:"no_auto_thread"`
	NotificationIds      interface{} `xmlrpc:"notification_ids"`
	ParentId             interface{} `xmlrpc:"parent_id"`
	PartnerIds           interface{} `xmlrpc:"partner_ids"`
	RecordName           interface{} `xmlrpc:"record_name"`
	ReplyTo              interface{} `xmlrpc:"reply_to"`
	ResId                interface{} `xmlrpc:"res_id"`
	Starred              bool        `xmlrpc:"starred"`
	StarredPartnerIds    interface{} `xmlrpc:"starred_partner_ids"`
	Subject              interface{} `xmlrpc:"subject"`
	SubtypeId            interface{} `xmlrpc:"subtype_id"`
	TrackingValueIds     interface{} `xmlrpc:"tracking_value_ids"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var MailMessageModel string = "mail.message"

type MailMessages []MailMessage

type MailMessagesNil []MailMessageNil

func (s *MailMessage) NilableType_() interface{} {
	return &MailMessageNil{}
}

func (ns *MailMessageNil) Type_() interface{} {
	s := &MailMessage{}
	return load(ns, s)
}

func (s *MailMessages) NilableType_() interface{} {
	return &MailMessagesNil{}
}

func (ns *MailMessagesNil) Type_() interface{} {
	s := &MailMessages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMessage))
	}
	return s
}
