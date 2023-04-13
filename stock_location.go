package odoo

import (
	"fmt"
)

// StockLocation represents stock.location model.
type StockLocation struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	Barcode               *String    `xmlrpc:"barcode,omptempty"`
	ChildIds              *Relation  `xmlrpc:"child_ids,omptempty"`
	Comment               *String    `xmlrpc:"comment,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CompleteName          *String    `xmlrpc:"complete_name,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	LocationId            *Many2One  `xmlrpc:"location_id,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	ParentLeft            *Int       `xmlrpc:"parent_left,omptempty"`
	ParentRight           *Int       `xmlrpc:"parent_right,omptempty"`
	PartnerId             *Many2One  `xmlrpc:"partner_id,omptempty"`
	Posx                  *Int       `xmlrpc:"posx,omptempty"`
	Posy                  *Int       `xmlrpc:"posy,omptempty"`
	Posz                  *Int       `xmlrpc:"posz,omptempty"`
	PutawayStrategyId     *Many2One  `xmlrpc:"putaway_strategy_id,omptempty"`
	QuantIds              *Relation  `xmlrpc:"quant_ids,omptempty"`
	RemovalStrategyId     *Many2One  `xmlrpc:"removal_strategy_id,omptempty"`
	ReturnLocation        *Bool      `xmlrpc:"return_location,omptempty"`
	ScrapLocation         *Bool      `xmlrpc:"scrap_location,omptempty"`
	Usage                 *Selection `xmlrpc:"usage,omptempty"`
	ValuationInAccountId  *Many2One  `xmlrpc:"valuation_in_account_id,omptempty"`
	ValuationOutAccountId *Many2One  `xmlrpc:"valuation_out_account_id,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockLocations represents array of stock.location model.
type StockLocations []StockLocation

// StockLocationModel is the odoo model name.
const StockLocationModel = "stock.location"

// Many2One convert StockLocation to *Many2One.
func (sl *StockLocation) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateStockLocation creates a new stock.location model and returns its id.
func (c *Client) CreateStockLocation(sl *StockLocation) (int64, error) {
	ids, err := c.CreateStockLocations([]*StockLocation{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLocation creates a new stock.location model and returns its id.
func (c *Client) CreateStockLocations(sls []*StockLocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(StockLocationModel, vv)
}

// UpdateStockLocation updates an existing stock.location record.
func (c *Client) UpdateStockLocation(sl *StockLocation) error {
	return c.UpdateStockLocations([]int64{sl.Id.Get()}, sl)
}

// UpdateStockLocations updates existing stock.location records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockLocations(ids []int64, sl *StockLocation) error {
	return c.Update(StockLocationModel, ids, sl)
}

// DeleteStockLocation deletes an existing stock.location record.
func (c *Client) DeleteStockLocation(id int64) error {
	return c.DeleteStockLocations([]int64{id})
}

// DeleteStockLocations deletes existing stock.location records.
func (c *Client) DeleteStockLocations(ids []int64) error {
	return c.Delete(StockLocationModel, ids)
}

// GetStockLocation gets stock.location existing record.
func (c *Client) GetStockLocation(id int64) (*StockLocation, error) {
	sls, err := c.GetStockLocations([]int64{id})
	if err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.location not found", id)
}

// GetStockLocations gets stock.location existing records.
func (c *Client) GetStockLocations(ids []int64) (*StockLocations, error) {
	sls := &StockLocations{}
	if err := c.Read(StockLocationModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLocation finds stock.location record by querying it with criteria.
func (c *Client) FindStockLocation(criteria *Criteria) (*StockLocation, error) {
	sls := &StockLocations{}
	if err := c.SearchRead(StockLocationModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("stock.location was not found with criteria %v", criteria)
}

// FindStockLocations finds stock.location records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocations(criteria *Criteria, options *Options) (*StockLocations, error) {
	sls := &StockLocations{}
	if err := c.SearchRead(StockLocationModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockLocationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockLocationId finds record id by querying it with criteria.
func (c *Client) FindStockLocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.location was not found with criteria %v and options %v", criteria, options)
}
