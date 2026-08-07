package odoo

// ReportProductReportProducttemplatelabel4X7 represents report.product.report_producttemplatelabel4x7 model.
type ReportProductReportProducttemplatelabel4X7 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportProductReportProducttemplatelabel4X7s represents array of report.product.report_producttemplatelabel4x7 model.
type ReportProductReportProducttemplatelabel4X7s []ReportProductReportProducttemplatelabel4X7

// ReportProductReportProducttemplatelabel4X7Model is the odoo model name.
const ReportProductReportProducttemplatelabel4X7Model = "report.product.report_producttemplatelabel4x7"

// Many2One convert ReportProductReportProducttemplatelabel4X7 to *Many2One.
func (rpr *ReportProductReportProducttemplatelabel4X7) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportProducttemplatelabel4X7 creates a new report.product.report_producttemplatelabel4x7 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X7(rpr *ReportProductReportProducttemplatelabel4X7) (int64, error) {
	ids, err := c.CreateReportProductReportProducttemplatelabel4X7s([]*ReportProductReportProducttemplatelabel4X7{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportProducttemplatelabel4X7 creates a new report.product.report_producttemplatelabel4x7 model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X7s(rprs []*ReportProductReportProducttemplatelabel4X7) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportProducttemplatelabel4X7Model, vv, nil)
}

// UpdateReportProductReportProducttemplatelabel4X7 updates an existing report.product.report_producttemplatelabel4x7 record.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X7(rpr *ReportProductReportProducttemplatelabel4X7) error {
	return c.UpdateReportProductReportProducttemplatelabel4X7s([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportProducttemplatelabel4X7s updates existing report.product.report_producttemplatelabel4x7 records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X7s(ids []int64, rpr *ReportProductReportProducttemplatelabel4X7) error {
	return c.Update(ReportProductReportProducttemplatelabel4X7Model, ids, rpr, nil)
}

// DeleteReportProductReportProducttemplatelabel4X7 deletes an existing report.product.report_producttemplatelabel4x7 record.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X7(id int64) error {
	return c.DeleteReportProductReportProducttemplatelabel4X7s([]int64{id})
}

// DeleteReportProductReportProducttemplatelabel4X7s deletes existing report.product.report_producttemplatelabel4x7 records.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X7s(ids []int64) error {
	return c.Delete(ReportProductReportProducttemplatelabel4X7Model, ids)
}

// GetReportProductReportProducttemplatelabel4X7 gets report.product.report_producttemplatelabel4x7 existing record.
func (c *Client) GetReportProductReportProducttemplatelabel4X7(id int64) (*ReportProductReportProducttemplatelabel4X7, error) {
	rprs, err := c.GetReportProductReportProducttemplatelabel4X7s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportProductReportProducttemplatelabel4X7s gets report.product.report_producttemplatelabel4x7 existing records.
func (c *Client) GetReportProductReportProducttemplatelabel4X7s(ids []int64) (*ReportProductReportProducttemplatelabel4X7s, error) {
	rprs := &ReportProductReportProducttemplatelabel4X7s{}
	if err := c.Read(ReportProductReportProducttemplatelabel4X7Model, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X7 finds report.product.report_producttemplatelabel4x7 record by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X7(criteria *Criteria) (*ReportProductReportProducttemplatelabel4X7, error) {
	rprs := &ReportProductReportProducttemplatelabel4X7s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X7Model, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportProductReportProducttemplatelabel4X7s finds report.product.report_producttemplatelabel4x7 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X7s(criteria *Criteria, options *Options) (*ReportProductReportProducttemplatelabel4X7s, error) {
	rprs := &ReportProductReportProducttemplatelabel4X7s{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X7Model, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X7Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X7Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProductReportProducttemplatelabel4X7Model, criteria, options)
}

// FindReportProductReportProducttemplatelabel4X7Id finds record id by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X7Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportProducttemplatelabel4X7Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
