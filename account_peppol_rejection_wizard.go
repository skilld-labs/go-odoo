package odoo

// AccountPeppolRejectionWizard represents account.peppol.rejection.wizard model.
type AccountPeppolRejectionWizard struct {
	ActionIds   *Relation `xmlrpc:"action_ids,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MoveIds     *Relation `xmlrpc:"move_ids,omitempty"`
	ReasonIds   *Relation `xmlrpc:"reason_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountPeppolRejectionWizards represents array of account.peppol.rejection.wizard model.
type AccountPeppolRejectionWizards []AccountPeppolRejectionWizard

// AccountPeppolRejectionWizardModel is the odoo model name.
const AccountPeppolRejectionWizardModel = "account.peppol.rejection.wizard"

// Many2One convert AccountPeppolRejectionWizard to *Many2One.
func (aprw *AccountPeppolRejectionWizard) Many2One() *Many2One {
	return NewMany2One(aprw.Id.Get(), "")
}

// CreateAccountPeppolRejectionWizard creates a new account.peppol.rejection.wizard model and returns its id.
func (c *Client) CreateAccountPeppolRejectionWizard(aprw *AccountPeppolRejectionWizard) (int64, error) {
	ids, err := c.CreateAccountPeppolRejectionWizards([]*AccountPeppolRejectionWizard{aprw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPeppolRejectionWizard creates a new account.peppol.rejection.wizard model and returns its id.
func (c *Client) CreateAccountPeppolRejectionWizards(aprws []*AccountPeppolRejectionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aprws {
		vv = append(vv, v)
	}
	return c.Create(AccountPeppolRejectionWizardModel, vv, nil)
}

// UpdateAccountPeppolRejectionWizard updates an existing account.peppol.rejection.wizard record.
func (c *Client) UpdateAccountPeppolRejectionWizard(aprw *AccountPeppolRejectionWizard) error {
	return c.UpdateAccountPeppolRejectionWizards([]int64{aprw.Id.Get()}, aprw)
}

// UpdateAccountPeppolRejectionWizards updates existing account.peppol.rejection.wizard records.
// All records (represented by ids) will be updated by aprw values.
func (c *Client) UpdateAccountPeppolRejectionWizards(ids []int64, aprw *AccountPeppolRejectionWizard) error {
	return c.Update(AccountPeppolRejectionWizardModel, ids, aprw, nil)
}

// DeleteAccountPeppolRejectionWizard deletes an existing account.peppol.rejection.wizard record.
func (c *Client) DeleteAccountPeppolRejectionWizard(id int64) error {
	return c.DeleteAccountPeppolRejectionWizards([]int64{id})
}

// DeleteAccountPeppolRejectionWizards deletes existing account.peppol.rejection.wizard records.
func (c *Client) DeleteAccountPeppolRejectionWizards(ids []int64) error {
	return c.Delete(AccountPeppolRejectionWizardModel, ids)
}

// GetAccountPeppolRejectionWizard gets account.peppol.rejection.wizard existing record.
func (c *Client) GetAccountPeppolRejectionWizard(id int64) (*AccountPeppolRejectionWizard, error) {
	aprws, err := c.GetAccountPeppolRejectionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aprws)[0]), nil
}

// GetAccountPeppolRejectionWizards gets account.peppol.rejection.wizard existing records.
func (c *Client) GetAccountPeppolRejectionWizards(ids []int64) (*AccountPeppolRejectionWizards, error) {
	aprws := &AccountPeppolRejectionWizards{}
	if err := c.Read(AccountPeppolRejectionWizardModel, ids, nil, aprws); err != nil {
		return nil, err
	}
	return aprws, nil
}

// FindAccountPeppolRejectionWizard finds account.peppol.rejection.wizard record by querying it with criteria.
func (c *Client) FindAccountPeppolRejectionWizard(criteria *Criteria) (*AccountPeppolRejectionWizard, error) {
	aprws := &AccountPeppolRejectionWizards{}
	if err := c.SearchRead(AccountPeppolRejectionWizardModel, criteria, NewOptions().Limit(1), aprws); err != nil {
		return nil, err
	}
	return &((*aprws)[0]), nil
}

// FindAccountPeppolRejectionWizards finds account.peppol.rejection.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolRejectionWizards(criteria *Criteria, options *Options) (*AccountPeppolRejectionWizards, error) {
	aprws := &AccountPeppolRejectionWizards{}
	if err := c.SearchRead(AccountPeppolRejectionWizardModel, criteria, options, aprws); err != nil {
		return nil, err
	}
	return aprws, nil
}

// FindAccountPeppolRejectionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolRejectionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPeppolRejectionWizardModel, criteria, options)
}

// FindAccountPeppolRejectionWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountPeppolRejectionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPeppolRejectionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
