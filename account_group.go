package odoo

// AccountGroup represents account.group model.
type AccountGroup struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CodePrefix  *String   `xmlrpc:"code_prefix,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	ParentLeft  *Int      `xmlrpc:"parent_left,omitempty"`
	ParentRight *Int      `xmlrpc:"parent_right,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountGroups represents array of account.group model.
type AccountGroups []AccountGroup

// AccountGroupModel is the odoo model name.
const AccountGroupModel = "account.group"

// Many2One convert AccountGroup to *Many2One.
func (ag *AccountGroup) Many2One() *Many2One {
	return NewMany2One(ag.Id.Get(), "")
}

// CreateAccountGroup creates a new account.group model and returns its id.
func (c *Client) CreateAccountGroup(ag *AccountGroup) (int64, error) {
	ids, err := c.CreateAccountGroups([]*AccountGroup{ag})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGroups creates a new account.group model and returns its id.
func (c *Client) CreateAccountGroups(ags []*AccountGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range ags {
		vv = append(vv, v)
	}
	return c.Create(AccountGroupModel, vv, nil)
}

// UpdateAccountGroup updates an existing account.group record.
func (c *Client) UpdateAccountGroup(ag *AccountGroup) error {
	return c.UpdateAccountGroups([]int64{ag.Id.Get()}, ag)
}

// UpdateAccountGroups updates existing account.group records.
// All records (represented by ids) will be updated by ag values.
func (c *Client) UpdateAccountGroups(ids []int64, ag *AccountGroup) error {
	return c.Update(AccountGroupModel, ids, ag, nil)
}

// DeleteAccountGroup deletes an existing account.group record.
func (c *Client) DeleteAccountGroup(id int64) error {
	return c.DeleteAccountGroups([]int64{id})
}

// DeleteAccountGroups deletes existing account.group records.
func (c *Client) DeleteAccountGroups(ids []int64) error {
	return c.Delete(AccountGroupModel, ids)
}

// GetAccountGroup gets account.group existing record.
func (c *Client) GetAccountGroup(id int64) (*AccountGroup, error) {
	ags, err := c.GetAccountGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ags)[0]), nil
}

// GetAccountGroups gets account.group existing records.
func (c *Client) GetAccountGroups(ids []int64) (*AccountGroups, error) {
	ags := &AccountGroups{}
	if err := c.Read(AccountGroupModel, ids, nil, ags); err != nil {
		return nil, err
	}
	return ags, nil
}

// FindAccountGroup finds account.group record by querying it with criteria.
func (c *Client) FindAccountGroup(criteria *Criteria) (*AccountGroup, error) {
	ags := &AccountGroups{}
	if err := c.SearchRead(AccountGroupModel, criteria, NewOptions().Limit(1), ags); err != nil {
		return nil, err
	}
	return &((*ags)[0]), nil
}

// FindAccountGroups finds account.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroups(criteria *Criteria, options *Options) (*AccountGroups, error) {
	ags := &AccountGroups{}
	if err := c.SearchRead(AccountGroupModel, criteria, options, ags); err != nil {
		return nil, err
	}
	return ags, nil
}

// FindAccountGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGroupModel, criteria, options)
}

// FindAccountGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
