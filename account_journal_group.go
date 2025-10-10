package odoo

// AccountJournalGroup represents account.journal.group model.
type AccountJournalGroup struct {
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	ExcludedJournalIds *Relation `xmlrpc:"excluded_journal_ids,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountJournalGroups represents array of account.journal.group model.
type AccountJournalGroups []AccountJournalGroup

// AccountJournalGroupModel is the odoo model name.
const AccountJournalGroupModel = "account.journal.group"

// Many2One convert AccountJournalGroup to *Many2One.
func (ajg *AccountJournalGroup) Many2One() *Many2One {
	return NewMany2One(ajg.Id.Get(), "")
}

// CreateAccountJournalGroup creates a new account.journal.group model and returns its id.
func (c *Client) CreateAccountJournalGroup(ajg *AccountJournalGroup) (int64, error) {
	ids, err := c.CreateAccountJournalGroups([]*AccountJournalGroup{ajg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountJournalGroup creates a new account.journal.group model and returns its id.
func (c *Client) CreateAccountJournalGroups(ajgs []*AccountJournalGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range ajgs {
		vv = append(vv, v)
	}
	return c.Create(AccountJournalGroupModel, vv, nil)
}

// UpdateAccountJournalGroup updates an existing account.journal.group record.
func (c *Client) UpdateAccountJournalGroup(ajg *AccountJournalGroup) error {
	return c.UpdateAccountJournalGroups([]int64{ajg.Id.Get()}, ajg)
}

// UpdateAccountJournalGroups updates existing account.journal.group records.
// All records (represented by ids) will be updated by ajg values.
func (c *Client) UpdateAccountJournalGroups(ids []int64, ajg *AccountJournalGroup) error {
	return c.Update(AccountJournalGroupModel, ids, ajg, nil)
}

// DeleteAccountJournalGroup deletes an existing account.journal.group record.
func (c *Client) DeleteAccountJournalGroup(id int64) error {
	return c.DeleteAccountJournalGroups([]int64{id})
}

// DeleteAccountJournalGroups deletes existing account.journal.group records.
func (c *Client) DeleteAccountJournalGroups(ids []int64) error {
	return c.Delete(AccountJournalGroupModel, ids)
}

// GetAccountJournalGroup gets account.journal.group existing record.
func (c *Client) GetAccountJournalGroup(id int64) (*AccountJournalGroup, error) {
	ajgs, err := c.GetAccountJournalGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ajgs)[0]), nil
}

// GetAccountJournalGroups gets account.journal.group existing records.
func (c *Client) GetAccountJournalGroups(ids []int64) (*AccountJournalGroups, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.Read(AccountJournalGroupModel, ids, nil, ajgs); err != nil {
		return nil, err
	}
	return ajgs, nil
}

// FindAccountJournalGroup finds account.journal.group record by querying it with criteria.
func (c *Client) FindAccountJournalGroup(criteria *Criteria) (*AccountJournalGroup, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.SearchRead(AccountJournalGroupModel, criteria, NewOptions().Limit(1), ajgs); err != nil {
		return nil, err
	}
	return &((*ajgs)[0]), nil
}

// FindAccountJournalGroups finds account.journal.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalGroups(criteria *Criteria, options *Options) (*AccountJournalGroups, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.SearchRead(AccountJournalGroupModel, criteria, options, ajgs); err != nil {
		return nil, err
	}
	return ajgs, nil
}

// FindAccountJournalGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountJournalGroupModel, criteria, options)
}

// FindAccountJournalGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountJournalGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
