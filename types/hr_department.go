package types

import (
	"time"
)

type HrDepartment struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AbsenceOfToday           int64     `xmlrpc:"absence_of_today"`
	Active                   bool      `xmlrpc:"active"`
	AllocationToApproveCount int64     `xmlrpc:"allocation_to_approve_count"`
	ChildIds                 []int64   `xmlrpc:"child_ids"`
	Color                    int64     `xmlrpc:"color"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
	JobsIds                  []int64   `xmlrpc:"jobs_ids"`
	LeaveToApproveCount      int64     `xmlrpc:"leave_to_approve_count"`
	ManagerId                Many2One  `xmlrpc:"manager_id"`
	MemberIds                []int64   `xmlrpc:"member_ids"`
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
	Note                     string    `xmlrpc:"note"`
	ParentId                 Many2One  `xmlrpc:"parent_id"`
	TotalEmployee            int64     `xmlrpc:"total_employee"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type HrDepartmentNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AbsenceOfToday           interface{} `xmlrpc:"absence_of_today"`
	Active                   bool        `xmlrpc:"active"`
	AllocationToApproveCount interface{} `xmlrpc:"allocation_to_approve_count"`
	ChildIds                 interface{} `xmlrpc:"child_ids"`
	Color                    interface{} `xmlrpc:"color"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
	JobsIds                  interface{} `xmlrpc:"jobs_ids"`
	LeaveToApproveCount      interface{} `xmlrpc:"leave_to_approve_count"`
	ManagerId                interface{} `xmlrpc:"manager_id"`
	MemberIds                interface{} `xmlrpc:"member_ids"`
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
	Note                     interface{} `xmlrpc:"note"`
	ParentId                 interface{} `xmlrpc:"parent_id"`
	TotalEmployee            interface{} `xmlrpc:"total_employee"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var HrDepartmentModel string = "hr.department"

type HrDepartments []HrDepartment

type HrDepartmentsNil []HrDepartmentNil

func (s *HrDepartment) NilableType_() interface{} {
	return &HrDepartmentNil{}
}

func (ns *HrDepartmentNil) Type_() interface{} {
	s := &HrDepartment{}
	return load(ns, s)
}

func (s *HrDepartments) NilableType_() interface{} {
	return &HrDepartmentsNil{}
}

func (ns *HrDepartmentsNil) Type_() interface{} {
	s := &HrDepartments{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrDepartment))
	}
	return s
}
