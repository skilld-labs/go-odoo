package types

import (
	"time"
)

type BoardBoard struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type BoardBoardNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var BoardBoardModel string = "board.board"

type BoardBoards []BoardBoard

type BoardBoardsNil []BoardBoardNil

func (s *BoardBoard) NilableType_() interface{} {
	return &BoardBoardNil{}
}

func (ns *BoardBoardNil) Type_() interface{} {
	s := &BoardBoard{}
	return load(ns, s)
}

func (s *BoardBoards) NilableType_() interface{} {
	return &BoardBoardsNil{}
}

func (ns *BoardBoardsNil) Type_() interface{} {
	s := &BoardBoards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BoardBoard))
	}
	return s
}
