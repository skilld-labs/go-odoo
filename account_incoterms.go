package odoo

// AccountIncoterms represents account.incoterms model.
type AccountIncoterms struct {
	Active      *Bool     `xmlrpc:"active,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountIncotermss represents array of account.incoterms model.
type AccountIncotermss []AccountIncoterms

// AccountIncotermsModel is the odoo model name.
const AccountIncotermsModel = "account.incoterms"

// Many2One convert AccountIncoterms to *Many2One.
func (ai *AccountIncoterms) Many2One() *Many2One {
	return NewMany2One(ai.Id.Get(), "")
}

// CreateAccountIncoterms creates a new account.incoterms model and returns its id.
func (c *Client) CreateAccountIncoterms(ai *AccountIncoterms) (int64, error) {
	ids, err := c.CreateAccountIncotermss([]*AccountIncoterms{ai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountIncoterms creates a new account.incoterms model and returns its id.
func (c *Client) CreateAccountIncotermss(ais []*AccountIncoterms) ([]int64, error) {
	var vv []interface{}
	for _, v := range ais {
		vv = append(vv, v)
	}
	return c.Create(AccountIncotermsModel, vv, nil)
}

// UpdateAccountIncoterms updates an existing account.incoterms record.
func (c *Client) UpdateAccountIncoterms(ai *AccountIncoterms) error {
	return c.UpdateAccountIncotermss([]int64{ai.Id.Get()}, ai)
}

// UpdateAccountIncotermss updates existing account.incoterms records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAccountIncotermss(ids []int64, ai *AccountIncoterms) error {
	return c.Update(AccountIncotermsModel, ids, ai, nil)
}

// DeleteAccountIncoterms deletes an existing account.incoterms record.
func (c *Client) DeleteAccountIncoterms(id int64) error {
	return c.DeleteAccountIncotermss([]int64{id})
}

// DeleteAccountIncotermss deletes existing account.incoterms records.
func (c *Client) DeleteAccountIncotermss(ids []int64) error {
	return c.Delete(AccountIncotermsModel, ids)
}

// GetAccountIncoterms gets account.incoterms existing record.
func (c *Client) GetAccountIncoterms(id int64) (*AccountIncoterms, error) {
	ais, err := c.GetAccountIncotermss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ais)[0]), nil
}

// GetAccountIncotermss gets account.incoterms existing records.
func (c *Client) GetAccountIncotermss(ids []int64) (*AccountIncotermss, error) {
	ais := &AccountIncotermss{}
	if err := c.Read(AccountIncotermsModel, ids, nil, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountIncoterms finds account.incoterms record by querying it with criteria.
func (c *Client) FindAccountIncoterms(criteria *Criteria) (*AccountIncoterms, error) {
	ais := &AccountIncotermss{}
	if err := c.SearchRead(AccountIncotermsModel, criteria, NewOptions().Limit(1), ais); err != nil {
		return nil, err
	}
	return &((*ais)[0]), nil
}

// FindAccountIncotermss finds account.incoterms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIncotermss(criteria *Criteria, options *Options) (*AccountIncotermss, error) {
	ais := &AccountIncotermss{}
	if err := c.SearchRead(AccountIncotermsModel, criteria, options, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountIncotermsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIncotermsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountIncotermsModel, criteria, options)
}

// FindAccountIncotermsId finds record id by querying it with criteria.
func (c *Client) FindAccountIncotermsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountIncotermsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
