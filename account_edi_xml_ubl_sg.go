package odoo

// AccountEdiXmlUblSg represents account.edi.xml.ubl_sg model.
type AccountEdiXmlUblSg struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblSgs represents array of account.edi.xml.ubl_sg model.
type AccountEdiXmlUblSgs []AccountEdiXmlUblSg

// AccountEdiXmlUblSgModel is the odoo model name.
const AccountEdiXmlUblSgModel = "account.edi.xml.ubl_sg"

// Many2One convert AccountEdiXmlUblSg to *Many2One.
func (aexu *AccountEdiXmlUblSg) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblSg creates a new account.edi.xml.ubl_sg model and returns its id.
func (c *Client) CreateAccountEdiXmlUblSg(aexu *AccountEdiXmlUblSg) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblSgs([]*AccountEdiXmlUblSg{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblSg creates a new account.edi.xml.ubl_sg model and returns its id.
func (c *Client) CreateAccountEdiXmlUblSgs(aexus []*AccountEdiXmlUblSg) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblSgModel, vv, nil)
}

// UpdateAccountEdiXmlUblSg updates an existing account.edi.xml.ubl_sg record.
func (c *Client) UpdateAccountEdiXmlUblSg(aexu *AccountEdiXmlUblSg) error {
	return c.UpdateAccountEdiXmlUblSgs([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblSgs updates existing account.edi.xml.ubl_sg records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblSgs(ids []int64, aexu *AccountEdiXmlUblSg) error {
	return c.Update(AccountEdiXmlUblSgModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblSg deletes an existing account.edi.xml.ubl_sg record.
func (c *Client) DeleteAccountEdiXmlUblSg(id int64) error {
	return c.DeleteAccountEdiXmlUblSgs([]int64{id})
}

// DeleteAccountEdiXmlUblSgs deletes existing account.edi.xml.ubl_sg records.
func (c *Client) DeleteAccountEdiXmlUblSgs(ids []int64) error {
	return c.Delete(AccountEdiXmlUblSgModel, ids)
}

// GetAccountEdiXmlUblSg gets account.edi.xml.ubl_sg existing record.
func (c *Client) GetAccountEdiXmlUblSg(id int64) (*AccountEdiXmlUblSg, error) {
	aexus, err := c.GetAccountEdiXmlUblSgs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblSgs gets account.edi.xml.ubl_sg existing records.
func (c *Client) GetAccountEdiXmlUblSgs(ids []int64) (*AccountEdiXmlUblSgs, error) {
	aexus := &AccountEdiXmlUblSgs{}
	if err := c.Read(AccountEdiXmlUblSgModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblSg finds account.edi.xml.ubl_sg record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblSg(criteria *Criteria) (*AccountEdiXmlUblSg, error) {
	aexus := &AccountEdiXmlUblSgs{}
	if err := c.SearchRead(AccountEdiXmlUblSgModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblSgs finds account.edi.xml.ubl_sg records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblSgs(criteria *Criteria, options *Options) (*AccountEdiXmlUblSgs, error) {
	aexus := &AccountEdiXmlUblSgs{}
	if err := c.SearchRead(AccountEdiXmlUblSgModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblSgIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblSgIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblSgModel, criteria, options)
}

// FindAccountEdiXmlUblSgId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblSgId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblSgModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
