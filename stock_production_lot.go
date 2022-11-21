package odoo

import (
	"fmt"
)

// StockProductionLot represents stock.production.lot model.
type StockProductionLot struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omptempty"`
	ProductQty               *Float    `xmlrpc:"product_qty,omptempty"`
	ProductUomId             *Many2One `xmlrpc:"product_uom_id,omptempty"`
	QuantIds                 *Relation `xmlrpc:"quant_ids,omptempty"`
	Ref                      *String   `xmlrpc:"ref,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockProductionLots represents array of stock.production.lot model.
type StockProductionLots []StockProductionLot

// StockProductionLotModel is the odoo model name.
const StockProductionLotModel = "stock.production.lot"

// Many2One convert StockProductionLot to *Many2One.
func (spl *StockProductionLot) Many2One() *Many2One {
	return NewMany2One(spl.Id.Get(), "")
}

// CreateStockProductionLot creates a new stock.production.lot model and returns its id.
func (c *Client) CreateStockProductionLot(spl *StockProductionLot) (int64, error) {
	return c.Create(StockProductionLotModel, spl)
}

// UpdateStockProductionLot updates an existing stock.production.lot record.
func (c *Client) UpdateStockProductionLot(spl *StockProductionLot) error {
	return c.UpdateStockProductionLots([]int64{spl.Id.Get()}, spl)
}

// UpdateStockProductionLots updates existing stock.production.lot records.
// All records (represented by ids) will be updated by spl values.
func (c *Client) UpdateStockProductionLots(ids []int64, spl *StockProductionLot) error {
	return c.Update(StockProductionLotModel, ids, spl)
}

// DeleteStockProductionLot deletes an existing stock.production.lot record.
func (c *Client) DeleteStockProductionLot(id int64) error {
	return c.DeleteStockProductionLots([]int64{id})
}

// DeleteStockProductionLots deletes existing stock.production.lot records.
func (c *Client) DeleteStockProductionLots(ids []int64) error {
	return c.Delete(StockProductionLotModel, ids)
}

// GetStockProductionLot gets stock.production.lot existing record.
func (c *Client) GetStockProductionLot(id int64) (*StockProductionLot, error) {
	spls, err := c.GetStockProductionLots([]int64{id})
	if err != nil {
		return nil, err
	}
	if spls != nil && len(*spls) > 0 {
		return &((*spls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.production.lot not found", id)
}

// GetStockProductionLots gets stock.production.lot existing records.
func (c *Client) GetStockProductionLots(ids []int64) (*StockProductionLots, error) {
	spls := &StockProductionLots{}
	if err := c.Read(StockProductionLotModel, ids, nil, spls); err != nil {
		return nil, err
	}
	return spls, nil
}

// FindStockProductionLot finds stock.production.lot record by querying it with criteria.
func (c *Client) FindStockProductionLot(criteria *Criteria) (*StockProductionLot, error) {
	spls := &StockProductionLots{}
	if err := c.SearchRead(StockProductionLotModel, criteria, NewOptions().Limit(1), spls); err != nil {
		return nil, err
	}
	if spls != nil && len(*spls) > 0 {
		return &((*spls)[0]), nil
	}
	return nil, fmt.Errorf("no stock.production.lot was found with criteria %v", criteria)
}

// FindStockProductionLots finds stock.production.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockProductionLots(criteria *Criteria, options *Options) (*StockProductionLots, error) {
	spls := &StockProductionLots{}
	if err := c.SearchRead(StockProductionLotModel, criteria, options, spls); err != nil {
		return nil, err
	}
	return spls, nil
}

// FindStockProductionLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockProductionLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockProductionLotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockProductionLotId finds record id by querying it with criteria.
func (c *Client) FindStockProductionLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockProductionLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.production.lot was found with criteria %v and options %v", criteria, options)
}
