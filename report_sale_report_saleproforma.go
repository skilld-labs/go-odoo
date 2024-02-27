package odoo

// ReportSaleReportSaleproforma represents report.sale.report_saleproforma model.
type ReportSaleReportSaleproforma struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportSaleReportSaleproformas represents array of report.sale.report_saleproforma model.
type ReportSaleReportSaleproformas []ReportSaleReportSaleproforma

// ReportSaleReportSaleproformaModel is the odoo model name.
const ReportSaleReportSaleproformaModel = "report.sale.report_saleproforma"

// Many2One convert ReportSaleReportSaleproforma to *Many2One.
func (rsr *ReportSaleReportSaleproforma) Many2One() *Many2One {
	return NewMany2One(rsr.Id.Get(), "")
}

// CreateReportSaleReportSaleproforma creates a new report.sale.report_saleproforma model and returns its id.
func (c *Client) CreateReportSaleReportSaleproforma(rsr *ReportSaleReportSaleproforma) (int64, error) {
	ids, err := c.CreateReportSaleReportSaleproformas([]*ReportSaleReportSaleproforma{rsr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportSaleReportSaleproforma creates a new report.sale.report_saleproforma model and returns its id.
func (c *Client) CreateReportSaleReportSaleproformas(rsrs []*ReportSaleReportSaleproforma) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsrs {
		vv = append(vv, v)
	}
	return c.Create(ReportSaleReportSaleproformaModel, vv, nil)
}

// UpdateReportSaleReportSaleproforma updates an existing report.sale.report_saleproforma record.
func (c *Client) UpdateReportSaleReportSaleproforma(rsr *ReportSaleReportSaleproforma) error {
	return c.UpdateReportSaleReportSaleproformas([]int64{rsr.Id.Get()}, rsr)
}

// UpdateReportSaleReportSaleproformas updates existing report.sale.report_saleproforma records.
// All records (represented by ids) will be updated by rsr values.
func (c *Client) UpdateReportSaleReportSaleproformas(ids []int64, rsr *ReportSaleReportSaleproforma) error {
	return c.Update(ReportSaleReportSaleproformaModel, ids, rsr, nil)
}

// DeleteReportSaleReportSaleproforma deletes an existing report.sale.report_saleproforma record.
func (c *Client) DeleteReportSaleReportSaleproforma(id int64) error {
	return c.DeleteReportSaleReportSaleproformas([]int64{id})
}

// DeleteReportSaleReportSaleproformas deletes existing report.sale.report_saleproforma records.
func (c *Client) DeleteReportSaleReportSaleproformas(ids []int64) error {
	return c.Delete(ReportSaleReportSaleproformaModel, ids)
}

// GetReportSaleReportSaleproforma gets report.sale.report_saleproforma existing record.
func (c *Client) GetReportSaleReportSaleproforma(id int64) (*ReportSaleReportSaleproforma, error) {
	rsrs, err := c.GetReportSaleReportSaleproformas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// GetReportSaleReportSaleproformas gets report.sale.report_saleproforma existing records.
func (c *Client) GetReportSaleReportSaleproformas(ids []int64) (*ReportSaleReportSaleproformas, error) {
	rsrs := &ReportSaleReportSaleproformas{}
	if err := c.Read(ReportSaleReportSaleproformaModel, ids, nil, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportSaleReportSaleproforma finds report.sale.report_saleproforma record by querying it with criteria.
func (c *Client) FindReportSaleReportSaleproforma(criteria *Criteria) (*ReportSaleReportSaleproforma, error) {
	rsrs := &ReportSaleReportSaleproformas{}
	if err := c.SearchRead(ReportSaleReportSaleproformaModel, criteria, NewOptions().Limit(1), rsrs); err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// FindReportSaleReportSaleproformas finds report.sale.report_saleproforma records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportSaleReportSaleproformas(criteria *Criteria, options *Options) (*ReportSaleReportSaleproformas, error) {
	rsrs := &ReportSaleReportSaleproformas{}
	if err := c.SearchRead(ReportSaleReportSaleproformaModel, criteria, options, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportSaleReportSaleproformaIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportSaleReportSaleproformaIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportSaleReportSaleproformaModel, criteria, options)
}

// FindReportSaleReportSaleproformaId finds record id by querying it with criteria.
func (c *Client) FindReportSaleReportSaleproformaId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportSaleReportSaleproformaModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
