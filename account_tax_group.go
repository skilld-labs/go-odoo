package odoo

// AccountTaxGroup represents account.tax.group model.
type AccountTaxGroup struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxGroups represents array of account.tax.group model.
type AccountTaxGroups []AccountTaxGroup

// AccountTaxGroupModel is the odoo model name.
const AccountTaxGroupModel = "account.tax.group"

// Many2One convert AccountTaxGroup to *Many2One.
func (atg *AccountTaxGroup) Many2One() *Many2One {
	return NewMany2One(atg.Id.Get(), "")
}

// CreateAccountTaxGroup creates a new account.tax.group model and returns its id.
func (c *Client) CreateAccountTaxGroup(atg *AccountTaxGroup) (int64, error) {
	ids, err := c.CreateAccountTaxGroups([]*AccountTaxGroup{atg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTaxGroups creates a new account.tax.group model and returns its id.
func (c *Client) CreateAccountTaxGroups(atgs []*AccountTaxGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range atgs {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxGroupModel, vv, nil)
}

// UpdateAccountTaxGroup updates an existing account.tax.group record.
func (c *Client) UpdateAccountTaxGroup(atg *AccountTaxGroup) error {
	return c.UpdateAccountTaxGroups([]int64{atg.Id.Get()}, atg)
}

// UpdateAccountTaxGroups updates existing account.tax.group records.
// All records (represented by ids) will be updated by atg values.
func (c *Client) UpdateAccountTaxGroups(ids []int64, atg *AccountTaxGroup) error {
	return c.Update(AccountTaxGroupModel, ids, atg, nil)
}

// DeleteAccountTaxGroup deletes an existing account.tax.group record.
func (c *Client) DeleteAccountTaxGroup(id int64) error {
	return c.DeleteAccountTaxGroups([]int64{id})
}

// DeleteAccountTaxGroups deletes existing account.tax.group records.
func (c *Client) DeleteAccountTaxGroups(ids []int64) error {
	return c.Delete(AccountTaxGroupModel, ids)
}

// GetAccountTaxGroup gets account.tax.group existing record.
func (c *Client) GetAccountTaxGroup(id int64) (*AccountTaxGroup, error) {
	atgs, err := c.GetAccountTaxGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atgs)[0]), nil
}

// GetAccountTaxGroups gets account.tax.group existing records.
func (c *Client) GetAccountTaxGroups(ids []int64) (*AccountTaxGroups, error) {
	atgs := &AccountTaxGroups{}
	if err := c.Read(AccountTaxGroupModel, ids, nil, atgs); err != nil {
		return nil, err
	}
	return atgs, nil
}

// FindAccountTaxGroup finds account.tax.group record by querying it with criteria.
func (c *Client) FindAccountTaxGroup(criteria *Criteria) (*AccountTaxGroup, error) {
	atgs := &AccountTaxGroups{}
	if err := c.SearchRead(AccountTaxGroupModel, criteria, NewOptions().Limit(1), atgs); err != nil {
		return nil, err
	}
	return &((*atgs)[0]), nil
}

// FindAccountTaxGroups finds account.tax.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxGroups(criteria *Criteria, options *Options) (*AccountTaxGroups, error) {
	atgs := &AccountTaxGroups{}
	if err := c.SearchRead(AccountTaxGroupModel, criteria, options, atgs); err != nil {
		return nil, err
	}
	return atgs, nil
}

// FindAccountTaxGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTaxGroupModel, criteria, options)
}

// FindAccountTaxGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
