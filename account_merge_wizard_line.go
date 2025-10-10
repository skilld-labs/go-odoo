package odoo

// AccountMergeWizardLine represents account.merge.wizard.line model.
type AccountMergeWizardLine struct {
	AccountHasHashedEntries *Bool      `xmlrpc:"account_has_hashed_entries,omitempty"`
	AccountId               *Many2One  `xmlrpc:"account_id,omitempty"`
	CompanyIds              *Relation  `xmlrpc:"company_ids,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty"`
	DisplayType             *Selection `xmlrpc:"display_type,omitempty"`
	GroupingKey             *String    `xmlrpc:"grouping_key,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty"`
	Info                    *String    `xmlrpc:"info,omitempty"`
	IsSelected              *Bool      `xmlrpc:"is_selected,omitempty"`
	Sequence                *Int       `xmlrpc:"sequence,omitempty"`
	WizardId                *Many2One  `xmlrpc:"wizard_id,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountMergeWizardLines represents array of account.merge.wizard.line model.
type AccountMergeWizardLines []AccountMergeWizardLine

// AccountMergeWizardLineModel is the odoo model name.
const AccountMergeWizardLineModel = "account.merge.wizard.line"

// Many2One convert AccountMergeWizardLine to *Many2One.
func (amwl *AccountMergeWizardLine) Many2One() *Many2One {
	return NewMany2One(amwl.Id.Get(), "")
}

// CreateAccountMergeWizardLine creates a new account.merge.wizard.line model and returns its id.
func (c *Client) CreateAccountMergeWizardLine(amwl *AccountMergeWizardLine) (int64, error) {
	ids, err := c.CreateAccountMergeWizardLines([]*AccountMergeWizardLine{amwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMergeWizardLine creates a new account.merge.wizard.line model and returns its id.
func (c *Client) CreateAccountMergeWizardLines(amwls []*AccountMergeWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range amwls {
		vv = append(vv, v)
	}
	return c.Create(AccountMergeWizardLineModel, vv, nil)
}

// UpdateAccountMergeWizardLine updates an existing account.merge.wizard.line record.
func (c *Client) UpdateAccountMergeWizardLine(amwl *AccountMergeWizardLine) error {
	return c.UpdateAccountMergeWizardLines([]int64{amwl.Id.Get()}, amwl)
}

// UpdateAccountMergeWizardLines updates existing account.merge.wizard.line records.
// All records (represented by ids) will be updated by amwl values.
func (c *Client) UpdateAccountMergeWizardLines(ids []int64, amwl *AccountMergeWizardLine) error {
	return c.Update(AccountMergeWizardLineModel, ids, amwl, nil)
}

// DeleteAccountMergeWizardLine deletes an existing account.merge.wizard.line record.
func (c *Client) DeleteAccountMergeWizardLine(id int64) error {
	return c.DeleteAccountMergeWizardLines([]int64{id})
}

// DeleteAccountMergeWizardLines deletes existing account.merge.wizard.line records.
func (c *Client) DeleteAccountMergeWizardLines(ids []int64) error {
	return c.Delete(AccountMergeWizardLineModel, ids)
}

// GetAccountMergeWizardLine gets account.merge.wizard.line existing record.
func (c *Client) GetAccountMergeWizardLine(id int64) (*AccountMergeWizardLine, error) {
	amwls, err := c.GetAccountMergeWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amwls)[0]), nil
}

// GetAccountMergeWizardLines gets account.merge.wizard.line existing records.
func (c *Client) GetAccountMergeWizardLines(ids []int64) (*AccountMergeWizardLines, error) {
	amwls := &AccountMergeWizardLines{}
	if err := c.Read(AccountMergeWizardLineModel, ids, nil, amwls); err != nil {
		return nil, err
	}
	return amwls, nil
}

// FindAccountMergeWizardLine finds account.merge.wizard.line record by querying it with criteria.
func (c *Client) FindAccountMergeWizardLine(criteria *Criteria) (*AccountMergeWizardLine, error) {
	amwls := &AccountMergeWizardLines{}
	if err := c.SearchRead(AccountMergeWizardLineModel, criteria, NewOptions().Limit(1), amwls); err != nil {
		return nil, err
	}
	return &((*amwls)[0]), nil
}

// FindAccountMergeWizardLines finds account.merge.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMergeWizardLines(criteria *Criteria, options *Options) (*AccountMergeWizardLines, error) {
	amwls := &AccountMergeWizardLines{}
	if err := c.SearchRead(AccountMergeWizardLineModel, criteria, options, amwls); err != nil {
		return nil, err
	}
	return amwls, nil
}

// FindAccountMergeWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMergeWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMergeWizardLineModel, criteria, options)
}

// FindAccountMergeWizardLineId finds record id by querying it with criteria.
func (c *Client) FindAccountMergeWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMergeWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
