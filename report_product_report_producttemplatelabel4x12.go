package odoo

// ReportProductReportProducttemplatelabel4X12 represents report.product.report_producttemplatelabel4x12 model.
type ReportProductReportProducttemplatelabel4X12 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportProductReportProducttemplatelabel4X12s represents array of report.product.report_producttemplatelabel4x12 model.
type ReportProductReportProducttemplatelabel4X12s []ReportProductReportProducttemplatelabel4X12

// ReportProductReportProducttemplatelabel4X12Model is the odoo model name.
const ReportProductReportProducttemplatelabel4X12Model = "report.product.report_producttemplatelabel4x12"

// Many2One convert ReportProductReportProducttemplatelabel4X12 to *Many2One.
func (rpr *ReportProductReportProducttemplatelabel4X12) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportProducttemplatelabel4X12 creates a new report.product.report_producttemplatelabel4x12 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X12(rpr *ReportProductReportProducttemplatelabel4X12) (int64, error) {
	ids, err := c.CreateReportProductReportProducttemplatelabel4X12s([]*ReportProductReportProducttemplatelabel4X12{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportProducttemplatelabel4X12 creates a new report.product.report_producttemplatelabel4x12 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X12s(rprs []*ReportProductReportProducttemplatelabel4X12) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportProducttemplatelabel4X12Model, vv, nil)
}

// UpdateReportProductReportProducttemplatelabel4X12 updates an existing report.product.report_producttemplatelabel4x12 record.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X12(rpr *ReportProductReportProducttemplatelabel4X12) error {
	return c.UpdateReportProductReportProducttemplatelabel4X12s([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportProducttemplatelabel4X12s updates existing report.product.report_producttemplatelabel4x12 records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X12s(ids []int64, rpr *ReportProductReportProducttemplatelabel4X12) error {
	return c.Update(ReportProductReportProducttemplatelabel4X12Model, ids, rpr, nil)
}

// DeleteReportProductReportProducttemplatelabel4X12 deletes an existing report.product.report_producttemplatelabel4x12 record.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X12(id int64) error {
	return c.DeleteReportProductReportProducttemplatelabel4X12s([]int64{id})
}

// DeleteReportProductReportProducttemplatelabel4X12s deletes existing report.product.report_producttemplatelabel4x12 records.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X12s(ids []int64) error {
	return c.Delete(ReportProductReportProducttemplatelabel4X12Model, ids)
}

// GetReportProductReportProducttemplatelabel4X12 gets report.product.report_producttemplatelabel4x12 existing record.
func (c *Client) GetReportProductReportProducttemplatelabel4X12(id int64) (*ReportProductReportProducttemplatelabel4X12, error) {
	rprs, err := c.GetReportProductReportProducttemplatelabel4X12s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportProductReportProducttemplatelabel4X12s gets report.product.report_producttemplatelabel4x12 existing records.
func (c *Client) GetReportProductReportProducttemplatelabel4X12s(ids []int64) (*ReportProductReportProducttemplatelabel4X12s, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12s{}
	if err := c.Read(ReportProductReportProducttemplatelabel4X12Model, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X12 finds report.product.report_producttemplatelabel4x12 record by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X12(criteria *Criteria) (*ReportProductReportProducttemplatelabel4X12, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X12Model, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportProductReportProducttemplatelabel4X12s finds report.product.report_producttemplatelabel4x12 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X12s(criteria *Criteria, options *Options) (*ReportProductReportProducttemplatelabel4X12s, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X12Model, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X12Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X12Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProductReportProducttemplatelabel4X12Model, criteria, options)
}

// FindReportProductReportProducttemplatelabel4X12Id finds record id by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X12Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportProducttemplatelabel4X12Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
