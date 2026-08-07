package odoo

// AccountEdiXmlUblBis3 represents account.edi.xml.ubl_bis3 model.
type AccountEdiXmlUblBis3 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblBis3s represents array of account.edi.xml.ubl_bis3 model.
type AccountEdiXmlUblBis3s []AccountEdiXmlUblBis3

// AccountEdiXmlUblBis3Model is the odoo model name.
const AccountEdiXmlUblBis3Model = "account.edi.xml.ubl_bis3"

// Many2One convert AccountEdiXmlUblBis3 to *Many2One.
func (aexu *AccountEdiXmlUblBis3) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblBis3 creates a new account.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreateAccountEdiXmlUblBis3(aexu *AccountEdiXmlUblBis3) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblBis3s([]*AccountEdiXmlUblBis3{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblBis3 creates a new account.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreateAccountEdiXmlUblBis3s(aexus []*AccountEdiXmlUblBis3) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblBis3Model, vv, nil)
}

// UpdateAccountEdiXmlUblBis3 updates an existing account.edi.xml.ubl_bis3 record.
func (c *Client) UpdateAccountEdiXmlUblBis3(aexu *AccountEdiXmlUblBis3) error {
	return c.UpdateAccountEdiXmlUblBis3s([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblBis3s updates existing account.edi.xml.ubl_bis3 records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblBis3s(ids []int64, aexu *AccountEdiXmlUblBis3) error {
	return c.Update(AccountEdiXmlUblBis3Model, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblBis3 deletes an existing account.edi.xml.ubl_bis3 record.
func (c *Client) DeleteAccountEdiXmlUblBis3(id int64) error {
	return c.DeleteAccountEdiXmlUblBis3s([]int64{id})
}

// DeleteAccountEdiXmlUblBis3s deletes existing account.edi.xml.ubl_bis3 records.
func (c *Client) DeleteAccountEdiXmlUblBis3s(ids []int64) error {
	return c.Delete(AccountEdiXmlUblBis3Model, ids)
}

// GetAccountEdiXmlUblBis3 gets account.edi.xml.ubl_bis3 existing record.
func (c *Client) GetAccountEdiXmlUblBis3(id int64) (*AccountEdiXmlUblBis3, error) {
	aexus, err := c.GetAccountEdiXmlUblBis3s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblBis3s gets account.edi.xml.ubl_bis3 existing records.
func (c *Client) GetAccountEdiXmlUblBis3s(ids []int64) (*AccountEdiXmlUblBis3s, error) {
	aexus := &AccountEdiXmlUblBis3s{}
	if err := c.Read(AccountEdiXmlUblBis3Model, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblBis3 finds account.edi.xml.ubl_bis3 record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblBis3(criteria *Criteria) (*AccountEdiXmlUblBis3, error) {
	aexus := &AccountEdiXmlUblBis3s{}
	if err := c.SearchRead(AccountEdiXmlUblBis3Model, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblBis3s finds account.edi.xml.ubl_bis3 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblBis3s(criteria *Criteria, options *Options) (*AccountEdiXmlUblBis3s, error) {
	aexus := &AccountEdiXmlUblBis3s{}
	if err := c.SearchRead(AccountEdiXmlUblBis3Model, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblBis3Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblBis3Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblBis3Model, criteria, options)
}

// FindAccountEdiXmlUblBis3Id finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblBis3Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblBis3Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
