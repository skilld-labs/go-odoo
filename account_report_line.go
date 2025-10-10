package odoo

// AccountReportLine represents account.report.line model.
type AccountReportLine struct {
	AccountCodesFormula *String    `xmlrpc:"account_codes_formula,omitempty"`
	ActionId            *Many2One  `xmlrpc:"action_id,omitempty"`
	AggregationFormula  *String    `xmlrpc:"aggregation_formula,omitempty"`
	ChildrenIds         *Relation  `xmlrpc:"children_ids,omitempty"`
	Code                *String    `xmlrpc:"code,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	DomainFormula       *String    `xmlrpc:"domain_formula,omitempty"`
	ExpressionIds       *Relation  `xmlrpc:"expression_ids,omitempty"`
	ExternalFormula     *String    `xmlrpc:"external_formula,omitempty"`
	Foldable            *Bool      `xmlrpc:"foldable,omitempty"`
	Groupby             *String    `xmlrpc:"groupby,omitempty"`
	HideIfZero          *Bool      `xmlrpc:"hide_if_zero,omitempty"`
	HierarchyLevel      *Int       `xmlrpc:"hierarchy_level,omitempty"`
	HorizontalSplitSide *Selection `xmlrpc:"horizontal_split_side,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	ParentId            *Many2One  `xmlrpc:"parent_id,omitempty"`
	PrintOnNewPage      *Bool      `xmlrpc:"print_on_new_page,omitempty"`
	ReportId            *Many2One  `xmlrpc:"report_id,omitempty"`
	Sequence            *Int       `xmlrpc:"sequence,omitempty"`
	TaxTagsFormula      *String    `xmlrpc:"tax_tags_formula,omitempty"`
	UserGroupby         *String    `xmlrpc:"user_groupby,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReportLines represents array of account.report.line model.
type AccountReportLines []AccountReportLine

// AccountReportLineModel is the odoo model name.
const AccountReportLineModel = "account.report.line"

// Many2One convert AccountReportLine to *Many2One.
func (arl *AccountReportLine) Many2One() *Many2One {
	return NewMany2One(arl.Id.Get(), "")
}

// CreateAccountReportLine creates a new account.report.line model and returns its id.
func (c *Client) CreateAccountReportLine(arl *AccountReportLine) (int64, error) {
	ids, err := c.CreateAccountReportLines([]*AccountReportLine{arl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportLine creates a new account.report.line model and returns its id.
func (c *Client) CreateAccountReportLines(arls []*AccountReportLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range arls {
		vv = append(vv, v)
	}
	return c.Create(AccountReportLineModel, vv, nil)
}

// UpdateAccountReportLine updates an existing account.report.line record.
func (c *Client) UpdateAccountReportLine(arl *AccountReportLine) error {
	return c.UpdateAccountReportLines([]int64{arl.Id.Get()}, arl)
}

// UpdateAccountReportLines updates existing account.report.line records.
// All records (represented by ids) will be updated by arl values.
func (c *Client) UpdateAccountReportLines(ids []int64, arl *AccountReportLine) error {
	return c.Update(AccountReportLineModel, ids, arl, nil)
}

// DeleteAccountReportLine deletes an existing account.report.line record.
func (c *Client) DeleteAccountReportLine(id int64) error {
	return c.DeleteAccountReportLines([]int64{id})
}

// DeleteAccountReportLines deletes existing account.report.line records.
func (c *Client) DeleteAccountReportLines(ids []int64) error {
	return c.Delete(AccountReportLineModel, ids)
}

// GetAccountReportLine gets account.report.line existing record.
func (c *Client) GetAccountReportLine(id int64) (*AccountReportLine, error) {
	arls, err := c.GetAccountReportLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arls)[0]), nil
}

// GetAccountReportLines gets account.report.line existing records.
func (c *Client) GetAccountReportLines(ids []int64) (*AccountReportLines, error) {
	arls := &AccountReportLines{}
	if err := c.Read(AccountReportLineModel, ids, nil, arls); err != nil {
		return nil, err
	}
	return arls, nil
}

// FindAccountReportLine finds account.report.line record by querying it with criteria.
func (c *Client) FindAccountReportLine(criteria *Criteria) (*AccountReportLine, error) {
	arls := &AccountReportLines{}
	if err := c.SearchRead(AccountReportLineModel, criteria, NewOptions().Limit(1), arls); err != nil {
		return nil, err
	}
	return &((*arls)[0]), nil
}

// FindAccountReportLines finds account.report.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportLines(criteria *Criteria, options *Options) (*AccountReportLines, error) {
	arls := &AccountReportLines{}
	if err := c.SearchRead(AccountReportLineModel, criteria, options, arls); err != nil {
		return nil, err
	}
	return arls, nil
}

// FindAccountReportLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportLineModel, criteria, options)
}

// FindAccountReportLineId finds record id by querying it with criteria.
func (c *Client) FindAccountReportLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
