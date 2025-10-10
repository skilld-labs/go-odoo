package odoo

// StockValuationLayer represents stock.valuation.layer model.
type StockValuationLayer struct {
	AccountMoveId          *Many2One `xmlrpc:"account_move_id,omitempty"`
	AccountMoveLineId      *Many2One `xmlrpc:"account_move_line_id,omitempty"`
	CategId                *Many2One `xmlrpc:"categ_id,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId             *Many2One `xmlrpc:"currency_id,omitempty"`
	Description            *String   `xmlrpc:"description,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	LotId                  *Many2One `xmlrpc:"lot_id,omitempty"`
	PriceDiffValue         *Float    `xmlrpc:"price_diff_value,omitempty"`
	ProductId              *Many2One `xmlrpc:"product_id,omitempty"`
	ProductTmplId          *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	Quantity               *Float    `xmlrpc:"quantity,omitempty"`
	Reference              *String   `xmlrpc:"reference,omitempty"`
	RemainingQty           *Float    `xmlrpc:"remaining_qty,omitempty"`
	RemainingValue         *Float    `xmlrpc:"remaining_value,omitempty"`
	StockMoveId            *Many2One `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerId  *Many2One `xmlrpc:"stock_valuation_layer_id,omitempty"`
	StockValuationLayerIds *Relation `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	UnitCost               *Float    `xmlrpc:"unit_cost,omitempty"`
	UomId                  *Many2One `xmlrpc:"uom_id,omitempty"`
	Value                  *Float    `xmlrpc:"value,omitempty"`
	WarehouseId            *Many2One `xmlrpc:"warehouse_id,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockValuationLayers represents array of stock.valuation.layer model.
type StockValuationLayers []StockValuationLayer

// StockValuationLayerModel is the odoo model name.
const StockValuationLayerModel = "stock.valuation.layer"

// Many2One convert StockValuationLayer to *Many2One.
func (svl *StockValuationLayer) Many2One() *Many2One {
	return NewMany2One(svl.Id.Get(), "")
}

// CreateStockValuationLayer creates a new stock.valuation.layer model and returns its id.
func (c *Client) CreateStockValuationLayer(svl *StockValuationLayer) (int64, error) {
	ids, err := c.CreateStockValuationLayers([]*StockValuationLayer{svl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockValuationLayer creates a new stock.valuation.layer model and returns its id.
func (c *Client) CreateStockValuationLayers(svls []*StockValuationLayer) ([]int64, error) {
	var vv []interface{}
	for _, v := range svls {
		vv = append(vv, v)
	}
	return c.Create(StockValuationLayerModel, vv, nil)
}

// UpdateStockValuationLayer updates an existing stock.valuation.layer record.
func (c *Client) UpdateStockValuationLayer(svl *StockValuationLayer) error {
	return c.UpdateStockValuationLayers([]int64{svl.Id.Get()}, svl)
}

// UpdateStockValuationLayers updates existing stock.valuation.layer records.
// All records (represented by ids) will be updated by svl values.
func (c *Client) UpdateStockValuationLayers(ids []int64, svl *StockValuationLayer) error {
	return c.Update(StockValuationLayerModel, ids, svl, nil)
}

// DeleteStockValuationLayer deletes an existing stock.valuation.layer record.
func (c *Client) DeleteStockValuationLayer(id int64) error {
	return c.DeleteStockValuationLayers([]int64{id})
}

// DeleteStockValuationLayers deletes existing stock.valuation.layer records.
func (c *Client) DeleteStockValuationLayers(ids []int64) error {
	return c.Delete(StockValuationLayerModel, ids)
}

// GetStockValuationLayer gets stock.valuation.layer existing record.
func (c *Client) GetStockValuationLayer(id int64) (*StockValuationLayer, error) {
	svls, err := c.GetStockValuationLayers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*svls)[0]), nil
}

// GetStockValuationLayers gets stock.valuation.layer existing records.
func (c *Client) GetStockValuationLayers(ids []int64) (*StockValuationLayers, error) {
	svls := &StockValuationLayers{}
	if err := c.Read(StockValuationLayerModel, ids, nil, svls); err != nil {
		return nil, err
	}
	return svls, nil
}

// FindStockValuationLayer finds stock.valuation.layer record by querying it with criteria.
func (c *Client) FindStockValuationLayer(criteria *Criteria) (*StockValuationLayer, error) {
	svls := &StockValuationLayers{}
	if err := c.SearchRead(StockValuationLayerModel, criteria, NewOptions().Limit(1), svls); err != nil {
		return nil, err
	}
	return &((*svls)[0]), nil
}

// FindStockValuationLayers finds stock.valuation.layer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayers(criteria *Criteria, options *Options) (*StockValuationLayers, error) {
	svls := &StockValuationLayers{}
	if err := c.SearchRead(StockValuationLayerModel, criteria, options, svls); err != nil {
		return nil, err
	}
	return svls, nil
}

// FindStockValuationLayerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockValuationLayerModel, criteria, options)
}

// FindStockValuationLayerId finds record id by querying it with criteria.
func (c *Client) FindStockValuationLayerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockValuationLayerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
