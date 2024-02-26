package odoo

// StockLocationRoute represents stock.location.route model.
type StockLocationRoute struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool     `xmlrpc:"active,omptempty"`
	CategIds               *Relation `xmlrpc:"categ_ids,omptempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	Name                   *String   `xmlrpc:"name,omptempty"`
	ProductCategSelectable *Bool     `xmlrpc:"product_categ_selectable,omptempty"`
	ProductIds             *Relation `xmlrpc:"product_ids,omptempty"`
	ProductSelectable      *Bool     `xmlrpc:"product_selectable,omptempty"`
	PullIds                *Relation `xmlrpc:"pull_ids,omptempty"`
	PushIds                *Relation `xmlrpc:"push_ids,omptempty"`
	SaleSelectable         *Bool     `xmlrpc:"sale_selectable,omptempty"`
	Sequence               *Int      `xmlrpc:"sequence,omptempty"`
	SuppliedWhId           *Many2One `xmlrpc:"supplied_wh_id,omptempty"`
	SupplierWhId           *Many2One `xmlrpc:"supplier_wh_id,omptempty"`
	WarehouseIds           *Relation `xmlrpc:"warehouse_ids,omptempty"`
	WarehouseSelectable    *Bool     `xmlrpc:"warehouse_selectable,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockLocationRoutes represents array of stock.location.route model.
type StockLocationRoutes []StockLocationRoute

// StockLocationRouteModel is the odoo model name.
const StockLocationRouteModel = "stock.location.route"

// Many2One convert StockLocationRoute to *Many2One.
func (slr *StockLocationRoute) Many2One() *Many2One {
	return NewMany2One(slr.Id.Get(), "")
}

// CreateStockLocationRoute creates a new stock.location.route model and returns its id.
func (c *Client) CreateStockLocationRoute(slr *StockLocationRoute) (int64, error) {
	ids, err := c.CreateStockLocationRoutes([]*StockLocationRoute{slr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLocationRoute creates a new stock.location.route model and returns its id.
func (c *Client) CreateStockLocationRoutes(slrs []*StockLocationRoute) ([]int64, error) {
	var vv []interface{}
	for _, v := range slrs {
		vv = append(vv, v)
	}
	return c.Create(StockLocationRouteModel, vv, nil)
}

// UpdateStockLocationRoute updates an existing stock.location.route record.
func (c *Client) UpdateStockLocationRoute(slr *StockLocationRoute) error {
	return c.UpdateStockLocationRoutes([]int64{slr.Id.Get()}, slr)
}

// UpdateStockLocationRoutes updates existing stock.location.route records.
// All records (represented by ids) will be updated by slr values.
func (c *Client) UpdateStockLocationRoutes(ids []int64, slr *StockLocationRoute) error {
	return c.Update(StockLocationRouteModel, ids, slr, nil)
}

// DeleteStockLocationRoute deletes an existing stock.location.route record.
func (c *Client) DeleteStockLocationRoute(id int64) error {
	return c.DeleteStockLocationRoutes([]int64{id})
}

// DeleteStockLocationRoutes deletes existing stock.location.route records.
func (c *Client) DeleteStockLocationRoutes(ids []int64) error {
	return c.Delete(StockLocationRouteModel, ids)
}

// GetStockLocationRoute gets stock.location.route existing record.
func (c *Client) GetStockLocationRoute(id int64) (*StockLocationRoute, error) {
	slrs, err := c.GetStockLocationRoutes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slrs)[0]), nil
}

// GetStockLocationRoutes gets stock.location.route existing records.
func (c *Client) GetStockLocationRoutes(ids []int64) (*StockLocationRoutes, error) {
	slrs := &StockLocationRoutes{}
	if err := c.Read(StockLocationRouteModel, ids, nil, slrs); err != nil {
		return nil, err
	}
	return slrs, nil
}

// FindStockLocationRoute finds stock.location.route record by querying it with criteria.
func (c *Client) FindStockLocationRoute(criteria *Criteria) (*StockLocationRoute, error) {
	slrs := &StockLocationRoutes{}
	if err := c.SearchRead(StockLocationRouteModel, criteria, NewOptions().Limit(1), slrs); err != nil {
		return nil, err
	}
	return &((*slrs)[0]), nil
}

// FindStockLocationRoutes finds stock.location.route records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationRoutes(criteria *Criteria, options *Options) (*StockLocationRoutes, error) {
	slrs := &StockLocationRoutes{}
	if err := c.SearchRead(StockLocationRouteModel, criteria, options, slrs); err != nil {
		return nil, err
	}
	return slrs, nil
}

// FindStockLocationRouteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationRouteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockLocationRouteModel, criteria, options)
}

// FindStockLocationRouteId finds record id by querying it with criteria.
func (c *Client) FindStockLocationRouteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLocationRouteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
