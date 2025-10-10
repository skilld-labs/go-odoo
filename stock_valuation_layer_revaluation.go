package odoo

// StockValuationLayerRevaluation represents stock.valuation.layer.revaluation model.
type StockValuationLayerRevaluation struct {
	AccountId          *Many2One  `xmlrpc:"account_id,omitempty"`
	AccountJournalId   *Many2One  `xmlrpc:"account_journal_id,omitempty"`
	AddedValue         *Float     `xmlrpc:"added_value,omitempty"`
	AdjustedLayerIds   *Relation  `xmlrpc:"adjusted_layer_ids,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrentQuantitySvl *Float     `xmlrpc:"current_quantity_svl,omitempty"`
	CurrentValueSvl    *Float     `xmlrpc:"current_value_svl,omitempty"`
	Date               *Time      `xmlrpc:"date,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	LotId              *Many2One  `xmlrpc:"lot_id,omitempty"`
	NewValue           *Float     `xmlrpc:"new_value,omitempty"`
	NewValueByQty      *Float     `xmlrpc:"new_value_by_qty,omitempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomName     *String    `xmlrpc:"product_uom_name,omitempty"`
	PropertyValuation  *Selection `xmlrpc:"property_valuation,omitempty"`
	Reason             *String    `xmlrpc:"reason,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockValuationLayerRevaluations represents array of stock.valuation.layer.revaluation model.
type StockValuationLayerRevaluations []StockValuationLayerRevaluation

// StockValuationLayerRevaluationModel is the odoo model name.
const StockValuationLayerRevaluationModel = "stock.valuation.layer.revaluation"

// Many2One convert StockValuationLayerRevaluation to *Many2One.
func (svlr *StockValuationLayerRevaluation) Many2One() *Many2One {
	return NewMany2One(svlr.Id.Get(), "")
}

// CreateStockValuationLayerRevaluation creates a new stock.valuation.layer.revaluation model and returns its id.
func (c *Client) CreateStockValuationLayerRevaluation(svlr *StockValuationLayerRevaluation) (int64, error) {
	ids, err := c.CreateStockValuationLayerRevaluations([]*StockValuationLayerRevaluation{svlr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockValuationLayerRevaluation creates a new stock.valuation.layer.revaluation model and returns its id.
func (c *Client) CreateStockValuationLayerRevaluations(svlrs []*StockValuationLayerRevaluation) ([]int64, error) {
	var vv []interface{}
	for _, v := range svlrs {
		vv = append(vv, v)
	}
	return c.Create(StockValuationLayerRevaluationModel, vv, nil)
}

// UpdateStockValuationLayerRevaluation updates an existing stock.valuation.layer.revaluation record.
func (c *Client) UpdateStockValuationLayerRevaluation(svlr *StockValuationLayerRevaluation) error {
	return c.UpdateStockValuationLayerRevaluations([]int64{svlr.Id.Get()}, svlr)
}

// UpdateStockValuationLayerRevaluations updates existing stock.valuation.layer.revaluation records.
// All records (represented by ids) will be updated by svlr values.
func (c *Client) UpdateStockValuationLayerRevaluations(ids []int64, svlr *StockValuationLayerRevaluation) error {
	return c.Update(StockValuationLayerRevaluationModel, ids, svlr, nil)
}

// DeleteStockValuationLayerRevaluation deletes an existing stock.valuation.layer.revaluation record.
func (c *Client) DeleteStockValuationLayerRevaluation(id int64) error {
	return c.DeleteStockValuationLayerRevaluations([]int64{id})
}

// DeleteStockValuationLayerRevaluations deletes existing stock.valuation.layer.revaluation records.
func (c *Client) DeleteStockValuationLayerRevaluations(ids []int64) error {
	return c.Delete(StockValuationLayerRevaluationModel, ids)
}

// GetStockValuationLayerRevaluation gets stock.valuation.layer.revaluation existing record.
func (c *Client) GetStockValuationLayerRevaluation(id int64) (*StockValuationLayerRevaluation, error) {
	svlrs, err := c.GetStockValuationLayerRevaluations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*svlrs)[0]), nil
}

// GetStockValuationLayerRevaluations gets stock.valuation.layer.revaluation existing records.
func (c *Client) GetStockValuationLayerRevaluations(ids []int64) (*StockValuationLayerRevaluations, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.Read(StockValuationLayerRevaluationModel, ids, nil, svlrs); err != nil {
		return nil, err
	}
	return svlrs, nil
}

// FindStockValuationLayerRevaluation finds stock.valuation.layer.revaluation record by querying it with criteria.
func (c *Client) FindStockValuationLayerRevaluation(criteria *Criteria) (*StockValuationLayerRevaluation, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.SearchRead(StockValuationLayerRevaluationModel, criteria, NewOptions().Limit(1), svlrs); err != nil {
		return nil, err
	}
	return &((*svlrs)[0]), nil
}

// FindStockValuationLayerRevaluations finds stock.valuation.layer.revaluation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerRevaluations(criteria *Criteria, options *Options) (*StockValuationLayerRevaluations, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.SearchRead(StockValuationLayerRevaluationModel, criteria, options, svlrs); err != nil {
		return nil, err
	}
	return svlrs, nil
}

// FindStockValuationLayerRevaluationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerRevaluationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockValuationLayerRevaluationModel, criteria, options)
}

// FindStockValuationLayerRevaluationId finds record id by querying it with criteria.
func (c *Client) FindStockValuationLayerRevaluationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockValuationLayerRevaluationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
