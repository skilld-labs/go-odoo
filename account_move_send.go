package odoo

// AccountMoveSend represents account.move.send model.
type AccountMoveSend struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountMoveSends represents array of account.move.send model.
type AccountMoveSends []AccountMoveSend

// AccountMoveSendModel is the odoo model name.
const AccountMoveSendModel = "account.move.send"

// Many2One convert AccountMoveSend to *Many2One.
func (ams *AccountMoveSend) Many2One() *Many2One {
	return NewMany2One(ams.Id.Get(), "")
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSend(ams *AccountMoveSend) (int64, error) {
	ids, err := c.CreateAccountMoveSends([]*AccountMoveSend{ams})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSends(amss []*AccountMoveSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range amss {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveSendModel, vv, nil)
}

// UpdateAccountMoveSend updates an existing account.move.send record.
func (c *Client) UpdateAccountMoveSend(ams *AccountMoveSend) error {
	return c.UpdateAccountMoveSends([]int64{ams.Id.Get()}, ams)
}

// UpdateAccountMoveSends updates existing account.move.send records.
// All records (represented by ids) will be updated by ams values.
func (c *Client) UpdateAccountMoveSends(ids []int64, ams *AccountMoveSend) error {
	return c.Update(AccountMoveSendModel, ids, ams, nil)
}

// DeleteAccountMoveSend deletes an existing account.move.send record.
func (c *Client) DeleteAccountMoveSend(id int64) error {
	return c.DeleteAccountMoveSends([]int64{id})
}

// DeleteAccountMoveSends deletes existing account.move.send records.
func (c *Client) DeleteAccountMoveSends(ids []int64) error {
	return c.Delete(AccountMoveSendModel, ids)
}

// GetAccountMoveSend gets account.move.send existing record.
func (c *Client) GetAccountMoveSend(id int64) (*AccountMoveSend, error) {
	amss, err := c.GetAccountMoveSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amss)[0]), nil
}

// GetAccountMoveSends gets account.move.send existing records.
func (c *Client) GetAccountMoveSends(ids []int64) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.Read(AccountMoveSendModel, ids, nil, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSend finds account.move.send record by querying it with criteria.
func (c *Client) FindAccountMoveSend(criteria *Criteria) (*AccountMoveSend, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, NewOptions().Limit(1), amss); err != nil {
		return nil, err
	}
	return &((*amss)[0]), nil
}

// FindAccountMoveSends finds account.move.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSends(criteria *Criteria, options *Options) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, options, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveSendModel, criteria, options)
}

// FindAccountMoveSendId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
