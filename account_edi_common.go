package odoo

// AccountEdiCommon represents account.edi.common model.
type AccountEdiCommon struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiCommons represents array of account.edi.common model.
type AccountEdiCommons []AccountEdiCommon

// AccountEdiCommonModel is the odoo model name.
const AccountEdiCommonModel = "account.edi.common"

// Many2One convert AccountEdiCommon to *Many2One.
func (aec *AccountEdiCommon) Many2One() *Many2One {
	return NewMany2One(aec.Id.Get(), "")
}

// CreateAccountEdiCommon creates a new account.edi.common model and returns its id.
func (c *Client) CreateAccountEdiCommon(aec *AccountEdiCommon) (int64, error) {
	ids, err := c.CreateAccountEdiCommons([]*AccountEdiCommon{aec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiCommon creates a new account.edi.common model and returns its id.
func (c *Client) CreateAccountEdiCommons(aecs []*AccountEdiCommon) ([]int64, error) {
	var vv []interface{}
	for _, v := range aecs {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiCommonModel, vv, nil)
}

// UpdateAccountEdiCommon updates an existing account.edi.common record.
func (c *Client) UpdateAccountEdiCommon(aec *AccountEdiCommon) error {
	return c.UpdateAccountEdiCommons([]int64{aec.Id.Get()}, aec)
}

// UpdateAccountEdiCommons updates existing account.edi.common records.
// All records (represented by ids) will be updated by aec values.
func (c *Client) UpdateAccountEdiCommons(ids []int64, aec *AccountEdiCommon) error {
	return c.Update(AccountEdiCommonModel, ids, aec, nil)
}

// DeleteAccountEdiCommon deletes an existing account.edi.common record.
func (c *Client) DeleteAccountEdiCommon(id int64) error {
	return c.DeleteAccountEdiCommons([]int64{id})
}

// DeleteAccountEdiCommons deletes existing account.edi.common records.
func (c *Client) DeleteAccountEdiCommons(ids []int64) error {
	return c.Delete(AccountEdiCommonModel, ids)
}

// GetAccountEdiCommon gets account.edi.common existing record.
func (c *Client) GetAccountEdiCommon(id int64) (*AccountEdiCommon, error) {
	aecs, err := c.GetAccountEdiCommons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aecs)[0]), nil
}

// GetAccountEdiCommons gets account.edi.common existing records.
func (c *Client) GetAccountEdiCommons(ids []int64) (*AccountEdiCommons, error) {
	aecs := &AccountEdiCommons{}
	if err := c.Read(AccountEdiCommonModel, ids, nil, aecs); err != nil {
		return nil, err
	}
	return aecs, nil
}

// FindAccountEdiCommon finds account.edi.common record by querying it with criteria.
func (c *Client) FindAccountEdiCommon(criteria *Criteria) (*AccountEdiCommon, error) {
	aecs := &AccountEdiCommons{}
	if err := c.SearchRead(AccountEdiCommonModel, criteria, NewOptions().Limit(1), aecs); err != nil {
		return nil, err
	}
	return &((*aecs)[0]), nil
}

// FindAccountEdiCommons finds account.edi.common records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiCommons(criteria *Criteria, options *Options) (*AccountEdiCommons, error) {
	aecs := &AccountEdiCommons{}
	if err := c.SearchRead(AccountEdiCommonModel, criteria, options, aecs); err != nil {
		return nil, err
	}
	return aecs, nil
}

// FindAccountEdiCommonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiCommonIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiCommonModel, criteria, options)
}

// FindAccountEdiCommonId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiCommonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiCommonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
