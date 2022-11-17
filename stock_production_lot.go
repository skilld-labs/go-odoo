package odoo

import (
	"fmt"
)

// StockProductionLot represents stock.production.lot model.
type StockProductionLot struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omitempty"`
	ProductQty               *Float    `xmlrpc:"product_qty,omitempty"`
	ProductUomId             *Many2One `xmlrpc:"product_uom_id,omitempty"`
	QuantIds                 *Relation `xmlrpc:"quant_ids,omitempty"`
	Ref                      *String   `xmlrpc:"ref,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
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
