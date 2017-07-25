package types

import (
	"time"
)

type ProcurementRule struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Action                 string    `xmlrpc:"action"`
	Active                 bool      `xmlrpc:"active"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	Delay                  int64     `xmlrpc:"delay"`
	DisplayName            string    `xmlrpc:"display_name"`
	GroupId                Many2One  `xmlrpc:"group_id"`
	GroupPropagationOption string    `xmlrpc:"group_propagation_option"`
	Id                     int64     `xmlrpc:"id"`
	LocationId             Many2One  `xmlrpc:"location_id"`
	LocationSrcId          Many2One  `xmlrpc:"location_src_id"`
	Name                   string    `xmlrpc:"name"`
	PartnerAddressId       Many2One  `xmlrpc:"partner_address_id"`
	PickingTypeId          Many2One  `xmlrpc:"picking_type_id"`
	ProcureMethod          string    `xmlrpc:"procure_method"`
	Propagate              bool      `xmlrpc:"propagate"`
	PropagateWarehouseId   Many2One  `xmlrpc:"propagate_warehouse_id"`
	RouteId                Many2One  `xmlrpc:"route_id"`
	RouteSequence          int64     `xmlrpc:"route_sequence"`
	Sequence               int64     `xmlrpc:"sequence"`
	WarehouseId            Many2One  `xmlrpc:"warehouse_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type ProcurementRuleNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Action                 interface{} `xmlrpc:"action"`
	Active                 bool        `xmlrpc:"active"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	Delay                  interface{} `xmlrpc:"delay"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	GroupId                interface{} `xmlrpc:"group_id"`
	GroupPropagationOption interface{} `xmlrpc:"group_propagation_option"`
	Id                     interface{} `xmlrpc:"id"`
	LocationId             interface{} `xmlrpc:"location_id"`
	LocationSrcId          interface{} `xmlrpc:"location_src_id"`
	Name                   interface{} `xmlrpc:"name"`
	PartnerAddressId       interface{} `xmlrpc:"partner_address_id"`
	PickingTypeId          interface{} `xmlrpc:"picking_type_id"`
	ProcureMethod          interface{} `xmlrpc:"procure_method"`
	Propagate              bool        `xmlrpc:"propagate"`
	PropagateWarehouseId   interface{} `xmlrpc:"propagate_warehouse_id"`
	RouteId                interface{} `xmlrpc:"route_id"`
	RouteSequence          interface{} `xmlrpc:"route_sequence"`
	Sequence               interface{} `xmlrpc:"sequence"`
	WarehouseId            interface{} `xmlrpc:"warehouse_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var ProcurementRuleModel string = "procurement.rule"

type ProcurementRules []ProcurementRule

type ProcurementRulesNil []ProcurementRuleNil

func (s *ProcurementRule) NilableType_() interface{} {
	return &ProcurementRuleNil{}
}

func (ns *ProcurementRuleNil) Type_() interface{} {
	s := &ProcurementRule{}
	return load(ns, s)
}

func (s *ProcurementRules) NilableType_() interface{} {
	return &ProcurementRulesNil{}
}

func (ns *ProcurementRulesNil) Type_() interface{} {
	s := &ProcurementRules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProcurementRule))
	}
	return s
}
