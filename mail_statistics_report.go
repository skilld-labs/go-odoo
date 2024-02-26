package odoo

// MailStatisticsReport represents mail.statistics.report model.
type MailStatisticsReport struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	Bounced       *Int       `xmlrpc:"bounced,omptempty"`
	Campaign      *String    `xmlrpc:"campaign,omptempty"`
	Delivered     *Int       `xmlrpc:"delivered,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom     *String    `xmlrpc:"email_from,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	Opened        *Int       `xmlrpc:"opened,omptempty"`
	Replied       *Int       `xmlrpc:"replied,omptempty"`
	ScheduledDate *Time      `xmlrpc:"scheduled_date,omptempty"`
	Sent          *Int       `xmlrpc:"sent,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
}

// MailStatisticsReports represents array of mail.statistics.report model.
type MailStatisticsReports []MailStatisticsReport

// MailStatisticsReportModel is the odoo model name.
const MailStatisticsReportModel = "mail.statistics.report"

// Many2One convert MailStatisticsReport to *Many2One.
func (msr *MailStatisticsReport) Many2One() *Many2One {
	return NewMany2One(msr.Id.Get(), "")
}

// CreateMailStatisticsReport creates a new mail.statistics.report model and returns its id.
func (c *Client) CreateMailStatisticsReport(msr *MailStatisticsReport) (int64, error) {
	ids, err := c.CreateMailStatisticsReports([]*MailStatisticsReport{msr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailStatisticsReport creates a new mail.statistics.report model and returns its id.
func (c *Client) CreateMailStatisticsReports(msrs []*MailStatisticsReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range msrs {
		vv = append(vv, v)
	}
	return c.Create(MailStatisticsReportModel, vv, nil)
}

// UpdateMailStatisticsReport updates an existing mail.statistics.report record.
func (c *Client) UpdateMailStatisticsReport(msr *MailStatisticsReport) error {
	return c.UpdateMailStatisticsReports([]int64{msr.Id.Get()}, msr)
}

// UpdateMailStatisticsReports updates existing mail.statistics.report records.
// All records (represented by ids) will be updated by msr values.
func (c *Client) UpdateMailStatisticsReports(ids []int64, msr *MailStatisticsReport) error {
	return c.Update(MailStatisticsReportModel, ids, msr, nil)
}

// DeleteMailStatisticsReport deletes an existing mail.statistics.report record.
func (c *Client) DeleteMailStatisticsReport(id int64) error {
	return c.DeleteMailStatisticsReports([]int64{id})
}

// DeleteMailStatisticsReports deletes existing mail.statistics.report records.
func (c *Client) DeleteMailStatisticsReports(ids []int64) error {
	return c.Delete(MailStatisticsReportModel, ids)
}

// GetMailStatisticsReport gets mail.statistics.report existing record.
func (c *Client) GetMailStatisticsReport(id int64) (*MailStatisticsReport, error) {
	msrs, err := c.GetMailStatisticsReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*msrs)[0]), nil
}

// GetMailStatisticsReports gets mail.statistics.report existing records.
func (c *Client) GetMailStatisticsReports(ids []int64) (*MailStatisticsReports, error) {
	msrs := &MailStatisticsReports{}
	if err := c.Read(MailStatisticsReportModel, ids, nil, msrs); err != nil {
		return nil, err
	}
	return msrs, nil
}

// FindMailStatisticsReport finds mail.statistics.report record by querying it with criteria.
func (c *Client) FindMailStatisticsReport(criteria *Criteria) (*MailStatisticsReport, error) {
	msrs := &MailStatisticsReports{}
	if err := c.SearchRead(MailStatisticsReportModel, criteria, NewOptions().Limit(1), msrs); err != nil {
		return nil, err
	}
	return &((*msrs)[0]), nil
}

// FindMailStatisticsReports finds mail.statistics.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailStatisticsReports(criteria *Criteria, options *Options) (*MailStatisticsReports, error) {
	msrs := &MailStatisticsReports{}
	if err := c.SearchRead(MailStatisticsReportModel, criteria, options, msrs); err != nil {
		return nil, err
	}
	return msrs, nil
}

// FindMailStatisticsReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailStatisticsReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailStatisticsReportModel, criteria, options)
}

// FindMailStatisticsReportId finds record id by querying it with criteria.
func (c *Client) FindMailStatisticsReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailStatisticsReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
