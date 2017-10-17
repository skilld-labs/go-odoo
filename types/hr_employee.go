package types

import (
	"time"
)

type HrEmployee struct {
	Active                   bool      `xmlrpc:"active"`
	AddressHomeId            Many2One  `xmlrpc:"address_home_id"`
	AddressId                Many2One  `xmlrpc:"address_id"`
	BankAccountId            Many2One  `xmlrpc:"bank_account_id"`
	Birthday                 time.Time `xmlrpc:"birthday"`
	CategoryIds              []int64   `xmlrpc:"category_ids"`
	ChildIds                 []int64   `xmlrpc:"child_ids"`
	CoachId                  Many2One  `xmlrpc:"coach_id"`
	Color                    int64     `xmlrpc:"color"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CountryId                Many2One  `xmlrpc:"country_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	CurrentLeaveId           Many2One  `xmlrpc:"current_leave_id"`
	CurrentLeaveState        string    `xmlrpc:"current_leave_state"`
	DepartmentId             Many2One  `xmlrpc:"department_id"`
	DisplayName              string    `xmlrpc:"display_name"`
	Gender                   string    `xmlrpc:"gender"`
	Id                       int64     `xmlrpc:"id"`
	IdentificationId         string    `xmlrpc:"identification_id"`
	Image                    string    `xmlrpc:"image"`
	ImageMedium              string    `xmlrpc:"image_medium"`
	ImageSmall               string    `xmlrpc:"image_small"`
	IsAbsentTotay            bool      `xmlrpc:"is_absent_totay"`
	IsAddressHomeACompany    bool      `xmlrpc:"is_address_home_a_company"`
	JobId                    Many2One  `xmlrpc:"job_id"`
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	LeaveDateFrom            time.Time `xmlrpc:"leave_date_from"`
	LeaveDateTo              time.Time `xmlrpc:"leave_date_to"`
	LeavesCount              float64   `xmlrpc:"leaves_count"`
	Marital                  string    `xmlrpc:"marital"`
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
	MobilePhone              string    `xmlrpc:"mobile_phone"`
	Name                     string    `xmlrpc:"name"`
	Notes                    string    `xmlrpc:"notes"`
	ParentId                 Many2One  `xmlrpc:"parent_id"`
	PassportId               string    `xmlrpc:"passport_id"`
	PermitNo                 string    `xmlrpc:"permit_no"`
	RemainingLeaves          float64   `xmlrpc:"remaining_leaves"`
	ResourceCalendarId       Many2One  `xmlrpc:"resource_calendar_id"`
	ResourceId               Many2One  `xmlrpc:"resource_id"`
	ShowLeaves               bool      `xmlrpc:"show_leaves"`
	Sinid                    string    `xmlrpc:"sinid"`
	Ssnid                    string    `xmlrpc:"ssnid"`
	TimesheetCost            float64   `xmlrpc:"timesheet_cost"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	VisaExpire               time.Time `xmlrpc:"visa_expire"`
	VisaNo                   string    `xmlrpc:"visa_no"`
	WebsiteMessageIds        []int64   `xmlrpc:"website_message_ids"`
	WorkEmail                string    `xmlrpc:"work_email"`
	WorkLocation             string    `xmlrpc:"work_location"`
	WorkPhone                string    `xmlrpc:"work_phone"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type HrEmployeeNil struct {
	Active                   bool        `xmlrpc:"active"`
	AddressHomeId            interface{} `xmlrpc:"address_home_id"`
	AddressId                interface{} `xmlrpc:"address_id"`
	BankAccountId            interface{} `xmlrpc:"bank_account_id"`
	Birthday                 interface{} `xmlrpc:"birthday"`
	CategoryIds              interface{} `xmlrpc:"category_ids"`
	ChildIds                 interface{} `xmlrpc:"child_ids"`
	CoachId                  interface{} `xmlrpc:"coach_id"`
	Color                    interface{} `xmlrpc:"color"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CountryId                interface{} `xmlrpc:"country_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	CurrentLeaveId           interface{} `xmlrpc:"current_leave_id"`
	CurrentLeaveState        interface{} `xmlrpc:"current_leave_state"`
	DepartmentId             interface{} `xmlrpc:"department_id"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Gender                   interface{} `xmlrpc:"gender"`
	Id                       interface{} `xmlrpc:"id"`
	IdentificationId         interface{} `xmlrpc:"identification_id"`
	Image                    interface{} `xmlrpc:"image"`
	ImageMedium              interface{} `xmlrpc:"image_medium"`
	ImageSmall               interface{} `xmlrpc:"image_small"`
	IsAbsentTotay            bool        `xmlrpc:"is_absent_totay"`
	IsAddressHomeACompany    bool        `xmlrpc:"is_address_home_a_company"`
	JobId                    interface{} `xmlrpc:"job_id"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	LeaveDateFrom            interface{} `xmlrpc:"leave_date_from"`
	LeaveDateTo              interface{} `xmlrpc:"leave_date_to"`
	LeavesCount              interface{} `xmlrpc:"leaves_count"`
	Marital                  interface{} `xmlrpc:"marital"`
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
	MobilePhone              interface{} `xmlrpc:"mobile_phone"`
	Name                     interface{} `xmlrpc:"name"`
	Notes                    interface{} `xmlrpc:"notes"`
	ParentId                 interface{} `xmlrpc:"parent_id"`
	PassportId               interface{} `xmlrpc:"passport_id"`
	PermitNo                 interface{} `xmlrpc:"permit_no"`
	RemainingLeaves          interface{} `xmlrpc:"remaining_leaves"`
	ResourceCalendarId       interface{} `xmlrpc:"resource_calendar_id"`
	ResourceId               interface{} `xmlrpc:"resource_id"`
	ShowLeaves               bool        `xmlrpc:"show_leaves"`
	Sinid                    interface{} `xmlrpc:"sinid"`
	Ssnid                    interface{} `xmlrpc:"ssnid"`
	TimesheetCost            interface{} `xmlrpc:"timesheet_cost"`
	UserId                   interface{} `xmlrpc:"user_id"`
	VisaExpire               interface{} `xmlrpc:"visa_expire"`
	VisaNo                   interface{} `xmlrpc:"visa_no"`
	WebsiteMessageIds        interface{} `xmlrpc:"website_message_ids"`
	WorkEmail                interface{} `xmlrpc:"work_email"`
	WorkLocation             interface{} `xmlrpc:"work_location"`
	WorkPhone                interface{} `xmlrpc:"work_phone"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var HrEmployeeModel string = "hr.employee"

type HrEmployees []HrEmployee

type HrEmployeesNil []HrEmployeeNil

func (s *HrEmployee) NilableType_() interface{} {
	return &HrEmployeeNil{}
}

func (ns *HrEmployeeNil) Type_() interface{} {
	s := &HrEmployee{}
	return load(ns, s)
}

func (s *HrEmployees) NilableType_() interface{} {
	return &HrEmployeesNil{}
}

func (ns *HrEmployeesNil) Type_() interface{} {
	s := &HrEmployees{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrEmployee))
	}
	return s
}
