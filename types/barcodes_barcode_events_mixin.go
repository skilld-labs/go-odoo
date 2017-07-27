package types

import (
	"time"
)

type BarcodesBarcodeEventsMixin struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	BarcodeScanned string    `xmlrpc:"_barcode_scanned"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
}

type BarcodesBarcodeEventsMixinNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	BarcodeScanned interface{} `xmlrpc:"_barcode_scanned"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
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
