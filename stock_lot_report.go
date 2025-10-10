package odoo

// StockLotReport represents stock.lot.report model.
type StockLotReport struct {
	Address      *String   `xmlrpc:"address,omitempty"`
	DeliveryDate *Time     `xmlrpc:"delivery_date,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	HasReturn    *Bool     `xmlrpc:"has_return,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	LotId        *Many2One `xmlrpc:"lot_id,omitempty"`
	PartnerId    *Many2One `xmlrpc:"partner_id,omitempty"`
	PickingId    *Many2One `xmlrpc:"picking_id,omitempty"`
	ProductId    *Many2One `xmlrpc:"product_id,omitempty"`
	Quantity     *Float    `xmlrpc:"quantity,omitempty"`
	UomId        *Many2One `xmlrpc:"uom_id,omitempty"`
}

// StockLotReports represents array of stock.lot.report model.
type StockLotReports []StockLotReport

// StockLotReportModel is the odoo model name.
const StockLotReportModel = "stock.lot.report"

// Many2One convert StockLotReport to *Many2One.
func (slr *StockLotReport) Many2One() *Many2One {
	return NewMany2One(slr.Id.Get(), "")
}

// CreateStockLotReport creates a new stock.lot.report model and returns its id.
func (c *Client) CreateStockLotReport(slr *StockLotReport) (int64, error) {
	ids, err := c.CreateStockLotReports([]*StockLotReport{slr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLotReport creates a new stock.lot.report model and returns its id.
func (c *Client) CreateStockLotReports(slrs []*StockLotReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range slrs {
		vv = append(vv, v)
	}
	return c.Create(StockLotReportModel, vv, nil)
}

// UpdateStockLotReport updates an existing stock.lot.report record.
func (c *Client) UpdateStockLotReport(slr *StockLotReport) error {
	return c.UpdateStockLotReports([]int64{slr.Id.Get()}, slr)
}

// UpdateStockLotReports updates existing stock.lot.report records.
// All records (represented by ids) will be updated by slr values.
func (c *Client) UpdateStockLotReports(ids []int64, slr *StockLotReport) error {
	return c.Update(StockLotReportModel, ids, slr, nil)
}

// DeleteStockLotReport deletes an existing stock.lot.report record.
func (c *Client) DeleteStockLotReport(id int64) error {
	return c.DeleteStockLotReports([]int64{id})
}

// DeleteStockLotReports deletes existing stock.lot.report records.
func (c *Client) DeleteStockLotReports(ids []int64) error {
	return c.Delete(StockLotReportModel, ids)
}

// GetStockLotReport gets stock.lot.report existing record.
func (c *Client) GetStockLotReport(id int64) (*StockLotReport, error) {
	slrs, err := c.GetStockLotReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slrs)[0]), nil
}

// GetStockLotReports gets stock.lot.report existing records.
func (c *Client) GetStockLotReports(ids []int64) (*StockLotReports, error) {
	slrs := &StockLotReports{}
	if err := c.Read(StockLotReportModel, ids, nil, slrs); err != nil {
		return nil, err
	}
	return slrs, nil
}

// FindStockLotReport finds stock.lot.report record by querying it with criteria.
func (c *Client) FindStockLotReport(criteria *Criteria) (*StockLotReport, error) {
	slrs := &StockLotReports{}
	if err := c.SearchRead(StockLotReportModel, criteria, NewOptions().Limit(1), slrs); err != nil {
		return nil, err
	}
	return &((*slrs)[0]), nil
}

// FindStockLotReports finds stock.lot.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLotReports(criteria *Criteria, options *Options) (*StockLotReports, error) {
	slrs := &StockLotReports{}
	if err := c.SearchRead(StockLotReportModel, criteria, options, slrs); err != nil {
		return nil, err
	}
	return slrs, nil
}

// FindStockLotReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLotReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockLotReportModel, criteria, options)
}

// FindStockLotReportId finds record id by querying it with criteria.
func (c *Client) FindStockLotReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLotReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
