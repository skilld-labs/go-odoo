package odoo

// StockForecastedProductProduct represents stock.forecasted_product_product model.
type StockForecastedProductProduct struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// StockForecastedProductProducts represents array of stock.forecasted_product_product model.
type StockForecastedProductProducts []StockForecastedProductProduct

// StockForecastedProductProductModel is the odoo model name.
const StockForecastedProductProductModel = "stock.forecasted_product_product"

// Many2One convert StockForecastedProductProduct to *Many2One.
func (sf *StockForecastedProductProduct) Many2One() *Many2One {
	return NewMany2One(sf.Id.Get(), "")
}

// CreateStockForecastedProductProduct creates a new stock.forecasted_product_product model and returns its id.
func (c *Client) CreateStockForecastedProductProduct(sf *StockForecastedProductProduct) (int64, error) {
	ids, err := c.CreateStockForecastedProductProducts([]*StockForecastedProductProduct{sf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockForecastedProductProduct creates a new stock.forecasted_product_product model and returns its id.
func (c *Client) CreateStockForecastedProductProducts(sfs []*StockForecastedProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range sfs {
		vv = append(vv, v)
	}
	return c.Create(StockForecastedProductProductModel, vv, nil)
}

// UpdateStockForecastedProductProduct updates an existing stock.forecasted_product_product record.
func (c *Client) UpdateStockForecastedProductProduct(sf *StockForecastedProductProduct) error {
	return c.UpdateStockForecastedProductProducts([]int64{sf.Id.Get()}, sf)
}

// UpdateStockForecastedProductProducts updates existing stock.forecasted_product_product records.
// All records (represented by ids) will be updated by sf values.
func (c *Client) UpdateStockForecastedProductProducts(ids []int64, sf *StockForecastedProductProduct) error {
	return c.Update(StockForecastedProductProductModel, ids, sf, nil)
}

// DeleteStockForecastedProductProduct deletes an existing stock.forecasted_product_product record.
func (c *Client) DeleteStockForecastedProductProduct(id int64) error {
	return c.DeleteStockForecastedProductProducts([]int64{id})
}

// DeleteStockForecastedProductProducts deletes existing stock.forecasted_product_product records.
func (c *Client) DeleteStockForecastedProductProducts(ids []int64) error {
	return c.Delete(StockForecastedProductProductModel, ids)
}

// GetStockForecastedProductProduct gets stock.forecasted_product_product existing record.
func (c *Client) GetStockForecastedProductProduct(id int64) (*StockForecastedProductProduct, error) {
	sfs, err := c.GetStockForecastedProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sfs)[0]), nil
}

// GetStockForecastedProductProducts gets stock.forecasted_product_product existing records.
func (c *Client) GetStockForecastedProductProducts(ids []int64) (*StockForecastedProductProducts, error) {
	sfs := &StockForecastedProductProducts{}
	if err := c.Read(StockForecastedProductProductModel, ids, nil, sfs); err != nil {
		return nil, err
	}
	return sfs, nil
}

// FindStockForecastedProductProduct finds stock.forecasted_product_product record by querying it with criteria.
func (c *Client) FindStockForecastedProductProduct(criteria *Criteria) (*StockForecastedProductProduct, error) {
	sfs := &StockForecastedProductProducts{}
	if err := c.SearchRead(StockForecastedProductProductModel, criteria, NewOptions().Limit(1), sfs); err != nil {
		return nil, err
	}
	return &((*sfs)[0]), nil
}

// FindStockForecastedProductProducts finds stock.forecasted_product_product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockForecastedProductProducts(criteria *Criteria, options *Options) (*StockForecastedProductProducts, error) {
	sfs := &StockForecastedProductProducts{}
	if err := c.SearchRead(StockForecastedProductProductModel, criteria, options, sfs); err != nil {
		return nil, err
	}
	return sfs, nil
}

// FindStockForecastedProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockForecastedProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockForecastedProductProductModel, criteria, options)
}

// FindStockForecastedProductProductId finds record id by querying it with criteria.
func (c *Client) FindStockForecastedProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockForecastedProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
