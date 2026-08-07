package odoo

// AccountEdiUbl represents account.edi.ubl model.
type AccountEdiUbl struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiUbls represents array of account.edi.ubl model.
type AccountEdiUbls []AccountEdiUbl

// AccountEdiUblModel is the odoo model name.
const AccountEdiUblModel = "account.edi.ubl"

// Many2One convert AccountEdiUbl to *Many2One.
func (aeu *AccountEdiUbl) Many2One() *Many2One {
	return NewMany2One(aeu.Id.Get(), "")
}

// CreateAccountEdiUbl creates a new account.edi.ubl model and returns its id.
func (c *Client) CreateAccountEdiUbl(aeu *AccountEdiUbl) (int64, error) {
	ids, err := c.CreateAccountEdiUbls([]*AccountEdiUbl{aeu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiUbl creates a new account.edi.ubl model and returns its id.
func (c *Client) CreateAccountEdiUbls(aeus []*AccountEdiUbl) ([]int64, error) {
	var vv []interface{}
	for _, v := range aeus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiUblModel, vv, nil)
}

// UpdateAccountEdiUbl updates an existing account.edi.ubl record.
func (c *Client) UpdateAccountEdiUbl(aeu *AccountEdiUbl) error {
	return c.UpdateAccountEdiUbls([]int64{aeu.Id.Get()}, aeu)
}

// UpdateAccountEdiUbls updates existing account.edi.ubl records.
// All records (represented by ids) will be updated by aeu values.
func (c *Client) UpdateAccountEdiUbls(ids []int64, aeu *AccountEdiUbl) error {
	return c.Update(AccountEdiUblModel, ids, aeu, nil)
}

// DeleteAccountEdiUbl deletes an existing account.edi.ubl record.
func (c *Client) DeleteAccountEdiUbl(id int64) error {
	return c.DeleteAccountEdiUbls([]int64{id})
}

// DeleteAccountEdiUbls deletes existing account.edi.ubl records.
func (c *Client) DeleteAccountEdiUbls(ids []int64) error {
	return c.Delete(AccountEdiUblModel, ids)
}

// GetAccountEdiUbl gets account.edi.ubl existing record.
func (c *Client) GetAccountEdiUbl(id int64) (*AccountEdiUbl, error) {
	aeus, err := c.GetAccountEdiUbls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// GetAccountEdiUbls gets account.edi.ubl existing records.
func (c *Client) GetAccountEdiUbls(ids []int64) (*AccountEdiUbls, error) {
	aeus := &AccountEdiUbls{}
	if err := c.Read(AccountEdiUblModel, ids, nil, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUbl finds account.edi.ubl record by querying it with criteria.
func (c *Client) FindAccountEdiUbl(criteria *Criteria) (*AccountEdiUbl, error) {
	aeus := &AccountEdiUbls{}
	if err := c.SearchRead(AccountEdiUblModel, criteria, NewOptions().Limit(1), aeus); err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// FindAccountEdiUbls finds account.edi.ubl records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUbls(criteria *Criteria, options *Options) (*AccountEdiUbls, error) {
	aeus := &AccountEdiUbls{}
	if err := c.SearchRead(AccountEdiUblModel, criteria, options, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiUblModel, criteria, options)
}

// FindAccountEdiUblId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiUblId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiUblModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
