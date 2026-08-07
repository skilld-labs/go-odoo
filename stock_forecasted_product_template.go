package odoo

// StockForecastedProductTemplate represents stock.forecasted_product_template model.
type StockForecastedProductTemplate struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// StockForecastedProductTemplates represents array of stock.forecasted_product_template model.
type StockForecastedProductTemplates []StockForecastedProductTemplate

// StockForecastedProductTemplateModel is the odoo model name.
const StockForecastedProductTemplateModel = "stock.forecasted_product_template"

// Many2One convert StockForecastedProductTemplate to *Many2One.
func (sf *StockForecastedProductTemplate) Many2One() *Many2One {
	return NewMany2One(sf.Id.Get(), "")
}

// CreateStockForecastedProductTemplate creates a new stock.forecasted_product_template model and returns its id.
func (c *Client) CreateStockForecastedProductTemplate(sf *StockForecastedProductTemplate) (int64, error) {
	ids, err := c.CreateStockForecastedProductTemplates([]*StockForecastedProductTemplate{sf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockForecastedProductTemplate creates a new stock.forecasted_product_template model and returns its id.
func (c *Client) CreateStockForecastedProductTemplates(sfs []*StockForecastedProductTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sfs {
		vv = append(vv, v)
	}
	return c.Create(StockForecastedProductTemplateModel, vv, nil)
}

// UpdateStockForecastedProductTemplate updates an existing stock.forecasted_product_template record.
func (c *Client) UpdateStockForecastedProductTemplate(sf *StockForecastedProductTemplate) error {
	return c.UpdateStockForecastedProductTemplates([]int64{sf.Id.Get()}, sf)
}

// UpdateStockForecastedProductTemplates updates existing stock.forecasted_product_template records.
// All records (represented by ids) will be updated by sf values.
func (c *Client) UpdateStockForecastedProductTemplates(ids []int64, sf *StockForecastedProductTemplate) error {
	return c.Update(StockForecastedProductTemplateModel, ids, sf, nil)
}

// DeleteStockForecastedProductTemplate deletes an existing stock.forecasted_product_template record.
func (c *Client) DeleteStockForecastedProductTemplate(id int64) error {
	return c.DeleteStockForecastedProductTemplates([]int64{id})
}

// DeleteStockForecastedProductTemplates deletes existing stock.forecasted_product_template records.
func (c *Client) DeleteStockForecastedProductTemplates(ids []int64) error {
	return c.Delete(StockForecastedProductTemplateModel, ids)
}

// GetStockForecastedProductTemplate gets stock.forecasted_product_template existing record.
func (c *Client) GetStockForecastedProductTemplate(id int64) (*StockForecastedProductTemplate, error) {
	sfs, err := c.GetStockForecastedProductTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sfs)[0]), nil
}

// GetStockForecastedProductTemplates gets stock.forecasted_product_template existing records.
func (c *Client) GetStockForecastedProductTemplates(ids []int64) (*StockForecastedProductTemplates, error) {
	sfs := &StockForecastedProductTemplates{}
	if err := c.Read(StockForecastedProductTemplateModel, ids, nil, sfs); err != nil {
		return nil, err
	}
	return sfs, nil
}

// FindStockForecastedProductTemplate finds stock.forecasted_product_template record by querying it with criteria.
func (c *Client) FindStockForecastedProductTemplate(criteria *Criteria) (*StockForecastedProductTemplate, error) {
	sfs := &StockForecastedProductTemplates{}
	if err := c.SearchRead(StockForecastedProductTemplateModel, criteria, NewOptions().Limit(1), sfs); err != nil {
		return nil, err
	}
	return &((*sfs)[0]), nil
}

// FindStockForecastedProductTemplates finds stock.forecasted_product_template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockForecastedProductTemplates(criteria *Criteria, options *Options) (*StockForecastedProductTemplates, error) {
	sfs := &StockForecastedProductTemplates{}
	if err := c.SearchRead(StockForecastedProductTemplateModel, criteria, options, sfs); err != nil {
		return nil, err
	}
	return sfs, nil
}

// FindStockForecastedProductTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockForecastedProductTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockForecastedProductTemplateModel, criteria, options)
}

// FindStockForecastedProductTemplateId finds record id by querying it with criteria.
func (c *Client) FindStockForecastedProductTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockForecastedProductTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
