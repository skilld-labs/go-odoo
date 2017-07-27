package types

import (
	"time"
)

type HrHolidays struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	CanReset                 bool      `xmlrpc:"can_reset"`
	CategoryId               Many2One  `xmlrpc:"category_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DateFrom                 time.Time `xmlrpc:"date_from"`
	DateTo                   time.Time `xmlrpc:"date_to"`
	DepartmentId             Many2One  `xmlrpc:"department_id"`
	DisplayName              string    `xmlrpc:"display_name"`
	DoubleValidation         bool      `xmlrpc:"double_validation"`
	EmployeeId               Many2One  `xmlrpc:"employee_id"`
	HolidayStatusId          Many2One  `xmlrpc:"holiday_status_id"`
	HolidayType              string    `xmlrpc:"holiday_type"`
	Id                       int64     `xmlrpc:"id"`
	LinkedRequestIds         []int64   `xmlrpc:"linked_request_ids"`
	ManagerId                Many2One  `xmlrpc:"manager_id"`
	ManagerId2               Many2One  `xmlrpc:"manager_id2"`
	MeetingId                Many2One  `xmlrpc:"meeting_id"`
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
	Name                     string    `xmlrpc:"name"`
	Notes                    string    `xmlrpc:"notes"`
	NumberOfDays             float64   `xmlrpc:"number_of_days"`
	NumberOfDaysTemp         float64   `xmlrpc:"number_of_days_temp"`
	ParentId                 Many2One  `xmlrpc:"parent_id"`
	PayslipStatus            bool      `xmlrpc:"payslip_status"`
	ReportNote               string    `xmlrpc:"report_note"`
	State                    string    `xmlrpc:"state"`
	Type                     string    `xmlrpc:"type"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type HrHolidaysNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	CanReset                 bool        `xmlrpc:"can_reset"`
	CategoryId               interface{} `xmlrpc:"category_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DateFrom                 interface{} `xmlrpc:"date_from"`
	DateTo                   interface{} `xmlrpc:"date_to"`
	DepartmentId             interface{} `xmlrpc:"department_id"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	DoubleValidation         bool        `xmlrpc:"double_validation"`
	EmployeeId               interface{} `xmlrpc:"employee_id"`
	HolidayStatusId          interface{} `xmlrpc:"holiday_status_id"`
	HolidayType              interface{} `xmlrpc:"holiday_type"`
	Id                       interface{} `xmlrpc:"id"`
	LinkedRequestIds         interface{} `xmlrpc:"linked_request_ids"`
	ManagerId                interface{} `xmlrpc:"manager_id"`
	ManagerId2               interface{} `xmlrpc:"manager_id2"`
	MeetingId                interface{} `xmlrpc:"meeting_id"`
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
	Name                     interface{} `xmlrpc:"name"`
	Notes                    interface{} `xmlrpc:"notes"`
	NumberOfDays             interface{} `xmlrpc:"number_of_days"`
	NumberOfDaysTemp         interface{} `xmlrpc:"number_of_days_temp"`
	ParentId                 interface{} `xmlrpc:"parent_id"`
	PayslipStatus            bool        `xmlrpc:"payslip_status"`
	ReportNote               interface{} `xmlrpc:"report_note"`
	State                    interface{} `xmlrpc:"state"`
	Type                     interface{} `xmlrpc:"type"`
	UserId                   interface{} `xmlrpc:"user_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var HrHolidaysModel string = "hr.holidays"

type HrHolidayss []HrHolidays

type HrHolidayssNil []HrHolidaysNil

func (s *HrHolidays) NilableType_() interface{} {
	return &HrHolidaysNil{}
}

func (ns *HrHolidaysNil) Type_() interface{} {
	s := &HrHolidays{}
	return load(ns, s)
}

func (s *HrHolidayss) NilableType_() interface{} {
	return &HrHolidayssNil{}
}

func (ns *HrHolidayssNil) Type_() interface{} {
	s := &HrHolidayss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrHolidays))
	}
	return s
}
