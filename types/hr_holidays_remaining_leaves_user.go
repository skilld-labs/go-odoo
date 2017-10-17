package types

import (
	"time"
)

type HrHolidaysRemainingLeavesUser struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	LeaveType   string    `xmlrpc:"leave_type"`
	Name        string    `xmlrpc:"name"`
	NoOfLeaves  int64     `xmlrpc:"no_of_leaves"`
	UserId      Many2One  `xmlrpc:"user_id"`
}

type HrHolidaysRemainingLeavesUserNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	LeaveType   interface{} `xmlrpc:"leave_type"`
	Name        interface{} `xmlrpc:"name"`
	NoOfLeaves  interface{} `xmlrpc:"no_of_leaves"`
	UserId      interface{} `xmlrpc:"user_id"`
}

var HrHolidaysRemainingLeavesUserModel string = "hr.holidays.remaining.leaves.user"

type HrHolidaysRemainingLeavesUsers []HrHolidaysRemainingLeavesUser

type HrHolidaysRemainingLeavesUsersNil []HrHolidaysRemainingLeavesUserNil

func (s *HrHolidaysRemainingLeavesUser) NilableType_() interface{} {
	return &HrHolidaysRemainingLeavesUserNil{}
}

func (ns *HrHolidaysRemainingLeavesUserNil) Type_() interface{} {
	s := &HrHolidaysRemainingLeavesUser{}
	return load(ns, s)
}

func (s *HrHolidaysRemainingLeavesUsers) NilableType_() interface{} {
	return &HrHolidaysRemainingLeavesUsersNil{}
}

func (ns *HrHolidaysRemainingLeavesUsersNil) Type_() interface{} {
	s := &HrHolidaysRemainingLeavesUsers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrHolidaysRemainingLeavesUser))
	}
	return s
}
