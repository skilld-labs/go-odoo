package odoo

// ReportBaseReportIrmodulereference represents report.base.report_irmodulereference model.
type ReportBaseReportIrmodulereference struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportBaseReportIrmodulereferences represents array of report.base.report_irmodulereference model.
type ReportBaseReportIrmodulereferences []ReportBaseReportIrmodulereference

// ReportBaseReportIrmodulereferenceModel is the odoo model name.
const ReportBaseReportIrmodulereferenceModel = "report.base.report_irmodulereference"

// Many2One convert ReportBaseReportIrmodulereference to *Many2One.
func (rbr *ReportBaseReportIrmodulereference) Many2One() *Many2One {
	return NewMany2One(rbr.Id.Get(), "")
}

// CreateReportBaseReportIrmodulereference creates a new report.base.report_irmodulereference model and returns its id.
func (c *Client) CreateReportBaseReportIrmodulereference(rbr *ReportBaseReportIrmodulereference) (int64, error) {
	ids, err := c.CreateReportBaseReportIrmodulereferences([]*ReportBaseReportIrmodulereference{rbr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportBaseReportIrmodulereference creates a new report.base.report_irmodulereference model and returns its id.
func (c *Client) CreateReportBaseReportIrmodulereferences(rbrs []*ReportBaseReportIrmodulereference) ([]int64, error) {
	var vv []interface{}
	for _, v := range rbrs {
		vv = append(vv, v)
	}
	return c.Create(ReportBaseReportIrmodulereferenceModel, vv, nil)
}

// UpdateReportBaseReportIrmodulereference updates an existing report.base.report_irmodulereference record.
func (c *Client) UpdateReportBaseReportIrmodulereference(rbr *ReportBaseReportIrmodulereference) error {
	return c.UpdateReportBaseReportIrmodulereferences([]int64{rbr.Id.Get()}, rbr)
}

// UpdateReportBaseReportIrmodulereferences updates existing report.base.report_irmodulereference records.
// All records (represented by ids) will be updated by rbr values.
func (c *Client) UpdateReportBaseReportIrmodulereferences(ids []int64, rbr *ReportBaseReportIrmodulereference) error {
	return c.Update(ReportBaseReportIrmodulereferenceModel, ids, rbr, nil)
}

// DeleteReportBaseReportIrmodulereference deletes an existing report.base.report_irmodulereference record.
func (c *Client) DeleteReportBaseReportIrmodulereference(id int64) error {
	return c.DeleteReportBaseReportIrmodulereferences([]int64{id})
}

// DeleteReportBaseReportIrmodulereferences deletes existing report.base.report_irmodulereference records.
func (c *Client) DeleteReportBaseReportIrmodulereferences(ids []int64) error {
	return c.Delete(ReportBaseReportIrmodulereferenceModel, ids)
}

// GetReportBaseReportIrmodulereference gets report.base.report_irmodulereference existing record.
func (c *Client) GetReportBaseReportIrmodulereference(id int64) (*ReportBaseReportIrmodulereference, error) {
	rbrs, err := c.GetReportBaseReportIrmodulereferences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rbrs)[0]), nil
}

// GetReportBaseReportIrmodulereferences gets report.base.report_irmodulereference existing records.
func (c *Client) GetReportBaseReportIrmodulereferences(ids []int64) (*ReportBaseReportIrmodulereferences, error) {
	rbrs := &ReportBaseReportIrmodulereferences{}
	if err := c.Read(ReportBaseReportIrmodulereferenceModel, ids, nil, rbrs); err != nil {
		return nil, err
	}
	return rbrs, nil
}

// FindReportBaseReportIrmodulereference finds report.base.report_irmodulereference record by querying it with criteria.
func (c *Client) FindReportBaseReportIrmodulereference(criteria *Criteria) (*ReportBaseReportIrmodulereference, error) {
	rbrs := &ReportBaseReportIrmodulereferences{}
	if err := c.SearchRead(ReportBaseReportIrmodulereferenceModel, criteria, NewOptions().Limit(1), rbrs); err != nil {
		return nil, err
	}
	return &((*rbrs)[0]), nil
}

// FindReportBaseReportIrmodulereferences finds report.base.report_irmodulereference records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportBaseReportIrmodulereferences(criteria *Criteria, options *Options) (*ReportBaseReportIrmodulereferences, error) {
	rbrs := &ReportBaseReportIrmodulereferences{}
	if err := c.SearchRead(ReportBaseReportIrmodulereferenceModel, criteria, options, rbrs); err != nil {
		return nil, err
	}
	return rbrs, nil
}

// FindReportBaseReportIrmodulereferenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportBaseReportIrmodulereferenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportBaseReportIrmodulereferenceModel, criteria, options)
}

// FindReportBaseReportIrmodulereferenceId finds record id by querying it with criteria.
func (c *Client) FindReportBaseReportIrmodulereferenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportBaseReportIrmodulereferenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
