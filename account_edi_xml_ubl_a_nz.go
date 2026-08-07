package odoo

// AccountEdiXmlUblANz represents account.edi.xml.ubl_a_nz model.
type AccountEdiXmlUblANz struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblANzs represents array of account.edi.xml.ubl_a_nz model.
type AccountEdiXmlUblANzs []AccountEdiXmlUblANz

// AccountEdiXmlUblANzModel is the odoo model name.
const AccountEdiXmlUblANzModel = "account.edi.xml.ubl_a_nz"

// Many2One convert AccountEdiXmlUblANz to *Many2One.
func (aexu *AccountEdiXmlUblANz) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblANz creates a new account.edi.xml.ubl_a_nz model and returns its id.
func (c *Client) CreateAccountEdiXmlUblANz(aexu *AccountEdiXmlUblANz) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblANzs([]*AccountEdiXmlUblANz{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblANz creates a new account.edi.xml.ubl_a_nz model and returns its id.
func (c *Client) CreateAccountEdiXmlUblANzs(aexus []*AccountEdiXmlUblANz) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblANzModel, vv, nil)
}

// UpdateAccountEdiXmlUblANz updates an existing account.edi.xml.ubl_a_nz record.
func (c *Client) UpdateAccountEdiXmlUblANz(aexu *AccountEdiXmlUblANz) error {
	return c.UpdateAccountEdiXmlUblANzs([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblANzs updates existing account.edi.xml.ubl_a_nz records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblANzs(ids []int64, aexu *AccountEdiXmlUblANz) error {
	return c.Update(AccountEdiXmlUblANzModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblANz deletes an existing account.edi.xml.ubl_a_nz record.
func (c *Client) DeleteAccountEdiXmlUblANz(id int64) error {
	return c.DeleteAccountEdiXmlUblANzs([]int64{id})
}

// DeleteAccountEdiXmlUblANzs deletes existing account.edi.xml.ubl_a_nz records.
func (c *Client) DeleteAccountEdiXmlUblANzs(ids []int64) error {
	return c.Delete(AccountEdiXmlUblANzModel, ids)
}

// GetAccountEdiXmlUblANz gets account.edi.xml.ubl_a_nz existing record.
func (c *Client) GetAccountEdiXmlUblANz(id int64) (*AccountEdiXmlUblANz, error) {
	aexus, err := c.GetAccountEdiXmlUblANzs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblANzs gets account.edi.xml.ubl_a_nz existing records.
func (c *Client) GetAccountEdiXmlUblANzs(ids []int64) (*AccountEdiXmlUblANzs, error) {
	aexus := &AccountEdiXmlUblANzs{}
	if err := c.Read(AccountEdiXmlUblANzModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblANz finds account.edi.xml.ubl_a_nz record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblANz(criteria *Criteria) (*AccountEdiXmlUblANz, error) {
	aexus := &AccountEdiXmlUblANzs{}
	if err := c.SearchRead(AccountEdiXmlUblANzModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblANzs finds account.edi.xml.ubl_a_nz records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblANzs(criteria *Criteria, options *Options) (*AccountEdiXmlUblANzs, error) {
	aexus := &AccountEdiXmlUblANzs{}
	if err := c.SearchRead(AccountEdiXmlUblANzModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblANzIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblANzIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblANzModel, criteria, options)
}

// FindAccountEdiXmlUblANzId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblANzId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblANzModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
