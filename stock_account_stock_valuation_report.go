package odoo

// StockAccountStockValuationReport represents stock_account.stock.valuation.report model.
type StockAccountStockValuationReport struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// StockAccountStockValuationReports represents array of stock_account.stock.valuation.report model.
type StockAccountStockValuationReports []StockAccountStockValuationReport

// StockAccountStockValuationReportModel is the odoo model name.
const StockAccountStockValuationReportModel = "stock_account.stock.valuation.report"

// Many2One convert StockAccountStockValuationReport to *Many2One.
func (ssvr *StockAccountStockValuationReport) Many2One() *Many2One {
	return NewMany2One(ssvr.Id.Get(), "")
}

// CreateStockAccountStockValuationReport creates a new stock_account.stock.valuation.report model and returns its id.
func (c *Client) CreateStockAccountStockValuationReport(ssvr *StockAccountStockValuationReport) (int64, error) {
	ids, err := c.CreateStockAccountStockValuationReports([]*StockAccountStockValuationReport{ssvr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockAccountStockValuationReport creates a new stock_account.stock.valuation.report model and returns its id.
func (c *Client) CreateStockAccountStockValuationReports(ssvrs []*StockAccountStockValuationReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssvrs {
		vv = append(vv, v)
	}
	return c.Create(StockAccountStockValuationReportModel, vv, nil)
}

// UpdateStockAccountStockValuationReport updates an existing stock_account.stock.valuation.report record.
func (c *Client) UpdateStockAccountStockValuationReport(ssvr *StockAccountStockValuationReport) error {
	return c.UpdateStockAccountStockValuationReports([]int64{ssvr.Id.Get()}, ssvr)
}

// UpdateStockAccountStockValuationReports updates existing stock_account.stock.valuation.report records.
// All records (represented by ids) will be updated by ssvr values.
func (c *Client) UpdateStockAccountStockValuationReports(ids []int64, ssvr *StockAccountStockValuationReport) error {
	return c.Update(StockAccountStockValuationReportModel, ids, ssvr, nil)
}

// DeleteStockAccountStockValuationReport deletes an existing stock_account.stock.valuation.report record.
func (c *Client) DeleteStockAccountStockValuationReport(id int64) error {
	return c.DeleteStockAccountStockValuationReports([]int64{id})
}

// DeleteStockAccountStockValuationReports deletes existing stock_account.stock.valuation.report records.
func (c *Client) DeleteStockAccountStockValuationReports(ids []int64) error {
	return c.Delete(StockAccountStockValuationReportModel, ids)
}

// GetStockAccountStockValuationReport gets stock_account.stock.valuation.report existing record.
func (c *Client) GetStockAccountStockValuationReport(id int64) (*StockAccountStockValuationReport, error) {
	ssvrs, err := c.GetStockAccountStockValuationReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssvrs)[0]), nil
}

// GetStockAccountStockValuationReports gets stock_account.stock.valuation.report existing records.
func (c *Client) GetStockAccountStockValuationReports(ids []int64) (*StockAccountStockValuationReports, error) {
	ssvrs := &StockAccountStockValuationReports{}
	if err := c.Read(StockAccountStockValuationReportModel, ids, nil, ssvrs); err != nil {
		return nil, err
	}
	return ssvrs, nil
}

// FindStockAccountStockValuationReport finds stock_account.stock.valuation.report record by querying it with criteria.
func (c *Client) FindStockAccountStockValuationReport(criteria *Criteria) (*StockAccountStockValuationReport, error) {
	ssvrs := &StockAccountStockValuationReports{}
	if err := c.SearchRead(StockAccountStockValuationReportModel, criteria, NewOptions().Limit(1), ssvrs); err != nil {
		return nil, err
	}
	return &((*ssvrs)[0]), nil
}

// FindStockAccountStockValuationReports finds stock_account.stock.valuation.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAccountStockValuationReports(criteria *Criteria, options *Options) (*StockAccountStockValuationReports, error) {
	ssvrs := &StockAccountStockValuationReports{}
	if err := c.SearchRead(StockAccountStockValuationReportModel, criteria, options, ssvrs); err != nil {
		return nil, err
	}
	return ssvrs, nil
}

// FindStockAccountStockValuationReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAccountStockValuationReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockAccountStockValuationReportModel, criteria, options)
}

// FindStockAccountStockValuationReportId finds record id by querying it with criteria.
func (c *Client) FindStockAccountStockValuationReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockAccountStockValuationReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
