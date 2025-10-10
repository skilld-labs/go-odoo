package odoo

// AccountReportColumn represents account.report.column model.
type AccountReportColumn struct {
	BlankIfZero         *Bool      `xmlrpc:"blank_if_zero,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomAuditActionId *Many2One  `xmlrpc:"custom_audit_action_id,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	ExpressionLabel     *String    `xmlrpc:"expression_label,omitempty"`
	FigureType          *Selection `xmlrpc:"figure_type,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	ReportId            *Many2One  `xmlrpc:"report_id,omitempty"`
	Sequence            *Int       `xmlrpc:"sequence,omitempty"`
	Sortable            *Bool      `xmlrpc:"sortable,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReportColumns represents array of account.report.column model.
type AccountReportColumns []AccountReportColumn

// AccountReportColumnModel is the odoo model name.
const AccountReportColumnModel = "account.report.column"

// Many2One convert AccountReportColumn to *Many2One.
func (arc *AccountReportColumn) Many2One() *Many2One {
	return NewMany2One(arc.Id.Get(), "")
}

// CreateAccountReportColumn creates a new account.report.column model and returns its id.
func (c *Client) CreateAccountReportColumn(arc *AccountReportColumn) (int64, error) {
	ids, err := c.CreateAccountReportColumns([]*AccountReportColumn{arc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportColumn creates a new account.report.column model and returns its id.
func (c *Client) CreateAccountReportColumns(arcs []*AccountReportColumn) ([]int64, error) {
	var vv []interface{}
	for _, v := range arcs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportColumnModel, vv, nil)
}

// UpdateAccountReportColumn updates an existing account.report.column record.
func (c *Client) UpdateAccountReportColumn(arc *AccountReportColumn) error {
	return c.UpdateAccountReportColumns([]int64{arc.Id.Get()}, arc)
}

// UpdateAccountReportColumns updates existing account.report.column records.
// All records (represented by ids) will be updated by arc values.
func (c *Client) UpdateAccountReportColumns(ids []int64, arc *AccountReportColumn) error {
	return c.Update(AccountReportColumnModel, ids, arc, nil)
}

// DeleteAccountReportColumn deletes an existing account.report.column record.
func (c *Client) DeleteAccountReportColumn(id int64) error {
	return c.DeleteAccountReportColumns([]int64{id})
}

// DeleteAccountReportColumns deletes existing account.report.column records.
func (c *Client) DeleteAccountReportColumns(ids []int64) error {
	return c.Delete(AccountReportColumnModel, ids)
}

// GetAccountReportColumn gets account.report.column existing record.
func (c *Client) GetAccountReportColumn(id int64) (*AccountReportColumn, error) {
	arcs, err := c.GetAccountReportColumns([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arcs)[0]), nil
}

// GetAccountReportColumns gets account.report.column existing records.
func (c *Client) GetAccountReportColumns(ids []int64) (*AccountReportColumns, error) {
	arcs := &AccountReportColumns{}
	if err := c.Read(AccountReportColumnModel, ids, nil, arcs); err != nil {
		return nil, err
	}
	return arcs, nil
}

// FindAccountReportColumn finds account.report.column record by querying it with criteria.
func (c *Client) FindAccountReportColumn(criteria *Criteria) (*AccountReportColumn, error) {
	arcs := &AccountReportColumns{}
	if err := c.SearchRead(AccountReportColumnModel, criteria, NewOptions().Limit(1), arcs); err != nil {
		return nil, err
	}
	return &((*arcs)[0]), nil
}

// FindAccountReportColumns finds account.report.column records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportColumns(criteria *Criteria, options *Options) (*AccountReportColumns, error) {
	arcs := &AccountReportColumns{}
	if err := c.SearchRead(AccountReportColumnModel, criteria, options, arcs); err != nil {
		return nil, err
	}
	return arcs, nil
}

// FindAccountReportColumnIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportColumnIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportColumnModel, criteria, options)
}

// FindAccountReportColumnId finds record id by querying it with criteria.
func (c *Client) FindAccountReportColumnId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportColumnModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
