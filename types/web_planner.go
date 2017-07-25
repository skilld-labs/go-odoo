package types

import (
	"time"
)

type WebPlanner struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	Active             bool      `xmlrpc:"active"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	Data               string    `xmlrpc:"data"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	MenuId             Many2One  `xmlrpc:"menu_id"`
	Name               string    `xmlrpc:"name"`
	PlannerApplication string    `xmlrpc:"planner_application"`
	Progress           int64     `xmlrpc:"progress"`
	TooltipPlanner     string    `xmlrpc:"tooltip_planner"`
	ViewId             Many2One  `xmlrpc:"view_id"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type WebPlannerNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	Active             bool        `xmlrpc:"active"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	Data               interface{} `xmlrpc:"data"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	MenuId             interface{} `xmlrpc:"menu_id"`
	Name               interface{} `xmlrpc:"name"`
	PlannerApplication interface{} `xmlrpc:"planner_application"`
	Progress           interface{} `xmlrpc:"progress"`
	TooltipPlanner     interface{} `xmlrpc:"tooltip_planner"`
	ViewId             interface{} `xmlrpc:"view_id"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var WebPlannerModel string = "web.planner"

type WebPlanners []WebPlanner

type WebPlannersNil []WebPlannerNil

func (s *WebPlanner) NilableType_() interface{} {
	return &WebPlannerNil{}
}

func (ns *WebPlannerNil) Type_() interface{} {
	s := &WebPlanner{}
	return load(ns, s)
}

func (s *WebPlanners) NilableType_() interface{} {
	return &WebPlannersNil{}
}

func (ns *WebPlannersNil) Type_() interface{} {
	s := &WebPlanners{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WebPlanner))
	}
	return s
}
