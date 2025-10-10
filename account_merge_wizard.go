package odoo

// AccountMergeWizard represents account.merge.wizard model.
type AccountMergeWizard struct {
	AccountIds         *Relation `xmlrpc:"account_ids,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisableMergeButton *Bool     `xmlrpc:"disable_merge_button,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	IsGroupByName      *Bool     `xmlrpc:"is_group_by_name,omitempty"`
	WizardLineIds      *Relation `xmlrpc:"wizard_line_ids,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountMergeWizards represents array of account.merge.wizard model.
type AccountMergeWizards []AccountMergeWizard

// AccountMergeWizardModel is the odoo model name.
const AccountMergeWizardModel = "account.merge.wizard"

// Many2One convert AccountMergeWizard to *Many2One.
func (amw *AccountMergeWizard) Many2One() *Many2One {
	return NewMany2One(amw.Id.Get(), "")
}

// CreateAccountMergeWizard creates a new account.merge.wizard model and returns its id.
func (c *Client) CreateAccountMergeWizard(amw *AccountMergeWizard) (int64, error) {
	ids, err := c.CreateAccountMergeWizards([]*AccountMergeWizard{amw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMergeWizard creates a new account.merge.wizard model and returns its id.
func (c *Client) CreateAccountMergeWizards(amws []*AccountMergeWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amws {
		vv = append(vv, v)
	}
	return c.Create(AccountMergeWizardModel, vv, nil)
}

// UpdateAccountMergeWizard updates an existing account.merge.wizard record.
func (c *Client) UpdateAccountMergeWizard(amw *AccountMergeWizard) error {
	return c.UpdateAccountMergeWizards([]int64{amw.Id.Get()}, amw)
}

// UpdateAccountMergeWizards updates existing account.merge.wizard records.
// All records (represented by ids) will be updated by amw values.
func (c *Client) UpdateAccountMergeWizards(ids []int64, amw *AccountMergeWizard) error {
	return c.Update(AccountMergeWizardModel, ids, amw, nil)
}

// DeleteAccountMergeWizard deletes an existing account.merge.wizard record.
func (c *Client) DeleteAccountMergeWizard(id int64) error {
	return c.DeleteAccountMergeWizards([]int64{id})
}

// DeleteAccountMergeWizards deletes existing account.merge.wizard records.
func (c *Client) DeleteAccountMergeWizards(ids []int64) error {
	return c.Delete(AccountMergeWizardModel, ids)
}

// GetAccountMergeWizard gets account.merge.wizard existing record.
func (c *Client) GetAccountMergeWizard(id int64) (*AccountMergeWizard, error) {
	amws, err := c.GetAccountMergeWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amws)[0]), nil
}

// GetAccountMergeWizards gets account.merge.wizard existing records.
func (c *Client) GetAccountMergeWizards(ids []int64) (*AccountMergeWizards, error) {
	amws := &AccountMergeWizards{}
	if err := c.Read(AccountMergeWizardModel, ids, nil, amws); err != nil {
		return nil, err
	}
	return amws, nil
}

// FindAccountMergeWizard finds account.merge.wizard record by querying it with criteria.
func (c *Client) FindAccountMergeWizard(criteria *Criteria) (*AccountMergeWizard, error) {
	amws := &AccountMergeWizards{}
	if err := c.SearchRead(AccountMergeWizardModel, criteria, NewOptions().Limit(1), amws); err != nil {
		return nil, err
	}
	return &((*amws)[0]), nil
}

// FindAccountMergeWizards finds account.merge.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMergeWizards(criteria *Criteria, options *Options) (*AccountMergeWizards, error) {
	amws := &AccountMergeWizards{}
	if err := c.SearchRead(AccountMergeWizardModel, criteria, options, amws); err != nil {
		return nil, err
	}
	return amws, nil
}

// FindAccountMergeWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMergeWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMergeWizardModel, criteria, options)
}

// FindAccountMergeWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMergeWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMergeWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
