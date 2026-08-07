package odoo

// AccountEdiProxyClientUser represents account_edi_proxy_client.user model.
type AccountEdiProxyClientUser struct {
	Active            *Bool      `xmlrpc:"active,omitempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	EdiIdentification *String    `xmlrpc:"edi_identification,omitempty"`
	EdiMode           *Selection `xmlrpc:"edi_mode,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IdClient          *String    `xmlrpc:"id_client,omitempty"`
	IsTokenOutOfSync  *Bool      `xmlrpc:"is_token_out_of_sync,omitempty"`
	PrivateKeyId      *Many2One  `xmlrpc:"private_key_id,omitempty"`
	ProxyType         *Selection `xmlrpc:"proxy_type,omitempty"`
	RefreshToken      *String    `xmlrpc:"refresh_token,omitempty"`
	TokenSyncVersion  *Int       `xmlrpc:"token_sync_version,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountEdiProxyClientUsers represents array of account_edi_proxy_client.user model.
type AccountEdiProxyClientUsers []AccountEdiProxyClientUser

// AccountEdiProxyClientUserModel is the odoo model name.
const AccountEdiProxyClientUserModel = "account_edi_proxy_client.user"

// Many2One convert AccountEdiProxyClientUser to *Many2One.
func (au *AccountEdiProxyClientUser) Many2One() *Many2One {
	return NewMany2One(au.Id.Get(), "")
}

// CreateAccountEdiProxyClientUser creates a new account_edi_proxy_client.user model and returns its id.
func (c *Client) CreateAccountEdiProxyClientUser(au *AccountEdiProxyClientUser) (int64, error) {
	ids, err := c.CreateAccountEdiProxyClientUsers([]*AccountEdiProxyClientUser{au})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEdiProxyClientUser creates a new account_edi_proxy_client.user model and returns its id.
func (c *Client) CreateAccountEdiProxyClientUsers(aus []*AccountEdiProxyClientUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range aus {
		vv = append(vv, v)
	}
	return c.Create(AccountEdiProxyClientUserModel, vv, nil)
}

// UpdateAccountEdiProxyClientUser updates an existing account_edi_proxy_client.user record.
func (c *Client) UpdateAccountEdiProxyClientUser(au *AccountEdiProxyClientUser) error {
	return c.UpdateAccountEdiProxyClientUsers([]int64{au.Id.Get()}, au)
}

// UpdateAccountEdiProxyClientUsers updates existing account_edi_proxy_client.user records.
// All records (represented by ids) will be updated by au values.
func (c *Client) UpdateAccountEdiProxyClientUsers(ids []int64, au *AccountEdiProxyClientUser) error {
	return c.Update(AccountEdiProxyClientUserModel, ids, au, nil)
}

// DeleteAccountEdiProxyClientUser deletes an existing account_edi_proxy_client.user record.
func (c *Client) DeleteAccountEdiProxyClientUser(id int64) error {
	return c.DeleteAccountEdiProxyClientUsers([]int64{id})
}

// DeleteAccountEdiProxyClientUsers deletes existing account_edi_proxy_client.user records.
func (c *Client) DeleteAccountEdiProxyClientUsers(ids []int64) error {
	return c.Delete(AccountEdiProxyClientUserModel, ids)
}

// GetAccountEdiProxyClientUser gets account_edi_proxy_client.user existing record.
func (c *Client) GetAccountEdiProxyClientUser(id int64) (*AccountEdiProxyClientUser, error) {
	aus, err := c.GetAccountEdiProxyClientUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// GetAccountEdiProxyClientUsers gets account_edi_proxy_client.user existing records.
func (c *Client) GetAccountEdiProxyClientUsers(ids []int64) (*AccountEdiProxyClientUsers, error) {
	aus := &AccountEdiProxyClientUsers{}
	if err := c.Read(AccountEdiProxyClientUserModel, ids, nil, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindAccountEdiProxyClientUser finds account_edi_proxy_client.user record by querying it with criteria.
func (c *Client) FindAccountEdiProxyClientUser(criteria *Criteria) (*AccountEdiProxyClientUser, error) {
	aus := &AccountEdiProxyClientUsers{}
	if err := c.SearchRead(AccountEdiProxyClientUserModel, criteria, NewOptions().Limit(1), aus); err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// FindAccountEdiProxyClientUsers finds account_edi_proxy_client.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiProxyClientUsers(criteria *Criteria, options *Options) (*AccountEdiProxyClientUsers, error) {
	aus := &AccountEdiProxyClientUsers{}
	if err := c.SearchRead(AccountEdiProxyClientUserModel, criteria, options, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindAccountEdiProxyClientUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiProxyClientUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEdiProxyClientUserModel, criteria, options)
}

// FindAccountEdiProxyClientUserId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiProxyClientUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiProxyClientUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
