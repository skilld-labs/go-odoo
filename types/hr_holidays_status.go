package types

import (
	"time"
)

type HrHolidaysStatus struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Active                 bool      `xmlrpc:"active"`
	CategId                Many2One  `xmlrpc:"categ_id"`
	ColorName              string    `xmlrpc:"color_name"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	DisplayName            string    `xmlrpc:"display_name"`
	DoubleValidation       bool      `xmlrpc:"double_validation"`
	Id                     int64     `xmlrpc:"id"`
	LeavesTaken            float64   `xmlrpc:"leaves_taken"`
	Limit                  bool      `xmlrpc:"limit"`
	MaxLeaves              float64   `xmlrpc:"max_leaves"`
	Name                   string    `xmlrpc:"name"`
	RemainingLeaves        float64   `xmlrpc:"remaining_leaves"`
	VirtualRemainingLeaves float64   `xmlrpc:"virtual_remaining_leaves"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type HrHolidaysStatusNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Active                 bool        `xmlrpc:"active"`
	CategId                interface{} `xmlrpc:"categ_id"`
	ColorName              interface{} `xmlrpc:"color_name"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	DoubleValidation       bool        `xmlrpc:"double_validation"`
	Id                     interface{} `xmlrpc:"id"`
	LeavesTaken            interface{} `xmlrpc:"leaves_taken"`
	Limit                  bool        `xmlrpc:"limit"`
	MaxLeaves              interface{} `xmlrpc:"max_leaves"`
	Name                   interface{} `xmlrpc:"name"`
	RemainingLeaves        interface{} `xmlrpc:"remaining_leaves"`
	VirtualRemainingLeaves interface{} `xmlrpc:"virtual_remaining_leaves"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var HrHolidaysStatusModel string = "hr.holidays.status"

type HrHolidaysStatuss []HrHolidaysStatus

type HrHolidaysStatussNil []HrHolidaysStatusNil

func (s *HrHolidaysStatus) NilableType_() interface{} {
	return &HrHolidaysStatusNil{}
}

func (ns *HrHolidaysStatusNil) Type_() interface{} {
	s := &HrHolidaysStatus{}
	return load(ns, s)
}

func (s *HrHolidaysStatuss) NilableType_() interface{} {
	return &HrHolidaysStatussNil{}
}

func (ns *HrHolidaysStatussNil) Type_() interface{} {
	s := &HrHolidaysStatuss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*HrHolidaysStatus))
	}
	return s
}
