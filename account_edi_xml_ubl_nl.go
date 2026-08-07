package odoo

// AccountEdiXmlUblNl represents account.edi.xml.ubl_nl model.
type AccountEdiXmlUblNl struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblNls represents array of account.edi.xml.ubl_nl model.
type AccountEdiXmlUblNls []AccountEdiXmlUblNl

// AccountEdiXmlUblNlModel is the odoo model name.
const AccountEdiXmlUblNlModel = "account.edi.xml.ubl_nl"

// Many2One convert AccountEdiXmlUblNl to *Many2One.
func (aexu *AccountEdiXmlUblNl) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblNl creates a new account.edi.xml.ubl_nl model and returns its id.
func (c *Client) CreateAccountEdiXmlUblNl(aexu *AccountEdiXmlUblNl) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblNls([]*AccountEdiXmlUblNl{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblNl creates a new account.edi.xml.ubl_nl model and returns its id.
func (c *Client) CreateAccountEdiXmlUblNls(aexus []*AccountEdiXmlUblNl) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblNlModel, vv, nil)
}

// UpdateAccountEdiXmlUblNl updates an existing account.edi.xml.ubl_nl record.
func (c *Client) UpdateAccountEdiXmlUblNl(aexu *AccountEdiXmlUblNl) error {
	return c.UpdateAccountEdiXmlUblNls([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblNls updates existing account.edi.xml.ubl_nl records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblNls(ids []int64, aexu *AccountEdiXmlUblNl) error {
	return c.Update(AccountEdiXmlUblNlModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblNl deletes an existing account.edi.xml.ubl_nl record.
func (c *Client) DeleteAccountEdiXmlUblNl(id int64) error {
	return c.DeleteAccountEdiXmlUblNls([]int64{id})
}

// DeleteAccountEdiXmlUblNls deletes existing account.edi.xml.ubl_nl records.
func (c *Client) DeleteAccountEdiXmlUblNls(ids []int64) error {
	return c.Delete(AccountEdiXmlUblNlModel, ids)
}

// GetAccountEdiXmlUblNl gets account.edi.xml.ubl_nl existing record.
func (c *Client) GetAccountEdiXmlUblNl(id int64) (*AccountEdiXmlUblNl, error) {
	aexus, err := c.GetAccountEdiXmlUblNls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblNls gets account.edi.xml.ubl_nl existing records.
func (c *Client) GetAccountEdiXmlUblNls(ids []int64) (*AccountEdiXmlUblNls, error) {
	aexus := &AccountEdiXmlUblNls{}
	if err := c.Read(AccountEdiXmlUblNlModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblNl finds account.edi.xml.ubl_nl record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblNl(criteria *Criteria) (*AccountEdiXmlUblNl, error) {
	aexus := &AccountEdiXmlUblNls{}
	if err := c.SearchRead(AccountEdiXmlUblNlModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblNls finds account.edi.xml.ubl_nl records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblNls(criteria *Criteria, options *Options) (*AccountEdiXmlUblNls, error) {
	aexus := &AccountEdiXmlUblNls{}
	if err := c.SearchRead(AccountEdiXmlUblNlModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblNlIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblNlIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblNlModel, criteria, options)
}

// FindAccountEdiXmlUblNlId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblNlId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblNlModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
