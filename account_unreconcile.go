package odoo

// AccountUnreconcile represents account.unreconcile model.
type AccountUnreconcile struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountUnreconciles represents array of account.unreconcile model.
type AccountUnreconciles []AccountUnreconcile

// AccountUnreconcileModel is the odoo model name.
const AccountUnreconcileModel = "account.unreconcile"

// Many2One convert AccountUnreconcile to *Many2One.
func (au *AccountUnreconcile) Many2One() *Many2One {
	return NewMany2One(au.Id.Get(), "")
}

// CreateAccountUnreconcile creates a new account.unreconcile model and returns its id.
func (c *Client) CreateAccountUnreconcile(au *AccountUnreconcile) (int64, error) {
	ids, err := c.CreateAccountUnreconciles([]*AccountUnreconcile{au})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountUnreconcile creates a new account.unreconcile model and returns its id.
func (c *Client) CreateAccountUnreconciles(aus []*AccountUnreconcile) ([]int64, error) {
	var vv []interface{}
	for _, v := range aus {
		vv = append(vv, v)
	}
	return c.Create(AccountUnreconcileModel, vv, nil)
}

// UpdateAccountUnreconcile updates an existing account.unreconcile record.
func (c *Client) UpdateAccountUnreconcile(au *AccountUnreconcile) error {
	return c.UpdateAccountUnreconciles([]int64{au.Id.Get()}, au)
}

// UpdateAccountUnreconciles updates existing account.unreconcile records.
// All records (represented by ids) will be updated by au values.
func (c *Client) UpdateAccountUnreconciles(ids []int64, au *AccountUnreconcile) error {
	return c.Update(AccountUnreconcileModel, ids, au, nil)
}

// DeleteAccountUnreconcile deletes an existing account.unreconcile record.
func (c *Client) DeleteAccountUnreconcile(id int64) error {
	return c.DeleteAccountUnreconciles([]int64{id})
}

// DeleteAccountUnreconciles deletes existing account.unreconcile records.
func (c *Client) DeleteAccountUnreconciles(ids []int64) error {
	return c.Delete(AccountUnreconcileModel, ids)
}

// GetAccountUnreconcile gets account.unreconcile existing record.
func (c *Client) GetAccountUnreconcile(id int64) (*AccountUnreconcile, error) {
	aus, err := c.GetAccountUnreconciles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// GetAccountUnreconciles gets account.unreconcile existing records.
func (c *Client) GetAccountUnreconciles(ids []int64) (*AccountUnreconciles, error) {
	aus := &AccountUnreconciles{}
	if err := c.Read(AccountUnreconcileModel, ids, nil, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindAccountUnreconcile finds account.unreconcile record by querying it with criteria.
func (c *Client) FindAccountUnreconcile(criteria *Criteria) (*AccountUnreconcile, error) {
	aus := &AccountUnreconciles{}
	if err := c.SearchRead(AccountUnreconcileModel, criteria, NewOptions().Limit(1), aus); err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// FindAccountUnreconciles finds account.unreconcile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountUnreconciles(criteria *Criteria, options *Options) (*AccountUnreconciles, error) {
	aus := &AccountUnreconciles{}
	if err := c.SearchRead(AccountUnreconcileModel, criteria, options, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindAccountUnreconcileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountUnreconcileIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountUnreconcileModel, criteria, options)
}

// FindAccountUnreconcileId finds record id by querying it with criteria.
func (c *Client) FindAccountUnreconcileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountUnreconcileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
