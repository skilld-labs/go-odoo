package odoo

// AccountEdiCii represents account.edi.cii model.
type AccountEdiCii struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiCiis represents array of account.edi.cii model.
type AccountEdiCiis []AccountEdiCii

// AccountEdiCiiModel is the odoo model name.
const AccountEdiCiiModel = "account.edi.cii"

// Many2One convert AccountEdiCii to *Many2One.
func (aec *AccountEdiCii) Many2One() *Many2One {
	return NewMany2One(aec.Id.Get(), "")
}

// CreateAccountEdiCii creates a new account.edi.cii model and returns its id.
func (c *Client) CreateAccountEdiCii(aec *AccountEdiCii) (int64, error) {
	ids, err := c.CreateAccountEdiCiis([]*AccountEdiCii{aec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiCii creates a new account.edi.cii model and returns its id.
func (c *Client) CreateAccountEdiCiis(aecs []*AccountEdiCii) ([]int64, error) {
	var vv []interface{}
	for _, v := range aecs {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiCiiModel, vv, nil)
}

// UpdateAccountEdiCii updates an existing account.edi.cii record.
func (c *Client) UpdateAccountEdiCii(aec *AccountEdiCii) error {
	return c.UpdateAccountEdiCiis([]int64{aec.Id.Get()}, aec)
}

// UpdateAccountEdiCiis updates existing account.edi.cii records.
// All records (represented by ids) will be updated by aec values.
func (c *Client) UpdateAccountEdiCiis(ids []int64, aec *AccountEdiCii) error {
	return c.Update(AccountEdiCiiModel, ids, aec, nil)
}

// DeleteAccountEdiCii deletes an existing account.edi.cii record.
func (c *Client) DeleteAccountEdiCii(id int64) error {
	return c.DeleteAccountEdiCiis([]int64{id})
}

// DeleteAccountEdiCiis deletes existing account.edi.cii records.
func (c *Client) DeleteAccountEdiCiis(ids []int64) error {
	return c.Delete(AccountEdiCiiModel, ids)
}

// GetAccountEdiCii gets account.edi.cii existing record.
func (c *Client) GetAccountEdiCii(id int64) (*AccountEdiCii, error) {
	aecs, err := c.GetAccountEdiCiis([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aecs)[0]), nil
}

// GetAccountEdiCiis gets account.edi.cii existing records.
func (c *Client) GetAccountEdiCiis(ids []int64) (*AccountEdiCiis, error) {
	aecs := &AccountEdiCiis{}
	if err := c.Read(AccountEdiCiiModel, ids, nil, aecs); err != nil {
		return nil, err
	}
	return aecs, nil
}

// FindAccountEdiCii finds account.edi.cii record by querying it with criteria.
func (c *Client) FindAccountEdiCii(criteria *Criteria) (*AccountEdiCii, error) {
	aecs := &AccountEdiCiis{}
	if err := c.SearchRead(AccountEdiCiiModel, criteria, NewOptions().Limit(1), aecs); err != nil {
		return nil, err
	}
	return &((*aecs)[0]), nil
}

// FindAccountEdiCiis finds account.edi.cii records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiCiis(criteria *Criteria, options *Options) (*AccountEdiCiis, error) {
	aecs := &AccountEdiCiis{}
	if err := c.SearchRead(AccountEdiCiiModel, criteria, options, aecs); err != nil {
		return nil, err
	}
	return aecs, nil
}

// FindAccountEdiCiiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiCiiIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiCiiModel, criteria, options)
}

// FindAccountEdiCiiId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiCiiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiCiiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
