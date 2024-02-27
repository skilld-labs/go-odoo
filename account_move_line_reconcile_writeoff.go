package odoo

// AccountMoveLineReconcileWriteoff represents account.move.line.reconcile.writeoff model.
type AccountMoveLineReconcileWriteoff struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	AnalyticId    *Many2One `xmlrpc:"analytic_id,omitempty"`
	Comment       *String   `xmlrpc:"comment,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DateP         *Time     `xmlrpc:"date_p,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	JournalId     *Many2One `xmlrpc:"journal_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
	WriteoffAccId *Many2One `xmlrpc:"writeoff_acc_id,omitempty"`
}

// AccountMoveLineReconcileWriteoffs represents array of account.move.line.reconcile.writeoff model.
type AccountMoveLineReconcileWriteoffs []AccountMoveLineReconcileWriteoff

// AccountMoveLineReconcileWriteoffModel is the odoo model name.
const AccountMoveLineReconcileWriteoffModel = "account.move.line.reconcile.writeoff"

// Many2One convert AccountMoveLineReconcileWriteoff to *Many2One.
func (amlrw *AccountMoveLineReconcileWriteoff) Many2One() *Many2One {
	return NewMany2One(amlrw.Id.Get(), "")
}

// CreateAccountMoveLineReconcileWriteoff creates a new account.move.line.reconcile.writeoff model and returns its id.
func (c *Client) CreateAccountMoveLineReconcileWriteoff(amlrw *AccountMoveLineReconcileWriteoff) (int64, error) {
	ids, err := c.CreateAccountMoveLineReconcileWriteoffs([]*AccountMoveLineReconcileWriteoff{amlrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveLineReconcileWriteoffs creates a new account.move.line.reconcile.writeoff model and returns its id.
func (c *Client) CreateAccountMoveLineReconcileWriteoffs(amlrws []*AccountMoveLineReconcileWriteoff) ([]int64, error) {
	var vv []interface{}
	for _, v := range amlrws {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveLineReconcileWriteoffModel, vv, nil)
}

// UpdateAccountMoveLineReconcileWriteoff updates an existing account.move.line.reconcile.writeoff record.
func (c *Client) UpdateAccountMoveLineReconcileWriteoff(amlrw *AccountMoveLineReconcileWriteoff) error {
	return c.UpdateAccountMoveLineReconcileWriteoffs([]int64{amlrw.Id.Get()}, amlrw)
}

// UpdateAccountMoveLineReconcileWriteoffs updates existing account.move.line.reconcile.writeoff records.
// All records (represented by ids) will be updated by amlrw values.
func (c *Client) UpdateAccountMoveLineReconcileWriteoffs(ids []int64, amlrw *AccountMoveLineReconcileWriteoff) error {
	return c.Update(AccountMoveLineReconcileWriteoffModel, ids, amlrw, nil)
}

// DeleteAccountMoveLineReconcileWriteoff deletes an existing account.move.line.reconcile.writeoff record.
func (c *Client) DeleteAccountMoveLineReconcileWriteoff(id int64) error {
	return c.DeleteAccountMoveLineReconcileWriteoffs([]int64{id})
}

// DeleteAccountMoveLineReconcileWriteoffs deletes existing account.move.line.reconcile.writeoff records.
func (c *Client) DeleteAccountMoveLineReconcileWriteoffs(ids []int64) error {
	return c.Delete(AccountMoveLineReconcileWriteoffModel, ids)
}

// GetAccountMoveLineReconcileWriteoff gets account.move.line.reconcile.writeoff existing record.
func (c *Client) GetAccountMoveLineReconcileWriteoff(id int64) (*AccountMoveLineReconcileWriteoff, error) {
	amlrws, err := c.GetAccountMoveLineReconcileWriteoffs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amlrws)[0]), nil
}

// GetAccountMoveLineReconcileWriteoffs gets account.move.line.reconcile.writeoff existing records.
func (c *Client) GetAccountMoveLineReconcileWriteoffs(ids []int64) (*AccountMoveLineReconcileWriteoffs, error) {
	amlrws := &AccountMoveLineReconcileWriteoffs{}
	if err := c.Read(AccountMoveLineReconcileWriteoffModel, ids, nil, amlrws); err != nil {
		return nil, err
	}
	return amlrws, nil
}

// FindAccountMoveLineReconcileWriteoff finds account.move.line.reconcile.writeoff record by querying it with criteria.
func (c *Client) FindAccountMoveLineReconcileWriteoff(criteria *Criteria) (*AccountMoveLineReconcileWriteoff, error) {
	amlrws := &AccountMoveLineReconcileWriteoffs{}
	if err := c.SearchRead(AccountMoveLineReconcileWriteoffModel, criteria, NewOptions().Limit(1), amlrws); err != nil {
		return nil, err
	}
	return &((*amlrws)[0]), nil
}

// FindAccountMoveLineReconcileWriteoffs finds account.move.line.reconcile.writeoff records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineReconcileWriteoffs(criteria *Criteria, options *Options) (*AccountMoveLineReconcileWriteoffs, error) {
	amlrws := &AccountMoveLineReconcileWriteoffs{}
	if err := c.SearchRead(AccountMoveLineReconcileWriteoffModel, criteria, options, amlrws); err != nil {
		return nil, err
	}
	return amlrws, nil
}

// FindAccountMoveLineReconcileWriteoffIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineReconcileWriteoffIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveLineReconcileWriteoffModel, criteria, options)
}

// FindAccountMoveLineReconcileWriteoffId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineReconcileWriteoffId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineReconcileWriteoffModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
