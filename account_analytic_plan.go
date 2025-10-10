package odoo

// AccountAnalyticPlan represents account.analytic.plan model.
type AccountAnalyticPlan struct {
	AccountCount         *Int       `xmlrpc:"account_count,omitempty"`
	AccountIds           *Relation  `xmlrpc:"account_ids,omitempty"`
	AllAccountCount      *Int       `xmlrpc:"all_account_count,omitempty"`
	ApplicabilityIds     *Relation  `xmlrpc:"applicability_ids,omitempty"`
	ChildrenCount        *Int       `xmlrpc:"children_count,omitempty"`
	ChildrenIds          *Relation  `xmlrpc:"children_ids,omitempty"`
	Color                *Int       `xmlrpc:"color,omitempty"`
	CompleteName         *String    `xmlrpc:"complete_name,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultApplicability *Selection `xmlrpc:"default_applicability,omitempty"`
	Description          *String    `xmlrpc:"description,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentPath           *String    `xmlrpc:"parent_path,omitempty"`
	RootId               *Many2One  `xmlrpc:"root_id,omitempty"`
	Sequence             *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticPlans represents array of account.analytic.plan model.
type AccountAnalyticPlans []AccountAnalyticPlan

// AccountAnalyticPlanModel is the odoo model name.
const AccountAnalyticPlanModel = "account.analytic.plan"

// Many2One convert AccountAnalyticPlan to *Many2One.
func (aap *AccountAnalyticPlan) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAnalyticPlan creates a new account.analytic.plan model and returns its id.
func (c *Client) CreateAccountAnalyticPlan(aap *AccountAnalyticPlan) (int64, error) {
	ids, err := c.CreateAccountAnalyticPlans([]*AccountAnalyticPlan{aap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticPlan creates a new account.analytic.plan model and returns its id.
func (c *Client) CreateAccountAnalyticPlans(aaps []*AccountAnalyticPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaps {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticPlanModel, vv, nil)
}

// UpdateAccountAnalyticPlan updates an existing account.analytic.plan record.
func (c *Client) UpdateAccountAnalyticPlan(aap *AccountAnalyticPlan) error {
	return c.UpdateAccountAnalyticPlans([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAnalyticPlans updates existing account.analytic.plan records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAnalyticPlans(ids []int64, aap *AccountAnalyticPlan) error {
	return c.Update(AccountAnalyticPlanModel, ids, aap, nil)
}

// DeleteAccountAnalyticPlan deletes an existing account.analytic.plan record.
func (c *Client) DeleteAccountAnalyticPlan(id int64) error {
	return c.DeleteAccountAnalyticPlans([]int64{id})
}

// DeleteAccountAnalyticPlans deletes existing account.analytic.plan records.
func (c *Client) DeleteAccountAnalyticPlans(ids []int64) error {
	return c.Delete(AccountAnalyticPlanModel, ids)
}

// GetAccountAnalyticPlan gets account.analytic.plan existing record.
func (c *Client) GetAccountAnalyticPlan(id int64) (*AccountAnalyticPlan, error) {
	aaps, err := c.GetAccountAnalyticPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaps)[0]), nil
}

// GetAccountAnalyticPlans gets account.analytic.plan existing records.
func (c *Client) GetAccountAnalyticPlans(ids []int64) (*AccountAnalyticPlans, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.Read(AccountAnalyticPlanModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAnalyticPlan finds account.analytic.plan record by querying it with criteria.
func (c *Client) FindAccountAnalyticPlan(criteria *Criteria) (*AccountAnalyticPlan, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.SearchRead(AccountAnalyticPlanModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	return &((*aaps)[0]), nil
}

// FindAccountAnalyticPlans finds account.analytic.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticPlans(criteria *Criteria, options *Options) (*AccountAnalyticPlans, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.SearchRead(AccountAnalyticPlanModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAnalyticPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticPlanModel, criteria, options)
}

// FindAccountAnalyticPlanId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
