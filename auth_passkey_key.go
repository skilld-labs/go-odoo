package odoo

// AuthPasskeyKey represents auth.passkey.key model.
type AuthPasskeyKey struct {
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	CredentialIdentifier *String   `xmlrpc:"credential_identifier,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	PublicKey            *String   `xmlrpc:"public_key,omitempty"`
	SignCount            *Int      `xmlrpc:"sign_count,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AuthPasskeyKeys represents array of auth.passkey.key model.
type AuthPasskeyKeys []AuthPasskeyKey

// AuthPasskeyKeyModel is the odoo model name.
const AuthPasskeyKeyModel = "auth.passkey.key"

// Many2One convert AuthPasskeyKey to *Many2One.
func (apk *AuthPasskeyKey) Many2One() *Many2One {
	return NewMany2One(apk.Id.Get(), "")
}

// CreateAuthPasskeyKey creates a new auth.passkey.key model and returns its id.
func (c *Client) CreateAuthPasskeyKey(apk *AuthPasskeyKey) (int64, error) {
	ids, err := c.CreateAuthPasskeyKeys([]*AuthPasskeyKey{apk})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthPasskeyKey creates a new auth.passkey.key model and returns its id.
func (c *Client) CreateAuthPasskeyKeys(apks []*AuthPasskeyKey) ([]int64, error) {
	var vv []interface{}
	for _, v := range apks {
		vv = append(vv, v)
	}
	return c.Create(AuthPasskeyKeyModel, vv, nil)
}

// UpdateAuthPasskeyKey updates an existing auth.passkey.key record.
func (c *Client) UpdateAuthPasskeyKey(apk *AuthPasskeyKey) error {
	return c.UpdateAuthPasskeyKeys([]int64{apk.Id.Get()}, apk)
}

// UpdateAuthPasskeyKeys updates existing auth.passkey.key records.
// All records (represented by ids) will be updated by apk values.
func (c *Client) UpdateAuthPasskeyKeys(ids []int64, apk *AuthPasskeyKey) error {
	return c.Update(AuthPasskeyKeyModel, ids, apk, nil)
}

// DeleteAuthPasskeyKey deletes an existing auth.passkey.key record.
func (c *Client) DeleteAuthPasskeyKey(id int64) error {
	return c.DeleteAuthPasskeyKeys([]int64{id})
}

// DeleteAuthPasskeyKeys deletes existing auth.passkey.key records.
func (c *Client) DeleteAuthPasskeyKeys(ids []int64) error {
	return c.Delete(AuthPasskeyKeyModel, ids)
}

// GetAuthPasskeyKey gets auth.passkey.key existing record.
func (c *Client) GetAuthPasskeyKey(id int64) (*AuthPasskeyKey, error) {
	apks, err := c.GetAuthPasskeyKeys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apks)[0]), nil
}

// GetAuthPasskeyKeys gets auth.passkey.key existing records.
func (c *Client) GetAuthPasskeyKeys(ids []int64) (*AuthPasskeyKeys, error) {
	apks := &AuthPasskeyKeys{}
	if err := c.Read(AuthPasskeyKeyModel, ids, nil, apks); err != nil {
		return nil, err
	}
	return apks, nil
}

// FindAuthPasskeyKey finds auth.passkey.key record by querying it with criteria.
func (c *Client) FindAuthPasskeyKey(criteria *Criteria) (*AuthPasskeyKey, error) {
	apks := &AuthPasskeyKeys{}
	if err := c.SearchRead(AuthPasskeyKeyModel, criteria, NewOptions().Limit(1), apks); err != nil {
		return nil, err
	}
	return &((*apks)[0]), nil
}

// FindAuthPasskeyKeys finds auth.passkey.key records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthPasskeyKeys(criteria *Criteria, options *Options) (*AuthPasskeyKeys, error) {
	apks := &AuthPasskeyKeys{}
	if err := c.SearchRead(AuthPasskeyKeyModel, criteria, options, apks); err != nil {
		return nil, err
	}
	return apks, nil
}

// FindAuthPasskeyKeyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthPasskeyKeyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AuthPasskeyKeyModel, criteria, options)
}

// FindAuthPasskeyKeyId finds record id by querying it with criteria.
func (c *Client) FindAuthPasskeyKeyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthPasskeyKeyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
