package odoo

// MailingTraceReport represents mailing.trace.report model.
type MailingTraceReport struct {
	Bounced       *Int       `xmlrpc:"bounced,omitempty"`
	Campaign      *String    `xmlrpc:"campaign,omitempty"`
	Canceled      *Int       `xmlrpc:"canceled,omitempty"`
	Clicked       *Int       `xmlrpc:"clicked,omitempty"`
	Delivered     *Int       `xmlrpc:"delivered,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom     *String    `xmlrpc:"email_from,omitempty"`
	Error         *Int       `xmlrpc:"error,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	MailingType   *Selection `xmlrpc:"mailing_type,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	Opened        *Int       `xmlrpc:"opened,omitempty"`
	Pending       *Int       `xmlrpc:"pending,omitempty"`
	Processing    *Int       `xmlrpc:"processing,omitempty"`
	Replied       *Int       `xmlrpc:"replied,omitempty"`
	Scheduled     *Int       `xmlrpc:"scheduled,omitempty"`
	ScheduledDate *Time      `xmlrpc:"scheduled_date,omitempty"`
	Sent          *Int       `xmlrpc:"sent,omitempty"`
	State         *Selection `xmlrpc:"state,omitempty"`
}

// MailingTraceReports represents array of mailing.trace.report model.
type MailingTraceReports []MailingTraceReport

// MailingTraceReportModel is the odoo model name.
const MailingTraceReportModel = "mailing.trace.report"

// Many2One convert MailingTraceReport to *Many2One.
func (mtr *MailingTraceReport) Many2One() *Many2One {
	return NewMany2One(mtr.Id.Get(), "")
}

// CreateMailingTraceReport creates a new mailing.trace.report model and returns its id.
func (c *Client) CreateMailingTraceReport(mtr *MailingTraceReport) (int64, error) {
	ids, err := c.CreateMailingTraceReports([]*MailingTraceReport{mtr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingTraceReport creates a new mailing.trace.report model and returns its id.
func (c *Client) CreateMailingTraceReports(mtrs []*MailingTraceReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtrs {
		vv = append(vv, v)
	}
	return c.Create(MailingTraceReportModel, vv, nil)
}

// UpdateMailingTraceReport updates an existing mailing.trace.report record.
func (c *Client) UpdateMailingTraceReport(mtr *MailingTraceReport) error {
	return c.UpdateMailingTraceReports([]int64{mtr.Id.Get()}, mtr)
}

// UpdateMailingTraceReports updates existing mailing.trace.report records.
// All records (represented by ids) will be updated by mtr values.
func (c *Client) UpdateMailingTraceReports(ids []int64, mtr *MailingTraceReport) error {
	return c.Update(MailingTraceReportModel, ids, mtr, nil)
}

// DeleteMailingTraceReport deletes an existing mailing.trace.report record.
func (c *Client) DeleteMailingTraceReport(id int64) error {
	return c.DeleteMailingTraceReports([]int64{id})
}

// DeleteMailingTraceReports deletes existing mailing.trace.report records.
func (c *Client) DeleteMailingTraceReports(ids []int64) error {
	return c.Delete(MailingTraceReportModel, ids)
}

// GetMailingTraceReport gets mailing.trace.report existing record.
func (c *Client) GetMailingTraceReport(id int64) (*MailingTraceReport, error) {
	mtrs, err := c.GetMailingTraceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mtrs)[0]), nil
}

// GetMailingTraceReports gets mailing.trace.report existing records.
func (c *Client) GetMailingTraceReports(ids []int64) (*MailingTraceReports, error) {
	mtrs := &MailingTraceReports{}
	if err := c.Read(MailingTraceReportModel, ids, nil, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailingTraceReport finds mailing.trace.report record by querying it with criteria.
func (c *Client) FindMailingTraceReport(criteria *Criteria) (*MailingTraceReport, error) {
	mtrs := &MailingTraceReports{}
	if err := c.SearchRead(MailingTraceReportModel, criteria, NewOptions().Limit(1), mtrs); err != nil {
		return nil, err
	}
	return &((*mtrs)[0]), nil
}

// FindMailingTraceReports finds mailing.trace.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceReports(criteria *Criteria, options *Options) (*MailingTraceReports, error) {
	mtrs := &MailingTraceReports{}
	if err := c.SearchRead(MailingTraceReportModel, criteria, options, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailingTraceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingTraceReportModel, criteria, options)
}

// FindMailingTraceReportId finds record id by querying it with criteria.
func (c *Client) FindMailingTraceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingTraceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
