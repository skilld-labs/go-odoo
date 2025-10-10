package odoo

// AccountReportExpression represents account.report.expression model.
type AccountReportExpression struct {
	Auditable       *Bool      `xmlrpc:"auditable,omitempty"`
	BlankIfZero     *Bool      `xmlrpc:"blank_if_zero,omitempty"`
	CarryoverTarget *String    `xmlrpc:"carryover_target,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateScope       *Selection `xmlrpc:"date_scope,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Engine          *Selection `xmlrpc:"engine,omitempty"`
	FigureType      *Selection `xmlrpc:"figure_type,omitempty"`
	Formula         *String    `xmlrpc:"formula,omitempty"`
	GreenOnPositive *Bool      `xmlrpc:"green_on_positive,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Label           *String    `xmlrpc:"label,omitempty"`
	ReportLineId    *Many2One  `xmlrpc:"report_line_id,omitempty"`
	ReportLineName  *String    `xmlrpc:"report_line_name,omitempty"`
	Subformula      *String    `xmlrpc:"subformula,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReportExpressions represents array of account.report.expression model.
type AccountReportExpressions []AccountReportExpression

// AccountReportExpressionModel is the odoo model name.
const AccountReportExpressionModel = "account.report.expression"

// Many2One convert AccountReportExpression to *Many2One.
func (are *AccountReportExpression) Many2One() *Many2One {
	return NewMany2One(are.Id.Get(), "")
}

// CreateAccountReportExpression creates a new account.report.expression model and returns its id.
func (c *Client) CreateAccountReportExpression(are *AccountReportExpression) (int64, error) {
	ids, err := c.CreateAccountReportExpressions([]*AccountReportExpression{are})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportExpression creates a new account.report.expression model and returns its id.
func (c *Client) CreateAccountReportExpressions(ares []*AccountReportExpression) ([]int64, error) {
	var vv []interface{}
	for _, v := range ares {
		vv = append(vv, v)
	}
	return c.Create(AccountReportExpressionModel, vv, nil)
}

// UpdateAccountReportExpression updates an existing account.report.expression record.
func (c *Client) UpdateAccountReportExpression(are *AccountReportExpression) error {
	return c.UpdateAccountReportExpressions([]int64{are.Id.Get()}, are)
}

// UpdateAccountReportExpressions updates existing account.report.expression records.
// All records (represented by ids) will be updated by are values.
func (c *Client) UpdateAccountReportExpressions(ids []int64, are *AccountReportExpression) error {
	return c.Update(AccountReportExpressionModel, ids, are, nil)
}

// DeleteAccountReportExpression deletes an existing account.report.expression record.
func (c *Client) DeleteAccountReportExpression(id int64) error {
	return c.DeleteAccountReportExpressions([]int64{id})
}

// DeleteAccountReportExpressions deletes existing account.report.expression records.
func (c *Client) DeleteAccountReportExpressions(ids []int64) error {
	return c.Delete(AccountReportExpressionModel, ids)
}

// GetAccountReportExpression gets account.report.expression existing record.
func (c *Client) GetAccountReportExpression(id int64) (*AccountReportExpression, error) {
	ares, err := c.GetAccountReportExpressions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ares)[0]), nil
}

// GetAccountReportExpressions gets account.report.expression existing records.
func (c *Client) GetAccountReportExpressions(ids []int64) (*AccountReportExpressions, error) {
	ares := &AccountReportExpressions{}
	if err := c.Read(AccountReportExpressionModel, ids, nil, ares); err != nil {
		return nil, err
	}
	return ares, nil
}

// FindAccountReportExpression finds account.report.expression record by querying it with criteria.
func (c *Client) FindAccountReportExpression(criteria *Criteria) (*AccountReportExpression, error) {
	ares := &AccountReportExpressions{}
	if err := c.SearchRead(AccountReportExpressionModel, criteria, NewOptions().Limit(1), ares); err != nil {
		return nil, err
	}
	return &((*ares)[0]), nil
}

// FindAccountReportExpressions finds account.report.expression records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExpressions(criteria *Criteria, options *Options) (*AccountReportExpressions, error) {
	ares := &AccountReportExpressions{}
	if err := c.SearchRead(AccountReportExpressionModel, criteria, options, ares); err != nil {
		return nil, err
	}
	return ares, nil
}

// FindAccountReportExpressionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExpressionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportExpressionModel, criteria, options)
}

// FindAccountReportExpressionId finds record id by querying it with criteria.
func (c *Client) FindAccountReportExpressionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportExpressionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
