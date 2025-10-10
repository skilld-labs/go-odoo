package odoo

// AccountAutopostBillsWizard represents account.autopost.bills.wizard model.
type AccountAutopostBillsWizard struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	NbUnmodifiedBills *Int      `xmlrpc:"nb_unmodified_bills,omitempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omitempty"`
	PartnerName       *String   `xmlrpc:"partner_name,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAutopostBillsWizards represents array of account.autopost.bills.wizard model.
type AccountAutopostBillsWizards []AccountAutopostBillsWizard

// AccountAutopostBillsWizardModel is the odoo model name.
const AccountAutopostBillsWizardModel = "account.autopost.bills.wizard"

// Many2One convert AccountAutopostBillsWizard to *Many2One.
func (aabw *AccountAutopostBillsWizard) Many2One() *Many2One {
	return NewMany2One(aabw.Id.Get(), "")
}

// CreateAccountAutopostBillsWizard creates a new account.autopost.bills.wizard model and returns its id.
func (c *Client) CreateAccountAutopostBillsWizard(aabw *AccountAutopostBillsWizard) (int64, error) {
	ids, err := c.CreateAccountAutopostBillsWizards([]*AccountAutopostBillsWizard{aabw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAutopostBillsWizard creates a new account.autopost.bills.wizard model and returns its id.
func (c *Client) CreateAccountAutopostBillsWizards(aabws []*AccountAutopostBillsWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aabws {
		vv = append(vv, v)
	}
	return c.Create(AccountAutopostBillsWizardModel, vv, nil)
}

// UpdateAccountAutopostBillsWizard updates an existing account.autopost.bills.wizard record.
func (c *Client) UpdateAccountAutopostBillsWizard(aabw *AccountAutopostBillsWizard) error {
	return c.UpdateAccountAutopostBillsWizards([]int64{aabw.Id.Get()}, aabw)
}

// UpdateAccountAutopostBillsWizards updates existing account.autopost.bills.wizard records.
// All records (represented by ids) will be updated by aabw values.
func (c *Client) UpdateAccountAutopostBillsWizards(ids []int64, aabw *AccountAutopostBillsWizard) error {
	return c.Update(AccountAutopostBillsWizardModel, ids, aabw, nil)
}

// DeleteAccountAutopostBillsWizard deletes an existing account.autopost.bills.wizard record.
func (c *Client) DeleteAccountAutopostBillsWizard(id int64) error {
	return c.DeleteAccountAutopostBillsWizards([]int64{id})
}

// DeleteAccountAutopostBillsWizards deletes existing account.autopost.bills.wizard records.
func (c *Client) DeleteAccountAutopostBillsWizards(ids []int64) error {
	return c.Delete(AccountAutopostBillsWizardModel, ids)
}

// GetAccountAutopostBillsWizard gets account.autopost.bills.wizard existing record.
func (c *Client) GetAccountAutopostBillsWizard(id int64) (*AccountAutopostBillsWizard, error) {
	aabws, err := c.GetAccountAutopostBillsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aabws)[0]), nil
}

// GetAccountAutopostBillsWizards gets account.autopost.bills.wizard existing records.
func (c *Client) GetAccountAutopostBillsWizards(ids []int64) (*AccountAutopostBillsWizards, error) {
	aabws := &AccountAutopostBillsWizards{}
	if err := c.Read(AccountAutopostBillsWizardModel, ids, nil, aabws); err != nil {
		return nil, err
	}
	return aabws, nil
}

// FindAccountAutopostBillsWizard finds account.autopost.bills.wizard record by querying it with criteria.
func (c *Client) FindAccountAutopostBillsWizard(criteria *Criteria) (*AccountAutopostBillsWizard, error) {
	aabws := &AccountAutopostBillsWizards{}
	if err := c.SearchRead(AccountAutopostBillsWizardModel, criteria, NewOptions().Limit(1), aabws); err != nil {
		return nil, err
	}
	return &((*aabws)[0]), nil
}

// FindAccountAutopostBillsWizards finds account.autopost.bills.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutopostBillsWizards(criteria *Criteria, options *Options) (*AccountAutopostBillsWizards, error) {
	aabws := &AccountAutopostBillsWizards{}
	if err := c.SearchRead(AccountAutopostBillsWizardModel, criteria, options, aabws); err != nil {
		return nil, err
	}
	return aabws, nil
}

// FindAccountAutopostBillsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutopostBillsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAutopostBillsWizardModel, criteria, options)
}

// FindAccountAutopostBillsWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAutopostBillsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAutopostBillsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
