package types

import (
	"time"
)

type StockProductionLot struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	Name                     string    `xmlrpc:"name"`
	ProductId                Many2One  `xmlrpc:"product_id"`
	ProductQty               float64   `xmlrpc:"product_qty"`
	ProductUomId             Many2One  `xmlrpc:"product_uom_id"`
	QuantIds                 []int64   `xmlrpc:"quant_ids"`
	Ref                      string    `xmlrpc:"ref"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type StockProductionLotNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	Name                     interface{} `xmlrpc:"name"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProductQty               interface{} `xmlrpc:"product_qty"`
	ProductUomId             interface{} `xmlrpc:"product_uom_id"`
	QuantIds                 interface{} `xmlrpc:"quant_ids"`
	Ref                      interface{} `xmlrpc:"ref"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var StockProductionLotModel string = "stock.production.lot"

type StockProductionLots []StockProductionLot

type StockProductionLotsNil []StockProductionLotNil

func (s *StockProductionLot) NilableType_() interface{} {
	return &StockProductionLotNil{}
}

func (ns *StockProductionLotNil) Type_() interface{} {
	s := &StockProductionLot{}
	return load(ns, s)
}

func (s *StockProductionLots) NilableType_() interface{} {
	return &StockProductionLotsNil{}
}

func (ns *StockProductionLotsNil) Type_() interface{} {
	s := &StockProductionLots{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockProductionLot))
	}
	return s
}
