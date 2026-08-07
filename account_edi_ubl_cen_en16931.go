package odoo

// AccountEdiUblCenEn16931 represents account.edi.ubl_cen_en16931 model.
type AccountEdiUblCenEn16931 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEdiUblCenEn16931s represents array of account.edi.ubl_cen_en16931 model.
type AccountEdiUblCenEn16931s []AccountEdiUblCenEn16931

// AccountEdiUblCenEn16931Model is the odoo model name.
const AccountEdiUblCenEn16931Model = "account.edi.ubl_cen_en16931"

// Many2One convert AccountEdiUblCenEn16931 to *Many2One.
func (aeu *AccountEdiUblCenEn16931) Many2One() *Many2One {
	return NewMany2One(aeu.Id.Get(), "")
}

// CreateAccountEdiUblCenEn16931 creates a new account.edi.ubl_cen_en16931 model and returns its id.
func (c *Client) CreateAccountEdiUblCenEn16931(aeu *AccountEdiUblCenEn16931) (int64, error) {
	ids, err := c.CreateAccountEdiUblCenEn16931s([]*AccountEdiUblCenEn16931{aeu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiUblCenEn16931 creates a new account.edi.ubl_cen_en16931 model and returns its id.
func (c *Client) CreateAccountEdiUblCenEn16931s(aeus []*AccountEdiUblCenEn16931) ([]int64, error) {
	var vv []interface{}
	for _, v := range aeus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiUblCenEn16931Model, vv, nil)
}

// UpdateAccountEdiUblCenEn16931 updates an existing account.edi.ubl_cen_en16931 record.
func (c *Client) UpdateAccountEdiUblCenEn16931(aeu *AccountEdiUblCenEn16931) error {
	return c.UpdateAccountEdiUblCenEn16931s([]int64{aeu.Id.Get()}, aeu)
}

// UpdateAccountEdiUblCenEn16931s updates existing account.edi.ubl_cen_en16931 records.
// All records (represented by ids) will be updated by aeu values.
func (c *Client) UpdateAccountEdiUblCenEn16931s(ids []int64, aeu *AccountEdiUblCenEn16931) error {
	return c.Update(AccountEdiUblCenEn16931Model, ids, aeu, nil)
}

// DeleteAccountEdiUblCenEn16931 deletes an existing account.edi.ubl_cen_en16931 record.
func (c *Client) DeleteAccountEdiUblCenEn16931(id int64) error {
	return c.DeleteAccountEdiUblCenEn16931s([]int64{id})
}

// DeleteAccountEdiUblCenEn16931s deletes existing account.edi.ubl_cen_en16931 records.
func (c *Client) DeleteAccountEdiUblCenEn16931s(ids []int64) error {
	return c.Delete(AccountEdiUblCenEn16931Model, ids)
}

// GetAccountEdiUblCenEn16931 gets account.edi.ubl_cen_en16931 existing record.
func (c *Client) GetAccountEdiUblCenEn16931(id int64) (*AccountEdiUblCenEn16931, error) {
	aeus, err := c.GetAccountEdiUblCenEn16931s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// GetAccountEdiUblCenEn16931s gets account.edi.ubl_cen_en16931 existing records.
func (c *Client) GetAccountEdiUblCenEn16931s(ids []int64) (*AccountEdiUblCenEn16931s, error) {
	aeus := &AccountEdiUblCenEn16931s{}
	if err := c.Read(AccountEdiUblCenEn16931Model, ids, nil, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblCenEn16931 finds account.edi.ubl_cen_en16931 record by querying it with criteria.
func (c *Client) FindAccountEdiUblCenEn16931(criteria *Criteria) (*AccountEdiUblCenEn16931, error) {
	aeus := &AccountEdiUblCenEn16931s{}
	if err := c.SearchRead(AccountEdiUblCenEn16931Model, criteria, NewOptions().Limit(1), aeus); err != nil {
		return nil, err
	}
	return &((*aeus)[0]), nil
}

// FindAccountEdiUblCenEn16931s finds account.edi.ubl_cen_en16931 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblCenEn16931s(criteria *Criteria, options *Options) (*AccountEdiUblCenEn16931s, error) {
	aeus := &AccountEdiUblCenEn16931s{}
	if err := c.SearchRead(AccountEdiUblCenEn16931Model, criteria, options, aeus); err != nil {
		return nil, err
	}
	return aeus, nil
}

// FindAccountEdiUblCenEn16931Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiUblCenEn16931Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiUblCenEn16931Model, criteria, options)
}

// FindAccountEdiUblCenEn16931Id finds record id by querying it with criteria.
func (c *Client) FindAccountEdiUblCenEn16931Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiUblCenEn16931Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
