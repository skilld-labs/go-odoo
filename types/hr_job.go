package types

import (
	"time"
)

type HrJob struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DepartmentId             Many2One  `xmlrpc:"department_id"`
	Description              string    `xmlrpc:"description"`
	DisplayName              string    `xmlrpc:"display_name"`
	EmployeeIds              []int64   `xmlrpc:"employee_ids"`
	ExpectedEmployees        int64     `xmlrpc:"expected_employees"`
	Id                       int64     `xmlrpc:"id"`
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
	NoOfEmployee             int64     `xmlrpc:"no_of_employee"`
	NoOfHiredEmployee        int64     `xmlrpc:"no_of_hired_employee"`
	NoOfRecruitment          int64     `xmlrpc:"no_of_recruitment"`
	Requirements             string    `xmlrpc:"requirements"`
	State                    string    `xmlrpc:"state"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type HrJobNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DepartmentId             interface{} `xmlrpc:"department_id"`
	Description              interface{} `xmlrpc:"description"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	EmployeeIds              interface{} `xmlrpc:"employee_ids"`
	ExpectedEmployees        interface{} `xmlrpc:"expected_employees"`
	Id                       interface{} `xmlrpc:"id"`
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
	NoOfEmployee             interface{} `xmlrpc:"no_of_employee"`
	NoOfHiredEmployee        interface{} `xmlrpc:"no_of_hired_employee"`
	NoOfRecruitment          interface{} `xmlrpc:"no_of_recruitment"`
	Requirements             interface{} `xmlrpc:"requirements"`
	State                    interface{} `xmlrpc:"state"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var HrJobModel string = "hr.job"

type HrJobs []HrJob

type HrJobsNil []HrJobNil

func (s *HrJob) NilableType_() interface{} {
	return &HrJobNil{}
}

func (ns *HrJobNil) Type_() interface{} {
	s := &HrJob{}
	return load(ns, s)
}

func (s *HrJobs) NilableType_() interface{} {
	return &HrJobsNil{}
}

func (ns *HrJobsNil) Type_() interface{} {
	s := &HrJobs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrJob))
	}
	return s
}
