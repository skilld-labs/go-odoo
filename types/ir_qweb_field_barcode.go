package types

import (
	"time"
)

type IrQwebFieldBarcode struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldBarcodeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldBarcodeModel string = "ir.qweb.field.barcode"

type IrQwebFieldBarcodes []IrQwebFieldBarcode

type IrQwebFieldBarcodesNil []IrQwebFieldBarcodeNil

func (s *IrQwebFieldBarcode) NilableType_() interface{} {
	return &IrQwebFieldBarcodeNil{}
}

func (ns *IrQwebFieldBarcodeNil) Type_() interface{} {
	s := &IrQwebFieldBarcode{}
	return load(ns, s)
}

func (s *IrQwebFieldBarcodes) NilableType_() interface{} {
	return &IrQwebFieldBarcodesNil{}
}

func (ns *IrQwebFieldBarcodesNil) Type_() interface{} {
	s := &IrQwebFieldBarcodes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldBarcode))
	}
	return s
}
