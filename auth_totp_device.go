package odoo

// AuthTotpDevice represents auth_totp.device model.
type AuthTotpDevice struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	ExpirationDate *Time     `xmlrpc:"expiration_date,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Scope          *String   `xmlrpc:"scope,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
}

// AuthTotpDevices represents array of auth_totp.device model.
type AuthTotpDevices []AuthTotpDevice

// AuthTotpDeviceModel is the odoo model name.
const AuthTotpDeviceModel = "auth_totp.device"

// Many2One convert AuthTotpDevice to *Many2One.
func (ad *AuthTotpDevice) Many2One() *Many2One {
	return NewMany2One(ad.Id.Get(), "")
}

// CreateAuthTotpDevice creates a new auth_totp.device model and returns its id.
func (c *Client) CreateAuthTotpDevice(ad *AuthTotpDevice) (int64, error) {
	ids, err := c.CreateAuthTotpDevices([]*AuthTotpDevice{ad})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthTotpDevice creates a new auth_totp.device model and returns its id.
func (c *Client) CreateAuthTotpDevices(ads []*AuthTotpDevice) ([]int64, error) {
	var vv []interface{}
	for _, v := range ads {
		vv = append(vv, v)
	}
	return c.Create(AuthTotpDeviceModel, vv, nil)
}

// UpdateAuthTotpDevice updates an existing auth_totp.device record.
func (c *Client) UpdateAuthTotpDevice(ad *AuthTotpDevice) error {
	return c.UpdateAuthTotpDevices([]int64{ad.Id.Get()}, ad)
}

// UpdateAuthTotpDevices updates existing auth_totp.device records.
// All records (represented by ids) will be updated by ad values.
func (c *Client) UpdateAuthTotpDevices(ids []int64, ad *AuthTotpDevice) error {
	return c.Update(AuthTotpDeviceModel, ids, ad, nil)
}

// DeleteAuthTotpDevice deletes an existing auth_totp.device record.
func (c *Client) DeleteAuthTotpDevice(id int64) error {
	return c.DeleteAuthTotpDevices([]int64{id})
}

// DeleteAuthTotpDevices deletes existing auth_totp.device records.
func (c *Client) DeleteAuthTotpDevices(ids []int64) error {
	return c.Delete(AuthTotpDeviceModel, ids)
}

// GetAuthTotpDevice gets auth_totp.device existing record.
func (c *Client) GetAuthTotpDevice(id int64) (*AuthTotpDevice, error) {
	ads, err := c.GetAuthTotpDevices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ads)[0]), nil
}

// GetAuthTotpDevices gets auth_totp.device existing records.
func (c *Client) GetAuthTotpDevices(ids []int64) (*AuthTotpDevices, error) {
	ads := &AuthTotpDevices{}
	if err := c.Read(AuthTotpDeviceModel, ids, nil, ads); err != nil {
		return nil, err
	}
	return ads, nil
}

// FindAuthTotpDevice finds auth_totp.device record by querying it with criteria.
func (c *Client) FindAuthTotpDevice(criteria *Criteria) (*AuthTotpDevice, error) {
	ads := &AuthTotpDevices{}
	if err := c.SearchRead(AuthTotpDeviceModel, criteria, NewOptions().Limit(1), ads); err != nil {
		return nil, err
	}
	return &((*ads)[0]), nil
}

// FindAuthTotpDevices finds auth_totp.device records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpDevices(criteria *Criteria, options *Options) (*AuthTotpDevices, error) {
	ads := &AuthTotpDevices{}
	if err := c.SearchRead(AuthTotpDeviceModel, criteria, options, ads); err != nil {
		return nil, err
	}
	return ads, nil
}

// FindAuthTotpDeviceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpDeviceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AuthTotpDeviceModel, criteria, options)
}

// FindAuthTotpDeviceId finds record id by querying it with criteria.
func (c *Client) FindAuthTotpDeviceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthTotpDeviceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
