package odoo

// StockRulesReport represents stock.rules.report model.
type StockRulesReport struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	ProductHasVariants *Bool     `xmlrpc:"product_has_variants,omitempty"`
	ProductId          *Many2One `xmlrpc:"product_id,omitempty"`
	ProductTmplId      *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	SoRouteIds         *Relation `xmlrpc:"so_route_ids,omitempty"`
	WarehouseIds       *Relation `xmlrpc:"warehouse_ids,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockRulesReports represents array of stock.rules.report model.
type StockRulesReports []StockRulesReport

// StockRulesReportModel is the odoo model name.
const StockRulesReportModel = "stock.rules.report"

// Many2One convert StockRulesReport to *Many2One.
func (srr *StockRulesReport) Many2One() *Many2One {
	return NewMany2One(srr.Id.Get(), "")
}

// CreateStockRulesReport creates a new stock.rules.report model and returns its id.
func (c *Client) CreateStockRulesReport(srr *StockRulesReport) (int64, error) {
	ids, err := c.CreateStockRulesReports([]*StockRulesReport{srr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockRulesReport creates a new stock.rules.report model and returns its id.
func (c *Client) CreateStockRulesReports(srrs []*StockRulesReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range srrs {
		vv = append(vv, v)
	}
	return c.Create(StockRulesReportModel, vv, nil)
}

// UpdateStockRulesReport updates an existing stock.rules.report record.
func (c *Client) UpdateStockRulesReport(srr *StockRulesReport) error {
	return c.UpdateStockRulesReports([]int64{srr.Id.Get()}, srr)
}

// UpdateStockRulesReports updates existing stock.rules.report records.
// All records (represented by ids) will be updated by srr values.
func (c *Client) UpdateStockRulesReports(ids []int64, srr *StockRulesReport) error {
	return c.Update(StockRulesReportModel, ids, srr, nil)
}

// DeleteStockRulesReport deletes an existing stock.rules.report record.
func (c *Client) DeleteStockRulesReport(id int64) error {
	return c.DeleteStockRulesReports([]int64{id})
}

// DeleteStockRulesReports deletes existing stock.rules.report records.
func (c *Client) DeleteStockRulesReports(ids []int64) error {
	return c.Delete(StockRulesReportModel, ids)
}

// GetStockRulesReport gets stock.rules.report existing record.
func (c *Client) GetStockRulesReport(id int64) (*StockRulesReport, error) {
	srrs, err := c.GetStockRulesReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srrs)[0]), nil
}

// GetStockRulesReports gets stock.rules.report existing records.
func (c *Client) GetStockRulesReports(ids []int64) (*StockRulesReports, error) {
	srrs := &StockRulesReports{}
	if err := c.Read(StockRulesReportModel, ids, nil, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindStockRulesReport finds stock.rules.report record by querying it with criteria.
func (c *Client) FindStockRulesReport(criteria *Criteria) (*StockRulesReport, error) {
	srrs := &StockRulesReports{}
	if err := c.SearchRead(StockRulesReportModel, criteria, NewOptions().Limit(1), srrs); err != nil {
		return nil, err
	}
	return &((*srrs)[0]), nil
}

// FindStockRulesReports finds stock.rules.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRulesReports(criteria *Criteria, options *Options) (*StockRulesReports, error) {
	srrs := &StockRulesReports{}
	if err := c.SearchRead(StockRulesReportModel, criteria, options, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindStockRulesReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRulesReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockRulesReportModel, criteria, options)
}

// FindStockRulesReportId finds record id by querying it with criteria.
func (c *Client) FindStockRulesReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRulesReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
