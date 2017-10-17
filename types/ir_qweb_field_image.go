package types

import (
	"time"
)

type IrQwebFieldImage struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldImageNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
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
