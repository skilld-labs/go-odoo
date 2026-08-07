package odoo

// AccountEdiXmlUblDe represents account.edi.xml.ubl_de model.
type AccountEdiXmlUblDe struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblDes represents array of account.edi.xml.ubl_de model.
type AccountEdiXmlUblDes []AccountEdiXmlUblDe

// AccountEdiXmlUblDeModel is the odoo model name.
const AccountEdiXmlUblDeModel = "account.edi.xml.ubl_de"

// Many2One convert AccountEdiXmlUblDe to *Many2One.
func (aexu *AccountEdiXmlUblDe) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblDe creates a new account.edi.xml.ubl_de model and returns its id.
func (c *Client) CreateAccountEdiXmlUblDe(aexu *AccountEdiXmlUblDe) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblDes([]*AccountEdiXmlUblDe{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblDe creates a new account.edi.xml.ubl_de model and returns its id.
func (c *Client) CreateAccountEdiXmlUblDes(aexus []*AccountEdiXmlUblDe) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblDeModel, vv, nil)
}

// UpdateAccountEdiXmlUblDe updates an existing account.edi.xml.ubl_de record.
func (c *Client) UpdateAccountEdiXmlUblDe(aexu *AccountEdiXmlUblDe) error {
	return c.UpdateAccountEdiXmlUblDes([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblDes updates existing account.edi.xml.ubl_de records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblDes(ids []int64, aexu *AccountEdiXmlUblDe) error {
	return c.Update(AccountEdiXmlUblDeModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblDe deletes an existing account.edi.xml.ubl_de record.
func (c *Client) DeleteAccountEdiXmlUblDe(id int64) error {
	return c.DeleteAccountEdiXmlUblDes([]int64{id})
}

// DeleteAccountEdiXmlUblDes deletes existing account.edi.xml.ubl_de records.
func (c *Client) DeleteAccountEdiXmlUblDes(ids []int64) error {
	return c.Delete(AccountEdiXmlUblDeModel, ids)
}

// GetAccountEdiXmlUblDe gets account.edi.xml.ubl_de existing record.
func (c *Client) GetAccountEdiXmlUblDe(id int64) (*AccountEdiXmlUblDe, error) {
	aexus, err := c.GetAccountEdiXmlUblDes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblDes gets account.edi.xml.ubl_de existing records.
func (c *Client) GetAccountEdiXmlUblDes(ids []int64) (*AccountEdiXmlUblDes, error) {
	aexus := &AccountEdiXmlUblDes{}
	if err := c.Read(AccountEdiXmlUblDeModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblDe finds account.edi.xml.ubl_de record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblDe(criteria *Criteria) (*AccountEdiXmlUblDe, error) {
	aexus := &AccountEdiXmlUblDes{}
	if err := c.SearchRead(AccountEdiXmlUblDeModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblDes finds account.edi.xml.ubl_de records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblDes(criteria *Criteria, options *Options) (*AccountEdiXmlUblDes, error) {
	aexus := &AccountEdiXmlUblDes{}
	if err := c.SearchRead(AccountEdiXmlUblDeModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblDeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblDeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblDeModel, criteria, options)
}

// FindAccountEdiXmlUblDeId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblDeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblDeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
