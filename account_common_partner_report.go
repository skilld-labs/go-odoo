package odoo

// AccountCommonPartnerReport represents account.common.partner.report model.
type AccountCommonPartnerReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omitempty"`
	DateTo          *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	JournalIds      *Relation  `xmlrpc:"journal_ids,omitempty"`
	ResultSelection *Selection `xmlrpc:"result_selection,omitempty"`
	TargetMove      *Selection `xmlrpc:"target_move,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountCommonPartnerReports represents array of account.common.partner.report model.
type AccountCommonPartnerReports []AccountCommonPartnerReport

// AccountCommonPartnerReportModel is the odoo model name.
const AccountCommonPartnerReportModel = "account.common.partner.report"

// Many2One convert AccountCommonPartnerReport to *Many2One.
func (acpr *AccountCommonPartnerReport) Many2One() *Many2One {
	return NewMany2One(acpr.Id.Get(), "")
}

// CreateAccountCommonPartnerReport creates a new account.common.partner.report model and returns its id.
func (c *Client) CreateAccountCommonPartnerReport(acpr *AccountCommonPartnerReport) (int64, error) {
	ids, err := c.CreateAccountCommonPartnerReports([]*AccountCommonPartnerReport{acpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCommonPartnerReport creates a new account.common.partner.report model and returns its id.
func (c *Client) CreateAccountCommonPartnerReports(acprs []*AccountCommonPartnerReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range acprs {
		vv = append(vv, v)
	}
	return c.Create(AccountCommonPartnerReportModel, vv, nil)
}

// UpdateAccountCommonPartnerReport updates an existing account.common.partner.report record.
func (c *Client) UpdateAccountCommonPartnerReport(acpr *AccountCommonPartnerReport) error {
	return c.UpdateAccountCommonPartnerReports([]int64{acpr.Id.Get()}, acpr)
}

// UpdateAccountCommonPartnerReports updates existing account.common.partner.report records.
// All records (represented by ids) will be updated by acpr values.
func (c *Client) UpdateAccountCommonPartnerReports(ids []int64, acpr *AccountCommonPartnerReport) error {
	return c.Update(AccountCommonPartnerReportModel, ids, acpr, nil)
}

// DeleteAccountCommonPartnerReport deletes an existing account.common.partner.report record.
func (c *Client) DeleteAccountCommonPartnerReport(id int64) error {
	return c.DeleteAccountCommonPartnerReports([]int64{id})
}

// DeleteAccountCommonPartnerReports deletes existing account.common.partner.report records.
func (c *Client) DeleteAccountCommonPartnerReports(ids []int64) error {
	return c.Delete(AccountCommonPartnerReportModel, ids)
}

// GetAccountCommonPartnerReport gets account.common.partner.report existing record.
func (c *Client) GetAccountCommonPartnerReport(id int64) (*AccountCommonPartnerReport, error) {
	acprs, err := c.GetAccountCommonPartnerReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acprs)[0]), nil
}

// GetAccountCommonPartnerReports gets account.common.partner.report existing records.
func (c *Client) GetAccountCommonPartnerReports(ids []int64) (*AccountCommonPartnerReports, error) {
	acprs := &AccountCommonPartnerReports{}
	if err := c.Read(AccountCommonPartnerReportModel, ids, nil, acprs); err != nil {
		return nil, err
	}
	return acprs, nil
}

// FindAccountCommonPartnerReport finds account.common.partner.report record by querying it with criteria.
func (c *Client) FindAccountCommonPartnerReport(criteria *Criteria) (*AccountCommonPartnerReport, error) {
	acprs := &AccountCommonPartnerReports{}
	if err := c.SearchRead(AccountCommonPartnerReportModel, criteria, NewOptions().Limit(1), acprs); err != nil {
		return nil, err
	}
	return &((*acprs)[0]), nil
}

// FindAccountCommonPartnerReports finds account.common.partner.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonPartnerReports(criteria *Criteria, options *Options) (*AccountCommonPartnerReports, error) {
	acprs := &AccountCommonPartnerReports{}
	if err := c.SearchRead(AccountCommonPartnerReportModel, criteria, options, acprs); err != nil {
		return nil, err
	}
	return acprs, nil
}

// FindAccountCommonPartnerReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonPartnerReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCommonPartnerReportModel, criteria, options)
}

// FindAccountCommonPartnerReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCommonPartnerReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCommonPartnerReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
