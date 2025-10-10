package odoo

// AccountCodeMapping represents account.code.mapping model.
type AccountCodeMapping struct {
	AccountId   *Many2One `xmlrpc:"account_id,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
}

// AccountCodeMappings represents array of account.code.mapping model.
type AccountCodeMappings []AccountCodeMapping

// AccountCodeMappingModel is the odoo model name.
const AccountCodeMappingModel = "account.code.mapping"

// Many2One convert AccountCodeMapping to *Many2One.
func (acm *AccountCodeMapping) Many2One() *Many2One {
	return NewMany2One(acm.Id.Get(), "")
}

// CreateAccountCodeMapping creates a new account.code.mapping model and returns its id.
func (c *Client) CreateAccountCodeMapping(acm *AccountCodeMapping) (int64, error) {
	ids, err := c.CreateAccountCodeMappings([]*AccountCodeMapping{acm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCodeMapping creates a new account.code.mapping model and returns its id.
func (c *Client) CreateAccountCodeMappings(acms []*AccountCodeMapping) ([]int64, error) {
	var vv []interface{}
	for _, v := range acms {
		vv = append(vv, v)
	}
	return c.Create(AccountCodeMappingModel, vv, nil)
}

// UpdateAccountCodeMapping updates an existing account.code.mapping record.
func (c *Client) UpdateAccountCodeMapping(acm *AccountCodeMapping) error {
	return c.UpdateAccountCodeMappings([]int64{acm.Id.Get()}, acm)
}

// UpdateAccountCodeMappings updates existing account.code.mapping records.
// All records (represented by ids) will be updated by acm values.
func (c *Client) UpdateAccountCodeMappings(ids []int64, acm *AccountCodeMapping) error {
	return c.Update(AccountCodeMappingModel, ids, acm, nil)
}

// DeleteAccountCodeMapping deletes an existing account.code.mapping record.
func (c *Client) DeleteAccountCodeMapping(id int64) error {
	return c.DeleteAccountCodeMappings([]int64{id})
}

// DeleteAccountCodeMappings deletes existing account.code.mapping records.
func (c *Client) DeleteAccountCodeMappings(ids []int64) error {
	return c.Delete(AccountCodeMappingModel, ids)
}

// GetAccountCodeMapping gets account.code.mapping existing record.
func (c *Client) GetAccountCodeMapping(id int64) (*AccountCodeMapping, error) {
	acms, err := c.GetAccountCodeMappings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acms)[0]), nil
}

// GetAccountCodeMappings gets account.code.mapping existing records.
func (c *Client) GetAccountCodeMappings(ids []int64) (*AccountCodeMappings, error) {
	acms := &AccountCodeMappings{}
	if err := c.Read(AccountCodeMappingModel, ids, nil, acms); err != nil {
		return nil, err
	}
	return acms, nil
}

// FindAccountCodeMapping finds account.code.mapping record by querying it with criteria.
func (c *Client) FindAccountCodeMapping(criteria *Criteria) (*AccountCodeMapping, error) {
	acms := &AccountCodeMappings{}
	if err := c.SearchRead(AccountCodeMappingModel, criteria, NewOptions().Limit(1), acms); err != nil {
		return nil, err
	}
	return &((*acms)[0]), nil
}

// FindAccountCodeMappings finds account.code.mapping records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCodeMappings(criteria *Criteria, options *Options) (*AccountCodeMappings, error) {
	acms := &AccountCodeMappings{}
	if err := c.SearchRead(AccountCodeMappingModel, criteria, options, acms); err != nil {
		return nil, err
	}
	return acms, nil
}

// FindAccountCodeMappingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCodeMappingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCodeMappingModel, criteria, options)
}

// FindAccountCodeMappingId finds record id by querying it with criteria.
func (c *Client) FindAccountCodeMappingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCodeMappingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
