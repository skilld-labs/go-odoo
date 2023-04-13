package odoo

import (
	"fmt"
)

// StockWarehouseOrderpoint represents stock.warehouse.orderpoint model.
type StockWarehouseOrderpoint struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	Active        *Bool      `xmlrpc:"active,omptempty"`
	CompanyId     *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	GroupId       *Many2One  `xmlrpc:"group_id,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	LeadDays      *Int       `xmlrpc:"lead_days,omptempty"`
	LeadType      *Selection `xmlrpc:"lead_type,omptempty"`
	LocationId    *Many2One  `xmlrpc:"location_id,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	ProductId     *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductMaxQty *Float     `xmlrpc:"product_max_qty,omptempty"`
	ProductMinQty *Float     `xmlrpc:"product_min_qty,omptempty"`
	ProductUom    *Many2One  `xmlrpc:"product_uom,omptempty"`
	QtyMultiple   *Float     `xmlrpc:"qty_multiple,omptempty"`
	WarehouseId   *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockWarehouseOrderpoints represents array of stock.warehouse.orderpoint model.
type StockWarehouseOrderpoints []StockWarehouseOrderpoint

// StockWarehouseOrderpointModel is the odoo model name.
const StockWarehouseOrderpointModel = "stock.warehouse.orderpoint"

// Many2One convert StockWarehouseOrderpoint to *Many2One.
func (swo *StockWarehouseOrderpoint) Many2One() *Many2One {
	return NewMany2One(swo.Id.Get(), "")
}

// CreateStockWarehouseOrderpoint creates a new stock.warehouse.orderpoint model and returns its id.
func (c *Client) CreateStockWarehouseOrderpoint(swo *StockWarehouseOrderpoint) (int64, error) {
	ids, err := c.Create(StockWarehouseOrderpointModel, []interface{}{swo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockWarehouseOrderpoint creates a new stock.warehouse.orderpoint model and returns its id.
func (c *Client) CreateStockWarehouseOrderpoints(swos []*StockWarehouseOrderpoint) ([]int64, error) {
	var vv []interface{}
	for _, v := range swos {
		vv = append(vv, v)
	}
	return c.Create(StockWarehouseOrderpointModel, vv)
}

// UpdateStockWarehouseOrderpoint updates an existing stock.warehouse.orderpoint record.
func (c *Client) UpdateStockWarehouseOrderpoint(swo *StockWarehouseOrderpoint) error {
	return c.UpdateStockWarehouseOrderpoints([]int64{swo.Id.Get()}, swo)
}

// UpdateStockWarehouseOrderpoints updates existing stock.warehouse.orderpoint records.
// All records (represented by ids) will be updated by swo values.
func (c *Client) UpdateStockWarehouseOrderpoints(ids []int64, swo *StockWarehouseOrderpoint) error {
	return c.Update(StockWarehouseOrderpointModel, ids, swo)
}

// DeleteStockWarehouseOrderpoint deletes an existing stock.warehouse.orderpoint record.
func (c *Client) DeleteStockWarehouseOrderpoint(id int64) error {
	return c.DeleteStockWarehouseOrderpoints([]int64{id})
}

// DeleteStockWarehouseOrderpoints deletes existing stock.warehouse.orderpoint records.
func (c *Client) DeleteStockWarehouseOrderpoints(ids []int64) error {
	return c.Delete(StockWarehouseOrderpointModel, ids)
}

// GetStockWarehouseOrderpoint gets stock.warehouse.orderpoint existing record.
func (c *Client) GetStockWarehouseOrderpoint(id int64) (*StockWarehouseOrderpoint, error) {
	swos, err := c.GetStockWarehouseOrderpoints([]int64{id})
	if err != nil {
		return nil, err
	}
	if swos != nil && len(*swos) > 0 {
		return &((*swos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.warehouse.orderpoint not found", id)
}

// GetStockWarehouseOrderpoints gets stock.warehouse.orderpoint existing records.
func (c *Client) GetStockWarehouseOrderpoints(ids []int64) (*StockWarehouseOrderpoints, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.Read(StockWarehouseOrderpointModel, ids, nil, swos); err != nil {
		return nil, err
	}
	return swos, nil
}

// FindStockWarehouseOrderpoint finds stock.warehouse.orderpoint record by querying it with criteria.
func (c *Client) FindStockWarehouseOrderpoint(criteria *Criteria) (*StockWarehouseOrderpoint, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.SearchRead(StockWarehouseOrderpointModel, criteria, NewOptions().Limit(1), swos); err != nil {
		return nil, err
	}
	if swos != nil && len(*swos) > 0 {
		return &((*swos)[0]), nil
	}
	return nil, fmt.Errorf("stock.warehouse.orderpoint was not found with criteria %v", criteria)
}

// FindStockWarehouseOrderpoints finds stock.warehouse.orderpoint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouseOrderpoints(criteria *Criteria, options *Options) (*StockWarehouseOrderpoints, error) {
	swos := &StockWarehouseOrderpoints{}
	if err := c.SearchRead(StockWarehouseOrderpointModel, criteria, options, swos); err != nil {
		return nil, err
	}
	return swos, nil
}

// FindStockWarehouseOrderpointIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouseOrderpointIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockWarehouseOrderpointModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockWarehouseOrderpointId finds record id by querying it with criteria.
func (c *Client) FindStockWarehouseOrderpointId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarehouseOrderpointModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.warehouse.orderpoint was not found with criteria %v and options %v", criteria, options)
}
