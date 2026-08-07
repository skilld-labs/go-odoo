package odoo

// AccountEdiXmlCii represents account.edi.xml.cii model.
type AccountEdiXmlCii struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlCiis represents array of account.edi.xml.cii model.
type AccountEdiXmlCiis []AccountEdiXmlCii

// AccountEdiXmlCiiModel is the odoo model name.
const AccountEdiXmlCiiModel = "account.edi.xml.cii"

// Many2One convert AccountEdiXmlCii to *Many2One.
func (aexc *AccountEdiXmlCii) Many2One() *Many2One {
	return NewMany2One(aexc.Id.Get(), "")
}

// CreateAccountEdiXmlCii creates a new account.edi.xml.cii model and returns its id.
func (c *Client) CreateAccountEdiXmlCii(aexc *AccountEdiXmlCii) (int64, error) {
	ids, err := c.CreateAccountEdiXmlCiis([]*AccountEdiXmlCii{aexc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlCii creates a new account.edi.xml.cii model and returns its id.
func (c *Client) CreateAccountEdiXmlCiis(aexcs []*AccountEdiXmlCii) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexcs {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlCiiModel, vv, nil)
}

// UpdateAccountEdiXmlCii updates an existing account.edi.xml.cii record.
func (c *Client) UpdateAccountEdiXmlCii(aexc *AccountEdiXmlCii) error {
	return c.UpdateAccountEdiXmlCiis([]int64{aexc.Id.Get()}, aexc)
}

// UpdateAccountEdiXmlCiis updates existing account.edi.xml.cii records.
// All records (represented by ids) will be updated by aexc values.
func (c *Client) UpdateAccountEdiXmlCiis(ids []int64, aexc *AccountEdiXmlCii) error {
	return c.Update(AccountEdiXmlCiiModel, ids, aexc, nil)
}

// DeleteAccountEdiXmlCii deletes an existing account.edi.xml.cii record.
func (c *Client) DeleteAccountEdiXmlCii(id int64) error {
	return c.DeleteAccountEdiXmlCiis([]int64{id})
}

// DeleteAccountEdiXmlCiis deletes existing account.edi.xml.cii records.
func (c *Client) DeleteAccountEdiXmlCiis(ids []int64) error {
	return c.Delete(AccountEdiXmlCiiModel, ids)
}

// GetAccountEdiXmlCii gets account.edi.xml.cii existing record.
func (c *Client) GetAccountEdiXmlCii(id int64) (*AccountEdiXmlCii, error) {
	aexcs, err := c.GetAccountEdiXmlCiis([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexcs)[0]), nil
}

// GetAccountEdiXmlCiis gets account.edi.xml.cii existing records.
func (c *Client) GetAccountEdiXmlCiis(ids []int64) (*AccountEdiXmlCiis, error) {
	aexcs := &AccountEdiXmlCiis{}
	if err := c.Read(AccountEdiXmlCiiModel, ids, nil, aexcs); err != nil {
		return nil, err
	}
	return aexcs, nil
}

// FindAccountEdiXmlCii finds account.edi.xml.cii record by querying it with criteria.
func (c *Client) FindAccountEdiXmlCii(criteria *Criteria) (*AccountEdiXmlCii, error) {
	aexcs := &AccountEdiXmlCiis{}
	if err := c.SearchRead(AccountEdiXmlCiiModel, criteria, NewOptions().Limit(1), aexcs); err != nil {
		return nil, err
	}
	return &((*aexcs)[0]), nil
}

// FindAccountEdiXmlCiis finds account.edi.xml.cii records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlCiis(criteria *Criteria, options *Options) (*AccountEdiXmlCiis, error) {
	aexcs := &AccountEdiXmlCiis{}
	if err := c.SearchRead(AccountEdiXmlCiiModel, criteria, options, aexcs); err != nil {
		return nil, err
	}
	return aexcs, nil
}

// FindAccountEdiXmlCiiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlCiiIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlCiiModel, criteria, options)
}

// FindAccountEdiXmlCiiId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlCiiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlCiiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
