package types

import (
	"time"
)

type BarcodesBarcodeEventsMixin struct {
	BarcodeScanned string    `xmlrpc:"_barcode_scanned"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LastUpdate     time.Time `xmlrpc:"__last_update"`
}

type BarcodesBarcodeEventsMixinNil struct {
	BarcodeScanned interface{} `xmlrpc:"_barcode_scanned"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
}

var BarcodesBarcodeEventsMixinModel string = "barcodes.barcode_events_mixin"

type BarcodesBarcodeEventsMixins []BarcodesBarcodeEventsMixin

type BarcodesBarcodeEventsMixinsNil []BarcodesBarcodeEventsMixinNil

func (s *BarcodesBarcodeEventsMixin) NilableType_() interface{} {
	return &BarcodesBarcodeEventsMixinNil{}
}

func (ns *BarcodesBarcodeEventsMixinNil) Type_() interface{} {
	s := &BarcodesBarcodeEventsMixin{}
	return load(ns, s)
}

func (s *BarcodesBarcodeEventsMixins) NilableType_() interface{} {
	return &BarcodesBarcodeEventsMixinsNil{}
}

func (ns *BarcodesBarcodeEventsMixinsNil) Type_() interface{} {
	s := &BarcodesBarcodeEventsMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BarcodesBarcodeEventsMixin))
	}
	return s
}
