package odoo

// ReportProductReportProducttemplatelabel2X7 represents report.product.report_producttemplatelabel2x7 model.
type ReportProductReportProducttemplatelabel2X7 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportProductReportProducttemplatelabel2X7s represents array of report.product.report_producttemplatelabel2x7 model.
type ReportProductReportProducttemplatelabel2X7s []ReportProductReportProducttemplatelabel2X7

// ReportProductReportProducttemplatelabel2X7Model is the odoo model name.
const ReportProductReportProducttemplatelabel2X7Model = "report.product.report_producttemplatelabel2x7"

// Many2One convert ReportProductReportProducttemplatelabel2X7 to *Many2One.
func (rpr *ReportProductReportProducttemplatelabel2X7) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportProducttemplatelabel2X7 creates a new report.product.report_producttemplatelabel2x7 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel2X7(rpr *ReportProductReportProducttemplatelabel2X7) (int64, error) {
	ids, err := c.CreateReportProductReportProducttemplatelabel2X7s([]*ReportProductReportProducttemplatelabel2X7{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportProducttemplatelabel2X7 creates a new report.product.report_producttemplatelabel2x7 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel2X7s(rprs []*ReportProductReportProducttemplatelabel2X7) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportProducttemplatelabel2X7Model, vv, nil)
}

// UpdateReportProductReportProducttemplatelabel2X7 updates an existing report.product.report_producttemplatelabel2x7 record.
func (c *Client) UpdateReportProductReportProducttemplatelabel2X7(rpr *ReportProductReportProducttemplatelabel2X7) error {
	return c.UpdateReportProductReportProducttemplatelabel2X7s([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportProducttemplatelabel2X7s updates existing report.product.report_producttemplatelabel2x7 records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportProducttemplatelabel2X7s(ids []int64, rpr *ReportProductReportProducttemplatelabel2X7) error {
	return c.Update(ReportProductReportProducttemplatelabel2X7Model, ids, rpr, nil)
}

// DeleteReportProductReportProducttemplatelabel2X7 deletes an existing report.product.report_producttemplatelabel2x7 record.
func (c *Client) DeleteReportProductReportProducttemplatelabel2X7(id int64) error {
	return c.DeleteReportProductReportProducttemplatelabel2X7s([]int64{id})
}

// DeleteReportProductReportProducttemplatelabel2X7s deletes existing report.product.report_producttemplatelabel2x7 records.
func (c *Client) DeleteReportProductReportProducttemplatelabel2X7s(ids []int64) error {
	return c.Delete(ReportProductReportProducttemplatelabel2X7Model, ids)
}

// GetReportProductReportProducttemplatelabel2X7 gets report.product.report_producttemplatelabel2x7 existing record.
func (c *Client) GetReportProductReportProducttemplatelabel2X7(id int64) (*ReportProductReportProducttemplatelabel2X7, error) {
	rprs, err := c.GetReportProductReportProducttemplatelabel2X7s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportProductReportProducttemplatelabel2X7s gets report.product.report_producttemplatelabel2x7 existing records.
func (c *Client) GetReportProductReportProducttemplatelabel2X7s(ids []int64) (*ReportProductReportProducttemplatelabel2X7s, error) {
	rprs := &ReportProductReportProducttemplatelabel2X7s{}
	if err := c.Read(ReportProductReportProducttemplatelabel2X7Model, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel2X7 finds report.product.report_producttemplatelabel2x7 record by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel2X7(criteria *Criteria) (*ReportProductReportProducttemplatelabel2X7, error) {
	rprs := &ReportProductReportProducttemplatelabel2X7s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel2X7Model, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportProductReportProducttemplatelabel2X7s finds report.product.report_producttemplatelabel2x7 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel2X7s(criteria *Criteria, options *Options) (*ReportProductReportProducttemplatelabel2X7s, error) {
	rprs := &ReportProductReportProducttemplatelabel2X7s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel2X7Model, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel2X7Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel2X7Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProductReportProducttemplatelabel2X7Model, criteria, options)
}

// FindReportProductReportProducttemplatelabel2X7Id finds record id by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel2X7Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportProducttemplatelabel2X7Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
