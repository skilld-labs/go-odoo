package odoo

// AccountCommonAccountReport represents account.common.account.report model.
type AccountCommonAccountReport struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omptempty"`
	DateTo         *Time      `xmlrpc:"date_to,omptempty"`
	DisplayAccount *Selection `xmlrpc:"display_account,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	JournalIds     *Relation  `xmlrpc:"journal_ids,omptempty"`
	TargetMove     *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountCommonAccountReports represents array of account.common.account.report model.
type AccountCommonAccountReports []AccountCommonAccountReport

// AccountCommonAccountReportModel is the odoo model name.
const AccountCommonAccountReportModel = "account.common.account.report"

// Many2One convert AccountCommonAccountReport to *Many2One.
func (acar *AccountCommonAccountReport) Many2One() *Many2One {
	return NewMany2One(acar.Id.Get(), "")
}

// CreateAccountCommonAccountReport creates a new account.common.account.report model and returns its id.
func (c *Client) CreateAccountCommonAccountReport(acar *AccountCommonAccountReport) (int64, error) {
	ids, err := c.CreateAccountCommonAccountReports([]*AccountCommonAccountReport{acar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCommonAccountReport creates a new account.common.account.report model and returns its id.
func (c *Client) CreateAccountCommonAccountReports(acars []*AccountCommonAccountReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range acars {
		vv = append(vv, v)
	}
	return c.Create(AccountCommonAccountReportModel, vv, nil)
}

// UpdateAccountCommonAccountReport updates an existing account.common.account.report record.
func (c *Client) UpdateAccountCommonAccountReport(acar *AccountCommonAccountReport) error {
	return c.UpdateAccountCommonAccountReports([]int64{acar.Id.Get()}, acar)
}

// UpdateAccountCommonAccountReports updates existing account.common.account.report records.
// All records (represented by ids) will be updated by acar values.
func (c *Client) UpdateAccountCommonAccountReports(ids []int64, acar *AccountCommonAccountReport) error {
	return c.Update(AccountCommonAccountReportModel, ids, acar, nil)
}

// DeleteAccountCommonAccountReport deletes an existing account.common.account.report record.
func (c *Client) DeleteAccountCommonAccountReport(id int64) error {
	return c.DeleteAccountCommonAccountReports([]int64{id})
}

// DeleteAccountCommonAccountReports deletes existing account.common.account.report records.
func (c *Client) DeleteAccountCommonAccountReports(ids []int64) error {
	return c.Delete(AccountCommonAccountReportModel, ids)
}

// GetAccountCommonAccountReport gets account.common.account.report existing record.
func (c *Client) GetAccountCommonAccountReport(id int64) (*AccountCommonAccountReport, error) {
	acars, err := c.GetAccountCommonAccountReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acars)[0]), nil
}

// GetAccountCommonAccountReports gets account.common.account.report existing records.
func (c *Client) GetAccountCommonAccountReports(ids []int64) (*AccountCommonAccountReports, error) {
	acars := &AccountCommonAccountReports{}
	if err := c.Read(AccountCommonAccountReportModel, ids, nil, acars); err != nil {
		return nil, err
	}
	return acars, nil
}

// FindAccountCommonAccountReport finds account.common.account.report record by querying it with criteria.
func (c *Client) FindAccountCommonAccountReport(criteria *Criteria) (*AccountCommonAccountReport, error) {
	acars := &AccountCommonAccountReports{}
	if err := c.SearchRead(AccountCommonAccountReportModel, criteria, NewOptions().Limit(1), acars); err != nil {
		return nil, err
	}
	return &((*acars)[0]), nil
}

// FindAccountCommonAccountReports finds account.common.account.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonAccountReports(criteria *Criteria, options *Options) (*AccountCommonAccountReports, error) {
	acars := &AccountCommonAccountReports{}
	if err := c.SearchRead(AccountCommonAccountReportModel, criteria, options, acars); err != nil {
		return nil, err
	}
	return acars, nil
}

// FindAccountCommonAccountReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonAccountReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCommonAccountReportModel, criteria, options)
}

// FindAccountCommonAccountReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCommonAccountReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCommonAccountReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
