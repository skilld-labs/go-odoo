package odoo

// AccountEdiXmlUbl21Fr represents account.edi.xml.ubl_21_fr model.
type AccountEdiXmlUbl21Fr struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUbl21Frs represents array of account.edi.xml.ubl_21_fr model.
type AccountEdiXmlUbl21Frs []AccountEdiXmlUbl21Fr

// AccountEdiXmlUbl21FrModel is the odoo model name.
const AccountEdiXmlUbl21FrModel = "account.edi.xml.ubl_21_fr"

// Many2One convert AccountEdiXmlUbl21Fr to *Many2One.
func (aexu *AccountEdiXmlUbl21Fr) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUbl21Fr creates a new account.edi.xml.ubl_21_fr model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl21Fr(aexu *AccountEdiXmlUbl21Fr) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUbl21Frs([]*AccountEdiXmlUbl21Fr{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUbl21Fr creates a new account.edi.xml.ubl_21_fr model and returns its id.
func (c *Client) CreateAccountEdiXmlUbl21Frs(aexus []*AccountEdiXmlUbl21Fr) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUbl21FrModel, vv, nil)
}

// UpdateAccountEdiXmlUbl21Fr updates an existing account.edi.xml.ubl_21_fr record.
func (c *Client) UpdateAccountEdiXmlUbl21Fr(aexu *AccountEdiXmlUbl21Fr) error {
	return c.UpdateAccountEdiXmlUbl21Frs([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUbl21Frs updates existing account.edi.xml.ubl_21_fr records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUbl21Frs(ids []int64, aexu *AccountEdiXmlUbl21Fr) error {
	return c.Update(AccountEdiXmlUbl21FrModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUbl21Fr deletes an existing account.edi.xml.ubl_21_fr record.
func (c *Client) DeleteAccountEdiXmlUbl21Fr(id int64) error {
	return c.DeleteAccountEdiXmlUbl21Frs([]int64{id})
}

// DeleteAccountEdiXmlUbl21Frs deletes existing account.edi.xml.ubl_21_fr records.
func (c *Client) DeleteAccountEdiXmlUbl21Frs(ids []int64) error {
	return c.Delete(AccountEdiXmlUbl21FrModel, ids)
}

// GetAccountEdiXmlUbl21Fr gets account.edi.xml.ubl_21_fr existing record.
func (c *Client) GetAccountEdiXmlUbl21Fr(id int64) (*AccountEdiXmlUbl21Fr, error) {
	aexus, err := c.GetAccountEdiXmlUbl21Frs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUbl21Frs gets account.edi.xml.ubl_21_fr existing records.
func (c *Client) GetAccountEdiXmlUbl21Frs(ids []int64) (*AccountEdiXmlUbl21Frs, error) {
	aexus := &AccountEdiXmlUbl21Frs{}
	if err := c.Read(AccountEdiXmlUbl21FrModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl21Fr finds account.edi.xml.ubl_21_fr record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl21Fr(criteria *Criteria) (*AccountEdiXmlUbl21Fr, error) {
	aexus := &AccountEdiXmlUbl21Frs{}
	if err := c.SearchRead(AccountEdiXmlUbl21FrModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUbl21Frs finds account.edi.xml.ubl_21_fr records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl21Frs(criteria *Criteria, options *Options) (*AccountEdiXmlUbl21Frs, error) {
	aexus := &AccountEdiXmlUbl21Frs{}
	if err := c.SearchRead(AccountEdiXmlUbl21FrModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUbl21FrIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUbl21FrIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUbl21FrModel, criteria, options)
}

// FindAccountEdiXmlUbl21FrId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUbl21FrId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUbl21FrModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
