package odoo

// ReportStockLabelLotTemplateView represents report.stock.label_lot_template_view model.
type ReportStockLabelLotTemplateView struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportStockLabelLotTemplateViews represents array of report.stock.label_lot_template_view model.
type ReportStockLabelLotTemplateViews []ReportStockLabelLotTemplateView

// ReportStockLabelLotTemplateViewModel is the odoo model name.
const ReportStockLabelLotTemplateViewModel = "report.stock.label_lot_template_view"

// Many2One convert ReportStockLabelLotTemplateView to *Many2One.
func (rsl *ReportStockLabelLotTemplateView) Many2One() *Many2One {
	return NewMany2One(rsl.Id.Get(), "")
}

// CreateReportStockLabelLotTemplateView creates a new report.stock.label_lot_template_view model and returns its id.
func (c *Client) CreateReportStockLabelLotTemplateView(rsl *ReportStockLabelLotTemplateView) (int64, error) {
	ids, err := c.CreateReportStockLabelLotTemplateViews([]*ReportStockLabelLotTemplateView{rsl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportStockLabelLotTemplateView creates a new report.stock.label_lot_template_view model and returns its id.
func (c *Client) CreateReportStockLabelLotTemplateViews(rsls []*ReportStockLabelLotTemplateView) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsls {
		vv = append(vv, v)
	}
	return c.Create(ReportStockLabelLotTemplateViewModel, vv, nil)
}

// UpdateReportStockLabelLotTemplateView updates an existing report.stock.label_lot_template_view record.
func (c *Client) UpdateReportStockLabelLotTemplateView(rsl *ReportStockLabelLotTemplateView) error {
	return c.UpdateReportStockLabelLotTemplateViews([]int64{rsl.Id.Get()}, rsl)
}

// UpdateReportStockLabelLotTemplateViews updates existing report.stock.label_lot_template_view records.
// All records (represented by ids) will be updated by rsl values.
func (c *Client) UpdateReportStockLabelLotTemplateViews(ids []int64, rsl *ReportStockLabelLotTemplateView) error {
	return c.Update(ReportStockLabelLotTemplateViewModel, ids, rsl, nil)
}

// DeleteReportStockLabelLotTemplateView deletes an existing report.stock.label_lot_template_view record.
func (c *Client) DeleteReportStockLabelLotTemplateView(id int64) error {
	return c.DeleteReportStockLabelLotTemplateViews([]int64{id})
}

// DeleteReportStockLabelLotTemplateViews deletes existing report.stock.label_lot_template_view records.
func (c *Client) DeleteReportStockLabelLotTemplateViews(ids []int64) error {
	return c.Delete(ReportStockLabelLotTemplateViewModel, ids)
}

// GetReportStockLabelLotTemplateView gets report.stock.label_lot_template_view existing record.
func (c *Client) GetReportStockLabelLotTemplateView(id int64) (*ReportStockLabelLotTemplateView, error) {
	rsls, err := c.GetReportStockLabelLotTemplateViews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsls)[0]), nil
}

// GetReportStockLabelLotTemplateViews gets report.stock.label_lot_template_view existing records.
func (c *Client) GetReportStockLabelLotTemplateViews(ids []int64) (*ReportStockLabelLotTemplateViews, error) {
	rsls := &ReportStockLabelLotTemplateViews{}
	if err := c.Read(ReportStockLabelLotTemplateViewModel, ids, nil, rsls); err != nil {
		return nil, err
	}
	return rsls, nil
}

// FindReportStockLabelLotTemplateView finds report.stock.label_lot_template_view record by querying it with criteria.
func (c *Client) FindReportStockLabelLotTemplateView(criteria *Criteria) (*ReportStockLabelLotTemplateView, error) {
	rsls := &ReportStockLabelLotTemplateViews{}
	if err := c.SearchRead(ReportStockLabelLotTemplateViewModel, criteria, NewOptions().Limit(1), rsls); err != nil {
		return nil, err
	}
	return &((*rsls)[0]), nil
}

// FindReportStockLabelLotTemplateViews finds report.stock.label_lot_template_view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockLabelLotTemplateViews(criteria *Criteria, options *Options) (*ReportStockLabelLotTemplateViews, error) {
	rsls := &ReportStockLabelLotTemplateViews{}
	if err := c.SearchRead(ReportStockLabelLotTemplateViewModel, criteria, options, rsls); err != nil {
		return nil, err
	}
	return rsls, nil
}

// FindReportStockLabelLotTemplateViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockLabelLotTemplateViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportStockLabelLotTemplateViewModel, criteria, options)
}

// FindReportStockLabelLotTemplateViewId finds record id by querying it with criteria.
func (c *Client) FindReportStockLabelLotTemplateViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockLabelLotTemplateViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
