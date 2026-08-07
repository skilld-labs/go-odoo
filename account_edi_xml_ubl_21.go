package odoo

// AccountEdiXmlUbl21 represents account.edi.xml.ubl_21 model.
type AccountEdiXmlUbl21 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUbl21s represents array of account.edi.xml.ubl_21 model.
type AccountEdiXmlUbl21s []AccountEdiXmlUbl21

// AccountEdiXmlUbl21Model is the odoo model name.
const AccountEdiXmlUbl21Model = "account.edi.xml.ubl_21"

// Many2One convert AccountEdiXmlUbl21 to *Many2One.
func (aexu *AccountEdiXmlUbl21) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUbl21 creates a new account.edi.xml.ubl_21 model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl21(aexu *AccountEdiXmlUbl21) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUbl21s([]*AccountEdiXmlUbl21{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUbl21 creates a new account.edi.xml.ubl_21 model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl21s(aexus []*AccountEdiXmlUbl21) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUbl21Model, vv, nil)
}

// UpdateAccountEdiXmlUbl21 updates an existing account.edi.xml.ubl_21 record.
func (c *Client) UpdateAccountEdiXmlUbl21(aexu *AccountEdiXmlUbl21) error {
	return c.UpdateAccountEdiXmlUbl21s([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUbl21s updates existing account.edi.xml.ubl_21 records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUbl21s(ids []int64, aexu *AccountEdiXmlUbl21) error {
	return c.Update(AccountEdiXmlUbl21Model, ids, aexu, nil)
}

// DeleteAccountEdiXmlUbl21 deletes an existing account.edi.xml.ubl_21 record.
func (c *Client) DeleteAccountEdiXmlUbl21(id int64) error {
	return c.DeleteAccountEdiXmlUbl21s([]int64{id})
}

// DeleteAccountEdiXmlUbl21s deletes existing account.edi.xml.ubl_21 records.
func (c *Client) DeleteAccountEdiXmlUbl21s(ids []int64) error {
	return c.Delete(AccountEdiXmlUbl21Model, ids)
}

// GetAccountEdiXmlUbl21 gets account.edi.xml.ubl_21 existing record.
func (c *Client) GetAccountEdiXmlUbl21(id int64) (*AccountEdiXmlUbl21, error) {
	aexus, err := c.GetAccountEdiXmlUbl21s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUbl21s gets account.edi.xml.ubl_21 existing records.
func (c *Client) GetAccountEdiXmlUbl21s(ids []int64) (*AccountEdiXmlUbl21s, error) {
	aexus := &AccountEdiXmlUbl21s{}
	if err := c.Read(AccountEdiXmlUbl21Model, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl21 finds account.edi.xml.ubl_21 record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl21(criteria *Criteria) (*AccountEdiXmlUbl21, error) {
	aexus := &AccountEdiXmlUbl21s{}
	if err := c.SearchRead(AccountEdiXmlUbl21Model, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUbl21s finds account.edi.xml.ubl_21 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl21s(criteria *Criteria, options *Options) (*AccountEdiXmlUbl21s, error) {
	aexus := &AccountEdiXmlUbl21s{}
	if err := c.SearchRead(AccountEdiXmlUbl21Model, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl21Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl21Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUbl21Model, criteria, options)
}

// FindAccountEdiXmlUbl21Id finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl21Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUbl21Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
