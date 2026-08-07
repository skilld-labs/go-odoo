package odoo

// StockAvcoReport represents stock.avco.report model.
type StockAvcoReport struct {
	AddedValue    *Float     `xmlrpc:"added_value,omitempty"`
	AvcoValue     *Float     `xmlrpc:"avco_value,omitempty"`
	CompanyId     *Many2One  `xmlrpc:"company_id,omitempty"`
	CurrencyId    *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date          *Time      `xmlrpc:"date,omitempty"`
	Description   *String    `xmlrpc:"description,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	Justification *String    `xmlrpc:"justification,omitempty"`
	ProductId     *Many2One  `xmlrpc:"product_id,omitempty"`
	Quantity      *Float     `xmlrpc:"quantity,omitempty"`
	Reference     *String    `xmlrpc:"reference,omitempty"`
	ResModelName  *Selection `xmlrpc:"res_model_name,omitempty"`
	TotalQuantity *Float     `xmlrpc:"total_quantity,omitempty"`
	TotalValue    *Float     `xmlrpc:"total_value,omitempty"`
	UserId        *Many2One  `xmlrpc:"user_id,omitempty"`
	Value         *Float     `xmlrpc:"value,omitempty"`
}

// StockAvcoReports represents array of stock.avco.report model.
type StockAvcoReports []StockAvcoReport

// StockAvcoReportModel is the odoo model name.
const StockAvcoReportModel = "stock.avco.report"

// Many2One convert StockAvcoReport to *Many2One.
func (sar *StockAvcoReport) Many2One() *Many2One {
	return NewMany2One(sar.Id.Get(), "")
}

// CreateStockAvcoReport creates a new stock.avco.report model and returns its id.
func (c *Client) CreateStockAvcoReport(sar *StockAvcoReport) (int64, error) {
	ids, err := c.CreateStockAvcoReports([]*StockAvcoReport{sar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockAvcoReport creates a new stock.avco.report model and returns its id.
func (c *Client) CreateStockAvcoReports(sars []*StockAvcoReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range sars {
		vv = append(vv, v)
	}
	return c.Create(StockAvcoReportModel, vv, nil)
}

// UpdateStockAvcoReport updates an existing stock.avco.report record.
func (c *Client) UpdateStockAvcoReport(sar *StockAvcoReport) error {
	return c.UpdateStockAvcoReports([]int64{sar.Id.Get()}, sar)
}

// UpdateStockAvcoReports updates existing stock.avco.report records.
// All records (represented by ids) will be updated by sar values.
func (c *Client) UpdateStockAvcoReports(ids []int64, sar *StockAvcoReport) error {
	return c.Update(StockAvcoReportModel, ids, sar, nil)
}

// DeleteStockAvcoReport deletes an existing stock.avco.report record.
func (c *Client) DeleteStockAvcoReport(id int64) error {
	return c.DeleteStockAvcoReports([]int64{id})
}

// DeleteStockAvcoReports deletes existing stock.avco.report records.
func (c *Client) DeleteStockAvcoReports(ids []int64) error {
	return c.Delete(StockAvcoReportModel, ids)
}

// GetStockAvcoReport gets stock.avco.report existing record.
func (c *Client) GetStockAvcoReport(id int64) (*StockAvcoReport, error) {
	sars, err := c.GetStockAvcoReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// GetStockAvcoReports gets stock.avco.report existing records.
func (c *Client) GetStockAvcoReports(ids []int64) (*StockAvcoReports, error) {
	sars := &StockAvcoReports{}
	if err := c.Read(StockAvcoReportModel, ids, nil, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStockAvcoReport finds stock.avco.report record by querying it with criteria.
func (c *Client) FindStockAvcoReport(criteria *Criteria) (*StockAvcoReport, error) {
	sars := &StockAvcoReports{}
	if err := c.SearchRead(StockAvcoReportModel, criteria, NewOptions().Limit(1), sars); err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// FindStockAvcoReports finds stock.avco.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAvcoReports(criteria *Criteria, options *Options) (*StockAvcoReports, error) {
	sars := &StockAvcoReports{}
	if err := c.SearchRead(StockAvcoReportModel, criteria, options, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStockAvcoReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAvcoReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockAvcoReportModel, criteria, options)
}

// FindStockAvcoReportId finds record id by querying it with criteria.
func (c *Client) FindStockAvcoReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockAvcoReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
