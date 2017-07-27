package types

import (
	"time"
)

type ReportPaperformat struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	Default       bool      `xmlrpc:"default"`
	DisplayName   string    `xmlrpc:"display_name"`
	Dpi           int64     `xmlrpc:"dpi"`
	Format        string    `xmlrpc:"format"`
	HeaderLine    bool      `xmlrpc:"header_line"`
	HeaderSpacing int64     `xmlrpc:"header_spacing"`
	Id            int64     `xmlrpc:"id"`
	MarginBottom  float64   `xmlrpc:"margin_bottom"`
	MarginLeft    float64   `xmlrpc:"margin_left"`
	MarginRight   float64   `xmlrpc:"margin_right"`
	MarginTop     float64   `xmlrpc:"margin_top"`
	Name          string    `xmlrpc:"name"`
	Orientation   string    `xmlrpc:"orientation"`
	PageHeight    int64     `xmlrpc:"page_height"`
	PageWidth     int64     `xmlrpc:"page_width"`
	ReportIds     []int64   `xmlrpc:"report_ids"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type ReportPaperformatNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	Default       bool        `xmlrpc:"default"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Dpi           interface{} `xmlrpc:"dpi"`
	Format        interface{} `xmlrpc:"format"`
	HeaderLine    bool        `xmlrpc:"header_line"`
	HeaderSpacing interface{} `xmlrpc:"header_spacing"`
	Id            interface{} `xmlrpc:"id"`
	MarginBottom  interface{} `xmlrpc:"margin_bottom"`
	MarginLeft    interface{} `xmlrpc:"margin_left"`
	MarginRight   interface{} `xmlrpc:"margin_right"`
	MarginTop     interface{} `xmlrpc:"margin_top"`
	Name          interface{} `xmlrpc:"name"`
	Orientation   interface{} `xmlrpc:"orientation"`
	PageHeight    interface{} `xmlrpc:"page_height"`
	PageWidth     interface{} `xmlrpc:"page_width"`
	ReportIds     interface{} `xmlrpc:"report_ids"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var ReportPaperformatModel string = "report.paperformat"

type ReportPaperformats []ReportPaperformat

type ReportPaperformatsNil []ReportPaperformatNil

func (s *ReportPaperformat) NilableType_() interface{} {
	return &ReportPaperformatNil{}
}

func (ns *ReportPaperformatNil) Type_() interface{} {
	s := &ReportPaperformat{}
	return load(ns, s)
}

func (s *ReportPaperformats) NilableType_() interface{} {
	return &ReportPaperformatsNil{}
}

func (ns *ReportPaperformatsNil) Type_() interface{} {
	s := &ReportPaperformats{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportPaperformat))
	}
	return s
}
