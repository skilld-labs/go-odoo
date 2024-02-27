package odoo

// ReportAccountReportPartnerledger represents report.account.report_partnerledger model.
type ReportAccountReportPartnerledger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportPartnerledgers represents array of report.account.report_partnerledger model.
type ReportAccountReportPartnerledgers []ReportAccountReportPartnerledger

// ReportAccountReportPartnerledgerModel is the odoo model name.
const ReportAccountReportPartnerledgerModel = "report.account.report_partnerledger"

// Many2One convert ReportAccountReportPartnerledger to *Many2One.
func (rar *ReportAccountReportPartnerledger) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportPartnerledger creates a new report.account.report_partnerledger model and returns its id.
func (c *Client) CreateReportAccountReportPartnerledger(rar *ReportAccountReportPartnerledger) (int64, error) {
	ids, err := c.CreateReportAccountReportPartnerledgers([]*ReportAccountReportPartnerledger{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportPartnerledger creates a new report.account.report_partnerledger model and returns its id.
func (c *Client) CreateReportAccountReportPartnerledgers(rars []*ReportAccountReportPartnerledger) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportPartnerledgerModel, vv, nil)
}

// UpdateReportAccountReportPartnerledger updates an existing report.account.report_partnerledger record.
func (c *Client) UpdateReportAccountReportPartnerledger(rar *ReportAccountReportPartnerledger) error {
	return c.UpdateReportAccountReportPartnerledgers([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportPartnerledgers updates existing report.account.report_partnerledger records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportPartnerledgers(ids []int64, rar *ReportAccountReportPartnerledger) error {
	return c.Update(ReportAccountReportPartnerledgerModel, ids, rar, nil)
}

// DeleteReportAccountReportPartnerledger deletes an existing report.account.report_partnerledger record.
func (c *Client) DeleteReportAccountReportPartnerledger(id int64) error {
	return c.DeleteReportAccountReportPartnerledgers([]int64{id})
}

// DeleteReportAccountReportPartnerledgers deletes existing report.account.report_partnerledger records.
func (c *Client) DeleteReportAccountReportPartnerledgers(ids []int64) error {
	return c.Delete(ReportAccountReportPartnerledgerModel, ids)
}

// GetReportAccountReportPartnerledger gets report.account.report_partnerledger existing record.
func (c *Client) GetReportAccountReportPartnerledger(id int64) (*ReportAccountReportPartnerledger, error) {
	rars, err := c.GetReportAccountReportPartnerledgers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// GetReportAccountReportPartnerledgers gets report.account.report_partnerledger existing records.
func (c *Client) GetReportAccountReportPartnerledgers(ids []int64) (*ReportAccountReportPartnerledgers, error) {
	rars := &ReportAccountReportPartnerledgers{}
	if err := c.Read(ReportAccountReportPartnerledgerModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportPartnerledger finds report.account.report_partnerledger record by querying it with criteria.
func (c *Client) FindReportAccountReportPartnerledger(criteria *Criteria) (*ReportAccountReportPartnerledger, error) {
	rars := &ReportAccountReportPartnerledgers{}
	if err := c.SearchRead(ReportAccountReportPartnerledgerModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// FindReportAccountReportPartnerledgers finds report.account.report_partnerledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportPartnerledgers(criteria *Criteria, options *Options) (*ReportAccountReportPartnerledgers, error) {
	rars := &ReportAccountReportPartnerledgers{}
	if err := c.SearchRead(ReportAccountReportPartnerledgerModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportPartnerledgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportPartnerledgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAccountReportPartnerledgerModel, criteria, options)
}

// FindReportAccountReportPartnerledgerId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportPartnerledgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportPartnerledgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
