package odoo

// AccountEdiXmlUbl20 represents account.edi.xml.ubl_20 model.
type AccountEdiXmlUbl20 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUbl20s represents array of account.edi.xml.ubl_20 model.
type AccountEdiXmlUbl20s []AccountEdiXmlUbl20

// AccountEdiXmlUbl20Model is the odoo model name.
const AccountEdiXmlUbl20Model = "account.edi.xml.ubl_20"

// Many2One convert AccountEdiXmlUbl20 to *Many2One.
func (aexu *AccountEdiXmlUbl20) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUbl20 creates a new account.edi.xml.ubl_20 model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl20(aexu *AccountEdiXmlUbl20) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUbl20s([]*AccountEdiXmlUbl20{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUbl20 creates a new account.edi.xml.ubl_20 model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl20s(aexus []*AccountEdiXmlUbl20) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUbl20Model, vv, nil)
}

// UpdateAccountEdiXmlUbl20 updates an existing account.edi.xml.ubl_20 record.
func (c *Client) UpdateAccountEdiXmlUbl20(aexu *AccountEdiXmlUbl20) error {
	return c.UpdateAccountEdiXmlUbl20s([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUbl20s updates existing account.edi.xml.ubl_20 records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUbl20s(ids []int64, aexu *AccountEdiXmlUbl20) error {
	return c.Update(AccountEdiXmlUbl20Model, ids, aexu, nil)
}

// DeleteAccountEdiXmlUbl20 deletes an existing account.edi.xml.ubl_20 record.
func (c *Client) DeleteAccountEdiXmlUbl20(id int64) error {
	return c.DeleteAccountEdiXmlUbl20s([]int64{id})
}

// DeleteAccountEdiXmlUbl20s deletes existing account.edi.xml.ubl_20 records.
func (c *Client) DeleteAccountEdiXmlUbl20s(ids []int64) error {
	return c.Delete(AccountEdiXmlUbl20Model, ids)
}

// GetAccountEdiXmlUbl20 gets account.edi.xml.ubl_20 existing record.
func (c *Client) GetAccountEdiXmlUbl20(id int64) (*AccountEdiXmlUbl20, error) {
	aexus, err := c.GetAccountEdiXmlUbl20s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUbl20s gets account.edi.xml.ubl_20 existing records.
func (c *Client) GetAccountEdiXmlUbl20s(ids []int64) (*AccountEdiXmlUbl20s, error) {
	aexus := &AccountEdiXmlUbl20s{}
	if err := c.Read(AccountEdiXmlUbl20Model, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl20 finds account.edi.xml.ubl_20 record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl20(criteria *Criteria) (*AccountEdiXmlUbl20, error) {
	aexus := &AccountEdiXmlUbl20s{}
	if err := c.SearchRead(AccountEdiXmlUbl20Model, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUbl20s finds account.edi.xml.ubl_20 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl20s(criteria *Criteria, options *Options) (*AccountEdiXmlUbl20s, error) {
	aexus := &AccountEdiXmlUbl20s{}
	if err := c.SearchRead(AccountEdiXmlUbl20Model, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl20Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl20Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUbl20Model, criteria, options)
}

// FindAccountEdiXmlUbl20Id finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl20Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUbl20Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
