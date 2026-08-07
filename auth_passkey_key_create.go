package odoo

// AuthPasskeyKeyCreate represents auth.passkey.key.create model.
type AuthPasskeyKeyCreate struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AuthPasskeyKeyCreates represents array of auth.passkey.key.create model.
type AuthPasskeyKeyCreates []AuthPasskeyKeyCreate

// AuthPasskeyKeyCreateModel is the odoo model name.
const AuthPasskeyKeyCreateModel = "auth.passkey.key.create"

// Many2One convert AuthPasskeyKeyCreate to *Many2One.
func (apkc *AuthPasskeyKeyCreate) Many2One() *Many2One {
	return NewMany2One(apkc.Id.Get(), "")
}

// CreateAuthPasskeyKeyCreate creates a new auth.passkey.key.create model and returns its id.
func (c *Client) CreateAuthPasskeyKeyCreate(apkc *AuthPasskeyKeyCreate) (int64, error) {
	ids, err := c.CreateAuthPasskeyKeyCreates([]*AuthPasskeyKeyCreate{apkc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthPasskeyKeyCreate creates a new auth.passkey.key.create model and returns its id.
func (c *Client) CreateAuthPasskeyKeyCreates(apkcs []*AuthPasskeyKeyCreate) ([]int64, error) {
	var vv []interface{}
	for _, v := range apkcs {
		vv = append(vv, v)
	}
	return c.Create(AuthPasskeyKeyCreateModel, vv, nil)
}

// UpdateAuthPasskeyKeyCreate updates an existing auth.passkey.key.create record.
func (c *Client) UpdateAuthPasskeyKeyCreate(apkc *AuthPasskeyKeyCreate) error {
	return c.UpdateAuthPasskeyKeyCreates([]int64{apkc.Id.Get()}, apkc)
}

// UpdateAuthPasskeyKeyCreates updates existing auth.passkey.key.create records.
// All records (represented by ids) will be updated by apkc values.
func (c *Client) UpdateAuthPasskeyKeyCreates(ids []int64, apkc *AuthPasskeyKeyCreate) error {
	return c.Update(AuthPasskeyKeyCreateModel, ids, apkc, nil)
}

// DeleteAuthPasskeyKeyCreate deletes an existing auth.passkey.key.create record.
func (c *Client) DeleteAuthPasskeyKeyCreate(id int64) error {
	return c.DeleteAuthPasskeyKeyCreates([]int64{id})
}

// DeleteAuthPasskeyKeyCreates deletes existing auth.passkey.key.create records.
func (c *Client) DeleteAuthPasskeyKeyCreates(ids []int64) error {
	return c.Delete(AuthPasskeyKeyCreateModel, ids)
}

// GetAuthPasskeyKeyCreate gets auth.passkey.key.create existing record.
func (c *Client) GetAuthPasskeyKeyCreate(id int64) (*AuthPasskeyKeyCreate, error) {
	apkcs, err := c.GetAuthPasskeyKeyCreates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apkcs)[0]), nil
}

// GetAuthPasskeyKeyCreates gets auth.passkey.key.create existing records.
func (c *Client) GetAuthPasskeyKeyCreates(ids []int64) (*AuthPasskeyKeyCreates, error) {
	apkcs := &AuthPasskeyKeyCreates{}
	if err := c.Read(AuthPasskeyKeyCreateModel, ids, nil, apkcs); err != nil {
		return nil, err
	}
	return apkcs, nil
}

// FindAuthPasskeyKeyCreate finds auth.passkey.key.create record by querying it with criteria.
func (c *Client) FindAuthPasskeyKeyCreate(criteria *Criteria) (*AuthPasskeyKeyCreate, error) {
	apkcs := &AuthPasskeyKeyCreates{}
	if err := c.SearchRead(AuthPasskeyKeyCreateModel, criteria, NewOptions().Limit(1), apkcs); err != nil {
		return nil, err
	}
	return &((*apkcs)[0]), nil
}

// FindAuthPasskeyKeyCreates finds auth.passkey.key.create records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthPasskeyKeyCreates(criteria *Criteria, options *Options) (*AuthPasskeyKeyCreates, error) {
	apkcs := &AuthPasskeyKeyCreates{}
	if err := c.SearchRead(AuthPasskeyKeyCreateModel, criteria, options, apkcs); err != nil {
		return nil, err
	}
	return apkcs, nil
}

// FindAuthPasskeyKeyCreateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthPasskeyKeyCreateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AuthPasskeyKeyCreateModel, criteria, options)
}

// FindAuthPasskeyKeyCreateId finds record id by querying it with criteria.
func (c *Client) FindAuthPasskeyKeyCreateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthPasskeyKeyCreateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
