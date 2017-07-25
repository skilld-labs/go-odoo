package types

import (
	"time"
)

type IrQwebFieldImage struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldImageNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldImageModel string = "ir.qweb.field.image"

type IrQwebFieldImages []IrQwebFieldImage

type IrQwebFieldImagesNil []IrQwebFieldImageNil

func (s *IrQwebFieldImage) NilableType_() interface{} {
	return &IrQwebFieldImageNil{}
}

func (ns *IrQwebFieldImageNil) Type_() interface{} {
	s := &IrQwebFieldImage{}
	return load(ns, s)
}

func (s *IrQwebFieldImages) NilableType_() interface{} {
	return &IrQwebFieldImagesNil{}
}

func (ns *IrQwebFieldImagesNil) Type_() interface{} {
	s := &IrQwebFieldImages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldImage))
	}
	return s
}
