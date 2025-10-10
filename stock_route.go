package odoo

// StockRoute represents stock.route model.
type StockRoute struct {
	Active                 *Bool     `xmlrpc:"active,omitempty"`
	CategIds               *Relation `xmlrpc:"categ_ids,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	Name                   *String   `xmlrpc:"name,omitempty"`
	PackagingIds           *Relation `xmlrpc:"packaging_ids,omitempty"`
	PackagingSelectable    *Bool     `xmlrpc:"packaging_selectable,omitempty"`
	ProductCategSelectable *Bool     `xmlrpc:"product_categ_selectable,omitempty"`
	ProductIds             *Relation `xmlrpc:"product_ids,omitempty"`
	ProductSelectable      *Bool     `xmlrpc:"product_selectable,omitempty"`
	RuleIds                *Relation `xmlrpc:"rule_ids,omitempty"`
	SaleSelectable         *Bool     `xmlrpc:"sale_selectable,omitempty"`
	Sequence               *Int      `xmlrpc:"sequence,omitempty"`
	SuppliedWhId           *Many2One `xmlrpc:"supplied_wh_id,omitempty"`
	SupplierWhId           *Many2One `xmlrpc:"supplier_wh_id,omitempty"`
	WarehouseDomainIds     *Relation `xmlrpc:"warehouse_domain_ids,omitempty"`
	WarehouseIds           *Relation `xmlrpc:"warehouse_ids,omitempty"`
	WarehouseSelectable    *Bool     `xmlrpc:"warehouse_selectable,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockRoutes represents array of stock.route model.
type StockRoutes []StockRoute

// StockRouteModel is the odoo model name.
const StockRouteModel = "stock.route"

// Many2One convert StockRoute to *Many2One.
func (sr *StockRoute) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockRoute creates a new stock.route model and returns its id.
func (c *Client) CreateStockRoute(sr *StockRoute) (int64, error) {
	ids, err := c.CreateStockRoutes([]*StockRoute{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockRoute creates a new stock.route model and returns its id.
func (c *Client) CreateStockRoutes(srs []*StockRoute) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(StockRouteModel, vv, nil)
}

// UpdateStockRoute updates an existing stock.route record.
func (c *Client) UpdateStockRoute(sr *StockRoute) error {
	return c.UpdateStockRoutes([]int64{sr.Id.Get()}, sr)
}

// UpdateStockRoutes updates existing stock.route records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockRoutes(ids []int64, sr *StockRoute) error {
	return c.Update(StockRouteModel, ids, sr, nil)
}

// DeleteStockRoute deletes an existing stock.route record.
func (c *Client) DeleteStockRoute(id int64) error {
	return c.DeleteStockRoutes([]int64{id})
}

// DeleteStockRoutes deletes existing stock.route records.
func (c *Client) DeleteStockRoutes(ids []int64) error {
	return c.Delete(StockRouteModel, ids)
}

// GetStockRoute gets stock.route existing record.
func (c *Client) GetStockRoute(id int64) (*StockRoute, error) {
	srs, err := c.GetStockRoutes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetStockRoutes gets stock.route existing records.
func (c *Client) GetStockRoutes(ids []int64) (*StockRoutes, error) {
	srs := &StockRoutes{}
	if err := c.Read(StockRouteModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRoute finds stock.route record by querying it with criteria.
func (c *Client) FindStockRoute(criteria *Criteria) (*StockRoute, error) {
	srs := &StockRoutes{}
	if err := c.SearchRead(StockRouteModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindStockRoutes finds stock.route records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRoutes(criteria *Criteria, options *Options) (*StockRoutes, error) {
	srs := &StockRoutes{}
	if err := c.SearchRead(StockRouteModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRouteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRouteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockRouteModel, criteria, options)
}

// FindStockRouteId finds record id by querying it with criteria.
func (c *Client) FindStockRouteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRouteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
