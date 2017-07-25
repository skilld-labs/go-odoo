package types

import (
	"time"
)

type HrHolidaysSummaryDept struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	Depts       []int64   `xmlrpc:"depts"`
	DisplayName string    `xmlrpc:"display_name"`
	HolidayType string    `xmlrpc:"holiday_type"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type HrHolidaysSummaryDeptNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	Depts       interface{} `xmlrpc:"depts"`
	DisplayName interface{} `xmlrpc:"display_name"`
	HolidayType interface{} `xmlrpc:"holiday_type"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var HrHolidaysSummaryDeptModel string = "hr.holidays.summary.dept"

type HrHolidaysSummaryDepts []HrHolidaysSummaryDept

type HrHolidaysSummaryDeptsNil []HrHolidaysSummaryDeptNil

func (s *HrHolidaysSummaryDept) NilableType_() interface{} {
	return &HrHolidaysSummaryDeptNil{}
}

func (ns *HrHolidaysSummaryDeptNil) Type_() interface{} {
	s := &HrHolidaysSummaryDept{}
	return load(ns, s)
}

func (s *HrHolidaysSummaryDepts) NilableType_() interface{} {
	return &HrHolidaysSummaryDeptsNil{}
}

func (ns *HrHolidaysSummaryDeptsNil) Type_() interface{} {
	s := &HrHolidaysSummaryDepts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrHolidaysSummaryDept))
	}
	return s
}
