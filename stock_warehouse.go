package odoo

import (
	"fmt"
)

// StockWarehouse represents stock.warehouse model.
type StockWarehouse struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Active              *Bool      `xmlrpc:"active,omptempty"`
	BuyPullId           *Many2One  `xmlrpc:"buy_pull_id,omptempty"`
	BuyToResupply       *Bool      `xmlrpc:"buy_to_resupply,omptempty"`
	Code                *String    `xmlrpc:"code,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrossdockRouteId    *Many2One  `xmlrpc:"crossdock_route_id,omptempty"`
	DefaultResupplyWhId *Many2One  `xmlrpc:"default_resupply_wh_id,omptempty"`
	DeliveryRouteId     *Many2One  `xmlrpc:"delivery_route_id,omptempty"`
	DeliverySteps       *Selection `xmlrpc:"delivery_steps,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InTypeId            *Many2One  `xmlrpc:"in_type_id,omptempty"`
	IntTypeId           *Many2One  `xmlrpc:"int_type_id,omptempty"`
	LotStockId          *Many2One  `xmlrpc:"lot_stock_id,omptempty"`
	MtoPullId           *Many2One  `xmlrpc:"mto_pull_id,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	OutTypeId           *Many2One  `xmlrpc:"out_type_id,omptempty"`
	PackTypeId          *Many2One  `xmlrpc:"pack_type_id,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PickTypeId          *Many2One  `xmlrpc:"pick_type_id,omptempty"`
	ReceptionRouteId    *Many2One  `xmlrpc:"reception_route_id,omptempty"`
	ReceptionSteps      *Selection `xmlrpc:"reception_steps,omptempty"`
	ResupplyRouteIds    *Relation  `xmlrpc:"resupply_route_ids,omptempty"`
	ResupplyWhIds       *Relation  `xmlrpc:"resupply_wh_ids,omptempty"`
	RouteIds            *Relation  `xmlrpc:"route_ids,omptempty"`
	ViewLocationId      *Many2One  `xmlrpc:"view_location_id,omptempty"`
	WhInputStockLocId   *Many2One  `xmlrpc:"wh_input_stock_loc_id,omptempty"`
	WhOutputStockLocId  *Many2One  `xmlrpc:"wh_output_stock_loc_id,omptempty"`
	WhPackStockLocId    *Many2One  `xmlrpc:"wh_pack_stock_loc_id,omptempty"`
	WhQcStockLocId      *Many2One  `xmlrpc:"wh_qc_stock_loc_id,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockWarehouses represents array of stock.warehouse model.
type StockWarehouses []StockWarehouse

// StockWarehouseModel is the odoo model name.
const StockWarehouseModel = "stock.warehouse"

// Many2One convert StockWarehouse to *Many2One.
func (sw *StockWarehouse) Many2One() *Many2One {
	return NewMany2One(sw.Id.Get(), "")
}

// CreateStockWarehouse creates a new stock.warehouse model and returns its id.
func (c *Client) CreateStockWarehouse(sw *StockWarehouse) (int64, error) {
	ids, err := c.CreateStockWarehouses([]*StockWarehouse{sw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockWarehouse creates a new stock.warehouse model and returns its id.
func (c *Client) CreateStockWarehouses(sws []*StockWarehouse) ([]int64, error) {
	var vv []interface{}
	for _, v := range sws {
		vv = append(vv, v)
	}
	return c.Create(StockWarehouseModel, vv)
}

// UpdateStockWarehouse updates an existing stock.warehouse record.
func (c *Client) UpdateStockWarehouse(sw *StockWarehouse) error {
	return c.UpdateStockWarehouses([]int64{sw.Id.Get()}, sw)
}

// UpdateStockWarehouses updates existing stock.warehouse records.
// All records (represented by ids) will be updated by sw values.
func (c *Client) UpdateStockWarehouses(ids []int64, sw *StockWarehouse) error {
	return c.Update(StockWarehouseModel, ids, sw)
}

// DeleteStockWarehouse deletes an existing stock.warehouse record.
func (c *Client) DeleteStockWarehouse(id int64) error {
	return c.DeleteStockWarehouses([]int64{id})
}

// DeleteStockWarehouses deletes existing stock.warehouse records.
func (c *Client) DeleteStockWarehouses(ids []int64) error {
	return c.Delete(StockWarehouseModel, ids)
}

// GetStockWarehouse gets stock.warehouse existing record.
func (c *Client) GetStockWarehouse(id int64) (*StockWarehouse, error) {
	sws, err := c.GetStockWarehouses([]int64{id})
	if err != nil {
		return nil, err
	}
	if sws != nil && len(*sws) > 0 {
		return &((*sws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.warehouse not found", id)
}

// GetStockWarehouses gets stock.warehouse existing records.
func (c *Client) GetStockWarehouses(ids []int64) (*StockWarehouses, error) {
	sws := &StockWarehouses{}
	if err := c.Read(StockWarehouseModel, ids, nil, sws); err != nil {
		return nil, err
	}
	return sws, nil
}

// FindStockWarehouse finds stock.warehouse record by querying it with criteria.
func (c *Client) FindStockWarehouse(criteria *Criteria) (*StockWarehouse, error) {
	sws := &StockWarehouses{}
	if err := c.SearchRead(StockWarehouseModel, criteria, NewOptions().Limit(1), sws); err != nil {
		return nil, err
	}
	if sws != nil && len(*sws) > 0 {
		return &((*sws)[0]), nil
	}
	return nil, fmt.Errorf("stock.warehouse was not found with criteria %v", criteria)
}

// FindStockWarehouses finds stock.warehouse records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouses(criteria *Criteria, options *Options) (*StockWarehouses, error) {
	sws := &StockWarehouses{}
	if err := c.SearchRead(StockWarehouseModel, criteria, options, sws); err != nil {
		return nil, err
	}
	return sws, nil
}

// FindStockWarehouseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarehouseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockWarehouseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockWarehouseId finds record id by querying it with criteria.
func (c *Client) FindStockWarehouseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarehouseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.warehouse was not found with criteria %v and options %v", criteria, options)
}
