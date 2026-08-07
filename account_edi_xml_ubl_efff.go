package odoo

// AccountEdiXmlUblEfff represents account.edi.xml.ubl_efff model.
type AccountEdiXmlUblEfff struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiXmlUblEfffs represents array of account.edi.xml.ubl_efff model.
type AccountEdiXmlUblEfffs []AccountEdiXmlUblEfff

// AccountEdiXmlUblEfffModel is the odoo model name.
const AccountEdiXmlUblEfffModel = "account.edi.xml.ubl_efff"

// Many2One convert AccountEdiXmlUblEfff to *Many2One.
func (aexu *AccountEdiXmlUblEfff) Many2One() *Many2One {
	return NewMany2One(aexu.Id.Get(), "")
}

// CreateAccountEdiXmlUblEfff creates a new account.edi.xml.ubl_efff model and returns its id.
func (c *Client) CreateAccountEdiXmlUblEfff(aexu *AccountEdiXmlUblEfff) (int64, error) {
	ids, err := c.CreateAccountEdiXmlUblEfffs([]*AccountEdiXmlUblEfff{aexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiXmlUblEfff creates a new account.edi.xml.ubl_efff model and returns its id.
func (c *Client) CreateAccountEdiXmlUblEfffs(aexus []*AccountEdiXmlUblEfff) ([]int64, error) {
	var vv []interface{}
	for _, v := range aexus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiXmlUblEfffModel, vv, nil)
}

// UpdateAccountEdiXmlUblEfff updates an existing account.edi.xml.ubl_efff record.
func (c *Client) UpdateAccountEdiXmlUblEfff(aexu *AccountEdiXmlUblEfff) error {
	return c.UpdateAccountEdiXmlUblEfffs([]int64{aexu.Id.Get()}, aexu)
}

// UpdateAccountEdiXmlUblEfffs updates existing account.edi.xml.ubl_efff records.
// All records (represented by ids) will be updated by aexu values.
func (c *Client) UpdateAccountEdiXmlUblEfffs(ids []int64, aexu *AccountEdiXmlUblEfff) error {
	return c.Update(AccountEdiXmlUblEfffModel, ids, aexu, nil)
}

// DeleteAccountEdiXmlUblEfff deletes an existing account.edi.xml.ubl_efff record.
func (c *Client) DeleteAccountEdiXmlUblEfff(id int64) error {
	return c.DeleteAccountEdiXmlUblEfffs([]int64{id})
}

// DeleteAccountEdiXmlUblEfffs deletes existing account.edi.xml.ubl_efff records.
func (c *Client) DeleteAccountEdiXmlUblEfffs(ids []int64) error {
	return c.Delete(AccountEdiXmlUblEfffModel, ids)
}

// GetAccountEdiXmlUblEfff gets account.edi.xml.ubl_efff existing record.
func (c *Client) GetAccountEdiXmlUblEfff(id int64) (*AccountEdiXmlUblEfff, error) {
	aexus, err := c.GetAccountEdiXmlUblEfffs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// GetAccountEdiXmlUblEfffs gets account.edi.xml.ubl_efff existing records.
func (c *Client) GetAccountEdiXmlUblEfffs(ids []int64) (*AccountEdiXmlUblEfffs, error) {
	aexus := &AccountEdiXmlUblEfffs{}
	if err := c.Read(AccountEdiXmlUblEfffModel, ids, nil, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblEfff finds account.edi.xml.ubl_efff record by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblEfff(criteria *Criteria) (*AccountEdiXmlUblEfff, error) {
	aexus := &AccountEdiXmlUblEfffs{}
	if err := c.SearchRead(AccountEdiXmlUblEfffModel, criteria, NewOptions().Limit(1), aexus); err != nil {
		return nil, err
	}
	return &((*aexus)[0]), nil
}

// FindAccountEdiXmlUblEfffs finds account.edi.xml.ubl_efff records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblEfffs(criteria *Criteria, options *Options) (*AccountEdiXmlUblEfffs, error) {
	aexus := &AccountEdiXmlUblEfffs{}
	if err := c.SearchRead(AccountEdiXmlUblEfffModel, criteria, options, aexus); err != nil {
		return nil, err
	}
	return aexus, nil
}

// FindAccountEdiXmlUblEfffIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiXmlUblEfffIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiXmlUblEfffModel, criteria, options)
}

// FindAccountEdiXmlUblEfffId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiXmlUblEfffId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiXmlUblEfffModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
