package odoo

// AccountAnalyticLineCalendarEmployee represents account.analytic.line.calendar.employee model.
type AccountAnalyticLineCalendarEmployee struct {
	Active      *Bool     `xmlrpc:"active,omitempty"`
	Checked     *Bool     `xmlrpc:"checked,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticLineCalendarEmployees represents array of account.analytic.line.calendar.employee model.
type AccountAnalyticLineCalendarEmployees []AccountAnalyticLineCalendarEmployee

// AccountAnalyticLineCalendarEmployeeModel is the odoo model name.
const AccountAnalyticLineCalendarEmployeeModel = "account.analytic.line.calendar.employee"

// Many2One convert AccountAnalyticLineCalendarEmployee to *Many2One.
func (aalce *AccountAnalyticLineCalendarEmployee) Many2One() *Many2One {
	return NewMany2One(aalce.Id.Get(), "")
}

// CreateAccountAnalyticLineCalendarEmployee creates a new account.analytic.line.calendar.employee model and returns its id.
func (c *Client) CreateAccountAnalyticLineCalendarEmployee(aalce *AccountAnalyticLineCalendarEmployee) (int64, error) {
	ids, err := c.CreateAccountAnalyticLineCalendarEmployees([]*AccountAnalyticLineCalendarEmployee{aalce})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticLineCalendarEmployee creates a new account.analytic.line.calendar.employee model and returns its id.
func (c *Client) CreateAccountAnalyticLineCalendarEmployees(aalces []*AccountAnalyticLineCalendarEmployee) ([]int64, error) {
	var vv []interface{}
	for _, v := range aalces {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticLineCalendarEmployeeModel, vv, nil)
}

// UpdateAccountAnalyticLineCalendarEmployee updates an existing account.analytic.line.calendar.employee record.
func (c *Client) UpdateAccountAnalyticLineCalendarEmployee(aalce *AccountAnalyticLineCalendarEmployee) error {
	return c.UpdateAccountAnalyticLineCalendarEmployees([]int64{aalce.Id.Get()}, aalce)
}

// UpdateAccountAnalyticLineCalendarEmployees updates existing account.analytic.line.calendar.employee records.
// All records (represented by ids) will be updated by aalce values.
func (c *Client) UpdateAccountAnalyticLineCalendarEmployees(ids []int64, aalce *AccountAnalyticLineCalendarEmployee) error {
	return c.Update(AccountAnalyticLineCalendarEmployeeModel, ids, aalce, nil)
}

// DeleteAccountAnalyticLineCalendarEmployee deletes an existing account.analytic.line.calendar.employee record.
func (c *Client) DeleteAccountAnalyticLineCalendarEmployee(id int64) error {
	return c.DeleteAccountAnalyticLineCalendarEmployees([]int64{id})
}

// DeleteAccountAnalyticLineCalendarEmployees deletes existing account.analytic.line.calendar.employee records.
func (c *Client) DeleteAccountAnalyticLineCalendarEmployees(ids []int64) error {
	return c.Delete(AccountAnalyticLineCalendarEmployeeModel, ids)
}

// GetAccountAnalyticLineCalendarEmployee gets account.analytic.line.calendar.employee existing record.
func (c *Client) GetAccountAnalyticLineCalendarEmployee(id int64) (*AccountAnalyticLineCalendarEmployee, error) {
	aalces, err := c.GetAccountAnalyticLineCalendarEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aalces)[0]), nil
}

// GetAccountAnalyticLineCalendarEmployees gets account.analytic.line.calendar.employee existing records.
func (c *Client) GetAccountAnalyticLineCalendarEmployees(ids []int64) (*AccountAnalyticLineCalendarEmployees, error) {
	aalces := &AccountAnalyticLineCalendarEmployees{}
	if err := c.Read(AccountAnalyticLineCalendarEmployeeModel, ids, nil, aalces); err != nil {
		return nil, err
	}
	return aalces, nil
}

// FindAccountAnalyticLineCalendarEmployee finds account.analytic.line.calendar.employee record by querying it with criteria.
func (c *Client) FindAccountAnalyticLineCalendarEmployee(criteria *Criteria) (*AccountAnalyticLineCalendarEmployee, error) {
	aalces := &AccountAnalyticLineCalendarEmployees{}
	if err := c.SearchRead(AccountAnalyticLineCalendarEmployeeModel, criteria, NewOptions().Limit(1), aalces); err != nil {
		return nil, err
	}
	return &((*aalces)[0]), nil
}

// FindAccountAnalyticLineCalendarEmployees finds account.analytic.line.calendar.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLineCalendarEmployees(criteria *Criteria, options *Options) (*AccountAnalyticLineCalendarEmployees, error) {
	aalces := &AccountAnalyticLineCalendarEmployees{}
	if err := c.SearchRead(AccountAnalyticLineCalendarEmployeeModel, criteria, options, aalces); err != nil {
		return nil, err
	}
	return aalces, nil
}

// FindAccountAnalyticLineCalendarEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLineCalendarEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticLineCalendarEmployeeModel, criteria, options)
}

// FindAccountAnalyticLineCalendarEmployeeId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticLineCalendarEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticLineCalendarEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
