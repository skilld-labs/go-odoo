package odoo

// AccountSecureEntriesWizard represents account.secure.entries.wizard model.
type AccountSecureEntriesWizard struct {
	ChainsToHashWithGaps             interface{} `xmlrpc:"chains_to_hash_with_gaps,omitempty"`
	CompanyId                        *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryCode                      *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                       *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                        *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                      *String     `xmlrpc:"display_name,omitempty"`
	HashDate                         *Time       `xmlrpc:"hash_date,omitempty"`
	Id                               *Int        `xmlrpc:"id,omitempty"`
	MaxHashDate                      *Time       `xmlrpc:"max_hash_date,omitempty"`
	MoveToHashIds                    *Relation   `xmlrpc:"move_to_hash_ids,omitempty"`
	NotHashableUnlockedMoveIds       *Relation   `xmlrpc:"not_hashable_unlocked_move_ids,omitempty"`
	UnreconciledBankStatementLineIds *Relation   `xmlrpc:"unreconciled_bank_statement_line_ids,omitempty"`
	Warnings                         interface{} `xmlrpc:"warnings,omitempty"`
	WriteDate                        *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                         *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountSecureEntriesWizards represents array of account.secure.entries.wizard model.
type AccountSecureEntriesWizards []AccountSecureEntriesWizard

// AccountSecureEntriesWizardModel is the odoo model name.
const AccountSecureEntriesWizardModel = "account.secure.entries.wizard"

// Many2One convert AccountSecureEntriesWizard to *Many2One.
func (asew *AccountSecureEntriesWizard) Many2One() *Many2One {
	return NewMany2One(asew.Id.Get(), "")
}

// CreateAccountSecureEntriesWizard creates a new account.secure.entries.wizard model and returns its id.
func (c *Client) CreateAccountSecureEntriesWizard(asew *AccountSecureEntriesWizard) (int64, error) {
	ids, err := c.CreateAccountSecureEntriesWizards([]*AccountSecureEntriesWizard{asew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountSecureEntriesWizard creates a new account.secure.entries.wizard model and returns its id.
func (c *Client) CreateAccountSecureEntriesWizards(asews []*AccountSecureEntriesWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range asews {
		vv = append(vv, v)
	}
	return c.Create(AccountSecureEntriesWizardModel, vv, nil)
}

// UpdateAccountSecureEntriesWizard updates an existing account.secure.entries.wizard record.
func (c *Client) UpdateAccountSecureEntriesWizard(asew *AccountSecureEntriesWizard) error {
	return c.UpdateAccountSecureEntriesWizards([]int64{asew.Id.Get()}, asew)
}

// UpdateAccountSecureEntriesWizards updates existing account.secure.entries.wizard records.
// All records (represented by ids) will be updated by asew values.
func (c *Client) UpdateAccountSecureEntriesWizards(ids []int64, asew *AccountSecureEntriesWizard) error {
	return c.Update(AccountSecureEntriesWizardModel, ids, asew, nil)
}

// DeleteAccountSecureEntriesWizard deletes an existing account.secure.entries.wizard record.
func (c *Client) DeleteAccountSecureEntriesWizard(id int64) error {
	return c.DeleteAccountSecureEntriesWizards([]int64{id})
}

// DeleteAccountSecureEntriesWizards deletes existing account.secure.entries.wizard records.
func (c *Client) DeleteAccountSecureEntriesWizards(ids []int64) error {
	return c.Delete(AccountSecureEntriesWizardModel, ids)
}

// GetAccountSecureEntriesWizard gets account.secure.entries.wizard existing record.
func (c *Client) GetAccountSecureEntriesWizard(id int64) (*AccountSecureEntriesWizard, error) {
	asews, err := c.GetAccountSecureEntriesWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*asews)[0]), nil
}

// GetAccountSecureEntriesWizards gets account.secure.entries.wizard existing records.
func (c *Client) GetAccountSecureEntriesWizards(ids []int64) (*AccountSecureEntriesWizards, error) {
	asews := &AccountSecureEntriesWizards{}
	if err := c.Read(AccountSecureEntriesWizardModel, ids, nil, asews); err != nil {
		return nil, err
	}
	return asews, nil
}

// FindAccountSecureEntriesWizard finds account.secure.entries.wizard record by querying it with criteria.
func (c *Client) FindAccountSecureEntriesWizard(criteria *Criteria) (*AccountSecureEntriesWizard, error) {
	asews := &AccountSecureEntriesWizards{}
	if err := c.SearchRead(AccountSecureEntriesWizardModel, criteria, NewOptions().Limit(1), asews); err != nil {
		return nil, err
	}
	return &((*asews)[0]), nil
}

// FindAccountSecureEntriesWizards finds account.secure.entries.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSecureEntriesWizards(criteria *Criteria, options *Options) (*AccountSecureEntriesWizards, error) {
	asews := &AccountSecureEntriesWizards{}
	if err := c.SearchRead(AccountSecureEntriesWizardModel, criteria, options, asews); err != nil {
		return nil, err
	}
	return asews, nil
}

// FindAccountSecureEntriesWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSecureEntriesWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountSecureEntriesWizardModel, criteria, options)
}

// FindAccountSecureEntriesWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountSecureEntriesWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSecureEntriesWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
