package odoo

// AccountEdiUblPint represents account.edi.ubl_pint model.
type AccountEdiUblPint struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiUblPints represents array of account.edi.ubl_pint model.
type AccountEdiUblPints []AccountEdiUblPint

// AccountEdiUblPintModel is the odoo model name.
const AccountEdiUblPintModel = "account.edi.ubl_pint"

// Many2One convert AccountEdiUblPint to *Many2One.
func (aeu *AccountEdiUblPint) Many2One() *Many2One {
	return NewMany2One(aeu.Id.Get(), "")
}

// CreateAccountEdiUblPint creates a new account.edi.ubl_pint model and returns its id.
func (c *Client) CreateAccountEdiUblPint(aeu *AccountEdiUblPint) (int64, error) {
	ids, err := c.CreateAccountEdiUblPints([]*AccountEdiUblPint{aeu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiUblPint creates a new account.edi.ubl_pint model and returns its id.
func (c *Client) CreateAccountEdiUblPints(aeus []*AccountEdiUblPint) ([]int64, error) {
	var vv []interface{}
	for _, v := range aeus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiUblPintModel, vv, nil)
}

// UpdateAccountEdiUblPint updates an existing account.edi.ubl_pint record.
func (c *Client) UpdateAccountEdiUblPint(aeu *AccountEdiUblPint) error {
	return c.UpdateAccountEdiUblPints([]int64{aeu.Id.Get()}, aeu)
}

// UpdateAccountEdiUblPints updates existing account.edi.ubl_pint records.
// All records (represented by ids) will be updated by aeu values.
func (c *Client) UpdateAccountEdiUblPints(ids []int64, aeu *AccountEdiUblPint) error {
	return c.Update(AccountEdiUblPintModel, ids, aeu, nil)
}

// DeleteAccountEdiUblPint deletes an existing account.edi.ubl_pint record.
func (c *Client) DeleteAccountEdiUblPint(id int64) error {
	return c.DeleteAccountEdiUblPints([]int64{id})
}

// DeleteAccountEdiUblPints deletes existing account.edi.ubl_pint records.
func (c *Client) DeleteAccountEdiUblPints(ids []int64) error {
	return c.Delete(AccountEdiUblPintModel, ids)
}

// GetAccountEdiUblPint gets account.edi.ubl_pint existing record.
func (c *Client) GetAccountEdiUblPint(id int64) (*AccountEdiUblPint, error) {
	aeus, err := c.GetAccountEdiUblPints([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// GetAccountEdiUblPints gets account.edi.ubl_pint existing records.
func (c *Client) GetAccountEdiUblPints(ids []int64) (*AccountEdiUblPints, error) {
	aeus := &AccountEdiUblPints{}
	if err := c.Read(AccountEdiUblPintModel, ids, nil, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblPint finds account.edi.ubl_pint record by querying it with criteria.
func (c *Client) FindAccountEdiUblPint(criteria *Criteria) (*AccountEdiUblPint, error) {
	aeus := &AccountEdiUblPints{}
	if err := c.SearchRead(AccountEdiUblPintModel, criteria, NewOptions().Limit(1), aeus); err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// FindAccountEdiUblPints finds account.edi.ubl_pint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblPints(criteria *Criteria, options *Options) (*AccountEdiUblPints, error) {
	aeus := &AccountEdiUblPints{}
	if err := c.SearchRead(AccountEdiUblPintModel, criteria, options, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblPintIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblPintIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiUblPintModel, criteria, options)
}

// FindAccountEdiUblPintId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiUblPintId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiUblPintModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
