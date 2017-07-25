package types

import (
	"time"
)

type HrEmployeeCategory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	EmployeeIds []int64   `xmlrpc:"employee_ids"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type HrEmployeeCategoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	EmployeeIds interface{} `xmlrpc:"employee_ids"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var HrEmployeeCategoryModel string = "hr.employee.category"

type HrEmployeeCategorys []HrEmployeeCategory

type HrEmployeeCategorysNil []HrEmployeeCategoryNil

func (s *HrEmployeeCategory) NilableType_() interface{} {
	return &HrEmployeeCategoryNil{}
}

func (ns *HrEmployeeCategoryNil) Type_() interface{} {
	s := &HrEmployeeCategory{}
	return load(ns, s)
}

func (s *HrEmployeeCategorys) NilableType_() interface{} {
	return &HrEmployeeCategorysNil{}
}

func (ns *HrEmployeeCategorysNil) Type_() interface{} {
	s := &HrEmployeeCategorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrEmployeeCategory))
	}
	return s
}
