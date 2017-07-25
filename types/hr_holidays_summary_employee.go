package types

import (
	"time"
)

type HrHolidaysSummaryEmployee struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DisplayName string    `xmlrpc:"display_name"`
	Emp         []int64   `xmlrpc:"emp"`
	HolidayType string    `xmlrpc:"holiday_type"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type HrHolidaysSummaryEmployeeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Emp         interface{} `xmlrpc:"emp"`
	HolidayType interface{} `xmlrpc:"holiday_type"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var HrHolidaysSummaryEmployeeModel string = "hr.holidays.summary.employee"

type HrHolidaysSummaryEmployees []HrHolidaysSummaryEmployee

type HrHolidaysSummaryEmployeesNil []HrHolidaysSummaryEmployeeNil

func (s *HrHolidaysSummaryEmployee) NilableType_() interface{} {
	return &HrHolidaysSummaryEmployeeNil{}
}

func (ns *HrHolidaysSummaryEmployeeNil) Type_() interface{} {
	s := &HrHolidaysSummaryEmployee{}
	return load(ns, s)
}

func (s *HrHolidaysSummaryEmployees) NilableType_() interface{} {
	return &HrHolidaysSummaryEmployeesNil{}
}

func (ns *HrHolidaysSummaryEmployeesNil) Type_() interface{} {
	s := &HrHolidaysSummaryEmployees{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrHolidaysSummaryEmployee))
	}
	return s
}
