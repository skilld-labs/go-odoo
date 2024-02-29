package odoo

// AccountAccountTag represents account.account.tag model.
type AccountAccountTag struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	Active        *Bool      `xmlrpc:"active,omitempty"`
	Applicability *Selection `xmlrpc:"applicability,omitempty"`
	Color         *Int       `xmlrpc:"color,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAccountTags represents array of account.account.tag model.
type AccountAccountTags []AccountAccountTag

// AccountAccountTagModel is the odoo model name.
const AccountAccountTagModel = "account.account.tag"

// Many2One convert AccountAccountTag to *Many2One.
func (aat *AccountAccountTag) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAccountTag creates a new account.account.tag model and returns its id.
func (c *Client) CreateAccountAccountTag(aat *AccountAccountTag) (int64, error) {
	ids, err := c.CreateAccountAccountTags([]*AccountAccountTag{aat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAccountTags creates a new account.account.tag model and returns its id.
func (c *Client) CreateAccountAccountTags(aats []*AccountAccountTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range aats {
		vv = append(vv, v)
	}
	return c.Create(AccountAccountTagModel, vv, nil)
}

// UpdateAccountAccountTag updates an existing account.account.tag record.
func (c *Client) UpdateAccountAccountTag(aat *AccountAccountTag) error {
	return c.UpdateAccountAccountTags([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAccountTags updates existing account.account.tag records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAccountTags(ids []int64, aat *AccountAccountTag) error {
	return c.Update(AccountAccountTagModel, ids, aat, nil)
}

// DeleteAccountAccountTag deletes an existing account.account.tag record.
func (c *Client) DeleteAccountAccountTag(id int64) error {
	return c.DeleteAccountAccountTags([]int64{id})
}

// DeleteAccountAccountTags deletes existing account.account.tag records.
func (c *Client) DeleteAccountAccountTags(ids []int64) error {
	return c.Delete(AccountAccountTagModel, ids)
}

// GetAccountAccountTag gets account.account.tag existing record.
func (c *Client) GetAccountAccountTag(id int64) (*AccountAccountTag, error) {
	aats, err := c.GetAccountAccountTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aats)[0]), nil
}

// GetAccountAccountTags gets account.account.tag existing records.
func (c *Client) GetAccountAccountTags(ids []int64) (*AccountAccountTags, error) {
	aats := &AccountAccountTags{}
	if err := c.Read(AccountAccountTagModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTag finds account.account.tag record by querying it with criteria.
func (c *Client) FindAccountAccountTag(criteria *Criteria) (*AccountAccountTag, error) {
	aats := &AccountAccountTags{}
	if err := c.SearchRead(AccountAccountTagModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	return &((*aats)[0]), nil
}

// FindAccountAccountTags finds account.account.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTags(criteria *Criteria, options *Options) (*AccountAccountTags, error) {
	aats := &AccountAccountTags{}
	if err := c.SearchRead(AccountAccountTagModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAccountTagModel, criteria, options)
}

// FindAccountAccountTagId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
