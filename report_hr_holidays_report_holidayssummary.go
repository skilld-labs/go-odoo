package odoo

// ReportHrHolidaysReportHolidayssummary represents report.hr_holidays.report_holidayssummary model.
type ReportHrHolidaysReportHolidayssummary struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// ReportHrHolidaysReportHolidayssummarys represents array of report.hr_holidays.report_holidayssummary model.
type ReportHrHolidaysReportHolidayssummarys []ReportHrHolidaysReportHolidayssummary

// ReportHrHolidaysReportHolidayssummaryModel is the odoo model name.
const ReportHrHolidaysReportHolidayssummaryModel = "report.hr_holidays.report_holidayssummary"

// Many2One convert ReportHrHolidaysReportHolidayssummary to *Many2One.
func (rhr *ReportHrHolidaysReportHolidayssummary) Many2One() *Many2One {
	return NewMany2One(rhr.Id.Get(), "")
}

// CreateReportHrHolidaysReportHolidayssummary creates a new report.hr_holidays.report_holidayssummary model and returns its id.
func (c *Client) CreateReportHrHolidaysReportHolidayssummary(rhr *ReportHrHolidaysReportHolidayssummary) (int64, error) {
	ids, err := c.CreateReportHrHolidaysReportHolidayssummarys([]*ReportHrHolidaysReportHolidayssummary{rhr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportHrHolidaysReportHolidayssummary creates a new report.hr_holidays.report_holidayssummary model and returns its id.
func (c *Client) CreateReportHrHolidaysReportHolidayssummarys(rhrs []*ReportHrHolidaysReportHolidayssummary) ([]int64, error) {
	var vv []interface{}
	for _, v := range rhrs {
		vv = append(vv, v)
	}
	return c.Create(ReportHrHolidaysReportHolidayssummaryModel, vv, nil)
}

// UpdateReportHrHolidaysReportHolidayssummary updates an existing report.hr_holidays.report_holidayssummary record.
func (c *Client) UpdateReportHrHolidaysReportHolidayssummary(rhr *ReportHrHolidaysReportHolidayssummary) error {
	return c.UpdateReportHrHolidaysReportHolidayssummarys([]int64{rhr.Id.Get()}, rhr)
}

// UpdateReportHrHolidaysReportHolidayssummarys updates existing report.hr_holidays.report_holidayssummary records.
// All records (represented by ids) will be updated by rhr values.
func (c *Client) UpdateReportHrHolidaysReportHolidayssummarys(ids []int64, rhr *ReportHrHolidaysReportHolidayssummary) error {
	return c.Update(ReportHrHolidaysReportHolidayssummaryModel, ids, rhr, nil)
}

// DeleteReportHrHolidaysReportHolidayssummary deletes an existing report.hr_holidays.report_holidayssummary record.
func (c *Client) DeleteReportHrHolidaysReportHolidayssummary(id int64) error {
	return c.DeleteReportHrHolidaysReportHolidayssummarys([]int64{id})
}

// DeleteReportHrHolidaysReportHolidayssummarys deletes existing report.hr_holidays.report_holidayssummary records.
func (c *Client) DeleteReportHrHolidaysReportHolidayssummarys(ids []int64) error {
	return c.Delete(ReportHrHolidaysReportHolidayssummaryModel, ids)
}

// GetReportHrHolidaysReportHolidayssummary gets report.hr_holidays.report_holidayssummary existing record.
func (c *Client) GetReportHrHolidaysReportHolidayssummary(id int64) (*ReportHrHolidaysReportHolidayssummary, error) {
	rhrs, err := c.GetReportHrHolidaysReportHolidayssummarys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rhrs)[0]), nil
}

// GetReportHrHolidaysReportHolidayssummarys gets report.hr_holidays.report_holidayssummary existing records.
func (c *Client) GetReportHrHolidaysReportHolidayssummarys(ids []int64) (*ReportHrHolidaysReportHolidayssummarys, error) {
	rhrs := &ReportHrHolidaysReportHolidayssummarys{}
	if err := c.Read(ReportHrHolidaysReportHolidayssummaryModel, ids, nil, rhrs); err != nil {
		return nil, err
	}
	return rhrs, nil
}

// FindReportHrHolidaysReportHolidayssummary finds report.hr_holidays.report_holidayssummary record by querying it with criteria.
func (c *Client) FindReportHrHolidaysReportHolidayssummary(criteria *Criteria) (*ReportHrHolidaysReportHolidayssummary, error) {
	rhrs := &ReportHrHolidaysReportHolidayssummarys{}
	if err := c.SearchRead(ReportHrHolidaysReportHolidayssummaryModel, criteria, NewOptions().Limit(1), rhrs); err != nil {
		return nil, err
	}
	return &((*rhrs)[0]), nil
}

// FindReportHrHolidaysReportHolidayssummarys finds report.hr_holidays.report_holidayssummary records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportHrHolidaysReportHolidayssummarys(criteria *Criteria, options *Options) (*ReportHrHolidaysReportHolidayssummarys, error) {
	rhrs := &ReportHrHolidaysReportHolidayssummarys{}
	if err := c.SearchRead(ReportHrHolidaysReportHolidayssummaryModel, criteria, options, rhrs); err != nil {
		return nil, err
	}
	return rhrs, nil
}

// FindReportHrHolidaysReportHolidayssummaryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportHrHolidaysReportHolidayssummaryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportHrHolidaysReportHolidayssummaryModel, criteria, options)
}

// FindReportHrHolidaysReportHolidayssummaryId finds record id by querying it with criteria.
func (c *Client) FindReportHrHolidaysReportHolidayssummaryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportHrHolidaysReportHolidayssummaryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
