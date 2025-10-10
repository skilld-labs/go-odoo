package odoo

// ResDevice represents res.device model.
type ResDevice struct {
	Browser           *String    `xmlrpc:"browser,omitempty"`
	City              *String    `xmlrpc:"city,omitempty"`
	Country           *String    `xmlrpc:"country,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DeviceType        *Selection `xmlrpc:"device_type,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	FirstActivity     *Time      `xmlrpc:"first_activity,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IpAddress         *String    `xmlrpc:"ip_address,omitempty"`
	IsCurrent         *Bool      `xmlrpc:"is_current,omitempty"`
	LastActivity      *Time      `xmlrpc:"last_activity,omitempty"`
	LinkedIpAddresses *String    `xmlrpc:"linked_ip_addresses,omitempty"`
	Platform          *String    `xmlrpc:"platform,omitempty"`
	Revoked           *Bool      `xmlrpc:"revoked,omitempty"`
	SessionIdentifier *String    `xmlrpc:"session_identifier,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResDevices represents array of res.device model.
type ResDevices []ResDevice

// ResDeviceModel is the odoo model name.
const ResDeviceModel = "res.device"

// Many2One convert ResDevice to *Many2One.
func (rd *ResDevice) Many2One() *Many2One {
	return NewMany2One(rd.Id.Get(), "")
}

// CreateResDevice creates a new res.device model and returns its id.
func (c *Client) CreateResDevice(rd *ResDevice) (int64, error) {
	ids, err := c.CreateResDevices([]*ResDevice{rd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResDevice creates a new res.device model and returns its id.
func (c *Client) CreateResDevices(rds []*ResDevice) ([]int64, error) {
	var vv []interface{}
	for _, v := range rds {
		vv = append(vv, v)
	}
	return c.Create(ResDeviceModel, vv, nil)
}

// UpdateResDevice updates an existing res.device record.
func (c *Client) UpdateResDevice(rd *ResDevice) error {
	return c.UpdateResDevices([]int64{rd.Id.Get()}, rd)
}

// UpdateResDevices updates existing res.device records.
// All records (represented by ids) will be updated by rd values.
func (c *Client) UpdateResDevices(ids []int64, rd *ResDevice) error {
	return c.Update(ResDeviceModel, ids, rd, nil)
}

// DeleteResDevice deletes an existing res.device record.
func (c *Client) DeleteResDevice(id int64) error {
	return c.DeleteResDevices([]int64{id})
}

// DeleteResDevices deletes existing res.device records.
func (c *Client) DeleteResDevices(ids []int64) error {
	return c.Delete(ResDeviceModel, ids)
}

// GetResDevice gets res.device existing record.
func (c *Client) GetResDevice(id int64) (*ResDevice, error) {
	rds, err := c.GetResDevices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rds)[0]), nil
}

// GetResDevices gets res.device existing records.
func (c *Client) GetResDevices(ids []int64) (*ResDevices, error) {
	rds := &ResDevices{}
	if err := c.Read(ResDeviceModel, ids, nil, rds); err != nil {
		return nil, err
	}
	return rds, nil
}

// FindResDevice finds res.device record by querying it with criteria.
func (c *Client) FindResDevice(criteria *Criteria) (*ResDevice, error) {
	rds := &ResDevices{}
	if err := c.SearchRead(ResDeviceModel, criteria, NewOptions().Limit(1), rds); err != nil {
		return nil, err
	}
	return &((*rds)[0]), nil
}

// FindResDevices finds res.device records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResDevices(criteria *Criteria, options *Options) (*ResDevices, error) {
	rds := &ResDevices{}
	if err := c.SearchRead(ResDeviceModel, criteria, options, rds); err != nil {
		return nil, err
	}
	return rds, nil
}

// FindResDeviceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResDeviceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResDeviceModel, criteria, options)
}

// FindResDeviceId finds record id by querying it with criteria.
func (c *Client) FindResDeviceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResDeviceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
