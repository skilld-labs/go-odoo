package odoo

// ReportStockLabelProductProductView represents report.stock.label_product_product_view model.
type ReportStockLabelProductProductView struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportStockLabelProductProductViews represents array of report.stock.label_product_product_view model.
type ReportStockLabelProductProductViews []ReportStockLabelProductProductView

// ReportStockLabelProductProductViewModel is the odoo model name.
const ReportStockLabelProductProductViewModel = "report.stock.label_product_product_view"

// Many2One convert ReportStockLabelProductProductView to *Many2One.
func (rsl *ReportStockLabelProductProductView) Many2One() *Many2One {
	return NewMany2One(rsl.Id.Get(), "")
}

// CreateReportStockLabelProductProductView creates a new report.stock.label_product_product_view model and returns its id.
func (c *Client) CreateReportStockLabelProductProductView(rsl *ReportStockLabelProductProductView) (int64, error) {
	ids, err := c.CreateReportStockLabelProductProductViews([]*ReportStockLabelProductProductView{rsl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportStockLabelProductProductView creates a new report.stock.label_product_product_view model and returns its id.
func (c *Client) CreateReportStockLabelProductProductViews(rsls []*ReportStockLabelProductProductView) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsls {
		vv = append(vv, v)
	}
	return c.Create(ReportStockLabelProductProductViewModel, vv, nil)
}

// UpdateReportStockLabelProductProductView updates an existing report.stock.label_product_product_view record.
func (c *Client) UpdateReportStockLabelProductProductView(rsl *ReportStockLabelProductProductView) error {
	return c.UpdateReportStockLabelProductProductViews([]int64{rsl.Id.Get()}, rsl)
}

// UpdateReportStockLabelProductProductViews updates existing report.stock.label_product_product_view records.
// All records (represented by ids) will be updated by rsl values.
func (c *Client) UpdateReportStockLabelProductProductViews(ids []int64, rsl *ReportStockLabelProductProductView) error {
	return c.Update(ReportStockLabelProductProductViewModel, ids, rsl, nil)
}

// DeleteReportStockLabelProductProductView deletes an existing report.stock.label_product_product_view record.
func (c *Client) DeleteReportStockLabelProductProductView(id int64) error {
	return c.DeleteReportStockLabelProductProductViews([]int64{id})
}

// DeleteReportStockLabelProductProductViews deletes existing report.stock.label_product_product_view records.
func (c *Client) DeleteReportStockLabelProductProductViews(ids []int64) error {
	return c.Delete(ReportStockLabelProductProductViewModel, ids)
}

// GetReportStockLabelProductProductView gets report.stock.label_product_product_view existing record.
func (c *Client) GetReportStockLabelProductProductView(id int64) (*ReportStockLabelProductProductView, error) {
	rsls, err := c.GetReportStockLabelProductProductViews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsls)[0]), nil
}

// GetReportStockLabelProductProductViews gets report.stock.label_product_product_view existing records.
func (c *Client) GetReportStockLabelProductProductViews(ids []int64) (*ReportStockLabelProductProductViews, error) {
	rsls := &ReportStockLabelProductProductViews{}
	if err := c.Read(ReportStockLabelProductProductViewModel, ids, nil, rsls); err != nil {
		return nil, err
	}
	return rsls, nil
}

// FindReportStockLabelProductProductView finds report.stock.label_product_product_view record by querying it with criteria.
func (c *Client) FindReportStockLabelProductProductView(criteria *Criteria) (*ReportStockLabelProductProductView, error) {
	rsls := &ReportStockLabelProductProductViews{}
	if err := c.SearchRead(ReportStockLabelProductProductViewModel, criteria, NewOptions().Limit(1), rsls); err != nil {
		return nil, err
	}
	return &((*rsls)[0]), nil
}

// FindReportStockLabelProductProductViews finds report.stock.label_product_product_view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockLabelProductProductViews(criteria *Criteria, options *Options) (*ReportStockLabelProductProductViews, error) {
	rsls := &ReportStockLabelProductProductViews{}
	if err := c.SearchRead(ReportStockLabelProductProductViewModel, criteria, options, rsls); err != nil {
		return nil, err
	}
	return rsls, nil
}

// FindReportStockLabelProductProductViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockLabelProductProductViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportStockLabelProductProductViewModel, criteria, options)
}

// FindReportStockLabelProductProductViewId finds record id by querying it with criteria.
func (c *Client) FindReportStockLabelProductProductViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockLabelProductProductViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
