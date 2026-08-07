package odoo

// AccountEdiUblPintEu represents account.edi.ubl_pint_eu model.
type AccountEdiUblPintEu struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiUblPintEus represents array of account.edi.ubl_pint_eu model.
type AccountEdiUblPintEus []AccountEdiUblPintEu

// AccountEdiUblPintEuModel is the odoo model name.
const AccountEdiUblPintEuModel = "account.edi.ubl_pint_eu"

// Many2One convert AccountEdiUblPintEu to *Many2One.
func (aeu *AccountEdiUblPintEu) Many2One() *Many2One {
	return NewMany2One(aeu.Id.Get(), "")
}

// CreateAccountEdiUblPintEu creates a new account.edi.ubl_pint_eu model and returns its id.
func (c *Client) CreateAccountEdiUblPintEu(aeu *AccountEdiUblPintEu) (int64, error) {
	ids, err := c.CreateAccountEdiUblPintEus([]*AccountEdiUblPintEu{aeu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiUblPintEu creates a new account.edi.ubl_pint_eu model and returns its id.
func (c *Client) CreateAccountEdiUblPintEus(aeus []*AccountEdiUblPintEu) ([]int64, error) {
	var vv []interface{}
	for _, v := range aeus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiUblPintEuModel, vv, nil)
}

// UpdateAccountEdiUblPintEu updates an existing account.edi.ubl_pint_eu record.
func (c *Client) UpdateAccountEdiUblPintEu(aeu *AccountEdiUblPintEu) error {
	return c.UpdateAccountEdiUblPintEus([]int64{aeu.Id.Get()}, aeu)
}

// UpdateAccountEdiUblPintEus updates existing account.edi.ubl_pint_eu records.
// All records (represented by ids) will be updated by aeu values.
func (c *Client) UpdateAccountEdiUblPintEus(ids []int64, aeu *AccountEdiUblPintEu) error {
	return c.Update(AccountEdiUblPintEuModel, ids, aeu, nil)
}

// DeleteAccountEdiUblPintEu deletes an existing account.edi.ubl_pint_eu record.
func (c *Client) DeleteAccountEdiUblPintEu(id int64) error {
	return c.DeleteAccountEdiUblPintEus([]int64{id})
}

// DeleteAccountEdiUblPintEus deletes existing account.edi.ubl_pint_eu records.
func (c *Client) DeleteAccountEdiUblPintEus(ids []int64) error {
	return c.Delete(AccountEdiUblPintEuModel, ids)
}

// GetAccountEdiUblPintEu gets account.edi.ubl_pint_eu existing record.
func (c *Client) GetAccountEdiUblPintEu(id int64) (*AccountEdiUblPintEu, error) {
	aeus, err := c.GetAccountEdiUblPintEus([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// GetAccountEdiUblPintEus gets account.edi.ubl_pint_eu existing records.
func (c *Client) GetAccountEdiUblPintEus(ids []int64) (*AccountEdiUblPintEus, error) {
	aeus := &AccountEdiUblPintEus{}
	if err := c.Read(AccountEdiUblPintEuModel, ids, nil, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblPintEu finds account.edi.ubl_pint_eu record by querying it with criteria.
func (c *Client) FindAccountEdiUblPintEu(criteria *Criteria) (*AccountEdiUblPintEu, error) {
	aeus := &AccountEdiUblPintEus{}
	if err := c.SearchRead(AccountEdiUblPintEuModel, criteria, NewOptions().Limit(1), aeus); err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// FindAccountEdiUblPintEus finds account.edi.ubl_pint_eu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblPintEus(criteria *Criteria, options *Options) (*AccountEdiUblPintEus, error) {
	aeus := &AccountEdiUblPintEus{}
	if err := c.SearchRead(AccountEdiUblPintEuModel, criteria, options, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblPintEuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblPintEuIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiUblPintEuModel, criteria, options)
}

// FindAccountEdiUblPintEuId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiUblPintEuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiUblPintEuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
