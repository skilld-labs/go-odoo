package odoo

// AccountAccountType represents account.account.type model.
type AccountAccountType struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	IncludeInitialBalance *Bool      `xmlrpc:"include_initial_balance,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	Note                  *String    `xmlrpc:"note,omptempty"`
	Type                  *Selection `xmlrpc:"type,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAccountTypes represents array of account.account.type model.
type AccountAccountTypes []AccountAccountType

// AccountAccountTypeModel is the odoo model name.
const AccountAccountTypeModel = "account.account.type"

// Many2One convert AccountAccountType to *Many2One.
func (aat *AccountAccountType) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAccountType creates a new account.account.type model and returns its id.
func (c *Client) CreateAccountAccountType(aat *AccountAccountType) (int64, error) {
	ids, err := c.CreateAccountAccountTypes([]*AccountAccountType{aat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAccountType creates a new account.account.type model and returns its id.
func (c *Client) CreateAccountAccountTypes(aats []*AccountAccountType) ([]int64, error) {
	var vv []interface{}
	for _, v := range aats {
		vv = append(vv, v)
	}
	return c.Create(AccountAccountTypeModel, vv, nil)
}

// UpdateAccountAccountType updates an existing account.account.type record.
func (c *Client) UpdateAccountAccountType(aat *AccountAccountType) error {
	return c.UpdateAccountAccountTypes([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAccountTypes updates existing account.account.type records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAccountTypes(ids []int64, aat *AccountAccountType) error {
	return c.Update(AccountAccountTypeModel, ids, aat, nil)
}

// DeleteAccountAccountType deletes an existing account.account.type record.
func (c *Client) DeleteAccountAccountType(id int64) error {
	return c.DeleteAccountAccountTypes([]int64{id})
}

// DeleteAccountAccountTypes deletes existing account.account.type records.
func (c *Client) DeleteAccountAccountTypes(ids []int64) error {
	return c.Delete(AccountAccountTypeModel, ids)
}

// GetAccountAccountType gets account.account.type existing record.
func (c *Client) GetAccountAccountType(id int64) (*AccountAccountType, error) {
	aats, err := c.GetAccountAccountTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aats)[0]), nil
}

// GetAccountAccountTypes gets account.account.type existing records.
func (c *Client) GetAccountAccountTypes(ids []int64) (*AccountAccountTypes, error) {
	aats := &AccountAccountTypes{}
	if err := c.Read(AccountAccountTypeModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountType finds account.account.type record by querying it with criteria.
func (c *Client) FindAccountAccountType(criteria *Criteria) (*AccountAccountType, error) {
	aats := &AccountAccountTypes{}
	if err := c.SearchRead(AccountAccountTypeModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	return &((*aats)[0]), nil
}

// FindAccountAccountTypes finds account.account.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTypes(criteria *Criteria, options *Options) (*AccountAccountTypes, error) {
	aats := &AccountAccountTypes{}
	if err := c.SearchRead(AccountAccountTypeModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAccountTypeModel, criteria, options)
}

// FindAccountAccountTypeId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
