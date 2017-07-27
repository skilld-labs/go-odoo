package types

import (
	"time"
)

type MailTrackingValue struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	Field            string    `xmlrpc:"field"`
	FieldDesc        string    `xmlrpc:"field_desc"`
	FieldType        string    `xmlrpc:"field_type"`
	Id               int64     `xmlrpc:"id"`
	MailMessageId    Many2One  `xmlrpc:"mail_message_id"`
	NewValueChar     string    `xmlrpc:"new_value_char"`
	NewValueDatetime time.Time `xmlrpc:"new_value_datetime"`
	NewValueFloat    float64   `xmlrpc:"new_value_float"`
	NewValueInteger  int64     `xmlrpc:"new_value_integer"`
	NewValueMonetary float64   `xmlrpc:"new_value_monetary"`
	NewValueText     string    `xmlrpc:"new_value_text"`
	OldValueChar     string    `xmlrpc:"old_value_char"`
	OldValueDatetime time.Time `xmlrpc:"old_value_datetime"`
	OldValueFloat    float64   `xmlrpc:"old_value_float"`
	OldValueInteger  int64     `xmlrpc:"old_value_integer"`
	OldValueMonetary float64   `xmlrpc:"old_value_monetary"`
	OldValueText     string    `xmlrpc:"old_value_text"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type MailTrackingValueNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Field            interface{} `xmlrpc:"field"`
	FieldDesc        interface{} `xmlrpc:"field_desc"`
	FieldType        interface{} `xmlrpc:"field_type"`
	Id               interface{} `xmlrpc:"id"`
	MailMessageId    interface{} `xmlrpc:"mail_message_id"`
	NewValueChar     interface{} `xmlrpc:"new_value_char"`
	NewValueDatetime interface{} `xmlrpc:"new_value_datetime"`
	NewValueFloat    interface{} `xmlrpc:"new_value_float"`
	NewValueInteger  interface{} `xmlrpc:"new_value_integer"`
	NewValueMonetary interface{} `xmlrpc:"new_value_monetary"`
	NewValueText     interface{} `xmlrpc:"new_value_text"`
	OldValueChar     interface{} `xmlrpc:"old_value_char"`
	OldValueDatetime interface{} `xmlrpc:"old_value_datetime"`
	OldValueFloat    interface{} `xmlrpc:"old_value_float"`
	OldValueInteger  interface{} `xmlrpc:"old_value_integer"`
	OldValueMonetary interface{} `xmlrpc:"old_value_monetary"`
	OldValueText     interface{} `xmlrpc:"old_value_text"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var MailTrackingValueModel string = "mail.tracking.value"

type MailTrackingValues []MailTrackingValue

type MailTrackingValuesNil []MailTrackingValueNil

func (s *MailTrackingValue) NilableType_() interface{} {
	return &MailTrackingValueNil{}
}

func (ns *MailTrackingValueNil) Type_() interface{} {
	s := &MailTrackingValue{}
	return load(ns, s)
}

func (s *MailTrackingValues) NilableType_() interface{} {
	return &MailTrackingValuesNil{}
}

func (ns *MailTrackingValuesNil) Type_() interface{} {
	s := &MailTrackingValues{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailTrackingValue))
	}
	return s
}
