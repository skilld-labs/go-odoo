package odoo

// AccountPeppolClarification represents account.peppol.clarification model.
type AccountPeppolClarification struct {
	Code           *String    `xmlrpc:"code,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description    *String    `xmlrpc:"description,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	ListIdentifier *Selection `xmlrpc:"list_identifier,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountPeppolClarifications represents array of account.peppol.clarification model.
type AccountPeppolClarifications []AccountPeppolClarification

// AccountPeppolClarificationModel is the odoo model name.
const AccountPeppolClarificationModel = "account.peppol.clarification"

// Many2One convert AccountPeppolClarification to *Many2One.
func (apc *AccountPeppolClarification) Many2One() *Many2One {
	return NewMany2One(apc.Id.Get(), "")
}

// CreateAccountPeppolClarification creates a new account.peppol.clarification model and returns its id.
func (c *Client) CreateAccountPeppolClarification(apc *AccountPeppolClarification) (int64, error) {
	ids, err := c.CreateAccountPeppolClarifications([]*AccountPeppolClarification{apc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPeppolClarification creates a new account.peppol.clarification model and returns its id.
func (c *Client) CreateAccountPeppolClarifications(apcs []*AccountPeppolClarification) ([]int64, error) {
	var vv []interface{}
	for _, v := range apcs {
		vv = append(vv, v)
	}
	return c.Create(AccountPeppolClarificationModel, vv, nil)
}

// UpdateAccountPeppolClarification updates an existing account.peppol.clarification record.
func (c *Client) UpdateAccountPeppolClarification(apc *AccountPeppolClarification) error {
	return c.UpdateAccountPeppolClarifications([]int64{apc.Id.Get()}, apc)
}

// UpdateAccountPeppolClarifications updates existing account.peppol.clarification records.
// All records (represented by ids) will be updated by apc values.
func (c *Client) UpdateAccountPeppolClarifications(ids []int64, apc *AccountPeppolClarification) error {
	return c.Update(AccountPeppolClarificationModel, ids, apc, nil)
}

// DeleteAccountPeppolClarification deletes an existing account.peppol.clarification record.
func (c *Client) DeleteAccountPeppolClarification(id int64) error {
	return c.DeleteAccountPeppolClarifications([]int64{id})
}

// DeleteAccountPeppolClarifications deletes existing account.peppol.clarification records.
func (c *Client) DeleteAccountPeppolClarifications(ids []int64) error {
	return c.Delete(AccountPeppolClarificationModel, ids)
}

// GetAccountPeppolClarification gets account.peppol.clarification existing record.
func (c *Client) GetAccountPeppolClarification(id int64) (*AccountPeppolClarification, error) {
	apcs, err := c.GetAccountPeppolClarifications([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apcs)[0]), nil
}

// GetAccountPeppolClarifications gets account.peppol.clarification existing records.
func (c *Client) GetAccountPeppolClarifications(ids []int64) (*AccountPeppolClarifications, error) {
	apcs := &AccountPeppolClarifications{}
	if err := c.Read(AccountPeppolClarificationModel, ids, nil, apcs); err != nil {
		return nil, err
	}
	return apcs, nil
}

// FindAccountPeppolClarification finds account.peppol.clarification record by querying it with criteria.
func (c *Client) FindAccountPeppolClarification(criteria *Criteria) (*AccountPeppolClarification, error) {
	apcs := &AccountPeppolClarifications{}
	if err := c.SearchRead(AccountPeppolClarificationModel, criteria, NewOptions().Limit(1), apcs); err != nil {
		return nil, err
	}
	return &((*apcs)[0]), nil
}

// FindAccountPeppolClarifications finds account.peppol.clarification records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolClarifications(criteria *Criteria, options *Options) (*AccountPeppolClarifications, error) {
	apcs := &AccountPeppolClarifications{}
	if err := c.SearchRead(AccountPeppolClarificationModel, criteria, options, apcs); err != nil {
		return nil, err
	}
	return apcs, nil
}

// FindAccountPeppolClarificationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolClarificationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPeppolClarificationModel, criteria, options)
}

// FindAccountPeppolClarificationId finds record id by querying it with criteria.
func (c *Client) FindAccountPeppolClarificationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPeppolClarificationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
