package types

import (
	"time"
)

type MailComposeMessage struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	ActiveDomain          string    `xmlrpc:"active_domain"`
	AttachmentIds         []int64   `xmlrpc:"attachment_ids"`
	AuthorAvatar          string    `xmlrpc:"author_avatar"`
	AuthorId              Many2One  `xmlrpc:"author_id"`
	AutoDelete            bool      `xmlrpc:"auto_delete"`
	AutoDeleteMessage     bool      `xmlrpc:"auto_delete_message"`
	Body                  string    `xmlrpc:"body"`
	ChannelIds            []int64   `xmlrpc:"channel_ids"`
	ChildIds              []int64   `xmlrpc:"child_ids"`
	CompositionMode       string    `xmlrpc:"composition_mode"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	Date                  time.Time `xmlrpc:"date"`
	DisplayName           string    `xmlrpc:"display_name"`
	EmailFrom             string    `xmlrpc:"email_from"`
	Id                    int64     `xmlrpc:"id"`
	IsLog                 bool      `xmlrpc:"is_log"`
	MailServerId          Many2One  `xmlrpc:"mail_server_id"`
	MailingListIds        []int64   `xmlrpc:"mailing_list_ids"`
	MassMailingCampaignId Many2One  `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         Many2One  `xmlrpc:"mass_mailing_id"`
	MassMailingName       string    `xmlrpc:"mass_mailing_name"`
	MessageId             string    `xmlrpc:"message_id"`
	MessageType           string    `xmlrpc:"message_type"`
	Model                 string    `xmlrpc:"model"`
	Needaction            bool      `xmlrpc:"needaction"`
	NeedactionPartnerIds  []int64   `xmlrpc:"needaction_partner_ids"`
	NoAutoThread          bool      `xmlrpc:"no_auto_thread"`
	NotificationIds       []int64   `xmlrpc:"notification_ids"`
	Notify                bool      `xmlrpc:"notify"`
	ParentId              Many2One  `xmlrpc:"parent_id"`
	PartnerIds            []int64   `xmlrpc:"partner_ids"`
	RecordName            string    `xmlrpc:"record_name"`
	ReplyTo               string    `xmlrpc:"reply_to"`
	ResId                 int64     `xmlrpc:"res_id"`
	Starred               bool      `xmlrpc:"starred"`
	StarredPartnerIds     []int64   `xmlrpc:"starred_partner_ids"`
	Subject               string    `xmlrpc:"subject"`
	SubtypeId             Many2One  `xmlrpc:"subtype_id"`
	TemplateId            Many2One  `xmlrpc:"template_id"`
	TrackingValueIds      []int64   `xmlrpc:"tracking_value_ids"`
	UseActiveDomain       bool      `xmlrpc:"use_active_domain"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type MailComposeMessageNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	ActiveDomain          interface{} `xmlrpc:"active_domain"`
	AttachmentIds         interface{} `xmlrpc:"attachment_ids"`
	AuthorAvatar          interface{} `xmlrpc:"author_avatar"`
	AuthorId              interface{} `xmlrpc:"author_id"`
	AutoDelete            bool        `xmlrpc:"auto_delete"`
	AutoDeleteMessage     bool        `xmlrpc:"auto_delete_message"`
	Body                  interface{} `xmlrpc:"body"`
	ChannelIds            interface{} `xmlrpc:"channel_ids"`
	ChildIds              interface{} `xmlrpc:"child_ids"`
	CompositionMode       interface{} `xmlrpc:"composition_mode"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	Date                  interface{} `xmlrpc:"date"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	EmailFrom             interface{} `xmlrpc:"email_from"`
	Id                    interface{} `xmlrpc:"id"`
	IsLog                 bool        `xmlrpc:"is_log"`
	MailServerId          interface{} `xmlrpc:"mail_server_id"`
	MailingListIds        interface{} `xmlrpc:"mailing_list_ids"`
	MassMailingCampaignId interface{} `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         interface{} `xmlrpc:"mass_mailing_id"`
	MassMailingName       interface{} `xmlrpc:"mass_mailing_name"`
	MessageId             interface{} `xmlrpc:"message_id"`
	MessageType           interface{} `xmlrpc:"message_type"`
	Model                 interface{} `xmlrpc:"model"`
	Needaction            bool        `xmlrpc:"needaction"`
	NeedactionPartnerIds  interface{} `xmlrpc:"needaction_partner_ids"`
	NoAutoThread          bool        `xmlrpc:"no_auto_thread"`
	NotificationIds       interface{} `xmlrpc:"notification_ids"`
	Notify                bool        `xmlrpc:"notify"`
	ParentId              interface{} `xmlrpc:"parent_id"`
	PartnerIds            interface{} `xmlrpc:"partner_ids"`
	RecordName            interface{} `xmlrpc:"record_name"`
	ReplyTo               interface{} `xmlrpc:"reply_to"`
	ResId                 interface{} `xmlrpc:"res_id"`
	Starred               bool        `xmlrpc:"starred"`
	StarredPartnerIds     interface{} `xmlrpc:"starred_partner_ids"`
	Subject               interface{} `xmlrpc:"subject"`
	SubtypeId             interface{} `xmlrpc:"subtype_id"`
	TemplateId            interface{} `xmlrpc:"template_id"`
	TrackingValueIds      interface{} `xmlrpc:"tracking_value_ids"`
	UseActiveDomain       bool        `xmlrpc:"use_active_domain"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var MailComposeMessageModel string = "mail.compose.message"

type MailComposeMessages []MailComposeMessage

type MailComposeMessagesNil []MailComposeMessageNil

func (s *MailComposeMessage) NilableType_() interface{} {
	return &MailComposeMessageNil{}
}

func (ns *MailComposeMessageNil) Type_() interface{} {
	s := &MailComposeMessage{}
	return load(ns, s)
}

func (s *MailComposeMessages) NilableType_() interface{} {
	return &MailComposeMessagesNil{}
}

func (ns *MailComposeMessagesNil) Type_() interface{} {
	s := &MailComposeMessages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailComposeMessage))
	}
	return s
}
