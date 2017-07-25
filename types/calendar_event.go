package types

import (
	"time"
)

type CalendarEvent struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	Active                   bool      `xmlrpc:"active"`
	AlarmIds                 []int64   `xmlrpc:"alarm_ids"`
	Allday                   bool      `xmlrpc:"allday"`
	AttendeeIds              []int64   `xmlrpc:"attendee_ids"`
	AttendeeStatus           string    `xmlrpc:"attendee_status"`
	Byday                    string    `xmlrpc:"byday"`
	CategIds                 []int64   `xmlrpc:"categ_ids"`
	ColorPartnerId           int64     `xmlrpc:"color_partner_id"`
	Count                    int64     `xmlrpc:"count"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	Day                      int64     `xmlrpc:"day"`
	Description              string    `xmlrpc:"description"`
	DisplayName              string    `xmlrpc:"display_name"`
	DisplayStart             string    `xmlrpc:"display_start"`
	DisplayTime              string    `xmlrpc:"display_time"`
	Duration                 float64   `xmlrpc:"duration"`
	EndType                  string    `xmlrpc:"end_type"`
	FinalDate                time.Time `xmlrpc:"final_date"`
	Fr                       bool      `xmlrpc:"fr"`
	Id                       int64     `xmlrpc:"id"`
	Interval                 int64     `xmlrpc:"interval"`
	IsAttendee               bool      `xmlrpc:"is_attendee"`
	Location                 string    `xmlrpc:"location"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	Mo                       bool      `xmlrpc:"mo"`
	MonthBy                  string    `xmlrpc:"month_by"`
	Name                     string    `xmlrpc:"name"`
	OpportunityId            Many2One  `xmlrpc:"opportunity_id"`
	PartnerIds               []int64   `xmlrpc:"partner_ids"`
	Privacy                  string    `xmlrpc:"privacy"`
	Recurrency               bool      `xmlrpc:"recurrency"`
	RecurrentId              int64     `xmlrpc:"recurrent_id"`
	RecurrentIdDate          time.Time `xmlrpc:"recurrent_id_date"`
	Rrule                    string    `xmlrpc:"rrule"`
	RruleType                string    `xmlrpc:"rrule_type"`
	Sa                       bool      `xmlrpc:"sa"`
	ShowAs                   string    `xmlrpc:"show_as"`
	Start                    time.Time `xmlrpc:"start"`
	StartDate                time.Time `xmlrpc:"start_date"`
	StartDatetime            time.Time `xmlrpc:"start_datetime"`
	State                    string    `xmlrpc:"state"`
	Stop                     time.Time `xmlrpc:"stop"`
	StopDate                 time.Time `xmlrpc:"stop_date"`
	StopDatetime             time.Time `xmlrpc:"stop_datetime"`
	Su                       bool      `xmlrpc:"su"`
	Th                       bool      `xmlrpc:"th"`
	Tu                       bool      `xmlrpc:"tu"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	We                       bool      `xmlrpc:"we"`
	WeekList                 string    `xmlrpc:"week_list"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type CalendarEventNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	Active                   bool        `xmlrpc:"active"`
	AlarmIds                 interface{} `xmlrpc:"alarm_ids"`
	Allday                   bool        `xmlrpc:"allday"`
	AttendeeIds              interface{} `xmlrpc:"attendee_ids"`
	AttendeeStatus           interface{} `xmlrpc:"attendee_status"`
	Byday                    interface{} `xmlrpc:"byday"`
	CategIds                 interface{} `xmlrpc:"categ_ids"`
	ColorPartnerId           interface{} `xmlrpc:"color_partner_id"`
	Count                    interface{} `xmlrpc:"count"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Day                      interface{} `xmlrpc:"day"`
	Description              interface{} `xmlrpc:"description"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	DisplayStart             interface{} `xmlrpc:"display_start"`
	DisplayTime              interface{} `xmlrpc:"display_time"`
	Duration                 interface{} `xmlrpc:"duration"`
	EndType                  interface{} `xmlrpc:"end_type"`
	FinalDate                interface{} `xmlrpc:"final_date"`
	Fr                       bool        `xmlrpc:"fr"`
	Id                       interface{} `xmlrpc:"id"`
	Interval                 interface{} `xmlrpc:"interval"`
	IsAttendee               bool        `xmlrpc:"is_attendee"`
	Location                 interface{} `xmlrpc:"location"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	Mo                       bool        `xmlrpc:"mo"`
	MonthBy                  interface{} `xmlrpc:"month_by"`
	Name                     interface{} `xmlrpc:"name"`
	OpportunityId            interface{} `xmlrpc:"opportunity_id"`
	PartnerIds               interface{} `xmlrpc:"partner_ids"`
	Privacy                  interface{} `xmlrpc:"privacy"`
	Recurrency               bool        `xmlrpc:"recurrency"`
	RecurrentId              interface{} `xmlrpc:"recurrent_id"`
	RecurrentIdDate          interface{} `xmlrpc:"recurrent_id_date"`
	Rrule                    interface{} `xmlrpc:"rrule"`
	RruleType                interface{} `xmlrpc:"rrule_type"`
	Sa                       bool        `xmlrpc:"sa"`
	ShowAs                   interface{} `xmlrpc:"show_as"`
	Start                    interface{} `xmlrpc:"start"`
	StartDate                interface{} `xmlrpc:"start_date"`
	StartDatetime            interface{} `xmlrpc:"start_datetime"`
	State                    interface{} `xmlrpc:"state"`
	Stop                     interface{} `xmlrpc:"stop"`
	StopDate                 interface{} `xmlrpc:"stop_date"`
	StopDatetime             interface{} `xmlrpc:"stop_datetime"`
	Su                       bool        `xmlrpc:"su"`
	Th                       bool        `xmlrpc:"th"`
	Tu                       bool        `xmlrpc:"tu"`
	UserId                   interface{} `xmlrpc:"user_id"`
	We                       bool        `xmlrpc:"we"`
	WeekList                 interface{} `xmlrpc:"week_list"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var CalendarEventModel string = "calendar.event"

type CalendarEvents []CalendarEvent

type CalendarEventsNil []CalendarEventNil

func (s *CalendarEvent) NilableType_() interface{} {
	return &CalendarEventNil{}
}

func (ns *CalendarEventNil) Type_() interface{} {
	s := &CalendarEvent{}
	return load(ns, s)
}

func (s *CalendarEvents) NilableType_() interface{} {
	return &CalendarEventsNil{}
}

func (ns *CalendarEventsNil) Type_() interface{} {
	s := &CalendarEvents{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarEvent))
	}
	return s
}
