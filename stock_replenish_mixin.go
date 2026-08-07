package odoo

// StockReplenishMixin represents stock.replenish.mixin model.
type StockReplenishMixin struct {
	AllowedRouteIds *Relation `xmlrpc:"allowed_route_ids,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	RouteId         *Many2One `xmlrpc:"route_id,omitempty"`
	ShowVendor      *Bool     `xmlrpc:"show_vendor,omitempty"`
	SupplierId      *Many2One `xmlrpc:"supplier_id,omitempty"`
}

// StockReplenishMixins represents array of stock.replenish.mixin model.
type StockReplenishMixins []StockReplenishMixin

// StockReplenishMixinModel is the odoo model name.
const StockReplenishMixinModel = "stock.replenish.mixin"

// Many2One convert StockReplenishMixin to *Many2One.
func (srm *StockReplenishMixin) Many2One() *Many2One {
	return NewMany2One(srm.Id.Get(), "")
}

// CreateStockReplenishMixin creates a new stock.replenish.mixin model and returns its id.
func (c *Client) CreateStockReplenishMixin(srm *StockReplenishMixin) (int64, error) {
	ids, err := c.CreateStockReplenishMixins([]*StockReplenishMixin{srm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReplenishMixin creates a new stock.replenish.mixin model and returns its id.
func (c *Client) CreateStockReplenishMixins(srms []*StockReplenishMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range srms {
		vv = append(vv, v)
	}
	return c.Create(StockReplenishMixinModel, vv, nil)
}

// UpdateStockReplenishMixin updates an existing stock.replenish.mixin record.
func (c *Client) UpdateStockReplenishMixin(srm *StockReplenishMixin) error {
	return c.UpdateStockReplenishMixins([]int64{srm.Id.Get()}, srm)
}

// UpdateStockReplenishMixins updates existing stock.replenish.mixin records.
// All records (represented by ids) will be updated by srm values.
func (c *Client) UpdateStockReplenishMixins(ids []int64, srm *StockReplenishMixin) error {
	return c.Update(StockReplenishMixinModel, ids, srm, nil)
}

// DeleteStockReplenishMixin deletes an existing stock.replenish.mixin record.
func (c *Client) DeleteStockReplenishMixin(id int64) error {
	return c.DeleteStockReplenishMixins([]int64{id})
}

// DeleteStockReplenishMixins deletes existing stock.replenish.mixin records.
func (c *Client) DeleteStockReplenishMixins(ids []int64) error {
	return c.Delete(StockReplenishMixinModel, ids)
}

// GetStockReplenishMixin gets stock.replenish.mixin existing record.
func (c *Client) GetStockReplenishMixin(id int64) (*StockReplenishMixin, error) {
	srms, err := c.GetStockReplenishMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srms)[0]), nil
}

// GetStockReplenishMixins gets stock.replenish.mixin existing records.
func (c *Client) GetStockReplenishMixins(ids []int64) (*StockReplenishMixins, error) {
	srms := &StockReplenishMixins{}
	if err := c.Read(StockReplenishMixinModel, ids, nil, srms); err != nil {
		return nil, err
	}
	return srms, nil
}

// FindStockReplenishMixin finds stock.replenish.mixin record by querying it with criteria.
func (c *Client) FindStockReplenishMixin(criteria *Criteria) (*StockReplenishMixin, error) {
	srms := &StockReplenishMixins{}
	if err := c.SearchRead(StockReplenishMixinModel, criteria, NewOptions().Limit(1), srms); err != nil {
		return nil, err
	}
	return &((*srms)[0]), nil
}

// FindStockReplenishMixins finds stock.replenish.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishMixins(criteria *Criteria, options *Options) (*StockReplenishMixins, error) {
	srms := &StockReplenishMixins{}
	if err := c.SearchRead(StockReplenishMixinModel, criteria, options, srms); err != nil {
		return nil, err
	}
	return srms, nil
}

// FindStockReplenishMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReplenishMixinModel, criteria, options)
}

// FindStockReplenishMixinId finds record id by querying it with criteria.
func (c *Client) FindStockReplenishMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReplenishMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
