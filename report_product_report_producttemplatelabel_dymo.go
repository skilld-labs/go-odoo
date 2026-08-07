package odoo

// ReportProductReportProducttemplatelabelDymo represents report.product.report_producttemplatelabel_dymo model.
type ReportProductReportProducttemplatelabelDymo struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportProductReportProducttemplatelabelDymos represents array of report.product.report_producttemplatelabel_dymo model.
type ReportProductReportProducttemplatelabelDymos []ReportProductReportProducttemplatelabelDymo

// ReportProductReportProducttemplatelabelDymoModel is the odoo model name.
const ReportProductReportProducttemplatelabelDymoModel = "report.product.report_producttemplatelabel_dymo"

// Many2One convert ReportProductReportProducttemplatelabelDymo to *Many2One.
func (rpr *ReportProductReportProducttemplatelabelDymo) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportProducttemplatelabelDymo creates a new report.product.report_producttemplatelabel_dymo model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabelDymo(rpr *ReportProductReportProducttemplatelabelDymo) (int64, error) {
	ids, err := c.CreateReportProductReportProducttemplatelabelDymos([]*ReportProductReportProducttemplatelabelDymo{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportProducttemplatelabelDymo creates a new report.product.report_producttemplatelabel_dymo model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabelDymos(rprs []*ReportProductReportProducttemplatelabelDymo) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportProducttemplatelabelDymoModel, vv, nil)
}

// UpdateReportProductReportProducttemplatelabelDymo updates an existing report.product.report_producttemplatelabel_dymo record.
func (c *Client) UpdateReportProductReportProducttemplatelabelDymo(rpr *ReportProductReportProducttemplatelabelDymo) error {
	return c.UpdateReportProductReportProducttemplatelabelDymos([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportProducttemplatelabelDymos updates existing report.product.report_producttemplatelabel_dymo records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportProducttemplatelabelDymos(ids []int64, rpr *ReportProductReportProducttemplatelabelDymo) error {
	return c.Update(ReportProductReportProducttemplatelabelDymoModel, ids, rpr, nil)
}

// DeleteReportProductReportProducttemplatelabelDymo deletes an existing report.product.report_producttemplatelabel_dymo record.
func (c *Client) DeleteReportProductReportProducttemplatelabelDymo(id int64) error {
	return c.DeleteReportProductReportProducttemplatelabelDymos([]int64{id})
}

// DeleteReportProductReportProducttemplatelabelDymos deletes existing report.product.report_producttemplatelabel_dymo records.
func (c *Client) DeleteReportProductReportProducttemplatelabelDymos(ids []int64) error {
	return c.Delete(ReportProductReportProducttemplatelabelDymoModel, ids)
}

// GetReportProductReportProducttemplatelabelDymo gets report.product.report_producttemplatelabel_dymo existing record.
func (c *Client) GetReportProductReportProducttemplatelabelDymo(id int64) (*ReportProductReportProducttemplatelabelDymo, error) {
	rprs, err := c.GetReportProductReportProducttemplatelabelDymos([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportProductReportProducttemplatelabelDymos gets report.product.report_producttemplatelabel_dymo existing records.
func (c *Client) GetReportProductReportProducttemplatelabelDymos(ids []int64) (*ReportProductReportProducttemplatelabelDymos, error) {
	rprs := &ReportProductReportProducttemplatelabelDymos{}
	if err := c.Read(ReportProductReportProducttemplatelabelDymoModel, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabelDymo finds report.product.report_producttemplatelabel_dymo record by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabelDymo(criteria *Criteria) (*ReportProductReportProducttemplatelabelDymo, error) {
	rprs := &ReportProductReportProducttemplatelabelDymos{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabelDymoModel, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportProductReportProducttemplatelabelDymos finds report.product.report_producttemplatelabel_dymo records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabelDymos(criteria *Criteria, options *Options) (*ReportProductReportProducttemplatelabelDymos, error) {
	rprs := &ReportProductReportProducttemplatelabelDymos{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabelDymoModel, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabelDymoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabelDymoIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProductReportProducttemplatelabelDymoModel, criteria, options)
}

// FindReportProductReportProducttemplatelabelDymoId finds record id by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabelDymoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportProducttemplatelabelDymoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
