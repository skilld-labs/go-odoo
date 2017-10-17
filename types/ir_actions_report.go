package types

import (
	"time"
)

type IrActionsReport struct {
	Attachment      string    `xmlrpc:"attachment"`
	AttachmentUse   bool      `xmlrpc:"attachment_use"`
	BindingModelId  Many2One  `xmlrpc:"binding_model_id"`
	BindingType     string    `xmlrpc:"binding_type"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	GroupsId        []int64   `xmlrpc:"groups_id"`
	Help            string    `xmlrpc:"help"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Model           string    `xmlrpc:"model"`
	Multi           bool      `xmlrpc:"multi"`
	Name            string    `xmlrpc:"name"`
	PaperformatId   Many2One  `xmlrpc:"paperformat_id"`
	PrintReportName string    `xmlrpc:"print_report_name"`
	ReportFile      string    `xmlrpc:"report_file"`
	ReportName      string    `xmlrpc:"report_name"`
	ReportType      string    `xmlrpc:"report_type"`
	Type            string    `xmlrpc:"type"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
	XmlId           string    `xmlrpc:"xml_id"`
}

type IrActionsReportNil struct {
	Attachment      interface{} `xmlrpc:"attachment"`
	AttachmentUse   bool        `xmlrpc:"attachment_use"`
	BindingModelId  interface{} `xmlrpc:"binding_model_id"`
	BindingType     interface{} `xmlrpc:"binding_type"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	GroupsId        interface{} `xmlrpc:"groups_id"`
	Help            interface{} `xmlrpc:"help"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Model           interface{} `xmlrpc:"model"`
	Multi           bool        `xmlrpc:"multi"`
	Name            interface{} `xmlrpc:"name"`
	PaperformatId   interface{} `xmlrpc:"paperformat_id"`
	PrintReportName interface{} `xmlrpc:"print_report_name"`
	ReportFile      interface{} `xmlrpc:"report_file"`
	ReportName      interface{} `xmlrpc:"report_name"`
	ReportType      interface{} `xmlrpc:"report_type"`
	Type            interface{} `xmlrpc:"type"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
	XmlId           interface{} `xmlrpc:"xml_id"`
}

var IrActionsReportModel string = "ir.actions.report"

type IrActionsReports []IrActionsReport

type IrActionsReportsNil []IrActionsReportNil

func (s *IrActionsReport) NilableType_() interface{} {
	return &IrActionsReportNil{}
}

func (ns *IrActionsReportNil) Type_() interface{} {
	s := &IrActionsReport{}
	return load(ns, s)
}

func (s *IrActionsReports) NilableType_() interface{} {
	return &IrActionsReportsNil{}
}

func (ns *IrActionsReportsNil) Type_() interface{} {
	s := &IrActionsReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsReport))
	}
	return s
}
