package odoo

// ReportAccountReportJournal represents report.account.report_journal model.
type ReportAccountReportJournal struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// ReportAccountReportJournals represents array of report.account.report_journal model.
type ReportAccountReportJournals []ReportAccountReportJournal

// ReportAccountReportJournalModel is the odoo model name.
const ReportAccountReportJournalModel = "report.account.report_journal"

// Many2One convert ReportAccountReportJournal to *Many2One.
func (rar *ReportAccountReportJournal) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportJournal creates a new report.account.report_journal model and returns its id.
func (c *Client) CreateReportAccountReportJournal(rar *ReportAccountReportJournal) (int64, error) {
	ids, err := c.CreateReportAccountReportJournals([]*ReportAccountReportJournal{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportJournal creates a new report.account.report_journal model and returns its id.
func (c *Client) CreateReportAccountReportJournals(rars []*ReportAccountReportJournal) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportJournalModel, vv, nil)
}

// UpdateReportAccountReportJournal updates an existing report.account.report_journal record.
func (c *Client) UpdateReportAccountReportJournal(rar *ReportAccountReportJournal) error {
	return c.UpdateReportAccountReportJournals([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportJournals updates existing report.account.report_journal records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportJournals(ids []int64, rar *ReportAccountReportJournal) error {
	return c.Update(ReportAccountReportJournalModel, ids, rar, nil)
}

// DeleteReportAccountReportJournal deletes an existing report.account.report_journal record.
func (c *Client) DeleteReportAccountReportJournal(id int64) error {
	return c.DeleteReportAccountReportJournals([]int64{id})
}

// DeleteReportAccountReportJournals deletes existing report.account.report_journal records.
func (c *Client) DeleteReportAccountReportJournals(ids []int64) error {
	return c.Delete(ReportAccountReportJournalModel, ids)
}

// GetReportAccountReportJournal gets report.account.report_journal existing record.
func (c *Client) GetReportAccountReportJournal(id int64) (*ReportAccountReportJournal, error) {
	rars, err := c.GetReportAccountReportJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// GetReportAccountReportJournals gets report.account.report_journal existing records.
func (c *Client) GetReportAccountReportJournals(ids []int64) (*ReportAccountReportJournals, error) {
	rars := &ReportAccountReportJournals{}
	if err := c.Read(ReportAccountReportJournalModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportJournal finds report.account.report_journal record by querying it with criteria.
func (c *Client) FindReportAccountReportJournal(criteria *Criteria) (*ReportAccountReportJournal, error) {
	rars := &ReportAccountReportJournals{}
	if err := c.SearchRead(ReportAccountReportJournalModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// FindReportAccountReportJournals finds report.account.report_journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportJournals(criteria *Criteria, options *Options) (*ReportAccountReportJournals, error) {
	rars := &ReportAccountReportJournals{}
	if err := c.SearchRead(ReportAccountReportJournalModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAccountReportJournalModel, criteria, options)
}

// FindReportAccountReportJournalId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
