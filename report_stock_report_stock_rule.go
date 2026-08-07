package odoo

// ReportStockReportStockRule represents report.stock.report_stock_rule model.
type ReportStockReportStockRule struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportStockReportStockRules represents array of report.stock.report_stock_rule model.
type ReportStockReportStockRules []ReportStockReportStockRule

// ReportStockReportStockRuleModel is the odoo model name.
const ReportStockReportStockRuleModel = "report.stock.report_stock_rule"

// Many2One convert ReportStockReportStockRule to *Many2One.
func (rsr *ReportStockReportStockRule) Many2One() *Many2One {
	return NewMany2One(rsr.Id.Get(), "")
}

// CreateReportStockReportStockRule creates a new report.stock.report_stock_rule model and returns its id.
func (c *Client) CreateReportStockReportStockRule(rsr *ReportStockReportStockRule) (int64, error) {
	ids, err := c.CreateReportStockReportStockRules([]*ReportStockReportStockRule{rsr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportStockReportStockRule creates a new report.stock.report_stock_rule model and returns its id.
func (c *Client) CreateReportStockReportStockRules(rsrs []*ReportStockReportStockRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsrs {
		vv = append(vv, v)
	}
	return c.Create(ReportStockReportStockRuleModel, vv, nil)
}

// UpdateReportStockReportStockRule updates an existing report.stock.report_stock_rule record.
func (c *Client) UpdateReportStockReportStockRule(rsr *ReportStockReportStockRule) error {
	return c.UpdateReportStockReportStockRules([]int64{rsr.Id.Get()}, rsr)
}

// UpdateReportStockReportStockRules updates existing report.stock.report_stock_rule records.
// All records (represented by ids) will be updated by rsr values.
func (c *Client) UpdateReportStockReportStockRules(ids []int64, rsr *ReportStockReportStockRule) error {
	return c.Update(ReportStockReportStockRuleModel, ids, rsr, nil)
}

// DeleteReportStockReportStockRule deletes an existing report.stock.report_stock_rule record.
func (c *Client) DeleteReportStockReportStockRule(id int64) error {
	return c.DeleteReportStockReportStockRules([]int64{id})
}

// DeleteReportStockReportStockRules deletes existing report.stock.report_stock_rule records.
func (c *Client) DeleteReportStockReportStockRules(ids []int64) error {
	return c.Delete(ReportStockReportStockRuleModel, ids)
}

// GetReportStockReportStockRule gets report.stock.report_stock_rule existing record.
func (c *Client) GetReportStockReportStockRule(id int64) (*ReportStockReportStockRule, error) {
	rsrs, err := c.GetReportStockReportStockRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// GetReportStockReportStockRules gets report.stock.report_stock_rule existing records.
func (c *Client) GetReportStockReportStockRules(ids []int64) (*ReportStockReportStockRules, error) {
	rsrs := &ReportStockReportStockRules{}
	if err := c.Read(ReportStockReportStockRuleModel, ids, nil, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportStockReportStockRule finds report.stock.report_stock_rule record by querying it with criteria.
func (c *Client) FindReportStockReportStockRule(criteria *Criteria) (*ReportStockReportStockRule, error) {
	rsrs := &ReportStockReportStockRules{}
	if err := c.SearchRead(ReportStockReportStockRuleModel, criteria, NewOptions().Limit(1), rsrs); err != nil {
		return nil, err
	}
	return &((*rsrs)[0]), nil
}

// FindReportStockReportStockRules finds report.stock.report_stock_rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockReportStockRules(criteria *Criteria, options *Options) (*ReportStockReportStockRules, error) {
	rsrs := &ReportStockReportStockRules{}
	if err := c.SearchRead(ReportStockReportStockRuleModel, criteria, options, rsrs); err != nil {
		return nil, err
	}
	return rsrs, nil
}

// FindReportStockReportStockRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockReportStockRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportStockReportStockRuleModel, criteria, options)
}

// FindReportStockReportStockRuleId finds record id by querying it with criteria.
func (c *Client) FindReportStockReportStockRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockReportStockRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
