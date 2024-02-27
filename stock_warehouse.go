package odoo

// StockWarehouse represents stock.warehouse model.
type StockWarehouse struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	Active              *Bool      `xmlrpc:"active,omitempty"`
	BuyPullId           *Many2One  `xmlrpc:"buy_pull_id,omitempty"`
	BuyToResupply       *Bool      `xmlrpc:"buy_to_resupply,omitempty"`
	Code                *String    `xmlrpc:"code,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrossdockRouteId    *Many2One  `xmlrpc:"crossdock_route_id,omitempty"`
	DefaultResupplyWhId *Many2One  `xmlrpc:"default_resupply_wh_id,omitempty"`
	DeliveryRouteId     *Many2One  `xmlrpc:"delivery_route_id,omitempty"`
	DeliverySteps       *Selection `xmlrpc:"delivery_steps,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InTypeId            *Many2One  `xmlrpc:"in_type_id,omitempty"`
	IntTypeId           *Many2One  `xmlrpc:"int_type_id,omitempty"`
	LotStockId          *Many2One  `xmlrpc:"lot_stock_id,omitempty"`
	MtoPullId           *Many2One  `xmlrpc:"mto_pull_id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	OutTypeId           *Many2One  `xmlrpc:"out_type_id,omitempty"`
	PackTypeId          *Many2One  `xmlrpc:"pack_type_id,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickTypeId          *Many2One  `xmlrpc:"pick_type_id,omitempty"`
	ReceptionRouteId    *Many2One  `xmlrpc:"reception_route_id,omitempty"`
	ReceptionSteps      *Selection `xmlrpc:"reception_steps,omitempty"`
	ResupplyRouteIds    *Relation  `xmlrpc:"resupply_route_ids,omitempty"`
	ResupplyWhIds       *Relation  `xmlrpc:"resupply_wh_ids,omitempty"`
	RouteIds            *Relation  `xmlrpc:"route_ids,omitempty"`
	ViewLocationId      *Many2One  `xmlrpc:"view_location_id,omitempty"`
	WhInputStockLocId   *Many2One  `xmlrpc:"wh_input_stock_loc_id,omitempty"`
	WhOutputStockLocId  *Many2One  `xmlrpc:"wh_output_stock_loc_id,omitempty"`
	WhPackStockLocId    *Many2One  `xmlrpc:"wh_pack_stock_loc_id,omitempty"`
	WhQcStockLocId      *Many2One  `xmlrpc:"wh_qc_stock_loc_id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(StockWarehouseModel, vv, nil)
}

// UpdateStockWarehouse updates an existing stock.warehouse record.
func (c *Client) UpdateStockWarehouse(sw *StockWarehouse) error {
	return c.UpdateStockWarehouses([]int64{sw.Id.Get()}, sw)
}

// UpdateStockWarehouses updates existing stock.warehouse records.
// All records (represented by ids) will be updated by sw values.
func (c *Client) UpdateStockWarehouses(ids []int64, sw *StockWarehouse) error {
	return c.Update(StockWarehouseModel, ids, sw, nil)
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
	return &((*sws)[0]), nil
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
	return &((*sws)[0]), nil
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
	return c.Search(StockWarehouseModel, criteria, options)
}

// FindStockWarehouseId finds record id by querying it with criteria.
func (c *Client) FindStockWarehouseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarehouseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
