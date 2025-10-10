package odoo

// AccountReconcileModelLine represents account.reconcile.model.line model.
type AccountReconcileModelLine struct {
	AccountId                      *Many2One   `xmlrpc:"account_id,omitempty"`
	AllowPaymentTolerance          *Bool       `xmlrpc:"allow_payment_tolerance,omitempty"`
	Amount                         *Float      `xmlrpc:"amount,omitempty"`
	AmountString                   *String     `xmlrpc:"amount_string,omitempty"`
	AmountType                     *Selection  `xmlrpc:"amount_type,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	DistributionAnalyticAccountIds *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	ForceTaxIncluded               *Bool       `xmlrpc:"force_tax_included,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	JournalId                      *Many2One   `xmlrpc:"journal_id,omitempty"`
	Label                          *String     `xmlrpc:"label,omitempty"`
	ModelId                        *Many2One   `xmlrpc:"model_id,omitempty"`
	PaymentToleranceParam          *Float      `xmlrpc:"payment_tolerance_param,omitempty"`
	RuleType                       *Selection  `xmlrpc:"rule_type,omitempty"`
	Sequence                       *Int        `xmlrpc:"sequence,omitempty"`
	ShowForceTaxIncluded           *Bool       `xmlrpc:"show_force_tax_included,omitempty"`
	TaxIds                         *Relation   `xmlrpc:"tax_ids,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModelLines represents array of account.reconcile.model.line model.
type AccountReconcileModelLines []AccountReconcileModelLine

// AccountReconcileModelLineModel is the odoo model name.
const AccountReconcileModelLineModel = "account.reconcile.model.line"

// Many2One convert AccountReconcileModelLine to *Many2One.
func (arml *AccountReconcileModelLine) Many2One() *Many2One {
	return NewMany2One(arml.Id.Get(), "")
}

// CreateAccountReconcileModelLine creates a new account.reconcile.model.line model and returns its id.
func (c *Client) CreateAccountReconcileModelLine(arml *AccountReconcileModelLine) (int64, error) {
	ids, err := c.CreateAccountReconcileModelLines([]*AccountReconcileModelLine{arml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModelLine creates a new account.reconcile.model.line model and returns its id.
func (c *Client) CreateAccountReconcileModelLines(armls []*AccountReconcileModelLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range armls {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelLineModel, vv, nil)
}

// UpdateAccountReconcileModelLine updates an existing account.reconcile.model.line record.
func (c *Client) UpdateAccountReconcileModelLine(arml *AccountReconcileModelLine) error {
	return c.UpdateAccountReconcileModelLines([]int64{arml.Id.Get()}, arml)
}

// UpdateAccountReconcileModelLines updates existing account.reconcile.model.line records.
// All records (represented by ids) will be updated by arml values.
func (c *Client) UpdateAccountReconcileModelLines(ids []int64, arml *AccountReconcileModelLine) error {
	return c.Update(AccountReconcileModelLineModel, ids, arml, nil)
}

// DeleteAccountReconcileModelLine deletes an existing account.reconcile.model.line record.
func (c *Client) DeleteAccountReconcileModelLine(id int64) error {
	return c.DeleteAccountReconcileModelLines([]int64{id})
}

// DeleteAccountReconcileModelLines deletes existing account.reconcile.model.line records.
func (c *Client) DeleteAccountReconcileModelLines(ids []int64) error {
	return c.Delete(AccountReconcileModelLineModel, ids)
}

// GetAccountReconcileModelLine gets account.reconcile.model.line existing record.
func (c *Client) GetAccountReconcileModelLine(id int64) (*AccountReconcileModelLine, error) {
	armls, err := c.GetAccountReconcileModelLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*armls)[0]), nil
}

// GetAccountReconcileModelLines gets account.reconcile.model.line existing records.
func (c *Client) GetAccountReconcileModelLines(ids []int64) (*AccountReconcileModelLines, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.Read(AccountReconcileModelLineModel, ids, nil, armls); err != nil {
		return nil, err
	}
	return armls, nil
}

// FindAccountReconcileModelLine finds account.reconcile.model.line record by querying it with criteria.
func (c *Client) FindAccountReconcileModelLine(criteria *Criteria) (*AccountReconcileModelLine, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.SearchRead(AccountReconcileModelLineModel, criteria, NewOptions().Limit(1), armls); err != nil {
		return nil, err
	}
	return &((*armls)[0]), nil
}

// FindAccountReconcileModelLines finds account.reconcile.model.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLines(criteria *Criteria, options *Options) (*AccountReconcileModelLines, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.SearchRead(AccountReconcileModelLineModel, criteria, options, armls); err != nil {
		return nil, err
	}
	return armls, nil
}

// FindAccountReconcileModelLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelLineModel, criteria, options)
}

// FindAccountReconcileModelLineId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
