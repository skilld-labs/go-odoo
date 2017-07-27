package types

import (
	"time"
)

type IrAttachment struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	Checksum     string    `xmlrpc:"checksum"`
	CompanyId    Many2One  `xmlrpc:"company_id"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	Datas        string    `xmlrpc:"datas"`
	DatasFname   string    `xmlrpc:"datas_fname"`
	DbDatas      string    `xmlrpc:"db_datas"`
	Description  string    `xmlrpc:"description"`
	DisplayName  string    `xmlrpc:"display_name"`
	FileSize     int64     `xmlrpc:"file_size"`
	Id           int64     `xmlrpc:"id"`
	IndexContent string    `xmlrpc:"index_content"`
	LocalUrl     string    `xmlrpc:"local_url"`
	Mimetype     string    `xmlrpc:"mimetype"`
	Name         string    `xmlrpc:"name"`
	Public       bool      `xmlrpc:"public"`
	ResField     string    `xmlrpc:"res_field"`
	ResId        int64     `xmlrpc:"res_id"`
	ResModel     string    `xmlrpc:"res_model"`
	ResName      string    `xmlrpc:"res_name"`
	StoreFname   string    `xmlrpc:"store_fname"`
	Type         string    `xmlrpc:"type"`
	Url          string    `xmlrpc:"url"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type IrAttachmentNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	Checksum     interface{} `xmlrpc:"checksum"`
	CompanyId    interface{} `xmlrpc:"company_id"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	Datas        interface{} `xmlrpc:"datas"`
	DatasFname   interface{} `xmlrpc:"datas_fname"`
	DbDatas      interface{} `xmlrpc:"db_datas"`
	Description  interface{} `xmlrpc:"description"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	FileSize     interface{} `xmlrpc:"file_size"`
	Id           interface{} `xmlrpc:"id"`
	IndexContent interface{} `xmlrpc:"index_content"`
	LocalUrl     interface{} `xmlrpc:"local_url"`
	Mimetype     interface{} `xmlrpc:"mimetype"`
	Name         interface{} `xmlrpc:"name"`
	Public       bool        `xmlrpc:"public"`
	ResField     interface{} `xmlrpc:"res_field"`
	ResId        interface{} `xmlrpc:"res_id"`
	ResModel     interface{} `xmlrpc:"res_model"`
	ResName      interface{} `xmlrpc:"res_name"`
	StoreFname   interface{} `xmlrpc:"store_fname"`
	Type         interface{} `xmlrpc:"type"`
	Url          interface{} `xmlrpc:"url"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var IrAttachmentModel string = "ir.attachment"

type IrAttachments []IrAttachment

type IrAttachmentsNil []IrAttachmentNil

func (s *IrAttachment) NilableType_() interface{} {
	return &IrAttachmentNil{}
}

func (ns *IrAttachmentNil) Type_() interface{} {
	s := &IrAttachment{}
	return load(ns, s)
}

func (s *IrAttachments) NilableType_() interface{} {
	return &IrAttachmentsNil{}
}

func (ns *IrAttachmentsNil) Type_() interface{} {
	s := &IrAttachments{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrAttachment))
	}
	return s
}
