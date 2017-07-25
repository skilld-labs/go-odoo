package types

import (
	"time"
)

type IrActionsReportXml struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	Attachment           string    `xmlrpc:"attachment"`
	AttachmentUse        bool      `xmlrpc:"attachment_use"`
	Auto                 bool      `xmlrpc:"auto"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	DisplayName          string    `xmlrpc:"display_name"`
	GroupsId             []int64   `xmlrpc:"groups_id"`
	Header               bool      `xmlrpc:"header"`
	Help                 string    `xmlrpc:"help"`
	Id                   int64     `xmlrpc:"id"`
	IrValuesId           Many2One  `xmlrpc:"ir_values_id"`
	Model                string    `xmlrpc:"model"`
	Multi                bool      `xmlrpc:"multi"`
	Name                 string    `xmlrpc:"name"`
	PaperformatId        Many2One  `xmlrpc:"paperformat_id"`
	Parser               string    `xmlrpc:"parser"`
	PrintReportName      string    `xmlrpc:"print_report_name"`
	ReportFile           string    `xmlrpc:"report_file"`
	ReportName           string    `xmlrpc:"report_name"`
	ReportRml            string    `xmlrpc:"report_rml"`
	ReportRmlContent     string    `xmlrpc:"report_rml_content"`
	ReportRmlContentData string    `xmlrpc:"report_rml_content_data"`
	ReportSxw            string    `xmlrpc:"report_sxw"`
	ReportSxwContent     string    `xmlrpc:"report_sxw_content"`
	ReportSxwContentData string    `xmlrpc:"report_sxw_content_data"`
	ReportType           string    `xmlrpc:"report_type"`
	ReportXml            string    `xmlrpc:"report_xml"`
	ReportXsl            string    `xmlrpc:"report_xsl"`
	Type                 string    `xmlrpc:"type"`
	Usage                string    `xmlrpc:"usage"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
	XmlId                string    `xmlrpc:"xml_id"`
}

type IrActionsReportXmlNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	Attachment           interface{} `xmlrpc:"attachment"`
	AttachmentUse        bool        `xmlrpc:"attachment_use"`
	Auto                 bool        `xmlrpc:"auto"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	GroupsId             interface{} `xmlrpc:"groups_id"`
	Header               bool        `xmlrpc:"header"`
	Help                 interface{} `xmlrpc:"help"`
	Id                   interface{} `xmlrpc:"id"`
	IrValuesId           interface{} `xmlrpc:"ir_values_id"`
	Model                interface{} `xmlrpc:"model"`
	Multi                bool        `xmlrpc:"multi"`
	Name                 interface{} `xmlrpc:"name"`
	PaperformatId        interface{} `xmlrpc:"paperformat_id"`
	Parser               interface{} `xmlrpc:"parser"`
	PrintReportName      interface{} `xmlrpc:"print_report_name"`
	ReportFile           interface{} `xmlrpc:"report_file"`
	ReportName           interface{} `xmlrpc:"report_name"`
	ReportRml            interface{} `xmlrpc:"report_rml"`
	ReportRmlContent     interface{} `xmlrpc:"report_rml_content"`
	ReportRmlContentData interface{} `xmlrpc:"report_rml_content_data"`
	ReportSxw            interface{} `xmlrpc:"report_sxw"`
	ReportSxwContent     interface{} `xmlrpc:"report_sxw_content"`
	ReportSxwContentData interface{} `xmlrpc:"report_sxw_content_data"`
	ReportType           interface{} `xmlrpc:"report_type"`
	ReportXml            interface{} `xmlrpc:"report_xml"`
	ReportXsl            interface{} `xmlrpc:"report_xsl"`
	Type                 interface{} `xmlrpc:"type"`
	Usage                interface{} `xmlrpc:"usage"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
	XmlId                interface{} `xmlrpc:"xml_id"`
}

var IrActionsReportXmlModel string = "ir.actions.report.xml"

type IrActionsReportXmls []IrActionsReportXml

type IrActionsReportXmlsNil []IrActionsReportXmlNil

func (s *IrActionsReportXml) NilableType_() interface{} {
	return &IrActionsReportXmlNil{}
}

func (ns *IrActionsReportXmlNil) Type_() interface{} {
	s := &IrActionsReportXml{}
	return load(ns, s)
}

func (s *IrActionsReportXmls) NilableType_() interface{} {
	return &IrActionsReportXmlsNil{}
}

func (ns *IrActionsReportXmlsNil) Type_() interface{} {
	s := &IrActionsReportXmls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsReportXml))
	}
	return s
}
