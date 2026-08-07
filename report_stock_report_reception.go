package odoo

// ReportStockReportReception represents report.stock.report_reception model.
type ReportStockReportReception struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportStockReportReceptions represents array of report.stock.report_reception model.
type ReportStockReportReceptions []ReportStockReportReception

// ReportStockReportReceptionModel is the odoo model name.
const ReportStockReportReceptionModel = "report.stock.report_reception"

// Many2One convert ReportStockReportReception to *Many2One.
func (rsr *ReportStockReportReception) Many2One() *Many2One {
	return NewMany2One(rsr.Id.Get(), "")
}

// CreateReportStockReportReception creates a new report.stock.report_reception model and returns its id.
func (c *Client) CreateReportStockReportReception(rsr *ReportStockReportReception) (int64, error) {
	ids, err := c.CreateReportStockReportReceptions([]*ReportStockReportReception{rsr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportStockReportReception creates a new report.stock.report_reception model and returns its id.
func (c *Client) CreateReportStockReportReceptions(rsrs []*ReportStockReportReception) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsrs {
		vv = append(vv, v)
	}
	return c.Create(ReportStockReportReceptionModel, vv, nil)
}

// UpdateReportStockReportReception updates an existing report.stock.report_reception record.
func (c *Client) UpdateReportStockReportReception(rsr *ReportStockReportReception) error {
	return c.UpdateReportStockReportReceptions([]int64{rsr.Id.Get()}, rsr)
}

// UpdateReportStockReportReceptions updates existing report.stock.report_reception records.
// All records (represented by ids) will be updated by rsr values.
func (c *Client) UpdateReportStockReportReceptions(ids []int64, rsr *ReportStockReportReception) error {
	return c.Update(ReportStockReportReceptionModel, ids, rsr, nil)
}

// DeleteReportStockReportReception deletes an existing report.stock.report_reception record.
func (c *Client) DeleteReportStockReportReception(id int64) error {
	return c.DeleteReportStockReportReceptions([]int64{id})
}

// DeleteReportStockReportReceptions deletes existing report.stock.report_reception records.
func (c *Client) DeleteReportStockReportReceptions(ids []int64) error {
	return c.Delete(ReportStockReportReceptionModel, ids)
}

// GetReportStockReportReception gets report.stock.report_reception existing record.
func (c *Client) GetReportStockReportReception(id int64) (*ReportStockReportReception, error) {
	rsrs, err := c.GetReportStockReportReceptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// GetReportStockReportReceptions gets report.stock.report_reception existing records.
func (c *Client) GetReportStockReportReceptions(ids []int64) (*ReportStockReportReceptions, error) {
	rsrs := &ReportStockReportReceptions{}
	if err := c.Read(ReportStockReportReceptionModel, ids, nil, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportStockReportReception finds report.stock.report_reception record by querying it with criteria.
func (c *Client) FindReportStockReportReception(criteria *Criteria) (*ReportStockReportReception, error) {
	rsrs := &ReportStockReportReceptions{}
	if err := c.SearchRead(ReportStockReportReceptionModel, criteria, NewOptions().Limit(1), rsrs); err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// FindReportStockReportReceptions finds report.stock.report_reception records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockReportReceptions(criteria *Criteria, options *Options) (*ReportStockReportReceptions, error) {
	rsrs := &ReportStockReportReceptions{}
	if err := c.SearchRead(ReportStockReportReceptionModel, criteria, options, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportStockReportReceptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockReportReceptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportStockReportReceptionModel, criteria, options)
}

// FindReportStockReportReceptionId finds record id by querying it with criteria.
func (c *Client) FindReportStockReportReceptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockReportReceptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
